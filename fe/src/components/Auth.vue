<template>
    <div class="login-container">
        <el-form ref="form" :model="form" label-width="80px">
            <el-form-item label="用户名">
                <el-input v-model="form.username"></el-input>
            </el-form-item>
            <el-form-item label="密码">
                <el-input v-model="form.password"></el-input>
            </el-form-item>
            <el-form-item label="是否永久登陆">
                <el-switch v-model="form.isForever"></el-switch>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="onSubmit">立即创建</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
import api from '@/api/main.js'
import store from '@/store/store.js'

export default {
    name: 'Auth',
    data() {
        return {
            form: {
                username: '',
                password: '',
                isForever: ''
            }
        }
    },
    methods: {
        cancel() {
          this.$router.push({ name: 'Home'})
        },
        onSubmit() {
          api.auth(this.form, v => {
              store.setUserInfo(v);
              this.$router.push('/home');
          });
        },
    }
}
</script>

<style scoped>
</style>
