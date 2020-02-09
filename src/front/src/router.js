import Vue from 'vue'
import Router from "vue-router"
import Hello from "./views/Hello";
import NotFound from "./views/404"
import Login from "./views/auth/Login";
import Register from "./views/auth/Register";
import Top from "./views/Top";

Vue.use(Router);

export default new Router({
  mode: "history",
  routes: [
    {
      path: '/', component: Top
    },
    {
      path: '/Register', component: Register
    },
    {
      path: '/Login', component: Login
    },
    {
      path: '/hello', component: Hello
    },
    {
      path: '*', component: NotFound
    }
  ]
})

