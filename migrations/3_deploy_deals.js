var tsc_token = artifacts.require("./TSCToken.sol");
var iterableMapping = artifacts.require("./IterableMapping.sol");
var deals = artifacts.require("./Deals.sol");

module.exports = function(deployer) {
  deployer.deploy(iterableMapping);
  deployer.link(iterableMapping, deals);
  deployer.deploy(deals, tsc_token.address);
};
