// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../../src/BLSApkRegistry.sol";

// wrapper around the BLSApkRegistry contract that exposes internal functionality, for unit testing _other functionality_.
contract BLSApkRegistryHarness is BLSApkRegistry {
    constructor(
        ISlashingRegistryCoordinator _slashingRegistryCoordinator
    ) BLSApkRegistry(_slashingRegistryCoordinator) {}

    function setBLSPublicKey(address account, BN254.G1Point memory pk) external {
        bytes32 pubkeyHash = BN254.hashG1Point(pk);
        // store updates
        operatorToPubkeyHash[account] = pubkeyHash;
        pubkeyHashToOperator[pubkeyHash] = account;
        operatorToPubkey[account] = pk;
    }
}
