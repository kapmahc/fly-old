import React from 'react';
import { NavDropdown, MenuItem } from 'react-bootstrap';

const Widget = () => (
  <NavDropdown eventKey={3} title="language" id="header-language-bar">
    <MenuItem eventKey={3.1}>Action</MenuItem>
    <MenuItem eventKey={3.2}>Another action</MenuItem>
    <MenuItem eventKey={3.3}>Something else here</MenuItem>
    <MenuItem divider />
    <MenuItem eventKey={3.3}>Separated link</MenuItem>
  </NavDropdown>
)

export default Widget
