<!-- 资产列表 -->
<template>
  <div class="assets width1200">
    <div class="breadcrumb">
      <el-breadcrumb separator="/">
        <el-breadcrumb-item :to="{ path: '/index' }">首页</el-breadcrumb-item>
        <el-breadcrumb-item>资产</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <div class="contract-table">
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
        <el-table-column align="center" label="资产名称">
          <template slot-scope="scope">
            <el-link type="primary" @click="toDetail(scope.row.denom)">{{
              scope.row.denom
            }}</el-link>
          </template>
        </el-table-column>
        <el-table-column align="center" label="资产ID">
          <template slot-scope="scope">
            <el-tooltip class="item" effect="dark" placement="top">
              <div slot="content">{{ scope.row.ref_tx }}</div>
              <el-link type="primary"
                ><span
                  class="oneLine oneLine180"
                  @click="toDetail(scope.row.denom)"
                  >{{ scope.row.ref_tx }}</span
                ></el-link
              >
            </el-tooltip>
          </template>
        </el-table-column>
        <el-table-column
          align="center"
          prop="assetIssueTotal"
          label="资产发行总量"
          width="150"
        >
          <template slot-scope="scope">
            {{ scope.row.mint_amount + scope.row.burn_amount }}
          </template>
        </el-table-column>
        <el-table-column
          width="150"
          align="center"
          prop="mint_amount"
          label="资产可用总量"
        >
        </el-table-column>
        <el-table-column align="center" prop="time" label="发行时间">
        </el-table-column>
        <el-table-column align="center" label="资产发行人">
          <template slot-scope="scope">
            <el-tooltip class="item" effect="dark" placement="top">
              <div slot="content">{{ scope.row.initiator }}</div>
              <el-link type="primary"
                ><span
                  class="oneLine oneLine180"
                  @click="toDetail(scope.row.denom)"
                  >{{ scope.row.initiator }}</span
                ></el-link
              >
            </el-tooltip>
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
import { assetsList } from "@/server/assets/index.js";
export default {
  name: "assets",
  data() {
    return {
      tableData: [
        {
          assetAlais: "0xfdafjaifdjafdnafafa",
          assetID: "1",
          assetIssueTotal: 10000,
          assetAvailTotal: 1000,
          assetIssueTime: "2020-12-12 12-12-12",
          assetProducer: "0xdfafadfad8fjafnadfja",
        },
      ],
      page: {
        total: 0, // 总页数
        currentPage: 1, // 当前页数
        pageSize: 10, // 每页显示多少条,
      },
    };
  },
  created() {
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
    // 分页获取数据
    getList(page, params = {}) {
      let url = "/assets";
      let paramsData = Object.assign(
        {
          page_num: page.currentPage,
          page_size: page.pageSize,
        },
        params
      );
      assetsList(url, paramsData).then((res) => {
        console.log(res);
        this.tableData = res.assets;
        this.page.total = res.total;
      });
    },
    // 资产详情
    toDetail(name) {
      console.log(name);
      this.$router.push({
        path: "/assets/detail",
        query: {
          name,
        },
      });
    },
  },
};
</script>
<style scoped></style>
