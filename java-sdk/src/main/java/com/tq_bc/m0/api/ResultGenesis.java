package com.tq_bc.m0.api;

import cn.hutool.json.JSONObject;
import cn.hutool.json.JSONUtil;
import com.google.gson.annotations.SerializedName;
import tendermint.abci.Types.ConsensusParams;

import java.io.IOException;
import java.util.List;

public class ResultGenesis {
    class GenesisValidator {
        @SerializedName("address")
        public String address;
        @SerializedName("pub_key")
        public String pub_key;
        @SerializedName("power")
        public Long power;
        @SerializedName("name")
        public String name;
    }
    class GenesisDoc {
        @SerializedName("genesis_time")
        public String genesis_time;
        @SerializedName("chain_id")
        public String chain_id;
        @SerializedName("consensus_params")
        public ConsensusParams consensus_params;
        @SerializedName("validators")
        public List<GenesisValidator> validators;
        @SerializedName("app_hash")
        public String app_hash;
        @SerializedName("app_state")
        public JSONObject app_state;

        public GenesisDoc() {
        }

        public GenesisDoc(String jsonStr) throws IOException {
            fromJson(jsonStr);
        }

        public void fromJson(String str) throws IOException {
            JSONObject result = JSONUtil.parseObj(str);
            if (!result.isNull("genesis_time")) {
                genesis_time = result.getStr("genesis_time");
            }
            if (!result.isNull("chain_id")) {
                chain_id = result.getStr("chain_id");
            }
            if (!result.isNull("consensus_params")) {
                ConsensusParams.Builder builder = ConsensusParams.newBuilder();
                Api.deserializer.fromJson(builder, result.getStr("consensus_params"));
                consensus_params = builder.build();
            }
            if (!result.isNull("validators")) {
                validators = JSONUtil.toList(result.getJSONArray("validators"), GenesisValidator.class);
            }
            if (!result.isNull("app_hash")) {
                app_hash = result.getStr("app_hash");
            }
            if (!result.isNull("app_state")) {
                app_state = result.getJSONObject("app_state");
            }
        }

        public String toJson() throws IOException {
            JSONObject result = JSONUtil.createObj();
            result.set("genesis_time", genesis_time);
            result.set("chain_id", chain_id);
            if (consensus_params != null) {
                result.set("consensus_params", JSONUtil.parseObj(Api.deserializer.toJson(consensus_params)));
            } else {
                result.set("consensus_params", null);
            }
            result.set("validators", validators);
            result.set("app_hash", app_hash);
            result.set("app_state", app_state);
            return JSONUtil.toJsonPrettyStr(this);
        }
    }
    @SerializedName("genesis")
    public GenesisDoc genesis;

    public ResultGenesis() {
    }

    public ResultGenesis(String jsonStr) throws IOException {
        fromJson(jsonStr);
    }

    public void fromJson(String str) throws IOException {
        JSONObject result = JSONUtil.parseObj(str);
        if (!result.isNull("genesis")) {
            genesis = new GenesisDoc(result.getStr("genesis"));
        }
    }

    public String toJson() throws IOException {
        JSONObject result = JSONUtil.createObj();
        result.set("genesis", genesis.toJson());
        return result.toStringPretty();
    }

    public static ResultGenesis GetGenesis(Api client) throws Exception {
        return client.GetGenesis();
    }
}