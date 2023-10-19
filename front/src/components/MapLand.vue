<template>
  <span v-if="mode === 'edit'">
    <MapEdit />
    <div class="flex flex-col mt-4">
      <div>
        <TextEditor />
      </div>
      <div class="flex justify-around mt-10">
        <button class="border px-2 py-1 rounded-1 mt-1 dark:border-white border-black bg-green-300/90 dark:text-white text-black hover:bg-green-400/90" @click="handleUpload">上传</button>
        <button @click="mode = 'preview'" class="border px-2 py-1 rounded-1 mt-1 dark:border-white border-black bg-blue-300/90 dark:text-white text-black hover:bg-blue-400/90">切换预览模式</button>
      </div>
    </div>
  </span>
  <span v-else>
    <MapPreview />
    <div class="flex justify-around mt-10">
      <button @click="mode = 'edit'" class="border px-2 py-1 rounded-1 mt-1 dark:border-white border-black bg-blue-300/90 dark:text-white text-black hover:bg-blue-400/90">切换编辑模式</button>
    </div>
  </span>
</template>

<script setup lang="ts">
import { useHeroStore } from '@/store/HeroStore'
import { useMapStore } from '@/store/MapStore'
import axios from 'axios'
import { onMounted, ref } from 'vue'
type Mode = 'edit' | 'preview'
const mode = ref<Mode>('edit')

const mapStore = useMapStore()
const heroStore = useHeroStore()

const fetchResource = async () => {
  await mapStore.getAllMaps()
  await heroStore.getAllHeros()
}

onMounted(async () => {
  await fetchResource()
})

const handleUpload = async () => {
  if (mapStore.editHeroPosition.length === 0 || mapStore.editSkillPosition.length === 0) {
    alert('请先添加英雄和技能')
    return
  }
  const result = await axios.post('http://10.16.44.76:8001/position', {
    hid: heroStore.currentHeroId,
    mid: mapStore.currentMapId,
    skill: heroStore.currentSkillName,
    sx: mapStore.editHeroPosition[0],
    sy: mapStore.editHeroPosition[1],
    px: mapStore.editSkillPosition[0],
    py: mapStore.editSkillPosition[1],
    description: mapStore.description || '这家伙什么都没有写!',
  })
  if (result.data.code === 200) {
    alert('上传成功')
  } else {
    alert('上传失败')
  }
  mapStore.editHeroPosition = []
  mapStore.editSkillPosition = []
  mapStore.setDescription('')
  location.reload()
}
</script>
