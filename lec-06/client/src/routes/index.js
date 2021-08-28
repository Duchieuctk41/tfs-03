import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../pages/Home.vue'
import Auth from '../pages/Auth.vue'

Vue.use(VueRouter)

const router = new VueRouter({
  mode: 'history',
  routes: [
    { path: '', component: Home },
    { path: '/auth', component: Auth },
  ]
})

export default router
