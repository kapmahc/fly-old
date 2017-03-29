<template>
  <div>
    <h2>{{$t("auth.users.sign-in.title")}}</h2>
    <hr/>
    <br/>
    <shared-links />
  </div>
</template>

<script>
import Links from './Links'

export default {
  name: 'auth-sign-in',
  components: {
    'shared-links': Links
  }
}
</script>
