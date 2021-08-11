#include "mchain/contract.pb.h"
#include "mchain/block.h"

namespace pb = mchain::contract::sdk;

namespace mchain {

Block::Block() {}

Block::~Block() {}

void Block::init(const pb::Block& pbblock) {
    blockid = pbblock.blockid();
    pre_hash = pbblock.pre_hash();
    proposer = pbblock.proposer();
    sign = pbblock.sign();
    pubkey = pbblock.pubkey();
    height = pbblock.height();
    tx_count = pbblock.tx_count();

    for (int i = 0; i < pbblock.txids_size(); i++) {
        txids.emplace_back(pbblock.txids(i));
    }
}

}  // namespace mchain

