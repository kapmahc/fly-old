<template>
  <dashboard-layout>
    <form-panel :title="$t('auth.users.change-password.title')" :onSubmit="onSubmit">
      <div class="form-group">
        <label for="currentPassword">{{$t("attributes.currentPassword")}}</label>
        <input type="password" class="form-control" v-model="item.currentPassword" id="currentPassword">
      </div>
      <div class="form-group">
        <label for="newPassword">{{$t("attributes.newPassword")}}</label>
        <input type="password" class="form-control" v-model="item.newPassword" id="password" aria-describedby="passwordHelp">
        <small id="passwordHelp" class="form-text text-muted">{{$t("helps.password")}}</small>
      </div>
      <div class="form-group">
        <label for="passwordConfirmation">{{$t("attributes.passwordConfirmation")}}</label>
        <input type="password" class="form-control" v-model="item.passwordConfirmation" id="passwordConfirmation" aria-describedby="passwordConfirmationHelp">
        <small id="passwordConfirmationHelp" class="form-text text-muted">{{$t("helps.passwordConfirmation")}}</small>
      </div>
    </form-panel>
  </dashboard-layout>
</template>

<script>
import {post} from '@/ajax'
import Layout from '@/layouts/Dashboard'
import Form from '@/components/Form'

export default {
  name: 'auth-users-change-password',
  data () {
    return {
      item: {
        currentPassword: '',
        newPassword: '',
        passwordConfirmation: ''
      }
    }
  },
  components: {
    'dashboard-layout': Layout,
    'form-panel': Form
  },
  methods: {
    onSubmit () {
      var data = new FormData()
      data.append('currentPassword', this.item.currentPassword)
      data.append('newPassword', this.item.newPassword)
      data.append('passwordConfirmation', this.item.passwordConfirmation)

      post('/users/change-password', data).then(function (rst) {
        alert(this.$t('flashs.success'))
      }.bind(this)).catch((err) => {
        alert(err)
      })
    }
  }
}
</script>
