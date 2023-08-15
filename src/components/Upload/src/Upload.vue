<script setup lang="ts">
import {
  ElUpload,
  ElDialog,
  ElButton,
  ElTableColumn,
  ElTable,
  ElImage,
  ElIcon,
  ElPagination
} from 'element-plus'
import {
  defineComponent,
  PropType,
  ref,
  computed,
  unref,
  toRefs,
  watch,
  onMounted,
  emit
} from 'vue'
import { propTypes } from '@/utils/propTypes'
import { Picture as IconPicture } from '@element-plus/icons-vue'
import { useI18n } from '@/hooks/web/useI18n'
import { uploadImageApi } from '@/api/upload'
import { basicProps } from './props'
import { FileItem, UploadResultStatus } from './typing'
import { useMessage } from '@/hooks/web/useMessage'
import { buildUUID } from '@/utils/uuid'
import { checkImgType, getBase64WithFile } from './helper'
import type { UploadProps, UploadUserFile, UploadFiles } from 'element-plus'
// import { composeEventHandlers } from 'element-plus/es/utils'

const fileList = ref<UploadUserFile[]>()

const { t } = useI18n()

const props = {
  src: propTypes.string.def(''),
  modelValue: {
    type: String,
    required: true
  }
}
const thumbUrl = ref('')
const emit = defineEmits(['change', 'update:modelValue'])

watch(
  () => props.modelValue,
  (val: string) => {
    if (val === unref(thumbUrl)) return
    thumbUrl.value = val
  },
  {
    immediate: true
  }
)
// 监听
watch(
  () => thumbUrl.value,
  (val: string) => {
    emit('update:modelValue', val)
  }
)
console.log('dddddddddd')
console.log(props.src)
</script>

<template>
  <div class="block">
    <span class="demonstration">Custom --{{ thumbUrl }}</span>
    <el-image :src="thumbUrl.value">
      <template #error>
        <div class="image-slot">
          <el-icon><icon-picture /></el-icon>
        </div>
      </template>
    </el-image>
  </div>
</template>
<style scoped>
.demo-image__error .block {
  padding: 30px 0;
  text-align: center;
  border-right: solid 1px var(--el-border-color);
  display: inline-block;
  width: 49%;
  box-sizing: border-box;
  vertical-align: top;
}
.demo-image__error .demonstration {
  display: block;
  color: var(--el-text-color-secondary);
  font-size: 14px;
  margin-bottom: 20px;
}
.demo-image__error .el-image {
  padding: 0 5px;
  max-width: 300px;
  max-height: 200px;
  width: 100%;
  height: 200px;
}

.demo-image__error .image-slot {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
  background: var(--el-fill-color-light);
  color: var(--el-text-color-secondary);
  font-size: 30px;
}
.demo-image__error .image-slot .el-icon {
  font-size: 30px;
}
</style>
