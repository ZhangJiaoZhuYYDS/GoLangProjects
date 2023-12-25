import hyRequest from '..'

// 获取角色所有权限
export function getPermissionsOfRole(roleId: number) {
  return hyRequest.get({
    url: '/role/perms/' + roleId
  })
}
// 修改角色的权限
export function updateRolePermission(userIdAndPermissions: object) {
  return hyRequest.put({
    url: '/role/permission',

    data: userIdAndPermissions
  })
}
