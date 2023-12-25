<template>
  <template v-for="menu in menuList" :key="menu.path">
    <el-sub-menu v-if="menu.children && menu.children.length > 1" :index="menu.path">
      <template #title>
        <el-icon class="title-color">
          <component :is="menu.meta!.icon" />
        </el-icon>
        <span class="title-color">{{ menu.name }}</span>
      </template>

      <menu-item :menu-list="menu.children" />
    </el-sub-menu>

    <el-menu-item v-else :index="menu.path" @click="handleItemClick(menu)">
      <el-icon class="title-color">
        <component :is="menu.meta!.icon" />
      </el-icon>
      <span class="title-color">{{ menu.name }}</span>
    </el-menu-item>
  </template>
</template>

<script setup lang="ts">
import type { RouteRecordRaw } from 'vue-router'
import { useRoute, useRouter } from 'vue-router'

defineProps<{
  menuList: RouteRecordRaw[]
}>()
const router = useRouter()
function handleItemClick(submenu: any) {
  const name = submenu.name
  // console.log(name)

  router.push({ name: name })
}
</script>

<style scoped lang="less">
.title-color {
  color: #b7bdc3;
}
.el-sub-menu {
  background-color: #001529;
}

.el-menu-item {
  padding-left: 50px !important;
  background-color: #0c2135;
}

.el-menu-item:hover {
  background-color: #fff;
}

.el-menu-item.is-active {
  background-color: #0a60bd;
}
</style>
