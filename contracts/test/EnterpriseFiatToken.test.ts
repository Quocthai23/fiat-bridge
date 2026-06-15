import { describe, it } from "node:test";
import { expect } from "chai";
import hre from "hardhat";
import { getAddress, parseEther } from "viem";

describe("EnterpriseFiatToken", function () {
  async function deployTokenFixture() {
    const [owner, minter, otherAccount, blacklistedUser] = await hre.viem.getWalletClients();

    const token = await hre.viem.deployContract("EnterpriseFiatToken");

    const publicClient = await hre.viem.getPublicClient();

    const minterRole = await token.read.MINTER_ROLE();
    await token.write.grantRole([minterRole, minter.account.address]);

    return {
      token,
      owner,
      minter,
      otherAccount,
      blacklistedUser,
      publicClient,
    };
  }

  describe("Test-Case S1 (Blacklist Drain)", function () {
    it("Should revert Mint and Burn transactions for blacklisted users, preventing Gas Drain", async function () {
      const { token, minter, blacklistedUser } = await deployTokenFixture();

      const blacklistRole = await token.read.BLACKLISTER_ROLE();

      await token.write.blacklist([blacklistedUser.account.address]);

      const amount = parseEther("1000");

      const tokenAsMinter = await hre.viem.getContractAt(
        "EnterpriseFiatToken",
        token.address,
        { client: { wallet: minter } }
      );

      await expect(
        tokenAsMinter.write.mint(["core-tx-001", blacklistedUser.account.address, amount])
      ).to.be.rejectedWith("Recipient is blacklisted");

      const burnerRole = await token.read.BURNER_ROLE();
      await token.write.grantRole([burnerRole, minter.account.address]);

      await expect(
        tokenAsMinter.write.burn(["core-tx-002", blacklistedUser.account.address, amount])
      ).to.be.rejectedWith("Sender is blacklisted");

      const balance = await token.read.balanceOf([blacklistedUser.account.address]);
      expect(balance).to.equal(0n);
    });
  });

  describe("Test-Case S2 (Bypass Role)", function () {
    it("Should revert if an account without MINTER_ROLE tries to call mint", async function () {
      const { token, otherAccount } = await deployTokenFixture();

      const amount = parseEther("1000");

      const tokenAsOther = await hre.viem.getContractAt(
        "EnterpriseFiatToken",
        token.address,
        { client: { wallet: otherAccount } }
      );

      const minterRole = await token.read.MINTER_ROLE();
      const expectedError = `AccessControlUnauthorizedAccount("${getAddress(otherAccount.account.address)}", "${minterRole}")`;

      await expect(
        tokenAsOther.write.mint(["core-tx-003", otherAccount.account.address, amount])
      ).to.be.rejectedWith(expectedError);
    });
  });

  describe("Idempotency (processedCoreTxs)", function () {
    it("Should revert if the same coreTxId is processed twice", async function () {
      const { token, minter, otherAccount } = await deployTokenFixture();

      const amount = parseEther("100");

      const tokenAsMinter = await hre.viem.getContractAt(
        "EnterpriseFiatToken",
        token.address,
        { client: { wallet: minter } }
      );

      await tokenAsMinter.write.mint(["core-tx-111", otherAccount.account.address, amount]);

      await expect(
        tokenAsMinter.write.mint(["core-tx-111", otherAccount.account.address, amount])
      ).to.be.rejectedWith("Transaction already processed");
    });
  });
});
