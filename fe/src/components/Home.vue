<template>
  <div class="hello">
    <Header/>
    <el-button>home</el-button>
    <Footer/>
  </div>
</template>

<script>
import Header from '@/components/Header.vue'
import Footer from '@/components/Footer.vue'
import api from '@/api/main.js'


export default {
    name: 'Home',
  mounted() {
      let ret = api.getGameList().then(v => {

      })
      console.log(ret);
  },
    components: {
        Header,
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
