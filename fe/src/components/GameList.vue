<template>
  <div class="gameList">
      <el-row>
        <el-col v-for="(item) in gameList" :span="8">
          <div class="card">
            <el-image :src="src"></el-image>
          </div>
          {{item}}
        </el-col>
      </el-row>
  </div>
</template>
<script>
import api from "@/api/main";

export default {
  name: 'Game',
  data () {
    return {
      gameList:[],
      src: 'https://cube.elemecdn.com/6/94/4d3ea53c084bad6931a56d5158a48jpeg.jpeg'
    };
  },
     mounted() {
        api.getGameList().then(v => {
        this.gameList = v.lists
      });
    }

}
</script>
<style scoped>
.card{
  padding: 0;
  margin: 10px 10px;
}
</style>
