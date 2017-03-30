<template>
  <div>
    <h2>{{$t("auth.users.sign-up.title")}}</h2>
    <hr/>
    <form v-on:submit.prevent="onSubmit">
      <div class="form-group">
        <label for="name">{{$t("attributes.fullName")}}</label>
        <input type="text" class="form-control" id="name" v-model="item.name">
      </div>
      <div class="form-group">
        <label for="email">{{$t("attributes.email")}}</label>
        <input type="email" class="form-control" id="email" v-model="item.email">
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
      <form-buttons />
    </form>
  </div>
</template>

<script>
import {post} from '@/ajax'
import Buttons from '@/components/FormButtons'

export default {
  name: 'auth-users-sign-up',
  data () {
    return {
      item: {
        name: '',
        email: '',
        password: '',
        passwordConfirmation: ''
      }
    }
  },
  components: {
    'form-buttons': Buttons
  },
  methods: {
    onSubmit () {
      var data = new FormData()
      data.append('name', this.item.name)
      data.append('email', this.item.email)
      data.append('password', this.item.password)
      data.append('passwordConfirmation', this.item.passwordConfirmation)

      post('/users/sign-up', data).then(function (rst) {
        alert(this.$t('auth.messages.email-for-confirm'))
        this.$router.push({ name: 'auth.users.sign-in' })
      }.bind(this)).catch((err) => {
        alert(err)
      })
    }
  }
}
</script>
