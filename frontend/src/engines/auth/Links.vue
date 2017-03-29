<template>
  <ul>
    <li :key="a" v-for="a in actions">
      <router-link :to="{name: `auth.users.${a}`}">
        {{$t(`auth.users.${a}.title`)}}
      </router-link>
    </li>
    <li>
      <router-link :to="{name: 'home'}">{{$t("site.leave-words.new.title")}}</router-link>
    </li>
  </ul>
</template>

<script>
export default {
  name: 'auth-shared-links',
  data () {
    return {
      actions: ['sign-in', 'sign-up', 'forgot-password', 'confirm', 'unlock']
    }
  }
}
</script>
