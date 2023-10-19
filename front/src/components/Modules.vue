<template>
  <div class="flex gap-2 items-center border-r px-2 border-r-2 sm:border-none">
    <template v-for="item in modules" :key="item.link">
      <router-link
        :to="item.link"
        :class="{ active: (route.path === '/' && item.type === 'teach') || route.path === '/' + item.type }"
        class="flex gap-1 items-center justify-center px-2 py-1 border border-2 border-transparent hover:border-b-black dark:hover:border-b-white hover:text-gray-700 dark:hover:text-white"
      >
        <span>{{ item.text }}</span>
      </router-link>
    </template>
  </div>
</template>

<script setup lang="ts">
import modules from '@/content/module'
import { useRoute } from 'vue-router'

const route = useRoute()
</script>

<style scoped>
.dark .active {
  border-bottom-color: white;
}
.active {
  border-bottom-color: black;
}
</style>
