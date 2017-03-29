import Router from 'vue-router'

import auth from '@/engines/auth'
import site from '@/engines/site'

export default new Router({
  routes: [].concat(auth, site)
})
