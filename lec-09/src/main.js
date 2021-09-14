import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from 'axios'

Vue.config.productionTip = false

Vue.axios = axios
Vue.axios.defaults.baseURL = 'https://jsonplaceholder.typicode.com/'

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')
