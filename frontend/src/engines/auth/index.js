import SignIn from './SignIn'
import SignUp from './SignUp'

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
      component: SignIn
    },
    {
      path: '/users/confirm',
      name: 'auth.users.confirm',
      component: SignIn
    },
    {
      path: '/users/unlock',
      name: 'auth.users.unlock',
      component: SignIn
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
