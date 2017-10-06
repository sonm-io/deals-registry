var TSC = artifacts.require("./TSCToken.sol");
var Deals = artifacts.require("./Deals.sol");

contract('Deals', async function (accounts) {

  var deals;
  var token;
  var hub = accounts[1];
  var client = accounts[2];

  var test_spec = 2341234;
  var test_price = 10000;

  before(async function () {
    token = await TSC.deployed();
    deals = await Deals.deployed();

  });

  it('test null', async function () {
    let dealAmount = await deals.GetDealAmount.call();
    assert.equal(dealAmount.toNumber(), 0);
  });

  it('test CreateDeal', async function () {

    var dealOpenEvent = deals.DealOpened();
    dealOpenEvent.watch(async function (err, result) {
      if (err) {
        console.log(err);
        return;
      }
      let deal_id = result.args.index;

      let deal_info = await deals.GetDealInfo.call(deal_id);
      console.log("Open+dealId:" + deal_id);

      assert.equal(deal_info[2], hub);
      assert.equal(deal_info[1], client);
      assert.equal(deal_info[0].toNumber(), test_spec);
      assert.equal(deal_info[3].toNumber(), test_price);
      assert.equal(deal_info[4].toNumber(), 0);

      dealOpenEvent.stopWatching();
    });


    let deal_transaction = await deals.OpenDeal(client, test_spec, test_price, {from: hub});

    let dealAmount = await deals.GetDealAmount.call();
    assert.equal(dealAmount, 1);
  });

  it('test AcceptDeal', async function () {
    var dealOpenEvent = deals.DealOpened();
    var dealAcceptedEvent = deals.DealAccepted();
    var deal_id;

    dealOpenEvent.watch(async function (err, result) {
      if (err) {
        console.log(err);
        return;
      }
      deal_id = result.args.index;
      console.log("Accept1+dealId:" + deal_id);


      let transaction = await deals.AcceptDeal(deal_id, {from: client});
      dealOpenEvent.stopWatching();
    });

    dealAcceptedEvent.watch(async function (err, result) {
      if (err) {
        console.log(err);
        return;
      }
      let deal_info = await deals.GetDealInfo.call(deal_id);
      console.log("Accept2+dealId:" + deal_id);

      assert.equal(deal_info[2], hub);
      assert.equal(deal_info[1], client);
      assert.equal(deal_info[0].toNumber(), test_spec);
      assert.equal(deal_info[3].toNumber(), test_price);
      assert.equal(deal_info[4].toNumber(), 1);
      dealAcceptedEvent.stopWatching();
    });

    let deal_transaction = await deals.OpenDeal(client, test_spec, test_price, {from: hub});

  });

  it('test CloseDeal', async function () {
    var dealOpenEvent = deals.DealOpened();
    var dealAcceptedEvent = deals.DealAccepted();
    var dealClosedEvent = deals.DealClosed();
    var deal_id;

    dealOpenEvent.watch(async function (err, result) {
      if (err) {
        console.log(err);
        return;
      }
      deal_id = result.args.index;
      console.log("Close1+dealId:" + deal_id);


      let transaction = await deals.AcceptDeal(deal_id, {from: client});
      dealOpenEvent.stopWatching();
    });

    dealAcceptedEvent.watch(async function (err, result) {
      if (err) {
        console.log(err);
        return;
      }
      let deal_info = await deals.GetDealInfo.call(deal_id);
      console.log("Close2+dealId:" + deal_id);

      assert.equal(deal_info[2], hub);
      assert.equal(deal_info[1], client);
      assert.equal(deal_info[0].toNumber(), test_spec);
      assert.equal(deal_info[3].toNumber(), test_price);
      assert.equal(deal_info[4].toNumber(), 1);

      let transactions = await deals.CloseDeal(deal_id, {from: hub});
      dealAcceptedEvent.stopWatching();
    });

    dealClosedEvent.watch(async function (err, result) {
      if (err) {
        console.log(err);
        return;
      }
      let deal_info = await deals.GetDealInfo.call(deal_id);
      console.log("Close3+dealId:" + deal_id);

      assert.equal(deal_info[2], hub);
      assert.equal(deal_info[1], client);
      assert.equal(deal_info[0].toNumber(), test_spec);
      assert.equal(deal_info[3].toNumber(), test_price);
      assert.equal(deal_info[4].toNumber(), 2);
      dealClosedEvent.stopWatching();
    });

    let deal_transaction = await deals.OpenDeal(client, test_spec, test_price, {from: hub});

  });

  it('test GetDealAmount', async function () {
    let dealAmount = await deals.GetDealAmount.call();
    assert.equal(dealAmount.toNumber(), 3);
  });

});
