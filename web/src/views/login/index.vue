<template>
  <a-button @click="auth">授权</a-button>
</template>

<script lang="ts">
import {defineComponent, onMounted} from "vue"

export default defineComponent({
  name: "Login",
  setup() {
    // 申请授权
    const auth = () => {
      let clientId = 'f5763537e579c4f97c56c69b80489a17e250d8186e48efbe3e3fba4c4b6c9558'
      let redirectUri = 'http://127.0.0.1:8000/user/auth/gitee'
      let authUrl = 'https://gitee.com/oauth/authorize?client_id=' + clientId + '&redirect_uri=' + redirectUri + '&response_type=code'
      window.location.href = authUrl
    }
    return {
      auth
    }
  }
})
</script>

<style scoped>

</style>