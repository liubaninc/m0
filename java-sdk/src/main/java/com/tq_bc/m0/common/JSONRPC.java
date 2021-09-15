package com.tq_bc.m0.common;

import cn.hutool.http.HttpUtil;
import cn.hutool.json.JSONObject;
import cn.hutool.json.JSONUtil;
import java.util.Map;

/**
 * @author mengdexuan on 2021/2/4 15:07.
 */
public class JSONRPC {
	private String baseUrl;

	public JSONRPC(String baseUrl) {
		this.baseUrl = baseUrl;
		if (!baseUrl.endsWith("/")) {
			this.baseUrl += "/";
		}
	}

	public JSONObject request(String action, Map<String, Object> paramMap) throws M0Exception{
		String response = HttpUtil.post(this.baseUrl+action, paramMap);
		JSONObject obj = JSONUtil.parseObj(response);
		if (obj.containsKey("error")) {
			JSONObject error = obj.getJSONObject("error");
			System.out.println(error.toStringPretty());
			throw new M0Exception(error.getInt("code"), error.getStr("data"));
		}
		return obj.getJSONObject("result");
	}

	public class M0Exception extends Exception {
		public Integer code;
		public String message;

		public M0Exception(Integer code, String message) {
			this.code = code;
			this.message = message;
		}
	}
}
