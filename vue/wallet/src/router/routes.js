import Layout from "@/components/layout/index.vue";
import Viewport from "@/components/viewport/viewport.vue";
export let routes = [
  {
    path: "/",
    name: "index",
    component: Layout,
    redirect: "/wallet",
    children: [
      {
        path: "/dapps",
        component: Viewport,
        children: [
          {
            path: "evidence",
            component: Viewport,
            children: [
              {
                path: "/",
                name: "存证列表",
                meta: {
                  isNeedLogin: !0,
                },
                component: () =>
                  import(
                    /* webpackChunkName: "dapps/evidence/index" */ "@/views/dapps/evidence"
                  ),
              },
              {
                path: "detail",
                name: "存证详情",
                meta: {
                  isNeedLogin: !0,
                },
                component: () =>
                  import(
                    /* webpackChunkName: "dapps/evidence/detail" */ "@/views/dapps/evidence/detail"
                  ),
              },
              {
                path: "singleEvidence",
                name: "单签存证",
                meta: {
                  isNeedLogin: !0,
                },
                component: () =>
                  import(
                    /* webpackChunkName: "dapps/evidence/singleEvidence" */ "@/views/dapps/evidence/singleEvidence"
                  ),
              },
              {
                path: "moreEvidence",
                name: "多签存证",
                meta: {
                  isNeedLogin: !0,
                },
                component: () =>
                  import(
                    /* webpackChunkName: "dapps/evidence/moreEvidence" */ "@/views/dapps/evidence/moreEvidence"
                  ),
              },
              {
                path: "evidSuccess",
                name: "存证成功",
                meta: {
                  isNeedLogin: !0,
                },
                component: () =>
                  import(
                    /* webpackChunkName: "dapps/evidence/evidSuccess" */ "@/views/dapps/evidence/evidSuccess"
                  ),
              },
            ],
          },
        ],
      },
      {
        path: "walService",
        component: Viewport,
        children: [
          {
            path: "/",
            name: "单签钱包详情",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(
                /* webpackChunkName: "walService/singleDetail" */ "@/views/walService/singleDetail"
              ),
          },
          {
            path: "moreDetail",
            name: "单签钱包详情",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(
                /* webpackChunkName: "walService/moreDetail" */ "@/views/walService/moreDetail"
              ),
          },
          {
            path: "keyBackup",
            name: "秘钥备份",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(
                /* webpackChunkName: "walService/keyBackup" */ "@/views/walService/keyBackup"
              ),
          },
          {
            path: "backupInfo",
            name: "秘钥备份信息",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(
                /* webpackChunkName: "walService/backupInfo" */ "@/views/walService/backupInfo"
              ),
          },
        ],
      },
      {
        path: "/trx",
        name: "交易",
        component: Viewport,
        children: [
          {
            path: "/",
            name: "交易列表",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(
                /* webpackChunkName: "trx/index" */ "@/views/transactions"
              ),
          },
          {
            path: "detail",
            name: "交易详情",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(
                /* webpackChunkName: "trx/detail" */ "@/views/transactions/detail"
              ),
          },
        ],
      },
      {
        path: "/assets",
        name: "资产",
        component: Viewport,
        children: [
          {
            path: "/",
            name: "资产列表",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(/* webpackChunkName: "assets/index" */ "@/views/assets"),
          },
          {
            path: "detail",
            name: "资产详情",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(
                /* webpackChunkName: "assets/detail" */ "@/views/assets/detail"
              ),
          },
          {
            path: "receive",
            name: "转账接收",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(
                /* webpackChunkName: "assets/receive" */ "@/views/assets/receive"
              ),
          },
          {
            path: "transferOut",
            name: "单签资产转出",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(
                /* webpackChunkName: "assets/transferOut" */ "@/views/assets/transferout/transferOut"
              ),
          },
          {
            path: "transferOutMore",
            name: "多签资产转出",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(
                /* webpackChunkName: "assets/transferOutMore" */ "@/views/assets/transferout/transferOutMore"
              ),
          },
          {
            path: "transferOutSuccess",
            name: "资产转出成功",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(
                /* webpackChunkName: "assets/transferOutSuccess" */ "@/views/assets/transferout/transferOutSuccess"
              ),
          },
          {
            path: "transferOutIng",
            name: "资产转出进行中",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(
                /* webpackChunkName: "assets/transferOutIng" */ "@/views/assets/transferout/transferOutIng"
              ),
          },
          {
            path: "publishAsset",
            name: "发行资产",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(
                /* webpackChunkName: "assets/publishAsset" */ "@/views/assets/publish"
              ),
          },
          {
            path: "publishMore",
            name: "发行多资产",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(
                /* webpackChunkName: "assets/publishMore" */ "@/views/assets/publish/publishMore"
              ),
          },
          {
            path: "publicIng",
            name: "发行多资产已提交",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(
                /* webpackChunkName: "assets/publicIng" */ "@/views/assets/publish/publicIng"
              ),
          },
          {
            path: "publicSuccess",
            name: "发行多资产成功",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(
                /* webpackChunkName: "assets/publicSuccess" */ "@/views/assets/publish/publicSuccess"
              ),
          },
          {
            path: "addSessets",
            name: "增发资产",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(
                /* webpackChunkName: "increase/addSessets" */ "@/views/assets/increase"
              ),
          },
          {
            path: "addAssetsMore",
            name: "增发资产多签",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(
                /* webpackChunkName: "increase/addAssetsMore" */ "@/views/assets/increase/addAssetsMore"
              ),
          },
          {
            path: "destoryAssets",
            name: "销毁资产",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(
                /* webpackChunkName: "destory/destoryAssets" */ "@/views/assets/destory"
              ),
          },
          {
            path: "destoryAssetsMore",
            name: "多签销毁资产",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(
                /* webpackChunkName: "destory/destoryAssetsMore" */ "@/views/assets/destory/destoryAssetsMore"
              ),
          },
          {
            path: "destoryIng",
            name: "销毁资产已提交",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(
                /* webpackChunkName: "destory/destoryIng" */ "@/views/assets/destory/destoryIng"
              ),
          },
          {
            path: "destorySuccess",
            name: "销毁资产成功",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(
                /* webpackChunkName: "destory/destorySuccess" */ "@/views/assets/destory/destorySuccess"
              ),
          },
          // {
          //   path: "signMore",
          //   name: "多签交易签名",
          //   meta: {
          //     isNeedLogin: !0,
          //   },
          //   component: () =>
          //     import(
          //       /* webpackChunkName: "destory/signMore" */ "@/views/assets/signature/index2"
          //     ),
          // },
          {
            path: "signature",
            name: "多签交易签名",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(
                /* webpackChunkName: "destory/signature" */ "@/views/assets/signature"
              ),
          },
        ],
      },
      {
        path: "/download",
        component: Viewport,
        children: [
          {
            path: "",
            name: "下载中心",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(/* webpackChunkName: "download/index" */ "@/views/download"),
          },
        ]
      },
      {
        path: "/wallet",
        component: Viewport,
        children: [
          {
            path: "",
            name: "钱包列表",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(/* webpackChunkName: "wallet/index" */ "@/views/wallet"),
          },
          {
            path: "createTypes",
            name: "钱包创建类型选择",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(
                /* webpackChunkName: "wallet/createTypes" */ "@/views/wallet/createTypes"
              ),
          },
          {
            path: "importTypes",
            name: "钱包导入类型选择",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(
                /* webpackChunkName: "wallet/walTypes" */ "@/views/wallet/importTypes"
              ),
          },
          {
            path: "walSingleCreate",
            name: "单签钱包创建",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(
                /* webpackChunkName: "wallet/walSingleCreate" */ "@/views/wallet/walSingleCreate"
              ),
          },
          {
            path: "walMoreCreate",
            name: "多签钱包创建",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(
                /* webpackChunkName: "wallet/walMoreCreate" */ "@/views/wallet/walMoreCreate"
              ),
          },
          {
            path: "walCreateSuccess",
            name: "walCreateSuccess",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(
                /* webpackChunkName: "wallet/walCreateSuccess" */ "@/views/wallet/walCreateSuccess"
              ),
          },
          {
            path: "walWordImport",
            name: "助记词导入",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(
                /* webpackChunkName: "wallet/walWordImport" */ "@/views/wallet/walWordImport"
              ),
          },
          {
            path: "walPrivateImport",
            name: "私钥导入",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(
                /* webpackChunkName: "wallet/walPrivateImport" */ "@/views/wallet/walPrivateImport"
              ),
          },
        ],
      },
      {
        path: "/ctractMarket",
        component: Viewport,
        children: [
          {
            path: "",
            name: "合约市场",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(/* webpackChunkName: "ctractMarket/contractTemp" */ "@/views/ctractMarket/contractTemp"),
          },
          {
            path: "useTemp",
            name: "使用合约模板",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(/* webpackChunkName: "ctractMarket/useTemp" */ "@/views/ctractMarket/contractTemp/useTemp"),
          },
        ],
      },

      {
        path: "/mycontract",
        component: Viewport,
        children: [
          {
            path: "",
            name: "我的合约",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(/* webpackChunkName: "mycontract/index" */ "@/views/myContract"),
          },
          {
            path: "detail",
            name: "合约详情",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(/* webpackChunkName: "mycontract/detail" */ "@/views/myContract/detail"),
          },
          {
            path: "create",
            name: "合约创建",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(/* webpackChunkName: "mycontract/create" */ "@/views/myContract/create"),
          },
          {
            path: "signature",
            name: "多签合约签名",
            meta: {
              isNeedLogin: !0,
            },
            component: Viewport,
            children: [
              {
                path: "",
                name: "多签合约签名",
                meta: {
                  isNeedLogin: !0,
                },
                component: () =>
                  import(/* webpackChunkName: "mycontract/signature" */ "@/views/myContract/signature"),
              },
              {
                path: "opSuccess",
                name: "多签合约签名成功",
                meta: {
                  isNeedLogin: !0,
                },
                component: () =>
                  import(/* webpackChunkName: "mycontract/opSuccess" */ "@/views/myContract/signature/opSuccess"),
              },
            ]
          },
          {
            path: "deploy",
            name: "部署模块",
            component: Viewport,
            children: [
              {
                path: "",
                name: "部署合约",
                meta: {
                  isNeedLogin: !0,
                },
                component: () =>
                  import(/* webpackChunkName: "mycontract/deploy" */ "@/views/myContract/deploy"),
              },
              {
                path: "moreDeploy",
                name: "多签部署合约",
                meta: {
                  isNeedLogin: !0,
                },
                component: () =>
                  import(/* webpackChunkName: "mycontract/moreDeploy" */ "@/views/myContract/deploy/moreDeploy"),
              },
              {
                path: "singleDeploy",
                name: "单签部署合约",
                meta: {
                  isNeedLogin: !0,
                },
                component: () =>
                  import(/* webpackChunkName: "mycontract/singleDeploy" */ "@/views/myContract/deploy/singleDeploy"),
              },
            ]

          },
          {
            path: "upgrade",
            name: "合约升级",
            meta: {
              isNeedLogin: !0,
            },
            component: () =>
              import(/* webpackChunkName: "mycontract/upgrade" */ "@/views/myContract/upgrade"),
          },
        ],
      }
    ],
  },

  {
    path: "/login",
    name: "登录",
    component: () =>
      import(/* webpackChunkName: "login/index" */ "@/views/login/index"),
  },
  {
    path: "/register",
    name: "注册",
    component: () =>
      import(/* webpackChunkName: "login/register" */ "@/views/login/register"),
  },
  {
    path: "/regSuccess",
    name: "注册成功",
    component: () =>
      import(
        /* webpackChunkName: "login/regSuccess" */ "@/views/login/regSuccess"
      ),
  },
];
