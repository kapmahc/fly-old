import auth from './auth'

const engines = {
  auth,
}

export default {
  routes: Object.keys(engines).reduce((a, k) => {    
    return a.concat(engines[k].routes)
  }, [])
};
