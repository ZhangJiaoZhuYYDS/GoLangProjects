import { createRouter, createWebHashHistory, createWebHistory } from 'vue-router'

import { localCache } from '@/utils/cache'
import { LOGIN_TOKEN, ROLE_ID } from '@/global/constants'

import { firstMenu } from '@/utils/map-menu'

const router = createRouter({
  // history: createWebHistory(import.meta.env.BASE_URL),
  history: createWebHashHistory(),
  routes: [
    {
      path: '/',
      redirect: '/main'
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/login/MyLogin.vue')
    },
    {
      path: '/main',
      name: 'main',
      component: () => import('../views/main/MyMain.vue')
    },
    {
      path: '/:pathMatch(.*)',
      component: () => import('../views/not-found/NotFound.vue')
    }
  ]
})

router.beforeEach((to) => {
  const token = localCache.getCache(LOGIN_TOKEN)

  if (to.path.startsWith('/main') && !token) {
    return '/login'
  }

  // 如果是进入到main
  if (to.path === '/main') {
    return firstMenu?.path
  }
})

export default router
