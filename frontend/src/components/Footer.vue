<template>
  <div class="row">
    <div class="col-md-12">
      <hr/>
      <footer>
        <p class="float-right">
          {{ $t("footer.other-languages") }}
          <button :key="l" class="btn btn-sm btn-link" v-on:click="setLanguage(l)" v-for="l in info.languages">
            {{$t(`languages.${l}`)}}
          </button>
        </p>
        <p>
          &copy; {{ $t("site.copyright") }}
          &middot; <router-link :to="{name: 'home'}">{{$t("footer.privacy")}}</router-link>
          &middot; <router-link :to="{name: 'home'}">{{$t("footer.teams")}}</router-link>
        </p>
      </footer>
    </div>
  </div>
</template>

<script>
import {get} from '@/ajax'
import {LOCALE} from '@/constants'

export default {
  name: 'application-footer',
  computed: {
    info () {
      return this.$store.state.siteInfo
    }
  },
  beforeCreate () {
    get('/site/info').then(function (rst) {
      document.title = rst.title
      this.$store.commit('refresh', rst)
    }.bind(this))
  },
  methods: {
    setLanguage (l) {
      localStorage.setItem(LOCALE, l)
      location.reload()
    }
  }
}
</script>
