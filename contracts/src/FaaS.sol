// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {ITaskMailbox, ITaskMailboxTypes} from "@hourglass-monorepo/src/interfaces/core/ITaskMailbox.sol";
import {OperatorSet} from "@eigenlayer-contracts/src/contracts/libraries/OperatorSetLib.sol";

contract FaaS {
    ITaskMailbox public immutable taskMailbox;
    address public immutable avs;
    uint32 public executorOperatorSetId;
    
    struct Function {
        bytes content;
        string url;
    }
    
    mapping(bytes32 => Function) public functions;
    
    event FunctionRegistered(bytes32 indexed functionId, address indexed registrar);
    event FunctionDeployed(bytes32 indexed functionId, string indexed url, bytes32 indexed digest, address registrar);
    event FunctionCalled(bytes32 indexed functionId, bytes32 indexed taskId, address indexed caller);
    
    struct FunctionCall {
        bytes fn;
        bytes32 fnId;
        bytes input;
        string url;
    }
    
    constructor(address _taskMailbox, address _avs, uint32 _executorOperatorSetId) {
        taskMailbox = ITaskMailbox(_taskMailbox);
        avs = _avs;
        executorOperatorSetId = _executorOperatorSetId;
    }
    
    function registerFunction(bytes memory content) external returns (bytes32) {
        bytes32 functionId = keccak256(content);
        functions[functionId] = Function({
            content: content,
            url: ""
        });
        
        emit FunctionRegistered(functionId, msg.sender);
        return functionId;
    }
    
    function deployFunction(string memory url, bytes32 digest) external returns (bytes32) {
        bytes32 functionId = digest;
        functions[functionId] = Function({
            content: "",
            url: url
        });
        
        emit FunctionDeployed(functionId, url, digest, msg.sender);
        return functionId;
    }
    
    function callFunction(bytes32 functionId, bytes memory arguments) external returns (bytes32) {
        Function memory func = functions[functionId];
        require(func.content.length > 0 || bytes(func.url).length > 0, "Function not found");
        
        FunctionCall memory call = FunctionCall({
            fn: func.content,
            fnId: functionId,
            input: arguments,
            url: func.url
        });
        
        bytes memory payload = abi.encode(call);
        
        OperatorSet memory executorOperatorSet = OperatorSet({
            avs: avs,
            id: executorOperatorSetId
        });
        
        ITaskMailboxTypes.TaskParams memory taskParams = ITaskMailboxTypes.TaskParams({
            refundCollector: msg.sender,
            avsFee: 0,
            executorOperatorSet: executorOperatorSet,
            payload: payload
        });
        
        bytes32 taskHash = taskMailbox.createTask(taskParams);
        
        emit FunctionCalled(functionId, taskHash, msg.sender);
        return taskHash;
    }
    
    function getFunction(bytes32 functionId) external view returns (Function memory) {
        return functions[functionId];
    }
    
    function getFunctionContent(bytes32 functionId) external view returns (bytes memory) {
        return functions[functionId].content;
    }
    
    function getFunctionUrl(bytes32 functionId) external view returns (string memory) {
        return functions[functionId].url;
    }
}