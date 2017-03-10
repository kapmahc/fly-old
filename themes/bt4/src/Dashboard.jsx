import React from 'react';
import engines from './engines'

const Widget = ({children}) => (
  <div>
    {engines.dashboard}
    <br/>
    {children}
  </div>
)

export default Widget
