<template>
  <div style="flex-grow: 1; height: 100%; overflow: auto;">

    <div v-show="numOfFiles==0" class="section" style="width:100%; height:100%;">
      <div class="notification is-centeredContent">
        <p>Empty filename list. Add some to start tinkering!</p>
      </div>
    </div>

    <table v-show="numOfFiles!=0" class="table is-hoverable" style="overflow: auto !important;">
      <thead>
        <tr>
<!--          <th><abbr title="Position">Pos</abbr></th> -->
          <th></th>
          <th v-for="(v, i) in headers" :key="i">{{v}}</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(v, i) in node.files" :key="i">
<!--          <td><a href="" :title="'See '+v.name+' metadata'">metadata</a></td>-->
          <td style="white-space: nowrap;">
            <span v-on:click="emit('select', i)" class="icon">
              <i :class="'mdi mdi-checkbox-'+(node.select[i]!=undefined?'marked':'blank')+'-outline'"/>
            </span>
            <span v-on:click="emit('viewer', i)" class="icon">
              <i class="mdi mdi-eye-outline"/>
            </span>
            <span v-on:click="emit('add2map', i)" class="icon">
              <i class="mdi mdi-arrow-right"/>
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
          </td>
          <td>{{v.name}}</td>
          <td>{{v.size}}</td>
          <td>{{v.modtime}}</td>
          <td>{{v.meta.uid}}</td>
          <td>{{v.meta.compression}}</td>
          <td>{{v.meta.device}}</td>
          <td>{{v.meta.board}}</td>
        </tr>
      </tbody>
      <tfoot>
        <tr>
          <th></th>
          <th v-for="(v, i) in headers" :key="i">{{v}}</th>
        </tr>
      </tfoot>
    </table>

  </div>
</template>

<script>
export default {
  name: 'Main',
  props: {
    node: Object,
  },
  computed: {
    numOfFiles () {
      return Object.keys(this.node.files).length;
    }
  },
  data () {
    return {
      headers: [
        'Filename',
        'Size',
        'ModTime',
        'UID',
        'Compression',
        'Device',
        'Board'
      ]
    }
  },
  methods: {
    emit (t, i) {
      this.$emit( 'trigger', t, i );
    }
  }
}
</script>

<style lang="sass" scoped>
  table
    margin-bottom: 0
  .is-centeredContent
    height: 100%
    display: flex
    align-items: center
    justify-content: center
  .icon
    cursor: pointer
  .icon:hover
    color: #00f
  .table td,
  .table th
    padding: .25rem .5rem
  .table td
    font-family: Inconsolata, monospace
</style>
