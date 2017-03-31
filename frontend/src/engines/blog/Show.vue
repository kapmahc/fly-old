<template>
  <application-layout>
    <h2>{{cur.title}}</h2>
    <hr/>
    <markdown-panel :body="cur.body"/>
  </application-layout>
</template>

<script>
import Application from '@/layouts/Application'
import List from './List'
import {get} from '@/ajax'
import Markdown from '@/components/Markdown'

export default {
  name: 'blog-show',
  components: {
    'application-layout': Application,
    'blog-list': List,
    'markdown-panel': Markdown
  },
  data () {
    return {
      cur: {
        title: '',
        body: ''
      }
    }
  },
  created () {
    this.fetchData()
  },
  watch: {
    '$route': 'fetchData'
  },
  methods: {
    fetchData () {
      get(`/blog/${this.$route.params[0]}`).then(function (rst) {
        this.cur = rst
      }.bind(this)).catch((err) => alert(err))
    }
  }
}
</script>
