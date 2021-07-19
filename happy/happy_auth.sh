#!/usr/bin/env bash


#提议新增账户
function proposeAddAccount() {
  from=$1
  to=$2
  roles=$3

  echo "propose-add-account $from $to "

  before_propose_list=$(m0d query authority list-account --home ~/output/node0/.m0/ --output json)
  before_propose_to=$(m0d query authority show-account $(m0d keys show $to -a --home ~/output/node0/.m0/) --home ~/output/node0/.m0/ --output json)

  #发送新增账户消息
  hash=$(m0d tx authority propose-add-account $(m0d keys show $to -a --home ~/output/node0/.m0/) $(m0d keys show $to -p --home ~/output/node0/.m0/) $roles --from $from --chain-id happy --home ~/output/node0/.m0/ -y --broadcast-mode block | jq ".txhash")

  tx=$(m0d query tx $(echo $hash | sed 's/\"//g') --home ~/output/node0/.m0/)
  #查看新增账户信息
  after_propose_list=$(m0d query authority list-account --home ~/output/node0/.m0/ --output json)
  after_propose_to=$(m0d query authority show-account $(m0d keys show $to -a --home ~/output/node0/.m0/) --home ~/output/node0/.m0/ --output json)

  echo "===before_propose_account_list=== $before_propose_list"
  echo "===before_propose_show_new_add_account=== $before_propose_to"

  echo "propose-add-account send message"
  echo "propose-add-account message tx hash: $hash"
  echo "propose-add-account message tx info: $tx"

  echo "***after_propose_account_list*** $after_propose_list"
  echo "***after_propose_show_new_add_account*** $after_propose_to"
  echo ""

  return 0
}

# 2/3的管理员账户授权`新增账户`提议，
function approveAddAccount() {

  from=$1
  to=$2

  echo "approve-add-account $from $to "

  before_approve_list=$(m0d query authority list-proposed-account --home ~/output/node0/.m0/ --output json)
  before_approve_to=$(m0d query authority show-proposed-account $(m0d keys show $to -a --home ~/output/node0/.m0/) --home ~/output/node0/.m0/)

  #发送新增账户消息
  hash=$(m0d tx authority approve-add-account $(m0d keys show $to -a --home ~/output/node0/.m0/) --from $from --chain-id happy --home ~/output/node0/.m0/ -y --broadcast-mode block | jq ".txhash")

  tx=$(m0d query tx $(echo $hash | sed 's/\"//g') --home ~/output/node0/.m0/)
  #查看新增账户信息
  after_approve_list=$(m0d query authority list-proposed-account --home ~/output/node0/.m0/ --output json)
  after_approve_to=$(m0d query authority list-proposed-account $(m0d keys show $to -a --home ~/output/node0/.m0/) --home ~/output/node0/.m0/ --output json)

  echo "===before_approve_account_list=== $before_approve_list"
  echo "===before-approve_to=== $before_approve_to"

  echo "approve-add-accountt send message"
  echo "approve-add-account message tx hash: $hash"
  echo "approve-add-account message tx info: $tx"

  echo "===after_approve_account_list=== $after_approve_list"
  echo "===after_approve_to=== $after_approve_to"
  echo ""

  return 0
}

#提议修改账户
function proposeModifyAccount() {
  from=$1
  to=$2
  roles=$3

  echo "propose-modify-account $from $to "

  before_propose_modify_list=$(m0d query authority list-account --home ~/output/node0/.m0/ --output json)
  before_propose_modify_to=$(m0d query authority show-account $(m0d keys show $to -a --home ~/output/node0/.m0/) --home ~/output/node0/.m0/ --output json)

  #发送修改账户消息
  hash=$(m0d tx authority propose-modify-account $(m0d keys show $to -a --home ~/output/node0/.m0/) $roles --from $from --chain-id happy --home ~/output/node0/.m0/ -y --broadcast-mode block | jq ".txhash")

  tx=$(m0d query tx $(echo $hash | sed 's/\"//g') --home ~/output/node0/.m0/)
  #查看账户信息
  after_propose_modify_list=$(m0d query authority list-proposed-account --home ~/output/node0/.m0/ --output json)
  after_propose_modify_to=$(m0d query authority show-proposed-account $(m0d keys show $to -a --home ~/output/node0/.m0/) --home ~/output/node0/.m0/ --output json)

  echo "===before_propose__modify_account_list=== $before_propose_modify_list"
  echo "===before_propose_modify_account_to=== $before_propose_modify_to"

  echo "propose-modify-account send message"
  echo "propose-modify-account message tx hash: $hash"
  echo "propose-modify-account message tx info: $tx"

  echo "***after_propose_modify_account_list*** $after_propose_modify_list"
  echo "***after_propose_modify_account_to*** $after_propose_modify_to"
  echo ""

  return 0
}

#授权修改账户
function approveModifyAccount() {
  from=$1
  to=$2

  echo "approve-modify-account $from $to "

  before_approve_modify_list=$(m0d query authority list-proposed-account --home ~/output/node0/.m0/ --output json)
  before_approve_modify_to=$(m0d query authority show-proposed-account $(m0d keys show $to -a --home ~/output/node0/.m0/) --home ~/output/node0/.m0/ --output json)

  #发送修改账户消息
  hash=$(m0d tx authority approve-modify-account $(m0d keys show $to -a --home ~/output/node0/.m0/) --from $from --chain-id happy --home ~/output/node0/.m0/ -y --broadcast-mode block | jq ".txhash")

  tx=$(m0d query tx $(echo $hash | sed 's/\"//g') --home ~/output/node0/.m0/)
  #查看账户信息
  after_approve_modify_list=$(m0d query authority list-account --home ~/output/node0/.m0/ --output json)
  after_approve_modify_to=$(m0d query authority show-account $(m0d keys show $to -a --home ~/output/node0/.m0/) --home ~/output/node0/.m0/ --output json)


  echo "===before_approve__modify_account_list=== $before_approve_modify_list"
  echo "===before_approve_modify_account_to=== $before_approve_modify_to"

  echo "approve-modify-account send message"
  echo "approve-modify-account message tx hash: $hash"
  echo "approve-modify-account message tx info: $tx"

  echo "***after_approve_modify_account_list*** $after_approve_modify_list"
  echo "***after_approve_modify_account_to*** $after_approve_modify_to"
  echo ""

  return 0
}

#提议删除账户
function proposeRevokeAccount() {
  from=$1
  to=$2

  echo "propose-revoke-account $from $to "

  before_propose_revoke_list=$(m0d query authority list-account --home ~/output/node0/.m0/ --output json)
  before_propose_revoke_to=$(m0d query authority show-account $(m0d keys show $to -a --home ~/output/node0/.m0/) --home ~/output/node0/.m0/ --output json)

  #发送删除账户消息
  hash=$(m0d tx authority propose-revoke-account $(m0d keys show $to -a --home ~/output/node0/.m0/) --from $from --chain-id happy --home ~/output/node0/.m0/ -y --broadcast-mode block | jq ".txhash")

  tx=$(m0d query tx $(echo $hash | sed 's/\"//g') --home ~/output/node0/.m0/)
  #查看账户信息
  after_propose_revoke_list=$(m0d query authority list-proposed-account-to-revoke --home ~/output/node0/.m0/ --output json)
  after_propose_revoke_to=$(m0d query authority show-proposed-account-to-revoke $(m0d keys show $to -a --home ~/output/node0/.m0/) --home ~/output/node0/.m0/ --output json)

  echo "===before_propose__revoke_account_list=== $before_propose_revoke_list"
  echo "===before_propose_revoke_account_to=== $before_propose_revoke_to"

  echo "propose-revoke-account send message"
  echo "propose-revoke-account message tx hash: $hash"
  echo "propose-revoke-account message tx info: $tx"

  echo "***after_propose_revoke_account_list*** $after_propose_revoke_list"
  echo "***after_propose_revoke_account_to*** $after_propose_revoke_to"
  echo ""

  return 0
}

#授权删除账户
function approveRevokeAccount() {
  from=$1
  to=$2

  echo "approve-revoke-account $from $to "

  before_approve_revoke_list=$(m0d query authority list-proposed-account-to-revoke --home ~/output/node0/.m0/ --output json)
  before_approve_revoke_to=$(m0d query authority show-proposed-account-to-revoke $(m0d keys show $to -a --home ~/output/node0/.m0/) --home ~/output/node0/.m0/ --output json)
  before_revoke_to=$(m0d query authority show-account $(m0d keys show $to -a --home ~/output/node0/.m0/) --home ~/output/node0/.m0/ --output json)

  #发送删除账户消息
  hash=$(m0d tx authority approve-revoke-account $(m0d keys show $to -a --home ~/output/node0/.m0/) --from $from --chain-id happy --home ~/output/node0/.m0/ -y --broadcast-mode block | jq ".txhash")

  tx=$(m0d query tx $(echo $hash | sed 's/\"//g') --home ~/output/node0/.m0/)
  #查看账户信息
  after_approve_revoke_list=$(m0d query authority list-proposed-account-to-revoke --home ~/output/node0/.m0/ --output json)
  after_approve_revoke_to=$(m0d query authority show-proposed-account-to-revoke $(m0d keys show $to -a --home ~/output/node0/.m0/) --home ~/output/node0/.m0/ --output json)
  after_revoke_to=$(m0d query authority show-account $(m0d keys show $to -a --home ~/output/node0/.m0/) --home ~/output/node0/.m0/ --output json)


  echo "===before_approve__revoke_account_list=== $before_approve_revoke_list"
  echo "===before_approve_revoke_account_to=== $before_approve_revoke_to"
  echo "===before_revoke_account_to=== $before_revoke_to"

  echo "approve-revoke-account send message"
  echo "approve-revoke-account message tx hash: $hash"
  echo "approve-revoke-account message tx info: $tx"

  echo "***after_approve_revoke_account_list*** $after_approve_revoke_list"
  echo "***after_approve_revoke_account_to*** $after_approve_revoke_to"
  echo "===after_revoke_account_to=== $after_revoke_to"
  echo ""

  return 0
}