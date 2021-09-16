// @ts-nocheck
import { post, get } from "@/utils/request";

/**
 *
 * @param {*} action page_size page_num
 * @returns
 */
export let queryPeers = async (params) => {
    try {
        let { code, data, msg } = await get(`/events/peer`, {
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
export let publishAssets = async function (params) {
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
