<template>
  <el-dialog title="设置角色" v-model="state.showSetRoleDialog">
    <el-select @change="upRole" v-model="state.currentRole">
      <el-option
        v-for="item in roles"
        :key="item.id"
        :label="item.name"
        :value="item.id"
      ></el-option>
    </el-select>
  </el-dialog>

  <el-table :data="state.users" style="width: 100%">
    <el-table-column prop="username" label="用户名" width="180" />
    <el-table-column prop="roleName" label="用户权限" width="180" />
    <el-table-column label="操作">
      <template #default="scope">
        <el-button @click="handleEditClick(scope.row)">编辑</el-button>
      </template>
    </el-table-column>
  </el-table>

  <!-- 分页 -->
  <el-row style="float: right"> </el-row>
</template>

<script setup lang="ts">
import { ref, reactive, nextTick } from 'vue'
import type { Ref } from 'vue'
import { useMainStore } from '@/stores/main/main'
import { getAllRoles } from '@/api/main/main'
import { setRole } from '@/api/main/users'
import { storeToRefs } from 'pinia'
const state = reactive<any>({
  users: [],
  showSetRoleDialog: false,
  currentUserId: 0,
  currentRole: ''
})
interface Role {
  id: number
  name: string
}
const roles: Ref<Array<Role> | null> = ref(null)
// const users = reactive<any>([])
const { users } = storeToRefs(useMainStore())
state.users.push(...users.value)
// console.log('system/MyAccount', users)
async function handleEditClick(user: any) {
  const res = await getAllRoles()
  roles.value = res.data
  state.currentUserId = user.id
  state.showSetRoleDialog = true
  state.currentRole = user.roleName
  console.log('system/MyAccount', res)
}

async function upRole(roleId: any) {
  console.log(roleId)

  const res = await setRole(state.currentUserId, roleId)
  if (res.data === 1) {
    console.log('角色设置成功')
  }
  // state.users = []
  // const menuStore = useMainStore()
  // await menuStore.getAllUsers()
  // const { roles } = storeToRefs(useMainStore())
  const newRoleName = roles.value!.filter((item: any) => item.id == roleId)[0].name

  state.users.forEach((user: any) => {
    if (user.id == state.currentUserId) {
      user.roleName = newRoleName
      // user.roleName

      // console.log(user, 88)
    }
  })

  // state.users.push(...users.value)
  console.log(
    'system/MyAccountstate.users',
    roles.value,
    state.currentUserId,
    users.value,
    state.users
  )
}
</script>

<style scoped lang="less">
.account {
  color: red;
}
</style>
