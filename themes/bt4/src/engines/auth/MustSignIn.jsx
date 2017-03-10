import React from 'react';

export const Info = () => (
  <div>user info</div>
)

export const Logs = () => (
  <div>logs</div>
)

export const ChangePassword = () => (
  <div>change password</div>
)

export const Layout = ({children}) => (
  <div>
    must sign in
    <br/>
    {children}
  </div>
)
