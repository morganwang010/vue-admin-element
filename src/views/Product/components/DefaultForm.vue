<script setup lang="ts">
import { Form } from '@/components/Form'
import { reactive, ref, onMounted, computed, watch } from 'vue'
import { useI18n } from '@/hooks/web/useI18n'
import { useIcon } from '@/hooks/web/useIcon'
import { ContentWrap } from '@/components/ContentWrap'
import { useAppStore } from '@/store/modules/app'
import { FormSchema } from '@/types/form'
import { ComponentOptions } from '@/types/components'
import { useForm } from '@/hooks/web/useForm'
const appStore = useAppStore()

const { t } = useI18n()
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
const isMobile = computed(() => appStore.getMobile)

const schema1 = reactive<FormSchema[]>([
  {
    field: 'name',
    label: t('formDemo.input'),
    component: 'Input'
  },
  {
    field: 'image',
    label: t('formDemo.input'),
    component: 'Input'
  },
  {
    field: 'image',
    label: t('formDemo.input'),
    component: 'Upload',
    value: props.currentRow.image
  },
  {
    field: 'field64',
    component: 'Upload',
    label: t('formDemo.default')
  }
])
const { register, methods, elFormRef } = useForm({
  // schema: props.formSchema
  schema: schema1
})
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
watch(
  () => props.currentRow,
  (currentRow) => {
    if (!currentRow) return
    console.log(currentRow)
    const { setValues } = methods
    setValues(currentRow)
  },
  {
    deep: true,
    immediate: true
  }
)
</script>

<template>
  <ContentWrap :title="t('formDemo.defaultForm')" :message="t('formDemo.formDes')">
    <Form
      :schema="schema"
      label-width="auto"
      :label-position="isMobile ? 'top' : 'right'"
      @register="register"
    />
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
