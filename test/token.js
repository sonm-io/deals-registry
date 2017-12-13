// var TSC = artifacts.require('./TSCToken.sol')
// var Deals = artifacts.require('./Deals.sol')
//
// contract('Token', async function (accounts) {
//   var token
//   var coinbase = accounts[0]
//   var hub = accounts[1]
//   var client = accounts[2]
//
//   var test_spec = 2341234
//   var test_price = 10000
//   var test_workTime = 10000
//   var test_PreviousDeal = 0 // always `0` - not implemented in platform
//
//   let amountOfTestDeals = 3
//
//   before(async function () {
//     token = await TSC.deployed()
//
//     await token.transfer(hub, 100000, {from: coinbase})
//     await token.transfer(client, 100000, {from: coinbase})
//   })
//
//   it('test balances', async function () {
//     let hubBalance = await token.balanceOf.call(hub)
//     assert.equal(hubBalance.toNumber(), 100000)
//     let clientBalance = await token.balanceOf.call(client)
//     assert.equal(clientBalance.toNumber(), 100000)
//   })
//
// })
//