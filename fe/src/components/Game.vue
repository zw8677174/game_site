<template>
  <div class="gameList">
      <el-row>
        <el-col v-for="(item) in gameList" :span="8">{{item}}</el-col>
      </el-row>
  </div>
</template>
<script>
import api from "@/api/main";

export default {
  name: 'Game',
  data () {
    return {
      gameList:[]
    };
  },
    mounted() {
       api.getGameList().then(v => {
         console.log(v)
        this.gameList = v.data.data.lists
         console.log(this.gameList[0].id)
      });
    }


}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
