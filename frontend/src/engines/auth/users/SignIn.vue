<template>
  <non-sign-in-layout :title="$t('auth.users.sign-in.title')" :onSubmit="onSubmit">
    <div class="form-group">
      <label for="email">{{$t("attributes.email")}}</label>
      <input type="email" class="form-control" id="email" v-model="item.email">
    </div>
    <div class="form-group">
      <label for="password">{{$t("attributes.password")}}</label>
      <input type="password" class="form-control" v-model="item.password" id="password" aria-describedby="passwordHelp">
    </div>
  </non-sign-in-layout>
</template>

<script>
import {post} from '@/ajax'
import Layout from './NonSignIn'

export default {
  name: 'auth-sign-in',
  data () {
    return {
      item: {
        email: '',
        password: ''
      }
    }
  },
  components: {
    'non-sign-in-layout': Layout
  },
  methods: {
    onSubmit () {
      var data = new FormData()
      data.append('email', this.item.email)
      data.append('password', this.item.password)

      post('/users/sign-in', data).then(function (rst) {
        this.$store.commit('signIn', rst.token)
        this.$router.push({ name: 'home' })
      }.bind(this)).catch((err) => {
        alert(err)
      })
    }
  }
}
</script>
