package com.tq_bc.m0.api;

import cn.hutool.json.JSONObject;
import cn.hutool.json.JSONUtil;
import com.google.gson.annotations.SerializedName;
import tendermint.types.BlockOuterClass.Block;
import tendermint.types.Types.BlockID;

import java.io.IOException;

public class ResultBlock {
    @SerializedName("block_id")
    public BlockID block_id;
    @SerializedName("block")
    public Block block;

    public ResultBlock() {
    }

    public ResultBlock(String jsonstr) throws IOException {
        fromJson(jsonstr);
    }

    public void fromJson(String str) throws IOException {
        JSONObject result = JSONUtil.parseObj(str);
        if (!result.isNull("block")) {
            Block.Builder builder = Block.newBuilder();
            Api.deserializer.fromJson(builder, result.getStr("block"));
            block = builder.build();
        }
        if (!result.isNull("block_id")) {
            BlockID.Builder builder = BlockID.newBuilder();
            Api.deserializer.fromJson(builder, result.getStr("block_id"));
            block_id = builder.build();
        }
    }

    public String toJson() throws IOException {
        JSONObject result = JSONUtil.createObj();
        if (block_id != null) {
            result.set("block_id", JSONUtil.parseObj(Api.deserializer.toJson(block_id)));
        } else {
            result.set("block_id", null);
        }
        if (block != null) {
            result.set("block", JSONUtil.parseObj(Api.deserializer.toJson(block)));
        } else {
            result.set("block", null);
        }
        return result.toStringPretty();
    }

    public static ResultBlock GetBlock(Api client, Long height) throws Exception {
        return client.GetBlock(height);
    }

    public static ResultBlock GetBlockLatest(Api client) throws Exception {
        return client.GetBlockLatest();
    }
}
