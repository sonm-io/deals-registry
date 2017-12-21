var SNMT_token = artifacts.require("./SNMTToken.sol");
var deals = artifacts.require("./Deals.sol");

module.exports = function(deployer) {
  deployer.deploy(deals, SNMT_token.address);
};
