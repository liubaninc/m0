// @ts-nocheck
import { post, get } from "@/utils/request";
/**
 * 用户注册
 * @param {*} params
 * @returns
 */
export async function registerUser(params) {
  try {
    let { code, data, msg } = await post("user/register", {
      name: "",
      password: "",
      nick: "",
      email: "",
      mobile: "",
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
}

/**
 * 用户登录
 * @param {*} params
 * @returns
 */
export async function loginUser(params) {
  try {
    let { code, data, msg } = await post("user/login", {
      name: "ttt",
      password: "ttt",
      exp_duration: 0,
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
}

/**
 *
 * @param {*} params
 */
export let getImgCode = async (params = {}) => {
  try {
    let { captchaId, imageUrl } = await get("captcha", {
      ...params,
    });
    imageUrl =
      (process.env.NODE_ENV == "development"
        ? process.env.VUE_APP_DEV_BASE_URL
        : process.env.VUE_APP_PRO_BASE_URL) + imageUrl;
    if (captchaId && imageUrl) {
      return {
        imageUrl,
        captchaId,
      };
    }
  } catch (error) {
    console.log(error);
  }
};

/**
 * 用户退出
 * @param {*} params
 * @returns
 */
export let layoutUser = async (params) => {
  try {
    let { code, data, msg } = await post("user/logout", {
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
