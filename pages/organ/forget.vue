<template>
    <Layout>
        <view class="page-container">
            <view class="header">
                <view class="header-content">
                    <text class="header-title">修改密码</text>
                </view>
            </view>

            <view class="form-container">
                <form @submit="submitForm">
                    <view class="form-item">
                        <view class="form-label">账号</view>
                        <input class="form-input" v-model="formData.account_name" readonly placeholder="请输入账号" />
                    </view>
                    <view class="form-item">
                        <view class="form-label">新密码</view>
                        <input class="form-input" v-model="formData.new_password" placeholder="请输入新密码" password />
                    </view>
                    <view class="form-item">
                        <view class="form-label">确认密码</view>
                        <input class="form-input" v-model="formData.confirm_password" placeholder="请输入确认密码" password />
                    </view>
                    <view class="submit-btn-container">
                        <button 
                            class="submit-btn"
                            @click="submitForm"
                            form-type="submit"
                        >
                            提交
                        </button>
                    </view>
                </form>
            </view>
        </view>
    </Layout>
</template>

<script lang="ts" setup>
import Layout from './layout.vue';
import { forgetPassword, getOraganInfo } from '@/api/organ';
import { ref, onMounted } from 'vue';

const formData = ref({
    account_name: '',
    new_password: '',
    confirm_password: ''
});

const fetchAccountName = async () => {
    const res = await getOraganInfo();
    formData.value.account_name = res.result.account_name;
};

onMounted(() => {
    fetchAccountName();
});

const submitForm = async () => {
    console.log("formData", formData.value);
    if (formData.value.new_password !== formData.value.confirm_password) {
        uni.showToast({
            title: '两次密码不一致',
            icon: 'none'
        });
        return;
    }
    const data = {
        account_name: formData.value.account_name,
        password: formData.value.new_password
    };
    const res = await forgetPassword(data);
    if (res.code === 0) {
        uni.showToast({
            title: '修改成功',
            icon: 'success'
        });
        setTimeout(() => {
            uni.navigateTo({
              url: '/pages/organ/login'
            });
        }, 1000);
    } else {
        uni.showToast({
            title: res.message,
            icon: 'none'
        });
    }
};
</script>

<style scoped>
/* 全局样式重置 */
::-webkit-inner-spin-button,
::-webkit-outer-spin-button {
    -webkit-appearance: none;
    margin: 0;
}

/* 页面容器 */
.page-container {
    min-height: 100vh;
    background-color: #ffffff;
    padding-bottom: 100rpx;
    display: flex;
    flex-direction: column;
}

/* 头部样式 */
.header {
    height: 80rpx;
    position: relative;
}

.header-content {
    position: fixed;
    left: 0;
    height: 80rpx;
    background: linear-gradient(to right, #ff1b6b, #ff9147);
    color: #ffffff;
    display: flex;
    align-items: center;
	justify-content: center;
	width: 100%;
}

.header-icon {
    font-size: 36rpx;
}

.header-title {
    font-size: 36rpx;
    font-weight: 600;
}

/* 表单容器 */
.form-container {
    padding: 30rpx;
}

/* 表单项 */
.form-item {
    display: flex;
    align-items: center;
    padding: 20rpx 0;
    border-bottom: 1rpx solid #eee;
}

.form-label {
    width: 150rpx;
    font-size: 28rpx;
    color: #333;
}

.form-input {
    flex: 1;
    font-size: 28rpx;
    padding: 10rpx 0;
}

/* 提交按钮 */
.submit-btn-container {
    margin-top: 40rpx;
    padding: 0 30rpx;
}

.submit-btn {
    background: linear-gradient(to right, #ff1b6b, #ff9147);
    color: white;
    width: 100%;
    border: none;
    border-radius: 50rpx;
    height: 80rpx;
    line-height: 80rpx;
    font-size: 32rpx;
}
</style>