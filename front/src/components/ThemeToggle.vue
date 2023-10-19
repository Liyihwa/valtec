<template>
  <div ref="toggleBtn" class="i-ri-sun-line dark:i-ri-moon-line hover:cursor-pointer"></div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { ref } from 'vue'

const toggleBtn = ref<HTMLElement>()
const theme = ref(localStorage.getItem('theme') || 'dark')

onMounted(() => {
  if (theme.value === 'dark') {
    document.documentElement.classList.add('dark')
  } else {
    document.documentElement.classList.remove('dark')
  }
  toggleBtn.value!.addEventListener('click', () => {
    if (theme.value === 'dark') {
      localStorage.setItem('theme', 'light')
      document.documentElement.classList.remove('dark')
      theme.value = 'light'
    } else {
      localStorage.setItem('theme', 'dark')
      document.documentElement.classList.add('dark')
      theme.value = 'dark'
    }
  })
})
</script>

<style scoped></style>
