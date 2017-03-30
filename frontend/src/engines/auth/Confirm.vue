<template>
  <auth-email-form action="confirm" />
</template>

<script>
import EmailForm from './EmailForm'

export default {
  name: 'auth-users-confirm',
  components: {
    'auth-email-form': EmailForm
  }
}
</script>
