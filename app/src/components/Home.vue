<template>
<div id="home">
  <div class="spacer"></div>
  <div id="hello">
    <div><img src="../../../assets/vue_logo.png"></div>
    <h1 class="title is-spaced">
      Go(lang) tools to exploit warm/cold boot in iCE40 FPGAs
    </h1>
    <p class="subtitle">
      Built with <a href="https://vuejs.org/" target="_blank">Vue.js</a>, <a href="https://golang.org/" target="_blank">Go(lang)</a> and <i class="heartbeat mdi mdi-heart"/>
    </p>
    <p class="subsubtitle">
      Packed with <a href="https://webpack.js.org/" target="_blank">webpack</a> and <a href="https://www.docker.com/" target="_blank">docker</a>, to put <a href="https://bulma.io" target="_blank">Bulma</a>, <a href="https://materialdesignicons.com/" target="_blank">mdi</a>, <a href="http://wavedrom.com/" target="_blank">WaveDrom</a>, <a href="https://github.com/surmon-china/vue-codemirror" target="_blank">vue-codemirror</a> (<a href="http://codemirror.net/" target="_blank">codemirror</a>)... together
    </p>
    <button v-on:click="$emit( 'home' )" class="button is-radiusless">Start Hacking!</button>
  </div>
  <div class="spacer"></div>
  <hr/>
  <div id="info">
    <div>
      <p>Check out the CLI version!</p>
      <p><a href="doc">Read the Docs!</a></p>
      <p>Contribute!</p>
    </div>
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
  display: flex;
  flex-direction: column;
  height: 100%;
  padding: 10px;
}
#hello {
  flex-grow: 0;
}
#info{
  flex-grow: 0;
  text-align: left;
}
.spacer {
  flex-grow: 1;
}
.title,
.subtitle {
  color: unset;
  margin: 1.5rem 0rem;
}
.subsubtitle {
  font-size: 1rem;
  margin: 1.5rem 0rem;
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
