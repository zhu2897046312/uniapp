<template>
  <view class="app-container">
    <!-- 主要内容区域 -->
    <view class="main-content">
      <!-- 使用 page-meta 确保页面高度正确 -->
      <page-meta page-style="padding-bottom: 100px"></page-meta>
      <slot></slot>
    </view>
    
    <!-- 底部导航栏（始终显示） -->
    <Navigation />
  </view>
</template>

<script lang="ts" setup>
import Navigation from './navigation.vue';
</script>

<style scoped>
.app-container {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  position: relative;
}

.main-content {
  flex: 1;
  width: 100%;
  padding-bottom: env(safe-area-inset-bottom); /* 兼容全面屏底部 */
}
</style>