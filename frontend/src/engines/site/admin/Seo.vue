<template>
  <dashboard-layout>
    <form-panel :title="$t('site.admin.seo.title')" :onSubmit="onSubmit">
      <div class="form-group">
        <label for="googleVerifyCode">{{$t("site.attributes.googleVerifyCode")}}</label>
        <input type="text" class="form-control" id="googleVerifyCode" v-model="item.googleVerifyCode">
      </div>
      <div class="form-group">
        <label for="baiduVerifyCode">{{$t("site.attributes.baiduVerifyCode")}}</label>
        <input type="text" class="form-control" id="baiduVerifyCode" v-model="item.baiduVerifyCode">
      </div>
    </form-panel>
  </dashboard-layout>
</template>

<script>
import {post, get} from '@/ajax'
import Layout from '@/layouts/Dashboard'
import Form from '@/components/Form'

export default {
  name: 'site-admin-seo',
  data () {
    return {
      item: {
        googleVerifyCode: '',
        baiduVerifyCode: ''
      }
    }
  },
  beforeCreate () {
    get('/admin/site/seo').then(function (rst) {
      this.item = rst
    }.bind(this))
  },
  components: {
    'dashboard-layout': Layout,
    'form-panel': Form
  },
  methods: {
    onSubmit () {
      var data = new FormData()
      data.append('googleVerifyCode', this.item.googleVerifyCode)
      data.append('baiduVerifyCode', this.item.baiduVerifyCode)

      post('/admin/site/seo', data).then(function (rst) {
        alert(this.$t('flashs.success'))
      }.bind(this)).catch((err) => {
        alert(err)
      })
    }
  }
}
</script>
