<style scoped lang="stylus">
  .mainSelectMenu
    height 500px
    width 100%
    overflow-y auto
    overflow-x hidden
    background-color rgba(69, 68, 73, 0.86)
    /*background-repeat no-repeat*/
    /*background-size 100% 100%*/
  .mainSelectMenu2
    position absolute
    right 0
    background-color rgba(15,15,30,.5)
    float right
    padding-left 40px
    cursor initial
    width 360px
    height 432px
  .mainSelectMenu3
    color white
    text-align left
    font-size 22px
    cursor pointer
    height 29.25px
    padding-bottom 10px
  .mainSelectMenu4
    margin-top 10px
    height 0.5px
    text-align left
    background-color white
  .mainSelectMenu5
    height 25.25px
    color white
    cursor default
    font-size 22px
    text-align left
    padding 5px 0 30px 0
  .mainSelectMenu6
    overflow-y auto
    height 100%
    max-height 332px
    scrollbar-width: none
  .mainSelectMenu6 table
    float left
    height 100%
    width 100%
  .mainSelectMenu6 table tbody
    display block
  .mainSelectMenu6 table tbody tr
    height 50px
    line-height 40px
    display block
  .mainSelectMenu6 table tbody tr td
    text-align left
    display block
    width 100%
    font-size 15px
    color white
    cursor pointer
    transition font-size 0.2s
  .mainSelectMenu6::-webkit-scrollbar {
    width 0
  }
  .publicSelectMenu
    float left
    display flex
    flex-direction row
    width 100%
    height 424px
    overflow-x hidden
    overflow-y hidden
    padding 30px 0 30px 0
    scrollbar-width: none
    scroll-behavior smooth

  .publicSelectMenu img
    float left
    border-radius 5px
    margin 0 5px 0 5px
    height 364px
    width 260px
    cursor pointer
    transition all 0.2s
  .publicSelectMenu img:hover
    transform scale(1.08)
  .moveButton
    position: absolute
    right: 360px
    height: 372px
    width: 50px
  .moveButton img
    height 50px
    width 25px
    opacity 0.6
    top 171px
    position relative
    cursor pointer
  .moveButton2
    position: absolute
    height: 372px
    width: 50px
  .moveButton2 img
    position relative
    height 50px
    width 25px
    opacity 0.6
    top 171px
    right 50px
    transition all 0s
    transform rotateY(180deg)
    cursor pointer
    &:hover
      transform rotateY(-180deg) scale(1.08)
  .mainSelectHead
    height 68px
    background-color rgba(255,255,255,.14)
    display flex
    flex-direction row
    flex-wrap nowrap
    justify-content space-between
    align-items center
    cursor default
  .headLogo
    position relative
    padding-left 90px
    display: flex
  .headLogo img
    height 40px
    width 200px
    cursor pointer
  .headSearch
    position relative
  .headSearchLabel
    display flex
    align-items center
    justify-content center
    cursor pointer
    width 50px
  .borderStyle
    display flex
    background-color rgba(255,255,255,.30)
    border-radius 20px
    padding 0 0 0 20px
  .borderStyle input
    padding 0 0 0 0
    border-width 0
    color white
    font-size large
    font-family "Times New Roman"
    height 40px
    background-color rgba(255,255,255, 0)
    outline none
    width  250px
  .borderStyle button
    border-width 0
    cursor: pointer
    background-color orangered
    color white
    border-radius 0 20px 20px 0
    outline none
    font-size 15px
    height 40px
    width 100px
    &:hover
      background-color #ff1827
  .headUser
    position relative
    align-items center
    margin-right 90px
    display: flex
  .headUser img
    cursor pointer
    height 40px
    width 40px
    border-radius 20px

</style>

<template>
  <div id= "mainSelectMenu" class="mainSelectMenu" :style="{'cursor': showPublicSelectMenu ? 'default' : 'pointer'}">
    <div class="mainSelectHead">
      <div class="headLogo">
        <img src="../../assets/logoMain.png" alt="loading">
      </div>
      <div class="headSearch">
        <div class="borderStyle">
          <input /><label class="headSearchLabel">res</label><button>全网搜</button>
        </div>
      </div>
      <div class="headUser">
        <img src="../../assets/user.png" alt="loading" @mouseenter="doUserLoginOrDetails(true)" @mouseleave="doUserLoginOrDetails(false)">
        <modal @mouseenter="doFocusInUserLogin(true)" @mouseleave="doFocusInUserLogin(false)"
          :show="showModal"
        >
          <template v-slot:head></template>
          <template v-slot:body>
            <login>
            </login>
          </template>
        </modal>
      </div>
    </div>
    <div v-if="showPublicSelectMenu" style="padding-left: 90px">
      <div id="publicSelectMenu" class="publicSelectMenu">
        <div v-for="(column, index) in MenuList" v-bind:key="index">
          <img :src="column.url" alt="loading ..." @click="$router.push({path: column.url})">
        </div>
      </div>
      <div class="moveButton" v-if="showRightMoveButton">
        <img src="../../assets/force.png" @click="doImageMove(300)" alt="loading">
      </div>
      <div class="moveButton2" v-if="showLeftMoveButton">
        <img src="../../assets/force.png" @click="doImageMove(-300)" alt="loading">
      </div>
    </div>
    <div v-else style="position: absolute; width: calc(100% - 340px); height: 430px" @click="doBackGroundClick"></div>
    <div class="mainSelectMenu2">
      <div id="mainSelectMenu3" class="mainSelectMenu3" @mouseenter="doMouseEnterList">大家在看</div>
      <div id="mainSelectMenu4" class="mainSelectMenu4"></div>
      <div id="mainSelectMenu5" class="mainSelectMenu5">重磅推荐</div>
      <div id="mainSelectMenu6" class="mainSelectMenu6">
        <table>
          <tbody>
          <tr v-for="(column, index) in MenuList" v-bind:key="index">
            <td :id="'indexOfMenu' + index" @mouseenter="doMouseEnter(index, column.url)" @mouseleave="doSetInterval(index)" @click="$router.push({path: column.url})">{{column.name}}</td>
          </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
import Modal from '../commonBase/Modal'
import Login from './Login'
export default {
  name: 'mainSelectMenu',
  components: { Login, Modal },
  data () {
    return {
      selectBox: 'public',
      MenuList: [
        { name: '周莲妃', url: 'https://ss1.bdstatic.com/70cFuXSh_Q1YnxGkpoWK1HF6hhy/it/u=4021452577,2359594280&fm=26&gp=0.jpg' },
        { name: '唐复琼', url: 'https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fimg1.gtimg.com%2F13%2F1391%2F139152%2F13915244_640x640_281.jpg&refer=http%3A%2F%2Fimg1.gtimg.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1620643463&t=072a5eccad1b1644d412ffc1dfa55747' },
        { name: '徐莉', url: 'https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fgame.people.com.cn%2FNMediaFile%2F2015%2F0107%2FMAIN201501070835000494764022963.jpg&refer=http%3A%2F%2Fgame.people.com.cn&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1620643463&t=d3ec21daf7976d9c52ab5401e71291fe' },
        { name: '莫林', url: 'https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fcomic.people.com.cn%2FNMediaFile%2F2014%2F1201%2FMAIN201412011401000273973235794.jpg&refer=http%3A%2F%2Fcomic.people.com.cn&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1620643463&t=0ff3f920d5f6fcf89d241ca62c86cd0b' },
        { name: '周莲妃', url: 'https://ss1.bdstatic.com/70cFuXSh_Q1YnxGkpoWK1HF6hhy/it/u=4021452577,2359594280&fm=26&gp=0.jpg' },
        { name: '唐复琼', url: 'https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fimg1.gtimg.com%2F13%2F1391%2F139152%2F13915244_640x640_281.jpg&refer=http%3A%2F%2Fimg1.gtimg.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1620643463&t=072a5eccad1b1644d412ffc1dfa55747' },
        { name: '徐莉', url: 'https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fgame.people.com.cn%2FNMediaFile%2F2015%2F0107%2FMAIN201501070835000494764022963.jpg&refer=http%3A%2F%2Fgame.people.com.cn&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1620643463&t=d3ec21daf7976d9c52ab5401e71291fe' },
        { name: '莫林', url: 'https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fcomic.people.com.cn%2FNMediaFile%2F2014%2F1201%2FMAIN201412011401000273973235794.jpg&refer=http%3A%2F%2Fcomic.people.com.cn&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1620643463&t=0ff3f920d5f6fcf89d241ca62c86cd0b' }
      ],
      MenuListIndex: 0,
      setInterval1: null,
      showPublicSelectMenu: true,
      showLeftMoveButton: false,
      showRightMoveButton: true,
      showModal: true,
      setTimeOutJob1: null
    }
  },
  mounted () {
    document.getElementById('mainSelectMenu3').style.color = 'indianred'
    // this.doSetInterval()
  },
  computed: {
  },
  methods: {
    doImageMove (way) {
      const doScroll = document.getElementById('publicSelectMenu')
      doScroll.scrollTo(doScroll.scrollLeft + way, 0)
      if (doScroll.scrollLeft + way > 0) this.showLeftMoveButton = true
      else this.showLeftMoveButton = false
      if (doScroll.scrollLeft + way < doScroll.scrollWidth - doScroll.clientWidth) this.showRightMoveButton = true
      else this.showRightMoveButton = false
    },
    doMouseEnterList () {
      this.selectBox = 'public'
      this.showPublicSelectMenu = true
      clearInterval(this.setInterval1)
      document.getElementById('mainSelectMenu3').style.color = 'indianred'
      document.getElementById('mainSelectMenu').style.backgroundImage = ''
      document.getElementById('mainSelectMenu').style.backgroundColor = 'rgba(69, 68, 73, 0.86)'
      document.getElementById('indexOfMenu' + this.MenuListIndex).style.color = 'white'
      document.getElementById('indexOfMenu' + this.MenuListIndex).style.fontSize = '15px'
    },
    doMouseEnter (index, url) {
      this.showPublicSelectMenu = false
      clearInterval(this.setInterval1)
      document.getElementById('mainSelectMenu3').style.color = 'white'
      let withdrawChange = document.getElementById('indexOfMenu' + index)
      if (this.selectBox === 'public') {
        this.selectBox = 'selectMenu'
        document.getElementById('indexOfMenu0').style.fontSize = '15px'
        document.getElementById('indexOfMenu0').style.color = 'white'
        withdrawChange.style.fontSize = '21px'
        withdrawChange.style.color = 'indianred'
        // eslint-disable-next-line no-useless-escape
        document.getElementById('mainSelectMenu').style.backgroundImage = 'url(\"' + url + '\")'
      } else {
        if (this.MenuListIndex !== index) {
          const becomeChange = document.getElementById('indexOfMenu' + this.MenuListIndex)
          becomeChange.style.fontSize = '15px'
          becomeChange.style.color = 'white'
          withdrawChange = document.getElementById('indexOfMenu' + index)
          withdrawChange.style.fontSize = '21px'
          withdrawChange.style.color = 'indianred'
          // eslint-disable-next-line no-useless-escape
          document.getElementById('mainSelectMenu').style.backgroundImage = 'url(\"' + url + '\")'
        }
      }
      this.MenuListIndex = index
    },
    doSetInterval () {
      this.setInterval1 = setInterval(() => {
        const oldIndex = this.MenuListIndex
        if (this.MenuListIndex === this.MenuList.length - 1) {
          this.MenuListIndex = -1
        }
        this.MenuListIndex++
        const becomeChange = document.getElementById('indexOfMenu' + this.MenuListIndex)
        becomeChange.style.fontSize = '21px'
        becomeChange.style.color = 'indianred'
        // eslint-disable-next-line no-useless-escape
        document.getElementById('mainSelectMenu').style.backgroundImage = 'url(\"' + this.MenuList[this.MenuListIndex].url + '\")'
        const withdrawChange = document.getElementById('indexOfMenu' + oldIndex)
        withdrawChange.style.fontSize = '15px'
        withdrawChange.style.color = 'white'
      }, 50000)
    },
    doBackGroundClick () {
      this.$router.push({ path: this.MenuList[this.MenuListIndex].url })
    },
    doSetTimeOut200ms (val) {
      this.setTimeOutJob1 = setTimeout(() => {
        this.showModal = val
      }, 200)
    },
    doUserLoginOrDetails (val) {
      clearTimeout(this.setTimeOutJob1)
      this.doSetTimeOut200ms(val)
    },
    doFocusInUserLogin (val) {
      if (val) clearTimeout(this.setTimeOutJob1)
      else this.doSetTimeOut200ms(false)
    }
  },
  beforeUnmount () {
    clearInterval(this.setInterval1)
  }
}
</script>
