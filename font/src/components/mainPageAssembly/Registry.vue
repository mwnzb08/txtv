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
      v-model:value="selectEntity.pwd2"
    >
      <template v-slot:prefix>
        <key-outlined></key-outlined>
      </template>
    </a-input>
    <a-button
      class="border-radius"
      type="primary"
      :disabled="selectEntity.user_id === '' || selectEntity.pwd ==='' || selectEntity.pwd2 === ''"
      @click="doRegistry"
    >
      注册
    </a-button>
  </div>
</template>

<script>
import { UserOutlined, KeyOutlined, EyeInvisibleFilled, EyeTwoTone, CloseCircleFilled, LoadingOutlined, CheckCircleFilled } from '@ant-design/icons-vue'
export default {
  name: 'Registry',
  components: {
    UserOutlined,
    KeyOutlined,
    EyeInvisibleFilled,
    EyeTwoTone,
    CloseCircleFilled,
    LoadingOutlined,
    CheckCircleFilled
  },
  data () {
    return {
      selectEntity: { user_id: '', pwd: '', pwd2: '' },
      passwordType: 'password',
      checkParams: { user_id: false, loading: false, success: false }
    }
  },
  methods: {
    doCheckRegistryUserId () {
      if (this.selectEntity.user_id.trim() !== '') {
        this.checkParams.user_id = false
        this.checkParams.loading = true
        this.$http.post('/checkRegistryUserId', this.selectEntity).then((res) => {
          this.checkParams.loading = false
          this.checkParams.success = !res.data.gridData
          this.checkParams.user_id = res.data.gridData
        })
      }
    },
    checkBeforeSubmit () {
      if (this.checkParams.user_id) {
        window.notification.error({ message: 'Error', description: '无效用户名' })
        this.$refs.userId.focus()
        return false
      }
      if (this.selectEntity.pwd.trim() !== this.selectEntity.pwd2.trim()) {
        this.selectEntity.pwd2 = ''
        window.notification.error({ message: 'Error', description: '两次输入密码不一致' })
        return false
      }
      return true
    },
    doRegistry () {
      if (!this.checkBeforeSubmit()) return
      this.$http.post('/registry', this.selectEntity).then((res) => {
        this.checkParams.loading = false
        this.checkParams.user_id = res.data.gridData
      })
    }
  }
}
</script>
