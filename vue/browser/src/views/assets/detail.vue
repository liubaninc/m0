<!-- 资产详情 -->
<template>
  <div class="assetsDetail width1200">
    <div class="breadcrumb">
      <el-breadcrumb separator="/">
        <el-breadcrumb-item :to="{ path: '/index' }">首页</el-breadcrumb-item>
        <el-breadcrumb-item :to="{ path: '/assets' }">资产</el-breadcrumb-item>
        <el-breadcrumb-item>资产详情</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <div class="Detail-message">
      <div class="title">资产别名 / {{ name }}</div>
      <el-row>
        <el-col :span="12">
          <div class="grid-content bg-purple-light">
            <span class="label">资产ID</span>
            <el-tooltip class="item" effect="dark" placement="top">
              <div slot="content">{{ detailData.ref_tx }}</div>
              <span class="oneLine">{{ detailData.ref_tx }}</span>
            </el-tooltip>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="grid-content bg-purple-light">
            <span class="label">资产别名</span>
            <span>{{ detailData.denom }}</span>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="grid-content">
            <span class="label">资产发行总量</span>
            <span>{{ detailData.mint_amount + detailData.burn_amount }}</span>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="grid-content">
            <span class="label">资产可用总量</span>
            <span>{{ detailData.mint_amount }}</span>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="grid-content bg-purple-light">
            <span class="label">资产发行时间</span>
            <span>{{ detailData.time }}</span>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="grid-content bg-purple-light">
            <span class="label">资产发行人</span>
            <span class="link oneLine" @click="toDetail(fxName)">{{
              fxName
            }}</span>
            <i
              class="el-icon-document-copy link ml30"
              v-clipboard:copy="fxName"
              v-clipboard:success="onCopy"
              v-clipboard:error="onError"
            ></i>
          </div>
        </el-col>
      </el-row>
    </div>
    <div class="table-tlt">
      <i class="el-icon-s-operation"></i>资产关联地址列表
    </div>
    <div class="Detail-table">
      <el-table
        :data="tableData"
        :stripe="true"
        :border="true"
        :header-cell-style="{
          background: '#f0f0f0',
          color: 'rgba(0, 0, 0, 0.85)',
          fontWeight: 500,
          textAlign: 'center',
        }"
      >
        <el-table-column align="center" label="用户地址">
          <template slot-scope="scope">
            <el-link type="primary" @click="toDetail(scope.row.address)">{{
              scope.row.address
            }}</el-link>
          </template>
        </el-table-column>
        <el-table-column align="center" label="资产金额" prop="">
          <template slot-scope="scope">
            <span>{{ scope.row.coins[0].amount }}</span>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        class="pagination"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="page.currentPage"
        :page-sizes="[10, 20, 30, 40]"
        :page-size="page.pageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :total="page.total"
        background
      >
      </el-pagination>
    </div>
  </div>
</template>

<script>
import { assetsDetail, detailList } from "@/server/assets/index.js";
export default {
  methods: {},
  data() {
    return {
      name: this.$route.query.name,
      detailData: {},
      fxName: "x234dfgg", // 资产发行人
      tableData: [],
      page: {
        total: 0, // 总页数
        currentPage: 1, // 当前页数
        pageSize: 10, // 每页显示多少条,
      },
    };
  },
  created() {
    this.getDetail(this.name);
    this.getList(this.page);
  },
  methods: {
    // 切换第几页
    handleCurrentChange: function(page) {
      this.page.currentPage = page;
      this.getList(this.page);
    },
    // 每页几条
    handleSizeChange: function(sizeNum) {
      this.page.pageSize = sizeNum;
      this.getList(this.page);
    },
    // 详情
    getDetail(name) {
      let url = "/assets/" + name;
      assetsDetail(url).then((res) => {
        this.detailData = res;
        this.fxName = res.initiator;
      });
    },
    // 分页获取数据-资产关联地址列表
    getList(page, params = {}) {
      let url = `/addresses`;
      let paramsData = Object.assign(
        {
          page_num: page.currentPage,
          page_size: page.pageSize,
          coin: this.name, // 资产别名
        },
        params
      );
      detailList(url, paramsData).then((res) => {
        this.tableData = res.addresses;
        this.page.total = res.total;
      });
    },
    // 资产地址
    toDetail(url) {
      this.$router.push({
        path: "/assets/address",
        query: {
          name: this.name,
          url: url,
        },
      });
    },
    // 复制
    onCopy: function(e) {
      console.log("You just copied: " + e.text);
    },
    onError: function(e) {
      console.log("Failed to copy texts");
    },
  },
};
</script>
<style scoped>
.table-tlt {
  margin: 20px 0;
  font-size: 16px;
}
.table-tlt i {
  color: #409eff;
  font-weight: 700;
  margin-right: 10px;
}
.el-icon-document-copy {
  cursor: pointer;
}
</style>
