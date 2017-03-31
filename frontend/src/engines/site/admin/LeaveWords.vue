<template>
  <dashboard-layout>
    <h2>
      {{$t('site.admin.leave-words.index.title')}}
    </h2>
    <hr/>

    <modal-form size='lg' :id="mid" :onRemove="onRemove" :title="$t('buttons.view')">
      <p>
        <pre><code>{{cur.body}}</code></pre>
      </p>
    </modal-form>

    <table class="table table-bordered table-hover">
      <thead>
        <tr>
          <th>{{$t('attributes.createdAt')}}</th>
          <th>{{$t('attributes.body')}}</th>
        </tr>
      </thead>
      <tbody>
        <tr :key="it.id" v-for="it in items"  data-toggle="modal" :data-target="`#${mid}`" v-on:click="onEdit(it)">
          <th scope="row">{{it.createdAt}}</th>
          <td><pre><code>{{it.body}}</code></pre></td>
        </tr>
      </tbody>
    </table>

  </dashboard-layout>
</template>

<script>
import $ from 'jquery'

import Layout from '@/layouts/Dashboard'
import {get, _delete} from '@/ajax'
import Modal from '@/components/Modal'

export default {
  name: 'site-admin-leave-words',
  data () {
    return {
      mid: 'editLeaveWordModal',
      items: [],
      cur: {
        body: ''
      }
    }
  },
  components: {
    'dashboard-layout': Layout,
    'modal-form': Modal
  },
  beforeCreate () {
    get('/leave-words').then(function (rst) {
      this.items = rst
    }.bind(this))
  },

  methods: {
    onRemove (it) {
      if (confirm(this.$t('are-you-sure'))) {
        var id = this.cur.id
        _delete(`/leave-words/${id}`).then(function (rst) {
          var items = this.items.filter((it) => it.id !== id)
          this.items = items
          this.cur = {body: ''}
          $(`#${this.mid}`).modal('hide')
        }.bind(this)).catch((err) => alert(err))
      }
    },
    onEdit (it) {
      this.cur = it
    }
  }
}
</script>
