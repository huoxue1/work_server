<template>
<div style="width: 100%;height: 100%;">
  <div class="header" style="width: 100%;margin-top: 20px">
    <div style="width: 30%;float: left">
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
    </div>
    <div style="width: 15%;float:left;">
      <el-upload
          :action="base+'/public/upload'"
          :auto-upload="true"
          :on-success="uploadSuccess"
          :on-progress="upload"
          :before-upload="beforeUpload"
          :data="{'work_id':selected_work_id,'token':token,'type':'file'}"
      >
        <el-button type="success" @click="upload">点击上传文件</el-button>

      </el-upload>

    </div>
    <div style="width: 15%;float: left">
      <el-button type="success" @click="uploadDir">点击上传文件夹</el-button>
    </div>


    <div style="width: 30%;float: right">
      <div style="float: left;margin-right: 100px">截至时间<el-link :href="link">：</el-link><span style="color: red" v-text="selected_work.end_time"></span></div>
      <span style="float:left;">上传人数： <span style="color: red" v-text="files.length"></span>人</span>
    </div>

  </div>
  <el-drawer :model-value="draw.enable" title="上传进度">
    <span>{{this.draw.file_name}}</span>
    <el-progress :percentage="draw.pro"></el-progress>
  </el-drawer>
  <div class="body">
  <el-table :data="files">
      <el-table-column sortable  prop="file_name" label="文件名"/>
    <el-table-column sortable prop="size" label="文件大小"/>
    <el-table-column sortable prop="upload_time" label="上传时间"/>
    <el-table-column  label="">
      <template #default="scope">
        <el-button size="mini" :disabled="!is_admin ? scope.row.token !== this.token : false" type="danger" @click="handRemove(scope.row.id)"
        >删除</el-button
        >
        <el-button size="mini" type="warning" :disabled="!is_admin ? scope.row.token !== this.token : false" @click="handleRename(scope.row.id,scope.row.file_name)"
        >重命名</el-button
        >
        <el-button size="mini" type="info" :disabled="!is_admin ? scope.row.token !== this.token : false" @click="handDownload(scope.row.id,scope.row.file_name)"
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
import axios from "axios";

import {ElMessage, ElMessageBox} from "element-plus";

export default {
  name: "Upload",
  components: {},
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
    uploadDir(){
      let dir = document.createElement("input")
      dir.webkitdirectory = true
      dir.type = "file"
      dir.ref = "file"
      dir.onchange = ()=>{
        let form = new FormData();
        let fileNames = [];
        for (let i = 0; i < dir.files.length; i++) {
          let f = dir.files.item(i)
          form.append("file",f)
          fileNames.push(f.webkitRelativePath)
        }
        form.set("fileNames",fileNames.join(","))
        form.set("work_id",this.selected_work_id)
        form.set("token",Api.get_token())
        form.set("type","dir")
        axios.post(Api.base+"/public/upload",form,{headers:{ "Content-Type": "multipart/form-data" },onUploadProgress:(e)=>{
            this.draw.enable = true
            this.draw.file_name = fileNames[0].split('/')[0]
            this.draw.pro = ((e.loaded/e.total)*100).toFixed(2).valueOf()
        }}).then(resp => {
          console.log(resp)
          this.flush_files()
          this.draw.enable = false
        }).catch(()=>{this.draw.enable = false})

      }
      dir.click()
    },

    handleRename(file_id,file_name){
      ElMessageBox.prompt("请输入更改后的文件名","输入",{cancelButtonText:"退出",confirmButtonText:"确认",inputValue:file_name}).then(value => {
        if (value.value === ""){
          return
        }
        Api.renameFile(this.selected_work_id,file_id,value.value).then(() => {
          ElMessage.success({message:"修改成功"})
          this.flush_files()
        }).catch(err => {
          ElMessage.error({message:err})
        })
      })
    },
    changesData(){
      console.log(this.$refs.file.files);
    },
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

.border_right{
  border-right: red 2px solid;
}

.header{
  width: 100%;
  height: 10%;
  border-bottom: #8592c7 2px solid;
}

.body{
  width: 100%;
  height: 85%;
}

</style>
