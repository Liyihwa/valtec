<template>
  <div class="rounded-2 w-full h-full overflow-hidden bg-black w-full h-full">
    <ol-map :loadTilesWhileAnimating="true" :loadTilesWhileInteracting="true" style="height: 100%; width: 100%">
      <ol-view ref="view" :center="center" :rotation="rotation" :zoom="zoom" :projection="projection" />
      <ol-zoom-control />
      <ol-attribution-control />
      <ol-image-layer>
        <ol-source-image-static :url="imgUrl" :imageSize="size" :imageExtent="extent" :projection="projection" :attributions="imgCopyright"></ol-source-image-static>
      </ol-image-layer>
      <ol-tile-layer>
        <ol-source-osm />
      </ol-tile-layer>

      <ol-context-menu-control :items="contextMenuItems" />

      <ol-vector-layer v-if="heroStore?.heroMarker">
        <ol-source-vector ref="HeroMarker"> </ol-source-vector>
        <ol-style>
          <ol-style-icon :src="heroStore?.heroMarker" :scale="0.4"></ol-style-icon>
        </ol-style>
      </ol-vector-layer>

      <ol-vector-layer v-if="heroStore?.skillMarker">
        <ol-source-vector ref="SkillMarker"> </ol-source-vector>
        <div>
          <ol-style>
            <ol-style-icon :src="heroStore?.skillMarker" :scale="0.5"></ol-style-icon>
          </ol-style>
        </div>
      </ol-vector-layer>
    </ol-map>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, inject } from 'vue'
import { useMapStore } from '@/store/MapStore'
import { useHeroStore } from '@/store/HeroStore'
import { SingleItem } from 'node_modules/ol-contextmenu/dist/types'
const mapStore = useMapStore()
const heroStore = useHeroStore()

const zoom = ref(1.5)
const rotation = ref(0)
const size = ref([1024, 968])
const center = ref([size.value[0] / 2, size.value[1] / 2])
const extent = ref([0, 0, ...size.value])

const projection = reactive({
  code: 'valorantTechLab',
  units: 'pixels',
  extent: extent,
})
const imgUrl = ref<string>()
const imgCopyright = ref<string>()

//fetch everything
const fetchResource = async () => {
  await mapStore.getAllMaps()
  await heroStore.getAllHeros()
}
onMounted(async () => {
  await fetchResource()
  imgUrl.value = mapStore.currentMap.url
  imgCopyright.value = '© <a href="https://github.com/blkcor">ValorantTechLab</a>'
})

const contextMenuItems = ref<SingleItem[]>([])

const HeroMarker = ref<any>(null)
const SkillMarker = ref<any>(null)
const view = ref<any>(null)

const Feature = inject<any>('ol-feature')
const Geom = inject<any>('ol-geom')

contextMenuItems.value = [
  {
    text: 'Add Hero',
    classname: 'some-style-class', // you can add this icon with a CSS class
    // instead of `icon` property (see next line)

    icon: '', // this can be relative or absolute
    callback: (val: any) => {
      HeroMarker.value!.source.clear()
      const feature = new Feature({
        geometry: new Geom.Point(val.coordinate),
      })
      mapStore.editHeroPosition = val.coordinate
      HeroMarker.value!.source.addFeature(feature)
    },
  },
  {
    text: 'Add Skill',
    classname: 'some-style-class', // you can add this icon with a CSS class
    // instead of `icon` property (see next line)

    icon: '', // this can be relative or absolute
    callback: (val: any) => {
      SkillMarker.value!.source.clear()
      const feature = new Feature({
        geometry: new Geom.Point(val.coordinate),
      })
      mapStore.editSkillPosition = val.coordinate
      SkillMarker.value!.source.addFeature(feature)
    },
  },
]
</script>

<style scoped></style>
