<template>
    <div class="font-size-settings">
      <div class="font-view">
      <div class="set_view">
      <h3>字体大小</h3>
        <div class="preview">
            <div class="message-box">
            <p :style="{ fontSize: `${this.fontSize}px` ,fontFamily:fontStyle }">预览字体大小</p>
            </div>
            <div class="avatar">
                <img :src="avatar">
            </div>
        </div>
        <div class="slider-container">
          <label >字体大小:</label>
            <input
            type="range"
            min="12"
            max="28"
            v-model="fontSize"
            @input="onFontSizeChange"
            />
            <span>{{ fontSize+'px' }}</span>
            <img src="@/assets/icon/comfirm.png" alt="Contact" class="icon" @click="saveFontSize"/>
        </div>
        <div class="description">
            <p>拖动上面的滑块，可设置字体大小</p>
        </div>

              <!-- 字体样式选择 -->
        <div class="font-style-container">
          <label for="fontStyle">字体样式:</label>
          <select ref="fontStyleSelect" v-model="fontStyle" @change="onFontStyleChange">
            <option value="Microsoft YaHei">微软雅黑</option>
            <option value="SimSun">宋体</option>
            <option value="SimHei">黑体</option>
            <option value="Kaiti">楷体</option>
            <option value="FangSong">仿宋</option>
            <option value="LiSu">隶书</option>
            <option value="YouYuan">幼圆</option>
            <option value="STSong">华文宋体</option>
            <option value="STXihei">华文细黑</option>
            <option value="STKaiti">华文楷体</option>
          </select>
          <img src="@/assets/icon/comfirm.png" alt="Contact" class="icon" @click="saveFontStyle"/>
        </div>
      </div>
      </div>

      <Windows 
        :visible="showModal"
        :message="modalMessage"
        @close="showModal = false"
      />
    </div>
  </template>

  
  <script>
  import Windows from '@/components/base/Windows.vue';
  import avatar from '@/assets/images/avatar.jpg';
  import { changeFontsize,changeFontstyle } from '@/services/settingView.js';

  export default {
    components:{
      Windows,
    },
    data() {
      return {
        avatar:this.$store.state.user.avatar,
        fontSize:this.$store.state.settings.fontSize, // 默认字体大小
        fontStyle:this.$store.state.settings.fontStyle,
        showModal:false,
        modalMessage:'',
      };
    },
    mounted() {
    // 在组件挂载时，设置下拉选择框的默认选项
    this.setDefaultFontStyle();
    },
    methods: {
      onFontSizeChange() {
        // 处理字体大小变化逻辑
        console.log('Font size changed:', this.fontSize);
      },
      onFontStyleChange() {
      // 处理字体样式变化逻辑
      console.log('Font style changed:', this.fontStyle);
    },
      async saveFontSize() {
        try{
          //向后端发送更改请求
          const response = await changeFontsize({fontSize:this.fontSize});
          if(response.success){
            // 保存字体大小逻辑
            console.log('Font size saved:', this.fontSize);
            // 调用一个方法来保存字体大小
            this.$store.commit('SET_FONTSIZE',`${this.fontSize}px`);
            this.$emit('updateUser', {fontsize:`${this.fontSize}px`});
            this.modalMessage = '字体大小修改成功';
            this.showModal = true;
          }else{
            this.modalMessage = response.message;
            this.showModal =true;
          }
        }catch(error){
          this.modalMessage = '保存失败请重试！';
          this.showModal = true;
          console.error(error);
        }
      },

      async saveFontStyle(){
        try{
          //向后端发送更改请求
          const response = await changeFontstyle({fontStyle:this.fontStyle});
          if(response.success){
            // 保存字体大小逻辑
            console.log('Font style saved:', this.fontStyle);
            // 调用一个方法来保存字体大小
            this.$store.commit('SET_FONTSTYLE',this.fontStyle);
            this.modalMessage = '字体样式修改成功';
            this.showModal = true;
          }else{
            this.modalMessage = response.message;
            this.showModal =true;
          }
        }catch(error){
          this.modalMessage = '样式保存失败请重试！';
          this.showModal = true;
          console.error(error);
        }
      },

      setDefaultFontStyle() {
      // 设置下拉选择框的默认选项
      const selectElement = this.$refs.fontStyleSelect;
      if (selectElement) {
        const options = selectElement.options;
        for (let i = 0; i < options.length; i++) {
          if (options[i].value === this.$store.state.settings.fontStyle) {
            selectElement.selectedIndex = i;
            break;
          }
        }
      }
    },
    },
  };
  </script>
  
  <style scoped>
  .font-size-settings {
    margin-top: 60px;
    padding: 20px;
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .font-view {
  padding: 20px;
  width: 100%;
  height: 80%;
  max-width: 500px; /* 设置最大宽度 */
  background-color: var(--background-color); /* 添加背景色 */
  border-radius: 8px; /* 添加圆角 */
  box-shadow: 0 2px 10px var(--background-color2); /* 添加阴影 */
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  gap:10px
  }

.set_view{
  display: flex;
  flex-direction: column;
  width: 350px;
  height: 400px;
  justify-content: space-around;
  align-items: center;
  gap:5px;
}
  
  h3 {
    font-size: var(--font-size-large);
    margin-top: 10px;
    margin-bottom: 15px;
    color: var(--text-color);
    background-color: var(--sidebar-background-color1);
    width: 350px;
    height: 33px;
    border-radius: 10px; /* 添加圆角 */
    box-shadow: 0 2px 10px var(--background-color2); /* 添加阴影*/ 
  }
  
  .preview {
    margin-bottom: 20px;
    text-align: center;
    align-items: flex-end;
    display: flex;
    margin-left: 80px;
  }

  .message-box {
  background-color: var(--sidebar-background-color); /* 浅蓝色背景 */
  border-radius: 8px; /* 圆角 */
  padding: 10px;
  margin: 0 auto; /* 居中 */
  max-width: 100%; /*最大宽度 */
  margin-right: 10px;
}

.message-box p {
  font-size: var(--font-size); /* 默认字体大小 */
  color: var(--sidebar-text-color);
  margin: 0;
}

.avatar {
  width: 50px;
  height: 50px;
  border-radius: 100%; /* 圆形头像 */
  overflow: hidden; /* 隐藏超出部分 */
}

.avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover; /* 保持图片比例 */
}
  
  .preview p {
    font-size: var(--font-size);
    color: var(--text-color);
  }
  
  .slider-container {
    font-size: var(--font-size);
    color: var(--text-color);
    display: flex;
    align-items: center;
    width: 350px;
  }
  
  input[type="range"] {
    flex: 1;
  }

  span{
    margin-right: 10px;
  }
  
  .description {
    margin-bottom: 20px;
  }

  select{
    flex: 1;
  }

  .font-style-container {
    font-size: var(--font-size);
    color:var(--text-color);
    width: fit-content;
    display: flex;
    align-items: center;
    margin-bottom: 20px;
    width: 350px;
    gap:5px;
  }

  .font-style-container label {
    margin-right: 10px;
  }

  .font-style-container select {
    padding: 5px;
    border-radius: 4px;
    border: 1px solid #ccc;
  }

  .icon{
    width:30px;
    height: 30px;
    margin-right: 10px;
  }

  .icon:hover{
    background-color: var(--button-background-color);
    border-radius: 15px;
  }

  label{
    margin-right: 10px;
  }
  </style>