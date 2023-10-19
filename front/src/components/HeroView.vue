<template>
  <div class="flex items-center gap-2 border-1 px-4 py-1 border-gray-500/50 w-full justify-around rounded-1">
    <img
      @click="heroStore.changeModelSeen(true)"
      class="w-15 h-15 hover:cursor-pointer border-1 border-blue-400/50 p-1 box-border hover:border-blue-400 rounded-[50%]"
      :src="heroStore.currentHero?.avatar"
      :alt="heroStore.currentHero?.name"
    />
    <Skills :skills="heroStore.currentHero?.skills" />
  </div>
  <!-- <template v-if="heroStore.modelSeen"> <HeroModel :data="heroStore.heros" /></template> -->
  <HeroDialog />
</template>

<script setup lang="ts">
import { useHeroStore } from '@/store/HeroStore'
import { onMounted } from 'vue'

const heroStore = useHeroStore()

const fetchCurrentHero = async () => {
  const heroStore = useHeroStore()
  await heroStore.getAllHeros()
}

onMounted(async () => {
  await fetchCurrentHero()
})
</script>

<style scoped></style>
