import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from './axios'
import moment from 'moment'
import { Button, Input, message } from 'ant-design-vue'
import './assets/css/global.styl'

message.config({ maxCount: 3 })
window.message = message
const app = createApp(App)
app.config.globalProperties.$http = axios
app.use(Input).use(Button)
app.use(store).use(router).use(moment)
  .mount('#app')
