<template>
  <el-table :data="works"  style="width: 100%;margin-left: 20px;margin-right: 20px">
    <el-table-column prop="id" label="ID"/>
    <el-table-column prop="name" label="任务名"/>
    <el-table-column prop="end_time" label="time"/>
    <el-table-column >
      <template #header>
        <el-input v-model="search" size="mini" placeholder="Type to search"/>
      </template>
      <template #default="scope">
        <el-button size="small" @click="handleDelete(scope.row.id)">删除</el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
import Api from "../utils/api";
import Utils from "../utils/utils";

export default {
  name: "WorkerManager",
  data(){
    return{
      works:[]
    }
  },
  created() {
    Api.get_works().then((data)=>{
      this.works = data
      for (let i = 0; i < this.works; i++) {
        this.works[i].format_time = Utils.format_time(this.works[i].end_time)
      }
    })
  },
  methods:{
    handleDelete(id){
      Api.delete_work(id).then(data=>{
        console.log(data)
        Api.get_works().then((data)=>{
          this.works = data
          for (let i = 0; i < this.works; i++) {
            this.works[i].format_time = Utils.format_time(this.works[i].end_time)
          }
        })
      })
    }
  }
}
</script>

<style scoped>


</style>
