#include "mchain/crypto.h"
#include <string>

extern "C" void xvm_hash(const char* name, const char* inputptr, int inputlen,
                         char* outputptr, int outputlen);
extern "C" void xvm_encode(const char* name, const char* inputptr, int inputlen,
                           char** outputpptr, int* outlen);
extern "C" int xvm_decode(const char* name, const char* inputptr, int inputlen,
                          char** outputpptr, int* outlen);
extern "C" int xvm_ecverify(const char* pubptr, int publen, const char* sigptr,
                            int siglen, const char* hashptr, int hashlen);
extern "C" char* xvm_addr_from_pubkey(const char* pubptr, int publen);

namespace mchain {
namespace crypto {
std::string sha256(const std::string& input) {
    char out[32];
    xvm_hash("sha256", (const char*)&input[0], input.size(), out, sizeof(out));
    return std::string(out, sizeof(out));
}

std::string hex_encode(const std::string& input) {
    char* out = NULL;
    int outlen = 0;
    xvm_encode("hex", (const char*)&input[0], input.size(), &out, &outlen);
    std::string ret(out, outlen);
    free(out);
    return ret;
}

bool hex_decode(const std::string& input, std::string* output) {
    char* out = NULL;
    int outlen = 0;
    int ret = 0;
    ret =
        xvm_decode("hex", (const char*)&input[0], input.size(), &out, &outlen);
    if (ret != 0) {
        return false;
    }
    output->assign(out, outlen);
    free(out);
    return true;
}

bool ecverify(const std::string& pubkey, const std::string& sign,
              const std::string& hash) {
    int ret = 0;
    ret = xvm_ecverify((const char*)&pubkey[0], pubkey.size(),
                       (const char*)&sign[0], sign.size(),
                       (const char*)&hash[0], hash.size());
    if (ret != 0) {
        return false;
    }
    return true;
}

bool addr_from_pubkey(const std::string& pubkey, std::string* addr) {
    char* out = NULL;
    out = xvm_addr_from_pubkey((const char*)&pubkey[0], pubkey.size());
    if (out == NULL) {
        return false;
    }
    addr->assign(out);
    free(out);
    return true;
}

}  // namespace crypto
}  // namespace mchain