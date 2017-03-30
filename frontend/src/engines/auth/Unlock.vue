<template>
  <auth-email-form action="unlock" />
</template>

<script>
import EmailForm from './EmailForm'

export default {
  name: 'auth-users-unlock',
  components: {
    'auth-email-form': EmailForm
  }
}
</script>
