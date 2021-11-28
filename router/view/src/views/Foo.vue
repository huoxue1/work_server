<template>
<div style="width: 100%;height: 100%;padding: 10px;">
<el-row>
  <el-col :span="8">
    <el-input v-model="name" placeholder="请输入任务名称" />
  </el-col>
  <el-col :span="8">
    <el-input v-model="end_time" placeholder="请输入截至时间" />
  </el-col>
  <el-col :span="5">
    <el-button type="success" @click="create_work">提交任务</el-button>
  </el-col>
</el-row>
</div>
</template>

<script>
import Api from "../utils/api";

export default {
  name: "Foo",
  data(){
    return{
      name:"",
      end_time:""
    }
  },
  created() {

  },
  methods:{
      create_work:function () {
        Api.create_work({name:this.name,end_time:new Date(this.end_time).getTime()/1000})
      }
  }
}
</script>

<style scoped>

</style>