<template>
  <non-sign-in-layout :title="$t('auth.users.reset-password.title')" :onSubmit="onSubmit">
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
  </non-sign-in-layout>
</template>

<script>
import {post} from '@/ajax'
import Layout from './NonSignIn'

export default {
  name: 'auth-users-reset-password',
  data () {
    return {
      item: {
        password: '',
        passwordConfirmation: ''
      }
    }
  },
  components: {
    'non-sign-in-layout': Layout
  },
  methods: {
    onSubmit () {
      var data = new FormData()
      data.append('token', this.$route.params.token)
      data.append('password', this.item.password)
      data.append('passwordConfirmation', this.item.passwordConfirmation)

      post('/users/reset-password', data).then(function (rst) {
        alert(this.$t('auth.messages.reset-password-success'))
        this.$router.push({ name: 'auth.users.sign-in' })
      }.bind(this)).catch((err) => {
        alert(err)
      })
    }
  }
}
</script>
