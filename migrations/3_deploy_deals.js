var tsc_token = artifacts.require("./TSCToken.sol");
var deals = artifacts.require("./Deals.sol");

module.exports = function(deployer) {
  deployer.deploy(deals, tsc_token.address);
};
