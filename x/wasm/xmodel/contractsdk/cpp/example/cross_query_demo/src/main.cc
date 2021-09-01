#include "mchain/mchain.h"

struct CrossQueryDemo : public mchain::Contract {};

DEFINE_METHOD(CrossQueryDemo, initialize) {
    mchain::Context* ctx = self.context();
    ctx->ok("initialize succeed");
}

DEFINE_METHOD(CrossQueryDemo, cross_query) {
    mchain::Context* ctx = self.context();
    mchain::Response response;
    ctx->cross_query("xuper://test.xuper?module=wasm&bcname=xuper&contract_name=counter&method_name=get", {{"key", "zq"}}, &response);
    *ctx->mutable_response() = response;   
    ctx->ok("ok");
}
