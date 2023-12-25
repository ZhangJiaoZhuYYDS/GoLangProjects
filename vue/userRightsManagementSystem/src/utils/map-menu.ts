const modules = import.meta.glob('../views/main/**/*.vue')

export let firstMenu: any = null
export const initMenu = (menu: any) => {
  menu.forEach((item: any) => {
    item.children.forEach((el: any) => {
      el.redirect = null
      el.path = '/main' + el.path
      if (el.component === 'Layout') {
        // el.component = () => import('@/views/main/MyMain.vue')
      } else {
        if (el.component == 'index/Index') {
          el.component = () => import('@/views/main/index/MyIndex.vue')
          // console.log(modules, 99, `@/views/main/${el.component}.vue`)
        } else if (el.component == 'system/Group') {
          el.component = () => import('@/views/main/system/MyRoles.vue')
          // console.log('ultils/map-menu', el.component)
        } else {
          el.component = el.component.split('/')[0] + '/My' + el.component.split('/')[1]
          el.component = modules[`../views/main/${el.component}.vue`]
          console.log('ultils/map-menu', modules, `../views/main/${el.component}.vue`)
        }
        if (!firstMenu) {
          firstMenu = el
        }
      }
    })
    if (item.children.length == 1) {
      item.path = item.children[0].path
      item.name = item.children[0].name
      item.meta = item.children[0].meta
    }
  })

  // menu.forEach((el: any) => {
  //   el.redirect = null
  //   el.path = '/main' + el.path
  //   if (el.component === 'Layout') {
  //     // el.component = () => import('@/views/main/MyMain.vue')
  //   } else {
  //     if (el.component == 'index/Index') {
  //       el.component = () => import('@/views/main/index/MyIndex.vue')
  //       // console.log(modules, 99, `@/views/main/${el.component}.vue`)
  //     } else if (el.component == 'system/Group') {
  //       el.component = () => import('@/views/main/system/MyRoles.vue')
  //     } else {
  //       el.component = el.component.split('/')[0] + '/My' + el.component.split('/')[1]
  //       el.component = modules[`../views/main/${el.component}.vue`]
  //     }
  //     if (!firstMenu) {
  //       firstMenu = el
  //     }
  //   }

  //   if (el.children != null && el.children.length) {
  //     initMenu(el.children)
  //   }
  // })

  return menu
}
interface IBreadCrumbs {
  id?: number
  path: string
  name: string
}

export function mapPathToBreadcrumbs(path: string, userMenu: any[]) {
  const breadCrumbs: IBreadCrumbs[] = []
  for (const menu of userMenu) {
    for (const subMenu of menu.children) {
      if (path == subMenu.path) {
        breadCrumbs.push({ path: menu.path, name: menu.name })

        if (menu.children.length > 1) {
          breadCrumbs.push({ path: subMenu.path, name: subMenu.name })
        }
      }
    }
  }

  return breadCrumbs
}
