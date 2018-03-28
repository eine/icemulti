<template>
  <div id="home">
    <section class="section">
  <div class="container">
    <h1 class="title">
      Yet another Digital Timing Diagram tool
    </h1>
    <p class="subtitle">
      Built with <a href="https://vuejs.org/" target="_blank">Vue.js</a>, <a href="https://golang.org/" target="_blank">Go(lang)</a> and <i class="heartbeat mdi mdi-heart"/>
    </p>
    <p class="subtitle">
      Integrating <a href="https://bulma.io" target="_blank">Bulma</a>, <a href="https://materialdesignicons.com/" target="_blank">mdi</a>, <a href="http://wavedrom.com/" target="_blank">WaveDrom</a>, <a href="https://github.com/surmon-china/vue-codemirror" target="_blank">vue-codemirror</a> (<a href="http://codemirror.net/" target="_blank">codemirror</a>)...
    </p>
    <p class="subtitle">
      Packed with <a href="https://webpack.js.org/" target="_blank">webpack</a> and <a href="https://www.docker.com/" target="_blank">docker</a>
    </p>
  </div>
  </section>
  <div class="hello">
    <img src="../../../assets/vue_logo.png">
    <p>
      <a href="https://github.com/vuejs/vue-cli/tree/dev/docs" target="_blank">vue-cli documentation</a>
      ;
      <b>Installed CLI Plugins</b>:
        <a href="https://github.com/vuejs/vue-cli/tree/dev/packages/%40vue/cli-plugin-babel" target="_blank">babel</a>,
        <a href="https://github.com/vuejs/vue-cli/tree/dev/packages/%40vue/cli-plugin-eslint" target="_blank">eslint</a>
      ;
      <b>Essential Links</b>:
        <a href="https://vuejs.org" target="_blank">Core Docs</a>,
        <a href="https://forum.vuejs.org" target="_blank">Forum</a>,
        <a href="https://chat.vuejs.org" target="_blank">Community Chat</a>,
        <a href="https://twitter.com/vuejs" target="_blank">Twitter</a>
      ;
      <b>Ecosystem</b>:
        <a href="https://router.vuejs.org/en/essentials/getting-started.html" target="_blank">vue-router</a>,
        <a href="https://vuex.vuejs.org/en/intro.html" target="_blank">vuex</a>,
        <a href="https://github.com/vuejs/vue-devtools#vue-devtools" target="_blank">vue-devtools</a>,
        <a href="https://vue-loader.vuejs.org/en" target="_blank">vue-loader</a>,
        <a href="https://github.com/vuejs/awesome-vue" target="_blank">awesome-vue</a>
    </p>
  </div>
  </div>
</template>

<script>
export default {
  name: 'Home'
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
.subtitle:not(:last-child) {
  margin-bottom: .5rem;
}
.title:not(.is-spaced) + .subtitle {
    margin-top: -1rem;
}
#home {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
.heartbeat {
  color: #ae2323;
  display: inline-block;
  position: relative;
  -webkit-animation: heartbeat 1s infinite linear;
  animation: heartbeat 1s infinite linear;
}
@keyframes heartbeat {
  0%   { transform: scale(  .85 ); }
  20%  { transform: scale( 1.1  ); }
  40%  { transform: scale(  .85 ); }
  60%  { transform: scale( 1.35 ); }
  80%  { transform: scale(  .85 ); }
  100% { transform: scale(  .85 ); }
}
</style>
