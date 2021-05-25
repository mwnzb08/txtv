<style scoped lang="stylus">
  .registryLayout
    padding 20vh 40% 20vh 40%

</style>

<template>
  <div class="registryLayout">
    <h2>欢迎注册</h2>
    <a-input
      class="list-padding-margin border-radius"
      placeholder="用户名"
      :maxlength="16"
      v-model:value="selectEntity.user_id"
      @blur="doCheckRegistryUserId"
      @focusin="() => {checkParams.user_id = false; checkParams.success = false}"
      ref="userId"
    >
      <template v-slot:prefix>
        <user-outlined></user-outlined>
      </template>
      <template v-slot:suffix>
        <loading-outlined v-show="checkParams.loading"></loading-outlined>
        <close-circle-filled v-show="checkParams.user_id" style="color: red"></close-circle-filled>
        <check-circle-filled v-show="checkParams.success" style="color: green"></check-circle-filled>
        <exclamation-circle-filled v-show="checkParams.warning" style="color: orange"></exclamation-circle-filled>
      </template>
    </a-input>
    <a-input
      class="list-padding-margin"
      placeholder="密码"
      :maxlength="16"
      :type="passwordType"
      v-model:value="selectEntity.pwd"
    >
      <template v-slot:prefix>
        <key-outlined></key-outlined>
      </template>
      <template v-slot:suffix>
        <eye-invisible-filled @click="passwordType = 'text'" v-if="passwordType === 'password'"></eye-invisible-filled>
        <eye-two-tone  @click="passwordType = 'password'" v-if="passwordType === 'text'"></eye-two-tone>
      </template>
    </a-input>
    <a-input
      class="list-padding-margin"
      placeholder="确认密码"
      :maxlength="16"
      type="password"
      v-model:value="selectEntityIgnore.pwd2"
    >
      <template v-slot:prefix>
        <key-outlined></key-outlined>
      </template>
    </a-input>
    <a-input
      class="list-padding-margin"
      placeholder="邮箱"
      :maxlength="32"
      type="text"
      v-model:value="selectEntity.email"
    >
      <template v-slot:prefix>
        <mail-outlined></mail-outlined>
      </template>
    </a-input>
    <a-button
      class="border-radius"
      type="primary"
      :disabled="selectEntity.user_id === '' || selectEntity.pwd ==='' || selectEntityIgnore.pwd2 === '' || selectEntity.email === ''"
      @click="doRegistry"
    >
      注册
    </a-button>
  </div>
</template>

<script>
import { UserOutlined, KeyOutlined, EyeInvisibleFilled, EyeTwoTone, CloseCircleFilled, LoadingOutlined, CheckCircleFilled, ExclamationCircleFilled, MailOutlined } from '@ant-design/icons-vue'
export default {
  name: 'Registry',
  components: {
    UserOutlined,
    KeyOutlined,
    EyeInvisibleFilled,
    EyeTwoTone,
    CloseCircleFilled,
    LoadingOutlined,
    CheckCircleFilled,
    ExclamationCircleFilled,
    MailOutlined
  },
  data () {
    return {
      selectEntity: { user_id: '', pwd: '', email: '' },
      selectEntityIgnore: { pwd2: '' },
      passwordType: 'password',
      checkParams: { user_id: false, loading: false, success: false, warning: false },
      timer1: null
    }
  },
  methods: {
    doCheckRegistryUserId () {
      if (this.selectEntity.user_id.trim().length >= 6 && this.checkParams.warning === false) {
        this.checkParams.user_id = false
        this.checkParams.success = false
        this.checkParams.warning = false// 暂时不想放到vuex，或许
        this.checkParams.loading = true
        this.$http.post('/checkRegistryUserId', this.selectEntity).then((res) => {
          this.checkParams.loading = false
          if (res.data.gridData === 'try again later') {
            this.checkParams.success = false
            this.checkParams.user_id = false
            this.checkParams.warning = true
          } else {
            this.checkParams.success = !res.data.gridData
            this.checkParams.user_id = res.data.gridData
          }
        })
      } else if (this.selectEntity.user_id.trim().length < 6) {
        this.checkParams.user_id = true
      }
    },
    checkBeforeSubmit () {
      if (this.selectEntity.user_id.length < 6 || this.selectEntity.pwd.length < 6) {
        window.notification.error({ message: 'Error', description: window.messageString['appMsg.E0003'] })
        return false
      }
      var regEmail = /^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/
      if (this.selectEntity.email !== '' && !regEmail.test(this.selectEntity.email)) {
        window.notification.error({ message: 'Error', description: window.messageString['appMsg.E0005'] })
        return false
      }
      if (this.checkParams.warning) {
        window.notification.error({ message: 'Error', description: window.messageString['appMsg.E0002'] })
        return false
      }
      if (this.checkParams.user_id) {
        window.notification.error({ message: 'Error', description: window.messageString['appMsg.E0001'] })
        this.$refs.userId.focus()
        return false
      }
      if (this.selectEntity.pwd.trim() !== this.selectEntityIgnore.pwd2.trim()) {
        this.selectEntityIgnore.pwd2 = ''
        window.notification.error({ message: 'Error', description: window.messageString['appMsg.E0004'] })
        return false
      }
      return true
    },
    doRegistry () {
      if (!this.checkBeforeSubmit()) return
      this.$http.post('/registry', this.selectEntity).then((res) => {
        // this.checkParams.loading = false
        // this.checkParams.user_id = res.data.gridData
        const data = res.data
        if (!data.isRegistry) {
          window.notification.error({ message: 'Error', description: window.messageString[data.result] })
        } else {
          window.notification.success({ message: 'Successful', description: window.messageString[data.result] })
          this.timer1 = setTimeout(() => {
            this.$router.push({ path: '/' })
          }, 2100)
        }
      })
    }
  },
  beforeUnmount () {
    clearTimeout(this.timer1)
  }
}
</script>
