package com.tq_bc.m0.api;

import cn.hutool.json.JSONObject;
import cn.hutool.json.JSONUtil;
import com.google.gson.annotations.SerializedName;
import tendermint.p2p.Types.DefaultNodeInfo;

import java.io.IOException;
public class ResultStatus {
    class SyncInfo {
        @SerializedName("latest_block_hash")
        public String latest_block_hash;
        @SerializedName("latest_app_hash")
        public String latest_app_hash;
        @SerializedName("latest_block_height")
        public Long latest_block_height;
        @SerializedName("latest_block_time")
        public String latest_block_time;

        @SerializedName("earliest_block_hash")
        public String earliest_block_hash;
        @SerializedName("earliest_app_hash")
        public String earliest_app_hash;
        @SerializedName("earliest_block_height")
        public Long earliest_block_height;
        @SerializedName("earliest_block_time")
        public String earliest_block_time;

        @SerializedName("catching_up")
        public boolean catching_up;
    }
    class ValidatorInfo{
        @SerializedName("address")
        public String address;
        @SerializedName("pub_key")
        public String pub_key;
        @SerializedName("power")
        public Long power;
    }

    @SerializedName("node_info")
    public DefaultNodeInfo node_info;
    @SerializedName("sync_info")
    public SyncInfo sync_info;
    @SerializedName("validator_info")
    public ValidatorInfo validator_info;

    public ResultStatus() {
    }

    public ResultStatus(String jsonStr) throws IOException {
        fromJson(jsonStr);
    }

    public void fromJson(String str) throws IOException {
        JSONObject result = JSONUtil.parseObj(str);
        if (!result.isNull("node_info")) {
            DefaultNodeInfo.Builder builder = DefaultNodeInfo.newBuilder();
            System.out.println(result.getStr("node_info"));
            Api.deserializer.fromJson(builder, result.getStr("node_info"));
            node_info = builder.build();
        }
        if (!result.isNull("sync_info")) {
            sync_info = JSONUtil.toBean(result.getJSONObject("sync_info"), SyncInfo.class);
        }
        if (!result.isNull("validator_info")) {
            validator_info = JSONUtil.toBean(result.getJSONObject("validator_info"), ValidatorInfo.class);
        }
    }

    public String toJson() throws IOException {
        JSONObject result = JSONUtil.createObj();
        if (node_info != null) {
            result.set("node_info", JSONUtil.parseObj(Api.deserializer.toJson(node_info)));
        } else {
            result.set("node_info", null);
        }
        result.set("sync_info", sync_info);
        result.set("validator_info", validator_info);
        return result.toStringPretty();
    }

    public static ResultStatus GetStatus(Api client) throws Exception {
        return client.GetStatus();
    }
}
