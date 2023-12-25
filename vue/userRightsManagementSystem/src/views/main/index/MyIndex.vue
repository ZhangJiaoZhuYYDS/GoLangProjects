<template>
  <div class="index1">
    <el-row :gutter="3">
      <el-col :span="8"><line-chart /></el-col>
      <el-col :span="8"><map-chart /></el-col>
      <el-col :span="8"><pie-chart /></el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import lineChart from './c-cpns/lineChart.vue'
import mapChart from './c-cpns/mapChart.vue'
import PieChart from './c-cpns/pieChart.vue'
</script>

<style scoped lang="less">
.index1 {
  width: 100%;
  height: 100%;
}
</style>
