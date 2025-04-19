<template>
  <Layout>
    <view class="page-container">
      <!-- 顶部标题栏 -->
      <view class="header">
        <text class="title">质保卡录入</text>
      </view>

      <!-- 表单内容 -->
      <view class="form-container">
        <form @submit="submitForm">
          <!-- 表单项目 -->
          <view class="form-item">
            <text class="form-label">质保卡号</text>
            <input 
              v-model="formData.card_number" 
              placeholder="请输入质保卡号" 
              class="form-input"
            />
          </view>
          
          <view class="form-item">
            <text class="form-label">产品序列号</text>
            <input 
              v-model="formData.serial_number" 
              placeholder="请输入产品序列号" 
              class="form-input"
            />
          </view>
		  <view class="form-item">
		    <text class="form-label">姓名</text>
		    <input 
		      v-model="formData.patient_name" 
		      placeholder="请输入姓名" 
		      class="form-input"
		    />
		  </view>
		  <view class="form-item">
		    <text class="form-label">年龄</text>
		    <input 
		      v-model="formData.patient_age" 
		      placeholder="请输入年龄" 
		      class="form-input"
		    />
		  </view>
		  <view class="form-item">
		    <text class="form-label">手术医生</text>
		    <input 
		      v-model="formData.surgeon_name" 
		      placeholder="请输入手术医生" 
		      class="form-input"
		    />
		  </view>
          
          <!-- 其他表单字段... -->
          
          <view class="form-item">
            <text class="form-label">手术日期</text>
            <picker 
              mode="date" 
              :value="formData.surgery_date" 
              @change="onDateChange"
              class="date-picker"
            >
              <view class="picker-text">
                {{ formData.surgery_date || '请选择手术日期' }}
              </view>
            </picker>
          </view>

          <!-- 图片上传区域 -->
          <view class="form-item">
            <text class="form-label">其他图片</text>
            <view class="uploader-container">
              <!-- 已上传图片预览 -->
              <view 
                v-for="(item, index) in imageList" 
                :key="index" 
                class="uploader-item"
              >
                <image 
                  :src="item.url" 
                  mode="aspectFill" 
                  class="uploader-image"
                />
                <view 
                  class="delete-btn" 
                  @click.stop="removeImage(index)"
                >
                  <text class="iconfont icon-close"></text>
                </view>
              </view>
              
              <!-- 添加图片按钮 -->
              <view 
                v-if="imageList.length < 9" 
                class="uploader-item add-btn" 
                @click="chooseImage"
              >
                <text class="iconfont icon-plus"></text>
                <text class="add-text">添加图片</text>
              </view>
            </view>
            <text class="upload-tip">最多可上传9张图片</text>
          </view>

          <!-- 提交按钮 -->
          <button 
            form-type="submit" 
            class="submit-btn"
            :disabled="isSubmitting"
          >
            {{ isSubmitting ? '提交中...' : '提交' }}
          </button>
        </form>
      </view>
    </view>
  </Layout>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import Layout from '../layout.vue';
import { addWarrantyCard } from '@/api/organ';

interface FormData {
  card_number: string;
  serial_number: string;
  patient_name: string;
  patient_age: string;
  surgeon_name: string;
  surgery_date: string;
  remark: string;
  credential: string[];
}

interface ImageItem {
  url: string;
  file?: any; // Uniapp 临时文件对象
  uploaded?: boolean; // 是否已上传
}

const formData = ref<FormData>({
  card_number: '',
  serial_number: '',
  patient_name: '',
  patient_age: '',
  surgeon_name: '',
  surgery_date: '',
  remark: '',
  credential: []
});

const imageList = ref<ImageItem[]>([]);
const isSubmitting = ref(false);

// 选择图片
const chooseImage = () => {
  uni.chooseImage({
    count: 9 - imageList.value.length,
    sizeType: ['compressed'], // 压缩图
    sourceType: ['album', 'camera'], // 相册和相机
    success: (res) => {
      res.tempFiles.forEach(file => {
        imageList.value.push({
          url: file.path,
          file: file,
          uploaded: false
        });
      });
    },
    fail: (err) => {
      console.error('选择图片失败:', err);
      uni.showToast({
        title: '选择图片失败',
        icon: 'none'
      });
    }
  });
};

// 删除图片
const removeImage = (index: number) => {
  imageList.value.splice(index, 1);
};

// 日期选择
const onDateChange = (e: any) => {
  formData.value.surgery_date = e.detail.value;
};

// 上传单张图片
const uploadSingleImage = async (filePath: string): Promise<string> => {
  return new Promise((resolve, reject) => {
    uni.uploadFile({
      url: '你的上传接口地址',
      filePath: filePath,
      name: 'file',
      formData: {},
      success: (uploadRes) => {
        const res = JSON.parse(uploadRes.data);
        if (res.code === 0) {
          resolve(res.result.url);
        } else {
          reject(new Error(res.message));
        }
      },
      fail: (err) => {
        reject(err);
      }
    });
  });
};

// 提交表单
const submitForm = async () => {
  if (!formData.value.card_number || !formData.value.serial_number) {
    uni.showToast({
      title: '请输入质保卡号和产品序列号',
      icon: 'none'
    });
    return;
  }

  isSubmitting.value = true;
  
  try {
    // 上传所有图片
    const uploadedUrls = [];
    for (const item of imageList.value) {
      if (!item.uploaded) {
        const url = await uploadSingleImage(item.url);
        uploadedUrls.push(url);
        item.uploaded = true;
      }
    }
    formData.value.credential = uploadedUrls;

    // 提交表单数据
    const res = await addWarrantyCard(formData.value);
    if (res.code === 0) {
      uni.showToast({
        title: '提交成功',
        icon: 'success'
      });
      setTimeout(() => {
        uni.navigateBack();
      }, 1500);
    } else {
      uni.showToast({
        title: res.message || '提交失败',
        icon: 'none'
      });
    }
  } catch (error) {
    console.error('提交失败:', error);
    uni.showToast({
      title: error.message || '提交失败',
      icon: 'none'
    });
  } finally {
    isSubmitting.value = false;
  }
};
</script>

<style scoped>
.page-container {
  min-height: 100vh;
  background-color: #f8f8f8;
}

.header {
  background: linear-gradient(to right, #ff1b6b, #ff9147);
  padding: 30rpx;
}

.title {
  font-size: 36rpx;
  font-weight: bold;
  color: #fff;
  margin-left: 20rpx;
}

.form-container {
  padding: 30rpx;
}

.form-item {
  margin-bottom: 40rpx;
  background-color: #fff;
  padding: 30rpx;
  border-radius: 16rpx;
  box-shadow: 0 2rpx 10rpx rgba(0, 0, 0, 0.05);
}

.form-label {
  display: block;
  font-size: 32rpx;
  color: #333;
  margin-bottom: 20rpx;
  font-weight: bold;
}

.form-input, .picker-text {
  width: 100%;
  height: 80rpx;
  border: 1rpx solid #e5e5e5;
  border-radius: 8rpx;
  font-size: 28rpx;
}

.picker-text {
  display: flex;
  align-items: center;
  color: #999;
}

.uploader-container {
  display: flex;
  flex-wrap: wrap;
  margin-top: 20rpx;
}

.uploader-item {
  position: relative;
  width: 160rpx;
  height: 160rpx;
  margin-right: 20rpx;
  margin-bottom: 20rpx;
  border-radius: 8rpx;
  overflow: hidden;
}

.uploader-image {
  width: 100%;
  height: 100%;
}

.add-btn {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  background-color: #f5f5f5;
  border: 1rpx dashed #ccc;
}

.icon-plus {
  font-size: 50rpx;
  color: #999;
  margin-bottom: 10rpx;
}

.add-text {
  font-size: 24rpx;
  color: #666;
}

.delete-btn {
  position: absolute;
  top: 0;
  right: 0;
  width: 40rpx;
  height: 40rpx;
  background-color: rgba(0, 0, 0, 0.5);
  border-radius: 0 0 0 8rpx;
  display: flex;
  justify-content: center;
  align-items: center;
}

.icon-close {
  color: #fff;
  font-size: 24rpx;
}

.upload-tip {
  display: block;
  font-size: 24rpx;
  color: #999;
  margin-top: 10rpx;
}

.submit-btn {
  margin-top: 60rpx;
  background: linear-gradient(to right, #ff1b6b, #ff9147);
  color: #fff;
  height: 90rpx;
  line-height: 90rpx;
  border-radius: 45rpx;
  font-size: 32rpx;
}

.submit-btn[disabled] {
  opacity: 0.7;
}
</style>