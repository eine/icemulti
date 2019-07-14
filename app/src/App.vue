<template>
  <div id="app">
    <Home
      v-show="home"
      v-on:home="home=!home"
    />
    <Viewer
      v-on:show="viewer.show=!viewer.show"
      :node="viewer"
    />
    <Dialog
      v-on:trigger="triggerDialog"
      :node="dialog"
    />
    <split-pane
      v-show="!home"
      v-on:resize="resize"
      :min-percent='20'
      :default-percent='30'
      split="vertical"
    >
      <template slot="paneL">
        <div style="display: flex; flex-direction: column; height: 100%;">
          <!-- Main -->
          <TaskBar
            v-on:trigger="triggerTaskBar"
            :isGroup="isGroup"
          />
          <Main
            v-on:trigger="triggerMain"
            :node="main"
          />
       </div>
      </template>
      <template slot="paneR">
        <split-pane split="vertical">
          <template slot="paneL">
            <split-pane split="horizontal">
              <template slot="paneL">
                <!-- CodeGen -->
                <CodeGen
                  :node="gens"
                />
              </template>
              <template slot="paneR">
                <!-- Tabs -->
                xterm/codemirror
              </template>
            </split-pane>
          </template>
          <template slot="paneR">
            <!-- MemMap -->
            <MemMap
              v-on:trigger="triggerMaps"
              :node="maps"
            />
          </template>
        </split-pane>
      </template>
     </split-pane>
  </div>
</template>

<script>
import 'bulma/css/bulma.css';
import '@mdi/font/css/materialdesignicons.min.css';

import splitPane from 'vue-splitpane'

import TaskBar from './components/taskbar/TaskBar.vue'
import Home from './components/Home.vue'
import Main from './components/Main.vue'
import Viewer from './components/Viewer.vue'
import CodeGen from './components/CodeGen.vue'
import MemMap from './components/MemMap.vue'
import Dialog from './components/Dialog.vue'

var t_main = {
  select: {},
  files: {}
};

var t_viewer = {
  show: false,
  name: "",
  bin: ""
};

var t_dialog = {
  show: false,
  title: "",
  content: "",
  func: "",
  args: ""
};

var t_gens = {
  active: "",
  nodes: [
    {
      memmap: "",
      setup: "stw-uc",
      platform: "arduino",
      function: "test"
    }
  ]
};

var t_map = {
  name: "",
  size: "",
  erasesize: "",
  buffersize: "",
  interface: "",
  files: []
}

var t_maps = {
  active: 0,
  nodes: [ t_map ]
};

export default {
  name: 'app',
  components: {
    splitPane,
    Home,
    Dialog,
    Main,
    Viewer,
    CodeGen,
    MemMap,
    TaskBar
  },
  data () {
    return {
      windowWidth: 0,
      home: false,
      cnt: 0,
      dialog: Object.assign({}, t_dialog),
      viewer: Object.assign({}, t_viewer),
      main: Object.assign({}, t_main),
      gens: Object.assign({}, t_gens),
      maps: Object.assign({}, t_maps)
    }
  },
  computed: {
    selArr() {
      return Object.keys(this.main.select);
    },
    isGroup () {
      return this.selArr.length>1;
    }
  },
  methods: {
    triggerTaskBar (t,i) {
      switch(t) {
        case 'home':
          this.home = !this.home;
          break;
        case 'loadList':
          this.loadList(i);
          break;
        case 'uploadFiles':
          this.loadFiles(i);
          break;
        case 'add2map':
          this.addSel2map();
          break;
        default:
          console.log(t,i);
      }
    },
    triggerDialog (t,i) {
      switch(t) {
        case 'close':
          break;
        case 'overwrite':
          this[this.dialog.func](this.dialog.args);
          break;
        default:
          console.log(t,i);
      }
      this.$set(this.dialog, 'show', false);
      this.cnt += 1;
    },
    triggerMain (t,i) {
      switch(t) {
      case 'select':
        var s = this.main.select;
        (s[i]===undefined) ? this.$set( s, i, true ) : this.$delete( s, i );
        break;
      case 'viewer':
        this.viewer.name = i;
        this.viewer.bin = this.main.files[i].bin;
        this.viewer.show = true;
        break;
      case 'delete':
        this.$delete( this.main.files, i);
        break;
      case 'add2map':
        this.add2map(i);
        break;
      default:
        console.log(t,i);
      }
    },
    addSel2map() {
      this.cnt = 0;
      var k = 1;
      this.add2map( this.selArr[0] );
      var timer = setInterval( function() {
        if (this.cnt === k) {
          if ( k >= this.selArr.length ) {
            clearInterval(timer);
            return;
          }
          this.add2map( this.selArr[k] );
          k++;
        }
      }.bind(this), 250);
    },
    add2map(i) {
      this.isPack( this.main.files[i].bin ) ?
        this.pack2mapDialog(i)
      :
        this.img2map(i)
      ;
    },
    img2map (i) {
      this.cnt += 1;
      this.maps.nodes[this.maps.active].files.push({
        name: i,
        size: this.main.files[i].size
      });
    },
    isPack(bin) {
      var b = atob(bin).slice(0,4);
      //7E AA 99 7E
      return (b.charCodeAt(0)==126) && (b.charCodeAt(1)==170) &&
             (b.charCodeAt(2)==153) && (b.charCodeAt(3)==126);
    },
    pack2mapDialog (i) {
      console.log("Packed file");
      this.dialog.title = "Create memory map from pack file <"+i+">";
      this.dialog.func = "pack2map";
      this.dialog.content = "THIS OPTION IS NOT SUPPORTED, YET. NONE OF THE BUTTONS WILL HAVE ANY EFFECT. This seems to be a packet file containing an applet and multiple bitstreams. A memory map will be created from it. Do you want to create a NEW memory map or do you want to OVERWRITE the active one?";
      this.dialog.args = this.main.files[i];
      this.$set(this.dialog, 'show', true);
    },
    pack2map (a) {
      console.log(a);
    },
    triggerMaps (t,i) {
      switch(t) {
        case 'active':
          this.maps.active = i;
          break;
        case 'Add,Empty':
          var m = Object.assign({}, t_map);
          m.files = [];
          this.maps.nodes.push(m);
          break;
        default:
          console.log(t,i);
      }
    },
    loadList (d) {
      this.$set(this.main, "select", {});
      this.$set(this.main, "files", d);
    },
    loadFiles (d) {
      for ( var x=0; x < d.length; x++ ) {
        var f = d[x];
        switch (f.type) {
          case "application/octet-stream":
            var reader = new FileReader();
            reader.binFile = {
              path: "",
              name: f.name,
              size: f.size,
              modtime: f.lastModified,
              meta: {},
              bin: ""
            };
            reader.onload = function(e) {
              var binFile = e.target.binFile;
              var b = e.target.result;
              binFile.bin = b.slice( b.indexOf("base64,")+7 );
              this.$set(this.main.files, binFile.name, binFile);
            }.bind(this);
            reader.readAsDataURL(f);
            break;
          default:
            console.log("Unknown file type");
        }
      }
    }
  }
}
</script>

<style lang="sass">
  @import ./sass/colors.sass
  html
    overflow-y: auto !important
  body
  #app
    position: absolute
    min-width: 100%
    height: 100%
    color: $icemulti-front
    background-color: $icemulti-back
  a
    color: $icemulti-front
  .button
    background-color: rgba($icemulti-front,.75)
    border-color: $icemulti-back
  .button a
    color: $icemulti-back
</style>
