<template>
  <div>
    <h2>{{$t("auth.users.sign-in.title")}}</h2>
    <hr/>
    <form v-on:submit.prevent="onSubmit">
      <div class="form-group">
        <label for="email">{{$t("attributes.email")}}</label>
        <input type="email" class="form-control" id="email" v-model="item.email">
      </div>
      <div class="form-group">
        <label for="password">{{$t("attributes.password")}}</label>
        <input type="password" class="form-control" v-model="item.password" id="password" aria-describedby="passwordHelp">
      </div>
      <form-buttons />
    </form>
    <br/>
    <shared-links />
  </div>
</template>

<script>
import Links from './Links'
import {post} from '@/ajax'
import Buttons from '@/components/FormButtons'

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
    'shared-links': Links,
    'form-buttons': Buttons
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
