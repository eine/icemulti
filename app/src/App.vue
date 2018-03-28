<template>
  <div id="app">
    <div style="display: flex;">
      <TaskBar v-on:trigger="taskbarEvent" :tasks="tasks" :activeTask="activeTask"/>
      <Tabs style="flex-grow: 1;" v-on:tab="tabEvent" :tasks="tasks" :active="activeTask"/>
    </div>
    <div class="layout" style="display:flex">
      <Workspace class="workspace" v-on:close="closeWorkspace" v-for="(item, index) in workspaces" :key="index" :id="index" :obj="item" v-show="index===activeTask"/>
    </div>
    <Settings v-on:show="showChange" :show="show"/>
    <Home v-show="activeTask===1000"/>
  </div>
</template>

<script>
import 'bulma/css/bulma.css';
import 'mdi/css/materialdesignicons.min.css';

import Home from './components/Home.vue'
import Settings from './components/Settings.vue'
import TaskBar from './components/taskbar/TaskBar.vue'
import Tabs from './components/layouts/Tabs.vue'
import Workspace from './components/workspaces/Workspace.vue'

export default {
  name: 'app',
  components: {
    Home,
    Settings,
    TaskBar,
    Tabs,
    Workspace
  },
  data () {
    return {
      workspaces: [],
      show: false,
      activeTask: 1000
    }
  },
  computed: {
    tasks: function () {
      var t = [];
      for (var x=0; x<this.workspaces.length ; x++) {
        var o = {}
        o.title = this.workspaces[x].title;
        o.type  = this.workspaces[x].type;
        t.push(o)
      }
      return t;
    }
  },
  methods: {
    showChange (v) {
      this.show = v;
    },
    tabEvent (v) {
      this.activeTask = v;
    },
    taskbarEvent (e,v) {
      switch (e) {
        case 'active':
          this.tabEvent(v);
        break;
        case 'home':
          this.tabEvent(1000);
        break;
        case 'settings':
          this.show = true;
        break;
        case 'trigger':
          this.trigger(v);
        break;
        case 'uploadFile':
          this.uploadFile(v);
        break;
      }
    },
    uploadFile (f) {
      var reader = new FileReader();
      reader.onload = function(e) {
        this.newWorkspace( f.name, 'wavedrom', {
          file: f,
          text: e.target.result
        });
      }.bind(this);
      reader.readAsText(f);
    },
    newWorkspace (title, type, data) {
      this.workspaces.push({
        'title': title,
        'type': type,
        'data': data
      });
      this.activeTask = this.workspaces.length-1;
    },
    closeWorkspace (id) {
      this.workspaces.splice( id, 1 )
    },
    trigger (d) {
      console.log('App trigger: ', d);
      if ( d === 'Add workspace,Empty' ) {
        this.newWorkspace('Empty', 'empty', '');
      }
      if ( d === 'Add workspace,Editor' ) {
        this.newWorkspace('Editor', 'editor', '');
      }
      if ( d === 'Add workspace,Xterm/log' ) {
        this.newWorkspace('Xterm/log', 'term', '');
      }
      if ( d === 'Add workspace,Graph/Schematic' ) {
        this.newWorkspace('Graph', 'graph', '');
      }
      if ( d === 'Add workspace,Tree' ) {
        this.newWorkspace('Tree', 'tree', '');
      }
    }
  }
}
</script>

<style lang="sass">
  @import ./sass/colors.sass
  html
    overflow-y: auto
  a
    color: $dtd-front
  .button
    background-color: rgba($dtd-front,.75)
    border-color: $dtd-back
  .button a
    color: $dtd-back
  .workspace
    border: 1px solid $dtd-back
    width: 100%
    margin: 2px
  .workspace:not(:last-child)
    margin-bottom: 2px
</style>
