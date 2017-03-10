import React from 'react';
import {NavDropdown, MenuItem} from 'react-bootstrap';
import i18next from 'i18next';
import {LinkContainer} from 'react-router-bootstrap'

const Widget = () => (<NavDropdown title={i18next.t('site.dashboard.title')} id="site-dashboard-dropdown">
  {['info', 'author', 'seo', 'smtp', 'status'].map((k, i)=>(<LinkContainer key={i}to={`/admin/site/${k}`}>
    <MenuItem>{i18next.t(`site.admin.${k}.title`)}</MenuItem>
  </LinkContainer>))}
  <MenuItem divider />
  <LinkContainer to="site/admin/locales">
    <MenuItem>{i18next.t('site.admin.locales.index.title')}</MenuItem>
  </LinkContainer>
  <LinkContainer to="site/admin/users">
    <MenuItem>{i18next.t('site.admin.users.index.title')}</MenuItem>
  </LinkContainer>
  <LinkContainer to="site/admin/notices">
    <MenuItem>{i18next.t('site.notices.index.title')}</MenuItem>
  </LinkContainer>
  <LinkContainer to="site/admin/leave-words">
    <MenuItem>{i18next.t('site.leave-words.index.title')}</MenuItem>
  </LinkContainer>
</NavDropdown>)

export default Widget
