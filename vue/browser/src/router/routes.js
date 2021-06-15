import Layout from "@/components/layout/index.vue";
import Viewport from "@/components/viewport/viewport.vue";
export let routes = [
  {
    path: "/",
    name: "index",
    component: Layout,
    redirect: "/index",
    children: [
      {
        path: "/index",
        name: "首页",
        component: () => import("@/views/index/index"),
      },
      {
        // path: "/block",
        path: "/block",
        component: Viewport,
        name: "区块",
        children: [
          {
            path: "/block/list",
            name: "区块列表",
            component: () => import("@/views/block/list"),
          },
          {
            path: "/block/info",
            name: "区块详情",
            component: () => import("@/views/block/info"),
          },
          {
            path: "/block/tradeInfo",
            name: "交易详情",
            component: () => import("@/views/block/tradeInfo"),
          },
          {
            path: "/block/tradeList",
            name: "交易列表",
            component: () => import("@/views/block/tradeList"),
          },
        ],
      },
      {
        path: "/assets",
        name: "资产",
        component: () => import("@/views/assets/index"),
      },
      {
        path: "/assets/detail",
        name: "资产详情",
        component: () => import("@/views/assets/detail"),
      },
      {
        path: "/assets/address",
        name: "资产地址",
        component: () => import("@/views/assets/address"),
      },
      {
        path: "/contract",
        name: "合约",
        component: () => import("@/views/contract/index"),
      },
      {
        path: "/contract/detail",
        name: "合约详情",
        component: () => import("@/views/contract/detail"),
      },
      {
        path: "/node",
        name: "节点",
        component: () => import("@/views/node/index"),
      },
    ],
  },
  // {
  // 	path:'',
  // 	redirect: home
  // }
];
