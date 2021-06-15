// 资产地址模块api
import { get } from "@/utils/request";
import { Message } from "element-ui";

// 地址资产详情
export let addressDetail = async (url, params) => {
  try {
    let data = await get(url, params);
    if (data.code == 200) {
      return data.data;
    } else {
      Message.error(data.msg || "获取资产信息失败");
    }
  } catch (error) {
    console.log(error);
  }
};
// 交易列表
export let addressList = async (url, params) => {
  try {
    let data = await get(url, params);
    if (data.code == 200) {
      return data.data;
    } else {
      Message.error(data.msg || "获取交易列表失败");
    }
  } catch (error) {
    console.log(error);
  }
};

// 资产名称
export let addressName = async (url, params) => {
  try {
    let data = await get(url, params);
    if (data.code == 200) {
      return data.data;
    } else {
      Message.error(data.msg || "获取资产名称失败");
    }
  } catch (error) {
    console.log(error);
  }
};
