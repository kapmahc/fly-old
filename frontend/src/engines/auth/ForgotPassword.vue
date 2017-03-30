<template>
  <auth-email-form action="forgot-password" />
</template>

<script>
import EmailForm from './EmailForm'

export default {
  name: 'auth-users-forgot-password',
  components: {
    'auth-email-form': EmailForm
  }
}
</script>
