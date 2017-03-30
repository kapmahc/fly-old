<template>
  <div :class="`alert alert-${action} alert-dismissible fade show`" role="alert">
    <button type="button" class="close" data-dismiss="alert" aria-label="Close">
      <span aria-hidden="true">&times;</span>
    </button>
    <h4 class="alert-heading">{{$t(`flashs.${action}`)}}</h4>
    <p>{{new Date()}}</p>
    <p class="mb-0">{{message}}</p>
  </div>
</template>

<script>
export default {
  name: 'alert-dialog',
  props: ['action', 'message']
}
</script>
