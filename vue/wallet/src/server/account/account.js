import { post, get } from "@/utils/request";

const PAGE_SIZE = 100;
/**
 * 随机生成一个助记词
 * @param {*} params
 * @returns
 */
export let createMemWords = async (params) => {
  try {
    let { code, data, msg } = await get("account/mnemonic", {
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
 * 创建一个单签钱包账户，支持导入私钥明文(hex)、助记词
 * @param {*} params
 * @returns
 */
export let createSingleAcc = async function(params) {
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
 * 新建多签钱包账户
 * @param {*} params
 * @returns
 */
export let createMultisAcc = async (params) => {
  try {
    let { code, data, msg } = await post("account/create_multisig", {
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
 * 导出一个钱包账户
 * @param {*} params
 * @returns
 */
export let exportAcc = async function(params) {
  try {
    let { code, data, msg } = await post("account/export", {
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
 * 查询所有钱包账户
 * @param {*} params
 * @returns
 */
export let queryAccLists = async (params) => {
  try {
    let { code, data, msg } = await get("accounts", {
      page_size: PAGE_SIZE,
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
