import React from 'react';
import { Route, IndexRoute } from 'react-router'

import Home from './Home'
import {Index as IndexNotices} from './Notices'

export default {
  navLinks: [
    {href:'/notices', label:'site.notices.index.title'},
  ],
  routes: [
    (<IndexRoute key="site.root" component={Home} />),
    (<Route key="site.home" path="/home" component={Home} />),
    (<Route key="site.notices" path="/notices">
      <IndexRoute component={IndexNotices} />
    </Route>)
  ]
}
