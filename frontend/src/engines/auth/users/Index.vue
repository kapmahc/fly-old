<template>
  <application-layout>
    <h2>{{$t("auth.users.index.title")}}</h2>
    <hr/>
    <div class="card-columns">
      <div class="card" :key="it.id" v-for="it in items">
        <img class="card-img-top img-fluid" :src="it.logo" :alt="it.logo">
        <div class="card-block">
          <h4 class="card-title">{{it.name}}</h4>
          <p class="card-text"></p>
          <a :href="it.home" target="_blank" class="card-link">{{$t('buttons.view')}}</a>
        </div>
      </div>
    </div>
  </application-layout>
</template>

<script>
import Application from '@/layouts/Application'
import {get} from '@/ajax'

export default {
  name: 'auth-users-index',
  data () {
    return {
      items: []
    }
  },
  components: {
    'application-layout': Application
  },
  beforeCreate () {
    get('/users').then(function (rst) {
      this.items = rst
    }.bind(this))
  }
}
</script>
