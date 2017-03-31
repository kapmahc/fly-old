import Home from './Home'
import Install from './Install'
import Dashboard from './Dashboard'
import Info from './admin/Info'
import Author from './admin/Author'
import Seo from './admin/Seo'
import Smtp from './admin/Smtp'
import Status from './admin/Status'
import Locales from './admin/Locales'

export default {
  routes: [
    {
      path: '/admin/site/locales',
      name: 'site.admin.locales',
      component: Locales
    },
    {
      path: '/admin/site/status',
      name: 'site.admin.status',
      component: Status
    },
    {
      path: '/admin/site/smtp',
      name: 'site.admin.smtp',
      component: Smtp
    },
    {
      path: '/admin/site/seo',
      name: 'site.admin.seo',
      component: Seo
    },
    {
      path: '/admin/site/author',
      name: 'site.admin.author',
      component: Author
    },
    {
      path: '/admin/site/info',
      name: 'site.admin.info',
      component: Info
    },
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: Dashboard
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
  dashboard (user) {
    if (user.uid && user.isAdmin) {
      return {
        label: 'site.dashboard.title',
        items: [
          {href: 'site.admin.status', label: 'site.admin.status.title'},
          null,
          {href: 'site.admin.info', label: 'site.admin.info.title'},
          {href: 'site.admin.author', label: 'site.admin.author.title'},
          {href: 'site.admin.seo', label: 'site.admin.seo.title'},
          {href: 'site.admin.smtp', label: 'site.admin.smtp.title'},
          null,
          {href: 'site.admin.locales', label: 'site.admin.locales.index.title'},
          {href: 'home', label: 'site.admin.users.index.title'},
          null,
          {href: 'home', label: 'site.admin.notices.index.title'},
          {href: 'home', label: 'site.admin.leave-words.index.title'}
        ]
      }
    }
    return null
  }
}
