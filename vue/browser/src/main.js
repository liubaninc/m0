import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import VueClipboard from "vue-clipboard2";
import ElementUI from "element-ui";
import "element-ui/lib/theme-chalk/index.css";
import filters from "@/utils/filters";
import JsonViewer from "vue-json-viewer";
Vue.use(ElementUI);
Vue.use(VueClipboard);
Vue.use(JsonViewer);
Vue.config.productionTip = false;

/**注册全局过滤器 */
Object.keys(filters).forEach((key) => {
  Vue.filter(key, filters[key]);
});

new Vue({
  render: (h) => h(App),
  router,
}).$mount("#app");
