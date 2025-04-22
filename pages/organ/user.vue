<template>
  <Layout>
    <view class="page-container">
      <!-- 顶部标题栏 -->
      <view class="header">
        <view class="flex items-center">
          <text class="ml-4 text-xl font-semibold">我的</text>
        </view>
      </view>

      <view class="content">
        <!-- 用户信息卡片 -->
        <view class="user-card">
          <image 
            :src="getDefaultAvatar(userInfo?.account_name || '')" 
            class="avatar"
            mode="aspectFill"
          />
          <view class="user-info">
            <text class="institution-name">{{ userInfo?.institution_name }}</text>
            <text class="account-name">{{ userInfo?.account_name }}</text>
          </view>
        </view>
        
        <!-- 联系信息卡片 - 美化后的样式 -->
        <view class="contact-info">
          <view class="info-item">
            <text class="info-label">机构名称</text>
            <text class="info-value">{{ userInfo?.institution_name || '未设置' }}</text>
          </view>
          <view class="info-divider"></view>
          <view class="info-item">
            <text class="info-label">联系人</text>
            <text class="info-value">{{ userInfo?.contact_person || '未设置' }}</text>
          </view>
          <view class="info-divider"></view>
          <view class="info-item">
            <text class="info-label">手机号码</text>
            <text class="info-value">{{ userInfo?.mobile || '未设置' }}</text>
          </view>
        </view>
        
        <!-- 操作列表 - 美化后的样式 -->
        <view class="action-list">
          <view class="action-item" @click="forget">
            <text>修改密码</text>
            <uni-icons type="arrowright" size="16" color="#999"></uni-icons>
          </view>
          <view class="action-divider"></view>
          <view class="action-item logout-item" @click="logout">
            <text>退出登录</text>
            <uni-icons type="arrowright" size="16" color="#999"></uni-icons>
          </view>
        </view>
      </view>
    </view>
  </Layout>
</template>

<script lang="ts" setup>
import Layout from './layout.vue';
import { onMounted, ref } from 'vue';
import { getOraganInfo } from '@/api/organ';

interface UserInfo {
  institution_name: string,
  account_name: string,
  mobile: string,
  contact_person: string,
  state: number
}

const userInfo = ref<UserInfo | null>(null);

// 生成默认头像
const getDefaultAvatar = (name: string) => {
  if (!name) return '';
  return `https://ui-avatars.com/api/?name=${encodeURIComponent(name)}&background=random&rounded=true`;
};

// 退出登录
const logout = () => {
  uni.showModal({
    title: '提示',
    content: '确定退出登录吗？',
    confirmText: '退出',
    cancelText: '取消',
    success: (res) => {
      if (res.confirm) {
        uni.removeStorageSync('Authorization');
        uni.removeStorageSync('account_name');
        uni.reLaunch({
          url: '/pages/organ/login'
        });
      }
    }
  });
};

// 修改密码
const forget = () => {
  uni.navigateTo({
    url: '/pages/organ/forget'
  });
};

// 获取用户信息
const fetchInfo = async () => {
  uni.showLoading({
    title: '加载中...',
    mask: true
  });
  
  try {
    const res = await getOraganInfo();
    userInfo.value = res.result;
  } catch (error) {
    uni.showToast({
      title: '加载失败，请重试',
      icon: 'none'
    });
    console.error('获取用户信息失败:', error);
  } finally {
    uni.hideLoading();
  }
};

onMounted(() => {
  fetchInfo();
});
</script>

<style scoped>
/* 保持与 myPage.vue 相同的整体布局 */
.page-container {
  min-height: 100vh;
  background-color: #f8f8f8;
  display: flex;
  flex-direction: column;
}

/* 顶部标题栏 - 完全匹配 myPage.vue 的渐变效果 */
.header {
  width: 100%;
  background: linear-gradient(to right, #ff1b6b, #ff9147);
  color: white;
  padding: 16px;
}

/* 内容区域 - 匹配 padding 和布局 */
.content {
  width: 100%;
  flex: 1;
}

/* 用户卡片 - 匹配 myPage.vue 的样式 */
.user-card {
  background-color: #fff;
  border-radius: 8px;
  margin-bottom: 16px;
  display: flex;
  align-items: center;
  box-shadow: 0 1px 5px rgba(0, 0, 0, 0.05);
}

.avatar {
  width: 64px;
  height: 64px;
  border-radius: 50%;
  margin-right: 16px;
}

.user-info {
  display: flex;
  flex-direction: column;
}

.institution-name {
  font-size: 18px;
  font-weight: 600;
  margin-bottom: 4px;
}

.account-name {
  font-size: 14px;
  color: #666;
}

.info-list {
  margin-bottom: 16px;
}
.contact-info {
  background-color: white;
  border-radius: 8px;
  padding: 0 16px;
  margin-bottom: 16px;
  box-shadow: 0 1px 5px rgba(0, 0, 0, 0.05);
}

.info-item {
  display: flex;
  justify-content: space-between;
  padding: 16px 0;
}

.info-label {
  font-size: 14px;
  color: #666;
}

.info-value {
  font-size: 14px;
  color: #333;
  font-weight: 500;
}

.info-divider {
  height: 1px;
  background-color: #f0f0f0;
  margin: 0 -16px;
}

/* 操作列表美化 */
.action-list {
  background-color: white;
  border-radius: 8px;
  padding: 0 16px;
  box-shadow: 0 1px 5px rgba(0, 0, 0, 0.05);
}

.action-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 0;
  font-size: 14px;
}

.action-item.logout-item {
  color: #ff1b6b;
}

.action-divider {
  height: 1px;
  background-color: #f0f0f0;
  margin: 0 -16px;
}
</style>