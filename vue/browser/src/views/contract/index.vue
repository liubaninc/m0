<!-- 合约列表 -->
<template>
  <div class="contract width1200">
    <div class="breadcrumb">
      <el-breadcrumb separator="/">
        <el-breadcrumb-item :to="{ path: '/index' }">首页</el-breadcrumb-item>
        <el-breadcrumb-item>合约</el-breadcrumb-item>
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
        <el-table-column align="center" label="合约名称">
          <template slot-scope="scope">
            <el-link type="primary" @click="toDetail(scope.row.name)">{{
              scope.row.name
            }}</el-link>
          </template>
        </el-table-column>
        <el-table-column align="center" label="合约版本" width="120">
          <template slot-scope="scope">
            <span>{{ scope.row.version }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="合约状态" width="100">
          <template slot-scope="scope">
            <el-tag v-if="scope.row.status == 0">正常</el-tag>
            <el-tag v-if="scope.row.status == 1" type="danger">异常</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="center" prop="time" label="合约更新时间">
        </el-table-column>
        <el-table-column align="center" label="合约发行人" width="400">
          <template slot-scope="scope">
            <el-tooltip class="item" effect="dark" placement="top">
              <div slot="content">{{ scope.row.initiator }}</div>
              <span class="oneLine">{{ scope.row.initiator }}</span>
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
import { contractList } from "@/server/contract/index.js";
export default {
  methods: {},
  data() {
    return {
      tableData: [
        {
          contractAlais: "0xfdafjaifdjafdnafafa",
          contractVersion: 1.0,
          contractStatus: 0,
          contractUpdateTime: "2020-12-12 12:12:12",
          contractProducer: "0xdfafadfad8fjafnadfja",
        },
      ],
      page: {
        total: 0, // 总页数
        currentPage: 1, // 当前页数
        pageSize: 20, // 每页显示多少条,
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
      let url = "/contracts";
      let paramsData = Object.assign(
        {
          page_num: page.currentPage,
          page_size: page.pageSize,
        },
        params
      );
      contractList(url, paramsData).then((res) => {
        this.tableData = res.contracts;
        this.page.total = res.total;
      });
    },
    // 合约详情
    toDetail(name) {
      this.$router.push({
        path: "/contract/detail",
        query: {
          name,
        },
      });
    },
  },
};
</script>
<style scoped></style>
