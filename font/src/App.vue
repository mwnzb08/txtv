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
  mounted () {
  },
  beforeMount () {
    if (window.sessionStorage.getItem('userSession')) {
      console.log('sssssssssssssss')
      this.$store.state.userSession = JSON.stringify(JSON.parse(window.sessionStorage.getItem('userSession')))
    }
  }
}
</script>
