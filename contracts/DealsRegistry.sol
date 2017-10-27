pragma solidity ^0.4.15;


import "./Deals.sol";


contract DealsRegistry {

    address[] dealsContracts;

    address token;

    function GetAllVersion() constant returns(address[]){
        return dealsContracts;
    }

    function GetLastVersion() constant returns(address){
        return dealsContracts[dealsContracts.length];
    }

    function Migrate(address newContract) {
        var amount = Deals(this.GetLastVersion()).GetDealsAmount();
        Deals(newContract).setDealAmount(amount);
        dealsContract.push[newContract];
    }
}
