import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import vuetify from './plugins/vuetify'
import Vuetify from 'vuetify/lib'
import axios from 'axios'
import ReadMore from 'vue-read-more';
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

const firebaseConfig = {
  apiKey: "AIzaSyB89PkzBUBxBX-pkoTA4osJ9_JpTywTmWE",
  authDomain: "me-nu-cus.firebaseapp.com",
  projectId: "me-nu-cus",
  storageBucket: "me-nu-cus.appspot.com",
  messagingSenderId: "857121317179",
  appId: "1:857121317179:web:be43d01a458db48c62171e",
  measurementId: "G-05VXGV9FBD"
};

Vue.config.productionTip = false

const app = initializeApp(firebaseConfig);
const analytics = getAnalytics(app);

// Vue.use(VueSocketio, io('http://localhost:8081/food/res/menu-food'));
Vue.use(Vuetify)
Vue.use(ReadMore)
new Vue({
  router,
  store,
  axios,
  vuetify,
  analytics,
  
  render: h => h(App)
}).$mount('#app')
