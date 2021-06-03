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
    const userSession = window.localStorage.getItem('userSession')
    if (userSession) {
      this.$http.post('/checkLoginStatus', JSON.parse(userSession)).then((res) => {
        if (res.data) {
          this.$store.state.userSession = JSON.parse(userSession)
        } else {
          window.localStorage.removeItem('userSession')
        }
      })
    }
    this.$http.get('/msg').then((res) => {
      window.messageString = res.data
    })
  },
  setup () {
  }
}
</script>
