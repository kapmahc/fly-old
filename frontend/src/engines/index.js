import auth from './auth'
import site from './site'
import blog from './blog'

const engines = {
  auth,
  blog,
  // forum,
  // reading
  site
}

export const dashboard = Object.keys(engines).map((k, i) => {
  return engines[k].dashboard
}, [])

export const links = Object.keys(engines).reduce((a, k) => {
  return a.concat(engines[k].links)
}, [])

export const routes = Object.keys(engines).reduce((a, k) => {
  return a.concat(engines[k].routes)
}, [])
