// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {ITaskMailbox, ITaskMailboxTypes} from "@hourglass-monorepo/src/interfaces/core/ITaskMailbox.sol";
import {OperatorSet} from "@eigenlayer-contracts/src/contracts/libraries/OperatorSetLib.sol";

contract FaaS {
    ITaskMailbox public immutable taskMailbox;
    address public immutable avs;
    uint32 public executorOperatorSetId;
    
    mapping(bytes32 => bytes) public functions;
    
    event FunctionRegistered(bytes32 indexed functionId, address indexed registrar);
    event FunctionCalled(bytes32 indexed functionId, bytes32 indexed taskId, address indexed caller);
    
    struct FunctionCall {
        bytes fn;
        bytes32 fnId;
        bytes input;
    }
    
    constructor(address _taskMailbox, address _avs, uint32 _executorOperatorSetId) {
        taskMailbox = ITaskMailbox(_taskMailbox);
        avs = _avs;
        executorOperatorSetId = _executorOperatorSetId;
    }
    
    function registerFunction(bytes memory content) external returns (bytes32) {
        bytes32 functionId = keccak256(content);
        functions[functionId] = content;
        
        emit FunctionRegistered(functionId, msg.sender);
        return functionId;
    }
    
    function callFunction(bytes32 functionId, bytes memory arguments) external returns (bytes32) {
        require(functions[functionId].length > 0, "Function not found");
        
        FunctionCall memory call = FunctionCall({
            fn: functions[functionId],
            fnId: functionId,
            input: arguments
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
    
    function getFunction(bytes32 functionId) external view returns (bytes memory) {
        return functions[functionId];
    }
}