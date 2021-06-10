#include <string>
#include <vector>
#include "xchain/xchain.h"

class UnifiedCheck : public xchain::Contract {};

// define delimiters  used for dividing args
const std::string DELIMITER_COMMA = ",";

// prefix for different types
const std::string PREFIX_IDENTITY = "ID_";
const std::string PREFIX_ADDRESS = "ADDRESS_";
const std::string PREFIX_CONTRACT = "CONTRACT_";
const std::string PREFIX_PEER = "PEER_";

bool isAdmin(xchain::Context* ctx, const std::string &caller);
// return 0 if register pass, otherwise register fail
int register1(xchain::Context* ctx, const std::string& key, const std::string& prefix);
// return 0 if unregister pass, otherwise unregister fail
int unregister(xchain::Context* ctx, const std::string& key, const std::string& prefix);
// return 0 if registered, otherwise unregistered
int registered(xchain::Context* ctx, const std::string& key, const std::string& prefix);
// return 0 if on, otherwise off
bool on_off(xchain::Context* ctx);

// initialize method provisioning contract
// note that creator is important for adding more address into identity list
DEFINE_METHOD(UnifiedCheck, initialize) {
	xchain::Context* ctx = self.context();
	const std::string& creator = ctx->arg("creator");
	if (creator.empty()) {
		ctx->error("missing creator");
		return;
	}
//	if (!ctx->put_object(PREFIX_IDENTITY + creator, "true")) {
//		ctx->error("put creator failed");
//		return;
//	}
	ctx->put_object(creator, "true");
	ctx->ok("success");
}

// register_address method register addresses to contract
DEFINE_METHOD(UnifiedCheck, register_address) {
	xchain::Context* ctx = self.context();
	if (!isAdmin(ctx, ctx->initiator())) {
        ctx->error("only the admin can invoke");
        return;
    }
	if (register1(ctx, "address", PREFIX_ADDRESS) != 0) {
		ctx->error("register address to contract error");
		return;
	}
	ctx->ok("success");
}
// unregister_address method unregister addresses from contract
DEFINE_METHOD(UnifiedCheck, unregister_address) {
	xchain::Context* ctx = self.context();
	if (!isAdmin(ctx, ctx->initiator())) {
        ctx->error("only the admin can invoke");
        return;
    }
	if (unregister(ctx, "address", PREFIX_ADDRESS) != 0) {
		ctx->error("unregister address to contract error");
		return;
	}
	ctx->ok("success");
}

// register_contract method register contracts to contract
DEFINE_METHOD(UnifiedCheck, register_contract) {
	xchain::Context* ctx = self.context();
	if (!isAdmin(ctx, ctx->initiator())) {
        ctx->error("only the admin can invoke");
        return;
    }
	if (register1(ctx, "contract", PREFIX_CONTRACT) != 0) {
		ctx->error("register contract to contract error");
		return;
	}
	ctx->ok("success");
}
// unregister_contract method unregister contracts to contract
DEFINE_METHOD(UnifiedCheck, unregister_contract) {
	xchain::Context* ctx = self.context();
    if (!isAdmin(ctx, ctx->initiator())) {
        ctx->error("only the admin can invoke");
        return;
    }
	if (unregister(ctx, "contract", PREFIX_CONTRACT) != 0) {
		ctx->error("unregister contract to contract error");
		return;
	}
	ctx->ok("success");
}

// register_peer method register peers to contract
DEFINE_METHOD(UnifiedCheck, register_peer) {
	xchain::Context* ctx = self.context();
    if (!isAdmin(ctx, ctx->initiator())) {
        ctx->error("only the admin can invoke");
        return;
    }
	if (register1(ctx, "peer", PREFIX_PEER) != 0) {
		ctx->error("register peer to contract error");
		return;
	}
	ctx->ok("success");
}
// unregister_peer method register peers to contract
DEFINE_METHOD(UnifiedCheck, unregister_peer) {
	xchain::Context* ctx = self.context();
	if (!isAdmin(ctx, ctx->initiator())) {
        ctx->error("only the admin can invoke");
        return;
    }
	if (unregister(ctx, "peer", PREFIX_PEER) != 0) {
		ctx->error("unregister peer to contract error");
		return;
	}
	ctx->ok("success");
}

// register_id method register ids to contract
DEFINE_METHOD(UnifiedCheck, register_id) {
	xchain::Context* ctx = self.context();
	if (!isAdmin(ctx, ctx->initiator())) {
        ctx->error("only the admin can invoke");
        return;
    }
	if (register1(ctx, "id", PREFIX_IDENTITY) != 0) {
		ctx->error("register id to contract error");
		return;
	}
	ctx->ok("success");
}
// unregister_id method register ids to contract
DEFINE_METHOD(UnifiedCheck, unregister_id) {
	xchain::Context* ctx = self.context();
	if (!isAdmin(ctx, ctx->initiator())) {
        ctx->error("only the admin can invoke");
        return;
    }
	if (unregister(ctx, "id", PREFIX_IDENTITY) != 0) {
		ctx->error("unregister id to contract error");
		return;
	}
	ctx->ok("success");
}

// register method register to contract
DEFINE_METHOD(UnifiedCheck, register1) {
	xchain::Context* ctx = self.context();
	if (!isAdmin(ctx, ctx->initiator())) {
        ctx->error("only the admin can invoke");
        return;
    }
    if (ctx->arg("address").empty() &&
        ctx->arg("contract").empty() &&
        ctx->arg("peer").empty() &&
        ctx->arg("id").empty()) {
        ctx->error("not support type");
        return;
    }
	if (register1(ctx, "address", PREFIX_ADDRESS) != 0) {
		ctx->error("register address to contract error");
		return;
	}
	if (register1(ctx, "contract", PREFIX_CONTRACT) != 0) {
		ctx->error("register contract to contract error");
		return;
	}
	if (register1(ctx, "peer", PREFIX_PEER) != 0) {
		ctx->error("register peer to contract error");
		return;
	}
	if (register1(ctx, "id", PREFIX_IDENTITY) != 0) {
		ctx->error("register id to contract error");
		return;
	}
	ctx->ok("success");
}
// unregister method register to contract
DEFINE_METHOD(UnifiedCheck, unregister) {
	xchain::Context* ctx = self.context();
	if (!isAdmin(ctx, ctx->initiator())) {
        ctx->error("only the admin can invoke");
        return;
    }
	if (ctx->arg("address").empty() &&
        ctx->arg("contract").empty() &&
        ctx->arg("peer").empty() &&
        ctx->arg("id").empty()) {
        ctx->error("not support type");
        return;
    }
	if (unregister(ctx, "address", PREFIX_ADDRESS) != 0) {
		ctx->error("unregister address to contract error");
		return;
	}
	if (unregister(ctx, "contract", PREFIX_CONTRACT) != 0) {
		ctx->error("unregister contract to contract error");
		return;
	}
	if (unregister(ctx, "peer", PREFIX_PEER) != 0) {
		ctx->error("unregister peer to contract error");
		return;
	}
	if (unregister(ctx, "id", PREFIX_IDENTITY) != 0) {
		ctx->error("unregister id to contract error");
		return;
	}
	ctx->ok("success");
}

// register method register to contract
DEFINE_METHOD(UnifiedCheck, on) {
    xchain::Context* ctx = self.context();
    std::string creator;
    if (!ctx->get_object("creator", &creator)) {
         ctx->error("get creator fail");
         return;
    }
    if (creator != ctx->initiator()) {
        ctx->error("only creator can invoke");
        return;
    }
    if (ctx->put_object("on", "true")) {
        ctx->ok("success");
    } else {
        ctx->error("fail");
    }
}

// register method register to contract
DEFINE_METHOD(UnifiedCheck, off) {
    xchain::Context* ctx = self.context();
    if (ctx->delete_object("on")) {
        ctx->ok("success");
    } else {
        ctx->error("fail");
    }
}

// identity_check return if is identity
// keep this method for convenience
DEFINE_METHOD(UnifiedCheck, identity_check) {
	xchain::Context* ctx = self.context();
	if (!on_off(ctx)) {
        ctx->ok("success");
        return;
	}

	if (ctx->arg("address").empty() &&
	    ctx->arg("contract").empty() &&
	    ctx->arg("peer").empty()) {
		ctx->error("not support type");
		return;
	}
	if (registered(ctx, "address", PREFIX_ADDRESS) == 0) {
		ctx->ok("success");
		return;
	}
	if (registered(ctx, "contract", PREFIX_CONTRACT) == 0) {
		ctx->ok("success");
		return;
	}
	if (registered(ctx, "peer", PREFIX_PEER) == 0) {
		ctx->ok("success");
		return;
	}
	ctx->error("not found");
}
// banned_check return if is banned
// keep this method for convenience
DEFINE_METHOD(UnifiedCheck, banned_check) {
	xchain::Context* ctx = self.context();
//	if (!on_off(ctx)) {
//        ctx->ok("success");
//        return;
//    }
	if (registered(ctx, "address", PREFIX_ADDRESS) == 0) {
		ctx->ok("success");
		return;
	}
	if (registered(ctx, "contract", PREFIX_CONTRACT) == 0) {
		ctx->ok("success");
		return;
	}
	if (registered(ctx, "peer", PREFIX_PEER) == 0) {
		ctx->ok("success");
		return;
	}
	ctx->error("not found");
}

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

bool isAdmin(xchain::Context* ctx, const std::string &caller) {
    if (caller.empty()) {
        ctx->logf("missing caller");
        return false;
    }
    std::string admin;
    if (!ctx->get_object(caller, &admin) || admin != "true") {
        return false;
    }
   return true;
}

bool on_off(xchain::Context* ctx) {
    std::string on;
    if (!ctx->get_object("on", &on) || on != "true") {
        return false;
    }
    return true;
}

// registered is registered return 0 if registered, otherwise unregistered
int registered(xchain::Context* ctx, const std::string& key, const std::string& prefix) {
	const std::string values = ctx->arg(key);
	std::vector<std::string> value_sets;
	string_split(values, value_sets, DELIMITER_COMMA);
	std::string val;
	// one of contracts has been banned, return directly
	for (auto iter = value_sets.begin(); iter != value_sets.end(); ++iter) {
		std::string value = prefix + (*iter);
		if (ctx->get_object(value, &val) && val == "true") {
			return 0;
		}
	}
	return -1;
}
// key register to identity contract
// return 0 if register pass, otherwise register fail
int register1(xchain::Context* ctx, const std::string& key, const std::string& prefix) {
	const std::string values = ctx->arg(key);
	std::vector<std::string> value_sets;
	string_split(values, value_sets, DELIMITER_COMMA);
	for (auto iter = value_sets.begin(); iter != value_sets.end(); ++iter) {
		std::string value = prefix + *iter;
		if (!ctx->put_object(value, "true")) {
			return -1;
		}
	}
	return 0;
}
// key unregister to identity contract
// return 0 if unregister pass, otherwise unregister fail
int unregister(xchain::Context* ctx,
                const std::string& key,
                const std::string& prefix) {
	const std::string values = ctx->arg(key);
	std::vector<std::string> value_sets;
	string_split(values, value_sets, DELIMITER_COMMA);
	for (auto iter = value_sets.begin(); iter != value_sets.end(); ++iter) {
		std::string value = prefix + *iter;
		if (!ctx->delete_object(value)) {
			return -1;
		}
	}
	return 0;
}