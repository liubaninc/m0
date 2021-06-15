<template>
  <div class="commonHeader">
    <div class="headerBox">
      <div class="left">
        <div class="log" @click="$router.push('/index')">
          <img src="../../assets/logo/logo.png" alt="" />
          <!-- <span>AK EXPLORER</span> -->
        </div>
      </div>
      <div class="nav">
        <!-- <router-link v-for="(item, index) in navLists" :key="index" :to="item.path">{{item.text}}</router-link> -->
        <el-menu
          :default-active="$route.path"
          exact
          class="el-menu-demo"
          mode="horizontal"
          @select="handleSelect"
          background-color="#242424"
          text-color="#fff"
          active-text-color="#ffd04b"
          router
        >
          <el-menu-item index="/index">首页</el-menu-item>
          <el-submenu index="/block">
            <template slot="title">区块</template>
            <el-menu-item index="/block/list">区块</el-menu-item>
            <el-menu-item index="/block/tradeList">交易</el-menu-item>
          </el-submenu>
          <el-menu-item index="/assets">资产</el-menu-item>
          <el-menu-item index="/contract">合约</el-menu-item>
          <el-menu-item index="/node">节点</el-menu-item>
        </el-menu>
      </div>
      <div class="right">
        <el-input placeholder="区块/交易/地址/合约/资产" v-model="searchInput">
          <template slot="append">
            <el-button
              type="primary"
              slot="append"
              icon="el-icon-search"
              @click="search"
            ></el-button>
          </template>
        </el-input>
      </div>
    </div>
  </div>
</template>

<script>
import { searchConetnt } from "@server/search";
import { addressDetail } from "@server/assets/address";
export default {
  data() {
    return {
      searchInput: "",
    };
  },
  methods: {
    handleSelect(key, keyPath) {
      console.log(key, keyPath);
    },
    async search() {
      if (!this.searchInput)
        return this.$message({ message: "请输入搜索内容" });
      let searchResult = await searchConetnt(`/search/${this.searchInput}`);
      if (searchResult) {
        let [type, id] = searchResult.split("/");
        let url = "/";
        switch (type) {
          case "blocks":
            url = `/block/info`;
            break;
          case "transactions":
            url = `/block/tradeInfo`;
            break;
          case "assets":
            url = `/assets/detail`;
            break;
          case "contracts":
            url = `/contract/detail`;
            break;
          case "addresses":
            let { coins } = await this.queryAssestAddress(`/addresses/${id}`);
            if (coins && coins.length) {
              url = `/assets/address?name=${coins[0].denom}&url=${id}`;
            } else {
              return this.$message.error({ message: "地址不存在" });
            }
            break;
        }
        if (type != "addresses") {
          this.$router.push({
            path: url,
            query: {
              name: id,
            },
          });
        } else {
          this.$router.push({
            path: url,
          });
        }
      }
    },
    async queryAssestAddress(url, params = {}) {
      let assetsDetail = await addressDetail(url, params);
      return assetsDetail;
    },
  },
};
</script>
<style scoped>
.commonHeader {
  padding: 0;
  width: 1200px;
  margin: 0 auto;
}
.headerBox {
  display: flex;
  align-items: center;
  align-content: center;
}
.left {
  width: 200px;
}
.left .log {
  cursor: pointer;
}
.left .log img {
  width: 120px;
}
.nav {
  flex-grow: 1;
}
.nav .el-menu--horizontal > .el-menu-item {
  padding: 0 30px;
}
.right {
  width: 360px;
}
.el-input-group__append .el-button--primary {
  color: #fff;
  background-color: #409eff;
  border-color: #409eff;
}
.el-menu.el-menu--horizontal {
  border-bottom: none;
}
</style>
