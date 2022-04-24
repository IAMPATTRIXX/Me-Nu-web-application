import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '@/views/Home.vue'
import NotFound from '@/components/404.vue'


Vue.use(VueRouter)

const routes = [
  {
    path: '/:tID/:id',
    name: 'Home',
    component: Home
  },
  {
    path: '/404',
    name: '404',
    component: NotFound
  },
  
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
