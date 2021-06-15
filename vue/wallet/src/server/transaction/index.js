import { post, get } from "@/utils/request";

const PAGE_SIZE = 20;
/**
 * @param {*} params
 * @returns
 */
export let queryTrLists = async (params) => {
  try {
    let { code, data, msg } = await get(
      `/addresses/${params.address}/transactions`,
      {
        coin: "",
        type: "",
        page_size: PAGE_SIZE,
        ...params,
      }
    );
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
 * 查询交易详情
 * @param {*} params
 * @returns
 */
export let queryTrxDetail = async (params) => {
  try {
    let { code, data, msg } = await get(`/transactions/${params.hash}`, {
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

/**
 * @param {*} params
 * @returns
 */
export let updownSignFile = async (params) => {
  try {
    let { code, data, msg } = await get(`/download/${params.hash}`, {
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

/**
 * @param {*} params
 * @returns
 */
export let trxSign = async function(params) {
  try {
    let { code, data, msg } = await post(`/tx/sign`, {
      ...params,
    });
    if (code == 200) {
      return data;
    } else {
      this.$message.error(msg);
      console.log(code, msg);
    }
  } catch (error) {
    console.log(error);
  }
};
