<script setup lang="ts">
import { defineComponent, ref, watch, computed, PropType, unref } from 'vue'
import { ListItem } from './data'
import { useDesign } from '@/hooks/web/useDesign'
import { ElDescriptions, ElAvatar, ElTag } from 'element-plus'
import ElTypography from 'element-plus'
import { isNumber } from '@/utils/is'

const props = defineProps({
  list: {
    type: Array as PropType<ListItem[]>,
    default: () => []
  },
  pageSize: {
    type: [Boolean, Number] as PropType<boolean | number>,
    default: 5
  },
  currentPage: {
    type: Number,
    default: 1
  },
  titleRows: {
    type: Number,
    default: 1
  },
  descRows: {
    type: Number,
    default: 2
  },
  onTitleClick: {
    type: Function as PropType<(Recordable) => void>
  }
})

const getPagination = () => {
  const { list, pageSize } = props
  if (pageSize > 0 && list && list.length > pageSize) {
    return {
      total: list.length,
      pageSize,
      //size: 'small',
      current: unref(current),
      onChange(page) {
        current.value = page
        emit('update:currentPage', page)
      }
    }
  } else {
    return false
  }
}

const { prefixCls } = useDesign('header-notify-list')
const current = ref(props.currentPage || 1)
const getData = computed(() => {
  const { pageSize, list } = props
  if (pageSize === false) return []
  let size = isNumber(pageSize) ? pageSize : 5
  return list.slice(size * (unref(current) - 1), size * unref(current))
})
watch(
  () => props.currentPage,
  (v) => {
    current.value = v
  }
)

// export default defineComponent({
//   components: {
//     [ElAvatar.name]: ElAvatar,
//     [List.name]: List,
//     [List.Item.name]: List.Item,
//     AListItemMeta: List.Item.Meta,
//     ATypographyParagraph: Typography.Paragraph,
//     [ElTag.name]: ElTag
//   },
  // props: {
  //   list: {
  //     type: Array as PropType<ListItem[]>,
  //     default: () => []
  //   },
  //   pageSize: {
  //     type: [Boolean, Number] as PropType<Boolean | Number>,
  //     default: 5
  //   },
  //   currentPage: {
  //     type: Number,
  //     default: 1
  //   },
  //   titleRows: {
  //     type: Number,
  //     default: 1
  //   },
  //   descRows: {
  //     type: Number,
  //     default: 2
  //   },
  //   onTitleClick: {
  //     type: Function as PropType<(Recordable) => void>
  //   }
//   // },
//   emits: ['update:currentPage'],
//   setup(props, { emit }) {
//     const { prefixCls } = useDesign('header-notify-list')
//     const current = ref(props.currentPage || 1)
//     const getData = computed(() => {
//       const { pageSize, list } = props
//       if (pageSize === false) return []
//       let size = isNumber(pageSize) ? pageSize : 5
//       return list.slice(size * (unref(current) - 1), size * unref(current))
//     })
//     watch(
//       () => props.currentPage,
//       (v) => {
//         current.value = v
//       }
//     )
//     const isTitleClickable = computed(() => !!props.onTitleClick)
//     const getPagination = computed(() => {
//       const { list, pageSize } = props
//       if (pageSize > 0 && list && list.length > pageSize) {
//         return {
//           total: list.length,
//           pageSize,
//           //size: 'small',
//           current: unref(current),
//           onChange(page) {
//             current.value = page
//             emit('update:currentPage', page)
//           }
//         }
//       } else {
//         return false
//       }
//     })

//     function handleTitleClick(item: ListItem) {
//       props.onTitleClick && props.onTitleClick(item)
//     }

//     return {
//       prefixCls,
//       getPagination,
//       getData,
//       handleTitleClick,
//       isTitleClickable
//     }
//   }
// })
</script>
<template>
  <el-descriptions :class="prefixCls" bordered :pagination="getPagination">
    <template v-for="item in getData" :key="item.id">
      <el-descriptions-item class="list-item">
        <template #title>
          <div class="title">
            html
            <el-typography
              :type="isTitleClickable ? 'link' : ''"
              :delete="!!item.titleDelete"
              :ellipsis="titleEllipsisOptions"
              :onClick="handleTitleClick(item)"
              style="width: 100%; margin-bottom: 0 !important; cursor: pointer"
            >
              {{ item.title }}
            </el-typography>
            <div class="extra" v-if="item.extra">
              <el-tag class="tag" :color="item.color">
                {{ item.extra }}
              </el-tag>
            </div>
          </div>
        </template>

        <template #avatar>
          <el-avatar v-if="item.avatar" class="avatar" :src="item.avatar" />
          <span v-else> {{ item.avatar }}</span>
        </template>

        <template #description>
          <div>
            <div class="description" v-if="item.description">
              <el-typography
                style="width: 100%; margin-bottom: 0 !important"
                :rows="$props.descRows && $props.descRows > 0 ? $props.descRows : null"
              >
                <template #default="{ ellipsis }">
                  <el-tooltip v-if="ellipsis" content="item.description">
                    {{ item.description }}
                  </el-tooltip>
                  <span v-else>{{ item.description }}</span>
                </template>
              </el-typography>
            </div>
            <div class="datetime">
              {{ item.datetime }}
            </div>
          </div>
        </template>
      </el-descriptions-item>
    </template>
  </el-descriptions>
</template>

<style lang="less" scoped>
@prefix-cls: ~'@{namespace}-header-notify-list';

.@{prefix-cls} {
  &::-webkit-scrollbar {
    display: none;
  }

  ::v-deep(.ant-pagination-disabled) {
    display: inline-block !important;
  }

  &-item {
    padding: 6px;
    overflow: hidden;
    transition: all 0.3s;
    cursor: pointer;

    .title {
      margin-bottom: 8px;
      font-weight: normal;

      .extra {
        margin-top: -1.5px;
        margin-right: 0;
        float: right;
        font-weight: normal;

        .tag {
          margin-right: 0;
        }
      }

      .avatar {
        margin-top: 4px;
      }

      .description {
        font-size: 12px;
        line-height: 18px;
      }

      .datetime {
        margin-top: 4px;
        font-size: 12px;
        line-height: 18px;
      }
    }
  }
}
</style>
