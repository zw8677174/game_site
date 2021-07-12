<template>
  <div class="hello">
    <Header/>
    <Game/>
    <Footer/>
  </div>
</template>

<script>
import Header from '@/components/Header.vue'
import Footer from '@/components/Footer.vue'
import Game from '@/components/Game.vue'
import api from '@/api/main.js'


export default {
    name: 'Home',
  mounted() {
  },
    components: {
        Header,
        Game,
        Footer
    }
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
</style>
