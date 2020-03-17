<template>
  <div>
    <div v-on:click="show = !show" :class="'dropdown'+(show===true?' is-active':'')">
      <div class="dropdown-trigger">
        <button class="button is-radiusless is-small" aria-haspopup="true" aria-controls="dropdown-menu">
          <a :class="'mdi mdi-'+active.icon"> Layout</a>
        </button>
      </div>
      <div class="dropdown-menu is-radiusless" id="dropdown-menu" role="menu">
        <div class="dropdown-content is-radiusless">
          <a v-for="(item, index) in nodes" :key="index" v-on:click="active=item" :class="(item.name===active.name?'is-active ':'')+'dropdown-item mdi mdi-'+item.icon"> {{ item.name }}</a>
          <!--
          <hr class="dropdown-divider">
          <a href="#" class="dropdown-item">With a divider</a>
          -->
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'DropDownSelect',
  props: {
    nodes: Array,
    parentActive: Object
  },
  data () {
    return {
      active: this.parentActive,
      show: false
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="sass" scoped>
  .dropdown-item
    padding-left: 8px
    padding-right: 8px
    line-height: 1.25
</style>
