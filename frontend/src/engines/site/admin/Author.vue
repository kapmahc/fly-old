<template>
  <dashboard-layout>
    <form-panel :title="$t('site.admin.author.title')" :onSubmit="onSubmit">
      <div class="form-group">
        <label for="name">{{$t("attributes.fullName")}}</label>
        <input type="text" class="form-control" id="name" v-model="item.name">
      </div>
      <div class="form-group">
        <label for="email">{{$t("attributes.email")}}</label>
        <input type="email" class="form-control" id="email" v-model="item.email">
      </div>      
    </form-panel>
  </dashboard-layout>
</template>

<script>
import {post, get} from '@/ajax'
import Layout from '@/layouts/Dashboard'
import Form from '@/components/Form'

export default {
  name: 'site-admin-author',
  data () {
    return {
      item: {
        name: '',
        email: ''
      }
    }
  },
  beforeCreate () {
    get('/site/info').then(function (rst) {
      this.item = rst.author
    }.bind(this))
  },
  components: {
    'dashboard-layout': Layout,
    'form-panel': Form
  },
  methods: {
    onSubmit () {
      var data = new FormData()
      data.append('name', this.item.name)
      data.append('email', this.item.email)

      post('/admin/site/author', data).then(function (rst) {
        alert(this.$t('flashs.success'))
      }.bind(this)).catch((err) => {
        alert(err)
      })
    }
  }
}
</script>
