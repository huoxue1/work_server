package com.gjs.controller;

import com.gjs.config.Config;
import com.gjs.entity.Work;
import com.gjs.service.WorkService;
import com.gjs.utils.Resp;
import com.gjs.utils.Util;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@RestController
public class WorkController {

    @Resource
    WorkService workService;

    @Resource
    Config config;

    @RequestMapping("/public/get_works")
    public Resp<List<Work>> getWorks(){
        return Resp.ok(workService.queryAll());
    }

    @RequestMapping("/public/get_work/{work_id}")
    public Resp<Work> getWork(@PathVariable Integer work_id){
        return Resp.ok(workService.queryById(work_id));
    }

    @RequestMapping("/check_token")
    public Map<String,Integer> checkToken(@RequestParam("token") String token){
        HashMap<String, Integer> map = new HashMap<>();
        if (token.equals(config.getToken())){
            map.put("code",200);
        }else {
            map.put("code",403);
        }
        return map;
    }
}
