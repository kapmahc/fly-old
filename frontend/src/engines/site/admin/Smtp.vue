<template>
  <dashboard-layout>
    <form-panel :title="$t('site.admin.info.title')" :onSubmit="onSubmit">
      <div class="form-group">
        <label for="host">{{$t("attributes.host")}}</label>
        <input type="text" class="form-control" id="host" v-model="item.host">
      </div>
      <div class="form-group">
        <label for="port">{{$t("attributes.port")}}</label>
        <select class="form-control" id="port" v-model="item.port">
          <option>25</option>
          <option>465</option>
          <option>587</option>
        </select>
      </div>
      <div class="form-group">
        <label for="username">{{$t("site.admin.smtp.sender")}}</label>
        <input type="email" class="form-control" id="username" v-model="item.username">
      </div>
      <div class="form-group">
        <label for="password">{{$t("attributes.password")}}</label>
        <input type="password" class="form-control" v-model="item.password" id="password" aria-describedby="passwordHelp">
        <small id="passwordHelp" class="form-text text-muted">{{$t("helps.password")}}</small>
      </div>
      <div class="form-group">
        <label for="passwordConfirmation">{{$t("attributes.passwordConfirmation")}}</label>
        <input type="password" class="form-control" v-model="item.passwordConfirmation" id="passwordConfirmation" aria-describedby="passwordConfirmationHelp">
        <small id="passwordConfirmationHelp" class="form-text text-muted">{{$t("helps.passwordConfirmation")}}</small>
      </div>
      <div class="form-check">
        <label class="form-check-label">
          <input type="checkbox" class="form-check-input" id="ssl" v-model="item.ssl">
          {{$t("attributes.ssl")}}
        </label>
      </div>
    </form-panel>
  </dashboard-layout>
</template>

<script>
import {post, get} from '@/ajax'
import Layout from '@/layouts/Dashboard'
import Form from '@/components/Form'

export default {
  name: 'site-admin-smtp',
  data () {
    return {
      item: {
        host: '',
        port: 25,
        username: '',
        password: '',
        ssl: false
      }
    }
  },
  beforeCreate () {
    get('/admin/site/smtp').then(function (rst) {
      this.item = rst
    }.bind(this))
  },
  components: {
    'dashboard-layout': Layout,
    'form-panel': Form
  },
  methods: {
    onSubmit () {
      var data = new FormData()
      data.append('host', this.item.host)
      data.append('port', this.item.port)
      data.append('username', this.item.username)
      data.append('password', this.item.password)
      data.append('passwordConfirmation', this.item.passwordConfirmation)
      data.append('ssl', this.item.ssl)

      post('/admin/site/smtp', data).then(function (rst) {
        alert(this.$t('flashs.success'))
      }.bind(this)).catch((err) => {
        alert(err)
      })
    }
  }
}
</script>
