// @ts-nocheck
import { post, get } from "@/utils/request";
/**
 *
 * @param {*} params
 * @returns
 */
export let upLoadFile = async (params) => {
  console.log("====>upLoadFile", params.get("file"));
  try {
    let { code, data, msg } = await post(`/tx/upload`, params);
    if (code == 200) {
      return data;
    } else {
      console.log(code, msg);
    }
  } catch (error) {
    console.log(error);
  }
};
