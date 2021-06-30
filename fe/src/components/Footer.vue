<template>
    <el-row class='footer'>
        <el-col class='icon' :span="8">
            <i class="el-icon-s-promotion"></i>
        </el-col>
        <el-col class='icon' :span="8">
            <i class="el-icon-s-home"></i>
        </el-col>
        <el-col class='icon' :span="8">
            <i class="el-icon-user-solid"></i>
        </el-col>
    </el-row>
</template>

<script>

export default {
    name: 'Footer',
}
</script>

<style scoped>
.footer{
    #height: 40px;
    background-color: green;
    position: fixed;
    left: 0;
    bottom: 0;
    width: 100%;
    background-color: red;
    color: white;
    text-align: center;
}
.icon{
    font-size: 26px;
}

</style>
