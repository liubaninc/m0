package com.tq_bc.m0.api;

import cn.hutool.json.JSONUtil;
import com.google.gson.annotations.SerializedName;

import java.io.IOException;

public class ResultBroadcastTx {
    public Integer code;
    public String data;
    public String log;
    public String codespace;
    public String hash;

    public ResultBroadcastTx() {
    }

    public ResultBroadcastTx(String jsonStr) throws IOException {
        fromJson(jsonStr);
    }

    public void fromJson(String str) {
        System.out.println(str);
        ResultBroadcastTx ret = JSONUtil.toBean(str, ResultBroadcastTx.class);
        System.out.println(ret);
        code = ret.code;
        data = ret.data;
        log = ret.log;
        codespace = ret.codespace;
        hash = ret.hash;
    }

    public String toJson() {
        return JSONUtil.toJsonPrettyStr(this);
    }


    public static ResultBroadcastTx BroadcastTxAsync(Api client, String tx) throws Exception {
        return client.BroadcastTxAsync(tx);
    }

    public static ResultBroadcastTx BroadcastTxSync(Api client, String tx) throws Exception {
        return client.BroadcastTxSync(tx);
    }
}
