package workSpace

import (
	"TalkHive/global"
	"TalkHive/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 回收站！！！

// GetTrashItems - 返回回收站列表√
func GetTrashItems(c *gin.Context) {
	// 获取用户 ID 参数
	userID := c.GetHeader("User-Id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	// 1. 查询回收站中属于该用户的内容，按时间降序排序
	var trashItems []models.Recycle
	if err := global.Db.Where("account_id = ?", global.ParseUint(userID)).
		Order("recycle_time DESC"). // 按时间降序排序
		Find(&trashItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch trash items"})
		return
	}

	// 2. 获取回收站文件名并格式化返回结果
	var result []map[string]interface{}
	for _, item := range trashItems {
		filename := ""

		// 根据 RecycleType 查询对应的文件名
		if item.RecycleType == "note" {
			var note models.Notes
			if err := global.Db.Where("note_id = ?", item.RecycleID).First(&note).Error; err == nil {
				filename = note.NoteName + ".md"
			}
		} else if item.RecycleType == "code" {
			var code models.Codes
			if err := global.Db.Where("code_id = ?", item.RecycleID).First(&code).Error; err == nil {
				filename = code.Name + code.Suffix
			}
		}

		// 构造结果
		result = append(result, map[string]interface{}{
			"recycle_id":   item.RecycleID,
			"type":         item.RecycleType,
			"recycle_time": item.RecycleTime.Format("2006-01-02 15:04"), // 格式化时间
			"filename":     filename,                                    // 文件名
		})
	}

	// 3. 返回 JSON 响应
	c.JSON(http.StatusOK, result)
}

// RestoreItem - 恢复回收站笔记/代码√
func RestoreItem(c *gin.Context) {
	// 解析前端传递的参数
	var request struct {
		Type      string `json:"type" binding:"required,oneof=code note"`
		RecycleID uint   `json:"recycle_id" binding:"required"`
	}

	// 绑定参数
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 获取用户 ID 参数
	userID := c.GetHeader("User-Id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	// 1.获取回收站中的数据
	var recycle models.Recycle
	if err := global.Db.Model(&models.Recycle{}).Where("recycle_id = ? AND recycle_type = ? AND account_id = ?",
		request.RecycleID, request.Type, global.ParseUint(userID)).First(&recycle).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recycle item not found"})
		return
	}

	// 2.根据type恢复对应的笔记或代码
	var err error
	switch request.Type {
	case "note":
		// 恢复笔记
		err = global.Db.Model(&models.Notes{}).Where("note_id = ? AND account_id = ?", request.RecycleID,
			global.ParseUint(userID)).Update("is_show", true).Error
	case "code":
		// 恢复代码
		err = global.Db.Model(&models.Codes{}).Where("code_id = ? AND account_id = ?", request.RecycleID,
			global.ParseUint(userID)).Update("is_show", true).Error
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid type"})
		return
	}

	// 处理恢复失败的情况
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to restore item"})
		return
	}

	// 3.删除回收站中的记录
	if err := global.Db.Model(&models.Recycle{}).Where("recycle_id = ?",
		request.RecycleID).Delete(&models.Recycle{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete recycle item"})
		return
	}

	// 4.返回成功消息
	c.JSON(http.StatusOK, gin.H{"message": "Item restored successfully"})
}

// DeletePermanently - 永久删除回收站中的文件√
func DeletePermanently(c *gin.Context) {
	// 解析前端传递的参数
	var request struct {
		Type      string `json:"type" binding:"required,oneof=code note"`
		RecycleID uint   `json:"recycle_id" binding:"required"`
	}

	// 绑定参数
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 获取用户 ID 参数
	//userID := c.Param("id")
	userID := c.GetHeader("User-Id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	// 1. 获取回收站中的数据
	var recycle models.Recycle
	if err := global.Db.Model(&models.Recycle{}).Where("recycle_id = ? AND recycle_type = ? AND account_id = ?",
		request.RecycleID, request.Type, global.ParseUint(userID)).First(&recycle).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recycle item not found"})
		return
	}

	// 2. 根据type永久删除笔记或代码
	var err error
	switch request.Type {
	case "note":
		// 永久删除笔记
		err = global.Db.Model(&models.Notes{}).Where("note_id = ? AND account_id = ?", request.RecycleID,
			global.ParseUint(userID)).Delete(&models.Notes{}).Error
	case "code":
		// 永久删除代码
		err = global.Db.Model(&models.Codes{}).Where("code_id = ? AND account_id = ?", request.RecycleID,
			global.ParseUint(userID)).Delete(&models.Codes{}).Error
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid type"})
		return
	}

	// 删除失败处理
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item permanently"})
		return
	}

	// 3. 删除回收站中的记录
	if err := global.Db.Model(&models.Recycle{}).Where("recycle_id = ?",
		request.RecycleID).Delete(&models.Recycle{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete recycle item"})
		return
	}

	// 4. 返回成功消息
	c.JSON(http.StatusOK, gin.H{"message": "Item deleted permanently"})
}
