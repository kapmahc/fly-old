<template>
  <dashboard-layout>
    <div class="card-columns">
      <div :key="k" class="card" v-for="k in ['os', 'network', 'database', 'jobs']">
        <div class="card-block">
          <h4 class="card-title">{{$t(`site.admin.status.${k}`)}}</h4>
          <p class="card-text"></p>
        </div>
        <ul class="list-group list-group-flush">
          <li class="list-group-item list-group-item-action" v-for="(it, id) in item[k]">
            {{id}}: {{it}}
          </li>
        </ul>
      </div>
      <div class="card">
        <div class="card-block">
          <h4 class="card-title">{{$t("site.admin.status.cache")}}</h4>
          <p class="card-text">
            <pre><code>{{item.cache}}</code></pre>
          </p>
        </div>
      </div>
    </div>
  </dashboard-layout>
</template>

<script>
import {get} from '@/ajax'
import Layout from '@/layouts/Dashboard'

export default {
  name: 'site-admin-status',
  data () {
    return {
      item: {
      }
    }
  },
  beforeCreate () {
    get('/admin/site/status').then(function (rst) {
      this.item = rst
    }.bind(this))
  },
  components: {
    'dashboard-layout': Layout
  }
}
</script>

<style scoped>
.card-columns {
  column-count: 2;
}
</style>
