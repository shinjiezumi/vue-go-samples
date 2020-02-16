import Vue from 'vue'
import Router from "vue-router"
import Hello from "./views/Hello";
import NotFound from "./views/404"
import Login from "./views/auth/Login";
import Register from "./views/auth/Register";
import Top from "./views/Top";
import Todo from "./views/Todo"
import store from "./store"
import { APP_NAME } from "./constants";

Vue.use(Router);

export default new Router({
  mode: "history",
  routes: [
    {
      path: '/',
      component: Top,
      meta: {
        title: APP_NAME,
        description: 'Vue/Vuex + Go/Gin + MySQLを使ったサンプルアプリです',
      }
    },
    {
      path: '/Register',
      component: Register,
      meta: {
        title: `会員登録｜${APP_NAME}`,
        description: '会員登録ページです',
      },
      beforeEnter(to, from, next) {
        if (store.getters['auth/check']) {
          next('/todo')
        } else {
          next()
        }
      }
    },
    {
      path: '/Login',
      component: Login,
      meta: {
        title: `ログイン｜${APP_NAME}`,
        description: 'ログインページです',
      },
      beforeEnter(to, from, next) {
        if (store.getters['auth/check']) {
          next('/todo')
        } else {
          next()
        }
      }
    },
    {
      path: '/todo',
      component: Todo,
      meta: {
        title: `TodoList｜${APP_NAME}`,
        description: 'Todoリストアプリです',
      },
      beforeEnter(to, from, next) {
        if (!store.getters['auth/check']) {
          next('/login')
        } else {
          next()
        }
      }
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

