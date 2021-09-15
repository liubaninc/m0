package com.tq_bc.m0.api;
import com.google.protobuf.Descriptors;
import com.tq_bc.m0.common.JSONRPC;
import com.tq_bc.m0.common.ProtoJsonUtil;
import org.bouncycastle.util.encoders.Hex;
import tendermint.crypto.Keys.PublicKey;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Api {
    public static ProtoJsonUtil deserializer;
    static {
        try {
            List<Descriptors.Descriptor> anyFieldDescriptor = new ArrayList<>();
            anyFieldDescriptor.add(PublicKey.getDescriptor());
            deserializer = new ProtoJsonUtil(anyFieldDescriptor);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    private JSONRPC client;
    public String chain_id;
    public Api(String baseURL) throws Exception {
        client = new JSONRPC(baseURL);
        chain_id = GetGenesis().genesis.chain_id;
    }

    public ResultABCIInfo GetABCIInfo() throws Exception {
        Map<String, Object> req = new HashMap<>();
        return new ResultABCIInfo(client.request("abci_info", req).toString());
    }

    public ResultABCIQuery GetABCIQuery(String path, byte[] data, Long height, boolean prove) throws Exception {
        Map<String, Object> req = new HashMap<>();
        req.put("path", "\""+path+"\"");
        if (data != null) {
            req.put("data",  "0x"+Hex.toHexString(data));
        } else {
            req.put("data", null);
        }
        req.put("height", height);
        req.put("prove", prove);
        ResultABCIQuery result = new ResultABCIQuery(client.request("abci_query", req).toString());
        if (!result.response.getLog().isEmpty()){
            throw new RuntimeException(result.toJson());
        }
        return result;
    }

    public ResultBlock GetBlock(Long height) throws Exception {
        Map<String, Object> req = new HashMap<>();
        req.put("height", height);
        return new ResultBlock(client.request("block", req).toString());
    }

    public ResultBlock GetBlockLatest() throws Exception {
        Map<String, Object> req = new HashMap<>();
        return new ResultBlock(client.request("block", req).toString());
    }

    public ResultBlockResults GetBlockResults(Long height) throws Exception {
        Map<String, Object> req = new HashMap<>();
        req.put("height", height);
        return new ResultBlockResults(client.request("block_results", req).toString());
    }

    public ResultBlockResults GetBlockResultsLatest() throws Exception {
        Map<String, Object> req = new HashMap<>();
        return new ResultBlockResults(client.request("block_results", req).toString());
    }

    public ResultBlockChainInfo GetBlockChainInfo(Long min_height, Long max_height) throws Exception {
        Map<String, Object> req = new HashMap<>();
        req.put("minHeight", min_height);
        req.put("maxHeight", max_height);
        return new ResultBlockChainInfo(client.request("blockchain", req).toString());
    }

    public ResultCommit GetCommit(Long height) throws Exception {
        Map<String, Object> req = new HashMap<>();
        req.put("height", height);
        return new ResultCommit(client.request("commit", req).toString());
    }

    public ResultCommit GetCommitLatest() throws Exception {
        Map<String, Object> req = new HashMap<>();
        return new ResultCommit(client.request("commit", req).toString());
    }

    public ResultConsensusParams GetConsensusParams(long height) throws Exception {
        Map<String, Object> req = new HashMap<>();
        req.put("height", height);
        return new ResultConsensusParams(client.request("consensus_params", req).toString());
    }

    public ResultConsensusParams GetConsensusParamsLatest() throws Exception {
        Map<String, Object> req = new HashMap<>();
        return new ResultConsensusParams(client.request("consensus_params", req).toString());
    }

    public ResultConsensusState GetConsensusState() throws Exception {
        Map<String, Object> req = new HashMap<>();
        return new ResultConsensusState(client.request("consensus_state", req).toString());
    }

    public ResultDumpConsensusState GetDumpConsensusState() throws Exception {
        Map<String, Object> req = new HashMap<>();
        return new ResultDumpConsensusState(client.request("dump_consensus_state", req).toString());
    }

    public ResultHealth GetHealth() throws Exception {
        Map<String, Object> req = new HashMap<>();
        return new ResultHealth(client.request("health", req).toString());
    }

    public ResultGenesis GetGenesis() throws Exception {
        Map<String, Object> req = new HashMap<>();
        return new ResultGenesis(client.request("genesis", req).toString());
    }

    public ResultValidators GetValidators(Long height, Integer page, Integer per_page) throws Exception {
        Map<String, Object> req = new HashMap<>();
        req.put("height", height);
        req.put("page", page);
        req.put("per_page", per_page);
        return new ResultValidators(client.request("validators", req).toString());
    }

    public ResultValidators GetValidatorsLatest(Integer page, Integer per_page) throws Exception {
        Map<String, Object> req = new HashMap<>();
        req.put("page", page);
        req.put("per_page", per_page);
        return new ResultValidators(client.request("validators", req).toString());
    }

    public ResultNetInfo GetNetInfo() throws Exception {
        Map<String, Object> req = new HashMap<>();
        return new ResultNetInfo(client.request("net_info", req).toString());
    }

    public ResultStatus GetStatus() throws Exception {
        Map<String, Object> req = new HashMap<>();
        return new ResultStatus(client.request("status", req).toString());
    }

    public ResultTx GetTx(String hash) throws Exception {
        Map<String, Object> req = new HashMap<>();
        if (hash.startsWith("0X") || hash.startsWith("0x")) {
        } else {
            req.put("hash", "0x"+hash);
        }
        req.put("prove", false);
        return new ResultTx(client.request("tx", req).toString());
    }

    public ResultTxSearch GetTxSearch(String query, Integer page, Integer per_page) throws Exception {
        Map<String, Object> req = new HashMap<>();
        req.put("query", query);
        req.put("prove", false);
        req.put("page", page);
        req.put("per_page", per_page);
        req.put("order_by", "");
        return new ResultTxSearch(client.request("tx_search", req).toString());
    }

    public ResultUnconfirmedTxs GetUnconfirmedTxs(Integer limit) throws Exception {
        Map<String, Object> req = new HashMap<>();
        req.put("limit", limit);
        return new ResultUnconfirmedTxs(client.request("unconfirmed_txs", req).toString());
    }

    public ResultUnconfirmedTxs GetNumUnconfirmedTxs() throws Exception {
        Map<String, Object> req = new HashMap<>();
        return new ResultUnconfirmedTxs(client.request("num_unconfirmed_txs", req).toString());
    }

    public ResultBroadcastTx BroadcastTxSync(String tx)  throws Exception {
        Map<String, Object> req = new HashMap<>();
        req.put("tx", tx);
        return new ResultBroadcastTx(client.request("broadcast_tx_sync", req).toString());
    }

    public ResultBroadcastTx BroadcastTxAsync(String tx)  throws Exception {
        Map<String, Object> req = new HashMap<>();
        req.put("tx", tx);
        return new ResultBroadcastTx(client.request("broadcast_tx_async", req).toString());
    }

    public ResultBroadcastTxCommit BroadcastTxCommit(String tx)  throws Exception {
        Map<String, Object> req = new HashMap<>();
        req.put("tx", tx);
        return new ResultBroadcastTxCommit(client.request("broadcast_tx_commit", req).toString());
    }
}

