import Vue from 'vue'
import VueRouter from 'vue-router'
import {routes} from './routes'
Vue.use(VueRouter);
const router = new VueRouter({
	mode:'hash',
	linkActiveClass:'Aactive',
	routes:[
		...routes
	],
	scrollBehavior (to, from, savedPosition) {
	  if (savedPosition) {
	    return savedPosition
	  } else {
	    return { x: 0, y: 0 }
	  }
	},
})

router.beforeEach((to, from, next) => {
    next()
})

export default router;
