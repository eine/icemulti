<template>
  <div class="hexview">
    <table class="table is-hoverable">
      <!--
      <thead>
        <tr>
          <th v-for="(v, i) in headers" :key="i">{{v}}</th>
        </tr>
      </thead>
      -->
      <tbody>
        <tr v-for="(v, i) in content" :key="i">
          <td v-for="(a, i) in v.adr" :key="i">{{a}}</td>
          <td v-for="(n, i) in v.val" :key="i">{{n}}</td>
        </tr>
      </tbody>
    </table>

    <!--<pre class="hexdump">{{ content }}</pre>-->
  </div>
</template>

<script>
//https://github.com/chaitin/passionfruit/blob/master/gui/src/components/HexView.vue
//https://brendanzagaeski.appspot.com/0006.html
//https://en.wikipedia.org/wiki/Hex_dump
export default {
  props: {
    raw: ArrayBuffer,
  },
  computed: {
    content() {
      return (this.raw && this.raw.byteLength) ? this.hexmat : {
        0: { adr: [ '0000','FFFF' ], val: [0,1] },
        1: { adr: [ '0000','FFFF' ], val: [2,3] }
      };
    },
    hexdump() {
      let dump = '      0  1  2  3  4  5  6  7  8  9  A  B  C  D  E  F 0123456789ABCDEF'
      let view = new DataView(this.raw)
      for (let i = 0; i < this.raw.byteLength; i += 16) {
        dump += `\n${('0000' + i.toString(16).toUpperCase()).slice(-4)} `
        for (let j = 0; j < 16; j++) {
          let ch = i + j > this.raw.byteLength - 1 ?
            '  ' :
            (0 + view.getUint8(i + j).toString(16).toUpperCase()).slice(-2)
          dump += `${ch} `
        }
        dump += String.fromCharCode.apply(null,
          new Uint8Array(this.raw.slice(i, i + 16)))
          .replace(/[^\x20-\x7E]/g, '.')
      }
      return dump
    },
    hexmat() {
//      let dump = '      0  1  2  3  4  5  6  7  8  9  A  B  C  D  E  F 0123456789ABCDEF'
      let view = new DataView(this.raw)
      var mat = {}
      for (let i = 0; i < this.raw.byteLength; i += 16) {
        var m = [];
        var z = true;
        for (let j = 0; j < 16; j++) {
          m[j] = (i + j > this.raw.byteLength - 1) ?
            '  ' :
            (0 + view.getUint8(i + j).toString(16).toUpperCase()).slice(-2);
          if (z && m[j]!='00') {
            z = false;
          }
        }
        if (z===false) {
          var adr = ('00000000' + i.toString(16).toUpperCase()).slice(-8);
          mat[i] = {
            adr: [ adr.slice(0,4), adr.slice(4,8) ],
            val: m
          };
        }
      }
      return mat
    }
  }
}
</script>

<style lang="sass" scoped>
  .table td
    padding: .25rem .25rem
    font-family: Inconsolata, monospace
</style>
