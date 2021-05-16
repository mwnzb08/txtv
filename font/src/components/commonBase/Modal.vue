<style scoped>

</style>

<template>
<div ref="el" style="position: absolute; z-index: 50000; background-color: white; height: 300px; width: 400px
; border-radius: 5px; left: -300px; top:54px" v-show="show">
  <div><slot name="head"></slot></div>
  <div><slot name="body"></slot></div>
    <div><slot name="foot"></slot></div>
</div>
</template>

<script>
export default {
  name: 'Modal',
  props: {
    show: {
      type: Boolean,
      default: false
    }
  },
  watch: {
    show (newValue, oldValue) {
      const el = this.$refs.el
      if (newValue === true) {
        el.onclick = function () { el.style.display = 'block' }
      }
    }
  }
}
</script>
