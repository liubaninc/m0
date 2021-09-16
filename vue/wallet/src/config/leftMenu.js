/**
 * 左侧菜单
 */
export let leftMenus = [
  {
    id: 100001,
    text: "核心数据",
    icon: "iconfont m0-shuju",
    url: "",
    children: [
      {
        id: "100001_1",
        text: "资产管理",
        icon: "iconfont m0-zichan",
        children: [],
        url: "/assets",
        preKey: "core_",
      },
      {
        id: "100001_2",
        text: "交易记录",
        icon: "iconfont m0-biaoqiankuozhan_jiaoyi-179",
        children: [],
        url: "/trx",
        preKey: "core_",
      },
      {
        id: "100001_3",
        text: "我的合约",
        icon: "iconfont m0-icon-",
        children: [],
        url: "/mycontract",
        preKey: "core_",
      },
    ],
  },
  {
    id: 100002,
    text: "钱包服务",
    icon: "iconfont m0-fuwu",
    url: "",
    children: [
      {
        id: "100002_1",
        text: "钱包服务",
        icon: "iconfont m0-dilanxianxingiconyihuifu_huabanfuben",
        children: [],
        url: "/walService",
        preKey: "wal_",
      },
      {
        id: "100002_2",
        text: "秘钥备份",
        icon: "iconfont m0-miyao",
        children: [],
        url: "/walService/keyBackup",
        preKey: "wal_",
      },
    ],
  },
  {
    id: 100003,
    text: "DAPPS",
    icon: "iconfont m0-caidan",
    url: "",
    children: [
      {
        id: "100003_1",
        text: "存证服务",
        icon: "iconfont m0-fuwu1",
        children: [],
        url: "/dapps/evidence",
        preKey: "dp_",
      },
    ],
  },
  {
    id: 100004,
    text: "合约市场",
    icon: "iconfont m0-shichang",
    url: "",
    children: [
      {
        id: "100004_1",
        text: "合约市场",
        icon: "iconfont m0-xiaoxi",
        children: [],
        url: "/ctractMarket",
        preKey: "contract_",
      },
    ],
  },
];
