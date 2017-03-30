<template>
  <application-layout>
    <div class="row" v-if="user.uid">
      <div :key="db.label" class="col-md-4" v-for="db in dashboard" v-if="db">
        <h4>{{$t(db.label)}}</h4>
        <hr/>
        <div class="list-group">
          <router-link :key="it.label" v-for="it in db.items" v-if="it" :to="{name: it.href}" class="list-group-item list-group-item-action">
            {{$t(it.label)}}
          </router-link>
        </div>
      </div>
    </div>

    <alert-dialog v-else action="danger" :message="$t('errors.forbidden')" />

  </application-layout>
</template>

<script>
import Layout from '@/layouts/Application'
import {dashboard} from '@/engines'
import Alert from '@/components/Alert'

export default {
  name: 'site-dashboard',
  components: {
    'alert-dialog': Alert,
    'application-layout': Layout
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
