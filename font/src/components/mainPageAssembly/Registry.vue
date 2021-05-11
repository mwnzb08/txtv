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
      maxlength="16"
      v-model:value="selectEntity.user_id"
    >
      <template v-slot:prefix>
        <user-outlined></user-outlined>
      </template>
    </a-input>
    <a-input
      class="list-padding-margin"
      placeholder="密码"
      maxlength="16"
      :type="password_type"
      v-model:value="selectEntity.pwd"
    >
      <template v-slot:prefix>
        <key-outlined></key-outlined>
      </template>
      <template v-slot:suffix>
        <eye-invisible-filled @click="password_type = 'text'" v-if="password_type === 'password'"></eye-invisible-filled>
        <eye-two-tone  @click="password_type = 'password'" v-if="password_type === 'text'"></eye-two-tone>
      </template>
    </a-input>
    <a-input
      class="list-padding-margin"
      placeholder="确认密码"
      maxlength="16"
      type="password"
      v-model:value="selectEntity.pwd2"
    >
      <template v-slot:prefix>
        <key-outlined></key-outlined>
      </template>
    </a-input>
    <a-button class="border-radius" type="primary">注册</a-button>
  </div>
</template>

<script>
import { UserOutlined, KeyOutlined, EyeInvisibleFilled, EyeTwoTone } from '@ant-design/icons-vue'
export default {
  name: 'Registry',
  components: {
    UserOutlined,
    KeyOutlined,
    EyeInvisibleFilled,
    EyeTwoTone
  },
  data () {
    return {
      selectEntity: { user_id: '', pwd: '' },
      password_type: 'password'
    }
  }
}
</script>
