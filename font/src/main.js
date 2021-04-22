import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import moment from 'moment'
import axios from 'axios'
createApp.prototype.$http = axios
createApp(App).use(store).use(router).use(moment)
  .mount('#app')
