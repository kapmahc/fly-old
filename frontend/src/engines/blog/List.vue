<template>
  <div class="list-group">
    <router-link :key="it.href" :to="{name: 'blog.show', params: [it.href]}" v-for="it in items" class="list-group-item list-group-item-action">
      {{it.title}}
    </router-link>
  </div>
</template>

<script>
import {get} from '@/ajax'

export default {
  name: 'blog-list',
  data () {
    return {
      items: []
    }
  },
  beforeCreate () {
    get('/blogs').then(function (rst) {
      this.items = rst
    }.bind(this))
  }
}
</script>
