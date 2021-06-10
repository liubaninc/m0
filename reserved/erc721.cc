#include "xchain/xchain.h"

struct ERC721 : public xchain::Contract {};

// split_string split source string into string vector results using delimiter
void string_split(const std::string& source, std::vector<std::string>& results,
                  const std::string& delimiter) {
	std::string::size_type pos1 = 0;
	std::string::size_type pos2 = source.find(delimiter);
	while (std::string::npos != pos2) {
		results.push_back(std::move(source.substr(pos1, pos2 - pos1)));
		pos1 = pos2 + delimiter.size();
		pos2 = source.find(delimiter, pos1);
	}
	if (pos1 != source.length()) {
		results.push_back(source.substr(pos1));
	}
}

// 初始化资产 发行资产NFTs到调用者账户
DEFINE_METHOD(ERC721, initialize) {
    xchain::Context* ctx = self.context();
    const std::string& caller = ctx->initiator();
    if (caller.empty()) {
        ctx->error("missing caller");
        return;
    }
    const std::string& totalSupply = ctx->arg("totalSupply");
    if (totalSupply.empty()) {
        ctx->error("missing totalSupply");
        return;
    }

    std::vector<std::string> value_sets;
	string_split(totalSupply, value_sets, ",");
	for (auto iter = value_sets.begin(); iter != value_sets.end(); ++iter) {
        std::string value = *iter;

        std::string issued;
        std::string supply = "totalSupply_" + value;
        if (ctx->get_object(supply, &issued)) {
            ctx->error(value + "issued");
            return;
        }
        if (!ctx->put_object(supply, value)) {
            ctx->error("failed to put key");
            return;
        }

        std::string balance = "balanceOf_" + caller + "_" +value;
        if (!ctx->put_object(balance, value)){
            ctx->error("failed to put key");
            return;
        }
    }

    ctx->put_object("master", caller);
    ctx->ok("initialize success");
}

// 增发资产 增发资产NFTs到调用者账户, 仅初始调用者可增发
DEFINE_METHOD(ERC721, mint) {
    xchain::Context* ctx = self.context();
    const std::string& caller = ctx->initiator();
    if (caller.empty()) {
        ctx->error("missing caller");
        return;
    }

    std::string master;
    if (!ctx->get_object("master", &master)) {
        ctx->error("missing master");
        return;
    }
    if (master != caller) {
        ctx->error("only the person who created the contract can mint");
        return;
    }

    const std::string& increaseSupply = ctx->arg("amount");
    if (increaseSupply.empty()) {
        ctx->error("missing amount");
        return;
    }

    std::vector<std::string> value_sets;
    string_split(increaseSupply, value_sets, ",");
    for (auto iter = value_sets.begin(); iter != value_sets.end(); ++iter) {
        std::string value = *iter;

        std::string issued;
        std::string supply = "totalSupply_" + value;
        if (ctx->get_object(supply, &issued)) {
            ctx->error(value + "issued");
            return;
        }
        if (!ctx->put_object(supply, value)) {
            ctx->error("failed to put key");
            return;
        }

        std::string balance = "balanceOf_" + caller + "_" +value;
        if (!ctx->put_object(balance, value)){
            ctx->error("failed to put key");
            return;
        }
    }

    ctx->ok("mint success");
}

// 查询NTFs资产发行数量
DEFINE_METHOD(ERC721, totalSupply) {
    xchain::Context* ctx = self.context();

    int total = 0;
     std::string prefix_key = "totalSupply_";
    auto iter = ctx->new_iterator(prefix_key, prefix_key + "~");
    while (iter->next()) {
        std::pair<std::string, std::string> res;
        iter->get(&res);
        total ++;
    }
    char buf[32];
    snprintf(buf, 32, "%d", total);
    ctx->ok(buf);
}

// 查询指定账户owner可用NTFs资产数量
DEFINE_METHOD(ERC721, balance) {
    xchain::Context* ctx = self.context();

    const std::string& caller = ctx->arg("owner");
    if (caller.empty()) {
        ctx->error("missing owner");
        return;
    }

    int total = 0;
    std::string prefix_key = "balanceOf_" + caller + "_";
    auto iter = ctx->new_iterator(prefix_key, prefix_key + "~");
    while (iter->next()) {
        total ++;
    }
    char buf[32];
    snprintf(buf, 32, "%d", total);
    ctx->ok(buf);
}

// 查询to可以从账户from中代理转出NTFs的可用资产数量
DEFINE_METHOD(ERC721, allowance) {
    xchain::Context* ctx = self.context();
    const std::string& from = ctx->arg("from");
    if (from.empty()) {
        ctx->error("missing from");
        return;
    }
   
    const std::string& to = ctx->arg("to");
    if (to.empty()) {
        ctx->error("missing to");
        return;
    }

    std::string key = "allowanceOf_" + from + "_" + to;
    int total = 0;
    std::string prefix_key = "allowanceOf_" + from + "_" + to + "_";
    auto iter = ctx->new_iterator(prefix_key, prefix_key + "~");
    while (iter->next()) {
        total ++;
    }
    char buf[32];
    snprintf(buf, 32, "%d", total);
    ctx->ok(buf);
}

// 资产转移 从调用者账户中往to账户转出NFT资产，
DEFINE_METHOD(ERC721, transfer) {
    xchain::Context* ctx = self.context();
    const std::string& from = ctx->initiator();
    if (from.empty()) {
        ctx->error("missing from");
        return;
    }
   
    const std::string& to = ctx->arg("to");
    if (to.empty()) {
        ctx->error("missing to");
        return;
    }

    const std::string& token_str = ctx->arg("id");
    if (token_str.empty()) {
        ctx->error("missing token id");
        return;
    }

    std::string from_key = "balanceOf_" + from + "_" + token_str;
    std::string value;
    if (ctx->get_object(from_key, &value)) {

    } else {
        ctx->error("token id not found in from");
        return;
    }
    std::string to_key = "balanceOf_" + to + "_" + token_str;

    if (!ctx->delete_object(from_key)) {
        ctx->error("failed to delete key:" + from_key);
        return;
    }
    if (!ctx->put_object(to_key, token_str)) {
        ctx->error("failed to put key:" + to_key);
        return;
    }
    ctx->emit_event("transfer", from + " " + to + " " + token_str);
    ctx->ok("transfer success");
}

// 允许调用者代理某人转移NFT资产。条件是from账户必须经过了approve
DEFINE_METHOD(ERC721, transferFrom) {
    xchain::Context* ctx = self.context();
    const std::string& from = ctx->arg("from");
    if (from.empty()) {
        ctx->error("missing from");
        return;
    }
  
    const std::string& caller = ctx->initiator();
    if (caller.empty()) {
        ctx->error("missing caller");
        return;
    }

    const std::string& to = ctx->arg("to");
    if (to.empty()) {
        ctx->error("missing to");
        return;
    }

    const std::string& token_str = ctx->arg("id");
    if (token_str.empty()) {
        ctx->error("missing token id");
        return;
    }

    std::string allowance_key = "allowanceOf_" + from + "_" + caller + "_" + token_str;
    std::string value;
    if (ctx->get_object(allowance_key, &value)) {

    } else {
        ctx->error("You need to add allowance from_to");
        return;
    }

    std::string from_key = "balanceOf_" + from + "_" + token_str;;
    if (ctx->get_object(from_key, &value)) {

    } else {
        ctx->error("From no balance");
        return;
    }

    std::string to_key = "balanceOf_" + to + "_" + token_str;

    if (!ctx->delete_object(allowance_key)) {
        ctx->error("failed to delete key:" + allowance_key);
        return;
    }
    if (!ctx->delete_object(from_key)) {
        ctx->error("failed to delete key:" + from_key);
        return;
    }
    if (!ctx->put_object(to_key, token_str)) {
        ctx->error("failed to put key:" + to_key);
        return;
    }

    ctx->emit_event("transfer", from + " " + to + " " + token_str);
    ctx->ok("transferFrom success");
}

// 批准to能从代理调用账户中转出NFT的资产
DEFINE_METHOD(ERC721, approve) {
    xchain::Context* ctx = self.context();
    const std::string& from = ctx->initiator();
    if (from.empty()) {
        ctx->error("missing from");
        return;
    }
   
    const std::string& to = ctx->arg("to");
    if (to.empty()) {
        ctx->error("missing to");
        return;
    }

    const std::string& token_str = ctx->arg("id");
    if (token_str.empty()) {
        ctx->error("missing token id");
        return;
    }

    std::string from_key = "balanceOf_" + from + "_" + token_str;;
    std::string value;
    if (ctx->get_object(from_key, &value)) {

    } else {
        ctx->error("From no balance");
        return;
    }

    std::string allowance_key = "allowanceOf_" + from + "_" + to + "_" + token_str;;
    if (!ctx->put_object(allowance_key, token_str)) {
        ctx->error("failed to put key:" + allowance_key);
        return;
    }

    ctx->emit_event("approve", from + " " + to + " " + token_str);
    ctx->ok("approve success");
}

DEFINE_METHOD(ERC721, approveAll) {
    xchain::Context* ctx = self.context();
    const std::string& from = ctx->initiator();
    if (from.empty()) {
        ctx->error("missing from");
        return;
    }
   
    const std::string& to = ctx->arg("to");
    if (to.empty()) {
        ctx->error("missing to");
        return;
    }

    std::string prefix_key = "balanceOf_" + from + "_";
    auto iter = ctx->new_iterator(prefix_key, prefix_key + "~");
    while (iter->next()) {
        std::pair<std::string, std::string> res;
        iter->get(&res);
        std::string token_str = res.second;
        std::string allowance_key = "allowanceOf_" + from + "_" + to + "_" + token_str;;
        if (!ctx->put_object(allowance_key, token_str)) {
            ctx->error("failed to put key:" + allowance_key);
            return;
        }
    }

    ctx->emit_event("approveAll", from + " " + to + " " + "true");
    ctx->ok("approve success");
}


