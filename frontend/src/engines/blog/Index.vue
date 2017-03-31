<template>
  <application-layout>
    <h2>{{$t('blog.index.title')}}</h2>
    <hr/>
    <blog-list />
  </application-layout>
</template>

<script>
import Application from '@/layouts/Application'
import List from './List'

export default {
  name: 'blog-index',
  components: {
    'application-layout': Application,
    'blog-list': List
  }
}
</script>
