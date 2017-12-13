var TSC = artifacts.require('./TSCToken.sol')
var Deals = artifacts.require('./Deals.sol')

contract('Deals', async function (accounts) {
  var deals
  var token
  var hub = accounts[1]
  var client = accounts[2]

  var test_spec = 2341234
  var test_price = 10000
  var test_workTime = 10000
  var test_PreviousDeal = 0 // always `0` - not implemented in platform

  let amountOfTestDeals = 2

  before(async function () {
    token = await TSC.deployed()
    deals = await Deals.deployed()

    await token.transfer(hub, 100000, {from: accounts[0]})
    await token.approve(deals.address, 100000, {from: hub})
    await token.transfer(client, 100000, {from: accounts[0]})
    await token.approve(deals.address, 100000, {from: client})

  })

  it('test null', async function () {
    let dealAmount = await deals.GetDealsAmount.call()
    assert.equal(dealAmount.toNumber(), 0)
  })

  it('test balances', async function () {
    let hubBalance = await token.balanceOf.call(hub)
    assert.equal(hubBalance.toNumber(), 100000)
    let clientBalance = await token.balanceOf.call(client)
    assert.equal(clientBalance.toNumber(), 100000)

    let hubAllowance = await token.allowance(hub, deals.address)
    // console.log(hubAllowance)
  })

  it('test CreateDeal', async function () {
    var dealOpenEvent = deals.DealOpened()
    dealOpenEvent.watch(async function (err, result) {
      if (err) {
        console.log(err)
        return
      }

      let deal_id = result.args.id
      let deal_info = await deals.GetDealInfo.call(deal_id)

      assert.equal(deal_info[2], hub)
      assert.equal(deal_info[1], client)
      assert.equal(deal_info[0].toNumber(), test_spec)
      assert.equal(deal_info[3].toNumber(), test_price)
      assert.equal(deal_info[4].toNumber(), 0)

      dealOpenEvent.stopWatching()
    })

    let deal_transaction = await deals.OpenDeal(hub, client, test_spec, test_price, test_workTime, {from: client})

    let dealAmount = await deals.GetDealsAmount.call()
    assert.equal(dealAmount, 1)
  })

  it('test AcceptDeal', async function () {
    var dealOpenEvent = deals.DealOpened()
    var dealAcceptedEvent = deals.DealAccepted()
    var deal_id

    dealOpenEvent.watch(async function (err, result) {
      if (err) {
        console.log(err)
        return
      }

      deal_id = result.args.id

      let transaction = await deals.AcceptDeal(deal_id, {from: hub})
      dealOpenEvent.stopWatching()
    })

    dealAcceptedEvent.watch(async function (err, result) {
      if (err) {
        console.log(err)
        return
      }

      let balanceOfClient = await token.balanceOf(client)
      let clientDeals = await deals.GetDeals(client)
      assert.equal(balanceOfClient, 100000 - test_price * clientDeals.length)

      let deal_info = await deals.GetDealInfo.call(deal_id)

      assert.equal(deal_info[2], hub)
      assert.equal(deal_info[1], client)
      assert.equal(deal_info[0].toNumber(), test_spec)
      assert.equal(deal_info[3].toNumber(), test_price)
      assert.equal(deal_info[5].toNumber(), 1)

      dealAcceptedEvent.stopWatching()
    })

    let deal_transaction = await deals.OpenDeal(hub, client, test_spec, test_price, test_workTime, {from: client})
  })

  it('test CloseDeal')

  it('test CancelDeal')

  it('test GetDealAmount', async function () {
    let dealAmount = await deals.GetDealsAmount.call({from: accounts[0]})
    assert.equal(dealAmount.toNumber(), amountOfTestDeals)
  })

  it('test GetClientDeals', async function () {
    let clientDeals = await deals.GetDeals.call(client)
    assert.equal(clientDeals.length, amountOfTestDeals)
  })

  it('test GetHubDeals')
})
