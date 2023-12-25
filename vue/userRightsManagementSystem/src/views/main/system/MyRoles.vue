<template>
  <el-dialog v-model="state.showSetPermissionDialog" title="Tips" width="30%" draggable>
    <el-form>
      <el-form-item>
        <el-tree
          :data="state.permissionTree"
          show-checkbox
          node-key="permission"
          :default-checked-keys="state.permissions"
          :props="{ children: 'children', label: 'menuName' }"
          ref="permissionRef"
        />
      </el-form-item>
      <el-form-item>
        <el-button @click="handleResetClick">重置</el-button>
        <el-button type="primary" @click="handleConfirmClick"> 确认 </el-button>
      </el-form-item>
    </el-form>
  </el-dialog>

  <el-table :data="state.roles" style="width: 100%">
    <el-table-column prop="name" label="角色名称" width="180" />
    <el-table-column prop="description" label="权限介绍" width="400" />
    <el-table-column label="操作">
      <template #default="scope">
        <el-button @click="handleEditClick(scope.row)">编辑</el-button>
        <el-button>删除</el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<script setup lang="ts">
import { nextTick, reactive, ref } from 'vue'
import { getAllMenus, getAllRoles } from '@/api/main/main'
import { useMainStore } from '@/stores/main/main'
import type { ElTree } from 'element-plus'

import { getPermissionsOfRole, updateRolePermission } from '@/api/main/role'

const permissionRef = ref<InstanceType<typeof ElTree>>()

const state = reactive<any>({
  roles: [],
  showSetPermissionDialog: false,
  formData: {
    id: 0,
    name: '',
    description: ''
  },
  permissionTree: [],
  permissions: []
})

const mainStore = useMainStore()
state.roles.push(...mainStore.roles)

function handleEditClick(selectRole: any) {
  getAllMenus().then((res) => {
    // 获取权限菜单
    state.permissionTree = res.data
  })
  state.formData = JSON.parse(JSON.stringify(selectRole))
  state.showSetPermissionDialog = true
  nextTick(() => {
    getPermissionsOfRole(selectRole.id).then((result: any) => {
      // 获取当前用户角色权限
      state.permissions = result.data
    })
  })
}

// 清空checked-box
const handleResetClick = () => {
  permissionRef.value!.setCheckedKeys([], false)
}

//确认
async function handleConfirmClick() {
  state.showSetPermissionDialog = false
  let nodes = permissionRef.value!.getCheckedNodes(false, false)
  let halfNodes = permissionRef.value!.getHalfCheckedNodes()
  // console.log('nodes:', nodes, halfNodes)
  nodes.unshift.apply(nodes, halfNodes)
  // nodes.unshift(...halfNodes)
  console.log('nodes.unshift:', nodes)

  let permissions: InstanceType<typeof ElTree>[] = []
  nodes.forEach((node) => {
    if (node.permission) {
      permissions.push(node.permission)
    }
  })

  let vo = {
    roleId: state.formData.id,
    permissions: permissions
  }

  await updateRolePermission(vo)
  state.showSetPermissionDialog = false

  console.log('system/MyRoles/permissions:', permissions)
}
</script>

<style scoped lang="less">
.roles {
  color: red;
}
</style>
