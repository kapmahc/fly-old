<template>
  <dashboard-layout>
    <form-panel :title="$t('auth.users.info.title')" :onSubmit="onSubmit">
      <div class="form-group">
        <label for="name">{{$t("attributes.fullName")}}</label>
        <input type="text" class="form-control" id="name" v-model="item.name">
      </div>
      <div class="form-group">
        <label for="email">{{$t("attributes.email")}}</label>
        <input readonly type="email" class="form-control" id="email" v-model="item.email">
      </div>
      <div class="form-group">
        <label for="logo">{{$t("auth.attributes.user.logo")}}</label>
        <input type="text" class="form-control" id="logo" v-model="item.logo">
      </div>
      <div class="form-group">
        <label for="home">{{$t("auth.attributes.user.home")}}</label>
        <input type="text" class="form-control" id="home" v-model="item.home">
      </div>
    </form-panel>
  </dashboard-layout>
</template>

<script>
import {post, get} from '@/ajax'
import Layout from '@/layouts/Dashboard'
import Form from '@/components/Form'

export default {
  name: 'auth-users-info',
  data () {
    return {
      item: {
        name: '',
        email: '',
        logo: '',
        home: ''
      }
    }
  },
  components: {
    'dashboard-layout': Layout,
    'form-panel': Form
  },
  beforeCreate () {
    get('/users/info').then(function (rst) {
      this.item = rst
    }.bind(this))
  },
  methods: {
    onSubmit () {
      var data = new FormData()
      data.append('name', this.item.name)
      data.append('logo', this.item.logo)
      data.append('home', this.item.home)

      post('/users/info', data).then(function (rst) {
        alert(this.$t('flashs.success'))
      }.bind(this)).catch((err) => {
        alert(err)
      })
    }
  }
}
</script>
