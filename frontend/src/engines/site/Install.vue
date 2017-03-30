<template>
  <form-layout :title="$t('site.install.title')" :onSubmit="onSubmit">
    <div class="form-group">
      <label for="title">{{$t("site.attributes.title")}}</label>
      <input type="text" class="form-control" id="title" v-model="item.title">
    </div>
    <div class="form-group">
      <label for="subTitle">{{$t("site.attributes.subTitle")}}</label>
      <input type="text" class="form-control" id="subTitle" v-model="item.subTitle">
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
  </form-layout>
</template>

<script>
import {post} from '@/ajax'
import Form from '@/layouts/Form'

export default {
  name: 'site-install',
  data () {
    return {
      item: {
        title: '',
        subTitle: '',
        email: '',
        password: '',
        passwordConfirmation: ''
      }
    }
  },
  components: {
    'form-layout': Form
  },
  methods: {
    onSubmit () {
      var data = new FormData()
      data.append('title', this.item.title)
      data.append('subTitle', this.item.subTitle)
      data.append('email', this.item.email)
      data.append('password', this.item.password)
      data.append('passwordConfirmation', this.item.passwordConfirmation)

      post('/install', data).then(function (rst) {
        alert(this.$t('flashs.success'))
        this.$router.push({ name: 'auth.users.sign-in' })
      }.bind(this)).catch((err) => {
        alert(err)
      })
    }
  }
}
</script>
