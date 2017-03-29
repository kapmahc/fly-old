<template>
  <div class="container">
    <hr class="featurette-divider">
      <footer>
        <p class="float-right">
          {{ $t("footer.other-languages") }}
          <button class="btn btn-sm btn-link" v-on:click="setLanguage(l)" v-for="l in info.languages">
            {{$t(`languages.${l}`)}}
          </button>
        </p>
        <p>
          &copy; {{ $t("site.copyright") }}
          &middot; <a href="#">Privacy</a>
          &middot; <a href="#">Terms</a>
        </p>
      </footer>
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
