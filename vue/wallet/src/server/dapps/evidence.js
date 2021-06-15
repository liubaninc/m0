import { post, get } from "@/utils/request";

/**
 * @param {*} params
 * @returns
 */
export let queryEvidenceList = async (params) => {
  try {
    let { code, data, msg } = await post(`/claims/${params.account}`, {
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
export let saveEvidence = async function(params) {
  try {
    let { code, data, msg } = await post(`/claims/${params.account}/tx`, {
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
export let queryEvidenceDetail = async function(params) {
  try {
    let { code, data, msg } = await post(`/claims/${params.account}/get`, {
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
export let verifyFile = async function(params) {
  try {
    let { code, data, msg } = await post(`/claims/${params.account}/verify`, {
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
