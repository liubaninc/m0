<!-- 合约详情 -->
<template>
  <div class="assetsDetail width1200">
    <div class="breadcrumb">
      <el-breadcrumb separator="/">
        <el-breadcrumb-item :to="{ path: '/index' }">首页</el-breadcrumb-item>
        <el-breadcrumb-item :to="{ path: '/contract' }"
          >合约</el-breadcrumb-item
        >
        <el-breadcrumb-item>合约详情</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <div class="Detail-message">
      <div class="title">合约名称 / {{ name }}</div>
      <el-row>
        <el-col :span="12">
          <div class="grid-content bg-purple-light">
            <span class="label">合约名称</span>
            <span class="link">{{ detailData.name }}</span>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="grid-content bg-purple-light">
            <span class="label">合约状态</span>
            <span>{{ detailData.status == 0 ? "正常" : "异常" }}</span>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="grid-content">
            <span class="label">合约更新时间</span>
            <span>{{ detailData.time }}</span>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="grid-content">
            <span class="label">合约版本</span>
            <span class="oneLine">{{ detailData.version }}</span>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="grid-content bg-purple-light">
            <span class="label">合约关联交易量</span>
            <span>{{ detailData.total }}</span>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="grid-content bg-purple-light">
            <span class="label">合约发行人</span>
            <span
              class="link oneLine"
              @click="toAddress(detailData.name, detailData.initiator)"
              >{{ detailData.initiator }}</span
            >
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
    <div class="table-tlt"></div>
    <div class="Detail-table">
      <el-tabs type="border-card" v-model="activeName" @tab-click="handleClick">
        <el-tab-pane label="合约关联交易" name="first">
          <span slot="label"><i class="el-icon-s-shop"></i>合约关联交易</span>
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
            <el-table-column align="center" label="交易哈希" width="360">
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
            <el-table-column align="center" prop="size" label="交易大小">
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
        </el-tab-pane>
        <el-tab-pane label="合约升级历史" name="second">
          <span slot="label"><i class="el-icon-upload"></i>合约升级历史</span>
          <el-table
            :data="tableData2"
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
              <template>
                <span>{{ name }}</span>
              </template>
            </el-table-column>
            <el-table-column align="center" prop="time" label="交易时间">
            </el-table-column>
            <el-table-column align="center" prop="version" label="合约版本">
            </el-table-column>
            <el-table-column align="center" label="所属交易" width="400">
              <template slot-scope="scope">
                <el-tooltip class="item" effect="dark" placement="top">
                  <div slot="content">{{ scope.row.hash }}</div>
                  <el-link type="primary" @click="toDetail(scope.row.hash)"
                    ><span class="oneLine">{{ scope.row.hash }}</span></el-link
                  >
                </el-tooltip>
              </template>
            </el-table-column>
          </el-table>
          <el-pagination
            class="pagination"
            @size-change="handleSizeChange2"
            @current-change="handleCurrentChange2"
            :current-page="page2.currentPage"
            :page-sizes="[10, 20, 30, 40]"
            :page-size="page2.pageSize"
            layout="total, sizes, prev, pager, next, jumper"
            :total="page2.total"
            background
          >
          </el-pagination>
        </el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>

<script>
import { contractDetail, detailList } from "@/server/contract/index.js";
export default {
  methods: {},
  data() {
    return {
      name: this.$route.query.name,
      detailData: {},
      activeName: "first", // // false,关联交易；true,升级历史
      typeTab: false, // false,关联交易；true,升级历史
      tableData: [],
      page: {
        total: 0, // 总页数
        currentPage: 1, // 当前页数
        pageSize: 10, // 每页显示多少条,
      },
      tableData2: [],
      page2: {
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
    // 切换第几页
    handleCurrentChange2: function(page) {
      this.page2.currentPage = page;
      this.getList(this.page2);
    },
    // 每页几条
    handleSizeChange2: function(sizeNum) {
      this.page2.pageSize = sizeNum;
      this.getList(this.page2);
    },
    // 详情数据
    getDetail(name) {
      let url = "/contracts/" + name;
      contractDetail(url).then((res) => {
        this.detailData = res;
      });
    },
    // 分页获取数据
    getList(page, params = {}) {
      let url = `/contracts/${this.name}/transactions`;
      let typeTab = this.typeTab;
      let paramsData = Object.assign(
        {
          page_num: page.currentPage,
          page_size: page.pageSize,
          invoke: typeTab,
        },
        params
      );
      detailList(url, paramsData).then((res) => {
        if (!typeTab) {
          this.tableData = res.txs;
          this.page.total = res.total;
        } else {
          this.tableData2 = res.txs;
          this.page2.total = res.total;
        }
      });
    },
    // 资产地址
    toDetail(name) {
      this.$router.push({
        path: "/block/tradeInfo",
        query: {
          name,
        },
      });
    },
    // toAddress
    toAddress(name, url) {
      this.$router.push({
        path: "/assets/address",
        query: {
          name,
          url,
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
    // tab切换
    handleClick(tab, event) {
      console.log(tab.name);
      if (tab.name == "first") {
        this.typeTab = false;
      } else if (tab.name == "second") {
        this.typeTab = true;
      }
      this.getList(this.page);
    },
  },
};
</script>
<style scoped>
.table-tlt {
  margin-top: 30px;
  margin-bottom: 15px;
  font-size: 16px;
}
.table-tlt span {
  display: inline-block;
  margin-right: 10px;
  cursor: pointer;
  font-weight: 700;
}
.table-tlt .tabActive {
  color: #409eff;
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
