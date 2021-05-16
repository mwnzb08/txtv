<style scoped lang="stylus">
  .login
    height 100%
    padding 0 20% 0 20%
  input
    float right
    border-radius 7px
</style>

<template>
  <div class="login">
    <div v-if="!userSession.isLogin">
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
      <div class="list-padding">
        <a-button class="border-radius" @click="login">登录</a-button>
      </div>
      <div>
        <label @click="$router.push({path: 'Registry'})">注册</label>
      </div>
    </div>
    <div v-if="userSession.isLogin" style="height: 100%; padding-top: 40%">
      <div>
        <a-button type="danger" @click="doLoginOut">登出</a-button>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'
export default {
  name: 'login',
  data () {
    return {
      selectEntity: { user_id: '', pwd: '', limit: 1 }
    }
  },
  computed: {
    ...mapState({
      userSession: state => state.userSession
    })
  },
  methods: {
    login () {
      this.$http.post('/login', this.selectEntity).then((res) => {
        if (res.data.userSession.isLogin) {
          this.$store.state.userSession = res.data.userSession
          window.sessionStorage.setItem('userSession', JSON.stringify(res.data.userSession))
          this.selectEntity.pwd = ''
        } else {
          window.notification.error({ message: 'Error', description: '用户名或者密码错误' })
        }
      })
    },
    doLoginOut () {
      this.userSession.isLogin = false
      window.sessionStorage.removeItem('userSession')
      this.$http.get('/loginOut', []).then((res) => {
        if (res.data.gridData) {
          window.message.success('login out success!')
        }
      })
    }
  }
}
</script>
