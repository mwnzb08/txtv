import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import moment from 'moment'
import axios from 'axios'
import { Button, Input, Icon } from 'ant-design-vue'
import './assets/css/global.styl'

const app = createApp(App)
app.use()
app.config.globalProperties.$http = axios
app.use(Input).use(Icon).use(Button)
app.use(store).use(router).use(moment)
  .mount('#app')
