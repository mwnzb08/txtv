<style scoped lang="stylus">
  .login
    height 100%
    padding 0 20% 0 20%
  input
    float right
    border-radius 7px
  .login-button
    border-radius 7px
</style>

<template>
  <div class="login">
    <div class="list-padding">
      <label>用户名
        <input v-model="selectEntity.user_id" maxlength="32">
      </label>
    </div>
    <div class="list-padding">
      <label>密码
        <input v-model="selectEntity.pwd" maxlength="32">
      </label>
    </div>
    <button class="login-button" @click="login">登录</button>
  </div>
</template>

<script>
export default {
  name: 'login',
  data () {
    return {
      selectEntity: {}
    }
  },
  methods: {
    login () {
      this.$http.post('/login', this.selectEntity).then((res) => {
        this.$store.state.userSession = res.data.userSession
        window.sessionStorage.setItem('userSession', JSON.stringify(res.data.userSession))
      })
    }
  }
}
</script>
