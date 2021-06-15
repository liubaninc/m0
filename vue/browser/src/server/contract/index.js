// 合约模块api
import { get } from "@/utils/request";
import { Message } from "element-ui";
// 合约列表
export let contractList = async (url, params) => {
  try {
    let data = await get(url, params);
    if (data.code == 200) {
      return data.data;
    } else {
      Message.error(data.msg || "获取节点列表失败");
    }
  } catch (error) {
    console.log(error);
  }
};
// 合约详情
export let contractDetail = async (url, params) => {
  try {
    let data = await get(url, params);
    if (data.code == 200) {
      return data.data;
    } else {
      Message.error(data.msg || "获取节点详情失败");
    }
  } catch (error) {
    console.log(error);
  }
};

// 详情列表
export let detailList = async (url, params) => {
  try {
    let data = await get(url, params);
    if (data.code == 200) {
      return data.data;
    } else {
      Message.error(data.msg || "获取节点详情失败");
    }
  } catch (error) {
    console.log(error);
  }
};
