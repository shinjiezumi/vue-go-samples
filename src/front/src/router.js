import Vue from 'vue'
import Router from "vue-router"
import Hello from "./views/Hello";
import NotFound from "./views/404"
import Login from "./views/auth/Login";
import Register from "./views/auth/Register";
import Top from "./views/Top";
import Todo from "./views/Todo"
import Searcher from "./views/Searcher.vue"
import store from "./store"

Vue.use(Router);

export default new Router({
  mode: "history",
  routes: [
    {
      path: '/',
      component: Top,
    },
    {
      path: '/register',
      component: Register,
      beforeEnter(to, from, next) {
        if (store.getters['auth/check']) {
          next('/todos')
        } else {
          next()
        }
      }
    },
    {
      path: '/login',
      component: Login,
      beforeEnter(to, from, next) {
        if (store.getters['auth/check']) {
          next('/todos')
        } else {
          next()
        }
      }
    },
    {
      path: '/todos',
      component: Todo,
      beforeEnter(to, from, next) {
        if (!store.getters['auth/check']) {
          next('/login')
        } else {
          next()
        }
      }
    },
    {
      path: '/searcher',
      component: Searcher,
    },
    {
      path: '/hello',
      component: Hello
    },
    {
      path: '*',
      component: NotFound
    }
  ]
})

