
<template>
    <Layout>
	<view class="min-h-screen bg-white pb-20 flex flex-col">
        <view class="h-16">
            <view class="fixed top-0 w-full z-50 bg-gradient-to-r from-[#ff1b6b] to-[#ff9147] text-white py-4 px-4">
                <view class="flex items-center">
                    <text class="fas fa-arrow-left text-xl" @click="uni.navigateBack()"></text>
                    <text class="ml-4 text-xl font-semibold">修改密码</text>
                </view>
            </view>
        </view>

        <view class="container">
            <form @submit="submitForm">
                <view class="uni-form-item">
                    <view class="uni-form-label">账号</view>
                    <input v-model="formData.account_name" readonly placeholder="请输入账号" />
                </view>
                <view class="uni-form-item">
                    <view class="uni-form-label">新密码</view>
                    <input v-model="formData.new_password" placeholder="请输入新密码" password />
                </view>
                <view class="uni-form-item">
                    <view class="uni-form-label">确认密码</view>
                    <input v-model="formData.confirm_password" placeholder="请输入确认密码" password />
                </view>
                <view class="mt-4 px-4">
                    <button 
                        :style="{ background: 'linear-gradient(to right, #ff1b6b, #ff9147)', color: 'white', width: '100%', border: 'none' }"
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
::-webkit-inner-spin-button,
::-webkit-outer-spin-button {
    -webkit-appearance: none;
    margin: 0;
}
.min-h-screen {
    min-height: calc(100vh - 64px); /* 减去底部导航高度 */
}
.uni-form-item {
    display: flex;
    align-items: center;
    padding: 10px 0;
}
.uni-form-label {
    width: 80px;
}

</style>
    