// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {Test, console} from "forge-std/Test.sol";
import {FaaS} from "@project/FaaS.sol";
import {ITaskMailboxTypes} from "@hourglass-monorepo/src/interfaces/core/ITaskMailbox.sol";
import {OperatorSet} from "@eigenlayer-contracts/src/contracts/libraries/OperatorSetLib.sol";

contract MockTaskMailbox {
    mapping(bytes32 => bytes) public tasks;
    uint256 private taskCounter = 1;
    
    function createTask(ITaskMailboxTypes.TaskParams memory taskParams) external returns (bytes32) {
        bytes32 taskHash = keccak256(abi.encodePacked(taskCounter, taskParams.payload));
        tasks[taskHash] = taskParams.payload;
        taskCounter++;
        return taskHash;
    }
    
    function getTaskInfo(bytes32 taskHash) external view returns (
        ITaskMailboxTypes.Task memory task
    ) {
        task.payload = tasks[taskHash];
        task.status = ITaskMailboxTypes.TaskStatus.Created;
        return task;
    }
    
    function getTaskResult(bytes32 /* taskHash */) external pure returns (bytes memory) {
        return ""; // Mock empty result
    }
}

contract FaaSTest is Test {
    FaaS public faas;
    MockTaskMailbox public mockTaskMailbox;
    address public user = address(0x1);
    address public mockAvs = address(0x2);
    uint32 public mockExecutorOperatorSetId = 1;
    
    function setUp() public {
        mockTaskMailbox = new MockTaskMailbox();
        faas = new FaaS(address(mockTaskMailbox), mockAvs, mockExecutorOperatorSetId);
    }
    
    function testRegisterFunction() public {
        bytes memory functionContent = "dummy tarball content";
        
        vm.prank(user);
        bytes32 functionId = faas.registerFunction(functionContent);
        
        assertEq(functionId, keccak256(functionContent));
        assertEq(faas.getFunction(functionId), functionContent);
    }
    
    function testCallFunction() public {
        bytes memory functionContent = "dummy tarball content";
        bytes memory arguments = abi.encode([1, 2, 3]);
        
        vm.prank(user);
        bytes32 functionId = faas.registerFunction(functionContent);
        
        vm.prank(user);
        bytes32 taskHash = faas.callFunction(functionId, arguments);
        
        assertTrue(taskHash != bytes32(0));
        
        bytes memory taskPayload = mockTaskMailbox.tasks(taskHash);
        assertTrue(taskPayload.length > 0);
        
        FaaS.FunctionCall memory call = abi.decode(taskPayload, (FaaS.FunctionCall));
        assertEq(call.fn, functionContent);
        assertEq(call.fnId, functionId);
        assertEq(call.input, arguments);
    }
    
    function testCallNonexistentFunction() public {
        bytes32 nonexistentFunctionId = keccak256("nonexistent");
        bytes memory arguments = abi.encode([1, 2, 3]);
        
        vm.prank(user);
        vm.expectRevert("Function not found");
        faas.callFunction(nonexistentFunctionId, arguments);
    }
    
    function testFunctionRegisteredEvent() public {
        bytes memory functionContent = "dummy tarball content";
        bytes32 expectedFunctionId = keccak256(functionContent);
        
        vm.expectEmit(true, true, false, true);
        emit FaaS.FunctionRegistered(expectedFunctionId, user);
        
        vm.prank(user);
        faas.registerFunction(functionContent);
    }
    
    function testFunctionCalledEvent() public {
        bytes memory functionContent = "dummy tarball content";
        bytes memory arguments = abi.encode([1, 2, 3]);
        
        vm.prank(user);
        bytes32 functionId = faas.registerFunction(functionContent);
        
        vm.expectEmit(true, false, true, false);
        emit FaaS.FunctionCalled(functionId, bytes32(0), user);
        
        vm.prank(user);
        faas.callFunction(functionId, arguments);
    }
    
    function testConstructorParameters() public view {
        assertEq(address(faas.taskMailbox()), address(mockTaskMailbox));
        assertEq(faas.avs(), mockAvs);
        assertEq(faas.executorOperatorSetId(), mockExecutorOperatorSetId);
    }
}
