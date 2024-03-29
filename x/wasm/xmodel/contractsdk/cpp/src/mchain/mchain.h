#ifndef MCHAIN_MCHAIN_H
#define MCHAIN_MCHAIN_H

#include <map>
#include <memory>
#include <string>
#include <vector>
#include "mchain/account.h"
#include "mchain/basic_iterator.h"
#include "mchain/block.h"
#include "mchain/transaction.h"

namespace mchain {

struct Response {
    int status;
    std::string message;
    std::string body;
};

const std::string kUnknownKey = "";

class Context {
public:
    virtual ~Context() {}
    virtual const std::map<std::string, std::string>& args() const = 0;
    virtual const std::string& arg(const std::string& name) const = 0;
    virtual const std::string& initiator() const = 0;
    virtual int auth_require_size() const = 0;
    virtual const std::string& auth_require(int idx) const = 0;
    virtual bool get_object(const std::string& key, std::string* value) = 0;
    virtual bool put_object(const std::string& key,
                            const std::string& value) = 0;
    virtual bool delete_object(const std::string& key) = 0;
    virtual bool query_tx(const std::string& txid, Transaction* tx) = 0;
    virtual bool query_block(const int64_t& blockid, Block* block) = 0;
    virtual void ok(const std::string& body) = 0;
    virtual void error(const std::string& body) = 0;
    virtual Response* mutable_response() = 0;
    virtual std::unique_ptr<Iterator> new_iterator(
        const std::string& start, const std::string& limit) = 0;
    virtual Account& sender() = 0;
    virtual const std::string& transfer_amount() const = 0;
    virtual bool call(const std::string& module, const std::string& contract,
                      const std::string& method,
                      const std::map<std::string, std::string>& args,
                      Response* response) = 0;
    virtual bool cross_query(const std::string& uri, 
                        const std::map<std::string, std::string>& args,
                        Response* response) = 0;                  
    virtual void logf(const char* fmt, ...) = 0;
    virtual bool emit_event(const std::string& name, const std::string& body) = 0;
};

class Contract {
public:
    Contract();
    virtual ~Contract();
    Context* context() { return _ctx; };

private:
    Context* _ctx;
};

}  // namespace mchain

#define DEFINE_METHOD(contract_class, method_name)        \
    static void cxx_##method_name(contract_class&);       \
    extern "C" void __attribute__((used)) method_name() { \
        contract_class self;                              \
        cxx_##method_name(self);                          \
    };                                                    \
    static void cxx_##method_name(contract_class& self)

#endif
