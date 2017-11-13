pragma solidity ^0.4.15;


import "./TSCToken.sol";


contract Deals {

    using SafeMath for uint256;

    TSCToken token;

    enum Status{
    Pending,
    Accepted,
    Closed
    }

    struct Deal {
    uint256 specificationHash;
    address client;
    address hub;
    uint256 price;
    uint startTime;
    uint workTime;
    uint endTime;
    Status status;
    }

    event DealOpened(address hub, address client, uint index);

    event DealAccepted(address hub, address client, uint index);

    event DealClosed(address hub, address client, uint index);

    uint dealAmount = 0;

    mapping (uint => Deal) deals;

    mapping (uint => uint256) blockedBalance;

    mapping (address => uint[]) dealsIndex;

    function Deals(TSCToken _token){
        token = _token;
    }

    function OpenDeal(address _hub, address _client, uint256 _specHash, uint256 _price, uint _workTime){
        require(msg.sender == _client);

        dealAmount = dealAmount + 1;

        // startTime - 0, setting in AcceptDeal()
        // workTime - in seconds
        // endTime - 0, setting in AcceptDeal()
        deals[dealAmount] = Deal(_specHash, _client, _hub, _price, 0, _workTime, 0, Status.Pending);

        require(token.transferFrom(_client, this, _price));
        blockedBalance[dealAmount] = blockedBalance[dealAmount].add(_price);

        dealsIndex[_client].push(dealAmount);

        DealOpened(_hub, _client, dealAmount);
    }

    function AcceptDeal(uint id) {
        require(msg.sender == deals[id].hub);
        require(deals[id].status == Status.Pending);

        deals[id].status = Status.Accepted;
        deals[id].startTime = now;
        deals[id].endTime = now + deals[id].workTime;

        DealAccepted(deals[id].hub, deals[id].client, id);
    }

    function CloseDeal(uint id) {
        require(msg.sender == deals[id].client);
        if (deals[id].status == Status.Accepted) {
            // Closing deal
            if (now > deals[id].endTime) {
                // After endTime
                require(token.transfer(deals[id].hub, deals[id].price));
                blockedBalance[id] = blockedBalance[id].sub(deals[id].price);
            } else {
                // Before endTime
                var paidAmount = (now - deals[id].startTime) * (deals[id].price / deals[id].workTime);
                require(token.transfer(deals[id].hub, paidAmount));
                blockedBalance[id] = blockedBalance[id].sub(paidAmount);
                require(token.transfer(deals[id].client, deals[id].price - paidAmount));
                blockedBalance[id] = blockedBalance[id].sub(deals[id].price - paidAmount);
                deals[id].endTime = now;
            }
            deals[id].status = Status.Closed;
            DealClosed(deals[id].hub, deals[id].client, id);
        }else if (deals[id].status == Status.Pending) {
            // Canceling deal
            require(token.transfer(deals[id].client, deals[id].price));
            blockedBalance[id] = blockedBalance[id].sub(deals[id].price);

            deals[id].status = Status.Closed;

            DealClosed(deals[id].hub, deals[id].client, id);
        } else {
            revert();
        }
    }

    function GetDealInfo(uint dealIndex) constant returns (uint specHach, address client, address hub, uint price, uint startTime, uint workTime, uint endTIme, uint status){
        Deal storage deal = deals[dealIndex];
        return (deal.specificationHash, deal.client, deal.hub, deal.price, deal.startTime, deal.workTime, deal.endTime, uint(deal.status));
    }

    function GetDeals(address addr) constant returns (uint[]){
        return dealsIndex[addr];
    }

    function GetDealsAmount() constant returns (uint){
        return dealAmount;
    }

}