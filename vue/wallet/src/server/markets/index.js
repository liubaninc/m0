import { post, get } from "@/utils/request";


import Vue from 'vue'
/**
 * 查询合约模板市场
 * @param {*} params 
 * @returns 
 */
export let queryContractMarket = async (params) => {
    try {
        let { code, data, msg } = await post(`mcontract/template/list`, {
            ...params
        });
        if (code == 200) {
            return data;
        } else {
            Vue.prototype.$message.error(msg)
            console.log(code, msg);
        }
    } catch (error) {
        console.log(error);
    }
};
/**
 * 查询合约详情
 * @param {*} params 
 * @returns 
 */
export let queryContractInfo = async (params) => {
    try {
        let { code, data, msg } = await get(`/mcontract/template/get/${params.id}`, {
            ...params
        });
        if (code == 200) {
            return data;
        } else {
            Vue.prototype.$message.error(msg)
            console.log(code, msg);
        }
    } catch (error) {
        console.log(error);
    }
};
