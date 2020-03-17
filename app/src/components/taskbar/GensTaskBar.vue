<template>
  <div style="flex-grow: 0; display:flex">
    <button class="button is-radiusless is-small"><a class="mdi mdi-plus-outline"/></button>
    <DropDownMenu v-on:trigger="menuTrigger" :node="loadMenu"/>
<!--    <DropDownSelect :nodes="layouts" :parentActive="activeLayout"/> -->

    <input id="file-input" v-show="false" type="file" v-on:change="uploadFile">
  </div>
</template>

<script>
import DropDownMenu from './DropDownMenu.vue'

var layouts = [
  { name: 'Microcontroller', icon: 'chip', nodes: [
    { name: 'Arduino', icon: '', nodes:[] },
    { name: 'CortexM', icon: '', nodes:[] },
  ]},
  { name: 'Embedded Microcontroller', icon: 'view-dashboard',  nodes: [
    { name: 'Lattuino', icon: '', nodes:[] },
    { name: 'VexRiscv', icon: '', nodes:[] }
  ]},
  { name: 'HDL', icon: 'code-braces', nodes:[] },
  { name: 'USB', icon: 'usb', nodes: [
    { name: 'FTDI', icon: '', nodes:[] },
    { name: 'Bootloader', icon: '', nodes:[] },
  ]}
];

var loadMenu = {
  name: 'Load', icon: 'upload', nodes: [
    {
      name: 'Local',
      icon: 'harddisk',
      title: 'Load collection from local file (tar, tgz, tar.gz, etc.)',
      nodes: []
    },
    { name: 'Remote',
      icon: 'web',
      title: 'Load collection from URL (github, sourceforge, bitbucket, tar, tgz, tar.gz, etc.)',
      nodes: []
    }
  ]
};

export default {
  name: 'GensTaskBar',
  components: {
    DropDownMenu
  },
  props: {
    tasks: Array,
    activeTask: Number,
  },
  data () {
    return {
      layouts: layouts,
      loadMenu: loadMenu,
      activeLayout: layouts[0],
    }
  },
  methods: {
    emit (t, i) {
      this.$emit( 'trigger', t, i );
    },
    ajax_post () {
      this.$http.post('/ajax/list/',"all").then((r) => {
        var result = "";
        if (r.status === 200) {
          result='Response: ' + r.body;
          this.$emit('files', r.body);
        } else {
          result='Request failed. Returned status of ' + r.status;
        }
        console.log(result);
      })
    },
    menuTrigger (e) {
      if ( e === 'Add workspace,Diagram,Load' ) {
        document.getElementById("file-input").click();
      } else {
        this.emit( 'trigger' , e )
      }
    },
    taskTrigger (e,v) {
      this.emit( e , v )
    },
    uploadFile (e) {
      var file = e.target.files[0];
      if (!file) { return; }
      this.emit( 'uploadFile' , file )
    },
    downloadFile (e) {
      console.log(e);
    }
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
</style>
