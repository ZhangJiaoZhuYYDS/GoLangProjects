import { ref, computed, reactive } from 'vue'
import { defineStore } from 'pinia'

import { getAllRoles, getUserMenus } from '@/api/main/main'
import { getAllSysUsers } from '@/api/main/users'

export const useMainStore = defineStore('main', () => {
  const roles = reactive<any>([])
  const users = reactive<any>([])
  async function getAllRolesData() {
    const res = await getAllRoles()
    roles.length = 0
    roles.push(...res.data)
    console.log('store/main/main', roles)
  }

  async function getAllUsers() {
    const res = await getAllSysUsers()
    users.length = 0
    users.push(...res.data)
    // console.log('store/main/main', users)
  }

  return { roles, users, getAllRolesData, getAllUsers }
})
