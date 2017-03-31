<template>
  <application-layout>
    <h2>{{$t("site.notices.index.title")}}</h2>
    <hr/>
    <blockquote class="blockquote" :key="it.id" v-for="it in items">
      <p class="mb-0">
        <markdown-panel :body="it.body"/>
      </p>
      <footer class="blockquote-footer">
        <time-ago :date="it.updatedAt"/>        
      </footer>
    </blockquote>
  </application-layout>
</template>

<script>
import Application from '@/layouts/Application'
import {get} from '@/ajax'
import Markdown from '@/components/Markdown'
import TimeAgo from '@/components/TimeAgo'

export default {
  name: 'site-notices-index',
  data () {
    return {
      items: []
    }
  },
  components: {
    'application-layout': Application,
    'markdown-panel': Markdown,
    'time-ago': TimeAgo
  },
  beforeCreate () {
    get('/notices').then(function (rst) {
      this.items = rst
    }.bind(this))
  }
}
</script>
