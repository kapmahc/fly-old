import SignIn from './SignIn'
import SignUp from './SignUp'

export default [
  {
    path: '/users/sign-up',
    name: 'users.sign-up',
    component: SignUp
  },
  {
    path: '/users/sign-in',
    name: 'users.sign-in',
    component: SignIn
  }
]
