package com.tq_bc.m0.api;

import cn.hutool.json.JSONArray;
import cn.hutool.json.JSONObject;
import cn.hutool.json.JSONUtil;
import com.google.common.collect.Lists;
import com.google.gson.annotations.SerializedName;

import java.io.IOException;
import java.util.List;

public class ResultTxSearch {
    @SerializedName("txs")
    public List<ResultTx> txs;
    @SerializedName("total_count")
    public Integer total_count;

    public ResultTxSearch() {
    }

    public ResultTxSearch(String jsonstr) throws IOException {
        fromJson(jsonstr);
    }

    public void fromJson(String str) throws IOException {
        JSONObject result = JSONUtil.parseObj(str);
        if (!result.isNull("txs")) {
            txs = Lists.newArrayList();
            JSONArray list = result.getJSONArray("txs");
            for (Object obj:list) {
                txs.add(new ResultTx(obj.toString()));
            }
        }
        if (!result.isNull("total_count")) {
            total_count = result.getInt("total_count");
        }
    }

    public String toJson() throws IOException {
        JSONObject result = JSONUtil.createObj();
        if (txs != null) {
            JSONArray list = JSONUtil.createArray();
            for (ResultTx tx:txs) {
                list.add(JSONUtil.parseObj(tx.toJson()));
            }
            result.set("txs", list);
        } else {
            result.set("txs", null);
        }
        result.set("total_count", total_count);
        return result.toStringPretty();
    }

    public static ResultTxSearch GetTxSearch(Api client, String query, Integer page, Integer per_page) throws Exception {
        return client.GetTxSearch(query, page, per_page);
    }
}
