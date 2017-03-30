<template>
  <nav class="navbar navbar-toggleable-md navbar-inverse fixed-top bg-inverse">
      <button class="navbar-toggler navbar-toggler-right" type="button" data-toggle="collapse" data-target="#navbarCollapse" aria-controls="navbarCollapse" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <router-link class="navbar-brand" :to="{name: 'home'}">
        {{ $t("site.subTitle") }}
      </router-link>
      <div class="collapse navbar-collapse" id="navbarCollapse">
        <ul class="navbar-nav mr-auto">
          <li class="nav-item">
            <router-link class="nav-link" :to="{name: 'home'}">
              {{ $t("header.home") }} <span class="sr-only">(current)</span>
            </router-link>
          </li>
          <li :key="l.href" v-for="l in links">
            <router-link class="nav-link" :to="{name: l.href}">{{ $t(l.label) }}</router-link>
          </li>
          <li v-if="user.uid">
            <router-link class="nav-link" :to="{name: 'dashboard'}">{{ $t('header.dashboard') }}</router-link>
          </li>
          <language-bar />
          <personal-bar />
        </ul>
        <search-form />
      </div>
    </nav>
</template>

<script>
import {links} from '@/engines'

import SearchForm from './SearchForm'
import PersonalBar from './PersonalBar'
import LanguageBar from './LanguageBar'

export default {
  name: 'application-header',
  data () {
    return {links}
  },
  components: {
    'search-form': SearchForm,
    'personal-bar': PersonalBar,
    'language-bar': LanguageBar
  },
  computed: {
    info () {
      return this.$store.state.siteInfo
    },
    user () {
      return this.$store.state.currentUser
    }
  }
}
</script>
