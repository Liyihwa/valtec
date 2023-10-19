<template>
  <el-dialog v-model="modelSeen" draggable title="" center>
    <div class="z-99 h-full w-full bg-gray-300/50 p-8 rounded-1 flex flex-col items-end">
      <div class="flex items-center flex-col justify-center mx-auto">
        <p v-html="description"></p>
      </div>
    </div>
  </el-dialog>
</template>

<script lang="ts" setup>
import { storeToRefs } from 'pinia'
import { useMapStore } from '@/store/MapStore'
import axios from 'axios'
import { onMounted, ref } from 'vue'

const mapStore = useMapStore()

const storeRef = storeToRefs(mapStore)
const modelSeen = storeRef.resourceModelSeen

const { id } = defineProps({
  id: {
    type: Number,
    required: true,
  },
})
const description = ref<string>()
onMounted(async () => {
  const result = await axios.get('http://10.16.44.76:8001/position/' + id)
  description.value = result.data.data.description
})
</script>
<style scoped>
.el-dialog_body {
  @apply bg-red-200;
}
.active {
  @apply border-blue-400 border-3 bg-red-200;
}
</style>
