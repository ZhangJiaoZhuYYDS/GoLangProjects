<template>
  <el-breadcrumb :separator-icon="ArrowRight">
    <template v-for="item in breadcrumbs" :key="item.name">
      <!-- v-for必须加key -->
      <el-breadcrumb-item>{{ item.name }}</el-breadcrumb-item>
    </template>
  </el-breadcrumb>
</template>

<script setup lang="ts">
import { ArrowRight } from '@element-plus/icons-vue'
import { computed, ref, watch } from 'vue'
import { useRoute } from 'vue-router'

import { useLoginStore } from '@/stores/login/login'
import { mapPathToBreadcrumbs } from '@/utils/map-menu'
const route = useRoute()
const { userMenu } = useLoginStore()

// userMenu

const breadcrumbs = computed(() => {
  console.log(route.path)
  return mapPathToBreadcrumbs(route.path, userMenu)
})
</script>

<style scoped lang="less">
.home {
  color: red;
}
</style>
