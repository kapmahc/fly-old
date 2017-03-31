import Install from './Install'
import Dashboard from './Dashboard'
import Info from './admin/Info'
import Author from './admin/Author'
import Seo from './admin/Seo'
import Smtp from './admin/Smtp'
import Status from './admin/Status'
import Locales from './admin/Locales'
import AdminNotices from './admin/Notices'
import AdminLeaveWords from './admin/LeaveWords'
import NewLeaveWord from './leave-words/New'
import AdminUsers from './admin/Users'
import IndexNotices from './notices/Index'

export default {
  routes: [
    {
      path: '/notices',
      name: 'site.notices.index',
      component: IndexNotices
    },
    {
      path: '/admin/users',
      name: 'site.admin.users',
      component: AdminUsers
    },
    {
      path: '/leave-words/new',
      name: 'site.leave-words.new',
      component: NewLeaveWord
    },
    {
      path: '/admin/leave-words',
      name: 'site.admin.leave-words',
      component: AdminLeaveWords
    },
    {
      path: '/admin/notices',
      name: 'site.admin.notices',
      component: AdminNotices
    },
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
      redirect: {name: process.env.HOME},
      name: 'home'
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
      href: {name: 'site.notices.index'},
      label: 'site.notices.index.title'
    }
  ],
  dashboard (user) {
    if (user.uid && user.isAdmin) {
      return {
        label: 'site.dashboard.title',
        items: [
          {href: {name: 'site.admin.status'}, label: 'site.admin.status.title'},
          null,
          {href: {name: 'site.admin.info'}, label: 'site.admin.info.title'},
          {href: {name: 'site.admin.author'}, label: 'site.admin.author.title'},
          {href: {name: 'site.admin.seo'}, label: 'site.admin.seo.title'},
          {href: {name: 'site.admin.smtp'}, label: 'site.admin.smtp.title'},
          null,
          {href: {name: 'site.admin.locales'}, label: 'site.admin.locales.index.title'},
          {href: {name: 'site.admin.users'}, label: 'site.admin.users.index.title'},
          null,
          {href: {name: 'site.admin.notices'}, label: 'site.admin.notices.index.title'},
          {href: {name: 'site.admin.leave-words'}, label: 'site.admin.leave-words.index.title'}
        ]
      }
    }
    return null
  }
}
