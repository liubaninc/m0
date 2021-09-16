// @ts-nocheck
import Vue from "vue";
import VueRouter from "vue-router";
import { routes } from "./routes";
Vue.use(VueRouter);
import { localCache } from "@/utils/utils";
import { queryWalletLists } from "@/server/wallet";

import NProgress from "nprogress";
import "nprogress/nprogress.css";

const router = new VueRouter({
  mode: "hash",
  linkActiveClass: "Aactive",
  routes: [...routes],
  scrollBehavior (to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition;
    } else {
      return { x: 0, y: 0 };
    }
  },
});

NProgress.configure({
  easing: "ease",
  speed: 500,
  trickleRate: 0.02,
  trickleSpeed: 500,
  showSpinner: true,
});

NProgress.inc(0.2)

router.beforeEach(async (to, from, next) => {
  NProgress.start();
  let { meta } = to;
  let authorization = localCache.get("authorization");
  if (meta && meta.isNeedLogin) {
    if (authorization) {
      let { invalidToken } = await queryWalletLists({ page_size: 1 });
      if (invalidToken) {
        return next("/login");
      }
      next();
    } else {
      next("/login");
    }
  } else {
    next();
  }
});

router.afterEach(() => {
  NProgress.done();
});

export default router;
