pragma solidity ^0.4.15;


import 'zeppelin-solidity/contracts/token/StandardToken.sol';
import "zeppelin-solidity/contracts/ownership/Ownable.sol";

contract SNMTToken is StandardToken, Ownable {

    string public name = "SONM test token";

    string public symbol = "SNMT";

    uint public decimals = 18;

    uint constant TOKEN_LIMIT = 444 * 10e6 * 10e18;

    function SNMTToken(){
        totalSupply = TOKEN_LIMIT;
        owner = msg.sender;
        balances[owner] = TOKEN_LIMIT;
    }
}
