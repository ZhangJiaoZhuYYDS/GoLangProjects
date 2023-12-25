<template>
  <div class="login">
    <div class="login-form">
      <header>
        <img src="@/assets/img/logo.png" alt="" />
        <h3>vue3-admin</h3>
      </header>
      <el-form id="login-el-form">
        <el-form-item v-model="loginForm">
          <el-input
            :prefix-icon="User"
            label="username"
            placeholder="username"
            v-model="loginForm.username"
            type="text"
          />
        </el-form-item>
        <el-form-item>
          <el-input
            :prefix-icon="Lock"
            label="password"
            placeholder="password"
            v-model="loginForm.password"
            type="password"
          />
        </el-form-item>
        <el-form-item prop="verifyCode">
          <el-input
            placeholder="验证码"
            v-model="loginForm.verifyCode"
            type="text"
            style="width: 40%; display: inline-block"
          />
          <div style="margin-left: 10px; display: inline-block; height: 33px">
            <img
              :src="codeUrl"
              @click="getValidCode"
              style="margin-bottom: -12px; width: 100%; height: 100%; object-fit: cover"
            />
          </div>
        </el-form-item>
        <el-form-item style="border: none; background: none">
          <el-button type="primary" style="width: 100%" @click="handleLogin">登录</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue'
import { User, Lock } from '@element-plus/icons-vue'
import { getCode } from '@/api/login/login'
import { useLoginStore } from '@/stores/login/login'
import { LOGIN_TOKEN } from '@/global/constants'
import { localCache } from '@/utils/cache'

const codeUrl = ref<string>()

const loginForm = reactive({
  username: '',
  password: '',
  uuid: '',
  verifyCode: ''
})

const getValidCode = () => {
  getCode().then((res) => {
    // console.log(res);
    codeUrl.value = res.data.image
    loginForm.uuid = res.data.uuid
  })
}

const handleLogin = () => {
  const loginStore = useLoginStore()
  loginStore.login(loginForm)
}

const handleTokenLogin = () => {
  // const token = localCache.getCache(LOGIN_TOKEN)
  // if (token != null) {
  //   const loginStore = useLoginStore()
  //   loginStore.loginToken(token)
  // }
}

onMounted(() => {
  getValidCode()
  handleTokenLogin()
})
</script>

<style scoped lang="less">
.login {
  width: 100vw;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: url('@/assets/img/login/1.jpg');
  .login-form {
    width: 400px;
    height: 400px;
    padding: 4vh;
    margin: calc(50vh - 200px) auto;
    background: url('@/assets/img/login/login_form.png');
    background-size: 100% 100%;
    border-radius: 10px;
    box-shadow: 0 2px 8px 0 rgba(7, 17, 27, 0.06);
    opacity: '0.2';
    header {
      display: flex;
      justify-content: center;
      align-items: center;
      margin-bottom: 20px;
      img {
        display: inline-block;
        width: 40px;
      }

      h3 {
        margin-bottom: 0;
        font-size: 24px;
        color: #fff;
        text-align: center;
      }
    }

    .el-form-item {
      display: flex;
      width: 270px;
      margin: 20px auto;
      flex-wrap: nowrap;
      // border: 1px solid rgba(255, 255, 255, 0.1);
      // background: rgba(0, 0, 0, 0.1);
      border-radius: 5px;
      color: #454545;
      // padding-left: 10px;
    }
  }
}
</style>

<style lang="less">
#login-el-form {
  .el-input__wrapper {
    background-color: transparent;
  }
}
</style>
