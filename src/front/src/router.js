import Vue from 'vue';
import Router from 'vue-router';
// eslint-disable-next-line import/no-unresolved,import/extensions
import Hello from './views/Hello';
// eslint-disable-next-line import/no-unresolved,import/extensions
import NotFound from './views/404';
// eslint-disable-next-line import/no-unresolved,import/extensions
import Login from './views/auth/Login';
// eslint-disable-next-line import/no-unresolved,import/extensions
import Register from './views/auth/Register';
// eslint-disable-next-line import/no-unresolved,import/extensions
import Top from './views/Top';
// eslint-disable-next-line import/no-unresolved,import/extensions
import Todo from './views/Todo';
import Searcher from './views/Searcher.vue';
// eslint-disable-next-line import/no-cycle
import store from './store';

Vue.use(Router);

export default new Router({
  mode: 'history',
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
          next('/todos');
        } else {
          next();
        }
      },
    },
    {
      path: '/login',
      component: Login,
      beforeEnter(to, from, next) {
        if (store.getters['auth/check']) {
          next('/todos');
        } else {
          next();
        }
      },
    },
    {
      path: '/todos',
      component: Todo,
      beforeEnter(to, from, next) {
        if (!store.getters['auth/check']) {
          next('/login');
        } else {
          next();
        }
      },
    },
    {
      path: '/searcher',
      component: Searcher,
    },
    {
      path: '/hello',
      component: Hello,
    },
    {
      path: '*',
      component: NotFound,
    },
  ],
});
