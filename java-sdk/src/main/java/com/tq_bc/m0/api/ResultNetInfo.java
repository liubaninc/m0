package com.tq_bc.m0.api;

import cn.hutool.json.JSONArray;
import cn.hutool.json.JSONObject;
import cn.hutool.json.JSONUtil;
import com.google.common.collect.Lists;
import com.google.gson.annotations.SerializedName;
import tendermint.p2p.Types.DefaultNodeInfo;
import java.io.IOException;
import java.util.List;

public class ResultNetInfo {
    class Peer {
        @SerializedName("node_info")
        public DefaultNodeInfo node_info;
        @SerializedName("is_outbound")
        public boolean is_outbound;
//        @SerializedName("connection_status")
//        public String connection_status;
        @SerializedName("remote_ip")
        public String remote_ip;

        public Peer() {
        }

        public Peer(String jsonstr) throws IOException {
            fromJson(jsonstr);
        }

        public void fromJson(String str) throws IOException {
            JSONObject result = JSONUtil.parseObj(str);
            if (!result.isNull("node_info")) {
                DefaultNodeInfo.Builder builder = DefaultNodeInfo.newBuilder();
                Api.deserializer.fromJson(builder, result.getStr("node_info"));
                node_info = builder.build();
            }
            if (!result.isNull("is_outbound")) {
                is_outbound = result.getBool("is_outbound");
            }
            if (!result.isNull("remote_ip")) {
                remote_ip = result.getStr("remote_ip");
            }
        }
        public String toJson() throws IOException {
            JSONObject result = JSONUtil.createObj();
            if (node_info != null) {
                result.set("node_info", JSONUtil.parseObj(Api.deserializer.toJson(node_info)));
            } else {
                result.set("node_info", null);
            }
            result.set("is_outbound", is_outbound);
            result.set("remote_ip", remote_ip);
            return result.toStringPretty();
        }
    }

    @SerializedName("listening")
    public boolean listening;
    @SerializedName("listeners")
    public List<String> listeners;
    @SerializedName("n_peers")
    public Integer n_peers;
    @SerializedName("peers")
    public List<Peer> peers;

    public ResultNetInfo() {
    }

    public ResultNetInfo(String jsonstr) throws IOException {
        fromJson(jsonstr);
    }

    public void fromJson(String str) throws IOException {
        JSONObject result = JSONUtil.parseObj(str);
        if (!result.isNull("listening")) {
            listening = result.getBool("listening");
        }
        if (!result.isNull("n_peers")) {
            n_peers = result.getInt("n_peers");
        }
        if (!result.isNull("listeners")) {
            listeners = JSONUtil.toList(result.getJSONArray("listeners"), String.class);
        }
        if (!result.isNull("peers")) {
            peers = Lists.newArrayList();
            JSONArray ps = result.getJSONArray("peers");
            for (Object obj:ps) {
                peers.add(new Peer(obj.toString()));
            }
        }
    }
    public String toJson() throws IOException {
        JSONObject result = JSONUtil.createObj();
        result.set("listening", listening);
        result.set("n_peers", n_peers);
        result.set("listeners", listeners);
        if (peers != null) {
            JSONArray ps = JSONUtil.createArray();
            for (Peer p:peers) {
                ps.add(JSONUtil.parseObj(p.toJson()));
            }
            result.set("peers", ps);
        } else {
            result.set("peers", null);
        }
        return result.toStringPretty();
    }

    public static ResultNetInfo GetNetInfo(Api client) throws Exception {
        return client.GetNetInfo();
    }
}
