package com.gjs.entity;

import com.fasterxml.jackson.annotation.JsonProperty;

import java.io.Serializable;

/**
 * (Work)实体类
 *
 * @author makejava
 * @since 2022-05-15 13:45:25
 */
public class Work implements Serializable {
    private static final long serialVersionUID = 680613630071086608L;

    private Integer id;

    private String name;

    @JsonProperty("end_time")
    private Integer endTime;


    public Integer getId() {
        return id;
    }

    public void setId(Integer id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public Integer getEndTime() {
        return endTime;
    }

    public void setEndTime(Integer endTime) {
        this.endTime = endTime;
    }

}

