<template>
  <el-dialog v-model="modelSeen" draggable title="Choose Map" center>
    <div class="z-99 h-full w-full bg-gray-300/50 p-8 rounded-1 flex flex-col items-end">
      <div class="grid grid-rows-3 gap-4 grid-flow-col">
        <template v-for="item in maps" :key="item.id">
          <div class="flex flex-col items-center">
            <span>{{ item.name }}</span>
            <img
              :class="{ active: item.id === beforConfirmId }"
              class="w-full border-1 rounded-1 hover:border-blue-400 hover-cursor-pointer p-1 border-[#343434]"
              :src="item.avatar"
              alt=""
              @click="handleChoosingMap(item.id)"
            />
          </div>
        </template>
      </div>
      <div class="flex justify-center gap-10 w-full mt-5">
        <button class="bg-green-400 px-3 py-1 rounded-1 hover:bg-green-300" @click="handleChooseMap">确认</button>
        <button class="bg-red-400 px-3 py-1 rounded-1 hover:bg-red-300" @click="handleCancel">取消</button>
      </div>
    </div>
  </el-dialog>
</template>

<script lang="ts" setup>
import { storeToRefs } from 'pinia'
import { useMapStore } from '@/store/MapStore'
import { ref } from 'vue'

const mapStore = useMapStore()
const maps = mapStore.maps

const beforConfirmId = ref<number>(mapStore?.currentMapId)

const handleCancel = () => {
  mapStore.changeModelSeen(false)
}
const handleChooseMap = () => {
  mapStore.changeModelSeen(false)
  mapStore.setCurrentMapId(beforConfirmId.value)
}

const handleChoosingMap = (id: number) => {
  beforConfirmId.value = id
}

const storeRef = storeToRefs(mapStore)
const modelSeen = storeRef.modelSeen
</script>
<style scoped>
.el-dialog_body {
  @apply bg-red-200;
}
.active {
  @apply border-blue-400 border-3 bg-red-200;
}
</style>
