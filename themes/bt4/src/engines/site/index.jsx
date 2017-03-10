import React from 'react';
import { Route, IndexRoute } from 'react-router'

import Home from './Home'
import {Index as IndexNotices} from './Notices'

export default {
  routes: [
    (<IndexRoute key="site.home" component={Home} />),
    (<Route key="site.notices" path="/notices">
      <IndexRoute component={IndexNotices} />
    </Route>)
  ]
}
