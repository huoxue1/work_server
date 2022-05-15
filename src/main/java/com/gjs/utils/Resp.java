package com.gjs.utils;

public class Resp<T> {

    public Integer code;
    public String msg;
    public T data;
    public String error;

    public static <T> Resp<T> ok(Integer code,T data){
        return new Resp<T>(code,"操作成功",data,"");
    }

    public static <T> Resp<T> ok(T data){
        return new Resp<T>(200,"操作成功",data,"");
    }

    public static <T> Resp<T> error(Integer code, String error){
        return new Resp<>(code,"操作失败",null,error);
    }

    public Integer getCode() {
        return code;
    }

    public void setCode(Integer code) {
        this.code = code;
    }

    public String getMsg() {
        return msg;
    }

    public void setMsg(String msg) {
        this.msg = msg;
    }

    public T getData() {
        return data;
    }

    public void setData(T data) {
        this.data = data;
    }

    public String getError() {
        return error;
    }

    public void setError(String error) {
        this.error = error;
    }

    public Resp() {
    }

    public Resp(Integer code, String msg, T data, String error) {
        this.code = code;
        this.msg = msg;
        this.data = data;
        this.error = error;
    }
}
