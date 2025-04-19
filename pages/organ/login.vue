<script lang="ts" setup>
import { ref } from 'vue';
import { login, getOrgenVerifyCode } from '@/api/organ';

const logoUrl = '/static/logo.png';

const tabs = [
  { label: '账号密码登录', value: 'password' },
  { label: '手机验证码登录', value: 'verify' }
];

const currentTab = ref('password');
const account = ref('');
const password = ref('');
const phone = ref('');
const verifyCode = ref('');
const showPassword = ref(false);
const loading = ref(false);
const counting = ref(false);
const counter = ref(60);
const handleLogin = async () => {
  loading.value = true;
  
  try {
	let res;
	if (!account.value || !password.value) {
		uni.showToast({ title: '请输入账号密码', icon: 'none' });
		return;
	}
	res = await login({
		account_name: account.value,
		password: password.value
	});
	console.log(res)
	if (res.code === 0) {
	  uni.setStorageSync('Authorization', res.result);
	  uni.navigateTo({ url: '/pages/organ/warranty/add' });
	} else {
	  uni.showToast({ title: res.message, icon: 'none' });
	}
  } finally {
	loading.value = false;
  }
};
</script>

<template>
  <view class="main-content">
    <!-- Logo区域 -->
    <view class="logo">
      <image :src="logoUrl" mode="aspectFit" class="logo-img"></image>
      <text class="logo-title">傲莱菈机构登录</text>
    </view>

    <!-- 登录区域 -->
    <view class="login-container">
      <view class="form-container">
        <view class="input-container">
          <input 
            type="text"
            v-model="account"
            placeholder="请输入账号"
            class="input-field"
          />
          <uni-icons type="person" size="20" color="#999" class="input-icon"></uni-icons>
        </view>
        <view class="input-container">
          <input 
            :type="showPassword ? 'text' : 'password'"
            v-model="password"
            placeholder="请输入密码"
            class="input-field"
          />
          <uni-icons 
            :type="showPassword ? 'eye-filled' : 'eye-slash-filled'" 
            size="20" 
            color="#999" 
            class="input-icon"
            @click="showPassword = !showPassword"
          ></uni-icons>
        </view>
      </view>

      <button 
        @click="handleLogin"
        :loading="loading"
        class="login-btn"
      >
        <text v-if="loading">登录中...</text>
        <text v-else>登录</text>
      </button>
    </view>
  </view>
</template>

<style lang="scss">
.main-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 64rpx 20rpx;
  background-color: #fff;
  min-height: 100vh;
}

.logo {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 60rpx;
  
  &-img {
    width: 140rpx;
    height: 140rpx;
    margin-bottom: 40rpx;
    border-radius: 20rpx;
  }
  
  &-title {
    color: #666;
    font-size: 36rpx;
    font-weight: bold;
  }
}

.login-container {
  width: 100%;
  max-width: 600rpx;
}

.form-container {
  display: flex;
  flex-direction: column;
  gap: 40rpx;
}

.input-container {
  position: relative;
  
  .input-field {
    width: 100%;
    height: 96rpx;
    border: 2rpx solid #eee;
    border-radius: 16rpx;
    font-size: 28rpx;
    
    &:focus {
      border-color: #ff1b6b;
    }
  }
  
  .input-icon {
    position: absolute;
    right: 30rpx;
    top: 50%;
    transform: translateY(-50%);
  }
  
  &.code-input {
    .input-field {
      padding-right: 200rpx;
    }
    
    .code-btn {
      position: absolute;
      right: 10rpx;
      top: 50%;
      transform: translateY(-50%);
      height: 60rpx;
      line-height: 60rpx;
      padding: 0 20rpx;
      background: transparent;
      color: #ff1b6b;
      font-size: 24rpx;
      
      &[disabled] {
        color: #999;
      }
    }
  }
}

.login-btn {
  width: 100%;
  height: 96rpx;
  margin-top: 60rpx;
  background-image: linear-gradient(to right, #ff1b6b, #ff9147);
  color: #fff;
  border-radius: 16rpx;
  font-size: 32rpx;
  
  &:active {
    opacity: 0.9;
  }
}
</style>