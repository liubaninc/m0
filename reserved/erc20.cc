#include "mchain/mchain.h"

struct ERC20 : public mchain::Contract {};

// 初始化资产 发行指定数量totalSupply的资产到调用者账户
DEFINE_METHOD(ERC20, initialize) {
    mchain::Context* ctx = self.context();
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
    if (atoi(totalSupply.c_str()) <= 0) {
        ctx->error("non positive totalSupply");
        return;
    }

    ctx->put_object("totalSupply", totalSupply);

    std::string key = "balanceOf_" + caller;
    ctx->put_object(key, totalSupply);

    ctx->put_object("master", caller);
    ctx->ok("initialize success");
}

// 增发资产 增发指定数量amount的资产到调用者账户, 仅初始调用者可增发
DEFINE_METHOD(ERC20, mint) {
    mchain::Context* ctx = self.context();
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

    int increaseSupplyint = atoi(increaseSupply.c_str());
    if (increaseSupplyint <= 0) {
        ctx->error("non positive amount");
        return;
    }

    std::string value;
    if (!ctx->get_object("totalSupply", &value)) {
        ctx->error("get totalSupply error");
        return;
    }
    int valueint = atoi(value.c_str());
    int totalSupplyint = increaseSupplyint + valueint;
    char buf[32];
    snprintf(buf, 32, "%d", totalSupplyint);
    ctx->put_object("totalSupply", buf); 
    
    std::string key = "balanceOf_" + caller;
    if (!ctx->get_object(key, &value)) {
        ctx->error("get caller balance error");
        return;
    }
    valueint = atoi(value.c_str());
    int callerint = increaseSupplyint + valueint;
    snprintf(buf, 32, "%d", callerint);
    ctx->put_object(key, buf); 
    
    ctx->ok(buf);
}

// 查询资产发行总量
DEFINE_METHOD(ERC20, totalSupply) {
    mchain::Context* ctx = self.context();
    std::string value;
    if (ctx->get_object("totalSupply", &value)) {
        ctx->ok(value);
    } else {
        ctx->error("key not found");
    }
}

// 查询指定账户owner可用资产数量
DEFINE_METHOD(ERC20, balance) {
    mchain::Context* ctx = self.context();
    const std::string& caller = ctx->arg("owner");
    if (caller.empty()) {
        ctx->error("missing owner");
        return;
    }
    
    std::string key = "balanceOf_" + caller;
    std::string value;
    if (ctx->get_object(key, &value)) {
        ctx->ok(value);
    } else {
        ctx->error("key not found");
    }
}

// 查询to可以从账户from中代理转出token的可用资产数量
DEFINE_METHOD(ERC20, allowance) {
    mchain::Context* ctx = self.context();
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
    std::string value;
    if (ctx->get_object(key, &value)) {
        ctx->ok(value);
    } else {
        ctx->error("key not found");
    }
}

// 资产转移 从调用者账户中往to账户转数量amount的资产，
DEFINE_METHOD(ERC20, transfer) {
    mchain::Context* ctx = self.context();
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

    const std::string& token_str = ctx->arg("amount");
    if (token_str.empty()) {
        ctx->error("missing amount");
        return;
    }
    int token = atoi(token_str.c_str());
    if (token <= 0) {
        ctx->error("non positive amount");
        return;
    }

    std::string from_key = "balanceOf_" + from;
    std::string value;
    int from_balance = 0;
    if (ctx->get_object(from_key, &value)) {
        from_balance = atoi(value.c_str()); 
        if (from_balance < token) {
            ctx->error("The balance of from not enough");
            return;
        }  
    } else {
        ctx->error("key not found");
        return;
    }

    std::string to_key = "balanceOf_" + to;
    int to_balance = 0;
    if (ctx->get_object(to_key, &value)) {
        to_balance = atoi(value.c_str());
    }
   
    from_balance = from_balance - token;
    to_balance = to_balance + token;
   
    char buf[32]; 
    snprintf(buf, 32, "%d", from_balance);
    ctx->put_object(from_key, buf);
    snprintf(buf, 32, "%d", to_balance);
    ctx->put_object(to_key, buf);

    ctx->emit_event("transfer", from + " " + to + " " + token_str);

    ctx->ok("transfer success");
}

// 允许调用者代理某人转移amount数量资产。条件是from账户必须经过了approve
DEFINE_METHOD(ERC20, transferFrom) {
    mchain::Context* ctx = self.context();
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

    const std::string& token_str = ctx->arg("amount");
    if (token_str.empty()) {
        ctx->error("missing amount");
        return;
    }
    int token = atoi(token_str.c_str());
    if (token <= 0) {
        ctx->error("non positive amount");
        return;
    }

    std::string allowance_key = "allowanceOf_" + from + "_" + caller;
    std::string value;
    int allowance_balance = 0;
    if (ctx->get_object(allowance_key, &value)) {
        allowance_balance = atoi(value.c_str()); 
        if (allowance_balance < token) {
            ctx->error("The allowance of from_to not enough");
            return;
        }  
    } else {
        ctx->error("You need to add allowance from_to");
        return;
    }

    std::string from_key = "balanceOf_" + from;
    int from_balance = 0;
    if (ctx->get_object(from_key, &value)) {
        from_balance = atoi(value.c_str()); 
        if (from_balance < token) {
            ctx->error("The balance of from not enough");
            return;
        }  
    } else {
        ctx->error("From no balance");
        return;
    }

    std::string to_key = "balanceOf_" + to;
    int to_balance = 0;
    if (ctx->get_object(to_key, &value)) {
        to_balance = atoi(value.c_str());
    }
   
    from_balance = from_balance - token;
    to_balance = to_balance + token;
    allowance_balance = allowance_balance - token;

    char buf[32]; 
    snprintf(buf, 32, "%d", from_balance);
    ctx->put_object(from_key, buf);
    snprintf(buf, 32, "%d", to_balance);
    ctx->put_object(to_key, buf);
    snprintf(buf, 32, "%d", allowance_balance);
    ctx->put_object(allowance_key, buf);

    ctx->emit_event("transfer", from + " " + to + " " + token_str);

    ctx->ok("transferFrom success");
}

// 批准to能从代理调用账户中转出数量为amount的资产
DEFINE_METHOD(ERC20, approve) {
    mchain::Context* ctx = self.context();
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

    const std::string& token_str = ctx->arg("amount");
    if (token_str.empty()) {
        ctx->error("missing amount");
        return;
    }
    int token = atoi(token_str.c_str());
    if (token <= 0) {
        ctx->error("non positive amount");
        return;
    }

    std::string from_key = "balanceOf_" + from;
    std::string value;
    if (ctx->get_object(from_key, &value)) {
        int from_balance = atoi(value.c_str()); 
        if (from_balance < token) {
            ctx->error("The balance of from not enough");
            return;
        }  
    } else {
        ctx->error("From no balance");
        return;
    }

    std::string allowance_key = "allowanceOf_" + from + "_" + to;
    int allowance_balance = 0;
    if (ctx->get_object(allowance_key, &value)) {
        allowance_balance = atoi(value.c_str()); 
    }

    allowance_balance = allowance_balance + token;
   
    char buf[32]; 
    snprintf(buf, 32, "%d", allowance_balance);
    ctx->put_object(allowance_key, buf);

    ctx->emit_event("approve", from + " " + to + " " + token_str);

    ctx->ok("approve success");
}



