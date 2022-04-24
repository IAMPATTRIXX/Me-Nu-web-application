import Vue from 'vue'
import VueRouter from 'vue-router'
// import Home from '@/views/Home.vue'
import Staff from '@/views/Staff.vue'
import formAddmenu from '@/components/Staff_Add.vue'
import formEditmenu from '@/components/Staff_Edit.vue'
import formRankmenu from '@/components/Staff_Ranking.vue'
import formQr from '@/components/Staff_Qr.vue'
import formBill from '@/components/Staff_Bill.vue'
import formShowqr from '@/components/Staff_ShowQR.vue'
import formAds from '@/components/Staff_Ads.vue'

Vue.use(VueRouter)

const routes = [
  // {
  //   path: '/',
  //   name: 'Home',
  //   component: Home
  // },
  {
    path: '/',
    name: 'Staff',
    component: Staff
  },
  {
    path: '/staff/add',
    name: 'add',
    component: formAddmenu

  },
  {
    path: '/staff/edit',
    name: 'edit',
    component: formEditmenu

  },
  {
    path: '/staff/rank',
    name: 'rank',
    component: formRankmenu

  },
  {
    path: '/staff/qr',
    name: 'qr',
    component: formQr

  },
  {
    path: '/staff/bill',
    name: 'bill',
    component: formBill
  },
  {
    path: '/staff/showqr',
    name: 'showqr',
    component: formShowqr
  },
  {
    path: '/staff/ads',
    name: 'ads',
    component: formAds
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
