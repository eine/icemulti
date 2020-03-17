<template>
  <div>
  <div :class="'modal'+(node.show===true?' is-active':'')">
    <div class="modal-background"></div>
      <div class="modal-card">
        <header class="modal-card-head">
          <p class="modal-card-title">Bitstream viewer [{{ node.name }}]</p>
          <button v-on:click="close()" class="delete" aria-label="close"></button>
        </header>
        <section class="modal-card-body">
          <!-- Content ... -->
          <HexView :raw="binArray"/>
        </section>
<!--
        <footer class="modal-card-foot">
          <button v-on:click="close()" class="button is-success">Save changes</button>
          <button v-on:click="close()" class="button">Cancel</button>
        </footer>
-->
      </div>
    </div>
  </div>
</template>

<script>
import HexView from './HexView.vue'

export default {
  name: 'Viewer',
  components: {
    HexView
  },
  props: {
    node: Object,
  },
  computed: {
    binArray () {
      var str = atob(this.node.bin);
      var buf = new ArrayBuffer(str.length*2);
      var bufView = new Uint8Array(buf);
      for (var i=0; i < str.length; i++) {
        bufView[i] = str.charCodeAt(i);
      }
      return buf;
    }
  },
  methods: {
    close () {
      this.$emit('show',false);
    }
  }
}
</script>

<style lang="sass" scoped>
  .modal-card
    width: 800px
</style>
