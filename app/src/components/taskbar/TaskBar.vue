<template>
  <div style="flex-grow: 0; display:flex; padding: .25rem;">
    <button v-on:click="emit( 'home' )" class="button is-radiusless is-small"><a class="mdi mdi-home"/></button>
    <DropDownMenu v-on:trigger="loadTrigger" :node="loadMenu"/>
    <button
      v-show="isGroup"
      v-for="(b, i) in groupBtns"
      :key="i"
      :title="b.title"
      class="button is-radiusless is-small"
      v-on:click="emit( b.emit )"
    ><a :class="'mdi mdi-'+b.icon"/></button>

    <input id="file-input" multiple v-show="false" type="file" v-on:change="uploadFiles">
  </div>
</template>

<script>
import DropDownMenu from './DropDownMenu.vue'

var groupBtns = [
  {
    title: 'Add to map',
    emit: 'add2map',
    icon: 'arrow-right'
  },
  {
    title: 'Download',
    emit: 'download',
    icon: 'download'
  },
  {
    title: 'Delete',
    emit: 'delete',
    icon: 'delete'
  }
];

var loadMenu = {
  name: 'Add',
  icon: 'plus-outline',
  title: 'Add bitstreams, packs and/or collections',
  nodes: [
    {
      name: 'Empty',
      icon: 'checkbox-blank-outline',
      title: 'Add empty file entry',
      nodes: []
    },
    {
      name: 'Reload volume',
      icon: 'refresh',
      title: 'Get content from data volume',
      nodes: []
    },
    {
      name: 'Load',
      icon: 'upload',
      title: 'Load bitstreams, packs and/or collections',
      nodes: [
        {
          name: 'Local',
          icon: 'harddisk',
          title: 'Load from local file (bin, icem, tar, tgz, tar.gz, etc.)',
          nodes: []
        },
        {
          name: 'Remote',
          icon: 'web',
          title: 'Load from URL (github, sourceforge, bitbucket, tar, tgz, tar.gz, etc.)',
          nodes: []
        }
      ]
    }
  ]
};

export default {
  name: 'TaskBar',
  components: {
    DropDownMenu,
  },
  props: {
    isGroup: Boolean,
  },
  data () {
    return {
      loadMenu: loadMenu,
      groupBtns: groupBtns,
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
          this.emit('loadList', r.body);
        } else {
          result='Request failed. Returned status of ' + r.status;
        }
        console.log(result);
      })
    },
    loadTrigger (e) {
      switch (e) {
        case 'Add,Load,Local':
          document.getElementById("file-input").click();
        break;
        case 'Add,Reload volume':
          this.ajax_post();
        break;
        default:
          this.emit( 'trigger' , e )
      }
    },
    uploadFiles (e) {
      this.emit( 'uploadFiles', e.target.files );
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
