package com.tq_bc.m0.api;

import cn.hutool.json.JSONUtil;
import com.google.gson.annotations.SerializedName;
import java.io.IOException;
import java.util.List;

public class ResultUnconfirmedTxs {
    @SerializedName("n_txs")
    public Integer n_txs;
    @SerializedName("total")
    public Integer total;
    @SerializedName("total_bytes")
    public Long total_bytes;
    @SerializedName("txs")
    public List<String> txs;

    public ResultUnconfirmedTxs() {
    }

    public ResultUnconfirmedTxs(String jsonstr) throws IOException {
        fromJson(jsonstr);
    }

    public void fromJson(String str) {
        ResultUnconfirmedTxs ret = JSONUtil.toBean(str, ResultUnconfirmedTxs.class);
        n_txs = ret.n_txs;
        total = ret.total;
        total_bytes = ret.total_bytes;
        txs = ret.txs;
    }

    public String toJson() {
       return JSONUtil.toJsonPrettyStr(this);
    }
    public static ResultUnconfirmedTxs GetUnconfirmedTxs(Api client, Integer limit) throws Exception {
        return client.GetUnconfirmedTxs(limit);
    }

    public static ResultUnconfirmedTxs GetNumUnconfirmedTxs(Api client) throws Exception {
        return client.GetNumUnconfirmedTxs();
    }
}