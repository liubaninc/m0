// @ts-nocheck
import { post, get } from "@/utils/request";

/**
 *
 * @param {*} params
 * @returns
 */
export let queryAssetLists = async (params) => {
  try {
    let { code, data, msg } = await get(`/addresses/${params.address}`, {
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
export let queryAssetInfo = async (params) => {
  try {
    let { code, data, msg } = await get(`/assets/${params.denom}`, {
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
 *
 * @param {*} params
 * @returns
 */
export let queryAssetsByAddress = async (params) => {
  try {
    let { code, data, msg } = await get(`/addresses/${params.address}`, {
      coin: "",
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
 *
 * @param {*} params
 * @returns
 */
export let transferOutAssets = async function(params) {
  try {
    let { code, data, msg } = await post(`/tx/transfer`, {
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

/**
 * @param {*} params
 * @returns
 */
export let publishAssets = async function(params) {
  try {
    let { code, data, msg } = await post(`/tx/mint`, {
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

/**
 * @param {*} params
 * @returns
 */
export let burnAssetsByAdd = async function(params) {
  try {
    let { code, data, msg } = await post(`/tx/burn`, {
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

/**
 * @param {*} params
 * @returns
 */
export let queryAssetsListPage = async function(params) {
  try {
    let { code, data, msg } = await get(`/addresses/${params.address}/assets`, {
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
