<template>
  <div @click="handleChangeMap" class="map-container hover:cursor-pointer rounded-2 overflow-hidden w-400px text-center flex items-center justify-center relative">
    <img class="map w-full object-cover transition duration-300 ease" :src="mapStore.currentMap?.avatar" :alt="mapStore.currentMap?.name" />
    <div class="fw-800 text-3xl font-news absolute flex items-center justify-center w-full h-full">
      <div class="text-center text-white">{{ mapStore.currentMap?.name }}</div>
    </div>
  </div>
  <MapDialog />
</template>

<script lang="ts" setup>
import { useMapStore } from '@/store/MapStore'

import { onMounted } from 'vue'

let mapStore = useMapStore()

const fetchCurrentHero = async () => {
  mapStore = useMapStore()
  await mapStore.getAllMaps()
}

onMounted(async () => {
  await fetchCurrentHero()
})

const handleChangeMap = () => {
  mapStore.setModelSeen(true)
}
</script>

<style scoped>
.map-container:hover .map {
  transform: scale(1.2);
}
</style>
