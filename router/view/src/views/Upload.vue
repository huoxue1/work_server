<template>
<div style="width: 100%;height: 100%;border: #42b983 1px solid">
  <div class="header">
      <span style="float: left">请选择上传位置：</span>
      <el-select style="float: left" placeholder="请选择上传位置" v-model="selected_work_id">
        <el-option
        v-for="work in works"
        :key="work.id"
        :value="work.id"
        :label="work.name"
        >
        </el-option>
      </el-select>
    <el-upload
        :action="base+'/public/upload'"
        :auto-upload="true"
        :on-success="uploadSuccess"
        :on-progress="upload"
        :before-upload="beforeUpload"
        :data="{'work_id':selected_work_id,'token':token}"
    >
      <el-button type="success" @click="upload">点击上传文件</el-button>
      <span>上传人数： <span v-text="files.length"></span></span>
    </el-upload>
    <el-drawer :model-value="draw.enable" title="上传进度">
      <span>{{this.draw.file_name}}</span>
      <el-progress :percentage="draw.pro"></el-progress>
    </el-drawer>


    <div style="float: right;margin-right: 100px">截至时间<el-link :href="link">：</el-link><span v-text="selected_work.end_time"></span></div>
  </div>
  <div class="body">
  <el-table :data="files">
      <el-table-column prop="file_name" label="fileName"/>
    <el-table-column prop="size" label="fileSize"/>
    <el-table-column prop="upload_time" label="uploadTime"/>
    <el-table-column  label="action">
      <template #default="scope">
        <el-button size="mini" :disabled="!is_admin ? scope.row.token !== this.token : false" type="danger" @click="handRemove(scope.row.id)"
        >删除</el-button
        >

        <el-button size="mini" :disabled="!is_admin ? scope.row.token !== this.token : false" @click="handDownload(scope.row.id,scope.row.file_name)"
        >下载</el-button
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
import {ElMessage, ElMessageBox} from "element-plus";

export default {
  name: "Upload",
  data(){
    return{
      works:[],
      selected_work_id:1,
      selected_work:{},
      files:[],
      link:"/admin/get_zip_result/"+this.selected_work_id+"?token="+localStorage.getItem("token"),
      token:"",
      base: Api.base,
      is_admin:false,

      draw:{
        file_name: "",
        enable: false,
        pro: 0,
      }
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
        this.files.sort((a,b)=>{if (a.upload_time <= b.upload_time){return 1}else {return -1}})
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
    Api.check_token().then(resp => {
      this.is_admin = resp.code === 200
      console.log(this.is_admin)
    })
    Api.get_works().then((data)=>{
      this.works = data
      this.selected_work_id = data[0].id
      this.link = Api.base+"/admin/get_zip_result/"+this.selected_work_id+"?token="+localStorage.getItem("token")

      this.selected_work = data[0]
      this.selected_work.end_time = Utils.format_time(this.selected_work.end_time,true)
      Api.get_files(data[0].id).then((resp)=>{
        this.files = resp
        this.files.sort((a,b)=>{if (a.upload_time <= b.upload_time){return 1}else {return -1}})
        for (let i = 0; i < this.files.length; i++) {
          this.files[i].size = Utils.get_size(this.files[i].size)
          this.files[i].upload_time = Utils.format_time(this.files[i].upload_time)
        }
      })
    })
  },
  methods:{
    handRemove(id){
      ElMessageBox.confirm('你确定要删除吗?', '警告！',{confirmButtonText:"确认",cancelButtonText:"取消",type:"warning"}).then(()=>{
        Api.handRemove(id,Api.get_token()).then(() => {
          this.flush_files()
        })
      }).catch(()=>{

      })

    },
    handDownload(id,file_name){
      let  a = document.createElement("a")
      a.href = Api.base+`/public/download/${this.selected_work_id}/${id}?token=${Api.get_token()}`
      a.download = file_name
      console.log(file_name)
      a.click()
    },
    click:function () {
      alert(1)
    },
    flush_files(){
      Api.get_files(this.selected_work_id).then((resp)=>{
        this.files = resp
        this.files.sort((a,b)=>{if (a.upload_time <= b.upload_time){return 1}else {return -1}})
        for (let i = 0; i < this.files.length; i++) {
          this.files[i].size = Utils.get_size(this.files[i].size)
          this.files[i].upload_time = Utils.format_time(this.files[i].upload_time)
        }

      })
    },
    uploadSuccess(){
      this.draw.enable = false
      ElMessage.success("文件上传成功")
      this.flush_files()
    },
    beforeUpload: function (file) {
      this.draw.enable = true
      this.draw.file_name = file.name
    },
    upload:async function (evt) {
      this.draw.pro = parseInt(evt.percent)
    }
  }
}
</script>


<style scoped>


@media screen and (max-width: 500px) {
  .el-upload{
    width: 100%;
  }
}

.header{
  width: 100%;
  height: 10%;
  border: #42b983 1px solid;
}

.body{
  width: 100%;
  height: 85%;
}

</style>
