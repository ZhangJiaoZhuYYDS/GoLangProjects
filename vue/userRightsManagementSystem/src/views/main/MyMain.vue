<template>
  <div class="common-layout">
    <el-container>
      <el-aside width="200px">
        <main-menu />
      </el-aside>
      <el-container>
        <el-header>
          <main-header />
        </el-header>
        <el-main>
          <router-view></router-view>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'

import mainHeader from '@/components/menu-header/mainHeader.vue'
import mainMenu from '@/components/main-menu/MainMenu.vue'

// console.log(useRouter().getRoutes())

// const collapsed = ref<boolean>(false)
// const autowidth = computed(() => {
//   if (collapsed.value) {
//     return '64px'
//   } else {
//     return '200px'
//   }
// })
</script>

<style scoped lang="less">
.common-layout {
  display: flex;
  height: 100vh;
  overflow: hidden;
  .el-header {
    display: flex;
    flex-wrap: nowrap;
    align-items: center;
    background-color: #b3c0d1;
    color: var(--el-text-color-primary);
  }

  .el-aside {
    background-color: rgb(50, 65, 87);
    color: var(--el-text-color-primary);
    text-align: center;
    line-height: 200px;
    width: 200px;
    height: 100%;
    // overflow: hidden;
    // overflow: scroll;
  }
  // .el-main {
  //   padding: 0;
  // }
}
</style>
