package com.tq_bc.m0;

import cn.hutool.core.codec.Base64;
import cn.hutool.core.io.FileUtil;
import cn.hutool.core.util.StrUtil;
import cn.hutool.json.JSONArray;
import cn.hutool.json.JSONObject;
import cn.hutool.json.JSONUtil;
import com.google.common.base.Charsets;
import com.google.common.collect.Lists;
import com.tq_bc.m0.api.Api;
import com.tq_bc.m0.api.ResultABCIQuery;
import com.tq_bc.m0.api.ResultBroadcastTx;
import com.tq_bc.m0.api.ResultBroadcastTxCommit;
import com.tq_bc.m0.common.Crypto;
import liubaninc.m0.utxo.InputOuterClass.*;
import liubaninc.m0.utxo.QueryOuterClass.*;
import com.google.protobuf.Any;
import liubaninc.m0.utxo.Tx.*;
import cosmos.base.v1beta1.CoinOuterClass.Coin;
import lombok.Builder;
import lombok.Data;
import org.bitcoinj.core.ECKey;
import org.bouncycastle.util.encoders.Hex;

import java.util.List;
import java.util.Map;

@Data
@Builder
public class Transaction {
    private Api client;
    private String memo;
    private boolean commit;
    public static final String address_prefix = "mk";

    private String tx(ECKey key, List<Any> msgs, List<Coin> fees) throws Exception {
        String from = Crypto.getAddressFromECKey(key, address_prefix);
        StdFee.Builder std_fee_builder = StdFee.newBuilder().setGas(0);
        for (Coin fee: fees) {
            std_fee_builder.addAmount(M0.Coin.newBuilder().setDenom(fee.denom).setAmount(fee.amount));
        }

        JSONArray msgs = JSONUtil.createArray();
        msgs.add(JSONUtil.parseObj(msg));


        JSONObject value =  JSONUtil.createObj();
        value.set("fee", JSONUtil.parseObj(Api.deserializer.toJson(std_fee_builder.build())));
        value.set("signatures", JSONUtil.parseArray("[{}]"));
        value.set("memo", memo);
        value.set("msg", msgs);

        JSONObject tx = JSONUtil.createObj();
        tx.set("type", "cosmos-sdk/StdTx");
        tx.set("value", value);

        ResultABCIQuery result_tx_encode = ResultABCIQuery.GetABCIQuery(client, "custom/utxo/encode", tx.toString().getBytes(),0l, false);

        ResultABCIQuery result_simulate_tx = ResultABCIQuery.GetABCIQuery(client, "/app/simulate2",  Hex.decode(result_tx_encode.response.getValue().toStringUtf8()),0l, false);

        JSONObject simulate = JSONUtil.parseObj(result_simulate_tx.response.getValue().toStringUtf8());
        String gasused = simulate.getJSONObject("GasInfo").getStr("GasUsed");
        std_fee_builder.setGas(Long.valueOf(gasused));
        value.set("fee", JSONUtil.parseObj(Api.deserializer.toJson(std_fee_builder.build())));

        Account acct = Account.Get(client, from);
        SignData sd = SignData.builder().
                chain_id(client.chain_id).
                account_number(acct.account_number).
                sequence("0").
                memo(memo).
                fee(value.getJSONObject("fee")).
                msgs(value.getJSONArray("msg")).build();

        ResultABCIQuery result_sign_data = ResultABCIQuery.GetABCIQuery(client, "custom/utxo/sort", JSONUtil.toJsonStr(sd).getBytes(Charsets.UTF_8),0l, false);

        byte[] signatureBytes = Crypto.sign(result_sign_data.response.getValue().toByteArray(), key);
        JSONObject signature = JSONUtil.createObj();
        JSONObject pubkey = JSONUtil.createObj();
        pubkey.set("type", "tendermint/PubKeySecp256k1");
        pubkey.set("value", Base64.encode(key.getPubKey()));
        signature.set("pub_key", pubkey);
        signature.set("signature", Base64.encode(signatureBytes));

        JSONArray signatures = JSONUtil.createArray();
        signatures.add(signature);
        value.set("signatures", signatures);

        ResultABCIQuery result_tx = ResultABCIQuery.GetABCIQuery(client, "custom/utxo/encode", tx.toString().getBytes(),0l, false);
        if (commit) {
            ResultBroadcastTxCommit result = ResultBroadcastTxCommit.BroadcastTxCommit(client,"0x" + result_tx.response.getValue().toStringUtf8());
            if (result.check_tx.getCode() != 0) {
                throw new RuntimeException(result.check_tx.getLog());
            }
            return result.hash;
        } else {
            ResultBroadcastTx result = ResultBroadcastTx.BroadcastTxSync(client,"0x" + result_tx.response.getValue().toStringUtf8());
            if (result.code != 0) {
                throw new RuntimeException(result.log);
            }
            return result.hash;
        }
    }

    public void fillInputsAndOutputs(String address, List<Coin> amounts, List<Input> inputs, List<Output> outputs) throws Exception {
        // TODO
    }

    public String Mint(ECKey key, Coin amount, String desc, Coin fee) throws Exception {
        String from = Crypto.getAddressFromECKey(key, address_prefix);

        List<Input> inputs = Lists.newArrayList();
        List<Output> outputs = Lists.newArrayList();
        // 支出
        List<Coin> total = Lists.newArrayList();
        total.add(fee);
        fillInputsAndOutputs(from, total, inputs, outputs);
        // 收入
        outputs.add(Output.newBuilder()
                .setToAddr(from)
                .setAmount(amount)
                .build());
        MsgIssue.Builder builder = MsgIssue.newBuilder();
        MsgIssue msg = builder.setCreator(from)
                .addAllInputs(inputs)
                .addAllOutputs(outputs)
                .setDesc(desc)
                .build();

        return tx(key, Lists.newArrayList(Any.pack(msg)), Lists.newArrayList(fee));
    }

    public String Burn(ECKey key, Coin amount, String desc, Coin fee) throws Exception {
        String from = Crypto.getAddressFromECKey(key, address_prefix);

        List<Input> inputs = Lists.newArrayList();
        List<Output> outputs = Lists.newArrayList();
        // 支出
        List<Coin> total = Lists.newArrayList();
        total.add(amount);
        total.add(fee);
        fillInputsAndOutputs(from, total, inputs, outputs);

        MsgDestroy.Builder builder = MsgDestroy.newBuilder();
        MsgDestroy msg = builder.setCreator(from)
                .addAllInputs(inputs)
                .addAllOutputs(outputs)
                .setDesc(desc)
                .build();

        return tx( key, Lists.newArrayList(Any.pack(msg)), Lists.newArrayList(fee));
    }

    public String Transfer(ECKey key, String to, Coin amount, String desc, Coin fee) throws Exception {
        String from = Crypto.getAddressFromECKey(key, address_prefix);

        List<Input> inputs = Lists.newArrayList();
        List<Output> outputs = Lists.newArrayList();
        // 支出
        List<Coin> total = Lists.newArrayList();
        total.add(amount);
        total.add(fee);
        fillInputsAndOutputs(from, total, inputs, outputs);
        // 收入
        outputs.add(Output.newBuilder()
                .setToAddr(to)
                .setAmount(amount)
                .build());

        MsgDestroy.Builder builder = MsgDestroy.newBuilder();
        MsgDestroy msg = builder.setCreator(from)
                .addAllInputs(inputs)
                .addAllOutputs(outputs)
                .setDesc(desc)
                .build();

        return tx(key, Lists.newArrayList(Any.pack(msg)), Lists.newArrayList(fee));
    }

//    public String Deploy(ECKey key, String contract_name, String code_file, String args,  String desc, Coin fee) throws Exception {
//        String from = Crypto.getAddressFromECKey(key, address_prefix);
//        JSONObject kargs = JSONUtil.createObj();
//        kargs.set("contract_name", Base64.encode(contract_name.getBytes()));
//        kargs.set("contract_code", Base64.encode(FileUtil.readBytes(code_file)));
//        WasmCodeDesc wasm_desc = WasmCodeDesc.newBuilder().setRuntime("c").setContractType("wasm").build();
//        kargs.set("contract_desc", Base64.encode(wasm_desc.toByteArray()));
//
//        JSONObject initargs = JSONUtil.createObj();
//        for (Map.Entry entry: JSONUtil.parseObj(args).entrySet()) {
//            initargs.set(entry.getKey().toString(), Base64.encode(entry.getValue().toString().getBytes()));
//        }
//        kargs.set("init_args", Base64.encode(initargs.toString().getBytes()));
//        InvokeRequest invoke = InvokeRequest.builder().
//                module_name("kernel").
//                contract_name("").
//                method_name("Deploy").
//                args(kargs.toString()).
//                amount(Lists.newArrayList()).
//                build();
//        WasmMsgRequest request = WasmMsgRequest.builder().
//                fees(Lists.newArrayList(fee)).
//                desc(desc).
//                lock(true).
//                request(InvokeRPCRequest.builder().
//                        initiator(from).
//                        auth_require(Lists.newArrayList()).
//                        requests(Lists.newArrayList(invoke)).
//                        build()).
//                build();
//
//        return tx(key, msg_invoke(request), Lists.newArrayList(fee));
//    }
//
//    public String Upgrade(ECKey key, String contract_name, String code_file,  String desc, Coin fee) throws Exception {
//        String from = Crypto.getAddressFromECKey(key, address_prefix);
//        JSONObject kargs = JSONUtil.createObj();
//        kargs.set("contract_name", Base64.encode(contract_name.getBytes()));
//        kargs.set("contract_code", Base64.encode(FileUtil.readBytes(code_file)));
//        InvokeRequest invoke = InvokeRequest.builder().
//                module_name("kernel").
//                contract_name("").
//                method_name("Upgrade").
//                args(kargs.toString()).
//                amount(Lists.newArrayList()).
//                build();
//        WasmMsgRequest request = WasmMsgRequest.builder().
//                fees(Lists.newArrayList(fee)).
//                desc(desc).
//                lock(true).
//                request(InvokeRPCRequest.builder().
//                        initiator(from).
//                        auth_require(Lists.newArrayList()).
//                        requests(Lists.newArrayList(invoke)).
//                        build()).
//                build();
//
//        return tx(key, msg_invoke(request), Lists.newArrayList(fee));
//    }
//
//    public String Invoke(ECKey key, String contract_name, String method_name, String args,  String desc, Coin fee) throws Exception {
//        String from = Crypto.getAddressFromECKey(key, address_prefix);
//        JSONObject margs = JSONUtil.createObj();
//        for (Map.Entry entry: JSONUtil.parseObj(args).entrySet()) {
//            margs.set(entry.getKey().toString(), Base64.encode(entry.getValue().toString().getBytes()));
//        }
//        InvokeRequest invoke = InvokeRequest.builder().
//                module_name("wasm").
//                contract_name(contract_name).
//                method_name(method_name).
//                args(margs.toString()).
//                build();
//        WasmMsgRequest request = WasmMsgRequest.builder().
//                fees(Lists.newArrayList(fee)).
//                desc(desc).
//                lock(true).
//                request(InvokeRPCRequest.builder().
//                        initiator(from).
//                        auth_require(Lists.newArrayList()).
//                        requests(Lists.newArrayList(invoke)).
//                        build()).
//                build();
//
//        return tx(key, msg_invoke(request), Lists.newArrayList(fee));
//    }
//
//    public String ValidatorUpdate(ECKey key, String pub_key, Integer power, String name,  String desc, Coin fee) throws Exception {
//        String from = Crypto.getAddressFromECKey(key, address_prefix);
//        JSONObject msg = JSONUtil.createObj();
//        if (power != 0) {
//            JSONObject value = JSONUtil.createObj();
//            msg.set("type", "m0/MsgCreateValidator");
//            value.set("creator", from);
//            value.set("pubKey", pub_key);
//            value.set("moniker", name);
//            msg.set("value", value);
//        } else {
//            JSONObject value = JSONUtil.createObj();
//            value.set("creator", from);
//            value.set("pubKey", pub_key);
//            msg.set("type", "m0/MsgLeaveValidator");
//            msg.set("value", value);
//        }
//
//        return tx(key, JSONUtil.toJsonStr(msg), Lists.newArrayList(fee));
//    }
//
//    public String Query(String contract_name, String method_name, String args) throws Exception {
//        JSONObject margs = JSONUtil.createObj();
//        for (Map.Entry entry: JSONUtil.parseObj(args).entrySet()) {
//            margs.set(entry.getKey().toString(), Base64.encode(entry.getValue().toString().getBytes()));
//        }
//        InvokeRequest invoke = InvokeRequest.builder().
//                module_name("wasm").
//                contract_name(contract_name).
//                method_name(method_name).
//                args(margs.toString()).
//                build();
//        InvokeRPCRequest query = InvokeRPCRequest.builder().
//                requests(Lists.newArrayList(invoke)).
//                build();
//
//        ResultABCIQuery result = ResultABCIQuery.GetABCIQuery(client, "custom/wasm/query", JSONUtil.toJsonStr(query).getBytes(Charsets.UTF_8) ,0l, false);
//
//        String body = "";
//        JSONObject result_obj = JSONUtil.parseObj(result.response.getValue().toStringUtf8());
//        JSONArray responses = JSONUtil.parseArray(result_obj.getStr("responses"));
//        for (Object obj:responses) {
//            JSONObject response = (JSONObject)obj;
//            boolean status = response.getInt("status") == 200;
//            if (!status) {
//                // TODO
//            }
//            body = Base64.decodeStr(response.getStr("body"));//  Base64.decode(response.getStr("body").getBytes(Charsets.UTF_8)).toString();
//        }
//        return body;
//    }
}
