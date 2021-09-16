<template>
  <div class="com-breadcrumb contract">
    <div class="trx-title">我的合约</div>
    <div class="contract-desc">
      <div class="cm-sub-bg market-main-desc">
        <div class="market-img">
          <img
            src="../../assets/images/market/market_icon.png"
            alt=""
            srcset=""
          />
        </div>
        <div class="market-desc">
          <div class="cm-text-overflow desc-name">Hi，欢迎使用智能合约</div>
          <div class="cm-text-overflow2 desc-text">
            轻松创建、部署、管理智能合约，提升研发效率，降低业务成本
          </div>
        </div>
        <div class="contract-btns">
          <template v-if="wallet.threshold > 0">
            <a
              href="javascript:;"
              class="cm-btn-border009F72 cm-btn-225px"
              @click="$router.push('/mycontract/signature')"
              ><span class="iconfont m0-bianjimian edit-icon"></span
              >智能合约多签签名</a
            >
          </template>
          <a
            href="javascript:;"
            class="cm-btn-bg4acb9b cm-btn-225px publish-assets"
            @click="$router.push(`/myContract/create`)"
            >创建智能合约</a
          >
        </div>
      </div>
    </div>
    <div class="trx-list">
      <template v-if="contractLists && contractLists.length">
        <template v-for="contract in contractLists">
          <div
            class="cm-module-bg trx-list-row assets-row-item"
            :key="contract.id"
          >
            <div class="list-row-top">
              <div class="row-top-left">
                <div class="assets-left">
                  <div class="assets-title">合约名称</div>
                  <div class="assets-name">{{ contract.name || '---' }}</div>
                </div>
                <div class="assets-left">
                  <div class="assets-title">版本</div>
                  <div class="assets-name">{{ contract.version || '---' }}</div>
                </div>
                <div class="assets-left">
                  <div class="assets-title">状态</div>
                  <div class="assets-name">
                    {{ contract.status | contractStatusText }}
                  </div>
                </div>
                <div class="assets-left">
                  <div class="assets-title">更新时间</div>
                  <div class="assets-name">
                    {{ contract.updated_at || '---' }}
                  </div>
                </div>
              </div>
              <div class="top-items">
                <div
                  class="tp-item-tab"
                  @click="$router.push(`/mycontract/detail?id=${contract.id}`)"
                >
                  <span class="iconfont m0-edit-fill"></span>
                  详情
                </div>
                <template v-if="contract.status == 0 || contract.status == 3">
                  <div
                    class="tp-item-tab"
                    @click="handleContract(contract, 'undeploy')"
                  >
                    <span class="iconfont m0-shanchu1"></span>
                    删除
                  </div>
                  <div
                    class="tp-item-tab"
                    @click="handleContract(contract, 'deploy')"
                  >
                    <span class="iconfont m0-xiazai1"></span>
                    部署
                  </div>
                </template>
                <template
                  v-if="
                    contract.status == 2 ||
                    contract.status == 5 ||
                    contract.status == 8
                  "
                >
                  <div
                    class="tp-item-tab"
                    @click="handleContract(contract, 'upgrade')"
                  >
                    <span class="iconfont m0-shengji"></span>
                    升级
                  </div>
                  <div
                    class="tp-item-tab"
                    @click="handleContract(contract, 'freeze')"
                  >
                    <span class="iconfont m0-dongjiejine"></span>
                    冻结
                  </div>
                </template>
                <template v-if="contract.status == 7">
                  <div
                    class="tp-item-tab"
                    @click="handleContract(contract, 'upgrade')"
                  >
                    <span class="iconfont m0-shengji"></span>
                    重新升级
                  </div>
                </template>
                <template v-if="contract.status == 4">
                  <div
                    class="tp-item-tab"
                    @click="handleContract(contract, 'unfreeze')"
                  >
                    <span class="iconfont m0-jiedong"></span>
                    解冻
                  </div>
                </template>
              </div>
            </div>
          </div>
        </template>
      </template>
      <template v-else>
        <div class="cm-module-bg assets-nodata">暂无合约数据</div>
      </template>

      <div v-if="contractLists && contractLists.length" class="pagination-main">
        <el-pagination
          class="pagination"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
          :current-page="page.pageNum"
          :page-sizes="[10, 20, 30, 40]"
          :page-size="page.pageSize"
          layout="total, sizes, prev, pager, next, jumper"
          :total="page.total"
          background
        >
        </el-pagination>
      </div>
    </div>

    <div class="dialog" v-if="isDelte">
      <div class="dialog-main del-dialog">
        <div class="dialog-main-title">
          <span class="main-title-name"> 删除确认</span>
          <a
            href="javascript:;"
            class="iconfont m0-guanbi"
            @click="isDelte = false"
          ></a>
        </div>
        <div class="dialog-main-content">
          <div class="dialog-main-desc">是否确认删除智能合约？</div>
          <div class="dialog-main-btns">
            <a
              href="javascript:;"
              class="cm-btn-138px cm-btn-border009F72"
              @click="isDelte = false"
              >取消</a
            >
            <a
              href="javascript:;"
              class="cm-btn-138px cm-btn-bg4acb9b"
              @click="conformContract"
              >确定</a
            >
          </div>
        </div>
      </div>
      <div class="dialog-yy"></div>
    </div>
  </div>
</template>
<script>
import {
  queryContractLists,
  contractOperate,
  delContract,
} from '@/server/contract'

import { localCache } from '@/utils/utils'
export default {
  data() {
    return {
      wallet: {},
      page: {
        pageSize: 10,
        total: 0,
        pageNum: 1,
      },
      contractLists: [],
      isDelte: false,
      isDelteInfo: {},
    }
  },
  created() {
    let wallet = localCache.get('wallet')
    if (wallet) {
      this.wallet = wallet
      this.getContractList(
        wallet.address,
        this.page.pageNum,
        this.page.pageSize
      )
    }
  },
  destroyed() {
    clearInterval(this.timer)
  },
  methods: {
    conformContract() {
      this.loading = this.$loading({
        lock: true,
        text: '删除中...',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)',
      })
      let { isDelteInfo } = this
      let delInfo = this.delContractById(isDelteInfo.id)
      if (delInfo) {
        this.loading.close()
        this.isDelte = false
        this.contractLists = this.contractLists.filter(
          (item) => item.id != isDelteInfo.id
        )
        this.isDelteInfo = {}
      }
    },
    handleContract(contract, mode) {
      if (mode == 'upgrade') {
        this.$router.push(`/myContract/upgrade?id=${contract.id}&mode=${mode}`)
      } else if (mode == 'undeploy') {
        this.isDelte = true
        this.isDelteInfo = contract
      } else {
        this.$router.push(`/myContract/deploy?id=${contract.id}&mode=${mode}`)
      }
    },
    handleSizeChange(pageSize) {
      this.getContractList(this.wallet.address, this.page.pageNum, pageSize)
    },
    handleCurrentChange(pageNum) {
      this.getContractList(this.wallet.address, pageNum, this.page.pageSize)
    },
    async getContractList(account, pageNum, pageSize) {
      let { page_size, page_num, total, list } = await queryContractLists({
        account,
        page_num: pageNum,
        page_size: pageSize,
      })
      if (list) {
        this.contractLists = list
        this.page.pageNum = page_num
        this.page.pageSize = page_size
        this.page.total = total
      }
    },
    async delContractById(id) {
      let delRes = await delContract({ id })
      if (delRes) {
        return delRes
      }
    },
  },
}
</script>
<style scoped>
.trx-title {
  font-family: 'PingFangSC-Medium', 'PingFang SC Medium', 'PingFang SC',
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 18px;
  color: #ffffff;
}
.trx-list {
  padding: 16px 0 0;
}
.trx-list-row {
  border-radius: 5px;
  height: 140px;
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
  margin: 30px 0 0;
}
.row-top-left {
  flex: 1;
  display: flex;
}
/**asset start */
.top-items {
  display: flex;
  height: 125px;
  text-align: center;
  align-items: center;
  width: 240px;
  padding: 0 0 0 26px;
}
.tp-item-tab {
  text-align: center;
  cursor: pointer;
  color: #63f7d4;
  /* width: 120px; */
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  font-size: 16px;
  margin: 0 0 0 30px;
}
.tp-item-tab:first-child {
  margin-left: 0;
}
.tp-item-tab span {
  font-size: 22px;
  margin: 0 0 10px;
}
.publish-assets {
  margin: 0 0 0 20px;
}

.edit-icon {
  margin: 0 10px 0 0;
  color: #63f7d4;
  font-size: 14px;
}
/**asset start */
.assets-nodata {
  min-height: 200px;
  text-align: center;
  line-height: 200px;
}
/**end asset */

.contract-desc {
  padding: 20px 0 0;
}
.market-main-desc {
  display: flex;
  align-items: center;
  height: 140px;
  padding: 0px 40px;
  border-radius: 11px;
}
.market-img {
  width: 60px;
  height: 60px;
}
.market-img img {
  width: 60px;
  height: 100%;
}
.market-desc {
  margin: 0 0 0 30px;
  flex: 1;
}
.desc-name {
  font-family: 'PingFangSC-Medium', 'PingFang SC Medium', 'PingFang SC',
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 24px;
  color: #62f7d4;
}
.desc-text {
  font-family: 'PingFangSC-Medium', 'PingFang SC Medium', 'PingFang SC',
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 14px;
  color: #cbcbcc;
  margin: 16px 0 0;
  width: 80%;
  line-height: 20px;
}

.contract-btns {
  display: flex;
}

.main-title-name {
  flex: 1;
}
.del-dialog {
  display: flex;
  flex-direction: column;
  padding: 0 0 38px;
}
.dialog-main-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-content: center;
  align-items: center;
}
.dialog-main-desc {
  flex: 1;
  min-height: 150px;
  display: flex;
  align-items: center;
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-size: 14px;
  color: #ffffff;
}
.dialog-main-btns {
  display: flex;
}
.dialog-main-btns a {
  margin: 0 0 0 20px;
  font-size: 14px;
}
.dialog-main-btns a:first-child {
  margin-left: 0;
}
</style>