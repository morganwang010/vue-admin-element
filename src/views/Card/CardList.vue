<script setup lang="ts">
import Card from './components/Card.vue'
import { getCardList } from '@/api/cardList'
import { message } from '@/utils/message'
import { ElMessageBox } from 'element-plus'
import { ElRow, ElCol, ElBacktop, ElDropdown, ElCard } from 'element-plus'
import { ref, onMounted, nextTick } from 'vue'
import dialogForm from './components/DialogForm.vue'
import { useRenderIcon } from '@/components/ReIcon/src/hooks'
import Search from '@iconify-icons/ep/search'
import AddFill from '@iconify-icons/ri/add-circle-line'

defineOptions({
  name: 'ListCard'
})

const svg = `
        <path class="path" d="
          M 30 15
          L 28 17
          M 25.61 25.61
          A 15 15, 0, 0, 1, 15 30
          A 15 15, 0, 1, 1, 27.99 7.5
          L 15 15
        " style="stroke-width: 4px; fill: rgba(0, 0, 0, 0)"/>
      `

const INITIAL_DATA = {
  name: '',
  status: '',
  description: '',
  type: '',
  mark: ''
}

const pagination = ref({ current: 1, pageSize: 12, total: 0 })

const productList = ref([])
const dataLoading = ref(true)

const getCardListData = async () => {
  try {
    const { data } = await getCardList()
    productList.value = data.list
    pagination.value = {
      ...pagination.value,
      total: data.list.length
    }
  } catch (e) {
    console.log(e)
  } finally {
    setTimeout(() => {
      dataLoading.value = false
    }, 500)
  }
}

onMounted(() => {
  getCardListData()
})

const formDialogVisible = ref(false)
const formData = ref({ ...INITIAL_DATA })
const searchValue = ref('')

const onPageSizeChange = (size: number) => {
  pagination.value.pageSize = size
  pagination.value.current = 1
}
const onCurrentChange = (current: number) => {
  pagination.value.current = current
}
const handleDeleteItem = (product) => {
  ElMessageBox.confirm(
    product ? `确认删除后${product.name}的所有产品信息将被清空, 且无法恢复` : '',
    '提示',
    {
      type: 'warning'
    }
  )
    .then(() => {
      message('删除成功', { type: 'success' })
    })
    .catch(() => {})
}
const handleManageProduct = (product) => {
  formDialogVisible.value = true
  nextTick(() => {
    formData.value = { ...product, status: product?.isSetup ? '1' : '0' }
  })
}
</script>

<template>
  <div class="main">
    <div
      v-loading="dataLoading"
      :element-loading-svg="svg"
      element-loading-svg-view-box="-10, -10, 50, 50"
    >
      <!-- <el-empty
        description="暂无数据"
        v-show="
          productList
            .slice(
              pagination.pageSize * (pagination.current - 1),
              pagination.pageSize * pagination.current
            )
            .filter((v) => v.name.toLowerCase().includes(searchValue.toLowerCase())).length === 0
        "
      /> -->
      <template v-if="pagination.total > 0">
        <el-row :gutter="16">
          <el-col
            v-for="(product, index) in productList.slice(
              pagination.pageSize * (pagination.current - 1),
              pagination.pageSize * pagination.current
            )"
            :key="index"
            :xs="24"
            :sm="12"
            :md="8"
            :lg="6"
            :xl="4"
          >
            <el-card :body-style="{ padding: '0px' }">
              <img
                src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                class="image"
              />
              <div>
                <a :href="product.url">
                  <p class="list-card-item_detail--name text-text_color_primary">
                    {{ product.url }}
                  </p>
                  <p class="list-card-item_detail--desc text-text_color_regular">
                    {{ product.description }}
                  </p>
                </a>
              </div>
            </el-card>
            <!-- <Card
              :product="product"
              @delete-item="handleDeleteItem"
              @manage-product="handleManageProduct"
            /> -->
          </el-col>
        </el-row>
        <!-- <el-pagination
          class="float-right"
          v-model:currentPage="pagination.current"
          :page-size="pagination.pageSize"
          :total="pagination.total"
          :page-sizes="[12, 24, 36]"
          :background="true"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="onPageSizeChange"
          @current-change="onCurrentChange"
        /> -->
      </template>
    </div>
  </div>
</template>
