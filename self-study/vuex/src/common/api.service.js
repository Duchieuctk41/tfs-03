import axios from 'axios'
import Vue from 'vue'
import VueAxios from 'vue-axios'
import API_URL from './config'

const ApiService = {
  init() {
    Vue.use(VueAxios, axios)
    Vue.axios.defaults.baseURL = API_URL
  },
  get(resource) {
    return Vue.axios.get(`${resource}`)
  },
}

export default ApiService
