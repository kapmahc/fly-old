import SignIn from './SignIn'
import SignUp from './SignUp'
import Confirm from './Confirm'
import ForgotPassword from './ForgotPassword'
import ResetPassword from './ResetPassword'
import Unlock from './Unlock'

export default {
  routes: [
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
      href: 'home',
      label: 'auth.users.index.title'
    }
  ],
  dashboard: []
}
