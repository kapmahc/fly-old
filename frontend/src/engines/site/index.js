import Home from './Home'

export default {
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
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
