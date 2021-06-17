#!/bin/sh

cd /vue/browser && npm install

cd /vue/wallet && npm install

cd /vue/browser && npm run serve &
cd /vue/wallet && npm run serve &

set -ex

synced start