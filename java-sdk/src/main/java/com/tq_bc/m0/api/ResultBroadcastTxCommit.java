package com.tq_bc.m0.api;

import cn.hutool.json.JSONObject;
import cn.hutool.json.JSONUtil;
import com.google.gson.annotations.SerializedName;
import tendermint.abci.Types.ResponseDeliverTx;
import tendermint.abci.Types.ResponseCheckTx;

import java.io.IOException;

public class ResultBroadcastTxCommit {
    @SerializedName("check_tx")
    public ResponseCheckTx check_tx;
    @SerializedName("deliver_tx")
    public ResponseDeliverTx deliver_tx;
    @SerializedName("hash")
    public String hash;
    @SerializedName("height")
    public Integer height;

    public ResultBroadcastTxCommit() {
    }

    public ResultBroadcastTxCommit(String jsonStr) throws IOException {
        fromJson(jsonStr);
    }

    public void fromJson(String str) throws IOException {
        JSONObject result = JSONUtil.parseObj(str);
        if (!result.isNull("check_tx")) {
            ResponseCheckTx.Builder builder = ResponseCheckTx.newBuilder();
            Api.deserializer.fromJson(builder, result.getStr("check_tx"));
            check_tx = builder.build();
        }

        if (!result.isNull("deliver_tx")) {
            ResponseDeliverTx.Builder builder = ResponseDeliverTx.newBuilder();
            Api.deserializer.fromJson(builder, result.getStr("deliver_tx"));
            deliver_tx = builder.build();
        }
        hash = result.getStr("hash");
        height = result.getInt("height");
    }

    public String toJson() throws IOException {
        JSONObject result = JSONUtil.createObj();
        if (deliver_tx != null) {
            result.set("deliver_tx", JSONUtil.parseObj(Api.deserializer.toJson(deliver_tx)));
        } else {
            result.set("deliver_tx", null);
        }
        if (check_tx != null) {
            result.set("check_tx", JSONUtil.parseObj(Api.deserializer.toJson(check_tx)));
        } else {
            result.set("check_tx", null);
        }
        result.set("height", height);
        result.set("hash", hash);
        return result.toStringPretty();
    }

    public static ResultBroadcastTxCommit BroadcastTxCommit(Api client, String tx) throws Exception {
        return client.BroadcastTxCommit(tx);
    }
}
