// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {ITaskMailbox, ITaskMailboxTypes} from "@hourglass-monorepo/src/interfaces/core/ITaskMailbox.sol";
import {OperatorSet} from "@eigenlayer-contracts/src/contracts/libraries/OperatorSetLib.sol";

contract FaaS {
    ITaskMailbox public immutable taskMailbox;
    address public immutable avs;
    uint32 public executorOperatorSetId;
    
    // Optimized storage layout - pack small data together
    struct FunctionMetadata {
        bool hasContent;    // 1 byte
        bool hasUrl;        // 1 byte 
        uint32 contentLength; // 4 bytes - fits in same slot with bools
    }
    
    // Legacy struct for compatibility
    struct Function {
        bytes content;
        string url;
    }
    
    // Separate mappings to avoid loading unnecessary data
    mapping(bytes32 => FunctionMetadata) public functionMetadata;
    mapping(bytes32 => bytes) public functionContent;
    mapping(bytes32 => string) public functionUrls;
    
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
        
        // Only store what's needed - no empty strings
        functionMetadata[functionId] = FunctionMetadata({
            hasContent: true,
            hasUrl: false,
            contentLength: uint32(content.length)
        });
        functionContent[functionId] = content;
        
        emit FunctionRegistered(functionId, msg.sender);
        return functionId;
    }
    
    function deployFunction(string memory url, bytes32 digest) external returns (bytes32) {
        // Only store what's needed - no empty bytes
        functionMetadata[digest] = FunctionMetadata({
            hasContent: false,
            hasUrl: true,
            contentLength: 0
        });
        functionUrls[digest] = url;
        
        emit FunctionDeployed(digest, url, digest, msg.sender);
        return digest;
    }
    
    function callFunction(bytes32 functionId, bytes memory arguments) external returns (bytes32) {
        FunctionMetadata storage metadata = functionMetadata[functionId];
        require(metadata.hasContent || metadata.hasUrl, "Function not found");
        
        bytes memory payload;
        
        // Use storage references to avoid copying to memory
        if (metadata.hasContent) {
            // Create minimal struct - avoid empty string allocation
            FunctionCall memory call;
            call.fn = functionContent[functionId];
            call.fnId = functionId;
            call.input = arguments;
            // call.url defaults to empty string
            payload = abi.encode(call);
        } else {
            // URL-only version - avoid loading content
            FunctionCall memory call;
            // call.fn defaults to empty bytes
            call.fnId = functionId;
            call.input = arguments;
            call.url = functionUrls[functionId];
            payload = abi.encode(call);
        }
        
        // Single storage read for operator set
        OperatorSet memory opSet = OperatorSet(avs, executorOperatorSetId);
        
        bytes32 taskHash = taskMailbox.createTask(ITaskMailboxTypes.TaskParams(
            msg.sender,
            0,
            opSet,
            payload
        ));
        
        emit FunctionCalled(functionId, taskHash, msg.sender);
        return taskHash;
    }
    
    // Legacy compatibility function - recreates Function struct
    function getFunction(bytes32 functionId) external view returns (Function memory) {
        FunctionMetadata memory metadata = functionMetadata[functionId];
        return Function({
            content: metadata.hasContent ? functionContent[functionId] : new bytes(0),
            url: metadata.hasUrl ? functionUrls[functionId] : ""
        });
    }
    
    function getFunctionContent(bytes32 functionId) external view returns (bytes memory) {
        return functionContent[functionId];
    }
    
    function getFunctionUrl(bytes32 functionId) external view returns (string memory) {
        return functionUrls[functionId];
    }
    
    // New optimized getter for metadata only
    function getFunctionMetadata(bytes32 functionId) external view returns (FunctionMetadata memory) {
        return functionMetadata[functionId];
    }
}