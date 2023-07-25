<script setup lang="ts">
import { ContentWrap } from '@/components/ContentWrap'
import { useI18n } from '@/hooks/web/useI18n'
import { Table } from '@/components/Table'
import { getTableListApi } from '@/api/contract'
import { ContractTableData } from '@/api/contract/types'
import { ref, h, reactive, unref } from 'vue'
import { ElTag, ElButton } from 'element-plus'
import { useTable } from '@/hooks/web/useTable'
import { Pagination, TableColumn, TableSlotDefault } from '@/types/table'

const { register, tableObject, methods, elTableRef } = useTable<ContractTableData>({
  getListApi: getTableListApi,
  response: {
    list: 'list',
    total: 'total'
  }
})

const { getList } = methods

getList()

const {
  register: register2,
  tableObject: tableObject2,
  methods: methods2
} = useTable<ContractTableData>({
  getListApi: getTableListApi,
  response: {
    list: 'list',
    total: 'total'
  }
})

const { getList: getList2 } = methods2

getList2()

const { t } = useI18n()

const columns = reactive<TableColumn[]>([
  {
    field: 'index',
    label: t('contactTable.index'),
    type: 'index'
  },
  {
    field: 'title',
    label: t('contactTable.company')
  },
  {
    field: 'author',
    label: t('contactTable.focal')
  },
  {
    field: 'display_time',
    label: t('contactTable.startDate')
  },
  {
    field: 'display_time',
    label: t('contactTable.endDate')
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
    field: 'pageviews',
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

const paginationObj = ref<Pagination>()
paginationObj.value = {
  total: tableObject.total
}
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
