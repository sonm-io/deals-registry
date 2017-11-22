// Import the page's CSS. Webpack will know what to do with it.
import '../stylesheets/app.css'

// Import libraries we need.
import { default as Web3 } from 'web3'
import { default as contract } from 'truffle-contract'

// Import our contract artifacts and turn them into usable abstractions.
import tokenArtifact from '../../build/contracts/TSCToken.json'
import dealsArtifact from '../../build/contracts/Deals.json'

// MetaCoin is our usable abstraction, which we'll use through the code below.
var Token = contract(tokenArtifact)
var Deals = contract(dealsArtifact)

// The following code is simple to show off interacting with your contracts.
// As your needs grow you will likely need to change its form and structure.
// For application bootstrapping, check out window.addEventListener below.
var accounts
var account

var client
var hub

var token
var deals

window.addEventListener('load', function () {
  // Checking if Web3 has been injected by the browser (Mist/MetaMask)
  if (typeof web3 !== 'undefined') {
    console.warn('Using web3 detected from external source. ' +
      'If you find that your accounts don\'t appear or you have 0 MetaCoin,' +
      ' ensure you\'ve configured that source properly.' +
      ' If using MetaMask, see the following link. ' +
      'Feel free to delete this warning. :) ' +
      'http://truffleframework.com/tutorials/truffle-and-metamask')
    // Use Mist/MetaMask's provider
    window.web3 = new Web3(window.web3.currentProvider)
  } else {
    console.warn('No web3 detected.' +
      'Falling back to http://localhost:8545. ' +
      'You should remove this fallback when you deploy live, as it\'s inherently insecure. ' +
      'Consider switching to Metamask for development. ' +
      'More info here: http://truffleframework.com/tutorials/truffle-and-metamask')
    // fallback - use your fallback strategy (local node / hosted node + in-dapp id mgmt / fail)
    window.web3 = new Web3(new Web3.providers.HttpProvider('http://127.0.0.1:8545'))
  }

  window.App.start()
})

var test_spec = 2341234
var test_price = 60
var test_workTime = 60

window.App = {
  start: async function () {
    var web3 = window.web3
    var self = this

    // Bootstrap the MetaCoin abstraction for Use.
    Token.setProvider(web3.currentProvider)
    Deals.setProvider(web3.currentProvider)

    // Get the initial account balance so it can be displayed.
    var accs = web3.eth.accounts

    if (accs.length === 0) {
      console.error('Couldn\'t get any accounts! Make sure your Ethereum client is configured correctly.')
      return
    }

    accounts = accs
    account = accounts[0]
    client = accounts[1]
    hub = accounts[2]

    token = await Token.deployed()
    deals = await Deals.deployed()

    self.syncView()
  },

  watch: function () {
    var event = token.Transfer()
    event.watch(function (err, result) {
      console.log(result)
    })
  },

  syncView: function () {
    this.loadDealsCount()
    this.loadCoinbaseBalance()
    this.loadClientBalance()
    this.loadHubBalance()
    this.loadDealsBalance()
    this.loadDeals()
  },

  loadDealsCount: function () {
    Deals.deployed().then(function (deals) {
      return deals.GetDealsAmount.call()
    }).then(function (result) {
      console.log(result.toString())
    })
  },

  sendCoin: function () {
    var self = this

    var amount = parseInt(document.getElementById('amount').value)
    var receiver = document.getElementById('receiver').value

    this.setStatus('Initiating transaction... (please wait)')

    var meta
    Token.deployed().then(function (instance) {
      meta = instance
      return meta.sendCoin(receiver, amount, {from: account})
    }).then(function () {
      self.setStatus('Transaction complete!')
      self.refreshBalance()
    }).catch(function (e) {
      console.log(e)
      self.setStatus('Error sending coin; see log.')
    })
  },

  sendToHub: async function () {
    var self = this

    await token.transfer(hub, 10000,  {from: account})
    await self.loadHubBalance()
    await self.loadCoinbaseBalance()
  },

  sendToClient: async function () {
    var self = this

    await token.transfer(client, 10000,  {from: account})
    await self.loadClientBalance()
    await self.loadCoinbaseBalance()
  },

  allowFromHub: async function () {
    var self = this

    await token.approve(deals.address, 10000,  {from: hub})
    await self.loadHubBalance()
    await self.loadCoinbaseBalance()
  },

  allowFromClient: async function () {
    var self = this

    await token.approve(deals.address, 20000,  {from: client})
    await self.loadClientBalance()
    await self.loadCoinbaseBalance()
  },

  loadCoinbaseBalance: async function () {
    var self = this

    let coinbaseBalance = await token.balanceOf(account)
    $('#coinbase-balance').html(coinbaseBalance.toString())
  },

  loadClientBalance: async function () {
    var self = this

    let clientBalance = await token.balanceOf(client)
    $('#client-balance').html(clientBalance.toString())
    let clientAllowance = await token.allowance(client, deals.address)
    $('#client-allowance').html(clientAllowance.toString())
  },

  loadHubBalance: async function () {
    var self = this

    let hubBalance = await token.balanceOf(hub)
    $('#hub-balance').html(hubBalance.toString())
    let hubAllowance = await token.allowance(hub, deals.address)
    $('#hub-allowance').html(hubAllowance.toString())
  },

  loadDealsBalance: async function () {
    var self = this

    let dealsBalance = await token.balanceOf(deals.address)
    $('#deals-balance').html(dealsBalance.toString())
  },

  loadDeals: async function(){
    var self = this

    var clientDeals = await deals.GetDeals(client)

    var html = ""

    for(var i = 0; i < clientDeals.length; i++){
      // console.log(clientDeals[i])
      var deal_id = clientDeals[i].toNumber()
      var deal = await deals.GetDealInfo(deal_id)
      console.log(deal)
      html += "<tr><td>"+deal_id+"</td><td>"+deal[3]+"</td><td>"+deal[4]+"</td><td>"+deal[5]+"</td><td>"+deal[6]+"</td><td>"+deal[7]+"</td></tr>"
    }
    $('#deals-table').html(html)
  },

  OpenDeal: async function (){
    await deals.OpenDeal(hub, client, test_spec, test_price, test_workTime, 0, {from: client, gas: 302000 })

    this.syncView()
  },

  AcceptDeal: async function(){
    let id = $("#dealId").val()

    await deals.AcceptDeal(id, {from: hub, gas: 90000})

    this.syncView()
  },

  CloseDeal: async function() {
    let id = $("#dealId").val()

    await deals.CloseDeal(id, {from: client, gas: 200000})

    this.syncView()
  },

  CancelDeal: async function(){
    let id = $("#dealId").val()

    await deals.CancelDeal(id, {from: client, gas: 80000})

    this.syncView()
  },

  LogAllEvents: async function(){
    token.allEvents(function(err, log){
      console.log(log)
    });
  }
}
