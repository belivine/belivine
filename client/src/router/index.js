import Vue from 'vue'
import VueRouter from 'vue-router'
import Navbar from "../components/nav/Navbar.vue"

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Dashboard',
    redirect: '/dashboard',
    component: Navbar,
    children: [
      {
        path: '/time/tracking',
        name: 'Tracking Time',
        component: () => import('../views/timesheet/Tracking.vue')
      },
      {
        path: '/about',
        name: 'About',
        component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
      },  
    ]
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/auth/Login.vue')
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
