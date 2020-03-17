<template>
<div>

  <div class="field is-grouped is-horizontal">
    <div class="field-label is-small">
      <label class="label">Name</label>
    </div>
    <div class="field-body is-small">
      <div class="field">
        <div class="control">
          <input class="input is-small" type="text" :placeholder="node.name">
        </div>
      </div>
    </div>
  </div>

  <div class="field is-grouped is-horizontal" v-for="(item, index) in sizes" :key="index">
    <div class="field-label is-small">
      <label class="label">{{index}} Size</label>
    </div>
    <div class="field-body is-small">
      <div class="field">
        <div class="control is-expanded">
          <input class="input is-small" type="text" v-model="node.size">
        </div>
      </div>
      <div class="field">
        <div class="control">
          <div class="select is-small">
            <select>
              <option></option>
              <option v-on:click="item.prefix='K'">(K)ilo</option>
              <option v-on:click="item.prefix='M'">(M)ega</option>
              <option v-on:click="item.prefix='G'">(G)iga</option>
            </select>
          </div>
        </div>
      </div>
      <div class="field">
        <div class="control">
          <div class="select is-small">
            <select>
              <option v-on:click="item.unit='b'">(b)its</option>
              <option v-on:click="item.unit='B'">(B)ytes</option>
            </select>
          </div>
        </div>
      </div>
    </div>
  </div>

  <div class="field is-grouped is-horizontal">
    <div class="field-label is-small">
      <label class="label">Interface</label>
    </div>
    <div class="field-body is-small">
      <div class="field">
        <div class="control">
          <div class="select is-small">
            <select>
              <option>SPI</option>
              <option>QUAD SPI</option>
            </select>
          </div>
        </div>
      </div>
    </div>
  </div>

  {{ node.interface }}

</div>
</template>

<script>
export default {
  name: 'MemMapSettings',
  props: {
    node: Object,
  },
  data () {
    return {
      sizes: {
        "": {
          prefix: '',
          unit: 'b'
        },
        "Erase ": {
          prefix: '',
          unit: 'b'
        },
        "Buffer ": {
          prefix: '',
          unit: 'b'
        }
      },
    }
  }
}
</script>

<style lang="sass" scoped>
  .label
    color: #96b2bf
</style>
