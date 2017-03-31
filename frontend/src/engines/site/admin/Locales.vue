<template>
  <dashboard-layout>
    <h2>
      {{$t('site.admin.locales.index.title')}}
      <button class="btn btn-primary" v-on:click="onEdit()" data-toggle="modal" :data-target="`#${mid}`">{{$t('buttons.new')}}</button>
    </h2>
    <hr/>

    <modal-form :id="mid" :onRemove="cur.id ? onRemove : null" :onSave="onSave" :title="$t(cur.id ? 'buttons.edit' : 'buttons.new')">
      <div class="form-group">
        <label for="code">{{$t("site.attributes.locale.code")}}</label>
        <input type="text" class="form-control" id="code" v-model="cur.code">
      </div>
      <div class="form-group">
        <label for="message">{{$t("site.attributes.locale.message")}}</label>
        <textarea class="form-control" id="message" v-model="cur.message" rows="3"></textarea>
      </div>
    </modal-form>

    <ul class="list-group">
      <li data-toggle="modal" :data-target="`#${mid}`" v-on:click="onEdit(it)" :key="it.code" v-for="it in items" class="list-group-item list-group-item-action">
        {{it.code}} = {{it.message}}
      </li>
    </ul>
  </dashboard-layout>
</template>

<script>
import Layout from '@/layouts/Dashboard'
import {get, post, _delete} from '@/ajax'
import Modal from '@/components/Modal'
import $ from 'jquery'

export default {
  name: 'site-admin-locales',
  data () {
    return {
      mid: 'editLocaleModal',
      items: [],
      cur: {
        code: '',
        message: ''
      }
    }
  },
  components: {
    'dashboard-layout': Layout,
    'modal-form': Modal
  },
  beforeCreate () {
    get('/admin/locales').then(function (rst) {
      this.items = rst
    }.bind(this))
  },

  methods: {
    onSave () {
      var data = new FormData()
      data.append('code', this.cur.code)
      data.append('message', this.cur.message)

      post('/admin/locales', data).then(function (rst) {
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
        _delete(`/admin/locales/${this.cur.id}`).then(function (rst) {
          var items = this.items.filter((it) => it.id !== this.cur.id)
          this.items = items
          this.cur = {code: '', message: ''}
          $(`#${this.mid}`).modal('hide')
        }.bind(this)).catch((err) => alert(err))
      }
    },
    onEdit (it) {
      if (it) {
        this.cur = it
      } else {
        this.cur = {code: '', message: ''}
      }
    }
  }
}
</script>
