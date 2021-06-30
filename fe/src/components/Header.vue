<template>
  <div class="header">
    <Avatar/>
  </div>
</template>

<script>
import Avatar from '@/components/Avatar.vue'
export default {
  name: 'Header',
    components: {
        Avatar,
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.header{
  background-color: green;
}

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
