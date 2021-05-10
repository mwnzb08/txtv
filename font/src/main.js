import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from './axios'
import moment from 'moment'
import { Button, Input, Icon, message } from 'ant-design-vue'
import './assets/css/global.styl'
const app = createApp(App)
message.config({ maxCount: 3 })
window.message = message
app.config.globalProperties.$http = axios
app.use(Input).use(Icon).use(Button)
app.use(store).use(router).use(moment)
  .mount('#app')
