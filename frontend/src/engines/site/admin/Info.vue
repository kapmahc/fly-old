<template>
  <dashboard-layout>
    <form-panel :title="$t('site.admin.info.title')" :onSubmit="onSubmit">
      <div class="form-group">
        <label for="title">{{$t("site.attributes.title")}}</label>
        <input type="text" class="form-control" id="title" v-model="item.title">
      </div>
      <div class="form-group">
        <label for="subTitle">{{$t("site.attributes.subTitle")}}</label>
        <input type="text" class="form-control" id="subTitle" v-model="item.subTitle">
      </div>
      <div class="form-group">
        <label for="keywords">{{$t("site.attributes.keywords")}}</label>
        <input type="text" class="form-control" id="keywords" v-model="item.keywords">
      </div>
      <div class="form-group">
        <label for="description">{{$t("site.attributes.description")}}</label>
        <textarea class="form-control" id="description" v-model="item.description" rows="3"></textarea>
      </div>
      <div class="form-group">
        <label for="copyright">{{$t("site.attributes.copyright")}}</label>
        <input type="text" class="form-control" id="copyright" v-model="item.copyright">
      </div>
    </form-panel>
  </dashboard-layout>
</template>

<script>
import {post, get} from '@/ajax'
import Layout from '@/layouts/Dashboard'
import Form from '@/components/Form'

export default {
  name: 'site-admin-info',
  data () {
    return {
      item: {
        title: '',
        subTitle: '',
        keywords: '',
        description: '',
        copyright: ''
      }
    }
  },
  beforeCreate () {
    get('/site/info').then(function (rst) {
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
      data.append('title', this.item.title)
      data.append('subTitle', this.item.subTitle)
      data.append('keywords', this.item.keywords)
      data.append('description', this.item.description)
      data.append('copyright', this.item.copyright)

      post('/admin/site/info', data).then(function (rst) {
        alert(this.$t('flashs.success'))
      }.bind(this)).catch((err) => {
        alert(err)
      })
    }
  }
}
</script>
