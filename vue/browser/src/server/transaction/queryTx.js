import { get } from "../../utils/request";

const PAGE_SIZE = 10;

/**
 * 查询最近交易信息
 * @param {*} url
 * @param {*} params
 */
export let queryTxInfo = async (
  url = "/charts",
  params = {
    page_num: 1,
  }
) => {
  try {
    let { code, data } = await get(url, { page_size: PAGE_SIZE, ...params });
    if (code == 200) {
      return data;
    } else {
      console.log(code);
    }
  } catch (error) {
    console.log(error);
  }
};

/**
 * 查询交易列表
 * @param {*} url
 * @param {*} params
 */
export let queryTxList = async (
  url = "/transactions",
  params = { page_num: 1 }
) => {
  try {
    let { code, data } = await get(url, { page_size: PAGE_SIZE, ...params });
    if (code == 200) {
      return data;
    } else {
      console.log(code);
    }
  } catch (error) {
    console.log(error);
  }
};
