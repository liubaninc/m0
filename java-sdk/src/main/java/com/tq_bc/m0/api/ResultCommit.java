package com.tq_bc.m0.api;

import cn.hutool.json.JSONObject;
import cn.hutool.json.JSONUtil;
import com.google.gson.annotations.SerializedName;
import tendermint.types.Types.SignedHeader;

import java.io.IOException;

public class ResultCommit {
    @SerializedName("signed_header")
    public SignedHeader signed_header;
    @SerializedName("canonical")
    public Boolean canonical;

    public ResultCommit() {
    }

    public ResultCommit(String jsonstr) throws IOException {
        fromJson(jsonstr);
    }

    public void fromJson(String str) throws IOException {
        JSONObject result = JSONUtil.parseObj(str);
        if (!result.isNull("signed_header")) {
            SignedHeader.Builder builder = SignedHeader.newBuilder();
            Api.deserializer.fromJson(builder, result.getStr("signed_header"));
            signed_header = builder.build();
        }
        canonical = result.getBool("canonical");
    }

    public String toJson() throws IOException {
        JSONObject result = JSONUtil.createObj();
        if (signed_header != null) {
            result.set("signed_header", JSONUtil.parseObj(Api.deserializer.toJson(signed_header)));
        } else {
            result.set("signed_header", null);
        }
        result.set("canonical", canonical);
        return result.toStringPretty();
    }

    public static ResultCommit GetCommit(Api client, Long height) throws Exception {
        return client.GetCommit(height);
    }

    public static ResultCommit GetCommitLatest(Api client) throws Exception {
        return client.GetCommitLatest();
    }
}
