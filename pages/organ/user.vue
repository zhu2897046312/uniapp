<template>
	<Layout>
  <view class="page-container">
    <!-- 顶部标题栏 -->
    <view class="header">
      <text class="title">我的</text>
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

      <!-- 用户信息列表 -->
      <uni-list>
        <uni-list-item title="机构名称" :rightText="userInfo?.institution_name || ''" />
        <uni-list-item title="联系人" :rightText="userInfo?.contact_person || ''" />
        <uni-list-item title="手机号码" :rightText="userInfo?.mobile || ''" />
      </uni-list>

      <!-- 操作列表 -->
      <uni-list class="action-list">
        <uni-list-item title="修改密码" showArrow @click="forget" />
        <uni-list-item title="退出登录" class="logout-item" @click="logout" />
      </uni-list>
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
.page-container {
  min-height: 100vh;
  background-color: #f8f8f8;
  display: flex;
  flex-direction: column;
}

.header {
  background: linear-gradient(to right, #ff1b6b, #ff9147);
  color: white;
  padding: 30rpx;
}

.title {
  font-size: 36rpx;
  font-weight: bold;
  margin-left: 20rpx;
}

.content {
  flex: 1;
  padding: 20rpx;
}

.user-card {
  background-color: #fff;
  border-radius: 16rpx;
  padding: 30rpx;
  margin-bottom: 30rpx;
  display: flex;
  align-items: center;
  box-shadow: 0 2rpx 10rpx rgba(0, 0, 0, 0.05);
}

.avatar {
  width: 120rpx;
  height: 120rpx;
  border-radius: 50%;
  margin-right: 30rpx;
}

.user-info {
  display: flex;
  flex-direction: column;
}

.institution-name {
  font-size: 32rpx;
  font-weight: bold;
  margin-bottom: 10rpx;
}

.account-name {
  font-size: 28rpx;
  color: #666;
}

.action-list {
  margin-top: 40rpx;
}

.logout-item {
  color: #ff1b6b;
}

/* 底部安全区域适配 */
.safe-area-inset-bottom {
  padding-bottom: env(safe-area-inset-bottom);
}
</style>