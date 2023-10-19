import { Map, Position } from '@/types'
import { defineStore } from 'pinia'
import axios from 'axios'
import { useHeroStore } from './HeroStore'

export const useMapStore = defineStore('map', {
  state: () => ({
    maps: [] as Map[],
    modelSeen: false,
    currentMapId: 0,
    editHeroPosition: [] as number[],
    editSkillPosition: [] as number[],
    positions: [] as Position[],
    resourceModelSeen: false,
    description: '这家伙什么都没写~',
  }),
  getters: {
    currentMap(): Map {
      return this.maps.find((item) => item.id === this.currentMapId)!
    },
  },
  actions: {
    async getAllMaps() {
      const maps = await axios.get('http://10.16.44.76:8001/maps')
      this.maps.length = 0
      maps.data.data.forEach((item: Map) => {
        this.maps.push(item)
      })
      this.setCurrentMapId(this.maps[0].id)
    },
    async getPositions() {
      const heroStore = useHeroStore()
      const positions = await axios.get(`http://10.16.44.76:8001/positions/${this.currentMapId}/${heroStore.currentHeroId}/${heroStore.currentSkillName}`)
      this.positions = positions.data.data
    },
    setCurrentMapId(id: number) {
      this.currentMapId = id
    },
    changeModelSeen(state: boolean) {
      this.modelSeen = state
    },
    setModelSeen(state: boolean) {
      this.modelSeen = state
    },
    setDescription(description: string) {
      this.description = description
    },
  },
})
