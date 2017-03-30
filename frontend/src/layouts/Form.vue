<template>
  <application-layout>
    <h2>{{ title }}</h2>
    <hr/>
    <form v-on:submit.prevent="onSubmit">
      <slot />      
      <button type="submit" class="btn btn-primary">{{$t("buttons.submit")}}</button>
      <button type="reset" class="btn btn-secondary">{{$t("buttons.reset")}}</button>
    </form>
  </application-layout>
</template>

<script>
import Application from './Application'

export default {
  name: 'form-layout',
  props: ['title', 'onSubmit'],
  components: {
    'application-layout': Application
  }
}
</script>
