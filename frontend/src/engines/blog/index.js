import Index from './Index'
import Show from './Show'

export default {
  routes: [
    {
      path: '/blog',
      name: 'blog.index',
      component: Index
    },
    {
      path: '/blog/*',
      name: 'blog.show',
      component: Show
    }
  ],
  links: [
    {
      href: {name: 'blog.index'},
      label: 'blog.index.title'
    }
  ],
  dashboard (user) {
    return null
  }
}
