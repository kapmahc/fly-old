<template>
  <dashboard-layout>
    <h2>
      {{$t('site.admin.notices.index.title')}}
      <button class="btn btn-primary" v-on:click="onEdit()" data-toggle="modal" :data-target="`#${mid}`">{{$t('buttons.new')}}</button>
    </h2>
    <hr/>

    <modal-form size='lg' :id="mid" :onRemove="cur.id ? onRemove : null" :onSave="onSave" :title="$t(cur.id ? 'buttons.edit' : 'buttons.new')">
      <div class="form-group">
        <label for="body">{{$t("attributes.body")}}</label>
        <textarea class="form-control" id="body" v-model="cur.body" rows="3" aria-describedby="bodyHelp"></textarea>
        <small id="bodyHelp" class="form-text text-muted">{{$t('helps.markdown')}}</small>
      </div>
    </modal-form>

    <table class="table table-bordered table-hover">
      <thead>
        <tr>
          <th>{{$t('attributes.updatedAt')}}</th>
          <th>{{$t('attributes.body')}}</th>
        </tr>
      </thead>
      <tbody>
        <tr :key="it.id" v-for="it in items"  data-toggle="modal" :data-target="`#${mid}`" v-on:click="onEdit(it)">
          <th scope="row">{{it.updatedAt}}</th>
          <td><pre><code>{{it.body}}</code></pre></td>
        </tr>
      </tbody>
    </table>

  </dashboard-layout>
</template>

<script>
import $ from 'jquery'

import Layout from '@/layouts/Dashboard'
import {get, post, _delete} from '@/ajax'
import Modal from '@/components/Modal'
import {MARKDOWN} from '@/constants'

export default {
  name: 'site-admin-notices',
  data () {
    return {
      mid: 'editNoticeModal',
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
    get('/notices').then(function (rst) {
      this.items = rst
    }.bind(this))
  },

  methods: {
    onSave () {
      var data = new FormData()
      data.append('type', MARKDOWN)
      data.append('body', this.cur.body)
      var id = this.cur.id
      post(id ? `/notices/${id}` : '/notices', data).then(function (rst) {
        alert(this.$t('flashs.success'))
        var items = this.items.filter((it) => it.id !== rst.id)
        items.unshift(rst)
        this.cur = rst
        this.items = items
        $(`#${this.mid}`).modal('hide')
      }.bind(this)).catch((err) => {
        alert(err)
      })
    },
    onRemove (it) {
      if (confirm(this.$t('are-you-sure'))) {
        var id = this.cur.id
        _delete(`/admin/locales/${id}`).then(function (rst) {
          var items = this.items.filter((it) => it.id !== id)
          this.items = items
          this.cur = {body: ''}
          $(`#${this.mid}`).modal('hide')
        }.bind(this)).catch((err) => alert(err))
      }
    },
    onEdit (it) {
      if (it) {
        this.cur = it
      } else {
        this.cur = {body: ''}
      }
    }
  }
}
</script>
