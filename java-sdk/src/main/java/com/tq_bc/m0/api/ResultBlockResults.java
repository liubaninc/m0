package com.tq_bc.m0.api;

import cn.hutool.json.JSONArray;
import cn.hutool.json.JSONObject;
import cn.hutool.json.JSONUtil;
import com.google.common.collect.Lists;
import com.google.gson.annotations.SerializedName;
import tendermint.abci.Types.ResponseDeliverTx;
import tendermint.abci.Types.Event;
import tendermint.abci.Types.ValidatorUpdate;
import tendermint.abci.Types.ConsensusParams;

import java.io.IOException;
import java.util.List;

public class ResultBlockResults {
    @SerializedName("height")
    public long height;
    @SerializedName("txs_results")
    public List<ResponseDeliverTx> txs_results;
    @SerializedName("begin_block_events")
    public List<Event> begin_block_events;
    @SerializedName("end_block_events")
    public List<Event> end_block_events;
    @SerializedName("validator_updates")
    public List<ValidatorUpdate> validator_updates;
    @SerializedName("consensus_param_updates")
    public ConsensusParams consensus_param_updates;

    public ResultBlockResults() {
    }

    public ResultBlockResults(String jsonStr) throws IOException {
        fromJson(jsonStr);
    }

    public void fromJson(String str) throws IOException {
        JSONObject result = JSONUtil.parseObj(str);
        height = result.getLong("height");
        if (!result.isNull("txs_results")) {
            txs_results = Lists.newArrayList();
            JSONArray list_txs = result.getJSONArray("txs_results");
            for (Object obj : list_txs) {
                ResponseDeliverTx.Builder builder = ResponseDeliverTx.newBuilder();
                Api.deserializer.fromJson(builder, JSONUtil.toJsonStr(obj));
                txs_results.add(builder.build());
            }
        }
        if (!result.isNull("begin_block_events")) {
            begin_block_events = Lists.newArrayList();
            JSONArray list_events_bg = result.getJSONArray("begin_block_events");
            for (Object obj : list_events_bg) {
                Event.Builder builder = Event.newBuilder();
                Api.deserializer.fromJson(builder, JSONUtil.toJsonStr(obj));
                begin_block_events.add(builder.build());
            }
        }
        if (!result.isNull("end_block_events")) {
            end_block_events = Lists.newArrayList();
            JSONArray list_events_ed = result.getJSONArray("end_block_events");
            for (Object obj : list_events_ed) {
                Event.Builder builder = Event.newBuilder();
                Api.deserializer.fromJson(builder, JSONUtil.toJsonStr(obj));
                end_block_events.add(builder.build());
            }
        }
        if (!result.isNull("validator_updates")) {
            validator_updates = Lists.newArrayList();
            JSONArray list_val = result.getJSONArray("validator_updates");
            for (Object obj : list_val) {
                ValidatorUpdate.Builder builder = ValidatorUpdate.newBuilder();
                Api.deserializer.fromJson(builder, JSONUtil.toJsonStr(obj));
                validator_updates.add(builder.build());
            }
        }
        if (!result.isNull("consensus_param_updates")) {
            ConsensusParams.Builder builder = ConsensusParams.newBuilder();
            Api.deserializer.fromJson(builder, result.getStr("consensus_param_updates"));
            consensus_param_updates = builder.build();
        }
    }

    public String toJson() throws IOException {
        JSONObject result = JSONUtil.createObj();
        result.set("height", height);
        if (txs_results != null) {
            JSONArray txs = JSONUtil.createArray();
            for (ResponseDeliverTx tx: txs_results) {
                txs.add(JSONUtil.parseObj(Api.deserializer.toJson(tx)));
            }
            result.set("txs_results", txs);
        } else {
            result.set("txs_results", null);
        }
        if (begin_block_events != null) {
            JSONArray events = JSONUtil.createArray();
            for (Event event: begin_block_events) {
                events.add(JSONUtil.parseObj(Api.deserializer.toJson(event)));
            }
            result.set("begin_block_events", events);
        } else {
            result.set("begin_block_events", null);
        }
        if (end_block_events != null) {
            JSONArray events = JSONUtil.createArray();
            for (Event event: end_block_events) {
                events.add(JSONUtil.parseObj(Api.deserializer.toJson(event)));
            }
            result.set("end_block_events", events);
        } else {
            result.set("end_block_events", null);
        }
        if (validator_updates != null) {
            JSONArray vals = JSONUtil.createArray();
            for (ValidatorUpdate val: validator_updates) {
                vals.add(JSONUtil.parseObj(Api.deserializer.toJson(val)));
            }
            result.set("validator_updates", vals);
        } else {
            result.set("validator_updates", null);
        }
        if (consensus_param_updates != null) {
            result.set("consensus_param_updates", JSONUtil.parseObj(Api.deserializer.toJson(consensus_param_updates)));
        } else {
            result.set("consensus_param_updates", null);
        }
        return result.toStringPretty();
    }

    public static ResultBlockResults GetBlockResults(Api client, Long height) throws Exception {
        return client.GetBlockResults(height);
    }

    public static ResultBlockResults GetBlockResultsLatest(Api client) throws Exception {
        return client.GetBlockResultsLatest();
    }
}
