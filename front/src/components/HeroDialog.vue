<template>
  <el-dialog v-model="modelSeen" draggable title="Choose Hero" center>
    <div class="z-99 p-8 rounded-1 flex flex-col items-around">
      <div class="grid grid-rows-3 lg:grid-rows-4 md:grid-rows-5 sm:grid-rows-8 gap-4 grid-flow-col">
        <template v-for="item in heros" :key="item.id">
          <img
            :class="{ active: item.id === beforConfirmId }"
            class="w-15 h-15 border-1 rounded-1 hover:border-blue-400 hover-cursor-pointer p-1 border-[#343434]"
            :src="item.avatar"
            alt=""
            @click="handleChoosingHero(item.id)"
          />
        </template>
      </div>
      <div class="flex justify-center gap-10 w-full mt-5">
        <button class="bg-green-400 px-3 py-1 rounded-1 hover:bg-green-300" @click="handleChooseHero">确认</button>
        <button class="bg-red-400 px-3 py-1 rounded-1 hover:bg-red-300" @click="handleCancel">取消</button>
      </div>
    </div>
  </el-dialog>
</template>

<script lang="ts" setup>
import { storeToRefs } from 'pinia'
import { useHeroStore } from '@/store/HeroStore'
import { ref } from 'vue'

const heroStore = useHeroStore()
const heros = heroStore.heros
const beforConfirmId = ref<number>(heroStore?.currentHero?.id)

const handleCancel = () => {
  heroStore.changeModelSeen(false)
}
const handleChooseHero = () => {
  heroStore.changeModelSeen(false)
  heroStore.setCurrentHeroId(beforConfirmId.value)
}

const handleChoosingHero = (id: number) => {
  beforConfirmId.value = id
}

const storeRef = storeToRefs(heroStore)
const modelSeen = storeRef.modelSeen
</script>
<style scoped>
.el-button--text {
  margin-right: 15px;
}
.el-select {
  width: 300px;
}
.el-input {
  width: 300px;
}
.dialog-footer button:first-child {
  margin-right: 10px;
}

.active {
  @apply border-blue-400 border-3 bg-red-200;
}
</style>
