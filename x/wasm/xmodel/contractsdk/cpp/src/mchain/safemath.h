#pragma once

// TODO: this just call the system assert, will implement in near future.
namespace mchain {
    bool safe_assert(bool value) {
        assert(value);
        return true;
    };
}
