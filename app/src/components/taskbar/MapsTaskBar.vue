<template>
  <div style="flex-grow: 0; display:flex">
    <button v-on:click="emit( 'settings' )" class="button is-radiusless is-small"><a class="mdi mdi-settings"/></button>
    <div>
      <div :class="'dropdown'+(show===true?' is-active':'')">
        <div class="dropdown-trigger">
          <button v-on:click="show = !show" class="button is-radiusless is-small" aria-haspopup="true" aria-controls="dropdown-menu">
            {{node.active}} {{taskNames[node.active]}}
          <!--  <a :class="'mdi mdi-'+active.icon"> Layout</a>-->
          </button>
        </div>
        <div class="dropdown-menu is-radiusless" id="dropdown-menu" role="menu">
          <div class="dropdown-content is-radiusless">
            <div class="dropdown-content is-radiusless">
              <aside class="menu">
                <MenuLevel v-on:trigger="emit" :parents="''" :node="loadMenu"/>
              </aside>
            </div>
            <hr class="dropdown-divider">
            <a v-for="(item, index) in node.nodes" :key="index" v-on:click="emit( 'active', index )" :class="(index===node.active?'is-active ':'')+'dropdown-item'">
              <span v-on:click="emit('viewer', i)" class="icon">
                <i class="mdi mdi-eye-outline"/>
              </span>
              <span v-on:click="emit('map2file', i)" class="icon">
                <i class="mdi mdi-arrow-left"/>
              </span>
              <span v-on:click="emit('download', i)" class="icon">
                <i class="mdi mdi-download"/>
              </span>
              <span v-on:click="emit('save', i)" class="icon">
                <i class="mdi mdi-content-save"/>
              </span>
              <span v-on:click="emit('mismatch', i)" class="icon">
                <i class="mdi mdi-alert-octagram"/>
              </span>
              <span v-on:click="emit('delete', i)" class="icon">
                <i class="mdi mdi-delete"/>
              </span>
              {{index}} {{ taskNames[index] }}
            </a>
            <!--
            mdi mdi-'+item.icon
            v-on:click="active=item"

            <hr class="dropdown-divider">
            <a href="#" class="dropdown-item">With a divider</a>
            -->
          </div>
        </div>
      </div>
    </div>

    <!--<input id="file-input" v-show="false" type="file" v-on:change="uploadFile">-->
  </div>
</template>

<script>
import MenuLevel from './MenuLevel.vue'

var loadMenu = { nodes: [{
  name: 'Add', icon: 'plus-outline', nodes: [
    {
      name: 'Empty',
      icon: 'checkbox-blank-outline',
      nodes: []
    },
    {
      name: 'Reload volume',
      icon: 'refresh',
      nodes: []
    },
    {
      name: 'Load',
      icon: 'upload',
      nodes: [
        {
          name: 'Local',
          icon: 'harddisk',
          title: 'Load collection from local file (tar, tgz, tar.gz, etc.)',
          nodes: []
        },
        {
          name: 'Remote',
          icon: 'web',
          title: 'Load collection from URL (github, sourceforge, bitbucket, tar, tgz, tar.gz, etc.)',
          nodes: []
        }
      ]
    }
  ]
}]};

export default {
  name: 'MapTaskBar',
  components: {
    MenuLevel
  },
  props: {
    node: Object,
  },
  data () {
    return {
      show: false,
      loadMenu: loadMenu
    }
  },
  computed: {
    taskNames () {
      var t = [];
      for ( var x=0; x < this.node.nodes.length ; x++ ) {
        var a = this.node.nodes[x].name;
        t[x] = ( a === "" ) ? '<empty>' : a ;
      }
      return t;
    }
  },
  methods: {
    emit (t, i) {
      this.$emit( 'trigger', t, i );
    },
    menuTrigger (t,i) {
      switch (t) {
        case 'Add,Load,Local':
          document.getElementById("file-input").click();
          break;
        default:
          this.emit( t, i )
      }
    },/*
    uploadFile (e) {
      var file = e.target.files[0];
      if (!file) { return; }
      this.emit( 'uploadFile', file )
    },
    downloadFile (e) {
      console.log(e);
    }*/
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="sass" scoped>
  /* write SASS! */
</style>

<style lang="sass">
  .taskbar
    padding: 2px
  .dropdown-content
    padding-top: 0px
    padding-bottom: 0px
  .dropdown-item
    padding-left: 8px
    padding-right: 8px
    line-height: 1.25
  .dropdown-divider
    margin-top: .25rem
    margin-bottom: 0
</style>
