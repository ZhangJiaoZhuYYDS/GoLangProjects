import { ref, reactive } from 'vue'
import { defineStore, storeToRefs } from 'pinia'
import router from '@/router'
import { ElMessage } from 'element-plus'
import 'element-plus/es/components/message/style/css'

import { login as loginApi, loginToken as loginTokenApi } from '@/api/login/login'
// import { useMainStore } from '../main/main.ts'
import { useMainStore } from '../main/main'
import { localCache } from '@/utils/cache'
import { LOGIN_TOKEN, MENU_LIST, ROLE_ID } from '@/global/constants'
import type { RouteRecordRaw } from 'vue-router'
import { getUserMenus } from '@/api/main/main'
import { initMenu } from '@/utils/map-menu'

export const useLoginStore = defineStore('login', () => {
  const token = ref('')
  let userInfo = reactive({
    avatar: '',
    username: '',
    roleName: '',
    status: 0
  })
  let userMenu = ref<RouteRecordRaw[]>([])
  const menuList = ref<any>()
  const roles = ref<number>()

  async function login(requestUser: API.loginForm) {
    const result = await loginApi(requestUser)
    if (result.data.status) {
      ElMessage({
        message: '登录成功',
        type: 'success'
      })
    }
    console.log(result)
    //个人信息
    localCache.setCache(LOGIN_TOKEN, result.data.token)
    localCache.setCache(ROLE_ID, result.data.roleId)

    //请求菜单信息
    menuList.value = await getUserMenus(result.data.roleId)
    console.log('menuList', menuList.value.data)
    localCache.setCache(MENU_LIST, menuList.value.data)

    //初始化菜单信息
    userMenu.value = initMenu(menuList.value.data)
    console.log('store/login', userMenu)
    userMenu.value.forEach((item) => {
      item.children?.forEach((subitem) => {
        subitem.children = []
        router.addRoute('main', subitem)
      })
    })

    const menuStore = useMainStore()
    //获取所有角色信息
    menuStore.getAllRolesData()
    //获取可以登录的用户信息
    menuStore.getAllUsers()

    router.push('/main')
  }

  function loadLocalCacheAction() {
    token.value = localCache.getCache(LOGIN_TOKEN)
    roles.value = localCache.getCache(ROLE_ID)
    menuList.value = localCache.getCache(MENU_LIST)
    if (token.value && roles.value && menuList.value) {
      //请求菜单信息
      // const menuList = await getUserMenus(roles.value)
      // console.log('menuList', menuList.data)

      //初始化菜单信息
      userMenu.value = initMenu(menuList.value)
      console.log('stores/login.ts', userMenu.value)
      // 动态添加路由

      // userMenu.value.forEach((item) => {
      //   if (item.children) {
      //     if (item.children.length == 1) {
      //       item.redirect = item.path + '/' + item.children[0].path
      //       item.meta = item.children[0].meta
      //       router.addRoute('main', item)
      //     } else {
      //       item.children.forEach((subitem) => {
      //         subitem.children = []
      //         router.addRoute('main', subitem)
      //       })
      //     }
      //   }
      // })

      userMenu.value.forEach((item) => {
        item.children?.forEach((subitem) => {
          subitem.children = []
          router.addRoute('main', subitem)
        })
      })

      const menuStore = useMainStore()
      //获取所有角色信息
      menuStore.getAllRolesData()
      //获取可以登录的用户信息
      menuStore.getAllUsers()
    }
  }

  return { token, userInfo, roles, userMenu, login, loadLocalCacheAction }
})
