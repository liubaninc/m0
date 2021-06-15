// 资产模块api
import { get } from "@/utils/request";
import { Message } from "element-ui";
// 资产列表
export let assetsList = async (url, params) => {
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
export let assetsDetail = async (url, params) => {
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
      Message.error(data.msg || "获取资产关联地址失败");
    }
  } catch (error) {
    console.log(error);
  }
};
// 资产地址
export let address = async (url, params) => {
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
