<script setup lang="ts">
import {
  ElUpload,
  ElDialog,
  ElButton,
  ElTableColumn,
  ElTable,
  ElImage,
  ElPagination
} from 'element-plus'
import { defineComponent, PropType, ref, computed, unref, watch, onMounted } from 'vue'
import { propTypes } from '@/utils/propTypes'
import { useI18n } from '@/hooks/web/useI18n'
import { uploadImageApi } from '@/api/upload'

const { t } = useI18n()

const tableData = {}

const dialogVisible2 = ref(false)
const handleStartUpload = async () => {
  console.log('ssssssssss')
  const { maxNumber } = props
  if ((fileListRef.value.length + props.previewFileList?.length ?? 0) > maxNumber) {
    return createMessage.warning(t('component.upload.maxNumber', [maxNumber]))
  }
  try {
    isUploading
    Ref.value = true
    // 只上传不是成功状态的
    const uploadFileList =
      fileListRef.value.filter((item) => item.status !== UploadResultStatus.SUCCESS) || []
    const data = await Promise.all(
      uploadFileList.map((item) => {
        return uploadApiByItem(item)
      })
    )
    isUploadingRef.value = false
    // 生产环境:抛出错误
    const errorList = data.filter((item: any) => !item.success)
    if (errorList.length > 0) throw errorList
  } catch (e) {
    isUploadingRef.value = false
    throw e
  }
}
const files = ref<string[]>([])
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
  const { size, name } = file
  const { maxSize } = props
  // 设置最大值，则判断
  if (maxSize && file.size / 1024 / 1024 >= maxSize) {
    createMessage.error(t('component.upload.maxSizeMultiple', [maxSize]))
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
</script>

<template>
  <ElDialog
    :fullscreen="false"
    destroy-on-close
    lock-scroll
    draggable
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
