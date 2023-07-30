<script setup lang="ts">
import { ContentWrap } from '@/components/ContentWrap'
import { Search } from '@/components/Search'
import { Dialog } from '@/components/Dialog'
import { useI18n } from '@/hooks/web/useI18n'
import { ElButton, ElTag } from 'element-plus'
import { Table } from '@/components/Table'
import { getContractListApi, updateContractApi, createContractApi } from '@/api/contract'
import { useTable } from '@/hooks/web/useTable'
import { ContractTableData } from '@/api/contract/types'
import { h, ref, unref, reactive } from 'vue'
import Write from './components/Write.vue'
import Detail from './components/Detail.vue'
import { CrudSchema, useCrudSchemas } from '@/hooks/web/useCrudSchemas'
import { TableColumn } from '@/types/table'
import dayjs from 'dayjs'

const { register, tableObject, methods, elTableRef, paginationObj } = useTable<ContractTableData>({
  getListApi: getContractListApi,
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
    label: t('contractTable.index'),
    type: 'index'
  },
  {
    field: 'cname',
    label: t('contractTable.company')
  },
  {
    field: 'contact',
    label: t('contractTable.focal')
  },
  {
    field: 'mobilephone',
    label: t('contractTable.phone')
  },
  {
    field: 'beginTime',
    label: t('contractTable.startDate'),
    formatter: (_: Recordable, __: TableColumn, cellValue: string) => {
      return h(() =>
        cellValue === '' ? t('contractTable.null') : dayjs(cellValue).format('YYYY-MM-DD')
      )
    }
  },
  {
    field: 'overTime',
    label: t('contractTable.endDate'),
    formatter: (_: Recordable, __: TableColumn, cellValue: string) => {
      return h(() =>
        cellValue === '' ? t('contractTable.null') : dayjs(cellValue).format('YYYY-MM-DD')
      )
    }
  },
  {
    field: 'remarks',
    label: t('contractTable.remarks')
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
    field: 'amount',
    label: t('contactTable.mount')
  },
  {
    field: 'action',
    label: t('tableDemo.action')
  }
])

const { allSchemas } = useCrudSchemas(crudSchemas)

const dialogVisible = ref(false)

const dialogTitle = ref('')

const AddAction = () => {
  dialogTitle.value = t('exampleDemo.add')
  tableObject.currentRow = null
  dialogVisible.value = true
  actionType.value = ''
}

const delLoading = ref(false)

const delData = async (row: ContractTableData | null, multiple: boolean) => {
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

const action = (row: ContractTableData, type: string) => {
  dialogTitle.value = t(type === 'edit' ? 'exampleDemo.edit' : 'exampleDemo.detail')
  actionType.value = type
  tableObject.currentRow = row
  dialogVisible.value = true
}

const writeRef = ref<ComponentRef<typeof Write>>()

const loading = ref(false)

const save = async () => {
  const write = unref(writeRef)
  await write?.elFormRef?.validate(async (isValid) => {
    if (isValid) {
      loading.value = true
      const data = (await write?.getFormData()) as ContractTableData
      console.log(actionType.value)

      if (actionType.value === 'edit') {
        const params = {
          id: data.id,
          name: data.name,
          amount: data.amount,
          beginTime: data.beginTime,
          overTime: data.overTime,
          remarks: data.remarks,
          cname: data.cname,
          status: data.status,
          cid: '77',
          productList: 'fff'
        }
        const res = await updateContractApi(data)
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
        const res = await createContractApi(data)
          .catch(() => {})
          .finally(() => {
            loading.value = false
          })
        if (res) {
          dialogVisible.value = false
          tableObject.currentPage = 1
          getList()
        }
      }

      // const res = await apiMethod(data)
      //   .catch(() => {})
      //   .finally(() => {
      //     loading.value = false
      //   })
      // if (res) {
      //   dialogVisible.value = false
      //   tableObject.currentPage = 1
      //   getList()
      // }
    }
  })
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
    </Table>
  </ContentWrap>

  <Dialog v-model="dialogVisible" :title="dialogTitle">
    <Write
      v-if="actionType !== 'detail'"
      ref="writeRef"
      :form-schema="allSchemas.formSchema"
      :current-row="tableObject.currentRow"
    />

    <Detail
      v-if="actionType === 'detail'"
      :detail-schema="allSchemas.detailSchema"
      :current-row="tableObject.currentRow"
    />

    <template #footer>
      <ElButton v-if="actionType !== 'detail'" type="primary" :loading="loading" @click="save">
        {{ t('exampleDemo.save') }}
      </ElButton>
      <ElButton @click="dialogVisible = false">{{ t('dialogDemo.close') }}</ElButton>
    </template>
  </Dialog>
</template>
