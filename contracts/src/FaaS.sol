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
        address registrar;  // 20 bytes - fits in next slot
        uint32 timestamp;   // 4 bytes - fits in same slot with registrar
    }
    
    // Legacy struct for compatibility
    struct Function {
        bytes content;
        string url;
    }
    
    // Function listing data for pagination
    struct FunctionInfo {
        bytes32 functionId;
        string name;
        address registrar;
        uint32 timestamp;
        bool hasContent;
        bool hasUrl;
        uint32 contentLength;
    }
    
    // Separate mappings to avoid loading unnecessary data
    mapping(bytes32 => FunctionMetadata) public functionMetadata;
    mapping(bytes32 => bytes) public functionContent;
    mapping(bytes32 => string) public functionUrls;
    mapping(bytes32 => string) public functionNames;
    
    // Function listing arrays for pagination
    bytes32[] public functionIds;
    mapping(bytes32 => uint256) public functionIndex; // 1-based index (0 means not exists)
    
    event FunctionRegistered(bytes32 indexed functionId, string name, address indexed registrar);
    event FunctionDeployed(bytes32 indexed functionId, string name, string indexed url, bytes32 indexed digest, address registrar);
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
    
    function registerFunction(string memory name, bytes memory content) external returns (bytes32) {
        require(bytes(name).length > 0, "Function name cannot be empty");
        
        bytes32 functionId = keccak256(abi.encodePacked(msg.sender, name, content));
        require(functionIndex[functionId] == 0, "Function already exists");
        
        // Store function metadata
        functionMetadata[functionId] = FunctionMetadata({
            hasContent: true,
            hasUrl: false,
            contentLength: uint32(content.length),
            registrar: msg.sender,
            timestamp: uint32(block.timestamp)
        });
        functionContent[functionId] = content;
        functionNames[functionId] = name;
        
        // Add to listing array
        functionIds.push(functionId);
        functionIndex[functionId] = functionIds.length; // 1-based index
        
        emit FunctionRegistered(functionId, name, msg.sender);
        return functionId;
    }
    
    function deployFunction(string memory name, string memory url, bytes32 digest) external returns (bytes32) {
        require(bytes(name).length > 0, "Function name cannot be empty");
        require(bytes(url).length > 0, "Function URL cannot be empty");
        require(functionIndex[digest] == 0, "Function already exists");
        
        // Store function metadata
        functionMetadata[digest] = FunctionMetadata({
            hasContent: false,
            hasUrl: true,
            contentLength: 0,
            registrar: msg.sender,
            timestamp: uint32(block.timestamp)
        });
        functionUrls[digest] = url;
        functionNames[digest] = name;
        
        // Add to listing array
        functionIds.push(digest);
        functionIndex[digest] = functionIds.length; // 1-based index
        
        emit FunctionDeployed(digest, name, url, digest, msg.sender);
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
    
    function getFunctionName(bytes32 functionId) external view returns (string memory) {
        return functionNames[functionId];
    }
    
    function getTotalFunctions() external view returns (uint256) {
        return functionIds.length;
    }
    
    function listFunctions(uint256 offset, uint256 limit) external view returns (FunctionInfo[] memory functions, uint256 total) {
        total = functionIds.length;
        
        if (offset >= total) {
            return (new FunctionInfo[](0), total);
        }
        
        uint256 end = offset + limit;
        if (end > total) {
            end = total;
        }
        
        uint256 length = end - offset;
        functions = new FunctionInfo[](length);
        
        for (uint256 i = 0; i < length; i++) {
            bytes32 functionId = functionIds[offset + i];
            FunctionMetadata memory metadata = functionMetadata[functionId];
            
            functions[i] = FunctionInfo({
                functionId: functionId,
                name: functionNames[functionId],
                registrar: metadata.registrar,
                timestamp: metadata.timestamp,
                hasContent: metadata.hasContent,
                hasUrl: metadata.hasUrl,
                contentLength: metadata.contentLength
            });
        }
        
        return (functions, total);
    }
    
    function listFunctionsByRegistrar(address registrar, uint256 offset, uint256 limit) external view returns (FunctionInfo[] memory functions, uint256 total) {
        // First count functions by this registrar
        uint256 count = 0;
        for (uint256 i = 0; i < functionIds.length; i++) {
            if (functionMetadata[functionIds[i]].registrar == registrar) {
                count++;
            }
        }
        
        total = count;
        if (offset >= total) {
            return (new FunctionInfo[](0), total);
        }
        
        uint256 end = offset + limit;
        if (end > total) {
            end = total;
        }
        
        uint256 length = end - offset;
        functions = new FunctionInfo[](length);
        
        uint256 found = 0;
        uint256 added = 0;
        
        for (uint256 i = 0; i < functionIds.length && added < length; i++) {
            bytes32 functionId = functionIds[i];
            FunctionMetadata memory metadata = functionMetadata[functionId];
            
            if (metadata.registrar == registrar) {
                if (found >= offset) {
                    functions[added] = FunctionInfo({
                        functionId: functionId,
                        name: functionNames[functionId],
                        registrar: metadata.registrar,
                        timestamp: metadata.timestamp,
                        hasContent: metadata.hasContent,
                        hasUrl: metadata.hasUrl,
                        contentLength: metadata.contentLength
                    });
                    added++;
                }
                found++;
            }
        }
        
        return (functions, total);
    }
}