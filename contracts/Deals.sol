pragma solidity ^0.4.14;


// не изучена ситуация с суммой сделок более чем allowance
// не изучена ситуация с трансфером до close

import "./IterableMapping.sol";
import "./TSCToken.sol";


contract Deals {

    uint dealSecuringPercentage = 10;

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
    Status status;
    }

    event DealOpened(address hub, address client, uint index);
    event DealAccepted(address hub, address client, uint index);
    event DealClosed(address hub, address client, uint index);

    uint dealAmount = 0;

    mapping (uint => Deal) deals;

    // TODO: must implement this
    mapping (address => uint256) blockedBalance;

    // Client => []Hubs => []Deals
    mapping (address => IterableMapping.itmap) dealsIndex;

    function Deals(TSCToken _token){
        token = _token;
    }

    // OpenDeal opening new deal
    // may to calling from everyone
    function OpenDeal(address _hub, uint256 _specHash, uint256 _price){
        // setting own parameters
        address hubAddress = _hub;
        address clientAddress = msg.sender;

        // if client has no allowed tokens to contract must be thrown
        require(IsPriceAllowed(clientAddress, _price));

        // define index for new deal
        uint dealIndex = dealAmount + 1;
        dealAmount = dealIndex;

        // create new deal
        deals[dealIndex] = Deal(_specHash, clientAddress, hubAddress, _price, Status.Pending);

        // create index of this deal
        // TODO: revert this
//        IterableMapping.insert(dealsIndex[_client], uint144(bytes20(hubAddress) >> 16), dealIndex);

        DealOpened(hubAddress, clientAddress, dealIndex);
    }

    function AcceptDeal(uint dealIndex) {
        require(msg.sender == deals[dealIndex].hub);

        // if hub has no allowed tokens to securingDeal must be thrown
        require(IsPriceAllowed(deals[dealIndex].hub, (deals[dealIndex].price / dealSecuringPercentage * 200000)));

        deals[dealIndex].status = Status.Accepted;

        DealAccepted(deals[dealIndex].hub, deals[dealIndex].client, dealIndex);
    }

    function CloseDeal(uint dealIndex) payable {
        var status = deals[dealIndex].status;

        // Its must be involve to revoke deal though CloseDeal
        if (status == Status.Pending){
            require(msg.sender == deals[dealIndex].hub);
        }else if(status == Status.Accepted){
            // TODO: test this
            // TODO: must be spilt by time
            token.transferFrom(deals[dealIndex].client, deals[dealIndex].hub, deals[dealIndex].price);
        }

        deals[dealIndex].status = Status.Closed;

        DealClosed(deals[dealIndex].hub, deals[dealIndex].client, dealIndex);
    }

    function GetDealInfo(uint dealIndex) returns (uint specHach, address client, address hub, uint price, uint status){
        Deal storage deal = deals[dealIndex];
        return (deal.specificationHash, deal.client, deal.hub, deal.price, uint(deal.status));
    }

    function GetDealAmount() returns (uint){
        return dealAmount;
    }

    // TODO: test this
    function IsPriceAllowed(address addr, uint _price) returns (bool){
        return ((token.allowance(addr, this)) - _price)  >= 0;
    }
}