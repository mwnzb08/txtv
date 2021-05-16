<style scoped lang="stylus">

</style>

<template>
<div>
  <main-select-menu style="margin-bottom: 30px; min-width: 1300px">
  </main-select-menu>
  <middle-menu style="height: 100px; margin-bottom: 40px">
  </middle-menu>
</div>
</template>

<script>
import MainSelectMenu from './mainPageAssembly/MainSelectMenu'
import MiddleMenu from './mainPageAssembly/MiddleMenu'
export default {
  name: 'MainPage',
  components: { MiddleMenu, MainSelectMenu },
  props: {},
  data () {
    return {}
  },
  mounted () {
  },
  methods: {
  },
  beforeUnmount () {
  }
}
</script>
