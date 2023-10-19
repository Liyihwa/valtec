<template>
  <div class="flex gap-2">
    <img
      :class="{ active: index === heroStore.currentSkill }"
      class="mix-blend-difference w-10 h-10 p-1 hover:border-1 box-border rounded-1 hover:cursor-pointer"
      v-for="(item, index) in skills"
      :key="index"
      :src="item.img"
      alt=""
      @click="handleChooseSkill(index)"
    />
  </div>
</template>

<script setup lang="ts">
import { Skill } from '@/types'
import { PropType } from 'vue'
import { useHeroStore } from '@/store/HeroStore'

const heroStore = useHeroStore()

const { skills } = defineProps({
  skills: {
    type: Object as PropType<Skill[]>,
    required: true,
  },
})

const handleChooseSkill = (index: number) => {
  heroStore.setCurrentSkill(index)
}
</script>

<style scoped>
.active {
  @apply border-1;
}
</style>
