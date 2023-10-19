import { Hero } from '@/types'
import { defineStore } from 'pinia'
import axios from 'axios'

export const useHeroStore = defineStore('hero', {
  state: () => ({
    heros: [] as Hero[],
    currentHeroId: 0,
    currentSkill: 0,
    modelSeen: false,
  }),
  getters: {
    currentHero(): Hero {
      return this.heros.find((item) => item.id === this.currentHeroId)!
    },
    heroMarker(): string {
      return this.currentHero?.avatar
    },
    skillMarker(): string {
      return this.currentHero?.skills[this.currentSkill].img
    },
    currentSkillName(): string {
      return this.currentHero?.skills[this.currentSkill].type
    },
  },
  actions: {
    async getAllHeros() {
      const heros = await axios.get('http://10.16.44.76:8001/heros')
      this.heros.length = 0
      heros.data.data.forEach((item: Hero) => {
        this.heros.push(item)
      })
      this.setCurrentHeroId(this.heros[0].id)
    },
    setCurrentHeroId(id: number) {
      this.currentHeroId = id
      this.currentSkill = 0
    },
    setCurrentSkill(index: number) {
      this.currentSkill = index
    },
    changeModelSeen(state: boolean) {
      this.modelSeen = state
    },
  },
})
