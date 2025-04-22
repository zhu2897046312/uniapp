<script lang="ts" setup>
import Layout from '../layout.vue';
import { ref, onMounted } from 'vue'
import { getInfoByWarrantyNo } from '@/api/organ'

interface CardInfo {
  card_number: string
  serial_number: string
  patient_name: string
  patient_age: string
  surgeon_name: string
  institution_name: string
  status: string
  statusText: string
  credential: string[]
  bind_mobile: string
  remark: string
  surgery_date: string
  product_name: string
  prosthesis_model: string
}

const warrantyNo = ref('')
const formData = ref<CardInfo | null>(null)
const doctorAvatar = ''
const surgeryImages = ref<string[]>([])

const fetchData = async (warrantyNo: string) => {
  try {
    const res = await getInfoByWarrantyNo(warrantyNo)
    formData.value = res.result
    
    if (formData.value?.credential && formData.value.credential.length > 0) {
      surgeryImages.value = formData.value.credential
    }
  } catch (error) {
    uni.showToast({ title: '加载失败', icon: 'none' })
    console.error(error)
  }
}

const showPreview = (index: number) => {
  uni.previewImage({
    current: index,
    urls: surgeryImages.value
  })
}

onMounted(() => {
  const pages = getCurrentPages()
  const currentPage = pages[pages.length - 1]
  warrantyNo.value = currentPage.options?.warrantyNo || ''
  
  if (warrantyNo.value) {
    fetchData(warrantyNo.value)
  }
})
</script>

<template>
  <Layout>
  <view class="page-container">
    <!-- 顶部导航栏 -->
    <view class="nav-bar">
      <view class="nav-content">
        <uni-icons 
          type="arrowleft" 
          size="24" 
          color="#fff"
          @click="uni.navigateBack()"
        />
        <text class="nav-title">质保详情</text>
      </view>
    </view>

    <!-- 主要内容区域 -->
    <scroll-view class="content-container" scroll-y>
      <!-- 质保卡信息 -->
      <view class="card">
        <view class="card-item">
          <text class="card-label">质保卡号</text>
          <text class="card-value">{{ formData?.card_number || '-' }}</text>
        </view>
        
        <view class="card-item">
          <text class="card-label">姓名</text>
          <text class="card-value">
            {{ formData?.patient_name || '-' }}
            <text v-if="formData?.patient_age" class="card-age">({{ formData.patient_age }}岁)</text>
          </text>
        </view>
        
        <view class="card-item">
          <text class="card-label">手术机构</text>
          <view class="card-value-with-icon">
            <uni-icons type="home" size="18" color="#ff1b6b" />
            <text>{{ formData?.institution_name || '-' }}</text>
          </view>
        </view>
        
        <view class="card-item">
          <text class="card-label">手术医生</text>
          <view class="card-value-with-icon">
            <image :src="doctorAvatar" class="doctor-avatar" />
            <text>{{ formData?.surgeon_name || '-' }}</text>
          </view>
        </view>
        
        <view class="card-item">
          <text class="card-label">手术日期</text>
          <text class="card-value">{{ formData?.surgery_date || '-' }}</text>
        </view>
        
        <!-- 图片列表 -->
        <view v-if="surgeryImages.length > 0" class="image-section">
          <text class="section-title">相关图片</text>
          <scroll-view class="image-list" scroll-x>
            <view 
              v-for="(img, index) in surgeryImages" 
              :key="index" 
              class="image-wrapper"
              @click="showPreview(index)"
            >
              <image :src="img" mode="aspectFill" class="image" />
            </view>
          </scroll-view>
        </view>
      </view>

      <!-- 产品信息 -->
      <view class="card product-info">
        <text class="section-title">产品信息</text>
        
        <view class="info-item">
          <text class="info-label">序列号</text>
          <text class="info-value code-font">{{ formData?.serial_number || '-' }}</text>
        </view>
        
        <view class="info-item">
          <text class="info-label">产品名称</text>
          <text class="info-value">{{ formData?.product_name || '-' }}</text>
        </view>
        
        <view class="info-item">
          <text class="info-label">产品规格</text>
          <text class="info-value">{{ formData?.prosthesis_model || '-' }}</text>
        </view>
      </view>
    </scroll-view>
  </view>
  </Layout>
</template>

<style scoped>
/* 页面基础样式 */
.page-container {
  display: flex;
  flex-direction: column;
  height: 100vh;
  background-color: #f5f5f5;
}

/* 导航栏样式 */
.nav-bar {
  background: linear-gradient(to right, #ff1b6b, #ff9147);
  color: white;
  left: 0;
  right: 0;
}

.nav-content {
  display: flex;
  align-items: center;
  height: 44px;
}

.nav-title {
  font-size: 18px;
  font-weight: bold;
  margin-left: 15px;
}

/* 内容区域 */
.content-container {
  flex: 1;
  padding: 15px 15px 15px; /* 留出导航栏空间 */
  box-sizing: border-box;
}

/* 卡片样式 */
.card {
  background-color: #fff;
  border-radius: 8px;
  padding: 15px;
  margin-bottom: 15px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.card-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 0;
  border-bottom: 1px solid #f0f0f0;
}

.card-item:last-child {
  border-bottom: none;
}

.card-label {
  color: #666;
  font-size: 14px;
}

.card-value {
  color: #333;
  font-size: 14px;
}

.card-age {
  color: #999;
  margin-left: 5px;
}

.card-value-with-icon {
  display: flex;
  align-items: center;
}

.doctor-avatar {
  width: 24px;
  height: 24px;
  border-radius: 50%;
  margin-right: 8px;
  background-color: #f0f0f0;
}

/* 图片区域 */
.image-section {
  margin-top: 15px;
}

.section-title {
  display: block;
  font-size: 16px;
  font-weight: bold;
  color: #333;
  margin-bottom: 10px;
}

.image-list {
  white-space: nowrap;
}

.image-wrapper {
  display: inline-block;
  width: 100px;
  height: 100px;
  margin-right: 10px;
}

.image {
  width: 100%;
  height: 100%;
  border-radius: 4px;
}

/* 产品信息 */
.product-info {
  margin-top: 15px;
}

.info-item {
  display: flex;
  justify-content: space-between;
  padding: 12px 0;
}

.info-label {
  color: #666;
}

.info-value {
  color: #333;
}

.code-font {
  font-family: 'Courier New', Courier, monospace;
}

/* 响应式调整 */
@media (min-width: 768px) {
  .content-container {
    padding: 74px 20px 20px;
  }
  
  .card {
    padding: 20px;
  }
  
  .image-wrapper {
    width: 120px;
    height: 120px;
  }
}
</style>