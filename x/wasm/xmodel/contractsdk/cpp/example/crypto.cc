#include "mchain/crypto.h"
#include "mchain/mchain.h"

using namespace mchain::crypto;

struct Crypto : public mchain::Contract {};

DEFINE_METHOD(Crypto, initialize) {}

DEFINE_METHOD(Crypto, sha256) {
    mchain::Context* ctx = self.context();
    const std::string& in = ctx->arg("in");
    std::string out = sha256(in);
    ctx->ok(hex_encode(out));
}

DEFINE_METHOD(Crypto, ecverify) {
    mchain::Context* ctx = self.context();
    const std::string& hash_hex = ctx->arg("hash");
    std::string hash;
    hex_decode(hash_hex, &hash);
    const std::string& sign_hex = ctx->arg("sign");
    std::string sign;
    hex_decode(sign_hex, &sign);
    const std::string& pubkey = ctx->arg("pubkey");

    bool ok = ecverify(pubkey, sign, hash);
    if (!ok) {
        ctx->error("fail");
        return;
    }
    ctx->ok("ok");
}

DEFINE_METHOD(Crypto, addr_from_pubkey) {
    mchain::Context* ctx = self.context();
    const std::string& pubkey = ctx->arg("pubkey");
    std::string addr;
    if (!addr_from_pubkey(pubkey, &addr)) {
        ctx->error("fail");
        return;
    }
    ctx->ok(addr);
}