<template>
  <application-layout>
    <form-panel :title="title" :onSubmit="onSubmit">
      <slot />
    </form-panel>
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
import Form from '@/components/Form'

export default {
  name: 'auth-non-sign-in-layout',
  props: ['title', 'onSubmit'],
  components: {
    'application-layout': Application,
    'form-panel': Form
  },
  data () {
    return {
      actions: ['sign-in', 'sign-up', 'forgot-password', 'confirm', 'unlock']
    }
  }
}
</script>
