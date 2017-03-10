import React from 'react';
import { Route } from 'react-router'

import {Index, SignIn,SignUp,Confirm,Unlock,ForgotPassword,ResetPassword, Layout as NonSignInL} from './non-sign-in'
import {Info, Logs, ChangePassword, Layout as MustSignInL} from './must-sign-in'

export default {
  navLinks: [
    {href:'/users', label:'auth.users.index.title'},
  ],
  routes: [
    (<Route key="auth.users.index" path="/users" component={Index} />),
    (<Route key="auth.non-sign-in" path="/users" component={NonSignInL}>
      <Route path="sign-in" component={SignIn}/>
      <Route path="sign-up" component={SignUp}/>
      <Route path="confirm" component={Confirm}/>
      <Route path="unlock" component={Unlock}/>
      <Route path="forgot-password" component={ForgotPassword}/>
      <Route path="reset-password/:token" component={ResetPassword}/>
    </Route>),
    (<Route key="auth-must-sign-in" path="/users" component={MustSignInL}>
      <Route path="info" component={Info}/>
      <Route path="logs" component={Logs}/>
      <Route path="change-password" component={ChangePassword}/>
    </Route>),
  ],
}
