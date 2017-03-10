import React from 'react';

export const SignIn = () => (
  <div>sign in</div>
)

export const SignUp = () => (
  <div>sign up</div>
)

export const Confirm = () => (
  <div>confirm</div>
)

export const Unlock = () => (
  <div>unlock</div>
)

export const ForgotPassword = () => (
  <div>forgot password</div>
)

export const ResetPassword = () => (
  <div>reset password</div>
)


export const Layout = ({children}) => (
  <div>
    non-sign in
    <br/>
    {children}
  </div>
)
