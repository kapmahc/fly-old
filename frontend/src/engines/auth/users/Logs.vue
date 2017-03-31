<template>
  <dashboard-layout>
    <h2>{{$t('auth.users.logs.title')}}</h2>
    <hr/>
    <ul class="list-group">
      <li v-for="it in items" class="list-group-item list-group-item-action">
        [{{it.ip}}] {{it.createdAt}}: {{it.message}}
      </li>
    </ul>
  </dashboard-layout>
</template>

<script>
import Layout from '@/layouts/Dashboard'
import {get} from '@/ajax'

export default {
  name: 'auth-users-logs',
  data () {
    return {
      items: []
    }
  },
  components: {
    'dashboard-layout': Layout
  },
  beforeCreate () {
    get('/users/logs').then(function (rst) {
      this.items = rst
    }.bind(this))
  }
}
</script>
