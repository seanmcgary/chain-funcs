// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {Script, console} from "forge-std/Script.sol";
import {stdJson} from "forge-std/StdJson.sol";

import {OperatorSet, OperatorSetLib} from "@eigenlayer-contracts/src/contracts/libraries/OperatorSetLib.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {ITaskMailbox, ITaskMailboxTypes} from "@hourglass-monorepo/src/interfaces/core/ITaskMailbox.sol";
import {IAVSTaskHook} from "@hourglass-monorepo/src/interfaces/avs/l2/IAVSTaskHook.sol";
import {IBN254CertificateVerifier} from "@hourglass-monorepo/src/interfaces/avs/l2/IBN254CertificateVerifier.sol";

contract SetupAVSTaskMailboxConfig is Script {
    using stdJson for string;

    function run(
        string memory environment,
        uint32 aggregatorOperatorSetId,
        uint32 executorOperatorSetId,
        uint96 taskSLA
    ) public {
        // Read addresses from config files
        address taskMailbox = _readHourglassConfigAddress(environment, "taskMailbox");
        console.log("Task Mailbox:", taskMailbox);

        // Read AVS L2 contract addresses
        address taskHook = _readAVSL2ConfigAddress(environment, "avsTaskHook");
        console.log("AVS Task Hook:", taskHook);
        address certificateVerifier = _readAVSL2ConfigAddress(environment, "bn254CertificateVerifier");
        console.log("BN254 Certificate Verifier:", certificateVerifier);

        // Load the private key from the environment variable
        uint256 avsPrivateKey = vm.envUint("PRIVATE_KEY_AVS");
        address avs = vm.addr(avsPrivateKey);

        vm.startBroadcast(avsPrivateKey);
        console.log("AVS address:", avs);

        // 1. Set the AVS config
        uint32[] memory executorOperatorSetIds = new uint32[](1);
        executorOperatorSetIds[0] = executorOperatorSetId;
        ITaskMailboxTypes.AvsConfig memory avsConfig = ITaskMailboxTypes.AvsConfig({
            aggregatorOperatorSetId: aggregatorOperatorSetId,
            executorOperatorSetIds: executorOperatorSetIds
        });
        ITaskMailbox(taskMailbox).setAvsConfig(avs, avsConfig);
        ITaskMailboxTypes.AvsConfig memory avsConfigStored = ITaskMailbox(taskMailbox).getAvsConfig(avs);
        console.log(
            "AVS config set:",
            avsConfigStored.aggregatorOperatorSetId,
            avsConfigStored.executorOperatorSetIds[0]
        );

        // 2. Set the Executor Operator Set Task Config
        ITaskMailboxTypes.ExecutorOperatorSetTaskConfig memory executorOperatorSetTaskConfig = ITaskMailboxTypes
            .ExecutorOperatorSetTaskConfig({
            certificateVerifier: certificateVerifier,
            taskHook: IAVSTaskHook(taskHook),
            feeToken: IERC20(address(0)),
            feeCollector: address(0),
            taskSLA: taskSLA,
            stakeProportionThreshold: 10_000,
            taskMetadata: bytes("")
        });
        ITaskMailbox(taskMailbox).setExecutorOperatorSetTaskConfig(
            OperatorSet(avs, executorOperatorSetId), executorOperatorSetTaskConfig
        );
        ITaskMailboxTypes.ExecutorOperatorSetTaskConfig memory executorOperatorSetTaskConfigStored =
            ITaskMailbox(taskMailbox).getExecutorOperatorSetTaskConfig(OperatorSet(avs, executorOperatorSetId));
        console.log(
            "Executor Operator Set Task Config set:",
            executorOperatorSetTaskConfigStored.certificateVerifier,
            address(executorOperatorSetTaskConfigStored.taskHook)
        );

        vm.stopBroadcast();
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

    function _readAVSL2ConfigAddress(string memory environment, string memory key) internal view returns (address) {
        // Load the AVS L2 config file
        string memory avsL2ConfigFile = string.concat("script/", environment, "/output/deploy_avs_l2_output.json");
        string memory avsL2Config = vm.readFile(avsL2ConfigFile);

        // Parse and return the address
        return stdJson.readAddress(avsL2Config, string.concat(".addresses.", key));
    }
}
