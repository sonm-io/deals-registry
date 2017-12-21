pragma solidity ^0.4.15;


import 'zeppelin-solidity/contracts/token/StandardToken.sol';
import "zeppelin-solidity/contracts/ownership/Ownable.sol";


contract SNMTToken is StandardToken, Ownable {

    string public name = "SONM test token";

    string public symbol = "SNMT";

    uint public decimals = 18;

    event GiveAway(address whom, uint amount);

    function SNMTToken(){
        owner = msg.sender;
    }

    function mintToken(address target, uint256 mintedAmount) onlyOwner public returns (bool){
        balances[target] = balances[target].add(mintedAmount);
        totalSupply = totalSupply.add(mintedAmount);
        GiveAway(target, mintedAmount);
    }

    // @dev Gives the sender 100 SNMT's,
    function getTokens() public returns (bool){
        uint256 value = 100 * (10 ** decimals);
        balances[msg.sender] = balances[msg.sender].add(value);
        totalSupply = totalSupply.add(value);
        GiveAway(msg.sender, value);
        return true;
    }

    function() payable {
        getTokens();
    }
}
