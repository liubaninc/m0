<!-- layout -->
<template>
  <div class="layout">
    <el-container>
      <el-header class="layout-header">
        <Header />
      </el-header>
      <el-container class="layout-main">
        <el-aside class="left-side" v-if="isAside">
          <LeftSide></LeftSide>
        </el-aside>
        <el-main>
          <div class="content">
            <router-view></router-view>
          </div>
        </el-main>
      </el-container>
      <el-footer>京ICP备19054130号 北京磁云唐泉金服科技有限公司</el-footer>
    </el-container>
  </div>
</template>

<script>
import Header from '@/components/header/header.vue'
import LeftSide from '@/components/leftSide'
export default {
  name: 'Layout',
  components: { Header, LeftSide },
  data() {
    return {
      isAside: true,
    }
  },
  watch: {
    $route(to, from) {
      if (/^\/(wallet)|(wallet\/)|\/(download)|(download\/)/.test(to.path)) {
        this.isAside = false
      } else {
        this.isAside = true
      }
    },
  },
  created() {
    let hash = window.location.hash
    if (/^#\/(wallet)|(wallet\/)|\/(download)|(download\/)/.test(hash)) {
      this.isAside = false
    } else {
      this.isAside = true
    }
  },
  methods: {},
}
</script>
<style scoped>
.layout {
  background: #1b2c42;
  color: #fff;
}
.el-header {
  padding: 0;
  background: #2a405c !important;
  height: 54px !important;
  display: flex;
  justify-items: center;
  align-items: center;
}
.el-footer {
  width: 100%;
  text-align: center;
  background: #1f2525;
  color: #b7b7b7;
  line-height: 43px;
  height: 43px !important;
  font-size: 12px;
}
.layout-main {
  min-height: calc(100vh - 97px);
  height: calc(100vh - 97px);
  /* overflow-y: auto; */
  overflow-y: hidden;
}

.left-side {
  width: 200px !important;
}
.content {
  margin: 0 auto;
  padding: 0px 40px 20px;
}

.el-main {
  /* padding: 20px 40px; */
  padding: 0;
}
</style>
