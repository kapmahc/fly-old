<template>
  <li v-if="user.uid" class="nav-item dropdown">
    <a class="nav-link dropdown-toggle" id="personal-bar" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
      {{$t("personal-bar.welcome")}}
    </a>
    <div class="dropdown-menu" aria-labelledby="personal-bar">

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
      <router-link :to="{name: 'home'}" class="dropdown-item">{{$t("site.leave-words.new.title")}}</router-link>
    </div>
  </li>
</template>

<script>
export default {
  name: 'personal-bar',
  data () {
    return {
      actions: ['sign-in', 'sign-up', 'forgot-password', 'confirm', 'unlock']
    }
  },
  computed: {
    user () {
      return this.$store.state.currentUser
    }
  }
}
</script>
