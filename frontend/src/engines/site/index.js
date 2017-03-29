import Home from './Home'
import Install from './Install'

export default {
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/install',
      name: 'site.install',
      component: Install
    }
  ],
  links: [
    {
      href: 'home',
      label: 'site.notices.index.title'
    }
  ],
  dashboard: []
}
