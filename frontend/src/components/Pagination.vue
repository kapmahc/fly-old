<template>
  <nav aria-label="Page navigation">
    <ul :class="`pagination ${size ? 'pagination-'+size : ''}`">
      <li class="page-item">
        <router-link class="page-link" :to="url(first)" aria-label="Previous">
          <span aria-hidden="true">&laquo;</span>
          <span class="sr-only">{{$t("buttons.first")}}</span>
        </router-link>
      </li>
      <li :class="`page-item ${it === item.page ? 'active' : ''}`" :key="it" v-for="it in this.pages">
        <router-link class="page-link" :to="url(it)">{{it}}</router-link>
      </li>
      <li class="page-item">
        <router-link class="page-link" :to="url(last)" aria-label="Next">
          <span aria-hidden="true">&raquo;</span>
          <span class="sr-only">{{$t("buttons.last")}}</span>
        </router-link>
      </li>
    </ul>
  </nav>
</template>

<script>
export default {
  name: 'pagination-panel',
  props: ['item', 'href', 'size'],
  data () {
    return {
      first: 1,
      last: 1,
      pages: []
    }
  },
  watch: {
    'item': 'refresh'
  },
  methods: {
    refresh () {
      var count = Math.ceil(this.item.total / this.item.size)
      var pages = []
      var len = 5
      var begin = 0
      var end = 0
      if (count <= len * 2) {
        begin = 1
        end = count
      } else {
        begin = this.item.page - len
        if (begin <= 0) {
          begin = 1
        }

        end = begin + len * 2
        if (end > count) {
          end = count
          begin = count - len * 2
        }
      }

      for (var i = begin; i <= end; i++) {
        pages.push(i)
      }

      this.last = count
      this.pages = pages
    },

    url (page) {
      var u = Object.assign({}, this.href)
      if (!u.params) {
        u.params = {}
      }
      u.params.page = page
      return u
    }
  }
}
</script>
