package com.tq_bc.m0.api;

import cn.hutool.json.JSONUtil;

public class ResultHealth {
    public ResultHealth() {
    }

    public ResultHealth(String jsonStr) {
        fromJson(jsonStr);
    }

    public void fromJson(String str) {

    }

    public String toJson() {
        return JSONUtil.toJsonPrettyStr(this);
    }

    public static ResultHealth GetHealth(Api client) throws Exception {
        return client.GetHealth();
    }
}
