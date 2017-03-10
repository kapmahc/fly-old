import { combineReducers } from 'redux'

import {USERS_SIGN_IN, USERS_SIGN_OUT, REFRESH_SITE_INFO} from './actions'

const currentUser = (state={}, action) => {
  switch(action.type){
    case USERS_SIGN_IN:
      console.log(action.token)
      return {
        uid: 'uuu',
        name: 'who am i'
      }
    case USERS_SIGN_OUT:
      return {}
    default:
      return state
  }
}

const siteInfo = (state={}, action) => {
  switch(action.type){
    case REFRESH_SITE_INFO:
      return Object.assign({}, action.info)
    default:
      return state;
  }
}

const app = combineReducers({
  currentUser,
  siteInfo
})

export default app
