#!/usr/bin/env bash

set -e

source $(dirname $0)/happy_auth.sh

# 一、propose-add-account 提议新增
m0d keys add addAccount --home ~/output/node0/.m0/
proposeAddAccount alice addAccount authority,wasm,utxo,mibc

#二、approve-add-account 授权新增账户
# 1.新增一个新账户，创建第三个管理员.三个管理账户：alice addAccount addAccount1
m0d keys add addAccount1 --home ~/output/node0/.m0/
proposeAddAccount alice addAccount1 authority,wasm,utxo,mibc

# 2.提议新增一个普通用户
m0d keys add addAccount2 --home ~/output/node0/.m0/
proposeAddAccount alice addAccount2 wasm,utxo,mibc
# 3.授权新增账户，alice已经授权，再使用其他管理员授权一次就超过2/3.
approveAddAccount addAccount addAccount2

# 三、 propose-modify-account
#提议修改权限
proposeModifyAccount addAccount addAccount2 wasm,utxo,mibc
#授权提议修改权限
approveModifyAccount alice addAccount2

# 四、 propose-revoke-account
#提议删除
proposeRevokeAccount addAccount addAccount2
#授权提议删除
approveRevokeAccount alice addAccount2