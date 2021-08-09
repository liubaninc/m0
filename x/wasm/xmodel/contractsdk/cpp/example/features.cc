#include "mchain/json/json.h"
#include "mchain/mchain.h"

struct Features : mchain::Contract {};

DEFINE_METHOD(Features, initialize) {
    mchain::Context* ctx = self.context();
    ctx->ok("init");
}

DEFINE_METHOD(Features, logging) {
    mchain::Context* ctx = self.context();
    ctx->logf("log from contract");
    ctx->ok("ok");
}

DEFINE_METHOD(Features, put) {
    mchain::Context* ctx = self.context();
    for (const auto& elem : ctx->args()) {
        ctx->put_object(elem.first, elem.second);
    }
    ctx->ok("ok");
}

DEFINE_METHOD(Features, get) {
    mchain::Context* ctx = self.context();
    const std::string& key = ctx->arg("key");
    std::string value;
    if (ctx->get_object(key, &value)) {
        ctx->ok(value);
        return;
    }
    ctx->error("failed");
}

DEFINE_METHOD(Features, iterator) {
    mchain::Context* ctx = self.context();
    const std::string& start = ctx->arg("start");
    const std::string& limit = ctx->arg("limit");
    std::string ret;
    auto iter = ctx->new_iterator(start, limit);
    mchain::ElemType elem;
    while (iter->next()) {
        iter->get(&elem);
        ret += elem.first + ":" + elem.second + ", ";
    }
    ctx->ok(ret);
}

DEFINE_METHOD(Features, caller) {
    mchain::Context* ctx = self.context();
    ctx->ok(ctx->sender().get_name());
}

DEFINE_METHOD(Features, call) {
    mchain::Context* ctx = self.context();
    mchain::Response resp;
    const std::string contract = ctx->arg("contract");
    const std::string method = ctx->arg("method");
    bool ret = ctx->call("wasm", contract, method, ctx->args(), &resp);
    if (!ret) {
        ctx->error("call failed");
        return;
    }
    *ctx->mutable_response() = resp;
}

DEFINE_METHOD(Features, json_load_dump) {
    mchain::Context* ctx = self.context();
    const std::string v = ctx->arg("value");
    auto j = mchain::json::parse(v);
    ctx->ok(j.dump());
}

DEFINE_METHOD(Features, json_literal) {
    mchain::Context* ctx = self.context();
    mchain::json j = {
        {"int", 3},
        {"float", 3.14},
        {"string", "hello"},
        {"array", {"hello", "world"}},
        {"object", {{"key", "value"}}},
        {"true", true},
        {"false", false},
        {"null", nullptr},
    };
    ctx->ok(j.dump());
}
