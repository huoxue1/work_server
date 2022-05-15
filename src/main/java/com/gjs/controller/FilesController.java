package com.gjs.controller;


import com.gjs.config.Config;
import com.gjs.entity.Files;
import com.gjs.entity.Work;
import com.gjs.service.FilesService;
import com.gjs.service.WorkService;
import com.gjs.utils.Resp;
import com.gjs.utils.Util;
import org.slf4j.LoggerFactory;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.multipart.MultipartFile;

import javax.annotation.Resource;
import javax.servlet.http.HttpServletResponse;
import java.io.File;
import java.io.IOException;
import java.util.Date;
import java.util.List;

@RestController
public class FilesController {

    private static final org.slf4j.Logger logger = LoggerFactory.getLogger(FilesController.class);


    @Resource
    FilesService filesService;

    @Resource
    WorkService workService;

    @RequestMapping(value = "/public/upload")
    public Resp<Boolean> upload(
            @RequestParam("file") MultipartFile multipartFile,
            Integer work_id, String token) throws IOException {

        return Resp.ok(filesService.upload(work_id,token,multipartFile));

    }


    @RequestMapping("/public/get_files/{work_id}")
    public Resp<List<Files>> getFiles(@PathVariable Integer work_id){
        logger.info(config.getStorage_path());
        return Resp.ok(filesService.queryByWorkId(work_id));
    }

    @RequestMapping("/public/get_file/{file_id}")
    public Resp<Files> getFile(@PathVariable Integer file_id){
        return Resp.ok(filesService.queryById(file_id));
    }

    @Resource
    Config config;

    @RequestMapping("/public/remove_file/{file_id}")
    public Resp<Boolean> removeFile(@PathVariable Integer file_id, @RequestParam String token){
        Files file = filesService.queryById(file_id);
        // token不等于文件本身的token,也不等于管理员token
        if ((!file.getToken().equals(token)) && (!token.equals(config.getToken()))){
            return Resp.error(403,"token check failed");
        }
        return Resp.ok(filesService.removeFile(file));
    }

    @RequestMapping("/")
    public void home(HttpServletResponse response) throws IOException {
        response.sendRedirect("/static/view/dist/home.html");
    }
}
