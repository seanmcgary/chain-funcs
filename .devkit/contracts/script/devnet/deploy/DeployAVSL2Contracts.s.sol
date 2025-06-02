// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {Script, console} from "forge-std/Script.sol";

import {AVSTaskHook} from "@project/l2-contracts/AVSTaskHook.sol";
import {BN254CertificateVerifier} from "@project/l2-contracts/BN254CertificateVerifier.sol";

contract DeployAVSL2Contracts is Script {
    function run(
        string memory environment
    ) public {
        // Load the private key from the environment variable
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY_DEPLOYER");
        address deployer = vm.addr(deployerPrivateKey);

        // Deploy the AVSTaskHook and CertificateVerifier contracts
        vm.startBroadcast(deployerPrivateKey);
        console.log("Deployer address:", deployer);

        AVSTaskHook avsTaskHook = new AVSTaskHook();
        console.log("AVSTaskHook deployed to:", address(avsTaskHook));

        BN254CertificateVerifier bn254CertificateVerifier = new BN254CertificateVerifier();
        console.log("BN254CertificateVerifier deployed to:", address(bn254CertificateVerifier));

        vm.stopBroadcast();

        // Write deployment info to output file
        _writeOutputToJson(environment, address(avsTaskHook), address(bn254CertificateVerifier));
    }

    function _writeOutputToJson(
        string memory environment,
        address avsTaskHook,
        address bn254CertificateVerifier
    ) internal {
        // Add the addresses object
        string memory addresses = "addresses";
        vm.serializeAddress(addresses, "avsTaskHook", avsTaskHook);
        addresses = vm.serializeAddress(addresses, "bn254CertificateVerifier", bn254CertificateVerifier);

        // Add the chainInfo object
        string memory chainInfo = "chainInfo";
        chainInfo = vm.serializeUint(chainInfo, "chainId", block.chainid);

        // Finalize the JSON
        string memory finalJson = "final";
        vm.serializeString(finalJson, "addresses", addresses);
        finalJson = vm.serializeString(finalJson, "chainInfo", chainInfo);

        // Write to output file
        string memory outputFile = string.concat("script/", environment, "/output/deploy_avs_l2_output.json");
        vm.writeJson(finalJson, outputFile);
    }
}
