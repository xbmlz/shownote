<template>
  <div class="login-modal">
    <div class="login">
      <div class="header">
        <svg-icon icon-class="logo" color="#24292f"></svg-icon>
        <h2>ShowNote</h2>
      </div>
      <div class="content">
        <p class="title">请选择快捷登录方式</p>
        <ul>
          <li v-for="item in list" :key="item.value" @click="login(item.value)">
            <svg-icon :icon-class="item.icon" :color="item.color"></svg-icon>
            <span>{{ item.label }}</span>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script lang="ts">

import {defineComponent, onMounted} from "vue"

export default defineComponent({
  name: "Login",
  data() {
    return {
      list: [
        {
          icon: 'gitee',
          color: '#c71d23',
          label: 'Gitee',
          value: 'gitee',
        }
      ]
    }
  },
  methods: {
    // 申请授权
    auth() {
      let clientId = import.meta.env.MODE === "development" ? 'f5763537e579c4f97c56c69b80489a17e250d8186e48efbe3e3fba4c4b6c9558' : '919935005b4504c4db77424865a3223a6a7962cab43070e0b5e79f7f4ec7080a'
      let redirectUri = import.meta.env.MODE === "development" ? 'http://127.0.0.1:8000/user/auth/gitee' : 'https://shownote.cn/api/user/auth/gitee'
      let authUrl = 'https://gitee.com/oauth/authorize?client_id=' + clientId + '&redirect_uri=' + redirectUri + '&response_type=code'
      window.location.href = authUrl
    },
    login(name) {
      switch (name) {
        case 'gitee':
          this.auth();
          break
        case 'github':
          this.$message.info('暂不支持Github登录');
          break;
      }
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
  color: rgb(36, 42, 49);
  width: 100%;
  margin: 0px;
  display: flex;
  padding: 0px;
  background: rgb(245, 247, 249);
  min-height: 100vh;
  flex-direction: column;
  -webkit-box-orient: vertical;
  -webkit-box-direction: normal;

  .login {
    border: 1px solid #d4dadf;
    overflow: hidden;
    width: 400px;
    background: #fff;
    border-radius: 4px;


    .header {
      margin: 0;
      display: flex;
      align-items: center;
      justify-content: space-between;
      padding: 24px 0;
      flex-direction: column;

      .svg-icon {
        font-size: 100px;
      }

      h2 {
        font-size: 24px;
      }
    }

    .content {
      padding: 24px 0 50px;
      border-top: 1px solid #d4dadf;

      .title {
        text-align: center;
        margin-bottom: 10px;
      }

      ul {
        display: flex;
        align-items: center;
        justify-content: center;
      }

      li {
        display: flex;
        align-items: center;
        justify-content: center;
        flex-direction: column;
        width: 100px;
        cursor: pointer;

        &:hover {
          .svg-icon, span {
            opacity: 1;
          }
        }

        .svg-icon {
          font-size: 60px;
          margin: 20px 0;
          opacity: 0.8;
        }

        span {
          font-size: 16px;
          opacity: 0.8;
        }
      }
    }
  }
}

</style>
