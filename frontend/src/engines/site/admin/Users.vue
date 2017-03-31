<template>
  <dashboard-layout>
    <h2>
      {{$t('site.admin.users.index.title')}}
    </h2>
    <hr/>

    <modal-form size="sm" :id="mid" :title="$t('buttons.view')">      
      <div class="card">
        <img class="card-img-top" :src="cur.logo" :alt="cur.logo">
        <div class="card-block">
          <h4 class="card-title">{{cur.name}}</h4>
          <p class="card-text">{{cur.email}}</p>
          <a :href="cur.home" class="card-link" target="_blank">{{$t("auth.attributes.user.home")}}</a>
        </div>
      </div>
    </modal-form>

    <table class="table table-bordered table-hover">
      <thead>
        <tr>
          <th>{{$t('auth.attributes.user.info')}}</th>
          <th>{{$t('auth.attributes.user.lastSignIn')}}</th>
          <th>{{$t('auth.attributes.user.currentSignIn')}}</th>
        </tr>
      </thead>
      <tbody>
        <tr :key="it.id" v-for="it in items"  data-toggle="modal" :data-target="`#${mid}`" v-on:click="onEdit(it)">
          <th scope="row">{{it.name}}&lt;{{it.email}}&gt;</th>
          <td>{{it.lastSignInIp}} - {{it.lastSignInAt}}</td>
          <td>{{it.currentSignInIp}} - {{it.currentSignInAt}}</td>
        </tr>
      </tbody>
    </table>

  </dashboard-layout>
</template>

<script>
import Layout from '@/layouts/Dashboard'
import {get} from '@/ajax'
import Modal from '@/components/Modal'

export default {
  name: 'site-admin-users',
  data () {
    return {
      mid: 'editUserModal',
      items: [],
      cur: {}
    }
  },
  components: {
    'dashboard-layout': Layout,
    'modal-form': Modal
  },
  beforeCreate () {
    get('/admin/users').then(function (rst) {
      this.items = rst
    }.bind(this))
  },

  methods: {
    onEdit (it) {
      this.cur = it
    }
  }
}
</script>
