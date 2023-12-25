import hyRequest from '..'

export function getUserMenus(roleId: number) {
  return hyRequest.get({
    url: '/sysMenu/selectUserMenu/' + roleId
  })
}

export function getAllRoles() {
  return hyRequest.get({
    url: '/role/list'
  })
}

export function getAllMenus() {
  return hyRequest.get({
    url: '/sysMenu/list'
  })
}
