<script setup lang="ts">
import {
  ElUpload,
  ElDialog,
  ElButton,
  ElTableColumn,
  ElTable,
  ElImage,
  ElIcon,
  ElPagination,
  ElInput
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
const baseUrl = import.meta.env.VITE_PRODUCT_IMAGE_BASE_URL
console.log(baseUrl)
const fileList = ref<UploadUserFile[]>()

const { t } = useI18n()

const props = defineProps({
  modelValue: propTypes.string.def(''),
  previewFileList: {
    type: Array as PropType<string[]>,
    default: () => []
  }
})
const thumbUrl = ref('')
const emit = defineEmits(['change', 'update:modelValue'])
// maxNumber: 3,
// maxSize: 1024,
const maxNumber = 3
const maxSize = 1024
watch(
  () => props.modelValue,
  (val: string) => {
    if (val === unref(thumbUrl)) return
    thumbUrl.value = props.modelValue
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
const loading = ref(false)
const dialogVisible2 = ref(false)
const upload = async () => {
  // console.log('ttttttttttt')
  dialogVisible2.value = true
}
const save = async () => {
  console.log('tttt')
  // thumbUrl.value = '1'
}
const message = useMessage()
//   是否正在上传
const isUploadingRef = ref(false)
const fileListRef = ref<FileItem[]>([])
const files = ref<string[]>([])
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
    const res = await uploadImageApi(formData)
    console.log(baseUrl)
    console.log(res.data.name)
    thumbUrl.value = baseUrl + '/upload/' + res.data.name
    isUploadingRef.value = false
    // 生产环境:抛出错误
    // const errorList = data.filter((item: any) => !item.success)
    // if (errorList.length > 0) throw errorList
  } catch (e) {
    isUploadingRef.value = false
    throw e
  }
}
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
const getStringAccept = 'file'

const handleRemove: UploadProps['onRemove'] = (file, uploadFiles) => {
  console.log(file, uploadFiles)
}

const handlePreview: UploadProps['onPreview'] = (uploadFile) => {
  console.log(uploadFile)
  console.log('00000000000000000000')
}
const handleChange = (uploadFile) => {
  fileList.value = uploadFile
}
</script>

<template>
  <div class="block">
    <!-- <el-input v-model="thumbUrl" />
    <span class="demonstration">Custom --{{ thumbUrl }}</span> -->
    <div class="leftImage"
      ><el-image :src="thumbUrl" style="width: 100px; height: 100px">
        <template #error>
          <div class="image-slot">
            <el-icon><icon-picture /></el-icon>
          </div>
        </template>
      </el-image>
    </div>
    <div class="upload"
      ><el-button type="primary" :loading="loading" @click="upload">
        {{ t('common.upload') }}
      </el-button>
    </div>
  </div>

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
.block {
  padding: 30px 0;
  text-align: center;
  border-right: solid 1px var(--el-border-color);
  /* display: inline-block; */
  /* width: 10%; */
  box-sizing: border-box;
  vertical-align: middle;
  display: flex;
  justify-content: space-between;
}
.upload {
  vertical-align: middle;
  align-items: center;
  display: flex;
}
/* .demo-image__error .demonstration {
  display: block;
  color: var(--el-text-color-secondary);
  font-size: 14px;
  margin-bottom: 20px;
}
.demo-image__error .el-image {
  padding: 0 5px;
  max-width: 300px;
  max-height: 200px;
  width: 150px;
  height: 50px;
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
} */
</style>
