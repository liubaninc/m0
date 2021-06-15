// @ts-nocheck
import Vue from "vue";
import ElementUI from "element-ui";
import "element-ui/lib/theme-chalk/index.css";
import VueClipboard from "vue-clipboard2";
import filters from "@/utils/filters";
import JsonViewer from "vue-json-viewer";
import App from "./App.vue";
import router from "./router";
Vue.config.productionTip = false;
/**注册全局过滤器 */
Object.keys(filters).forEach((key) => {
  Vue.filter(key, filters[key]);
});

Vue.use(ElementUI);
Vue.use(VueClipboard);
Vue.use(JsonViewer);

new Vue({
  render: (h) => h(App),
  router,
}).$mount("#app");
