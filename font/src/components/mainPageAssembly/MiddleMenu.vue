<style scoped lang="stylus">
.mian-head
  background-color rgba(244,233,244,.5)
.mian-title
  font-size 15px
  font-family "Times New Roman"
  font-weight bold
  min-width 70px
  cursor pointer
  text-align left
  margin 0 6px 0 15px
.mian-title:hover {
  color orange
}
.main-children
  font-size 15px
  text-align left
  min-width 70px
  cursor pointer
  color #666666
  font-family "Times New Roman"
.main-children:hover
  color indianred
.col-10-3
  display flex
  float left
  width 30%
  min-width 240px
  max-width 360px
  height 42px
  margin 0 10px 15px 10px
  border-radius 20px
  align-items center
.col-10-16
  display flex
  float left
  width 16.66%
  font-size 15px
  min-width 45px
  height 42px
  margin 0 0 15px 0
  border-radius 20px
  align-items center
.col-10-16-hover
  cursor pointer
.col-10-16-hover:hover
    color indianred
</style>

<template>
  <div style="display: flex">
    <div style="width: 45%; float: left; min-width: 800px; margin-left: 5%">
      <div class="col-10-3 mian-head" v-for="it in middleTittleColumn" v-bind:key="it.name">
        <span class="mian-title"><check-circle-filled style="width: 25px"></check-circle-filled>{{it.name}}</span>
        <span style="margin-right: 10px">|</span>
        <span class="main-children">{{it.liName1}}</span>
        <span style="width: 5px"></span>
        <span class="main-children">{{it.liName2}}</span>
      </div>
    </div>
    <div style="margin: 0 2% 0 1%; padding: 25px 0 25px 0">
      <div style="display: flex; height: 100%; background-color: #eee; width: 1px"></div>
    </div>
    <div style="float: left; width: 25%; min-width: 400px">
      <div class="col-10-16" v-for="it in middleRightColumn" v-bind:key="it.name"><span class="col-10-16-hover">{{it.name}}</span>
      </div>
    </div>
  </div>
</template>

<script>
import { CheckCircleFilled } from '@ant-design/icons-vue'
export default {
  name: 'MiddleMenu',
  components: {
    CheckCircleFilled
  },
  data () {
    return {
      middleTittleColumn: [
        { name: '电视剧', liName1: '海外剧', liName2: '网络剧' },
        { name: '综艺', liName1: '腾讯自制', liName2: '独播' },
        { name: '少儿', liName1: '益智', liName2: '儿歌' },
        { name: '电影', liName1: '影院大片', liName2: '自制电影' },
        { name: '动漫', liName1: '国漫', liName2: '日漫' },
        { name: '纪录片', liName1: '美食', liName2: '自然' }
      ],
      middleRightColumn: [
        { name: 'VIP会员' },
        { name: 'NBA' },
        { name: '娱乐' },
        { name: '新闻' },
        { name: '艺术' },
        { name: '汽车' },
        { name: '音乐' },
        { name: '体育' },
        { name: '生活' },
        { name: '游戏' },
        { name: '知识' },
        { name: '更多' }
      ]
    }
  }
}
</script>
