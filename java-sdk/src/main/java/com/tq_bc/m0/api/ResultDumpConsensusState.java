package com.tq_bc.m0.api;

import cn.hutool.json.JSONArray;
import cn.hutool.json.JSONObject;
import cn.hutool.json.JSONUtil;
import com.google.common.collect.Lists;
import com.google.gson.annotations.SerializedName;
import java.util.List;

public class ResultDumpConsensusState {
    class PeerStateInfo {
        @SerializedName("node_address")
        public String node_address;
        @SerializedName("peer_state")
        public JSONObject peer_state;

        public PeerStateInfo() {
        }

        public PeerStateInfo(String jsonStr) {
            fromJson(jsonStr);
        }

        public void fromJson(String str) {
            JSONObject result = JSONUtil.parseObj(str);
            if (!result.isNull("node_address")) {
                node_address = result.getStr("node_address");
            }
            if (!result.isNull("peer_state")) {
                peer_state = result.getJSONObject("peer_state");
            }
        }

        public String toJson() {
            JSONObject result = JSONUtil.createObj();
            result.set("node_address", node_address);
            if (peer_state != null) {
                result.set("peer_state", peer_state);
            } else {
                result.set("peer_state", null);
            }
            return result.toStringPretty();
        }
    }
    @SerializedName("round_state")
    public JSONObject round_state;
    @SerializedName("peers")
    public List<PeerStateInfo> peers;

    public ResultDumpConsensusState() {
    }

    public ResultDumpConsensusState(String jsonStr) {
        fromJson(jsonStr);
    }

    public void fromJson(String str) {
        JSONObject result = JSONUtil.parseObj(str);
        if (!result.isNull("round_state")) {
            round_state = result.getJSONObject("round_state");
        }
        if (!result.isNull("peers")) {
            peers = Lists.newArrayList();
            JSONArray list = result.getJSONArray("peers");
            for (Object obj : list) {
                peers.add(new PeerStateInfo(obj.toString()));
            }
        }
    }

    public String toJson() {
        JSONObject result = JSONUtil.createObj();
        if (round_state != null) {
            result.set("round_state", JSONUtil.parseObj(round_state));
        } else {
            result.set("round_state", null);
        }
        if (peers != null) {
            JSONArray ps =  JSONUtil.createArray();
            for (PeerStateInfo peer: peers) {
                ps.add(JSONUtil.parseObj(peer.toJson()));
            }
            result.set("peers", ps);
        } else {
            result.set("peers", null);
        }
        return result.toStringPretty();
    }

    public static ResultDumpConsensusState GetDumpConsensusState(Api client) throws Exception {
        return client.GetDumpConsensusState();
    }
}

