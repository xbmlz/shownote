<template>
  <div class="login-modal">
    <img class="duck" src="@/assets/duck.gif" alt="">
    <div class="login">
      <div class="header">
        <h2>授权登录</h2>
        <p>登录后会创建一个名为sn-repo的仓库，用来存储您的笔记文件。</p>
      </div>
      <ul class="content">
        <li>
          <el-button @click="auth" type="primary">
            <img src="@/assets/gitee.png" alt="LOGO">
            Gitee 登录
          </el-button>
        </li>
      </ul>
    </div>
    <img class="duck" src="@/assets/duck.gif" alt="">
  </div>

</template>

<script lang="ts">

import {defineComponent, onMounted} from "vue"

export default defineComponent({
  name: "Login",
  setup() {
    console.log(import.meta.env.MODE)
    // 申请授权
    const auth = () => {
      let clientId = import.meta.env.MODE === "development" ? 'f5763537e579c4f97c56c69b80489a17e250d8186e48efbe3e3fba4c4b6c9558': '919935005b4504c4db77424865a3223a6a7962cab43070e0b5e79f7f4ec7080a'
      let redirectUri = import.meta.env.MODE === "development" ? 'http://127.0.0.1:8000/user/auth/gitee' : 'https://shownote.cn/api/user/auth/gitee'
      let authUrl = 'https://gitee.com/oauth/authorize?client_id=' + clientId + '&redirect_uri=' + redirectUri + '&response_type=code'
      window.location.href = authUrl
    }
    return {
      auth
    }
  }
})
</script>

<style scoped lang="scss">
.login-modal {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  user-select: none;
  background: url("@/assets/login_bg.jpg") no-repeat center center;

  .login {
    width: 280px;
    border: 1px solid #ccc;
    padding: 50px 30px;
    box-shadow: 2px 2px 10px 10px #ddd;

    .header {
      h2 {
        font-size: 28px;
      }

      p {
        padding: 15px 4px 20px;
        font-size: 16px;
      }
    }

    ul {
      li {
        list-style-type: none;

        ::v-deep(.el-button) {
          width: 100%;


          & > span {
            display: flex;
            align-items: center;
            font-size: 16px;
          }

          img {
            position: relative;
            top: -1px;
            width: 20px;
            margin-right: 14px;
          }
        }
      }
    }

  }
}

</style>
