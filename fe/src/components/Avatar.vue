<template>
  <div class="avatar-container">
    <el-avatar :size="size" :src="circleUrl"></el-avatar>
  </div>
</template>
<script>
export default {
  name: 'Avatar',
  data () {
      return {
        circleUrl: "https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png",
        size: "medium"
      }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.avatar-container{
    display: flex;
    flex-direction:row-reverse;
}
</style>
