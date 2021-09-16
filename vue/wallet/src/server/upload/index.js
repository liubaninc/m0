// @ts-nocheck
import { post, get } from "@/utils/request";
/**
 *
 * @param {*} params
 * @returns
 */
export let upLoadFile = async (params) => {
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
/**
 * 多签合约上传
 * @param {*} params 
 * @returns 
 */
export let upLoadMultipleSign = async (params) => {
  try {
    let { code, data, msg } = await post(`/mcontract/tx/upload`, params);
    if (code == 200) {
      return data;
    } else {
      console.log(code, msg);
    }
  } catch (error) {
    console.log(error);
  }
};

