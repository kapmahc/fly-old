<template>
  <non-sign-in-layout :title="$t('site.leave-words.new.title')" :onSubmit="onSubmit">
    <div class="form-group">
      <label for="body">{{$t("attributes.body")}}</label>
      <textarea class="form-control" id="body" v-model="item.body" rows="3" aria-describedby="bodyHelp"></textarea>
      <small id="bodyHelp" class="form-text text-muted">{{$t('site.helps.leave-word.body')}}</small>
    </div>
  </non-sign-in-layout>
</template>

<script>
import {post} from '@/ajax'
import Layout from '@/engines/auth/users/NonSignIn'
import {MARKDOWN} from '@/constants'

export default {
  name: 'site-leave-words-new',
  data () {
    return {
      item: {
        body: ''
      }
    }
  },
  components: {
    'non-sign-in-layout': Layout
  },
  methods: {
    onSubmit () {
      var data = new FormData()
      data.append('type', MARKDOWN)
      data.append('body', this.item.body)

      post('/leave-words', data).then(function (rst) {
        alert(this.$t('flashs.success'))
        this.item.body = ''
      }.bind(this)).catch((err) => {
        alert(err)
      })
    }
  }
}
</script>
