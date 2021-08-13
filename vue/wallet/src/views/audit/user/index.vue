<template>
  <div class="com-breadcrumb trx">
    <div class="trx-title">用户管理</div>
    <div class="trx-condition">
      <div class="trx-condition-row">
        <span>操作行为</span>
        <el-select v-model="actionValue" @change="actionChange">
          <template v-if="actions.length">
            <el-option
              v-for="action in actions"
              :key="action.id"
              :label="action.label"
              :value="action.value"
            >
            </el-option>
          </template>
        </el-select>
      </div>
    </div>
    <div class="trx-list">
      <template v-if="peers.length">
        <template v-for="peer in peers">
          <div class="cm-module-bg trx-list-row" :key="peer.hash">
            <div class="list-row-top">
              <div class="top-assets assets-left">
                <div class="assets-title">操作人</div>
                <div class="assets-name">
                  <el-tooltip
                    effect="dark"
                    :content="peer.operator"
                    placement="top"
                  >
                    <span class="cm-text-overflow trx-hash">
                      {{ peer.operator || '---' }}</span
                    ></el-tooltip
                  >
                </div>
              </div>
              <div class="top-assets op-object">
                <div class="assets-title">操作对象</div>
                <div class="assets-name">
                  <el-tooltip
                    effect="dark"
                    :content="peer.object"
                    placement="top"
                  >
                    <span class="cm-text-overflow object-text">{{
                      peer.object || '---'
                    }}</span>
                  </el-tooltip>
                </div>
              </div>
              <div class="top-assets op-action">
                <div class="assets-title">操作行为</div>
                <div class="assets-name">{{ peer.action || '---' }}</div>
              </div>
            </div>
            <div class="list-row-bottom">
              <div>
                <span class="trx-hash-name">交易哈希</span>
                <el-tooltip effect="dark" :content="peer.hash" placement="top">
                  <span class="cm-text-overflow trx-hash">
                    {{ peer.hash || '---' }}</span
                  ></el-tooltip
                >
              </div>
              <div>
                <span class="trx-hash-name">交易时间</span>
                <span
                  >{{ peer.time | dateFormat('YYYY-mm-dd HH:MM:SS') }}
                </span>
              </div>
            </div>
          </div>
        </template>
      </template>
      <template v-else>
        <div class="cm-module-bg trx-list-row trx-no-data">
          暂无用户管理数据
        </div>
      </template>

      <div v-if="peers && peers.length" class="pagination-main">
        <el-pagination
          class="pagination"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
          :current-page="page.pageNum"
          :page-sizes="[10, 20, 30, 40]"
          :page-size="page.pageSize"
          layout="total, sizes, prev, pager, next, jumper"
          :total="page.pageTotal"
          background
        >
        </el-pagination>
      </div>
    </div>
  </div>
</template>
<script>
import { queryUsers } from '@/server/audit'
export default {
  data() {
    return {
      actionValue: '全部',
      actions: [
        {
          id: 1,
          label: '全部',
          value: '全部',
        },
        {
          id: 2,
          label: '设置用户权限',
          value: '设置用户权限',
        },
      ],
      peers: [],
      page: {
        pageSize: 10,
        pageTotal: 0,
        pageNum: 1,
      },
      walAdd: '',
    }
  },
  created() {
    let { address } = this.$route.query
    this.walAdd = address
    this.getPeers({ address })
  },
  methods: {
    handleSizeChange(pageSize) {
      this.getPeers({
        address: this.walAdd,
        action: this.actionValue,
        page_num: this.page.pageNum,
        page_size: pageSize,
      })
    },
    handleCurrentChange(pageNum) {
      this.getPeers({
        address: this.walAdd,
        action: this.actionValue,
        page_num: pageNum,
        page_size: this.page.pageSize,
      })
    },
    actionChange(action) {
      if (action == '全部') {
        action = ''
      }
      this.getPeers({ address: this.walAdd, action })
    },
    async getPeers({ action, address, pageNum = 1, pageSize = 10 }) {
      let { page_num, page_size, total, items } = await queryUsers({
        action,
        address,
        page_size: pageSize,
        page_num: pageNum,
      })

      if (items && items.length) {
        this.page.pageNum = pageNum
        this.page.pageSize = pageSize
        this.page.pageTotal = total
        this.peers = items
      }
    },
  },
}
</script>
<style >
.trx-title {
  font-family: 'PingFangSC-Medium', 'PingFang SC Medium', 'PingFang SC',
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 18px;
  color: #ffffff;
}
.trx-condition {
  display: flex;
}

.trx-condition-row {
  margin: 30px 0 15px;
  margin-right: 40px;
}
.trx-condition-row :last-child {
  margin-right: 0;
}
.trx-condition-row span {
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #ffffff;
  margin: 0 30px 0 0;
}
.trx-list {
  padding: 16px 0 0;
  min-height: 300px;
}
.trx-list-row {
  border-radius: 5px;
  height: 155px;
  display: flex;
  flex-direction: column;
  padding: 26px 30px;
  justify-content: center;
  margin: 0 0 20px;
}
.trx-list-row:last-child {
  margin-bottom: 0;
}
.list-row-top {
  display: flex;
  align-items: center;
  align-content: center;
}
.top-assets {
  width: 20%;
}
.assets-left {
  width: 40%;
}
.assets-title {
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #768ca8;
}
.assets-name {
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 18px;
  color: #ffffff;
  margin: 14px 0 0;
}
.list-row-bottom {
  display: flex;
  justify-content: space-between;
  margin: 30px 0 0;
  color: #cccccc;
}
.trx-btn-info {
  text-align: right;
  color: #62f7d4;
  cursor: pointer;
}
.trx-btn-info img {
  width: 20px;
  height: 20px;
  margin: 0 auto 18px;
  transform: translate(-16px, 0px);
}
.trx-hash-name {
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 12px;
  color: #768ca8;
  margin: 0 14px 0 0;
}
.trx-status {
  text-align: center;
}
.trx-hash {
  display: inline-block;
  width: 80%;
  vertical-align: middle;
}
.trx-no-data {
  text-align: center;
  font-size: 16px;
}
.op-object {
  width: 40%;
}
.op-action {
  width: 20%;
}
.object-text {
  width: 80%;
  display: inline-block;
}
</style>