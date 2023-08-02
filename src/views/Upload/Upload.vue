<script lang="ts">
import { Component, Prop } from 'vue-property-decorator';
import { Uploader } from 'element-plus';
import { Uploader as EUploader } from 'element-plus/src';

@Component
export default class UploaderComponent extends EUploader {
  @Prop() public url: string;
  @Prop() public multiple: boolean;
  @Prop() public asyncData: (file: File) => Promise<any> | undefined;
  @Prop() public formData: { [key: string]: string | Buffer | undefined } | undefined;
  @Prop() public preview: (file: File) => string | null | undefined;
  @Prop() public onFileChange: (file: File) => void;
  @Prop() public onProgress: (event: { loaded: number, total: number }) => void;
  @Prop() public onError: (error: Error) => void;
 

  mounted() {
    super.setOptions({
      url: this.url,
      multiple: this.multiple,
      asyncData: this.asyncData,
      formData: this.formData,
      preview: this.preview,
      onFileChange: this.onFileChange,
      onProgress: this.onProgress,
      onError: this.onError,
    });

    if (this.openDialog) {
      this.openDialog();
    }
  }

  async openDialog() {
    const dialog = await this.$refs.dialog;
    if (dialog) {
      await dialog.open();
    }
  }

  async upload(file: File) {
    if (this.upload) {
      await this.upload(file);
    }
  }
}
</script>

<template>
  <el-upload
    :file-list="fileList"
    :before-upload="beforeUpload"
    :on-success="onSuccess"
    :on-error="onError"
    :on-progress="onProgress"
    :multiple="multiple"
  >
    <template #default>
      <div>
        <el-button type="text" size="small" @click="openDialog">选择文件</el-button>
      </div>
    </template>
    <template #modified>
      <div>
        <el-button type="text" size="small" @click="openDialog">选择文件</el-button>
        <el-button type="text" size="small" @click="upload">开始上传</el-button>
      </div>
    </template>
  </el-upload>
</template>
