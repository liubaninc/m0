// 节点列表api
import { get } from "@/utils/request";
import { Message } from "element-ui";

export let nodeList = async (url, params) => {
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
