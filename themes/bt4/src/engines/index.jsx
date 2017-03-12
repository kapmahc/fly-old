import auth from './auth'
import site from './site'
import blog from './blog'
import reading from './reading'

const engines = {
  auth,
  site,
  blog,
  reading
}

export default {
  dashboard: Object.keys(engines).map((k, i) => {
    return engines[k].dashboard
  }, []),
  navLinks: Object.keys(engines).reduce((a, k) => {
    return a.concat(engines[k].navLinks)
  }, []),
  routes: Object.keys(engines).reduce((a, k) => {
    return a.concat(engines[k].routes)
  }, [])
};
