<script setup lang="ts">
import { Form } from '@/components/Form'
import { PropType, reactive, ref, onMounted, computed, watch } from 'vue'
import { ProductTableData } from '@/api/product/types'
import { useI18n } from '@/hooks/web/useI18n'
import { useIcon } from '@/hooks/web/useIcon'
import { ContentWrap } from '@/components/ContentWrap'
import { useAppStore } from '@/store/modules/app'
import { FormSchema } from '@/types/form'
import { ComponentOptions } from '@/types/components'
import { useForm } from '@/hooks/web/useForm'

const appStore = useAppStore()

const { t } = useI18n()

const isMobile = computed(() => appStore.getMobile)

const props = defineProps({
  currentRow: {
    type: Object as PropType<Nullable<ProductTableData>>,
    default: () => null
  },
  formSchema: {
    type: Array as PropType<FormSchema[]>,
    default: () => []
  }
})

const { register, methods, elFormRef } = useForm({
  schema: props.formSchema
})

watch(
  () => props.currentRow,
  (currentRow) => {
    if (!currentRow) return
    const { setValues } = methods
    setValues(currentRow)
  },
  {
    deep: true,
    immediate: true
  }
)

const restaurants = ref<Recordable[]>([])
const querySearch = (queryString: string, cb: Fn) => {
  const results = queryString
    ? restaurants.value.filter(createFilter(queryString))
    : restaurants.value
  // call callback function to return suggestions
  cb(results)
}
const createFilter = (queryString: string) => {
  return (restaurant: Recordable) => {
    return restaurant.value.toLowerCase().indexOf(queryString.toLowerCase()) === 0
  }
}
const loadAll = () => {
  return [
    { value: 'vue', link: 'https://github.com/vuejs/vue' },
    { value: 'element', link: 'https://github.com/ElemeFE/element' },
    { value: 'cooking', link: 'https://github.com/ElemeFE/cooking' },
    { value: 'mint-ui', link: 'https://github.com/ElemeFE/mint-ui' },
    { value: 'vuex', link: 'https://github.com/vuejs/vuex' },
    { value: 'vue-router', link: 'https://github.com/vuejs/vue-router' },
    { value: 'babel', link: 'https://github.com/babel/babel' }
  ]
}
const handleSelect = (item: Recordable) => {
  console.log(item)
}
onMounted(() => {
  restaurants.value = loadAll()
})

const initials = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j']
const options = ref<ComponentOptions[]>(
  Array.from({ length: 1000 }).map((_, idx) => ({
    value: `Option ${idx + 1}`,
    label: `${initials[idx % 10]}${idx}`
  }))
)
const options2 = ref<ComponentOptions[]>(
  Array.from({ length: 10 }).map((_, idx) => {
    const label = idx + 1
    return {
      value: `Group ${label}`,
      label: `Group ${label}`,
      options: Array.from({ length: 10 }).map((_, idx) => ({
        value: `Option ${idx + 1 + 10 * label}`,
        label: `${initials[idx % 10]}${idx + 1 + 10 * label}`
      }))
    }
  })
)

const generateData = () => {
  const data: {
    value: number
    desc: string
    disabled: boolean
  }[] = []
  for (let i = 1; i <= 15; i++) {
    data.push({
      value: i,
      desc: `Option ${i}`,
      disabled: i % 4 === 0
    })
  }
  return data
}

const holidays = [
  '2021-10-01',
  '2021-10-02',
  '2021-10-03',
  '2021-10-04',
  '2021-10-05',
  '2021-10-06',
  '2021-10-07'
]

const isHoliday = ({ dayjs }) => {
  return holidays.includes(dayjs.format('YYYY-MM-DD'))
}

const schema = reactive<FormSchema[]>([
  {
    field: 'name1',
    label: t('formDemo.input'),
    component: 'Upload'
  },
  {
    field: 'name',
    label: t('formDemo.input'),
    component: 'Input'
  },
  {
    field: 'image',
    component: 'Input',
    label: t('formDemo.timeSelect')
  }
])
</script>

<template>
  <ContentWrap :title="t('formDemo.defaultForm')" :message="t('formDemo.formDes')">
    <Form
      :schema="schema"
      label-width="auto"
      :label-position="isMobile ? 'top' : 'right'"
      @register="register"
    >
      <template #field4-prefix>
        <Icon icon="ep:calendar" class="el-input__icon" />
      </template>
      <template #field4-suffix>
        <Icon icon="ep:calendar" class="el-input__icon" />
      </template>
    </Form>
  </ContentWrap>
</template>

<style lang="less" scoped>
.cell {
  height: 30px;
  padding: 3px 0;
  box-sizing: border-box;

  .text {
    position: absolute;
    left: 50%;
    display: block;
    width: 24px;
    height: 24px;
    margin: 0 auto;
    line-height: 24px;
    border-radius: 50%;
    transform: translateX(-50%);
  }

  &.current {
    .text {
      color: #fff;
      background: purple;
    }
  }

  .holiday {
    position: absolute;
    bottom: 0px;
    left: 50%;
    width: 6px;
    height: 6px;
    background: red;
    border-radius: 50%;
    transform: translateX(-50%);
  }
}
</style>
