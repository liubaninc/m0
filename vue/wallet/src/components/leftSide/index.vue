<template>
  <div class="cm-module-bg slide">
    <div class="slide-warpper">
      <div class="silde-top">
        <a
          href="javascript:;"
          class="silde-top-title"
          @click="$router.push(`/wallet`)"
        >
          <span class="title-icon"></span>
          返回钱包列表</a
        >
        <div class="wallet-type">
          <template v-if="wallet.threshold > 0">
            <img
              src="../../assets/images/detail/wallet_more_icon.png"
              alt=""
              srcset=""
            />
          </template>
          <template v-else>
            <img
              src="../../assets/images/detail/wallet_single_icon.png"
              alt=""
              srcset=""
            />
          </template>
          <div class="cm-text-overflow silde-top-name">{{ wallet.name }}</div>
        </div>
      </div>
      <div class="silde-menu">
        <div class="slide-menu-row" v-for="item in menus" :key="item.id">
          <div class="menu-row-title">
            <!-- <img
              src="../../assets/images/detail/menu/menu_core_icon.png"
              alt=""
              srcset=""
            /> -->
            <span class="item-icon" :class="item.icon"></span>
            {{ item.text }}
          </div>
          <template v-if="item.children.length">
            <!-- :class="
                child.url && $route.fullPath.indexOf(child.url) != -1
                  ? 'curRow'
                  : ''
              " -->
            <!-- <div
              class="menu-row-list"
              v-for="(child, index) in item.children"
              :key="child.id"
              :class="child.url && $route.path == child.url ? 'curRow' : ''"
              @click="selTab(child, index)"
            > -->
            <div
              class="menu-row-list"
              v-for="(child, index) in item.children"
              :key="child.id"
              :class="child.url && $route.path == child.url ? 'curRow' : ''"
              @click="selTab(child, index)"
            >
              <router-link :to="`${child.url}?address=${wallet.address}`">
                <div class="row-list-tab">
                  <span class="item-icon" :class="child.icon"></span>

                  {{ child.text }}
                </div>
              </router-link>
            </div>
          </template>
        </div>
      </div>
      <div class="silde-bottom">
        <div class="slide-bt-status">钱包状态</div>
        <div class="slide-bt-link">链接状态 <span>链接正常</span></div>
        <div class="slide-bt-height">
          同步完成(高度:{{ blockInfo.block_num }})
        </div>
        <div class="slide-bt-process">
          <div class="bt-process-bar"></div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { queryBlockChain } from '@/server/block'
import { localCache } from '@/utils/utils'
import { leftMenus } from '@/config/leftMenu'
export default {
  data() {
    return {
      menus: leftMenus,
      wallet: {},
      blockInfo: {},
      curIndex: 'core_0',
    }
  },
  created() {
    let wallet = localCache.get('wallet')
    console.log('=======>', this.$route)
    if (wallet) {
      this.wallet = wallet
      this.getBlock()
    }
  },
  mounted() {
    this.timer = setInterval(() => {
      this.getBlock()
    }, 1000 * 60 * 2)
  },
  destroyed() {
    clearInterval(this.timer)
  },
  methods: {
    selTab(child, index) {
      this.curIndex = child.preKey + index
    },
    async getBlock() {
      let blockInfo = await queryBlockChain({})
      if (blockInfo) {
        this.blockInfo = blockInfo
      }
    },
  },
}
</script>
<style>
.slide {
  /* padding: 2px 0; */
  /* height: calc(100vh - 137px); */
  overflow: auto;
  margin: 2px 0;
  min-height: calc(100vh - 101px);
}
.slide-warpper {
  /* padding: 14px 20px 50px; */
  padding: 14px 20px 0px;
  /* min-height: 824px; */
}
.silde-top img {
  width: 36px;
  height: 28px;
}
.silde-top-title {
  font-family: PingFangSC-Medium, 'PingFang SC Medium', 'PingFang SC',
    sans-serif;
  font-weight: 500;
  font-style: normal;
  font-size: 14px;
  color: rgb(118, 140, 168);
  display: block;
  /* margin: 0 0 40px; */
  margin: 0 0 18px;
}
.silde-top-title .title-icon {
  width: 18px;
  height: 14px;
  display: inline-block;
  background: url('../../assets/images/detail/back_icon_default.png');
  background-size: 100% 100%;
  background-repeat: no-repeat;
  vertical-align: bottom;
}
.silde-top-title:hover {
  color: #fff;
}
.silde-top-title:hover .title-icon {
  background: url('../../assets/images/detail/back_icon_cur.png');
  background-size: 100% 100%;
}
.silde-top-name {
  font-family: ArialMT, Arial, sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 18px;
  color: rgb(255, 255, 255);
  /* margin: 10px 0 0; */
  width: 90%;
  margin: 0 0 0 10px;
}
/***menu start */
.silde-menu {
  border-top: 1px solid #797979;
  border-bottom: 1px solid #797979;
  margin: 20px 0;
  padding: 16px 0;
}
.slide-menu-row {
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #ffffff;
}
.menu-row-title {
  height: 40px;
  line-height: 40px;
  display: flex;
  justify-items: center;
  align-items: center;
}
.menu-row-title img {
  width: 16px;
  height: 16px;
  margin: 0 10px 0 0;
}
.menu-row-list {
  padding: 0 0 0 24px;
}
.menu-row-list div {
  height: 40px;
  line-height: 40px;
  cursor: pointer;
  display: flex;
  align-items: center;
}
.row-list-tab img {
  width: 14px;
  height: 14px;
  margin: 0 10px 0 0;
}
.slide-bt-status {
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 14px;
  color: #ffffff;
  line-height: 30px;
}
.slide-bt-link {
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 12px;
  color: #fff;
  line-height: 30px;
}
.slide-bt-link span {
  /* color: #ff5a58; */
  margin: 0 0 0 20px;
}
.slide-bt-height {
  font-family: 'PingFangSC-Regular', 'PingFang SC', sans-serif;
  font-weight: 400;
  font-style: normal;
  font-size: 12px;
  color: #768ca8;
}
.slide-bt-process {
  height: 30px;
  line-height: 30px;
}
.wallet-type {
  display: flex;
  align-items: center;
}

.list-tab-cur {
  color: #62f7d4;
}
.slide-bt-process {
  width: 100%;
  height: 5px;
  background-color: rgba(242, 242, 242, 1);
  margin: 30px 0 0;
  border-radius: 5px;
}
.bt-process-bar {
  width: 80%;
  height: 5px;
  background-color: rgba(0, 159, 114, 1);
  border-radius: 5px;
}
.item-icon {
  margin: 0 10px 0 0;
}
/***end menu */

.curRow {
  background-color: rgba(26, 44, 66, 1);
  font-size: 14px;
  color: #62f7d4;
}
::-webkit-scrollbar {
  display: none;
}
</style>