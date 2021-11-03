package com.tq_bc.m0.api;

import cn.hutool.json.JSONObject;
import cn.hutool.json.JSONUtil;
import com.google.gson.annotations.SerializedName;
import tendermint.abci.Types.ResponseInfo;
import java.io.IOException;

public class ResultABCIInfo {
    @SerializedName("response")
    public ResponseInfo response;

    public ResultABCIInfo() {
    }

    public ResultABCIInfo(String jsonStr) throws IOException {
        fromJson(jsonStr);
    }

    public void fromJson(String str) throws IOException {
        JSONObject result = JSONUtil.parseObj(str);
        if (!result.isNull("response")) {
            ResponseInfo.Builder builder = ResponseInfo.newBuilder();
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


    public static ResultABCIInfo GetABCIInfo(Api client) throws Exception {
        return client.GetABCIInfo();
    }
}
