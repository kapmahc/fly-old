import React from 'react';
import ReactDOM from 'react-dom';
import { createStore, combineReducers } from 'redux'
import { Provider } from 'react-redux'
import { Router, Route, browserHistory } from 'react-router'
import { syncHistoryWithStore, routerReducer } from 'react-router-redux'

import Layout from './Layout';
import reducers from './reducers'
import engines from './engines'

// Add the reducer to your store on the `routing` key
const store = createStore(
  combineReducers({
    ...reducers,
    routing: routerReducer
  })
)

// Create an enhanced history that syncs navigation events with the store
const history = syncHistoryWithStore(browserHistory, store)

function main() {  
  ReactDOM.render(
    <Provider store={store}>
      <Router history={history}>
        <Route path="/" component={Layout}>
          {engines.routes}
        </Route>
      </Router>
    </Provider>,
    document.getElementById('root')
  );
}

export default main
