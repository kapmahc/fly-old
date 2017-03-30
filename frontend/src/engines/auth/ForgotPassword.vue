<template>
  <div>
    <h2>{{$t("auth.users.forgot-password.title")}}</h2>
    <hr/>
    <form v-on:submit.prevent="onSubmit">
      <div class="form-group">
        <label for="email">{{$t("attributes.email")}}</label>
        <input type="email" class="form-control" id="email" v-model="item.email">
      </div>
      <form-buttons />
    </form>
    <br/>
    <shared-links />
  </div>
</template>

<script>
import {post} from '@/ajax'
import Links from './Links'
import Buttons from '@/components/FormButtons'

export default {
  name: 'auth-users-forgot-password',
  data () {
    return {
      item: {
        email: ''
      }
    }
  },
  components: {
    'shared-links': Links,
    'form-buttons': Buttons
  },
  methods: {
    onSubmit () {
      var data = new FormData()
      data.append('email', this.item.email)

      post('/users/forgot-password', data).then(function (rst) {
        alert(this.$t('auth.messages.email-for-reset-password'))
        this.$router.push({ name: 'auth.users.sign-in' })
      }.bind(this)).catch((err) => {
        alert(err)
      })
    }
  }
}
</script>
