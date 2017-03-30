<template>
  <application-layout>
    <h2>{{ title }}</h2>
    <hr/>
    <form v-on:submit.prevent="onSubmit">
      <slot />
      <button type="submit" class="btn btn-primary">{{$t("buttons.submit")}}</button>
      <button type="reset" class="btn btn-secondary">{{$t("buttons.reset")}}</button>
    </form>
    <br/>
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
  </application-layout>
</template>

<script>
import Application from '@/layouts/Application'

export default {
  name: 'auth-non-sign-in-layout',
  props: ['title', 'onSubmit'],
  components: {
    'application-layout': Application
  },
  data () {
    return {
      actions: ['sign-in', 'sign-up', 'forgot-password', 'confirm', 'unlock']
    }
  }
}
</script>
