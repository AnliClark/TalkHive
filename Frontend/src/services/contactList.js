import apiClient from '@/services/api';

// 搜索好友/群聊（key可能是id或者昵称）
export const searchStrangers = (key) => {
  return apiClient.post('/stranger/search', { key });  
};
// 搜索通讯录
export const searchContacts = (keyword) => {
  return apiClient.post('/contactList/search', { keyword });
};

// 获取好友请求列表接口
export const getFriendRequests = () => {
  return apiClient.get('/contactList/friendRequests');
};
// 处理好友请求接口
export const friendRequestPend = (account_id, accept) => {
  return apiClient.post('/contactList/friendRequests/pend', { account_id, accept});
};
// 添加好友
export const addFriend = (account_id, reason) => {
  return apiClient.post('/contactList/friendRequests/addFriend', { account_id, reason });
}


// 获取群聊请求列表接口
export const getGroupRequests = () => {
  return apiClient.get('/contactList/groupRequests');
};

// 处理群聊邀请请求接口
export const groupInvitationRequestPend = (account_id, group_id, accept) => {
  return apiClient.post('/contactList/groupRequests/invitationPend', { account_id, group_id, accept });
};
// 发送群聊
export const addGroup = (group_id, reason) => {
  return apiClient.post('/contactList/groupRequests/addGroup', { group_id, reason });
}

// 处理群聊申请请求接口
export const groupApplyRequestPend = (account_id, group_id, accept) => {
  return apiClient.post('/contactList/groupRequests/applyPend', { account_id, group_id, accept });
};

// 好友和群聊
export const changeRemark = (id, is_group_chat, remark)=>{
  return apiClient.post('/contactList/remark', { id, is_group_chat ,remark });
}

// *好友列表
// 获取好友列表接口
export const getFriends = () => {
  return apiClient.get('/contactList/friends');
};
// 删除好友
export const deleteFriend = (account_id) => {
  return apiClient.post('/contactList/friends/delete', { account_id });
};

// *分组
// 获取分组名
export const getDivides = (type) => {   // type: 'friends' or 'groups'
  return apiClient.get(`/contactList/${type}/divides`);
};
// 新建分组
export const createDivide = (type, divide) => {   // type: 'friends' or 'groups'
  return apiClient.post(`/contactList/${type}/divides/create`, { divide });
};
// 删除分组
export const deleteDivide = (type, divide) => {   // type: 'friends' or 'groups'
  return apiClient.post(`/contactList/${type}/divides/delete`,{ divide });
};
// 修改分组名称
export const renameDivide = (type, old_divide, new_divide) => {   // type: 'friends' or 'groups'
  return apiClient.post(`/contactList/${type}/divides/rename`, {old_divide, new_divide });
};
// 移动好友到分组
export const moveInDivide = (type, tid, divide) => {
  return apiClient.post(`/contactList/${type}/divides/moveIn`, { tid, divide });
};


// 获取黑名单列表接口
export const getBlackList = () => {
  return apiClient.get('/contactList/blackList');
};

// 从黑名单中移除用户接口
export const removeFromBlackList = (account_id) => {
  return apiClient.post('/contactList/blackList/remove', { account_id });
};
// 拉黑用户
export const addToBlackList = (account_id) => {
  return apiClient.post('/contactList/blackList/add', { account_id });
};


// * 群聊
// 获取群聊列表
export const getGroups = () => {
  return apiClient.get('/contactList/groups');
};
// 创建群聊
export const createGroup = (group_name,group_avatar,group_description,allow_invite,allow_id_search,allow_name_search) => {
  return apiClient.post('/contactList/groups/createGroup', { group_name,group_avatar,group_description,allow_invite,allow_id_search,allow_name_search });
};
// 解散群聊接口
export const dismissGroup = (group_id) => {
  return apiClient.post(`/contactList/groups/dismissGroup`, { group_id });
};
// 搜索不在群聊内的好友
export const fetchFriendsNotInGroup=(group_id)=>{
  return apiClient.post('/contactList/groups/friendsNotInGroup', {group_id});
}
// 邀请他人入群
export const inviteMember=(group_id, account_id)=>{
  return apiClient.post('/contactList/groups/invite', { account_id,group_id});
}
// 退出群聊
export const exitGroup = (group_id) => {
  return apiClient.post(`/contactList/groups/exit`, { group_id });
};
// 获取群聊详细信息
export const getGroupInfo = (group_id) => {
  return apiClient.get(`/contactList/groups/groupInfo/${group_id}`);
};
// 更改我在群聊内的昵称
export const changeGroupNickname=(group_id, group_nickname)=>{
  return apiClient.post(`/contactList/groups/changeNickname`, {group_id, group_nickname});
}
// // 搜索群成员
// export const searchGroupMember=(group_id, keyword)=>{
//   return apiClient.post('/contactList/groups/searchMember', {group_id, keyword});
// }

// 设置是否允许群成员邀请他人
export const setAllowInvite=(group_id, allow_invite)=>{
  return apiClient.post('/contactList/groups/setAllowInvite', {group_id, allow_invite});
}
// 设置是否允许通过ID搜索群
export const setAllowIdSearch=(group_id, allow_id_search)=>{
  return apiClient.post('/contactList/groups/setAllowIdSearch', {group_id, allow_id_search});
}
// 设置是否允许通过昵称搜索群
export const setAllowNameSearch=(group_id, allow_name_search)=>{
  return apiClient.post('/contactList/groups/setAllowNameSearch', {group_id, allow_name_search});
}
// 是否全体禁言
export const setAllBanned=(group_id, is_all_banned)=>{
  return apiClient.post('/contactList/groups/setAllBanned', {group_id, is_all_banned});
}
// 禁言某人
export const setBanned=(group_id, account_id, is_banned)=>{
  return apiClient.post('/contactList/groups/banMember',{group_id, account_id, is_banned});
}
// 移除某人
export const removeMember=(group_id, account_id)=>{
  return apiClient.post('/contactList/groups/removeMember', {group_id, account_id});
}
// 设为管理员
export const setAdmin=(group_id, account_id, is_admin)=>{
  return apiClient.post('/contactList/groups/setAdmin', {group_id, account_id, is_admin});
}
// 转让群主
export const transferOwner=(group_id, account_id)=>{
  return apiClient.post('/contactList/groups/transferOwner', {group_id, account_id});
}
// 更改群头像
export const changeGroupAvatar=(group_id, group_avatar)=>{
  return apiClient.post('/contactList/groups/changeAvatar', {group_id, group_avatar});
}
// 更改群名称
export const changeGroupName=(group_id, group_name)=>{
  return apiClient.post('/contactList/groups/changeGroupName', {group_id, group_name});
}
// 更改群介绍
export const changeGroupIntroduction=(group_id, group_description)=>{
  return apiClient.post('/contactList/groups/changeGroupIntroduction', {group_id, group_description});
} 