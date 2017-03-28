<template>
  <div id="app">
    <application-header></application-header>
    <div class="container">
        <router-view></router-view>
    </div>
    <application-footer></application-footer>
  </div>
</template>

<script>
import Header from '@/components/Header'
import Footer from '@/components/Footer'

export default {
  name: 'app',
  components: {
    'application-header': Header,
    'application-footer': Footer
  }
}
</script>
