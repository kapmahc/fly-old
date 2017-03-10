import React from 'react';
import engines from './engines'
import {Nav} from 'react-bootstrap';

const Widget = ({children}) => (
  <div>
    <Nav bsStyle="tabs">
      {engines.dashboard}
    </Nav>
    <br/>
    {children}
  </div>
)

export default Widget
