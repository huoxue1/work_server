package com.gjs.entity;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

import java.io.Serializable;

/**
 * (Files)实体类
 *
 * @author makejava
 * @since 2022-05-15 13:45:28
 */
public class Files implements Serializable {
    private static final long serialVersionUID = -49060249687019952L;

    public Files() {
    }


    private Integer id;

    private String data;

    @JsonProperty("work_id")
    private Integer workId;

    @JsonProperty("file_name")
    private String fileName;

    private Integer size;

    @JsonProperty("upload_time")
    private Integer uploadTime;

    private String token;


    public Integer getId() {
        return id;
    }

    public void setId(Integer id) {
        this.id = id;
    }

    public String getData() {
        return data;
    }

    public void setData(String data) {
        this.data = data;
    }

    public Integer getWorkId() {
        return workId;
    }

    public void setWorkId(Integer workId) {
        this.workId = workId;
    }

    public String getFileName() {
        return fileName;
    }

    public void setFileName(String fileName) {
        this.fileName = fileName;
    }

    public Integer getSize() {
        return size;
    }

    public void setSize(Integer size) {
        this.size = size;
    }

    public Integer getUploadTime() {
        return uploadTime;
    }

    public void setUploadTime(Integer uploadTime) {
        this.uploadTime = uploadTime;
    }

    public String getToken() {
        return token;
    }

    public void setToken(String token) {
        this.token = token;
    }

}

