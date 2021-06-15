import { get } from "../../utils/request";

const PAGE_SIZE = 10;
/**
 * 查询区块信息
 * @param {*} url
 * @param {*} params
 */
export let queryBlockChainInfo = async (url = "/blockchain", params) => {
  try {
    let { code, data } = await get(url, { ...params });
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
 * 查询区块列表
 * @param {*} url
 * @param {*} params
 */
export let queryBlocks = async (
  url = "/blocks",
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
 * 查询区块信息
 * @param {*} url
 * @param {*} params
 */
export let queryBlockInfo = async (url = `/blocks/:id`, params = {}) => {
  try {
    let { code, data } = await get(url, { ...params });
    if (code == 200) {
      return data;
    } else {
      console.log(code);
    }
  } catch (error) {
    console.log(error);
  }
};
