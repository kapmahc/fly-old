<template>
  <p v-html="marked(body)" />
</template>

<script>
import marked from 'marked'

export default {
  name: 'markdown-panel',
  props: ['body'],
  methods: {
    marked
  }
}
</script>
