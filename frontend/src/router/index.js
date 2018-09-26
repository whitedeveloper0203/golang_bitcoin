import Vue from 'vue'
import Router from 'vue-router'
import Callback from '@/components/Callback'
import Home from '@/components/Home'
import Trade from '@/components/Trade'
import Profile from '@/components/Profile'

Vue.use(Router)

const router = new Router({
  mode: 'history',
  routes: [
    {
      path: '/callback',
      name: 'Callback',
      component: Callback
    },
    {
      path: '/home',
      name: 'Home',
      component: Home
    },
    {
      path: '/trade',
      name: 'Trade',
      component: Trade
    },
    {
      path: '/profile',
      name: 'Profile',
      component: Profile
    },
    {
      path: '*',
      redirect: '/home'
    }
  ]
})

export default router
