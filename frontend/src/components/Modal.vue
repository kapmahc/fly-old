<template>
  <div class="modal fade" :id="id" tabindex="-1" role="dialog" :aria-labelledby="`${id}Label`" aria-hidden="true">
    <div :class="`modal-dialog ${size ? 'modal-'+size : ''}`" role="document">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title" :id="`${id}Label`">{{title}}</h5>
          <button type="button" class="close" data-dismiss="modal" aria-label="Close">
            <span aria-hidden="true">&times;</span>
          </button>
        </div>
        <div class="modal-body">
          <slot />
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-secondary" data-dismiss="modal">{{$t("buttons.close")}}</button>
          <button v-if="onSave" type="button" v-on:click="onSave" class="btn btn-primary">{{$t("buttons.save")}}</button>
          <button v-if="onRemove" type="button" v-on:click="onRemove" class="btn btn-danger">{{$t("buttons.remove")}}</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'modal-form',
  props: ['id', 'size', 'title', 'onSave', 'onRemove']
}
</script>
