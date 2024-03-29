#ifndef MCHAIN_CRYPTO_H
#define MCHAIN_CRYPTO_H

#include <string>

namespace mchain {
namespace crypto {
// sha256 returns the sha256 sum of input as bytes
std::string sha256(const std::string& input);
// hex_encode returns the hex encoding of input
std::string hex_encode(const std::string& intput);
// hex_decode returns the hex decoding of input
// if ret false, input is an invalid hex string
bool hex_decode(const std::string& intput, std::string* output);
// ecverify verify elliptic curve signature
// pubkey is the format of json in tx and block
bool ecverify(const std::string& pubkey, const std::string& sign,
              const std::string& hash);
// addr_from_pubkey recover address from public key
bool addr_from_pubkey(const std::string& pubkey, std::string* addr);
}  // namespace crypto
}  // namespace mchain

#endif