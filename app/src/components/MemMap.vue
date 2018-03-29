<template>
<div style="height: 100%;
display: flex;
flex-direction: column;">
  <MapsTaskBar :node="node" v-on:trigger="triggerTaskBar"/>
  <div v-show="settings" style="flex-grow: 0;">
    <hr/>
    <MemMapSettings :node="activeNode"/>
    <hr/>
  </div>

  <div v-show="numOfFiles==0" class="section" style="width:100%; height:100%;">
    <div class="notification is-centeredContent">
      <p>Empty filename list. Add some to start tinkering!</p>
    </div>
  </div>

  <div v-show="numOfFiles!=0" style="display: flex; flex-direction: column; overflow: auto; height: 100%;">

    <div style="padding: .25rem;">
      <progress class="progress is-primary" value="15" max="100">30%</progress>
      <hr/>
    </div>

    <div style="display: flex; overflow: auto; height: 100%;">
      <div style="flex-grow: 1;">
      <table class="table is-hoverable">
        <thead>
          <tr>
            <th v-for="(v, i) in headers" :key="i">{{v}}</th>
          </tr>
        </thead>
        <tfoot>
          <tr>
            <th v-for="(v, i) in headers" :key="i">{{v}}</th>
          </tr>
        </tfoot>
        <tbody>
          <tr v-for="(item, index) in activeNode.files" :key="index">
            <td>{{item.name}}</td>
            <td title="Set base address"><input type="text" v-model="strOffset[index]" class="offsetInput"></td>
            <td
              v-if="activeNode.hexends[index]!=''"
              :title="'Decimal: '+activeNode.ends[index]"
              style="white-space: nowrap;"
            >0x {{ activeNode.hexends[index].slice(0,4) }} {{ activeNode.hexends[index].slice(4,8) }}</td>
            <td
              v-if="activeNode.hexends[index]===''"
              title="NaN"
            ></td>
          </tr>
        </tbody>
      </table>
      </div>

      <div style="flex-grow: 0;">
        <svg id="mapsvg" width="100%" height="95%"></svg>
      </div>
    </div>
  </div>
</div>
</template>

<script>
import MapsTaskBar from './taskbar/MapsTaskBar.vue'
import MemMapSettings from './MemMapSettings.vue'
//import d3 from 'd3'
const d3 = require('d3');

export default {
  name: 'MemMap',
  components: {
    MapsTaskBar,
    MemMapSettings
  },
  props: {
    node: Object,
  },
  data () {
    return {
      settings: true,
      strOffset: [],
      offsets: [],
      headers: [
        'Filename',
        'Offset',
        'End'
      ]
    }
  },
  computed: {
    numOfFiles () {
      return this.activeNode.files.length;
    },
    numOffsets () {
      var o = [];
      for ( var x=0; x < this.strOffset.length; x++ ) {
        o[x] = parseInt(this.strOffset[x].trim().replace(/\s/g, ""));
      }
      return o;
    },
    activeNode () {
      var n = this.node;
      var a = n.nodes[this.node.active];
      a.ends = [];
      a.hexends = [];
      for ( var x=0; x < a.files.length; x++ ) {
        var o = this.numOffsets[x]
        if ( (o === undefined) || isNaN(o) ) {
          a.ends[x] = NaN;
          a.hexends[x] = '';
        } else {
          a.ends[x] = Number(o) + a.files[x].size;
          a.hexends[x] = ('00000000' + a.ends[x].toString(16).toUpperCase()).slice(-8);
        }
      }
      return a;
    }
  },
  /*
  watch: {
    node () {
      console.log('KAIXO')
    }
  },*/
  methods: {
    emit (t, i) {
      this.$emit( 'trigger', t, i );
    },
    triggerTaskBar(t,i) {
      switch(t) {
        case 'settings':
          this.settings = !this.settings;
          break;
        default:
          this.emit(t,i);
      }
    }
  },
  mounted: function() {
  d3.select(document.getElementById("mapsvg"))
    .append('rect')
    .attr("x", 0)
    .attr("y", 0)
    .attr("width", 100)
    .attr("height", 20)
    .attr("fill","red");
  }
}
</script>

<style lang="sass" scoped>
  .table td
    padding: .25rem .5rem
    font-family: Inconsolata, monospace
  .offsetInput
    font-family: Inconsolata, monospace
    text-align: right
    font-size: 1rem
    width: 7.75rem
</style>
