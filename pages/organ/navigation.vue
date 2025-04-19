<template>
  <!-- 底部导航栏 -->
  <view class="tab-bar">
    <!-- 质保卡录入 -->
    <navigator 
      url="/pages/organ/warranty/add" 
      class="tab-item"
      :class="{active: currentPath === 'pages/organ/warranty/add'}"
      open-type="navigate"
    >
      <view class="icon-container">
        <text class="iconfont icon-plus-circle"></text>
      </view>
      <text class="tab-text">录入</text>
    </navigator>
    
    <!-- 质保卡列表 -->
    <navigator 
      url="/pages/organ/warranty/list" 
      class="tab-item"
      :class="{active: currentPath === 'pages/organ/warranty/list'}"
      open-type="navigate"
    >
      <view class="icon-container">
        <text class="iconfont icon-list"></text>
      </view>
      <text class="tab-text">列表</text>
    </navigator>
    
    <!-- 质保卡查询 -->
    <navigator 
      url="/pages/organ/user" 
      class="tab-item"
      :class="{active: currentPath === 'pages/organ/user'}"
      open-type="navigate"
    >
      <view class="icon-container">
        <text class="iconfont icon-user"></text>
      </view>
      <text class="tab-text">我的</text>
    </navigator>
  </view>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue';

const currentPath = ref('');

onMounted(() => {
  const pages = getCurrentPages();
  currentPath.value = pages[pages.length - 1]?.route || '';
});
</script>

<style scoped>
.tab-bar {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  height: 100rpx;
  background-color: #ffffff;
  border-top: 1rpx solid #e5e5e5;
  display: flex;
  justify-content: space-around;
  align-items: center;
  z-index: 999;
}

.tab-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  color: #666666;
  text-decoration: none;
  flex: 1;
  height: 100%;
}

.tab-item.active {
  color: #ff1b6b;
}

.icon-container {
  font-size: 40rpx;
  margin-bottom: 4rpx;
}

.tab-text {
  font-size: 24rpx;
}

/* 点击效果 */
.tab-item:active {
  transform: scale(0.95);
  transition: transform 0.1s ease;
}

/* 图标激活状态颜色 */
.tab-item.active .iconfont {
  color: #ff1b6b;
}
</style>