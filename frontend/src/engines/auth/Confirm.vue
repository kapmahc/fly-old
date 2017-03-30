<template>
  <div>
    <h2>{{$t("auth.users.confirm.title")}}</h2>
    <hr/>
    <form v-on:submit.prevent="onSubmit">
      <div class="form-group">
        <label for="email">{{$t("attributes.email")}}</label>
        <input type="email" class="form-control" id="email" v-model="item.email">
      </div>
      <form-buttons />
    </form>
  </div>
</template>

<script>
import {post} from '@/ajax'
import Buttons from '@/components/FormButtons'

export default {
  name: 'auth-users-confirm',
  data () {
    return {
      item: {
        email: ''
      }
    }
  },
  components: {
    'form-buttons': Buttons
  },
  methods: {
    onSubmit () {
      var data = new FormData()
      data.append('email', this.item.email)

      post('/users/confirm', data).then(function (rst) {
        alert(this.$t('auth.messages.email-for-confirm'))
        this.$router.push({ name: 'auth.users.sign-in' })
      }.bind(this)).catch((err) => {
        alert(err)
      })
    }
  }
}
</script>
