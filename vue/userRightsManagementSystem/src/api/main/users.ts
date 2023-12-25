import hyRequest from '..'

export function getAllSysUsers() {
  return hyRequest.get({
    url: '/sysUser/list'
  })
}

// 设置角色
export function setRole(userId: number, roleId: number) {
  return hyRequest.put({
    url: '/sysUser/role/' + userId + '/' + roleId
  })
}
