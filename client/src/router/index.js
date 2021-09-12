import Vue from 'vue'
import VueRouter from 'vue-router'
import store from '../store'


Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Dashboard',
    redirect: '/dashboard',
    component: () => import("../components/nav/Navbar.vue"),
    meta: { requiresAuth: true },
    children: [
      {
        path: '/time/tracking',
        name: 'Tracking Time',
        component: () => import('../views/timesheet/Tracking.vue'),
        meta: { requiresAuth: true }
      },
      {
        path: '/about',
        name: 'About',
        component: () => import(/* webpackChunkName: "about" */ '../views/About.vue'),
        meta: { requiresAuth: true }
      },  
    ]
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/auth/Login.vue')
  },
  {
    path: '/sign_up',
    name: 'Sign Up',
    component: () => import('../views/auth/Sign_up.vue')
  },
  {path: "*", redirect: '/'}
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    let loggedIn = store.state.user.token;
    if(!loggedIn){
      store.dispatch("user/setUser");
    }
    if (!loggedIn) {
      store.dispatch('logOut')
      next({
        path: '/login',
        query: {redirect: to.fullPath}
      })
    }else[
      next()
    ]
  }else{
    next();
  }
})
export default router
