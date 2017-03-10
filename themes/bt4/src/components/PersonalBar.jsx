import React,{PropTypes} from 'react';
import { connect } from 'react-redux'
import { NavDropdown, MenuItem } from 'react-bootstrap';
import {LinkContainer} from 'react-router-bootstrap'
import i18next from 'i18next';

const Widget = ({user}) => (
  user.uid ?
  <NavDropdown title={i18next.t("personal-bar.welcome", {name:user.name})} id="header-personal-bar">
    <MenuItem>Action</MenuItem>
    <MenuItem>Another action</MenuItem>
    <MenuItem>Something else here</MenuItem>
    <MenuItem divider />
    <MenuItem>Separated link</MenuItem>
  </NavDropdown> :
  <NavDropdown title={i18next.t("personal-bar.sign-in-or-up")} id="header-personal-bar">
    {['sign-in', 'sign-up', 'confirm', 'unlock', 'forgot-password'].map((k,i)=>(
      <LinkContainer key={i} to={`/users/${k}`}>
        <MenuItem>{i18next.t(`auth.users.${k}.title`)}</MenuItem>
      </LinkContainer>
    ))}
  </NavDropdown>
)

Widget.propTypes = {
  user: PropTypes.object.isRequired
}

export default connect(
  state => ({user: state.currentUser})
)(Widget)
