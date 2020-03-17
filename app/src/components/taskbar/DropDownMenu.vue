<template>
  <div
    :class="'dropdown'+(show===true?' is-active':'')"
  >
    <div
      v-on:click="show = !show"
      class="dropdown-trigger"
    >
      <button
        class="button is-radiusless is-small"
        :title="node.title!=undefined ? node.title : ''"
        aria-haspopup="true"
        aria-controls="dropdown-menu"
      >
        <a :class="'mdi mdi-'+node.icon"> {{ node.name }}</a>
      </button>
    </div>
    <div
      class="dropdown-menu is-radiusless"
      id="dropdown-menu"
      role="menu"
    >
      <div class="dropdown-content is-radiusless">
        <aside class="menu">
          <MenuLevel
            v-on:trigger="trigger"
            :parents="''"
            :node="node"
          />
        </aside>
      </div>
    </div>
  </div>
</template>

<script>
import MenuLevel from './MenuLevel.vue'
export default {
  name: 'DropDownMenu',
  components: {
    MenuLevel
  },
  props: {
    node: Object,
  },
  data () {
    return {
      show: false
    }
  },
  methods: {
    trigger: function(n) {
      this.$emit( 'trigger', n )
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="sass" scoped>
</style>
