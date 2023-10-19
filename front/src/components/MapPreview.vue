<template>
  <div id="map" class="rounded-2 w-full h-full overflow-hidden bg-black w-full h-full">
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

      <ol-interaction-select @select="featureSelected" :condition="selectCondition">
        <ol-style>
          <ol-style-stroke color="green" :width="10"></ol-style-stroke>
          <ol-style-fill color="rgba(255,255,255,0.5)"></ol-style-fill>
          <ol-style-icon :src="heroStore?.heroMarker" :scale="0.05"></ol-style-icon>
        </ol-style>
      </ol-interaction-select>

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
        <ol-vector-layer v-for="(position, index) in mapStore.positions" :key="index">
          <ol-source-vector>
            <ol-feature>
              <ol-geom-line-string
                :id="position.id"
                :coordinates="[
                  [position.px, position.py],
                  [position.sx, position.sy],
                ]"
              ></ol-geom-line-string>
              <ol-style>
                <ol-style-stroke :color="'green'" :width="2"></ol-style-stroke>
              </ol-style>
            </ol-feature>
          </ol-source-vector>
        </ol-vector-layer>
      </ol-vector-layer>
    </ol-map>
  </div>
  <template v-if="mapStore.resourceModelSeen">
    <ResourceModel :id="clickedPositionId" />
  </template>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, inject, watch } from 'vue'
import { useMapStore } from '@/store/MapStore'
import { useHeroStore } from '@/store/HeroStore'
import { SingleItem } from 'node_modules/ol-contextmenu/dist/types'
import { View, Map } from 'ol'
import { Point } from 'ol/geom'
import TileLayer from 'ol/layer/Tile'
import VectorLayer from 'ol/layer/Vector'
import { OSM } from 'ol/source'
import VectorSource from 'ol/source/Vector'
import { Style, Fill, Stroke } from 'ol/style'
import CircleStyle from 'ol/style/Circle'
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
  await mapStore.getPositions()
}

const map = ref<Map>()
const clickedPositionId = ref<number>(0)
onMounted(async () => {
  await fetchResource()
  imgUrl.value = mapStore.currentMap.url
  imgCopyright.value = '© <a href="https://github.com/blkcor">ValorantTechLab</a>'
  initHeroPosition()
  initSkillPosition()

  const osmLayer = new TileLayer({
    source: new OSM(),
  })

  const vectorSource = new VectorSource()
  const vectorLayer = new VectorLayer({
    source: vectorSource,
    style: new Style({
      image: new CircleStyle({
        radius: 7,
        fill: new Fill({ color: 'red' }),
        stroke: new Stroke({ color: 'white', width: 2 }),
      }),
    }),
  })

  map.value = new Map({
    target: 'map',
    layers: [osmLayer, vectorLayer],
    view: new View({
      center: [0, 0],
      zoom: 2,
    }),
  })

  // Example: Add a point feature
  const pointFeature = new Feature(new Point([0, 0]))
  vectorSource.addFeature(pointFeature)
})

const contextMenuItems = ref<SingleItem[]>([])

const HeroMarker = ref<any>(null)
const SkillMarker = ref<any>(null)
const view = ref<any>(null)

const Feature = inject<any>('ol-feature')
const Geom = inject<any>('ol-geom')

const initHeroPosition = () => {
  HeroMarker.value!.source.clear()
  mapStore.positions.forEach((position) => {
    const feature = new Feature({
      geometry: new Geom.Point([position.sx, position.sy]),
      id: position.id,
    })
    HeroMarker.value!.source.addFeature(feature)
  })
}

const initSkillPosition = () => {
  SkillMarker.value!.source.clear()
  mapStore.positions.forEach((position) => {
    const feature = new Feature({
      geometry: new Geom.Point([position.px, position.py]),
      //设置资源
      id: position.id,
    })
    SkillMarker.value!.source.addFeature(feature)
  })
}

watch(
  () => mapStore.positions,
  async () => {
    await mapStore.getPositions()
    initHeroPosition()
    initSkillPosition()
  }
)

watch(
  () => mapStore.currentMapId,
  () => {
    imgUrl.value = mapStore.currentMap.url
  }
)

const selectConditions = inject('ol-selectconditions')

const selectCondition = selectConditions.singleClick

const featureSelected = (event: any) => {
  //cursor pointer
  const mapElement = document.querySelector('.ol-viewport') as HTMLElement
  if (event.selected.length > 0) {
    // 检查是否选择了一个feature
    mapElement.style.cursor = 'pointer' // 设置cursor为pointer
    clickedPositionId.value = event.selected[0].values_.id
    mapStore.resourceModelSeen = true
  } else {
    mapElement.style.cursor = '' // 重置cursor样式
  }
}
</script>

<style scoped></style>
