import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import vuetify from './plugins/vuetify'
import Vuetify from 'vuetify/lib'
import axios from 'axios'
import { initializeApp } from "firebase/app";
import { getAnalytics } from "firebase/analytics";
// import VueSocketio from 'vue-socket.io-extended';
// import io from 'socket.io-client';

window.axios = require('axios');

Vue.filter("readmore", val => {
  if (!val || typeof val != "string") {
    return "";
  }
  let result = val.slice(0, 10);
  result = result + "..";
  return result;
});

Vue.config.productionTip = false

const firebaseConfig = {
  apiKey: "AIzaSyACPuaOys-D0bYfaAReOMtvOmjas6E9Ohs",
  authDomain: "me-nu-res.firebaseapp.com",
  projectId: "me-nu-res",
  storageBucket: "me-nu-res.appspot.com",
  messagingSenderId: "887922284748",
  appId: "1:887922284748:web:6d2771f88d7d13c9454a69",
  measurementId: "G-28N52SXVJ9"
};

// Vue.use(VueSocketio, io('http://localhost:8081/food/res/menu-food'));
Vue.use(Vuetify)
const app = initializeApp(firebaseConfig);
const analytics = getAnalytics(app);
// Vue.use(ReadMore)
new Vue({
  router,
  store,
  axios,
  vuetify,
  analytics,
  
  render: h => h(App)
}).$mount('#app')
