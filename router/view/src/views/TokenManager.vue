<template>
<div style="width: 100%;height: 100%">
<el-input class="data"  v-model="token" placeholder="请输入token" label="TOKEN:"/>
  <el-button class="data" @click="click"   type="success">请绑定token</el-button>

</div>
</template>

<script>
import Api from "../utils/api";

export default {
  name: "Manager",
  data(){
    return{
      token: ""
    }
  },
  methods:{
    click:function () {
      localStorage.setItem("token",this.token)
    }
  },
  created() {
    this.token = Api.get_token()
  }
}
</script>

<style scoped>
.data{
  display:inline;
  width: 50%;
}

</style>