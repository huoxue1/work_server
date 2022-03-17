<template>
  <el-container :style="container_style">
    <el-aside :style="aside_width">
      <div
          style="height: 6%"
          class="label_style bor"
      >
        {{label}}
      </div>
      <el-menu
          router="router"
          class="el-menu-demo"
          background-color="#545c64"
          text-color="#fff"
          active-text-color="#ffd04b"
      >
        <el-menu-item index="upload">文件上传</el-menu-item>
        <el-menu-item index="token_manager">Token管理</el-menu-item>
        <el-menu-item :disabled="!admin"  index="work_manager">后台管理</el-menu-item>

      </el-menu>

    </el-aside>
    <el-container style="padding: 0">
      <el-header :style="header_height">

      </el-header>
      <el-main>
        <router-view/>
      </el-main>
      <el-footer class="bor">
        <Foo/>

      </el-footer>
    </el-container>
  </el-container>
</template>



<script>


import Foo from "./views/Foo";
import Api from "./utils/api";




export default {

  components: {Foo},
  data(){
    return {
      label:"作业收集系统",
      container_style :{
        margin: "0 auto",
        height : "100%",
      },
      header_height:{
        padding: 0,
        backgroundColor: "#545c64",
        height: "6%"
      },
      aside_width:{
        backgroundColor: "#545c64",
        padding: 0,
        width: "200px"
      },
      admin: false
    }
  },
  methods:{

  },
  computed : {

  },
  created() {
    this.$router.push("/upload")
    if (window.innerWidth<600){
      this.aside_width.width = "0px"
    }
    Api.check_token().then(data =>{
      this.admin = data.code === 200;
    })
  }
}

</script>


<style lang="scss">

@media screen and (max-width: 500px) {
  .el-aside{
    display: none;
  }
  .el-footer{
    display: none;
  }
  .el-header{
    display: none;
  }
}

.menu{
  font-size: 30px;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  width: 100%;
  height: 100%;
  padding: 0;
  margin: 0 auto;

}

html,
body {
  margin: 0 auto;
  padding: 0;
  height: 100%;
}

.label_style {
  font-size: large;
  font-family: "Microsoft YaHei UI Light",serif;
  font-weight: bold;
  padding-top: 10px;
  color: #42b983;
}

#content{

  width: 100%;
  height: 100%;
}

.bor{
  border: #42b983 1px solid;
}




</style>
