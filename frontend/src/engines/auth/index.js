import SignIn from './SignIn'
import SignUp from './SignUp'

export default [
  {
    path: '/users/sign-up',
    name: 'UsersSignUp',
    component: SignUp
  },
  {
    path: '/users/sign-in',
    name: 'UsersSignIn',
    component: SignIn
  }
]
