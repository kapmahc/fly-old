import auth from './auth'
import site from './site'

const engines = {
  auth,
  site
}

export default {
  routes: Object.keys(engines).reduce((a, k) => {
    return a.concat(engines[k].routes)
  }, [])
};
