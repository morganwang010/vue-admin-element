<script setup lang="ts">
import { Icon } from '@/components/Icon'
import { ElPopover, ElButton, ElBadge, ElAvatar, ElTable, ElTableColumn } from 'element-plus'
import type { TabsPaneContext } from 'element-plus'
import { ElTabs, ElTabPane, ElIcon } from 'element-plus'
import { propTypes } from '@/utils/propTypes'
import { useDesign } from '@/hooks/web/useDesign'
import NoticeList from './NoticeList.vue'
import { Bell } from '@element-plus/icons-vue'
import { Popover, Tabs, Badge, TabPane } from 'ant-design-vue'
import { BellOutlined } from '@ant-design/icons-vue'
import { ref } from 'vue'
const { getPrefixCls } = useDesign()
const activeKey = ref('1')
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
  }
]
const buttonWidth = 70
</script>
<template>
  <!-- <div :class="prefixCls"> -->
  <div>
    <Popover placement="bottom">
      <Badge :count="5">
        <BellOutlined />
      </Badge>

      <template #content>
        <Tabs v-model:activeKey="activeKey">
          <TabPane key="1" tab="Tab 1">Content of Tab Pane 1</TabPane>
          <TabPane key="2" tab="Tab 2" force-render>Content of Tab Pane 2</TabPane>
          <TabPane key="3" tab="Tab 3">Content of Tab Pane 3</TabPane>
        </Tabs>
      </template>
      <template #title>
        <span>Title333</span>
      </template>
    </Popover>
  </div>
</template>
<style>
#components-popover-demo-placement .ant-btn {
  margin-left: 0;
  margin-right: 8px;
  margin-bottom: 8px;
  width: 70px;
  text-align: center;
  padding: 0;
}
#components-popover-demo-placement .ant-btn {
  margin-left: 8px;
  margin-right: 0;
}
</style>
