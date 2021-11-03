package com.tq_bc.m0.api;

import cn.hutool.json.JSONObject;
import cn.hutool.json.JSONUtil;
import com.google.gson.annotations.SerializedName;
import tendermint.abci.Types.ResponseDeliverTx;
import tendermint.types.Types.TxProof;
import java.io.IOException;

public class ResultTx {
    @SerializedName("hash")
    public String hash;
    @SerializedName("height")
    public long height;
    @SerializedName("index")
    public Integer index;
    @SerializedName("tx_result")
    public ResponseDeliverTx tx_result;
    @SerializedName("tx")
    public String tx;
    @SerializedName("proof")
    public TxProof proof;

    public ResultTx() {
    }

    public ResultTx(String jsonstr) throws IOException {
        fromJson(jsonstr);
    }

    public void fromJson(String str) throws IOException {
        JSONObject result = JSONUtil.parseObj(str);
        if (!result.isNull("hash")) {
            hash = result.getStr("hash");
        }
        if (!result.isNull("height")) {
            height = result.getLong("height");
        }
        if (!result.isNull("index")) {
            index = result.getInt("index");
        }
        if (!result.isNull("tx_result")) {
            ResponseDeliverTx.Builder builder =ResponseDeliverTx.newBuilder();
            Api.deserializer.fromJson(builder, result.getStr("tx_result"));
            tx_result = builder.build();
        }
        if (!result.isNull("tx")) {
            tx = result.getStr("tx");
        }
        if (!result.isNull("proof")) {
            TxProof.Builder builder = TxProof.newBuilder();
            Api.deserializer.fromJson(builder, result.getStr("proof"));
            proof = builder.build();
        }
    }

    public String toJson() throws IOException {
        JSONObject result = JSONUtil.createObj();
        result.set("hash", hash);
        result.set("height", height);
        result.set("index", index);
        if (tx_result != null) {
            result.set("proof", Api.deserializer.toJson(tx_result));
        } else {
            result.set("txs", null);
        }
        result.set("tx", tx);
        if (proof != null) {
            result.set("proof", Api.deserializer.toJson(proof));
        } else {
            result.set("proof", null);
        }
        return result.toStringPretty();
    }

    public static ResultTx GetTx(Api client, String hash) throws Exception {
        return client.GetTx(hash);
    }
}
