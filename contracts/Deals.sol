pragma solidity ^0.4.15;


import "./SNMTToken.sol";


contract Deals {

    using SafeMath for uint256;

    SNMTToken token;

    uint constant STATUS_PENDING = 1;
    uint constant STATUS_ACCEPTED = 2;
    uint constant STATUS_CLOSED = 3;

    struct Deal {
    uint256 specificationHash;
    address client;
    address hub;
    uint256 price;
    uint startTime;
    uint workTime;
    uint endTime;
    uint status;
    }

    event DealOpened(address indexed hub, address indexed client, uint indexed id);

    event DealAccepted(address indexed hub, address indexed client, uint indexed id);

    event DealClosed(address indexed hub, address indexed client, uint indexed id);

    uint dealAmount = 0;

    mapping (uint => Deal) deals;

    mapping (uint => uint256) blockedBalance;

    mapping (address => uint[]) dealsIndex;

    function Deals(SNMTToken _token){
        token = _token;
    }

    function OpenDeal(address _hub, address _client, uint256 _specHash, uint256 _price, uint _workTime){
        require(msg.sender == _client);

        dealAmount = dealAmount + 1;

        // startTime - 0, setting in AcceptDeal()
        // workTime - in seconds
        // endTime - 0, setting in AcceptDeal()
        deals[dealAmount] = Deal(_specHash, _client, _hub, _price, 0, _workTime, 0, STATUS_PENDING);

        require(token.transferFrom(_client, this, _price));
        blockedBalance[dealAmount] = blockedBalance[dealAmount].add(_price);

        dealsIndex[_client].push(dealAmount);
        dealsIndex[_hub].push(dealAmount);

        DealOpened(_hub, _client, dealAmount);
    }

    function AcceptDeal(uint id) {
        require(msg.sender == deals[id].hub);
        require(deals[id].status == STATUS_PENDING);

        deals[id].status = STATUS_ACCEPTED;
        deals[id].startTime = now;
        deals[id].endTime = now + deals[id].workTime;

        DealAccepted(deals[id].hub, deals[id].client, id);
    }

    function CloseDeal(uint id) {
        if (deals[id].status == STATUS_ACCEPTED) {
            // Closing deal
            if (now > deals[id].endTime) {
                // After endTime
                require(token.transfer(deals[id].hub, deals[id].price));
                blockedBalance[id] = blockedBalance[id].sub(deals[id].price);
            } else {
                require(msg.sender == deals[id].client);
                // Before endTime
                var paidAmount = (now - deals[id].startTime) * (deals[id].price / deals[id].workTime);
                require(token.transfer(deals[id].hub, paidAmount));
                blockedBalance[id] = blockedBalance[id].sub(paidAmount);
                require(token.transfer(deals[id].client, deals[id].price - paidAmount));
                blockedBalance[id] = blockedBalance[id].sub(deals[id].price - paidAmount);
                deals[id].endTime = now;
            }
            deals[id].status = STATUS_CLOSED;
            DealClosed(deals[id].hub, deals[id].client, id);
        }else if (deals[id].status == STATUS_PENDING) {
            require(msg.sender == deals[id].client);
            // Canceling deal
            require(token.transfer(deals[id].client, deals[id].price));
            blockedBalance[id] = blockedBalance[id].sub(deals[id].price);

            deals[id].status = STATUS_CLOSED;

            DealClosed(deals[id].hub, deals[id].client, id);
        } else {
            revert();
        }
    }

    function GetDealInfo(uint dealIndex) constant returns (uint specHach, address client, address hub, uint price, uint startTime, uint workTime, uint endTIme, uint status){
        Deal storage deal = deals[dealIndex];
        return (deal.specificationHash, deal.client, deal.hub, deal.price, deal.startTime, deal.workTime, deal.endTime, deal.status);
    }

    function GetDeals(address addr) constant returns (uint[]){
        return dealsIndex[addr];
    }

    function GetDealsAmount() constant returns (uint){
        return dealAmount;
    }

}