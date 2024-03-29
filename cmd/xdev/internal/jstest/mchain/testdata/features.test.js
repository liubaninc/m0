var assert = require("assert");

function deploy() {
    return mchain.Deploy({
        name: "features",
        code: "./features.wasm",
        lang: "c",
        init_args: {}
    });
}

Test("deploy", function (t) {
    t.Run("file_not_found", function (tt) {
        assert.throws(function () {
            mchain.Deploy({
                name: "features",
                code: "./not_exists.wasm",
                lang: "c",
                init_args: {}
            })
        })
    })

    t.Run("bad_runtime", function (tt) {
        assert.throws(function () {
            mchain.Deploy({
                name: "features",
                code: "./features.wasm",
                lang: "go",
                init_args: {}
            })
        })
    })
    t.Run("ok", function (tt) {
        deploy();
    })
})

Test("put", function (t) {
    var c = deploy();
    c.Invoke("put", { "k1": "v1" });
    resp = c.Invoke("get", { "key": "k1" });
    assert.equal(resp.Body, "v1");
})

Test("get", function (t) {
    var c = deploy();
    t.Run("not_found", function (tt) {
        resp = c.Invoke("get", { "key": "k1" });
        assert.ok(resp.Status != 200);
    })

    t.Run("ok", function (tt) {
        c.Invoke("put", { "k1": "v1" });
        resp = c.Invoke("get", { "key": "k1" });
        assert.equal(resp.Body, "v1");
    })
})

Test("iterator", function (t) {
    var c = deploy();
    t.Run("empty", function (tt) {
        resp = c.Invoke("iterator", { "start": "t_", "limit": "t_\xff" })
        assert.equal(resp.Status, 200);
        assert.equal(resp.Body, "");
    })

    t.Run("ok", function (tt) {
        c.Invoke("put", { "t_k1": "v1", "t_k2": "v2", "t_k3": "v3" });
        resp = c.Invoke("iterator", { "start": "t_", "limit": "t_\xff" })
        assert.equal(resp.Status, 200);
        assert.equal(resp.Body, "t_k1:v1, t_k2:v2, t_k3:v3, ");
    })
})

Test("logging", function (t) {
    var c = deploy();
    c.Invoke("logging", {});
})

Test("call", function (t) {
    t.Run("contract_not_found", function (tt) {
        var c = deploy();
        resp = c.Invoke("call", { "contract": "not_exists" })
        assert.notEqual(resp.Status, 200)
    })

    t.Run("ok", function (tt) {
        c1 = mchain.Deploy({
            name: "c1",
            code: "./features.wasm",
            lang: "c",
            init_args: {}
        });
        c1.Invoke("put", { "k1": "v1" })

        c2 = mchain.Deploy({
            name: "c2",
            code: "./features.wasm",
            lang: "c",
            init_args: {}
        });
        resp = c2.Invoke("call", {
            "contract": "c1",
            "method": "get",
            "key": "k1",
        })
        assert.equal(resp.Body, "v1")
    })
})
