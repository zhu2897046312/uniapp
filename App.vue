<script setup lang="ts">
import { onMounted, onUnmounted } from 'vue';
import not_login_event from './events/not_login_event';
import { useRouter } from 'vue-router';

const router = useRouter()

onMounted(() => {
  not_login_event.on(() => {
    router.push('/organ/login')
  })
})

onUnmounted(() => {
  not_login_event.removeAll()
})


</script>

<template>
  <router-view></router-view>
</template>

