#ifndef MCHAIN_TRANSACTION_H
#define MCHAIN_TRANSACTION_H

#include <string>
#include <vector>

namespace mchain {
namespace contract {
namespace sdk {
class Transaction;
}
}  // namespace contract
}  // namespace mchain

namespace mchain {

struct TxInput {
    std::string ref_txid;
    int32_t ref_offset;
    std::string from_addr;
    std::string amount;

    TxInput(const std::string& reftxid, int32_t refoffset,
            const std::string& fromaddr, const std::string& amou)
        : ref_txid(reftxid),
          ref_offset(refoffset),
          from_addr(fromaddr),
          amount(amou) {}
};

struct TxOutput {
    std::string amount;
    std::string to_addr;

    TxOutput(const std::string& amou, const std::string& toaddr)
        : amount(amou), to_addr(toaddr) {}
};

class Transaction {
public:
    Transaction();
    virtual ~Transaction();
    void init(const mchain::contract::sdk::Transaction& pbtx);
    bool from_raw(const std::string& raw_tx);

public:
    std::string txid;
    std::string blockid;
    std::string desc;
    std::string initiator;
    std::vector<std::string> auth_require;
    std::vector<TxInput> tx_inputs;
    std::vector<TxOutput> tx_outputs;
};

}  // namespace mchain

#endif
