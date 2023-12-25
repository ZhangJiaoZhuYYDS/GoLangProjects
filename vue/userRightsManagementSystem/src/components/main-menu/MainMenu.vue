<template>
  <el-menu :collapse="collapse" :default-active="defaultActive">
    <menu-item :menuList="userMenu" />
  </el-menu>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { storeToRefs } from 'pinia'

import { useLoginStore } from '@/stores/login/login'
import menuItem from './MenuItem.vue'
import { useRoute } from 'vue-router'

const collapse = ref(false)
const loginStore = useLoginStore()
const { userMenu } = loginStore

const route = useRoute()

const defaultActive = route.path

// console.log(userMenu, 9)
</script>

<style scoped lang="less">
// .el-menu {
//   border-right: none;
//   user-select: none;
// }
</style>
