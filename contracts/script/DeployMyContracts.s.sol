// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {Script, console} from "forge-std/Script.sol";
import {stdJson} from "forge-std/StdJson.sol";

import {IAllocationManager} from "@eigenlayer-contracts/src/contracts/interfaces/IAllocationManager.sol";
import {IKeyRegistrar} from "@eigenlayer-contracts/src/contracts/interfaces/IKeyRegistrar.sol";
import {IBN254CertificateVerifier} from
    "@eigenlayer-contracts/src/contracts/interfaces/IBN254CertificateVerifier.sol";
import {ITaskMailbox} from "@hourglass-monorepo/src/interfaces/core/ITaskMailbox.sol";

import {TaskAVSRegistrar} from "@project/l1-contracts/TaskAVSRegistrar.sol";
import {AVSTaskHook} from "@project/l2-contracts/AVSTaskHook.sol";
import {HelloWorld} from "@project/HelloWorld.sol"; // Import your custom contract

contract DeployMyContracts is Script {
    using stdJson for string;

    struct Context {
        address avs;
        uint256 avsPrivateKey;
        uint256 deployerPrivateKey;
        IAllocationManager allocationManager;
        IKeyRegistrar keyRegistrar;
        IBN254CertificateVerifier certificateVerifier;
        ITaskMailbox taskMailbox;
        TaskAVSRegistrar taskAVSRegistrar;
        AVSTaskHook taskHook;
    }

    struct Output {
        string name;
        address contractAddress;
    }

    function run(string memory environment, string memory _context) public {
        // Read the context
        Context memory context = _readContext(environment, _context);

        vm.startBroadcast(context.deployerPrivateKey);
        console.log("Deployer address:", vm.addr(context.deployerPrivateKey));

        //TODO: Implement custom contracts deployment
        // CustomContract customContract = new CustomContract();
        // console.log("CustomContract deployed to:", address(customContract));
        HelloWorld helloWorld = new HelloWorld();
        console.log("HelloWorld deployed to:", address(helloWorld));

        vm.stopBroadcast();

        vm.startBroadcast(context.avsPrivateKey);
        console.log("AVS address:", context.avs);

        //TODO: Implement any additional AVS setup

        vm.stopBroadcast();

        //TODO: Write to output file
        Output[] memory outputs = new Output[](1);
        // outputs[0] = Output({name: "CustomContract", address: address(customContract)});
        // _writeOutputToJson(environment, outputs);
        outputs[0] = Output({name: "HelloWorld", contractAddress: address(helloWorld)});
        _writeOutputToJson(environment, outputs);
    }

    function _readContext(
        string memory environment,
        string memory _context
    ) internal view returns (Context memory) {
        // Parse the context
        Context memory context;
        context.avs = stdJson.readAddress(_context, ".context.avs.address");
        context.avsPrivateKey = uint256(stdJson.readBytes32(_context, ".context.avs.avs_private_key"));
        context.deployerPrivateKey = uint256(stdJson.readBytes32(_context, ".context.deployer_private_key"));
        context.allocationManager = IAllocationManager(stdJson.readAddress(_context, ".context.eigenlayer.l1.allocation_manager"));
        context.keyRegistrar = IKeyRegistrar(stdJson.readAddress(_context, ".context.eigenlayer.l1.key_registrar"));
        context.certificateVerifier = IBN254CertificateVerifier(stdJson.readAddress(_context, ".context.eigenlayer.l2.bn254_certificate_verifier"));
        context.taskMailbox = ITaskMailbox(_readHourglassConfigAddress(environment, "taskMailbox"));
        context.taskAVSRegistrar = TaskAVSRegistrar(_readAVSL1ConfigAddress(environment, "taskAVSRegistrar"));
        context.taskHook = AVSTaskHook(_readAVSL2ConfigAddress(environment, "avsTaskHook"));

        return context;
    }

    function _readHourglassConfigAddress(
        string memory environment,
        string memory key
    ) internal view returns (address) {
        // Load the Hourglass config file
        string memory hourglassConfigFile =
                            string.concat("script/", environment, "/output/deploy_hourglass_core_output.json");
        string memory hourglassConfig = vm.readFile(hourglassConfigFile);

        // Parse and return the address
        return stdJson.readAddress(hourglassConfig, string.concat(".addresses.", key));
    }

    function _readAVSL1ConfigAddress(string memory environment, string memory key) internal view returns (address) {
        // Load the AVS L1 config file
        string memory avsL1ConfigFile = string.concat("script/", environment, "/output/deploy_avs_l1_output.json");
        string memory avsL1Config = vm.readFile(avsL1ConfigFile);

        // Parse and return the address
        return stdJson.readAddress(avsL1Config, string.concat(".addresses.", key));
    }

    function _readAVSL2ConfigAddress(string memory environment, string memory key) internal view returns (address) {
        // Load the AVS L2 config file
        string memory avsL2ConfigFile = string.concat("script/", environment, "/output/deploy_avs_l2_output.json");
        string memory avsL2Config = vm.readFile(avsL2ConfigFile);

        // Parse and return the address
        return stdJson.readAddress(avsL2Config, string.concat(".addresses.", key));
    }

    function _writeOutputToJson(
        string memory environment,
        Output[] memory outputs
    ) internal {
        uint256 length = outputs.length;

        if (length > 0) {
            // Add the addresses object
            string memory addresses = "addresses";

            for (uint256 i = 0; i < outputs.length - 1; i++) {
                vm.serializeAddress(addresses, outputs[i].name, outputs[i].contractAddress);
            }
            addresses = vm.serializeAddress(addresses, outputs[length - 1].name, outputs[length - 1].contractAddress);

            // Add the chainInfo object
            string memory chainInfo = "chainInfo";
            chainInfo = vm.serializeUint(chainInfo, "chainId", block.chainid);

            // Finalize the JSON
            string memory finalJson = "final";
            vm.serializeString(finalJson, "addresses", addresses);
            finalJson = vm.serializeString(finalJson, "chainInfo", chainInfo);

            // Write to output file
            string memory outputFile = string.concat("script/", environment, "/output/deploy_custom_contracts_output.json");
            vm.writeJson(finalJson, outputFile);
        }
    }
}
