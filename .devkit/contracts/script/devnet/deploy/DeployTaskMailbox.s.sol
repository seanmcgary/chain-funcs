// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {Script, console} from "forge-std/Script.sol";

import {TaskMailbox} from "@hourglass-monorepo/src/core/TaskMailbox.sol";

contract DeployTaskMailbox is Script {
    function run(
        string memory environment
    ) public {
        // Load the private key from the environment variable
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY_DEPLOYER");
        address deployer = vm.addr(deployerPrivateKey);

        // Deploy the TaskMailbox contract
        vm.startBroadcast(deployerPrivateKey);
        console.log("Deployer address:", deployer);

        TaskMailbox taskMailbox = new TaskMailbox();
        console.log("TaskMailbox deployed to:", address(taskMailbox));

        vm.stopBroadcast();

        // Write deployment info to output file
        _writeOutputToJson(environment, address(taskMailbox));
    }

    function _writeOutputToJson(string memory environment, address taskMailbox) internal {
        // Add the addresses object
        string memory addresses = "addresses";
        addresses = vm.serializeAddress(addresses, "taskMailbox", taskMailbox);

        // Add the chainInfo object
        string memory chainInfo = "chainInfo";
        chainInfo =vm.serializeUint(chainInfo, "chainId", block.chainid);

        // Finalize the JSON
        string memory finalJson = "final";
        vm.serializeString(finalJson, "addresses", addresses);
        finalJson = vm.serializeString(finalJson, "chainInfo", chainInfo);

        // Write to output file
        string memory outputFile = string.concat("script/", environment, "/output/deploy_hourglass_core_output.json");
        vm.writeJson(finalJson, outputFile);
    }
}
