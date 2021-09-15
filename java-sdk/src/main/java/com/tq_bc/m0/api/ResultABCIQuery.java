package com.tq_bc.m0.api;

import cn.hutool.json.JSONObject;
import cn.hutool.json.JSONUtil;
import com.google.gson.annotations.SerializedName;
import tendermint.abci.Types.ResponseQuery;
import java.io.IOException;

public class ResultABCIQuery {
    @SerializedName("response")
    public ResponseQuery response;

    public ResultABCIQuery() {
    }

    public ResultABCIQuery(String jsonStr) throws IOException {
        fromJson(jsonStr);
    }

    public void fromJson(String str) throws IOException {
        JSONObject result = JSONUtil.parseObj(str);
        if (!result.isNull("response")) {
            ResponseQuery.Builder builder = ResponseQuery.newBuilder();
            Api.deserializer.fromJson(builder, result.getStr("response"));
            response = builder.build();
        }
    }

    public String toJson() throws IOException {
        JSONObject result = JSONUtil.createObj();
        if (response != null) {
            result.set("response", JSONUtil.parseObj(Api.deserializer.toJson(response)));
        } else {
            result.set("response", null);
        }
        return result.toStringPretty();
    }

    public static ResultABCIQuery GetABCIQuery(Api client, String path, byte[] data, Long height, boolean prove) throws Exception {
        return client.GetABCIQuery(path, data, height, prove);
    }
}
