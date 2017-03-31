import SignIn from './users/SignIn'
import SignUp from './users/SignUp'
import Confirm from './users/Confirm'
import ForgotPassword from './users/ForgotPassword'
import ResetPassword from './users/ResetPassword'
import Unlock from './users/Unlock'
import Logs from './users/Logs'
import Info from './users/Info'
import ChangePassword from './users/ChangePassword'
import IndexUsers from './users/Index'

export default {
  routes: [
    {
      path: '/users',
      name: 'auth.users.index',
      component: IndexUsers
    },
    {
      path: '/users/change-password',
      name: 'auth.users.change-password',
      component: ChangePassword
    },
    {
      path: '/users/info',
      name: 'auth.users.info',
      component: Info
    },
    {
      path: '/users/logs',
      name: 'auth.users.logs',
      component: Logs
    },
    {
      path: '/users/sign-up',
      name: 'auth.users.sign-up',
      component: SignUp
    },
    {
      path: '/users/sign-in',
      name: 'auth.users.sign-in',
      component: SignIn
    },
    {
      path: '/users/forgot-password',
      name: 'auth.users.forgot-password',
      component: ForgotPassword
    },
    {
      path: '/users/reset-password/:token',
      name: 'auth.users.reset-password',
      component: ResetPassword
    },
    {
      path: '/users/confirm',
      name: 'auth.users.confirm',
      component: Confirm
    },
    {
      path: '/users/unlock',
      name: 'auth.users.unlock',
      component: Unlock
    }
  ],
  links: [
    {
      href: 'auth.users.index',
      label: 'auth.users.index.title'
    }
  ],
  dashboard (user) {
    if (user.uid) {
      return {
        label: 'auth.dashboard.title',
        items: [
          {href: 'auth.users.info', label: 'auth.users.info.title'},
          {href: 'auth.users.change-password', label: 'auth.users.change-password.title'},
          null,
          {href: 'auth.users.logs', label: 'auth.users.logs.title'}
        ]
      }
    }
    return null
  }
}
