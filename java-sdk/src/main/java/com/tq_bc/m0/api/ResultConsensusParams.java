package com.tq_bc.m0.api;

import cn.hutool.json.JSONObject;
import cn.hutool.json.JSONUtil;
import com.google.gson.annotations.SerializedName;
import tendermint.abci.Types.ConsensusParams;
import java.io.IOException;

public class ResultConsensusParams {
    @SerializedName("block_height")
    public long block_height;
    @SerializedName("consensus_params")
    public ConsensusParams consensus_params;

    public ResultConsensusParams() {
    }

    public ResultConsensusParams(String jsonstr) throws IOException {
        fromJson(jsonstr);
    }

    public void fromJson(String str) throws IOException {
        JSONObject result = JSONUtil.parseObj(str);
        if (!result.isNull("consensus_params")) {
            ConsensusParams.Builder builder = ConsensusParams.newBuilder();
            Api.deserializer.fromJson(builder, result.getStr("consensus_params"));
            consensus_params = builder.build();
        }
        block_height = result.getLong("block_height");
    }

    public String toJson() throws IOException {
        JSONObject result = JSONUtil.createObj();
        if (consensus_params != null) {
            result.set("consensus_params", JSONUtil.parseObj(Api.deserializer.toJson(consensus_params)));
        } else {
            result.set("consensus_params", null);
        }
        result.set("block_height", block_height);
        return result.toStringPretty();
    }

    public static ResultConsensusParams GetConsensusParams(Api client, Long height) throws Exception {
        return client.GetConsensusParams(height);
    }

    public static ResultConsensusParams GetConsensusParamsLatest(Api client) throws Exception {
        return client.GetConsensusParamsLatest();
    }
}
