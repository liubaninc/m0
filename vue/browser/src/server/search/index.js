import { get } from "@/utils/request";

/**
 * @param {*} url
 * @param {*} params
 */
export let searchConetnt = async (url = "/search", params = {}) => {
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
