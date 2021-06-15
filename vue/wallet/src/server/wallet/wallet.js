// @ts-nocheck
import { post, get } from "@/utils/request";

const PAGE_SIZE = 20;
/**
 *
 * @param {*} params
 * @returns
 */
export let queryWalletLists = async (params) => {
  try {
    let { code, data, msg } = await post("accounts", {
      page_num: 1,
      page_size: PAGE_SIZE,
      ...params,
    });
    if (code == 200) {
      return data;
    } else if (code == 3000) {
      return {
        invalidToken: !0,
      };
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
export let genMnemonic = async (params) => {
  try {
    let { code, data, msg } = await post("account/mnemonic", {
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
export let createWallet = async function(params) {
  try {
    let { code, data, msg } = await post("account/create", {
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
 *
 * @param {*} params
 * @returns
 */
export let createMoreSignWallet = async function(params) {
  try {
    let { code, data, msg } = await post("account/create_multisig", {
      sort: true,
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
export let queryWalletInfo = async (params) => {
  try {
    let { code, data, msg } = await post(
      `/accounts/${encodeURI(params.name)}`,
      {
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
