import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify';
import router from "./router";
import store from "./store"
import axios from "axios"
import { getToken } from "./util";

Vue.config.productionTip = false;

axios.defaults.baseURL = process.env.VUE_APP_API_BASE_URL + '/api';
axios.interceptors.request.use(config => {
  config.headers['Authorization'] = `Bearer ${getToken()}`;
  return config;
});
axios.interceptors.response.use(
  response => response,
  error => error.response || error
);

const globalMixIn = {
  created() {
    let { title } = this.$options;
    if (title) {
      document.title = title;
    }
  }
};

Vue.mixin(globalMixIn);

const createApp = async () => {
  await store.dispatch('auth/currentUser');

  new Vue({
    vuetify,
    router,
    store,
    render: h => h(App)
  }).$mount('#app');
};

createApp();

