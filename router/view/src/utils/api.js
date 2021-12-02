const axios = require("axios");
const { v4: uuidv4 } = require('uuid');

module.exports = {
    base: process.env.VUE_APP_API,
    get_files:async function(work_id){
        let resp = await axios.post(this.base+"/public/get_files/" + work_id)
        return resp.data.data
    },
    get_works:async function(){
        let resp = await axios.post(this.base+"/public/get_works")
        return resp.data.data
    },
    get_work:async function(work_id){
        let resp = await axios.post(this.base+"/public/get_work/"+work_id)
        return resp.data.data
    },
    upload:async function(file){
        let resp = await axios.post(this.base+"/public/upload",file)
        return resp.data
    },
    get_token:function(){
        let token = localStorage.getItem("token");
        if (token === null){
            let uuid = uuidv4()
            localStorage.setItem("token",uuid)
            return uuid
        }else{
            return token
        }
    },
    handRemove: async function (id,token){
        let resp = await axios.post(this.base+"/public/remove_file/"+id+"?token="+token)
        return resp.data
    },
    create_work:async function(data){
        let resp = await axios.post(this.base+"/admin/create_work?token="+this.get_token(),data)
        return resp.data
    },
    delete_work:async function(id){
        let resp = await axios.post(this.base+"/admin/delete_work/"+id+"?token="+this.get_token())
        return resp.data
    }
}