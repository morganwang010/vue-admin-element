<script setup lang="ts">
import { ContentWrap } from '@/components/ContentWrap'
import { useI18n } from '@/hooks/web/useI18n'
import { Table } from '@/components/Table'
import { getCustomerListApi } from '@/api/customer'
import { ContractTableData } from '@/api/contract/types'
import { ref, h, reactive, unref } from 'vue'
import { ElTag, ElButton } from 'element-plus'
import { useTable } from '@/hooks/web/useTable'
import { Pagination, TableColumn, TableSlotDefault } from '@/types/table'

const { register, tableObject, methods, elTableRef, paginationObj } = useTable<ContractTableData>({
  getListApi: getCustomerListApi,
  response: {
    message: 'message',
    data: 'data',
    list: 'list',
    total: 'total'
  }
})

const { getList } = methods

getList()

const { t } = useI18n()

const columns = reactive<TableColumn[]>([
  {
    field: 'id',
    label: t('customerTable.index'),
    type: 'index'
  },
  {
    field: 'name',
    label: t('customerTable.company')
  },
  {
    field: 'contact',
    label: t('customerTable.focal')
  },
  {
    field: 'mobilephone',
    label: t('customerTable.phone')
  },
  {
    field: 'created',
    label: t('customerTable.startDate')
  },
  {
    field: 'overTime',
    label: t('customerTable.endDate')
  },
  {
    field: 'domain',
    label: t('customerTable.domain')
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

const actionFn = (data: TableSlotDefault) => {
  console.log(data)
}

// const paginationObj = ref<Pagination>()

// paginationObj.value = {
//   total: tableObject.total
// }
console.log(tableObject)
console.log(tableObject.total)
</script>

<template>
  <ContentWrap>
    <Table
      v-model:pageSize="tableObject.pageSize"
      v-model:currentPage="tableObject.currentPage"
      :columns="columns"
      :data="tableObject.tableList"
      :loading="tableObject.loading"
      :pagination="paginationObj"
      @register="register"
    >
      <template #action="data">
        <ElButton type="primary" @click="actionFn(data as TableSlotDefault)">
          {{ t('contactTable.edit') }}
        </ElButton>
        <Icon
          @click="actionFn(data as TableSlotDefault)"
          class="ml-5px"
          icon="carbon:skill-level-advanced"
          :size="14"
        />
        <Icon class="ml-5px" icon="bi:question-circle-fill" :size="14" />
      </template>

      <template #expand="data">
        <div class="ml-30px">
          <div>{{ t('tableDemo.title') }}：{{ data.row.title }}</div>
          <div>{{ t('tableDemo.author') }}：{{ data.row.author }}</div>
          <div>{{ t('tableDemo.displayTime') }}：{{ data.row.display_time }}</div>
        </div>
      </template>
    </Table>
  </ContentWrap>
</template>