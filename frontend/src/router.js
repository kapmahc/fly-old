import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

import auth from '@/engines/auth'
import site from '@/engines/site'

const router = new Router({
  routes: [].concat(auth, site)
})

export default router
