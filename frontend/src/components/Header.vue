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
          <li v-for="l in links">
            <router-link class="nav-link" :to="{name: l.href}">{{ $t(l.label) }}</router-link>
          </li>
        </ul>
        <form class="form-inline mt-2 mt-md-0">
          <input class="form-control mr-sm-2" type="text" placeholder="Search">
          <button class="btn btn-outline-success my-2 my-sm-0" type="submit">Search</button>
        </form>
      </div>
    </nav>
</template>

<script>
import {links} from '@/engines'

export default {
  name: 'application-header',
  data () {
    return {links}
  },
  computed: {
    info () {
      return this.$store.state.siteInfo
    }
  }
}
</script>
