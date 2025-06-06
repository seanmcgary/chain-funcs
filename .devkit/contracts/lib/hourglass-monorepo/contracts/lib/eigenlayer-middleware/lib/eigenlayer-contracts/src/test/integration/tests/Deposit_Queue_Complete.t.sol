// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/test/integration/users/User.t.sol";
import "src/test/integration/IntegrationChecks.t.sol";

contract Integration_Deposit_QueueWithdrawal_Complete is IntegrationCheckUtils {
    /// Randomly generates a user with different held assets. Then:
    /// 1. deposit into strategy
    /// 2. queueWithdrawal
    /// 3. completeQueuedWithdrawal"
    function testFuzz_deposit_queueWithdrawal_completeAsTokens(uint24 _random) public rand(_random) {
        // Create a staker with a nonzero balance and corresponding strategies
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();

        // 1. Deposit into strategy
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);
        check_Deposit_State(staker, strategies, shares);

        // Ensure staker is not delegated to anyone post deposit
        assertFalse(delegationManager.isDelegated(address(staker)), "Staker should not be delegated after deposit");

        // 2. Queue Withdrawal
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, shares);

        // 3. Complete Queued Withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            uint[] memory expectedTokens = _calculateExpectedTokens(strategies, shares);
            IERC20[] memory tokens = staker.completeWithdrawalAsTokens(withdrawals[i]);
            check_Withdrawal_AsTokens_State(staker, User(payable(0)), withdrawals[i], strategies, shares, tokens, expectedTokens);
        }

        // Ensure staker is still not delegated to anyone post withdrawal completion
        assertFalse(delegationManager.isDelegated(address(staker)), "Staker should still not be delegated after withdrawal");
    }

    function testFuzz_deposit_queueWithdrawal_completeAsShares(uint24 _random) public rand(_random) {
        // Create a staker with a nonzero balance and corresponding strategies
        (User staker, IStrategy[] memory strategies, uint[] memory tokenBalances) = _newRandomStaker();

        // 1. Deposit into strategy
        staker.depositIntoEigenlayer(strategies, tokenBalances);
        uint[] memory shares = _calculateExpectedShares(strategies, tokenBalances);
        check_Deposit_State(staker, strategies, shares);

        // Ensure staker is not delegated to anyone post deposit
        assertFalse(delegationManager.isDelegated(address(staker)), "Staker should not be delegated after deposit");

        // 2. Queue Withdrawal
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, shares);

        // 3. Complete Queued Withdrawal
        _rollBlocksForCompleteWithdrawals(withdrawals);
        for (uint i = 0; i < withdrawals.length; i++) {
            staker.completeWithdrawalAsShares(withdrawals[i]);
            check_Withdrawal_AsShares_State(staker, User(payable(0)), withdrawals[i], strategies, shares);
        }

        // Ensure staker is still not delegated to anyone post withdrawal completion
        assertFalse(delegationManager.isDelegated(address(staker)), "Staker should still not be delegated after withdrawal");
    }

    /// @notice Regression test to ensure that depositSharesToRemove is not allowed to overflow. Previously
    /// this vulnerability would allow a staker to queue a withdrawal overflowing amount of deposit shares to withdraw
    /// resulting them in inflating their existing deposit shares. They could then inflate an operators delegated shares as much as they wanted.
    function testFuzz_deposit_revertOverflow_queueWithdrawal(uint24 _random) public rand(_random) {
        _configAssetTypes(HOLDS_ETH);
        // Create a staker with a nonzero balance and corresponding strategies
        (User staker, IStrategy[] memory strategies, uint[] memory initDepositShares) = _newRandomStaker();
        cheats.assume(initDepositShares[0] >= 64 ether);

        // Deposit staker by verifying withdrawal credentials
        (uint40[] memory validators, uint64 beaconBalanceGwei) = staker.startValidators();
        beaconChain.advanceEpoch_NoRewards();
        staker.verifyWithdrawalCredentials(validators);
        check_VerifyWC_State(staker, validators, beaconBalanceGwei);

        // Queue withdrawal for depositSharesToRemove that overflows
        int podOwnerDepositSharesBefore = eigenPodManager.podOwnerDepositShares(address(staker));
        assertEq(uint(podOwnerDepositSharesBefore), initDepositShares[0]);

        uint[] memory depositSharesToRemove = new uint[](1);
        depositSharesToRemove[0] = uint(type(int).max) + 100_000e18;

        cheats.expectRevert("SafeCast: value doesn't fit in an int256");
        Withdrawal[] memory withdrawals = staker.queueWithdrawals(strategies, depositSharesToRemove);
        bytes32[] memory withdrawalRoots = _getWithdrawalHashes(withdrawals);

        int podOwnerDepositSharesAfter = eigenPodManager.podOwnerDepositShares(address(staker));
        assertEq(podOwnerDepositSharesAfter, podOwnerDepositSharesBefore);
    }
}
