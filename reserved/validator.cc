#include "mchain/mchain.h"
#include "mchain/json/json.h"

#define CHECK_ARG(argKey)                             \
    std::string argKey = ctx->arg(#argKey);           \
    if (argKey == "") {                               \
        ctx->error("missing required arg: " #argKey); \
        return;                                       \
    }

std::string Meta(std::string name) { return "M" + name; }

class Validator : public mchain::Contract {};

bool isAdmin(mchain::Context* ctx, const std::string &caller);
// initialize method provisioning contract
// note that creator is important for adding more address into identity list
DEFINE_METHOD(Validator, initialize) {
    mchain::Context* ctx = self.context();
    const std::string& creator = ctx->arg("creator");
    if (creator.empty()) {
        ctx->error("missing creator");
        return;
    }
    ctx->put_object(creator, "true");
    ctx->ok("initialize success");
}

DEFINE_METHOD(Validator, update) {
    mchain::Context* ctx = self.context();
    CHECK_ARG(pub_key);
    CHECK_ARG(power);
    CHECK_ARG(name);

    if (!isAdmin(ctx, ctx->initiator())) {
        ctx->error("only the admin can invoke");
        return;
    }

    mchain::json j;
    j["pub_key"] = pub_key;
    j["power"] = std::atoi(power.c_str());
    j["name"] = name;
    if (j["power"] <= 0) {
        ctx->error("invalid power, it should be greater than 0");
        return;
    }
    auto data = j.dump();
    if (!ctx->put_object(Meta(name), data)) {
        ctx->error("fail to put object");
        return;
    }
    ctx->emit_event("validator_update", data);
    ctx->ok(data);
}

DEFINE_METHOD(Validator, get) {
    mchain::Context* ctx = self.context();
    const std::string key = ctx->arg("name");
    std::string value;
    if (!ctx->get_object(key, &value) || value.empty()) {
        ctx->error("name does not exist");
        return;
    }
    ctx->ok(value);
}

bool isAdmin(mchain::Context* ctx, const std::string &caller) {
    if (caller.empty()) {
        ctx->logf("missing initiator");
        return false;
    }
    std::string admin;
    if (!ctx->get_object(caller, &admin) || admin != "true") {
        return false;
    }
   return true;
}