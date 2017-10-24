var TSC = artifacts.require('./TSCToken.sol')
var Deals = artifacts.require('./Deals.sol')

contract('Deals', async function (accounts) {
  var deals
  var token
  var hub = accounts[1]
  var client = accounts[2]
  var null_account = accounts[3]

  var test_spec = 2341234
  var test_price = 10000
  var test_securityAmount = 1000
  var test_workTime = 3600

  before(async function () {
    token = await TSC.deployed()
    deals = await Deals.deployed()

    await token.transfer(hub, 100000, {from: accounts[0]})
    await token.approve(deals.address, 100000, {from: hub})
    await token.transfer(client, 100000, {from: accounts[0]})
    await token.approve(deals.address, 100000, {from: client})

  })

  it('test null', async function () {
    let dealAmount = await deals.GetDealAmount.call()
    assert.equal(dealAmount.toNumber(), 0)
  })

  it('test balances', async function () {
    let hubBalance = await token.balanceOf.call(hub)
    assert.equal(hubBalance.toNumber(), 100000)
    let clientBalance = await token.balanceOf.call(client)
    assert.equal(clientBalance.toNumber(), 100000)

    let hubAllowance = await token.allowance(hub, deals.address)
    assert.equal(100000, hubAllowance.toNumber())
  })

  it('test CreateDeal', async function () {
    var clientBal = await token.balanceOf(client);

    var dealOpenEvent = deals.DealOpened()
    dealOpenEvent.watch(async function (err, result) {
      if (err) {
        console.log(err)
        return
      }

      let contractBalance = await token.balanceOf(deals.address)
      assert.equal(test_price, contractBalance.toNumber());

      let deal_id = result.args.id
      let deal_info = await deals.GetDealInfo.call(deal_id)
      assert.equal(deal_info[2], hub)
      assert.equal(deal_info[1], client)
      assert.equal(deal_info[0].toNumber(), test_spec)
      assert.equal(deal_info[3].toNumber(), test_price)
      assert.equal(deal_info[4].toNumber(), 0)
      assert.equal(deal_info[5].toNumber(), 0)
      assert.equal(deal_info[6].toNumber(), test_workTime)

      let dealBalance = await deals.GetBlockedBalance.call(deal_id)
      assert.equal(dealBalance.toNumber(), test_price)

      let clientBalance = await token.balanceOf(client)
      assert.equal(clientBalance.toNumber(), clientBal-test_price)

      dealOpenEvent.stopWatching()
    })

    let deal_transaction = await deals.OpenDeal(hub, test_spec, test_price, test_securityAmount, test_workTime, { from: client })

    let dealAmount = await deals.GetDealAmount.call()
    assert.equal(dealAmount, 1)
  })

  it('test CreateDeal -- fail', async function () {

    let nullBalance = await token.balanceOf(null_account)
    assert.equal(nullBalance.toNumber(), 0)

    try{
      let deal_transaction = await deals.OpenDeal(hub, test_spec, test_price, test_securityAmount, test_workTime, { from: null_account })
    }catch (e){
      let contractBalance = await token.balanceOf(deals.address)
      assert.equal(contractBalance.toNumber(), 10000)

      let dealAmount = await deals.GetDealAmount.call()
      assert.equal(dealAmount, 1)

      let dealBalance = await deals.GetBlockedBalance.call(2)
      assert.equal(0, dealBalance.toNumber())

      return
    }
    assert.Fail()
  })



  // it('test AcceptDeal', async function () {
  //   var dealOpenEvent = deals.DealOpened()
  //   var dealAcceptedEvent = deals.DealAccepted()
  //   var deal_id
  //
  //   dealOpenEvent.watch(async function (err, result) {
  //     if (err) {
  //       console.log(err)
  //       return
  //     }
  //
  //     deal_id = result.args.index
  //
  //     let transaction = await deals.AcceptDeal(deal_id, { from: hub })
  //     dealOpenEvent.stopWatching()
  //   })
  //
  //   dealAcceptedEvent.watch(async function (err, result) {
  //     if (err) {
  //       console.log(err)
  //       return
  //     }
  //
  //     let deal_info = await deals.GetDealInfo.call(deal_id)
  //
  //     assert.equal(deal_info[2], hub)
  //     assert.equal(deal_info[1], client)
  //     assert.equal(deal_info[0].toNumber(), test_spec)
  //     assert.equal(deal_info[3].toNumber(), test_price)
  //     assert.equal(deal_info[4].toNumber(), 1)
  //
  //     dealAcceptedEvent.stopWatching()
  //   })
  //
  //   let deal_transaction = await deals.OpenDeal(hub, test_spec, test_price, { from: client })
  // })
  //
  // it('test CloseDeal', async function () {
  //   var dealOpenEvent = deals.DealOpened()
  //   var dealAcceptedEvent = deals.DealAccepted()
  //   var dealClosedEvent = deals.DealClosed()
  //   var deal_id
  //
  //   dealOpenEvent.watch(async function (err, result) {
  //     if (err) {
  //       console.log(err)
  //       return
  //     }
  //     deal_id = result.args.index
  //
  //     let transaction = await deals.AcceptDeal(deal_id, { from: hub })
  //     dealOpenEvent.stopWatching()
  //   })
  //
  //   dealAcceptedEvent.watch(async function (err, result) {
  //     if (err) {
  //       console.log(err)
  //       return
  //     }
  //     let deal_info = await deals.GetDealInfo.call(deal_id)
  //
  //     assert.equal(deal_info[2], hub)
  //     assert.equal(deal_info[1], client)
  //     assert.equal(deal_info[0].toNumber(), test_spec)
  //     assert.equal(deal_info[3].toNumber(), test_price)
  //     assert.equal(deal_info[4].toNumber(), 1)
  //
  //     let transactions = await deals.CloseDeal(deal_id, { from: client })
  //     dealAcceptedEvent.stopWatching()
  //   })
  //
  //   dealClosedEvent.watch(async function (err, result) {
  //     if (err) {
  //       console.log(err)
  //       return
  //     }
  //     let deal_info = await deals.GetDealInfo.call(deal_id)
  //
  //     assert.equal(deal_info[2], hub)
  //     assert.equal(deal_info[1], client)
  //     assert.equal(deal_info[0].toNumber(), test_spec)
  //     assert.equal(deal_info[3].toNumber(), test_price)
  //     assert.equal(deal_info[4].toNumber(), 2)
  //
  //     let hubBalance = await token.balanceOf.call(hub)
  //     assert.equal(hubBalance.toNumber(), 110000)
  //     let clientBalance = await token.balanceOf.call(client)
  //     assert.equal(clientBalance.toNumber(), 70000)
  //
  //     dealClosedEvent.stopWatching()
  //   })
  //
  //   let deal_transaction = await deals.OpenDeal(hub, test_spec, test_price, { from: client })
  // })
  //
  // it('test GetDealAmount', async function () {
  //   let dealAmount = await deals.GetDealAmount.call({from: accounts[0]})
  //   assert.equal(dealAmount.toNumber(), 3)
  // })

  it('test GetClientDeals', async function () {
    let clientDeals = await deals.GetDealByClient.call(client)
    console.log('Client deals: ', clientDeals)
    assert.equal(clientDeals.length, 1)
  })

  it('test GetClientDeals', async function () {
    let hubDeals = await deals.GetDealByHubAddress.call(hub)
    console.log('Hub deals: ', hubDeals)
    assert.equal(hubDeals.length, 1)
  })
})
