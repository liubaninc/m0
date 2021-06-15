<!-- layout -->
<template>
  <div class="layout">
    <el-container>
      <el-header class="layout-header">
        <Header />
      </el-header>
      <el-main class="layout-main">
        <div class="content">
          <router-view></router-view>
        </div>
      </el-main>
      <el-footer
        >©2020 TimeFlow All Rights Reserved 京ICP备13017353号-3</el-footer
      >
    </el-container>
  </div>
</template>

<script>
import Header from "../header/header.vue";
export default {
  name: "Layout",
  components: { Header },
  data() {
    return {};
  },
  methods: {},
};
</script>
<style scoped>
.el-header {
  padding: 0;
  background-color: #242424;
}
.el-footer {
  position: fixed;
  width: 100%;
  bottom: 0;
  text-align: center;
  background: #242424;
  color: #ffffff;
  line-height: 60px;
  font-size: 16px;
}
.layout-main {
  height: calc(100vh - 120px);
  overflow-y: auto;
}
.content {
  margin: 0 auto;
}
</style>
