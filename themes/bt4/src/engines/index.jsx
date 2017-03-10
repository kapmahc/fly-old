import auth from './auth'
import site from './site'

const engines = {
  auth,
  site
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
