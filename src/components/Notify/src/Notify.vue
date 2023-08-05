<script setup lang="ts">
import { Icon } from '@/components/Icon'
import { ElPopover, ElButton, ElBadge, ElAvatar, ElTable, ElTableColumn } from 'element-plus'
import type { TabsPaneContext } from 'element-plus'
import { ElTabs, ElTabPane, ElIcon } from 'element-plus'
import { propTypes } from '@/utils/propTypes'
import { useDesign } from '@/hooks/web/useDesign'
import NoticeList from './NoticeList.vue'
import { Bell } from '@element-plus/icons-vue'
const { getPrefixCls } = useDesign()

const props = defineProps({
  color: propTypes.string.def('')
})
const gridData = [
  {
    date: '2016-05-02',
    name: 'Jack',
    address: 'New York City'
  },
  {
    date: '2016-05-04',
    name: 'Jack',
    address: 'New York City'
  },
  {
    date: '2016-05-01',
    name: 'Jack',
    address: 'New York City'
  },
  {
    date: '2016-05-03',
    name: 'Jack',
    address: 'New York City'
  }
]
</script>
<template>
  <!-- <div :class="prefixCls"> -->
  <div>
    <el-popover placement="top-start" :width="200" trigger="hover" :color="color">
      <template #reference>
        <el-badge :value="2" class="item" type="warning">
          <el-icon :size="20" color="#409EFC" class="no-inherit">
            <Bell />
          </el-icon>
        </el-badge>
      </template>
      <el-table :data="gridData">
        <el-table-column width="150" property="date" label="date" />
        <el-table-column width="100" property="name" label="name" />
        <el-table-column width="300" property="address" label="address" />
      </el-table>
    </el-popover>
  </div>
</template>
