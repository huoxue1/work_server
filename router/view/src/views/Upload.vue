<template>
<div style="width: 100%;height: 100%;border: #42b983 1px solid">
  <div class="header">
      <span style="float: left">请选择上传位置：</span>
      <el-select style="float: left" v-model="selected_work_id">
        <el-option
        v-for="work in works"
        :key="work.id"
        :value="work.id"
        :label="work.name"
        >

        </el-option>
      </el-select>
    <el-upload
        action="/up"
        auto-upload="false"
        :before-upload="upload"
    >
      <el-button type="success" @click="upload">点击上传文件</el-button>
    </el-upload>


    <div style="float: right;margin-right: 100px">截至时间<el-link :href="link">：</el-link><span v-text="selected_work.end_time"></span></div>
  </div>
  <div class="body">
  <el-table :data="files">
      <el-table-column prop="file_name" label="文件名"/>
    <el-table-column prop="size" label="文件大小"/>
    <el-table-column prop="upload_time" label="上次更新时间"/>
    <el-table-column  label="上次更新时间">
      <template #default="scope">
        <el-button size="mini" :disabled="scope.row.canRemove" @click="handRemove(scope.row.id)"
        >删除</el-button
        >
        >
      </template>
    </el-table-column>
  </el-table>
  </div>

</div>
</template>

<script>
import Api from "../utils/api";
import Utils from "../utils/utils";

export default {
  name: "Upload",
  data(){
    return{
      works:[],
      selected_work_id:1,
      selected_work:{},
      files:[],
      link:"/admin/get_zip_result/"+this.selected_work_id+"?token="+localStorage.getItem("token"),
      token:""
    }
  },
  watch:{
    selected_work_id(){
      this.link = Api.base+"/admin/get_zip_result/"+this.selected_work_id+"?token="+localStorage.getItem("token")
      Api.get_work(this.selected_work_id).then((data)=>{
        this.selected_work = data
        console.log(data)
        this.selected_work.end_time = Utils.format_time(this.selected_work.end_time,true)
      })

      Api.get_files(this.selected_work_id).then((resp)=>{
        this.files = resp
        for (let i = 0; i < this.files.length; i++) {
          this.files[i].size = Utils.get_size(this.files[i].size)
          this.files[i].upload_time = Utils.format_time(this.files[i].upload_time)
        }
      })
    }
  },
  created() {
    console.log(Api.base)
    this.token = Api.get_token()
    Api.get_works().then((data)=>{
      this.works = data
      this.selected_work_id = data[0].id
      this.link = Api.base+"/admin/get_zip_result/"+this.selected_work_id+"?token="+localStorage.getItem("token")

      this.selected_work = data[0]
      this.selected_work.end_time = Utils.format_time(this.selected_work.end_time,true)
      Api.get_files(data[0].id).then((resp)=>{
        this.files = resp
        for (let i = 0; i < this.files.length; i++) {
          this.files[i].size = Utils.get_size(this.files[i].size)
          this.files[i].upload_time = Utils.format_time(this.files[i].upload_time)
        }
      })
    })
  },
  methods:{
    handRemove(id){
      Api.handRemove(id,Api.get_token())
    },
    click:function () {
      alert(1)
    },
    upload:async function (file) {
      // let files = await window.showOpenFilePicker({
      //   multiple: false,
      // });
      // let file = await files[0].getFile();
      let upload = {}
      Utils.fileToBase64(file,async (bs4) => {
        upload.data = bs4.replace(/^data:.*?;base64,/, "")
        upload.size = file.size
        upload.file_name = file.name
        upload.work_id = this.selected_work_id
        upload.token = Api.get_token()
        await Api.upload(upload)
        Api.get_files(this.selected_work_id).then((resp)=>{
          this.files = resp
          for (let i = 0; i < this.files.length; i++) {
            this.files[i].size = Utils.get_size(this.files[i].size)
            this.files[i].upload_time = Utils.format_time(this.files[i].upload_time)

          }
          return false
        })
      })
    }
  }
}
</script>


<script>

</script>

<style scoped>
.header{
  width: 100%;
  height: 10%;
  border: #42b983 1px solid;
  padding-left: 20px;
  padding-top: 10px;
}

.body{
  width: 100%;
  height: 85%;
}

</style>