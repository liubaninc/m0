package com.tq_bc.m0.api;

import cn.hutool.json.JSONArray;
import cn.hutool.json.JSONObject;
import cn.hutool.json.JSONUtil;
import com.google.common.collect.Lists;
import com.google.gson.annotations.SerializedName;
import lombok.Data;

import java.io.IOException;
import java.util.List;

public class ResultValidators {

    @Data
    class Validator {
        String address;
        String voting_power;
        String proposer_priority;
    }

    @SerializedName("block_height")
    public Long block_height;
    @SerializedName("validators")
    public List<Validator> validators;
    @SerializedName("count")
    public Integer count;
    @SerializedName("total")
    public Integer total;

    public ResultValidators() {
    }

    public ResultValidators(String jsonstr) throws IOException {
        fromJson(jsonstr);
    }

    public void fromJson(String str) throws IOException {
        JSONObject result = JSONUtil.parseObj(str);
        if (!result.isNull("block_height")) {
            block_height = result.getLong("block_height");
        }
        if (!result.isNull("validators")) {
            validators = Lists.newArrayList();
            JSONArray vals =  result.getJSONArray("validators");
            for (Object obj:vals) {
//                Validator.Builder builder = Validator.newBuilder();
//                Api.deserializer.fromJson(builder, obj.toString());
//                validators.add(builder.build());
                validators.add(JSONUtil.toBean(JSONUtil.toJsonStr(obj), Validator.class));
            }
        }
        if (!result.isNull("count")) {
            count = result.getInt("count");
        }
        if (!result.isNull("total")) {
            total = result.getInt("total");
        }
    }

    public String toJson() throws IOException {
        JSONObject result = JSONUtil.createObj();
        result.set("block_height", block_height);
        result.set("count", count);
        result.set("total", total);
        if (validators != null) {
            JSONArray list = JSONUtil.createArray();
            for (Validator val:validators) {
//                list.add(JSONUtil.parseObj(Api.deserializer.toJson(val)));
                list.add(JSONUtil.parseObj(val));
            }
            result.set("validators", list);
        } else {
            result.set("validators", null);
        }
        return result.toStringPretty();
    }


    public static ResultValidators GetValidators(Api client, Long height, Integer page, Integer per_page) throws Exception {
        return client.GetValidators(height, page, per_page);
    }

    public static ResultValidators GetValidatorsLatest(Api client, Integer page, Integer per_page) throws Exception {
        return client.GetValidatorsLatest(page, per_page);
    }
}
