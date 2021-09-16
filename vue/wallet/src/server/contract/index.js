import { post, get } from "@/utils/request";
import Vue from 'vue'


/**
 * 查询合约列表
 * @param {*} params 
 * @returns 
 */
export let queryContractLists = async (params) => {
    try {
        let { code, data, msg } = await post(`mcontract/list/${params.account}`, {
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
 * 合约详情
 * @param {*} id 
 * @returns 
 */
export let queryContractDetail = async (params) => {
    try {
        let { code, data, msg } = await get(`mcontract/get/${params.id}`, {
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
}

/**
 * 查询合约版本列表
 * @param {*} params 
 * @returns 
 */
export let queryContractVersions = async (params) => {
    try {
        let { code, data, msg } = await get(`mcontract/history/list/${params.contractName}`, {
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
}

/**
 * 查询某个合约下的交易列表
 * @param {*} params 
 * @returns 
 */
export let queryContractTrxByName = async (params) => {
    try {
        let { code, data, msg } = await get(`contracts/${params.name}/transactions`, {
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
}

/**
 * 查询交易信息
 * @param {*} params 
 * @returns 
 */
export let queryContractTrxByHash = async (params) => {
    try {
        let { code, data, msg } = await get(`mcontract/transactions/${params.hash}`, {
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
}
/**
 * @param {*} params 
 * @returns 
 */
export let signContract = async (params) => {
    try {
        let { code, data, msg } = await post(`mcontract/tx/sign`, {
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
 *  创建合约
 * @param {*} params 
 * @returns 
 */
export let createContract = async (params) => {
    try {
        let fromData = new FormData();
        params &&
            Object.keys(params).forEach((key) => {
                fromData.append(key, params[key]);
            });
        let { code, data, msg } = await post(`mcontract/create`, fromData, {
            "Content-Type": "application/x-www-form-urlencoded",
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
 * undeploy删除合约 deploy部署合约 upgrade升级合约 freeze冻结合约 unfreeze解冻合约
 * @param {*} params 
 * @returns 
 */
export let contractOperate = async (params) => {
    try {
        let fromData = new FormData();
        params &&
            Object.keys(params).forEach((key) => {
                fromData.append(key, params[key]);
            });
        let { code, data, msg } = await post(`mcontract/operate`, fromData, {
            "Content-Type": "application/x-www-form-urlencoded",
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
 * 合约部署
 * @param {*} params 
 * @returns 
 */
export let publishContract = async (params) => {
    try {
        let fromData = new FormData();
        params &&
            Object.keys(params).forEach((key) => {
                fromData.append(key, params[key]);
            });
        let { code, data, msg } = await post(`mcontract/operate`, fromData, {
            "Content-Type": "application/x-www-form-urlencoded",
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
 * 
 * @param {*} params 
 * @returns 
 */
export let delContract = async (params) => {
    try {
        let { code, data, msg } = await post(`mcontract/delete/${params.id}`, {
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

