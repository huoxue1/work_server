package com.gjs.controller;

import com.gjs.config.Config;
import com.gjs.entity.Work;
import com.gjs.service.WorkService;
import com.gjs.utils.Resp;
import com.gjs.utils.Util;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import sun.reflect.generics.reflectiveObjects.NotImplementedException;

import javax.annotation.Resource;
import javax.servlet.http.HttpServletResponse;
import java.io.*;
import java.util.Arrays;
import java.util.Map;
import java.util.Objects;

@Controller
@RequestMapping("/admin")
public class AdminController {

    private static final org.slf4j.Logger logger = LoggerFactory.getLogger(AdminController.class);


    @Resource
    Config config;

    @Resource
    WorkService workService;




    @RequestMapping("/create_work")
    public Resp<Object> createWork(@RequestBody Work work){
        Work work1 = workService.insert(work);
        if (work1 != null){
            return Resp.ok(null);
        }else {
            return Resp.error(502,"数据库插入失败");
        }
    }


    @RequestMapping("/get_zip_result/{work_id}")
    public void getGzipResult(@PathVariable Integer work_id, HttpServletResponse response) throws IOException {
        Work work = workService.queryById(work_id);
//        Util.compressSingleFile(config.getStorage_path()+"/"+work.getName()+"/",
//                config.getStorage_path()+"/"+work.getName()+".zip");
        File file = new File(config.getStorage_path() + "/" + work.getName() + "/");
        String[] childrenList = file.list();
        if (childrenList == null) {
            logger.error("the file is empty");
            return;
        }
        String[] fileList = new String[childrenList.length];
        for (int i = 0; i < childrenList.length; i++) {
            fileList[i] = config.getStorage_path() + "/" + work.getName() + "/" + childrenList[i];
        }
        Util.compressMultipleFiles(config.getStorage_path()+"/"+work.getName()+".zip",
             fileList);
        response.setCharacterEncoding("utf-8");
        response.setContentType("multipart/form-data");
        response.setHeader("Content-Disposition", "attachment;fileName=" + work.getName()+".zip");
        //打开本地文件流
        try (
                InputStream inputStream = new FileInputStream(config.getStorage_path()+"/"+work.getName()+".zip");
                //激活下载操作
                OutputStream os = response.getOutputStream();
                ){
            //循环写入输出流
            byte[] b = new byte[2048];
            int length;
            while ((length = inputStream.read(b)) > 0) {
                os.write(b, 0, length);
            }
        }
    }

    @RequestMapping("/delete_work/{work_id}")
    public Resp<Boolean> deleteWOrk(@PathVariable Integer work_id){

        return Resp.ok(workService.deleteById(work_id));
    }
}
