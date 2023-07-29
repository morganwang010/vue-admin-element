<script setup lang="ts">
import { ContentWrap } from '@/components/ContentWrap'
import { Dialog } from '@/components/Dialog'
import { useI18n } from '@/hooks/web/useI18n'
import { Table } from '@/components/Table'
import { getContractListApi } from '@/api/contract'
import { getDictOneApi } from '@/api/common'
import { ContractTableData } from '@/api/contract/types'
import { ref, h, reactive, unref } from 'vue'
import { ElTag, ElButton } from 'element-plus'
import { useTable } from '@/hooks/web/useTable'
import { TableColumn, TableSlotDefault } from '@/types/table'
import { Form, FormExpose } from '@/components/Form'
import { FormSchema } from '@/types/form'
import { useValidator } from '@/hooks/web/useValidator'

import dayjs from 'dayjs'
const { required } = useValidator()
const { register, tableObject, methods, elTableRef, paginationObj } = useTable<ContractTableData>({
  getListApi: getContractListApi,
  response: {
    message: 'message',
    data: 'data',
    list: 'list',
    total: 'total'
  }
})
const dialogVisible2 = ref(false)

const formRef = ref<FormExpose>()

const formSubmit = () => {
  unref(formRef)
    ?.getElFormRef()
    ?.validate((valid) => {
      if (valid) {
        console.log('submit success')
      } else {
        console.log('submit fail')
      }
    })
}
const { getList } = methods

getList()

const { t } = useI18n()

const columns = reactive<TableColumn[]>([
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

const schema = reactive<FormSchema[]>([
  {
    field: 'field1',
    label: t('formDemo.input'),
    component: 'Input',
    formItemProps: {
      rules: [required()]
    }
  },
  {
    field: 'field2',
    label: t('formDemo.select'),
    component: 'Select',
    componentProps: {
      options: [
        {
          label: 'option1',
          value: '1'
        },
        {
          label: 'option2',
          value: '2'
        }
      ]
    }
  },
  {
    field: 'field3',
    label: t('formDemo.radio'),
    component: 'Radio',
    componentProps: {
      options: [
        {
          label: 'option-1',
          value: '1'
        },
        {
          label: 'option-2',
          value: '2'
        }
      ]
    }
  },
  {
    field: 'field4',
    label: t('formDemo.checkbox'),
    component: 'Checkbox',
    value: [],
    componentProps: {
      options: [
        {
          label: 'option-1',
          value: '1'
        },
        {
          label: 'option-2',
          value: '2'
        },
        {
          label: 'option-3',
          value: '3'
        }
      ]
    }
  },
  {
    field: 'field5',
    component: 'DatePicker',
    label: t('formDemo.datePicker'),
    componentProps: {
      type: 'date'
    }
  },
  {
    field: 'field6',
    component: 'TimeSelect',
    label: t('formDemo.timeSelect')
  }
])
const getDictOne = async () => {
  const res = await getDictOneApi()
  if (res) {
    schema[1].componentProps!.options = res.data
  }
}
const actionFn = (data: TableSlotDefault) => {
  // schema[1].componentProps!.options = data
  console.log(data)
}

// getDictOne()
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
        <ElButton type="primary" v-hasPermi="['example:dialog:edit']" @click="action(row, 'edit')">
          {{ t('exampleDemo.edit') }}
        </ElButton>
        <ElButton type="primary" @click="actionFn(data as TableSlotDefault)">
          {{ t('contractTable.edit') }}
        </ElButton>
        <ElButton type="primary" @click="dialogVisible2 = !dialogVisible2">
          {{ t('dialogDemo.combineWithForm') }}
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

    <Dialog v-model="dialogVisible2" :title="t('dialogDemo.dialog')">
      <Form ref="formRef" :schema="schema" />
      <template #footer>
        <ElButton type="primary" @click="formSubmit">{{ t('dialogDemo.submit') }}</ElButton>
        <ElButton @click="dialogVisible2 = false">{{ t('dialogDemo.close') }}</ElButton>
      </template>
    </Dialog>
  </ContentWrap>
</template>
