<style lang="stylus">
  #app
    font-family Avenir, Helvetica, Arial, sans-serif
    -webkit-font-smoothing antialiased
    -moz-osx-font-smoothing grayscale
    text-align center
    color #2c3e50
</style>

<template>
  <router-view />
</template>

<script>
export default {
  beforeMount () {
    if (window.sessionStorage.getItem('userSession')) {
      this.$store.state.userSession = JSON.parse(window.sessionStorage.getItem('userSession'))
      console.log(this.$store.state.userSession.isLogin)
    }
    this.$http.get('/msg').then((res) => {
      window.messageString = res.data
      // window.sessionStorage.setItem('messageString', JSON.stringify(res.data))
      //
    })
  },
  setup () {
  }
}
</script>
