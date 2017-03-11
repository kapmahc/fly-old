import React from 'react';
import { Route, IndexRoute } from 'react-router'

import Home from './Home'
import {Index as IndexNotices} from './Notices'
import Dashboard from './Dashboard'
import Root from '../../Dashboard'
import {
  Info as SiteInfo,
  Author as SiteAuthor,
  Seo as SiteSeo
} from './admin'

export default {
  navLinks: [
    {href:'/notices', label:'site.notices.index.title'},
  ],
  dashboard: <Dashboard key='site.dashboard'/>,
  routes: [
    (<IndexRoute key="site.home" component={Home} />),
    (<Route key="site.notices" path="/notices">
      <IndexRoute component={IndexNotices} />
    </Route>),
    (<Route key="site.admin" path="/admin" component={Root}>
      <Route path="site/info" component={SiteInfo}/>
      <Route path="site/author" component={SiteAuthor}/>
      <Route path="site/seo" component={SiteSeo}/>
    </Route>)
  ]
}
