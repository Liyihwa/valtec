import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import Pages from 'vite-plugin-pages'
import Components from 'unplugin-vue-components/vite'
import AutoImport from 'unplugin-auto-import/vite'
import path from 'path'
import Unocss from 'unocss/vite'
import { NaiveUiResolver } from 'unplugin-vue-components/resolvers'

const pathSrc = path.resolve(__dirname, 'src')
// https://vitejs.dev/config/
export default defineConfig({
  resolve: {
    alias: {
      '@': pathSrc,
    },
  },
  plugins: [
    vue({
      reactivityTransform: true,
    }),
    Unocss(),
    Pages(),
    Components({
      dts: path.resolve(pathSrc, 'typings/components.d.ts'),
    }),
    AutoImport({
      imports: ['vue', 'vue-router', '@vueuse/core'],
      dirs: [path.resolve(pathSrc, 'composables')],
      dts: path.resolve(pathSrc, 'typings/components.d.ts'),
    }),
    Components({
      resolvers: [NaiveUiResolver()],
    }),
  ],
})
