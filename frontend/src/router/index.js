import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Dashboard from '../views/Dashboard.vue'
import store from '../store';

Vue.use(VueRouter)

function loginGuard(to, from, next) {
    if(store.state.loggedIn == false) {
	next('/');
    } else {
	next();
    }
}

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: Dashboard,
    beforeEnter: loginGuard
  }
]

const router = new VueRouter({
  routes
})

router.afterEach((to, from) => {
    store.commit('clearAllMessages');
})

export default router
