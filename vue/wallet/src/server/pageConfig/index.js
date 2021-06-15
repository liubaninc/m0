import { post, get } from "@/utils/request";

/**
 * @param {*} params
 * @returns
 */
export let queryPageContext = async (params) => {
  try {
    let resContext = await get(`/pConfig/page.config.json`, params);
    return resContext;
  } catch (error) {
    console.log(error);
  }
};
