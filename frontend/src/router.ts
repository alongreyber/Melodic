import Vue from 'vue';
import Router from 'vue-router';
import Home from './views/Home.vue';
import LoggedIn from './views/LoggedIn.vue';
import SpotifyCallback from './views/SpotifyCallback.vue';

Vue.use(Router);

export default new Router({
    mode: 'history',
    base: process.env.BASE_URL,
    routes: [
	{
	    path: '/',
	    name: 'home',
	    component: Home,
	},
	{
	    path: '/about',
	    name: 'about',
	    // route level code-splitting
	    // this generates a separate chunk (about.[hash].js) for this route
	    // which is lazy-loaded when the route is visited.
	    component: () => import(/* webpackChunkName: "about" */ './views/About.vue'),
	},
	{
	    path: "/spotify_callback",
	    name: "spotify_callback",
	    component: SpotifyCallback,
	},
	{
	    path: "/logged_in",
	    name: "logged_in",
	    component: LoggedIn,
	},
    ],
});
