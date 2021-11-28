<template>
div
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
  }
}
</script>

<style scoped>

</style>