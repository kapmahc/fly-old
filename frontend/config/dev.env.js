var merge = require('webpack-merge')
var prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
  NODE_ENV: '"development"',
  BACKEND: '"http://localhost:3000"',
  LOCALE: '"en-US"'
})
