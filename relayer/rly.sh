#!/bin/bash

mnemonic="key erupt service six thing spy noise heart giggle year oil fuel rival drop goat deal moral require knee pact bind brain word nuclear"

# 清除之前启动的环境
killall m0d
sleep 10
rm -rf ~/.earth
rm -rf ~/.mars
rm -rf ~/.relayer

# 启动earth区块链
set -e
m0d testnet -n 1 --chain-id earth
mv mytestnet/node0/.m0 ~/.earth
rm -rf mytestnet

echo "$mnemonic" | m0d keys add alice --recover --home ~/.earth
nohup m0d start --pruning nothing --home ~/.earth --log_level debug > ~/.earth/earth.log 2>&1 &

# 启动mars区块链
m0d testnet -n 1 --chain-id mars
mv mytestnet/node0/.m0 ~/.mars
rm -rf mytestnet
# 解决端口冲突
sed -i ' ' 's/:26656/:26658/g' ~/.mars/config/config.toml
sed -i ' ' 's/:26657/:26659/g' ~/.mars/config/config.toml
sed -i ' ' 's/:6060/:6061/g' ~/.mars/config/config.toml
sed -i ' ' 's/:9090/:9091/g' ~/.mars/config/app.toml
sed -i ' ' 's/:1317/:1318/g' ~/.mars/config/app.toml
sed -i ' ' 's/:26660/:26661/g' ~/.mars/config/app.toml

echo "$mnemonic" | m0d keys add alice --recover --home ~/.mars
nohup m0d start --pruning nothing --home ~/.mars --log_level debug > ~/.mars/mars.log 2>&1 &


# 跨链配置
rly config init
rly config add-chains chains
rly config add-paths paths

rly chains edit earth key alice
rly chains edit mars key alice

rly keys restore earth alice "$mnemonic" 
rly keys restore mars alice "$mnemonic"

rly light init earth -f
rly light init mars -f
rly transact link earth-mars

# 监听跨链消息
set +e
while [ 1 ]
do
  rly tx relay-packets earth-mars -d
  rly tx relay-acknowledgements earth-mars -d
done

