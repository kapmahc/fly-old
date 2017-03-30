<template>
  <application-layout>
    <slot v-if="user.uid" />
    <alert-dialog v-else action="danger" :message="$t('errors.forbidden')" />
  </application-layout>
</template>

<script>
import Layout from './Application'
import Alert from '@/components/Alert'

export default {
  name: 'dashboard-layout',
  components: {
    'application-layout': Layout,
    'alert-dialog': Alert
  },
  computed: {
    user () {
      return this.$store.state.currentUser
    }
  }
}
</script>
