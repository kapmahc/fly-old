<template>
  <application-layout>
    <h2>{{$t('reading.books.index.title')}}</h2>
    <hr/>
    <pagination-panel :item="pager" :href="{name: 'reading.books.pages'}"/>
    <div class="card-columns">
      <div class="card" v-for="it in pager.items">
        <div class="card-block">
          <router-link :to="{name: 'reading.books.show', params: {id: it.id}}">
            <h4 class="card-title">{{it.title}}</h4>
          </router-link>
          <p class="card-text">{{it.description}}</p>
        </div>
        <ul class="list-group list-group-flush">
          <li :key="k" v-for="k in ['author', 'subject', 'publisher']" class="list-group-item list-group-item-action">
            {{$t(`reading.attributes.book.${k}`)}}: {{it[k]}}
          </li>
          <li class="list-group-item list-group-item-action">{{$t('reading.attributes.book.publishedAt')}}: {{new Date(it.publishedAt).toLocaleDateString()}}</li>
        </ul>
        <div class="card-block">
          <router-link class="btn btn-primary" :to="{name: 'reading.books.show', params: {id: it.id}}">
            {{$t("buttons.view")}}
          </router-link>
        </div>
      </div>
    </div>

  </application-layout>
</template>

<script>
import Application from '@/layouts/Application'
import Pagination from '@/components/Pagination'
import {get} from '@/ajax'

export default {
  name: 'reading-books-index',
  components: {
    'application-layout': Application,
    'pagination-panel': Pagination
  },
  data () {
    return {
      page: null,
      pager: {
        total: 0,
        size: 60
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
      var page = this.$route.params.page
      if (page) {
        get(`/reading/books?page=${page}`).then(function (rst) {
          this.pager = rst
        }.bind(this))
      }
    }
  }
}
</script>
