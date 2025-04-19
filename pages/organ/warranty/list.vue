<script lang="ts" setup>
import Layout from '../layout.vue';
import { ref, onMounted, watch } from 'vue';
import { getWarrantyCardList } from '@/api/organ'

const loaded = ref(false)

interface WarrantCardInfo {
  card_number: string,
  status: string,
  statusText: string,
  serial_number: string,
  patient_name: string,
  surgeon_name: string,
  surgery_date: string
}
const warrantyList = ref<WarrantCardInfo[] | null>(null);
const searchText = ref('');
const getStatusClass = (status: string) => {
  return {
    'status-valid': status === 'valid',
    'status-expired': status === 'expired'
  };
};

const goToDetail = (warrantyNo: string) => {
  console.log("warrantyNo", warrantyNo)
  if (!warrantyNo) {
    uni.showToast({
      title: '质保卡号不能为空',
      icon: 'none'
    })
    return
  }
  uni.navigateTo({
    url: `/pages/organ/warranty/detail?warrantyNo=${warrantyNo}&from=search`
  })
}

const currentPage = ref(1)
const pageSize = ref(4)
const total = ref(0)
const fetchWarrantyList = async () => {
  uni.showLoading({
    title: '加载中...',
    mask: true
  });
  
  try {
    const res = await getWarrantyCardList({
      cardNumber: searchText.value,
      page: currentPage.value,
      pageSize: pageSize.value
    });
    
    if (res.code === 0) {
      warrantyList.value = res.result;
      total.value = res.result?.length || 0;
      loaded.value = true;
    } else {
      uni.showToast({
        title: res.message || '请求失败',
        icon: 'none'
      });
    }
  } catch (error) {
    uni.showToast({
      title: error instanceof Error ? error.message : '请求异常',
      icon: 'none'
    });
  } finally {
    setTimeout(() => uni.hideLoading(), 1000);
  }
}

// 防抖搜索
let searchTimeout: ReturnType<typeof setTimeout>
const debounceSearch = (fn: Function, delay = 500) => {
  clearTimeout(searchTimeout)
  searchTimeout = setTimeout(fn, delay)
}

watch(searchText, () => {
  debounceSearch(() => {
    currentPage.value = 1 // 搜索时重置到第一页
    fetchWarrantyList()
  })
})

onMounted(() => {
  fetchWarrantyList()
})
</script>

<template>
	<Layout>
  <view class="page-container">
    <!-- 顶部搜索区域 -->
    <view class="search-header">
      <view class="search-input-container">
        <input
          type="text"
          v-model="searchText"
          placeholder="搜索质保卡号"
          class="search-input"
        />
        <text class="iconfont icon-search search-icon"></text>
      </view>
    </view>

    <!-- 列表内容区域 -->
    <template v-if="loaded">
      <view v-if="total > 0" class="list-container">
        <view 
          v-for="(item, index) in warrantyList" 
          :key="index" 
          class="warranty-card"
          @click="goToDetail(item.card_number)"
        >
          <view class="card-content">
            <view class="card-header">
              <text class="card-number">{{ item.card_number }}</text>
              <text class="status-tag" :class="getStatusClass(item.status)">
                {{ item.statusText }}
              </text>
            </view>
            <view class="info-list">
              <view class="info-item">
                <text class="info-label">序列号</text>
                <text class="info-value">{{ item.serial_number }}</text>
              </view>
              <view class="info-item">
                <text class="info-label">患者姓名</text>
                <text class="info-value">{{ item.patient_name }}</text>
              </view>
              <view class="info-item">
                <text class="info-label">手术医生</text>
                <text class="info-value">{{ item.surgeon_name }}</text>
              </view>
              <view class="info-item">
                <text class="info-label">手术时间</text>
                <text class="info-value">{{ item.surgery_date }}</text>
              </view>
            </view>
          </view>
          <text class="iconfont icon-arrow-right arrow-icon"></text>
        </view>
      </view>
      <view v-else class="empty-container">
        <text class="empty-text">暂无数据</text>
      </view>
    </template>
  </view>
  </Layout>
</template>

<style scoped>
.page-container {
  min-height: 100vh;
  background-color: #f8f8f8;
  padding-bottom: 80px;
}

.search-header {
  position: fixed;
  width: 100%;
  z-index: 50;
  padding: 12px 16px;
  background: linear-gradient(to right, #ff1b6b, #ff9147);
}

.search-input-container {
  position: relative;
  margin-right: 35px;
}

.search-input {
  width: 100%;
  height: 40px;
  font-size: 14px;
  border: none;
  border-radius: 20px;
  background-color: rgba(255, 255, 255, 0.9);
  color: #333;
}

.search-icon {
  position: absolute;
  left: 12px;
  top: 12px;
  color: #999;
  font-size: 16px;
}

.list-container {
  padding-top: 80px;
  padding-left: 16px;
  padding-right: 16px;
  padding-bottom: 16px;
}

.warranty-card {
  margin-bottom: 16px;
  background-color: #fff;
  border-radius: 12px;
  padding: 16px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  transition: transform 0.2s;
}

.warranty-card:active {
  transform: scale(0.98);
}

.card-header {
  display: flex;
  align-items: center;
  margin-bottom: 12px;
}

.card-number {
  font-size: 16px;
  font-weight: 600;
  color: #333;
}

.status-tag {
  margin-left: 8px;
  padding: 2px 8px;
  font-size: 12px;
  border-radius: 10px;
}

.status-valid {
  background-color: #e6f7e6;
  color: #52c41a;
}

.status-expired {
  background-color: #f5f5f5;
  color: #999;
}

.info-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.info-item {
  display: flex;
  align-items: center;
  font-size: 14px;
}

.info-label {
  width: 80px;
  color: #999;
}

.info-value {
  color: #666;
}

.arrow-icon {
  color: #ddd;
  font-size: 16px;
}

.empty-container {
  padding-top: 80px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.empty-text {
  color: #999;
  font-size: 14px;
}
</style>