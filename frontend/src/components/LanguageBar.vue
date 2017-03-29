<template>
  <li class="nav-item dropdown">
    <a class="nav-link dropdown-toggle" id="language-bar" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
      {{$t("language-bar.switch")}}
    </a>
    <div class="dropdown-menu" aria-labelledby="language-bar">
      <a :key="l" v-on:click="setLanguage(l)" class="dropdown-item" v-for="l in languages">
        {{$t(`languages.${l}`)}}
      </a>
    </div>
  </li>
</template>

<script>
import {LOCALE} from '@/constants'

export default {
  name: 'language-bar',
  computed: {
    languages () {
      return this.$store.state.siteInfo.languages
    }
  },
  methods: {
    setLanguage (l) {
      localStorage.setItem(LOCALE, l)
      location.reload()
    }
  }
}
</script>
