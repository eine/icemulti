<template>
  <div>
    <GensTaskBar/>
    <button class="button is-radiusless is-small">New CodeGen</button>
    <div v-for="(item, index) in node.nodes" :key="index">
      {{ item.memmap }}
      {{ item.setup }}
      {{ item.platform }}
      {{ item.function }}
    </div>
  </div>
</template>

<script>
/*
stw-uc
stw-euc
stw-hdl
stw-usb

uc
euc
hdl
usb
*/

import GensTaskBar from './taskbar/GensTaskBar.vue'

export default {
  name: 'CodeGen',
  components: [
    GensTaskBar
  ],
  props: {
    node: Object,
  },
}
</script>

<style scoped>
</style>
