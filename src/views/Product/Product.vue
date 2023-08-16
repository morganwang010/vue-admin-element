<script setup lang="ts">
import { ContentWrap } from '@/components/ContentWrap'
import { Search } from '@/components/Search'
import { Dialog } from '@/components/Dialog'
import { Upload } from '@/components/Upload'
import { useI18n } from '@/hooks/web/useI18n'
import { ElButton, ElTag, ElImage, ElUpload } from 'element-plus'
import { Table } from '@/components/Table'
import { getProductListApi, updateProductApi, createProductApi } from '@/api/product'
import { useTable } from '@/hooks/web/useTable'
import { ProductTableData } from '@/api/product/types'
import { h, ref, unref, reactive } from 'vue'
import Write from './components/Write.vue'
import DefaultForm from './components/DefaultForm.vue'
import Detail from './components/Detail.vue'
import { CrudSchema, useCrudSchemas } from '@/hooks/web/useCrudSchemas'
import { TableColumn } from '@/types/table'
import dayjs from 'dayjs'

const { register, tableObject, methods, elTableRef, paginationObj } = useTable<ProductTableData>({
  getListApi: getProductListApi,
  response: {
    message: 'message',
    data: 'data',
    list: 'list',
    total: 'total'
  }
})
// const { register, tableObject, methods } = useTable<TableData>({
//   getListApi: getTableListApi,
//   delListApi: delTableListApi,
//   response: {
//     list: 'list',
//     total: 'total'
//   },
//   defaultParams: {
//     title: 's'
//   }
// })

const { getList, setSearchParams } = methods

getList()

const { t } = useI18n()

const crudSchemas = reactive<CrudSchema[]>([
  {
    field: 'id',
    label: t('productTable.index'),
    type: 'index'
  },
  {
    field: 'name',
    label: t('productTable.name')
  },
  {
    field: 'unit',
    label: t('productTable.unit')
  },
  {
    field: 'price',
    label: t('productTable.price')
  },
  {
    field: 'description',
    label: t('productTable.describe')
  },
  {
    field: 'image',
    label: t('productTable.image'),
    formatter: (_: Recordable, __: TableColumn, cellValue: string) => {
      return h(ElImage, { src: cellValue }, () => cellValue)
    }
  },
  {
    field: 'created',
    label: t('productTable.createdDate')
    // formatter: (_: Recordable, __: TableColumn, cellValue: string) => {
    //   return h(() =>
    //     cellValue === '' ? t('productTable.null') : dayjs(cellValue).format('YYYY-MM-DD')
    //   )
    // }
  },
  {
    field: 'offdate',
    label: t('productTable.offDate'),
    formatter: (_: Recordable, __: TableColumn, cellValue: string) => {
      return h(() =>
        cellValue === '' ? t('productTable.null') : dayjs(cellValue).format('YYYY-MM-DD')
      )
    }
  },
  {
    field: 'importance',
    label: t('tableDemo.importance'),
    formatter: (_: Recordable, __: TableColumn, cellValue: number) => {
      return h(
        ElTag,
        {
          type: cellValue === 1 ? 'success' : cellValue === 2 ? 'warning' : 'danger'
        },
        () =>
          cellValue === 1
            ? t('tableDemo.important')
            : cellValue === 2
            ? t('tableDemo.good')
            : t('tableDemo.commonly')
      )
    }
  },
  {
    field: 'action',
    label: t('tableDemo.action')
  }
])

const { allSchemas } = useCrudSchemas(crudSchemas)

const dialogVisible = ref(false)
const dialogVisible2 = ref(false)
const dialogTitle = ref('')

const AddAction = () => {
  dialogTitle.value = t('exampleDemo.add')
  tableObject.currentRow = null
  dialogVisible.value = true
  actionType.value = ''
}

const delLoading = ref(false)

const delData = async (row: ProductTableData | null, multiple: boolean) => {
  tableObject.currentRow = row
  const { delList, getSelections } = methods
  const selections = await getSelections()
  delLoading.value = true
  await delList(
    multiple ? selections.map((v) => v.id) : [tableObject.currentRow?.id as string],
    multiple
  ).finally(() => {
    delLoading.value = false
  })
}

const actionType = ref('')

const action = (row: ProductTableData, type: string) => {
  dialogTitle.value = t(type === 'edit' ? 'exampleDemo.edit' : 'exampleDemo.detail')
  actionType.value = type
  tableObject.currentRow = row
  dialogVisible.value = true
  console.log(row)
}

const writeRef = ref<ComponentRef<typeof Write>>()
const defaultFormRef = ref<ComponentRef<typeof DefaultForm>>()
const loading = ref(false)
const save = async () => {
  // const write = unref(writeRef)
  const defaultForm = unref(defaultFormRef)
  await defaultForm?.elFormRef?.validate(async (isValid) => {
    // console.log(isValid)
    console.log('3333333333')
    // if (1 == 1) {
    if (isValid) {
      loading.value = true
      const data = (await defaultForm?.getFormData()) as ProductTableData
      console.log('dddddddddddddddfffffffffffff')
      console.log(data)
      let apiMethod

      if (actionType.value === 'edit') {
        apiMethod = updateProductApi
      } else {
        apiMethod = createProductApi
      }

      const res = await apiMethod(data)
        .catch(() => {})
        .finally(() => {
          loading.value = false
        })
      if (res) {
        dialogVisible.value = false
        tableObject.currentPage = 1
        getList()
      }
    } else {
      console.log('cccccccc')
    }
  })
}

const upload = async () => {
  // console.log('ttttttttttt')
  dialogVisible2.value = true
}
</script>

<template>
  <ContentWrap>
    <!-- <Search
      :model="{ title: 's' }"
      :schema="allSchemas.searchSchema"
      @search="setSearchParams"
      @reset="setSearchParams"
    /> -->

    <div class="mb-10px">
      <ElButton type="primary" @click="AddAction">{{ t('exampleDemo.add') }}</ElButton>
      <ElButton :loading="delLoading" type="danger" @click="delData(null, true)">
        {{ t('exampleDemo.del') }}
      </ElButton>
    </div>

    <Table
      v-model:pageSize="tableObject.pageSize"
      v-model:currentPage="tableObject.currentPage"
      :columns="allSchemas.tableColumns"
      :data="tableObject.tableList"
      :loading="tableObject.loading"
      :pagination="{
        total: tableObject.total
      }"
      @register="register"
    >
      <template #action="{ row }">
        <ElButton type="primary" v-hasPermi="['example:dialog:edit']" @click="action(row, 'edit')">
          {{ t('exampleDemo.edit') }}
        </ElButton>
        <ElButton
          type="success"
          v-hasPermi="['example:dialog:view']"
          @click="action(row, 'detail')"
        >
          {{ t('exampleDemo.detail') }}
        </ElButton>
        <ElButton type="danger" v-hasPermi="['example:dialog:delete']" @click="delData(row, false)">
          {{ t('exampleDemo.del') }}
        </ElButton>
      </template>
      <!-- 
      <template #image="{ row }">
        <ElImage :src="row.image" />
      </template> -->
      <!-- 如果预览视频时，需要用如下的代码 -->
      <!-- <template #video>
      <video width="180" height="180" controls>
        <source
          src="https://yiming_chang.gitee.io/pure-admin-doc/video/url.mov"
          type="video/mp4"
        />
      </video>
    </template> -->
    </Table>
  </ContentWrap>

  <Dialog v-model="dialogVisible" :title="dialogTitle">
    <!-- <Write
      v-if="actionType !== 'detail'"
      ref="writeRef"
      :form-schema="allSchemas.formSchema"
      :current-row="tableObject.currentRow"
    /> -->
    <DefaultForm
      v-if="actionType !== 'detail'"
      ref="defaultFormRef"
      :current-row="tableObject.currentRow"
    />
    <Detail
      v-if="actionType === 'detail'"
      :detail-schema="allSchemas.detailSchema"
      :current-row="tableObject.currentRow"
    />

    <template #footer>
      <ElButton v-if="actionType !== 'detail'" type="primary" :loading="loading" @click="save">
        8888888 {{ t('exampleDemo.save') }}
      </ElButton>
      <ElButton @click="dialogVisible = false">{{ t('dialogDemo.close') }}</ElButton>
    </template>
  </Dialog>
  <Upload v-if="dialogVisible2 === 'true'" title="上传文件" />
</template>
