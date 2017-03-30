<template>
  <application-layout v-if="user.uid">
    <ul class="nav nav-tabs">
      <li v-if="db" :key="db.label" class="nav-item dropdown" v-for="db in dashboard">
        <a class="nav-link dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">
          {{$t(db.label)}}
        </a>
        <div class="dropdown-menu">
          <dropdown-item :key="Math.random()" :item="it" v-for="it in db.items" />
        </div>
      </li>
    </ul>
    <br/>
    <h2>{{title}}</h2>    
    <slot />
  </application-layout>
  <application-layout v-else>
    <alert-dialog action="danger" :message="$t('errors.forbidden')" />
  </application-layout>
</template>

<script>
import Layout from './Application'
import Alert from '@/components/Alert'
import {dashboard} from '@/engines'

var DropdownItem = {
  name: 'dropdown-item',
  template: '<router-link v-if="item" :to="{name: item.href}" class="dropdown-item">{{$t(item.label)}}</router-link><div v-else class="dropdown-divider"></div>',
  props: ['item']
}

export default {
  name: 'dashboard-layout',
  props: ['title'],
  components: {
    'application-layout': Layout,
    'alert-dialog': Alert,
    'dropdown-item': DropdownItem
  },
  computed: {
    dashboard () {
      var user = this.$store.state.currentUser
      return dashboard.map(db => db(user))
    },
    user () {
      return this.$store.state.currentUser
    }
  }
}
</script>
