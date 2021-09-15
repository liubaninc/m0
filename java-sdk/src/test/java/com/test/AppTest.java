package com.test;

import cn.hutool.core.util.StrUtil;
import com.tq_bc.m0.api.*;
import com.tq_bc.m0.common.Crypto;
import org.bitcoinj.core.ECKey;
import org.bouncycastle.util.encoders.Hex;
import org.junit.Test;

import java.io.IOException;
import java.util.Arrays;

/**
 * Unit test for simple App.
 */
public class AppTest {
    static private Api Api;
    static ECKey key;

    static {
        try {
            Api = new Api("http://127.0.0.1:26657");

            String mnemonic = "key erupt service six thing spy noise heart giggle year oil fuel rival drop goat deal moral require knee pact bind brain word nuclear";
            String privateKey = Crypto.getPrivateKeyFromMnemonicCode(Arrays.asList(StrUtil.split(mnemonic, " ")));
            key = ECKey.fromPrivate(Hex.decode(privateKey));
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    @Test
    public void testABCIInfo() throws Exception {
        ResultABCIInfo info = ResultABCIInfo.GetABCIInfo(Api);
        System.out.println(info.toJson());
    }

    @Test
    public void testABCIQuery() throws Exception {
//        ResponseInfo info = Api.GetABCIQuery();
//        System.out.println(Api.deserializer.toJson(info));
    }

    @Test
    public void testBlockLatest() throws Exception {
        ResultBlock info = ResultBlock.GetBlockLatest(Api);
        System.out.println(info.toJson());
    }

    @Test
    public void testBlock() throws Exception {
        ResultBlock info = ResultBlock.GetBlock(Api,5L);
        System.out.println(info.toJson());
    }

    @Test
    public void testBlockResultsLatest() throws Exception {
        ResultBlockResults info = ResultBlockResults.GetBlockResultsLatest(Api);
        System.out.println(info.toJson());
    }

    @Test
    public void testBlockResults() throws Exception {
        ResultBlockResults info = ResultBlockResults.GetBlockResults(Api,5L);
        System.out.println(info.toJson());
    }

    @Test
    public void testBlockChainInfo() throws Exception {
        ResultBlockChainInfo info = ResultBlockChainInfo.GetBlockChainInfo(Api, 1l, 3l);
        System.out.println(info.toJson());
    }

    @Test
    public void testCommitLatest() throws Exception {
        ResultCommit info = ResultCommit.GetCommitLatest(Api);
        System.out.println(info.toJson());
    }

    @Test
    public void testCommit() throws Exception {
        ResultCommit info = ResultCommit.GetCommit(Api,2l);
        System.out.println(info.toJson());
    }

    @Test
    public void testConsensusParamsLatest() throws Exception {
        ResultConsensusParams info = ResultConsensusParams.GetConsensusParamsLatest(Api);
        System.out.println(info.toJson());
    }

    @Test
    public void testConsensusParams() throws Exception {
        ResultConsensusParams info = ResultConsensusParams.GetConsensusParams(Api, 2l);
        System.out.println(info.toJson());
    }

    @Test
    public void testConsensusState() throws Exception {
        ResultConsensusState info = ResultConsensusState.GetConsensusState(Api);
        System.out.println(info.toJson());
    }

    @Test
    public void testDumpConsensusState() throws Exception {
        ResultDumpConsensusState info = ResultDumpConsensusState.GetDumpConsensusState(Api);
        System.out.println(info.toJson());
    }

    @Test
    public void testHealth() throws Exception {
        ResultHealth info = ResultHealth.GetHealth(Api);
        System.out.println(info.toJson());
    }

    @Test
    public void testGenesis() throws Exception {
        ResultGenesis info = ResultGenesis.GetGenesis(Api);
        System.out.println(info.toJson());
    }

    @Test
    public void testValidators() throws Exception {
        ResultValidators info = ResultValidators.GetValidators(Api,2l, 1, 100);
        System.out.println(info.toJson());
    }

    @Test
    public void testValidatorsLatest() throws Exception {
        ResultValidators info = ResultValidators.GetValidatorsLatest(Api, 1, 100);
        System.out.println(info.toJson());
    }

    @Test
    public void testNetInfo() throws Exception {
        ResultNetInfo info = ResultNetInfo.GetNetInfo(Api);
        System.out.println(info.toJson());
    }

    @Test
    public void testStatus() throws Exception {
        ResultStatus info = ResultStatus.GetStatus(Api);
        System.out.println(info.toJson());
    }

    @Test
    public void testTxSearch() throws Exception {
        ResultTxSearch info = ResultTxSearch.GetTxSearch(Api, "message.module=\"utxo\"", 1, 100);
        System.out.println(info.toJson());
    }

    @Test
    public void testTx() throws Exception {
        ResultTx info = ResultTx.GetTx(Api, "7731592DB0DC51602576257F89181876AB7B022FB8B4E5C9789E43EBD200E32C");
        System.out.println(info.toJson());
    }

    @Test
    public void testUnconfirmedTxs() throws Exception {
        ResultUnconfirmedTxs info = ResultUnconfirmedTxs.GetNumUnconfirmedTxs(Api);
        System.out.println(info.toJson());
    }

//    @Test
//    public void testAccountGet() throws Exception {
//        Account acct =  Account.Get(Api,"mc19dzfuxxv8vjeajjq475ahgrl0meudwexdmrnye");
//        System.out.println(JSONUtil.toJsonStr(acct));
//    }

//    @Test
//    public void testTxMint() throws Exception {
//        // 发行一种资产
//        Transaction tx = Transaction.builder().client(Api).memo("java sdk").commit(false).build();
//        String hash = tx.Mint(key, Coin.builder().amount("100000").denom("m1token").build(), "m1token", Coin.builder().amount("10").denom("m0token").build());
//        System.out.println(hash);
//    }
//
//    @Test
//    public void testTxBurn() throws Exception {
//        // 销毁多种资产
//        Transaction tx = Transaction.builder().client(Api).memo("java sdk").commit(false).build();
//        String hash = tx.Burn(key, Lists.newArrayList(Coin.builder().amount("100000").denom("m1token").build()), "m1token", Coin.builder().amount("10").denom("m0token").build());
//        System.out.println(hash);
//    }
//
//    @Test
//    public void testTxTransfer() throws Exception {
//        // 转移多种资产
//        Transaction tx = Transaction.builder().client(Api).memo("java sdk").commit(false).build();
//        String hash = tx.Transfer(key, "tk1lp386nzsxk2sa8ml63y6xat8d8ffpzz9cmkq9r", Lists.newArrayList(Coin.builder().amount("100000").denom("m0token").build()), "m1token", Coin.builder().amount("10").denom("m0token").build());
//        System.out.println(hash);
//    }
//
//    @Test
//    public void testTxDeploy() throws Exception {
//        // 部署合约
//        Transaction tx = Transaction.builder().client(Api).memo("java sdk").commit(false).build();
//        String hash = tx.Deploy(key,"testcounter", "../../counter.wasm", "{\"creator\":\"aaa\"}", "deploy", Coin.builder().amount("10").denom("m0token").build());
//        System.out.println(hash);
//    }
//
//    @Test
//    public void testTxUpgrade() throws Exception {
//        // 升级合约
//        Transaction tx = Transaction.builder().client(Api).memo("java sdk").commit(false).build();
//        String hash = tx.Upgrade(key, "testcounter", "../../counter.wasm", "upgrade", Coin.builder().amount("10").denom("m0token").build());
//        System.out.println(hash);
//    }
//
//    @Test
//    public void testTxInvoke() throws Exception {
//        // 调用合约
//        Transaction tx = Transaction.builder().client(Api).memo("java sdk").commit(false).build();
//        String hash = tx.Invoke(key, "testcounter", "increase", "{\"key\":\"aaa\"}", "invoke",  Coin.builder().amount("10").denom("m0token").build());
//        System.out.println(hash);
//    }
//
//    @Test
//    public void testTxQuery() throws Exception {
//        // 查询合约
//        Transaction tx = Transaction.builder().client(Api).memo("java sdk").build();
//        String result = tx.Query("testcounter", "get", "{\"key\":\"aaa\"}");
//        System.out.println(result);
//    }
//
//    @Test
//    public void testTxValidatorUpdate() throws Exception {
//        // 节点升级
//        String pubkey = Hex.toHexString(Base64.decode("AxttQW+MZpVBxMVNgA6T2kYI8LXSCNVHhc09web5c7TD"));
//        System.out.println(pubkey);
//
//        Transaction tx = Transaction.builder().client(Api).memo("java sdk").commit(false).build();
//        String hash = tx.ValidatorUpdate(key, "03a42abc882a0e7cc86cce697297d723d28dfd92ef1e966aeb3dd4f6df9c674acc", 100, "ssss", "update",  Coin.builder().amount("10").denom("m0token").build());
//        System.out.println(hash);
//    }
}

