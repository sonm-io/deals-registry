pragma solidity ^0.4.14;


import "./TSCToken.sol";


contract iDeals {


    // --------------------
    // DEAL FUNCTION
    // --------------------

    // OpenDeal function to start new deal between hub and client
    // create new deal and inserting it into indexes mappings
    // returns bool, not more.
    // other params returns via DealOpened event, spending by success result
    function OpenDeal(address hub, address client, uint specificationHash, uint256 pricePerHour);

    event DealOpened(address hub, address client, uint id);

    // AcceptDeal function uses for assure deal with opponent
    // in conclusion we decide, that deal must be started with client
    // whatever opponent is hub.
    // success case of execution AcceptDeal signaling that work must be started
    // return bool, not more.
    // other params returns via DealAccepted event, spending by success result of execution.
    function AcceptDeal(uint id) returns (bool);

    event DealAccepted(address hub, address client, uint id);

    // CloseDeal function uses for abort work execution and put in order hub and client interact above this deal.
    // May be used by hub - that might that hub want to end this deal otherwise:
    // 1) hub found more paid and complex deal for our slot
    // 2) client not paid deal more
    // 3) other purpose about price, illegal use, not properly use of slot resourse
    // Or may be used by client - that might that client want to stop work.
    // если клиент завершает сделку - то хаб получает все заработанные деньги(через функцию PayDealOut() ),
    // а клиент получает заблокированные токены обратно.
    // CloseDeal signaling that work must be stopped, if was not done before.
    // In this case
    // return bool, not more.
    // other params returns via DealClosed event, spending by success result of execution.
    function CloseDeal() returns (bool);

    event DealClosed(address hub, address client, uint id);

    //    // PayDeal function - uses for deposit tokens for pay your deal
    //    // any arrange of tokens must be allowed for contract for success execution of this function
    //    // returns bool, not more.
    //    // other params returns via DealPaid event, spending by success result of execution.
    //    function PayDeal(uint id, uint256 amount) returns (bool);
    //    event DealPaid(uint id, uint256 amount);
    //
    //    // PayOutDeal function uses for withdraw token from deal to hub
    //    // PayDealOut withdraw tokens from deal to hub imedeatly by working time
    //    // returns bool, not more.
    //    // other params returns via DealPaidOut event, spending by success result of execution.
    //    function PayOutDeal(uint id) returns (bool);
    //    event DealPaidOut(uint id, uint256 amount);


    // --------------------
    // GETTERS
    // --------------------

    //    // GetHubByClient(address) returns (uint[])
    //    // return array of uint responding all client deals without status
    //    function GetDealsByClient(address client) returns (uint[]);
    //
    //    // GetDealsByHub(address) returns (uint[])
    //    // return array of uint responding all hubs deals without status
    //    function GetDealsByHub(address hub) returns (uint[]);
    //
    //    // GetDeal info by deal id
    //    // returns any additional info about deal
    //    function GetDealInfo(uint id) constant returns (uint specificicationHash, address client, address hub, uint256 pricePerHour, uint status);
}


contract Deals {

    // --------------------
    // MAIN STRUCTURES
    // --------------------

    struct Deal {
    uint specificationHash;
    address client;
    address hub;
    uint256 price;
    uint256 securityAmount;
    Status status;
    uint startTime;
    uint workTime;
    uint endTime;
    uint256 paidAmount;
    }

    // Status
    // Now we decide to 3 avaiable status to deal
    // Pending - after create, wait acepting - not working
    // Accepted - time to work! At these status work not be doing(imidiatly after accepting, balance for deal is down, hub is down)
    // Closed - deal is ended/or deal is canceled by client/hub by any reasons
    enum Status{
    Pending, // int(Status.Pending) = 0
    Accepted, // int(Status.Accepted) = 1
    Closed, // int(Status.Closed) = 2
    ReportedFailFromHub,
    ReportedFailFromClient
    }

    event DealOpened(address hub, address client, uint id);

    event DealAccepted(address hub, address client, uint id);

    event DealClosed(address hub, address client, uint id);

    event DealCanceled(address hub, address client, uint id);

    event WithdrawTokens(uint id, uint256 amount);

    // For trusted execution above token work using SafeMath
    using SafeMath for uint256;

    uint dealSecuringPercentage = 10;

    TSCToken token;

    uint dealAmount = 0;

    mapping (uint => Deal) deals;

    // TODO: must implement this
    mapping (uint => uint256) blockedBalance;

    mapping (uint => uint256) blockedSecurityBalance;

    // Client => []Hubs => []Deals
    mapping (address => uint[]) hubDealsIndex;

    mapping (address => uint[]) clientDealsIndex;

    function Deals(TSCToken _token){
        token = _token;
    }


    // --------------------
    // EXTERNAL
    // --------------------


    // OpenDeal opening new deal
    // may to calling from everyone
    function OpenDeal(address _hub, uint _specificationHash, uint256 _price, uint256 _securityAmount, uint _workTime){
        // define index for new deal
        var _client = msg.sender;

        uint dealIndex = dealAmount + 1;
        dealAmount = dealIndex;

        // create new deal
        deals[dealIndex] = Deal(_specificationHash, _client, _hub, _price, _securityAmount, Status.Pending, 0, _workTime, 0, 0);

        // if client has no allowed tokens to contract must be thrown
        require(token.transferFrom(_client, this, _price));
        blockedBalance[dealIndex] = blockedBalance[dealIndex].add(_price);

        // create index of this deal
        clientDealsIndex[_client].push(dealIndex);
        hubDealsIndex[_hub].push(dealIndex);

        DealOpened(_hub, _client, dealIndex);
    }

    function AcceptDeal(uint id) returns (bool){

        require(msg.sender == deals[id].hub);

        deals[id].status = Status.Accepted;
        deals[id].startTime = now;


        DealOpened(0x0000000000000000000000000000000000000000, 0x0000000000000000000000000000000000000000, id);

        return false;
    }

    function CloseDeal(uint id) returns (bool){
        DealClosed(0x0000000000000000000000000000000000000000, 0x0000000000000000000000000000000000000000, id);

        return false;
    }

    function withdrawTokens(uint id){
        // TODO: здесь должна быть проверка на статус сделки - если сделка Reported, то токены нельзя снять
        var pricePerSecond = deals[id].price / deals[id].workTime;

        var withdrawAmount = (now - deals[id].startTime) * pricePerSecond - deals[id].paidAmount;

        require(token.transfer(deals[id].hub, withdrawAmount));
        blockedBalance[id] = blockedBalance.sub(withdrawAmount);

        WithdrawTokens(id, withdrawAmount);
    }

    // Must be used for canceling Deal from client
    function CancelDeal(uint id){
        require(msg.sender == deals[id].client);


        if (deals[id].status == Status.Pending) {
            require(token.transfer(deals[id].client, deals[id].price));
            blockedBalance[id] = blockedBalance[id].sub(deals[id].price);
        } else {
            withdrawTokens(id);
            require(token.transfer(deals[id].client, blockedBalance[id]));
            blockedBalance[id] = blockedBalance[id].sub(deals[id].price);
            // TODO: здесь должно быть возвращение токенов обеспечения хабу
        }


        deals[id].status = Status.Closed;

        DealCanceled(deals[id].hub, deals[id].client, id);
    }

    function FailWork(){
        require(msg.sender == deals[id].hub);

    }

    function ReportFail(uint id){
        require(deals[id].status == Status.Accepted);

        if (msg.sender == deals[id].hub){
            deals[id].status = Status.ReportedFailFromHub;
            return;
        }

        if (msg.sender == deals[id].client){
            deals[id].status = Status.ReportedFailFromClient;
            return;
        }

        revert();
    }

    function ConfirmFail(){
        if (deals[id].status == Status.ReportedFailFromClient){
            withdrawToken
        }

        if (deals[id].status == Status.ReportedFailFromClient){

        }

        revert();
    }

    function RejectFail(){
        require((deals[id].status == Status.ReportedFailFromHub)||(deals[id].status == Status.ReportedFailFromClient));

    }

    //    function AcceptDeal(uint dealIndex) {
    //        require(msg.sender == deals[dealIndex].hub);
    //
    //        // if hub has no allowed tokens to securingDeal must be thrown
    ////        require(IsPriceAllowed(deals[dealIndex].hub, (deals[dealIndex].price / dealSecuringPercentage )));
    //
    //        deals[dealIndex].status = Status.Accepted;
    //        deals[dealIndex].status = Status.Accepted;
    //
    //        DealAccepted(deals[dealIndex].hub, deals[dealIndex].client, dealIndex);
    //    }
    //
    //    function CloseDeal(uint id) payable {
    //        var status = deals[id].status;
    //
    //        // Its must be involve to revoke deal though CloseDeal
    ////        if (status == Status.Pending){
    ////            require(msg.sender == deals[id].hub);
    ////        }else if(status == Status.Accepted){
    ////            // TODO: must be spilt by time
    ////            require(token.transfer(deals[id].hub, deals[id].pricePerHour));
    //////            blockedBalance[deals[id].client] = blockedBalance[deals[id].client].sub(deals[id].pricePerHour);
    ////        }
    //
    //        deals[id].status = Status.Closed;
    //
    //        DealClosed(deals[id].hub, deals[id].client, id);
    //    }
    //

    // --------------------
    // GETTERS
    // --------------------

    function GetDealInfo(uint id) constant returns (uint specificicationHash, address client, address hub, uint256 price, uint status, uint startTime, uint workTime){
        var deal = deals[id];
        return (deal.specificationHash, deal.client, deal.hub, deal.price, uint(deal.status), deal.startTime, deal.workTime);
    }

    function GetDealByClient(address _client) constant returns (uint[]){
        return clientDealsIndex[_client];
    }

    function GetDealByHubAddress(address _hub) constant returns (uint[]){
        return hubDealsIndex[_hub];
    }

    function GetDealAmount() constant returns (uint){
        return dealAmount;
    }

    function GetBlockedBalance(uint id) constant returns (uint256){
        return blockedBalance[id];
    }

}