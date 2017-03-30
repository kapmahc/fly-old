<template>
  <div>
    <application-header />
    <div class="container">
      <div class="row">
        <div class="col-md-12">
          <slot />
        </div>
      </div>
      <application-footer />
    </div>
  </div>
</template>

<script>
import Header from '@/components/ApplicationHeader'
import Footer from '@/components/Footer'

export default {
  name: 'application-layout',
  components: {
    'application-header': Header,
    'application-footer': Footer
  }
}
</script>
