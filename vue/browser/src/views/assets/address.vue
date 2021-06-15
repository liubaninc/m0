<!-- 资产地址 -->
<template>
  <div class="address width1200">
    <div class="breadcrumb">
      <el-breadcrumb separator="/">
        <el-breadcrumb-item :to="{ path: '/index' }">首页</el-breadcrumb-item>
        <el-breadcrumb-item :to="{ path: '/assets' }">资产</el-breadcrumb-item>
        <el-breadcrumb-item>地址资产</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <div class="Detail-message">
      <div class="title">链上地址：{{ url }}</div>
      <el-row>
        <el-col :span="12">
          <div class="grid-content bg-purple-light">
            <span class="label">资产名称</span>
            <el-select
              size="mini"
              v-model="initName"
              placeholder="请选择"
              @change="changeName()"
            >
              <el-option
                v-for="item in options"
                :key="item.denom"
                :label="item.denom"
                :value="item.denom"
              >
              </el-option>
            </el-select>
          </div>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="12">
          <div class="grid-content">
            <span class="label">资产ID</span>
            <span class="link oneLine">{{ detailData.ref_tx }}</span>
            <i
              class="el-icon-document-copy link ml30"
              v-clipboard:copy="detailData.ref_tx"
              v-clipboard:success="onCopy"
              v-clipboard:error="onError"
            ></i>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="grid-content">
            <span class="label">资产发行总量</span>
            <span>{{ detailData.mint_amount + detailData.burn_amount }}</span>
          </div>
        </el-col>
      </el-row>
      <el-row>
        <el-col :span="12">
          <div class="grid-content bg-purple-light">
            <span class="label">交易数量</span>
            <span>{{ detailData.mint_amount }}</span>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="grid-content bg-purple-light">
            <span class="label">资产发行人</span>
            <span class="link oneLine">{{ detailData.initiator }}</span>
            <i
              class="el-icon-document-copy link ml30"
              v-clipboard:copy="detailData.initiator"
              v-clipboard:success="onCopy"
              v-clipboard:error="onError"
            ></i>
          </div>
        </el-col>
      </el-row>
    </div>
    <div class="table-tlt"><i class="el-icon-s-operation"></i>交易列表</div>
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
        <el-table-column align="center" label="交易Hash" width="400">
          <template slot-scope="scope">
            <el-tooltip class="item" effect="dark" placement="top">
              <div slot="content">{{ scope.row.hash }}</div>
              <el-link type="primary" @click="toDetail(scope.row.hash)"
                ><span class="oneLine">{{ scope.row.hash }}</span></el-link
              >
            </el-tooltip>
          </template>
        </el-table-column>
        <el-table-column align="center" label="交易时间" prop="time">
        </el-table-column>
        <el-table-column align="center" label="交易大小">
          <template slot-scope="scope">
            <span>{{ scope.row.size | formateSize }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" prop="height" label="所属区块">
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
import {
  addressDetail,
  addressList,
  addressName,
} from "@/server/assets/address.js";
export default {
  methods: {},
  data() {
    return {
      url: this.$route.query.url, // 链上地址
      initName: this.$route.query.name, // 资产名称
      detailData: {}, // 资产详情
      options: [],
      value: "",
      tableData: [],
      page: {
        total: 0, // 总页数
        currentPage: 1, // 当前页数
        pageSize: 10, // 每页显示多少条,
      },
    };
  },
  created() {
    this.getName(this.url); // 获取资产名称下拉框
    this.getDetail(this.initName); // 获取详情
    this.getList(this.page); // 获取交易列表
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
    // 获取名称下拉框数据
    getName(url) {
      let newUrl = "/addresses/" + url;
      addressName(newUrl).then((res) => {
        this.options = (res && res.coins) || [];
      });
    },
    // 详情数据
    getDetail(name) {
      let toUrl = "/assets/" + name;
      addressDetail(toUrl).then((res) => {
        this.detailData = res;
      });
    },
    // 分页获取数据
    getList(page, params = {}) {
      let url = `/addresses/${this.url}/transactions`;
      let name = this.initName;
      let paramsData = Object.assign(
        {
          page_num: page.currentPage,
          page_size: page.pageSize,
          coin: name,
        },
        params
      );
      addressList(url, paramsData).then((res) => {
        this.tableData = res.txs;
        this.page.total = res.total;
      });
    },
    // 交易详情
    toDetail(hash) {
      this.$router.push({
        path: "/block/tradeInfo",
        query: {
          name: hash,
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
    // 修改资产名称后
    changeName() {
      this.getDetail(this.initName); // 获取详情
      this.getList(this.page); // 获取交易列表
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
  cursor: pointer;
}
.el-icon-document-copy {
  cursor: pointer;
}
</style>
