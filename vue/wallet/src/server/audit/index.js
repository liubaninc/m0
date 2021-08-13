// @ts-nocheck
import { post, get } from "@/utils/request";
import Vue from 'vue';
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
        } else if (code == 201) {
            // window.location.href = location.origin + "#/audit/noAuth"
            return {
                noAuth: true,
                ...data
            }

        } else {
            Vue.prototype.$message.error(`${msg}`)
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
export let queryCheckPerson = async (params) => {
    try {
        let { code, data, msg } = await get(`/events/validator`, {
            ...params,
        });
        if (code == 200) {
            return data;
        } else if (code == 201) {
            // window.location.href = location.origin + "#/audit/noAuth"
            return {
                noAuth: true,
                ...data
            }
        } else {
            Vue.prototype.$message.error(`${msg}`)
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
export let queryCerts = async (params) => {
    try {
        let { code, data, msg } = await get(`/events/cert`, {
            ...params,
        });
        if (code == 200) {
            return data;
        } else if (code == 201) {
            // window.location.href = location.origin + "#/audit/noAuth"
            return {
                noAuth: true,
                ...data
            }
        } else {
            Vue.prototype.$message.error(`${msg}`)
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
export let queryUsers = async (params) => {
    try {
        let { code, data, msg } = await get(`/events/account`, {
            ...params,
        });
        if (code == 200) {
            return data;
        } else if (code == 201) {
            // window.location.href = location.origin + "#/audit/noAuth"
            return {
                noAuth: true,
                ...data
            }
        } else {
            Vue.prototype.$message.error(`${msg}`)
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
export let queryContracts = async (params) => {
    try {
        let { code, data, msg } = await get(`/events/contract`, {
            ...params,
        });
        if (code == 200) {
            return data;
        } else if (code == 201) {
            // window.location.href = location.origin + "#/audit/noAuth"
            return {
                noAuth: true,
                ...data
            }
        } else {
            Vue.prototype.$message.error(`${msg}`)
            console.log(code, msg);
        }
    } catch (error) {
        console.log(error);
    }
};
