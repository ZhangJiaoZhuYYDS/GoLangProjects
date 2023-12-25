import { ref } from 'vue'
import { defineStore } from 'pinia'

import { chartData } from '@/api/main/chart'

export const useChartStore = defineStore('chart', () => {
  const res = ref()
  async function getChartData(url: string) {
    const temp: any = await chartData(url)
    res.value = temp
    console.log(res.value)
  }
  return { res, getChartData }
})
