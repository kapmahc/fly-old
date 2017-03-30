<template>
  <non-sign-in-layout :title="$t(`auth.users.${action}.title`)" :onSubmit="onSubmit">
    <div class="form-group">
      <label for="email">{{$t("attributes.email")}}</label>
      <input type="email" class="form-control" id="email" v-model="item.email">
    </div>
  </non-sign-in-layout>
</template>

<script>
import {post} from '@/ajax'
import Layout from './NonSignIn'

export default {
  name: 'auth-email-form',
  props: ['action'],
  data () {
    return {
      item: {
        email: ''
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

      post(`/users/${this.action}`, data).then(function (rst) {
        alert(this.$t(`auth.messages.email-for-${this.action}`))
        this.$router.push({ name: 'auth.users.sign-in' })
      }.bind(this)).catch((err) => {
        alert(err)
      })
    }
  }
}
</script>
