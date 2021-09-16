
/**
 * 合约的状态
 * @param {*} status 
 * @returns 
 */
export function contractStatusText (status) {
    return ({
        '0': "待部署",
        "1": "部署中",
        "2": "已部署",
        "3": "部署失败",
        "4": "已冻结",
        "5": "已部署",
        "6": "升级中",
        "7": "升级失败",
        "8": "已部署",
        "9": "已删除"
    })[status]
}

/**
 * 按钮显示状态
 * @param {*} mode 
 * @returns 
 */
export function btnStatus (mode) {
    return ({
        'undeploy': "立即删除",
        "deploy": "立即部署",
        "upgrade": "立即升级",
        "freeze": "立即冻结",
        "unfreeze": "立即解冻",
    })[mode] || '立即部署'
}

/**
 * @param {*} mode 
 * @returns 
 */
export function tipStatus (mode) {
    return ({
        'undeploy': "删除中...",
        "deploy": "部署中...",
        "upgrade": "升级中...",
        "freeze": "冻结中...",
        "unfreeze": "解冻中...",
    })[mode] || '部署中...'
}
/**
 * 合约操作
 * @param {*} mode 
 * @returns 
 */
export function txtStatus (mode) {
    return ({
        'undeploy': "删除合约",
        "deploy": "部署合约",
        "upgrade": "升级合约",
        "freeze": "冻结合约",
        "unfreeze": "解冻合约",
    })[mode] || '部署合约'
}




