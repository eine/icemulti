<template>
  <div>
  <div :class="'modal'+(node.show===true?' is-active':'')">
    <div class="modal-background"></div>
      <div class="modal-card">
        <header class="modal-card-head">
          <p class="modal-card-title">{{ node.title }}</p>
          <button v-on:click="emit('close')" class="delete" aria-label="close"></button>
        </header>
        <section class="modal-card-body">
          <!-- Content ... -->
          {{ node.content }}
        </section>
        <footer class="modal-card-foot">
          <button v-on:click="emit('new')" class="button is-primary">New</button>
          <button v-on:click="emit('overwrite')" class="button is-warning">Overwrite</button>
          <button v-on:click="emit('close')" class="button">Cancel</button>
        </footer>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Dialog',
  props: {
    node: Object,
  },
  methods: {
    emit (t,i) {
      this.$emit('trigger',t,i);
    }
  }
}
</script>

<style lang="sass" scoped>
</style>
