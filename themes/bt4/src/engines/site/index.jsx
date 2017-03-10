import React from 'react';
import { Route, IndexRoute } from 'react-router'

import Home from './Home'
import {Index as IndexNotices} from './Notices'
import Dashboard from './Dashboard'

export default {
  navLinks: [
    {href:'/notices', label:'site.notices.index.title'},
  ],
  dashboard: <Dashboard key='site.dashboard'/>,
  routes: [
    (<IndexRoute key="site.home" component={Home} />),    
    (<Route key="site.notices" path="/notices">
      <IndexRoute component={IndexNotices} />
    </Route>)
  ]
}
