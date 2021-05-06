import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import moment from 'moment'
import axios from 'axios'
import './assets/css/global.styl'
const app = createApp(App)
app.config.globalProperties.$http = axios
app.use(store).use(router).use(moment)
  .mount('#app')
