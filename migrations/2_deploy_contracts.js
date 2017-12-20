var SNMT_token = artifacts.require("./SNMTToken.sol");

module.exports = function(deployer) {
  deployer.deploy(SNMT_token);
};
