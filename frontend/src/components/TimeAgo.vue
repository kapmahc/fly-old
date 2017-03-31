<template>
  <cite>{{format(date)}}</cite>
</template>

<script>
import TimeAgo from 'timeago.js'

export default {
  name: 'time-ago',
  props: ['date'],
  methods: {
    format (dt) {
      return new TimeAgo().format(dt)
    }
  }
}
</script>
