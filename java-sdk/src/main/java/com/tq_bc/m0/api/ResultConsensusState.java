package com.tq_bc.m0.api;

import cn.hutool.json.JSONObject;
import cn.hutool.json.JSONUtil;
import com.google.gson.annotations.SerializedName;

public class ResultConsensusState {
    @SerializedName("round_state")
    public JSONObject round_state;

    public ResultConsensusState() {
    }

    public ResultConsensusState(String jsonStr) {
        fromJson(jsonStr);
    }

    public void fromJson(String str) {
        JSONObject result = JSONUtil.parseObj(str);
        if (!result.isNull("round_state")) {
            round_state = result.getJSONObject("round_state");
        }
    }

    public String toJson() {
        JSONObject result = JSONUtil.createObj();
        if (round_state != null) {
            result.set("round_state", JSONUtil.parseObj(round_state));
        } else {
            result.set("round_state", null);
        }
        return result.toStringPretty();
    }

    public static ResultConsensusState GetConsensusState(Api client) throws Exception {
        return client.GetConsensusState();
    }
}
