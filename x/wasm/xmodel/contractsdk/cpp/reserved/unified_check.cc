#include <string>
#include <vector>
#include "mchain/mchain.h"

class UnifiedCheck : public mchain::Contract {};

// define delimiters
const std::string DELIMITER_COMMA = ",";  // used for dividing contract args
const std::string DELIMITER_SLASH = "/";  // used for checking account/ak

// prefix for different types
const std::string PREFIX_IDENTITY = "ID_";
const std::string PREFIX_BANNED_CONTRACT = "BAN_";

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

// string_last return the substring from last delimiter to the end.
// if no delimiter found, return source
std::string string_last(const std::string& source,
                        const std::string& delimiter) {
    std::size_t found = source.rfind(delimiter);
    if (found != std::string::npos) {
        return source.substr(found + 1, std::string::npos);
    }
    return source;
}

int verify_address(mchain::Context* ctx, std::string address) {
    std::string value;
    std::string key = PREFIX_IDENTITY + address;
    if (!ctx->get_object(key, &value) || value != "true") {
        return -1;
    }
    return 0;
}

// verify if initiator and auth_require are in identify list
// return 0 if check pass, otherwise check fail
int verify_identity(mchain::Context* ctx) {
    // verify initiator
    mchain::Account initiator(ctx->initiator());
    if (initiator.type() == mchain::ADDRESS) {
        // initiator is address
        if (verify_address(ctx, initiator.get_name()) != 0) {
            return -1;
        }
    } else {
        // initiator is account
        std::vector<std::string> addresses;
        if (!initiator.get_addresses(&addresses)) {
            return -1;
        }
        for (int i = 0; i < addresses.size(); i++) {
            if (verify_address(ctx, addresses[i]) != 0) {
                return -1;
            }
        }
    }

    int auth_require_size = ctx->auth_require_size();
    std::vector<std::string> accounts;
    for (int iter = 0; iter < auth_require_size; ++iter) {
        std::string ak_name =
            string_last(ctx->auth_require(iter), DELIMITER_SLASH);
        if (verify_address(ctx, ak_name) != 0) {
            return -1;
        }
    }
    return 0;
}

// verify is one of the contracts is banned
// return 0 if not banned, otherwise check fail
int verify_banned(mchain::Context* ctx) {
    const std::string keys = ctx->arg("contract");

    std::vector<std::string> contracts;
    string_split(keys, contracts, DELIMITER_COMMA);

    std::string value;
    // one of contracts has been banned, return directly
    for (auto iter = contracts.begin(); iter != contracts.end(); ++iter) {
        std::string banned_name = PREFIX_BANNED_CONTRACT + (*iter);
        if (ctx->get_object(banned_name, &value) && value == "true") {
            return 1;
        }
    }
    return 0;
}

// initialize method provisioning contract
// note that creator is important for adding more address into identity list
DEFINE_METHOD(UnifiedCheck, initialize) {
    mchain::Context* ctx = self.context();
    const std::string& creator = ctx->arg("creator");
    if (creator.empty()) {
        ctx->error("missing creator");
        return;
    }
    if (!ctx->put_object(PREFIX_IDENTITY + creator, "true")) {
        ctx->error("put creator failed");
        return;
    }
    ctx->ok("success");
}

//////// identity contract write method ////////
// register_aks method register aks to identify contract
DEFINE_METHOD(UnifiedCheck, register_aks) {
    mchain::Context* ctx = self.context();

    // aks register to identity contract
    const std::string aks = ctx->arg("aks");
    std::vector<std::string> ak_sets;
    string_split(aks, ak_sets, DELIMITER_COMMA);

    for (auto iter = ak_sets.begin(); iter != ak_sets.end(); ++iter) {
        std::string ak_name = PREFIX_IDENTITY + *iter;
        if (!ctx->put_object(ak_name, "true")) {
            ctx->error("register aks to identify contract error");
            return;
        }
    }

    ctx->ok("success");
}

// unregister_aks method unregister aks from identify contract
DEFINE_METHOD(UnifiedCheck, unregister_aks) {
    mchain::Context* ctx = self.context();

    // aks unregister form identity contract
    const std::string aks = ctx->arg("aks");
    std::vector<std::string> ak_sets;
    string_split(aks, ak_sets, DELIMITER_COMMA);

    for (auto iter = ak_sets.begin(); iter != ak_sets.end(); ++iter) {
        std::string ak_name = PREFIX_IDENTITY + (*iter);
        if (!ctx->delete_object(ak_name)) {
            ctx->error("unregister from identify contract error");
            return;
        }
    }

    ctx->ok("success");
}

//////// banned contract write method ////////
// ban could add contracts to banned list
DEFINE_METHOD(UnifiedCheck, ban) {
    mchain::Context* ctx = self.context();
    const std::string keys = ctx->arg("contract");
    const std::string value = "true";

    std::vector<std::string> contracts;
    string_split(keys, contracts, DELIMITER_COMMA);

    for (auto iter = contracts.begin(); iter != contracts.end(); ++iter) {
        std::string banned_name = PREFIX_BANNED_CONTRACT + (*iter);
        if (!ctx->put_object(banned_name, value)) {
            ctx->error("ban contract failed");
            return;
        }
    }
    ctx->ok("success");
}

// unban could remove contracts from banned list
DEFINE_METHOD(UnifiedCheck, unban) {
    mchain::Context* ctx = self.context();
    const std::string keys = ctx->arg("contract");

    std::vector<std::string> contracts;
    string_split(keys, contracts, DELIMITER_COMMA);

    for (auto iter = contracts.begin(); iter != contracts.end(); ++iter) {
        std::string banned_name = PREFIX_BANNED_CONTRACT + (*iter);
        if (!ctx->delete_object(banned_name)) {
            ctx->error("release failed");
            return;
        }
    }
    ctx->ok("success");
}

//////// unified verify method ////////
// verify method verify whether the aks were identified
DEFINE_METHOD(UnifiedCheck, verify) {
    mchain::Context* ctx = self.context();
    if (verify_identity(ctx) != 0) {
        ctx->error("identity check failed");
        return;
    }

    if (verify_banned(ctx) != 0) {
        ctx->error("banned contract check failed");
        return;
    }

    ctx->ok("success");
}

// identity_check return if the initiator and auth_requires are in identiy list
// keep this method for convenience
DEFINE_METHOD(UnifiedCheck, identity_check) {
    mchain::Context* ctx = self.context();
    if (verify_identity(ctx) != 0) {
        ctx->error("identity check failed");
        return;
    }

    ctx->ok("success");
}

// identity_query return whether address is in identity list
// keep this method for convenience
DEFINE_METHOD(UnifiedCheck, identity_query) {
    mchain::Context* ctx = self.context();
    const std::string& address = ctx->arg("address");
    if (address.empty()) {
        ctx->error("missing address");
        return;
    }

    if (verify_address(ctx, address) == 0) {
        ctx->ok("Found");
        return;
    }

    ctx->ok("Not found");
}

// banned_check return if the contract is banned
// keep this method for convenience
DEFINE_METHOD(UnifiedCheck, banned_check) {
    mchain::Context* ctx = self.context();
    if (verify_banned(ctx) != 0) {
        ctx->error("banned contract check failed");
        return;
    }

    ctx->ok("success");
}
