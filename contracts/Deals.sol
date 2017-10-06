pragma solidity ^0.4.14;


import "./IterableMapping.sol";


contract Deals {

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

    // Client => []Hubs => []Deals
    mapping (address => IterableMapping.itmap) dealsIndex;

    function OpenDeal(address _client, uint256 _specHash, uint256 _price){
        // setting own parameters
        address hubAddress = msg.sender;
        address clientAddress = _client;

        // define index for new deal
        uint dealIndex = dealAmount + 1;
        dealAmount = dealIndex;

        // create new deal
        deals[dealIndex] = Deal(_specHash, clientAddress, hubAddress, _price, Status.Pending);

        // create index of this deal
        IterableMapping.insert(dealsIndex[_client], uint144(bytes20(hubAddress) >> 16), dealIndex);

        DealOpened(hubAddress, clientAddress, dealIndex);
    }

    function AcceptDeal(uint dealIndex) {
        require(msg.sender == deals[dealIndex].client);
        deals[dealIndex].status = Status.Accepted;

        DealAccepted(deals[dealIndex].hub, deals[dealIndex].client, dealIndex);
    }

    function CloseDeal(uint dealIndex) {
        require(msg.sender == deals[dealIndex].hub);
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
    
}