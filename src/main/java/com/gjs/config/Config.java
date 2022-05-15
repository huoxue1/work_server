package com.gjs.config;

public class Config {

    public void init(){

    }

    private String token;

    private String storage_path;

    public Config(String token, String storage_path) {
        this.token = token;
        this.storage_path = storage_path;
    }

    public Config() {
    }

    public String getToken() {
        return token;
    }

    public void setToken(String token) {
        this.token = token;
    }

    public String getStorage_path() {
        return storage_path;
    }

    public void setStorage_path(String storage_path) {
        this.storage_path = storage_path;
    }
}
