var tsc_token = artifacts.require("./TSCToken.sol");

module.exports = function(deployer) {
  deployer.deploy(tsc_token);
};
