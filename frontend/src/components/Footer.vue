<template>
  <div class="container">
    <hr class="featurette-divider">
      <footer>
        <p class="float-right">
          <a href="#">Back to top</a>
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
  }
}
</script>
