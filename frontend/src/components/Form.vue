<template>
  <form v-on:submit.prevent="onSubmit">
    <h2>{{ title }}</h2>
    <hr/>
    <slot />
    <button type="submit" class="btn btn-primary">{{$t("buttons.submit")}}</button>    
  </form>
</template>

<script>
export default {
  name: 'form-panel',
  props: ['title', 'onSubmit']
}
</script>
