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
      <a-input class="list-margin" v-model:value="selectEntity.user_id" maxlength="32" placeholder="用户名"></a-input>
      <a-input class="list-margin" v-model:value="selectEntity.pwd" maxlength="32" placeholder="密码"></a-input>
      <template v-if="userSession.countLogin && userSession.countLogin >= 5 && userSession.countLogin <= 10">
        <a-input class="list-margin" v-model:value="selectEntity.valid_code_ignore" maxlength="4" placeholder="验证码"></a-input>
        <span><a-button @click="doGetValidCode">获取验证码</a-button></span>
      </template>
      <a-button class="list-margin border-radius" type="primary"  @click="login" :disabled="loginDisable">登录</a-button>
      <div>
        <a @click="$router.push({path: 'Registry'})">注册</a>
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
      selectEntity: { user_id: '', pwd: '', valid_code_ignore: '', limit: 1 }
    }
  },
  computed: {
    ...mapState({
      userSession: state => state.userSession
    }),
    loginDisable () {
      return this.selectEntity.user_id === '' || this.selectEntity.pwd === '' ||
        (this.userSession.countLogin && this.userSession.countLogin >= 5 && this.userSession.countLogin <= 10 && this.selectEntity.valid_code_ignore === '')
    }
  },
  methods: {
    checkBeforeSubmit () {
      if (this.selectEntity.user_id.length < 6 || this.selectEntity.pwd < 6) {
        window.message.error('用户名或者密码过短, 必须 >=6 ')
        return false
      }
      if (this.userSession.countLogin > 10) {
        window.message.error('请求过于频繁, 请稍后刷新重试')
        return false
      }
      return true
    },
    login () {
      if (!this.checkBeforeSubmit()) return
      this.$http.post('/login', this.selectEntity).then((res) => {
        const data = res.data
        if (data.userSession && data.userSession.isLogin) {
          this.$store.state.userSession = data.userSession
          data.userSession.currentData = new Date()
          window.localStorage.setItem('userSession', JSON.stringify(data.userSession))
          this.selectEntity.pwd = ''
        } else { // return back service request
          if (data.userSession && data.userSession.countLogin) {
            this.$store.state.userSession = data.userSession
            window.localStorage.setItem('userSession', JSON.stringify(data.userSession))
            window.message.error(data.userSession.gridData)
          } else { // else
            window.message.error('用户名或者密码错误')
          }
        }
      })
    },
    doGetValidCode () {
      if (!this.selectEntity.user_id) {
        window.message.error('请输入用户名')
        return
      }
      this.$http.post('/validCodeToEmail', this.selectEntity).then((res) => {
        window.message.error('已发送')
      })
    },
    doLoginOut () {
      this.userSession.isLogin = false
      window.localStorage.removeItem('userSession')
      this.$http.get('/loginOut', []).then((res) => {
        if (res.data.gridData) {
          window.message.success('login out success!')
        }
      })
    }
  }
}
</script>
