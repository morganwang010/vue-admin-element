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
  ...basicProps,
  maxNumber: 3,
  maxSize: 1024,
  previewFileList: {
    type: Array as PropType<string[]>,
    default: () => []
  },
  modelValue: propTypes.string.def(''),
  value: propTypes.string.def('')
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

const message = useMessage()
//   是否正在上传
const isUploadingRef = ref(false)
const fileListRef = ref<FileItem[]>([])
const files = ref<string[]>([])
// const { accept, helpText, maxNumber, maxSize } = toRefs(props)
const loading = ref(false)
// const tableData = {}
const { maxNumber } = props
const dialogVisible2 = ref(false)

const handleStartUpload = async () => {
  // const m = 'you reach the maxnumber' + maxNumber
  // w.warn(m)

  if ((fileListRef.value.length + props.previewFileList?.length ?? 0) > maxNumber) {
    return message.warn('Your file is to large')
  }
  try {
    isUploadingRef.value = true
    // 只上传不是成功状态的
    // console.log(fileList)
    // const uploadFileList =
    //   fileList.value.filter((item) => item.status !== UploadResultStatus.SUCCESS) || []
    // console.log('ssssssssss')
    // console.log(uploadFileList)
    // const data = await Promise.all(
    //   // uploadFileList.map((item) => {
    //   //   console.log(item)
    //   //   console.log('mmmmmmmm')
    //     return uploadImageApi(item)
    //   // })
    // )
    //构建一个formdData，用于上传文件
    let formData = new FormData()
    formData.append('file', fileList.value.raw)
    const data = await uploadImageApi(formData)
    console.log(data.name)

    isUploadingRef.value = false
    // 生产环境:抛出错误
    // const errorList = data.filter((item: any) => !item.success)
    // if (errorList.length > 0) throw errorList
  } catch (e) {
    isUploadingRef.value = false
    throw e
  }
}

const handleRemove: UploadProps['onRemove'] = (file, uploadFiles) => {
  console.log(file, uploadFiles)
}

const handlePreview: UploadProps['onPreview'] = (uploadFile) => {
  console.log(uploadFile)
  console.log('00000000000000000000')
}
const save = async () => {
  console.log('tttt')
}
const getStringAccept = 'file'
// 上传前校验
const beforeUpload = (file: File) => {
  console.log('555555555555')
  const { size, name } = file
  const { maxSize } = props
  // 设置最大值，则判断
  if (maxSize && file.size / 1024 / 1024 >= maxSize) {
    message.error(t('component.upload.maxSizeMultiple'))
    return false
  }
  const commonItem = {
    uuid: buildUUID(),
    file,
    size,
    name,
    percent: 0,
    type: name.split('.').pop()
  }
  // 生成图片缩略图
  if (checkImgType(file)) {
    // beforeUpload，如果异步会调用自带上传方法
    // file.thumbUrl = await getBase64(file);
    getBase64WithFile(file).then(({ result: thumbUrl }) => {
      fileListRef.value = [
        ...unref(fileListRef),
        {
          thumbUrl,
          ...commonItem
        }
      ]
    })
  } else {
    fileListRef.value = [...unref(fileListRef), commonItem]
  }
  return false
}

const handleChange = (uploadFile) => {
  fileList.value = uploadFile
  console.log('dddddddd')
  console.log(uploadFile)
  console.log(fileList)
}
const upload = async () => {
  // console.log('ttttttttttt')
  dialogVisible2.value = true
}
console.log(props.modelValue)
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

  <!-- <ElButton type="primary" :loading="loading" @click="upload">
    {{ t('exampleDemo.save') }}
  </ElButton> -->
  <ElDialog
    :fullscreen="false"
    destroy-on-close
    lock-scroll
    draggable
    v-model="dialogVisible2"
    :close-on-click-modal="false"
  >
    <ElUpload
      :previewFileList="files"
      :on-preview="handlePreview"
      :on-remove="handleRemove"
      :accept="getStringAccept"
      :auto-upload="false"
      :before-upload="beforeUpload"
      :show-upload-list="false"
      list-type="picture"
      @change="handleChange"
    >
      <el-button type="primary">{{ t('common.selectFile') }}</el-button>
      <template #tip>
        <div class="el-upload__tip"> jpg/png files with a size less than 500KB. </div>
      </template>
    </ElUpload>
    <!-- <div>
      <el-table>
        <el-table-column prop="thumb" label="略缩图">
          <template #default="scope">
            <el-image :src="scope.row.thumb" />
          </template>
        </el-table-column>

        <el-table-column prop="name" label="文件名" />

        <el-table-column prop="size" label="文件大小" />

        <el-table-column prop="status" label="状态" />

        <el-table-column label="操作">
          <template #default>
            <el-button size="small">下载</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div> -->
    <template #footer>
      <ElButton @click="dialogVisible2 = false">{{ t('common.cancel') }}</ElButton>
      <ElButton type="primary" @click="handleStartUpload">{{ t('common.startUpload') }}</ElButton>
      <ElButton type="primary" @click="save">{{ t('common.save') }}</ElButton>
    </template>
  </ElDialog>
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
