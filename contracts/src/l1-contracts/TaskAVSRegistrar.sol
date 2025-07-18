// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {IAllocationManager} from "@eigenlayer-contracts/src/contracts/interfaces/IAllocationManager.sol";
import {IKeyRegistrar} from "@eigenlayer-contracts/src/contracts/interfaces/IKeyRegistrar.sol";

import {TaskAVSRegistrarBase} from "@hourglass-monorepo/src/avs/TaskAVSRegistrarBase.sol";

contract TaskAVSRegistrar is TaskAVSRegistrarBase {
    /**
     * @dev Constructor that passes parameters to parent TaskAVSRegistrarBase
     * @param _avs The address of the AVS
     * @param _allocationManager The AllocationManager contract address
     * @param _keyRegistrar The KeyRegistrar contract address
     */
    constructor(
        address _avs,
        IAllocationManager _allocationManager,
        IKeyRegistrar _keyRegistrar
    ) TaskAVSRegistrarBase(_avs, _allocationManager, _keyRegistrar) {}

    /**
     * @dev Initializer that calls parent initializer
     * @param _owner The owner of the contract
     * @param _initialConfig The initial AVS configuration
     */
    function initialize(address _owner, AvsConfig memory _initialConfig) external initializer {
        __TaskAVSRegistrarBase_init(_owner, _initialConfig);
    }

    // TODO: Implement
}
