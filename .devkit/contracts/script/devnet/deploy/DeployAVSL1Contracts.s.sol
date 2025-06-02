// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {Script, console} from "forge-std/Script.sol";

import {IAllocationManager} from "@eigenlayer-contracts/src/contracts/interfaces/IAllocationManager.sol";

import {TaskAVSRegistrar} from "@project/l1-contracts/TaskAVSRegistrar.sol";

contract DeployAVSL1Contracts is Script {
    function run(string memory environment, address avs, address allocationManager) public {
        // Load the private key from the environment variable
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY_DEPLOYER");
        address deployer = vm.addr(deployerPrivateKey);

        // Deploy the TaskAVSRegistrar middleware contract
        vm.startBroadcast(deployerPrivateKey);
        console.log("Deployer address:", deployer);

        TaskAVSRegistrar taskAVSRegistrar = new TaskAVSRegistrar(avs, IAllocationManager(allocationManager));
        console.log("TaskAVSRegistrar deployed to:", address(taskAVSRegistrar));

        vm.stopBroadcast();

        // Write deployment info to output file
        _writeOutputToJson(environment, address(taskAVSRegistrar));
    }

    function _writeOutputToJson(
        string memory environment,
        address taskAVSRegistrar
    ) internal {
        // Add the addresses object
        string memory addresses = "addresses";
        addresses = vm.serializeAddress(addresses, "taskAVSRegistrar", taskAVSRegistrar);

        // Add the chainInfo object
        string memory chainInfo = "chainInfo";
        chainInfo = vm.serializeUint(chainInfo, "chainId", block.chainid);

        // Finalize the JSON
        string memory finalJson = "final";
        vm.serializeString(finalJson, "addresses", addresses);
        finalJson = vm.serializeString(finalJson, "chainInfo", chainInfo);

        // Write to output file
        string memory outputFile = string.concat("script/", environment, "/output/deploy_avs_l1_output.json");
        vm.writeJson(finalJson, outputFile);
    }
}
