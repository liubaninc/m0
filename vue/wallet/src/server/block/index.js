import { post, get } from "@/utils/request";
/**
 * @param {*} params
 * @returns
 */
export let queryBlockChain = async function(params) {
  try {
    let { code, data, msg } = await get(`/blockchain`, {
      ...params,
    });
    if (code == 200) {
      return data;
    } else {
      console.log(code, msg);
    }
  } catch (error) {
    console.log(error);
  }
};
