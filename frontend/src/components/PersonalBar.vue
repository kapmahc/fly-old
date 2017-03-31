<template>
  <li v-if="user.uid" class="nav-item dropdown">
    <a class="nav-link dropdown-toggle" id="personal-bar" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
      {{$t("personal-bar.welcome", {name:user.name})}}
    </a>
    <div class="dropdown-menu" aria-labelledby="personal-bar">
      <router-link :to="{name: 'dashboard'}" class="dropdown-item">{{$t("personal-bar.dashboard")}}</router-link>
      <div class="dropdown-divider"></div>
      <a v-on:click="onSignOut" class="dropdown-item">{{ $t("personal-bar.sign-out") }}</a>
    </div>
  </li>
  <li v-else class="nav-item dropdown">
    <a class="nav-link dropdown-toggle" id="personal-bar" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
      {{$t("personal-bar.sign-in-or-up")}}
    </a>
    <div class="dropdown-menu" aria-labelledby="personal-bar">
      <router-link :key="a" :to="{name: `auth.users.${a}`}" class="dropdown-item" v-for="a in actions">
        {{$t(`auth.users.${a}.title`)}}
      </router-link>
      <router-link :to="{name: 'site.leave-words.new'}" class="dropdown-item">{{$t("site.leave-words.new.title")}}</router-link>
    </div>
  </li>
</template>

<script>
import {_delete} from '@/ajax'

export default {
  name: 'personal-bar',
  data () {
    return {
      actions: ['sign-in', 'sign-up', 'forgot-password', 'confirm', 'unlock']
    }
  },
  methods: {
    onSignOut () {
      _delete('/users/sign-out').then(function () {
        this.$store.commit('signOut')
        this.$router.push({ name: 'auth.users.sign-in' })
      }.bind(this))
    }
  },
  computed: {
    user () {
      return this.$store.state.currentUser
    }
  }
}
</script>
