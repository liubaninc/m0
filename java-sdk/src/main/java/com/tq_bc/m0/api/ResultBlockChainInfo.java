package com.tq_bc.m0.api;

import cn.hutool.json.JSONArray;
import cn.hutool.json.JSONObject;
import cn.hutool.json.JSONUtil;
import com.google.common.collect.Lists;
import com.google.gson.annotations.SerializedName;
import tendermint.types.Types.BlockMeta;

import java.io.IOException;
import java.util.List;

public class ResultBlockChainInfo {
    @SerializedName("last_height")
    public long last_height;
    @SerializedName("block_metas")
    public List<BlockMeta> block_metas;

    public ResultBlockChainInfo() {
    }

    public ResultBlockChainInfo(String jsonStr) throws IOException {
        fromJson(jsonStr);
    }

    public void fromJson(String str) throws IOException {
        JSONObject result = JSONUtil.parseObj(str);
        last_height = result.getLong("last_height");
        if (!result.isNull("block_metas")) {
            block_metas = Lists.newArrayList();
            JSONArray list = result.getJSONArray("block_metas");
            for (Object obj : list) {
                BlockMeta.Builder builder = BlockMeta.newBuilder();
                Api.deserializer.fromJson(builder, JSONUtil.toJsonStr(obj));
                block_metas.add(builder.build());
            }
        }
    }

    public String toJson() throws IOException {
        JSONObject result = JSONUtil.createObj();
        result.set("last_height", last_height);
        if (block_metas != null) {
            JSONArray blocks = JSONUtil.createArray();
            for (BlockMeta blk: block_metas) {
                blocks.add(JSONUtil.parseObj(Api.deserializer.toJson(blk)));
            }
            result.set("block_metas", blocks);
        } else {
            result.set("block_metas", null);
        }
        return result.toStringPretty();
    }

    public static ResultBlockChainInfo GetBlockChainInfo(Api client, Long min_height, Long max_height) throws Exception {
        return client.GetBlockChainInfo(min_height, max_height);
    }
}