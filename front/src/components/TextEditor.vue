<template>
  <div class="border-[1px] border-solid border-[#ccc]">
    <Toolbar style="border-bottom: 1px solid #ccc" :editor="editorRef" :defaultConfig="toolbarConfig" :mode="mode" />
    <Editor style="height: 500px; overflow-y: hidden" v-model="valueHtml" :defaultConfig="editorConfig" :mode="mode" @onCreated="handleCreated" />
  </div>
</template>

<script setup lang="ts">
//@ts-nocheck
import '@wangeditor/editor/dist/css/style.css' // 引入 css

import { onBeforeUnmount, ref, shallowRef, watch } from 'vue'
import { useMapStore } from '@/store/MapStore'
import { Editor, Toolbar, IDomEditor, IToolbarConfig, IEditorConfig } from '@wangeditor/editor-for-vue'

const mapStore = useMapStore()
// 编辑器实例，必须用 shallowRef
const editorRef = shallowRef<IDomEditor>()

// 内容 HTML
const valueHtml = ref()

// 模拟 ajax 异步获取内容

const toolbarConfig: Partial<IToolbarConfig> = {}
const editorConfig: Partial<IEditorConfig> = {
  placeholder: '请输入内容...',
  MENU_CONF: {
    uploadImage: {
      server: 'http://10.16.44.76:8002/upload',
      fieldName: 'file',
      headers: {},
    },
  },
}
const mode = 'default'
// 组件销毁时，也及时销毁编辑器
onBeforeUnmount(() => {
  const editor = editorRef.value
  if (editor == null) return
  editor.destroy()
})

const handleCreated = (editor: any) => {
  editorRef.value = editor // 记录 editor 实例，重要！
}

watch(valueHtml, () => {
  mapStore.setDescription(valueHtml.value)
})
</script>
