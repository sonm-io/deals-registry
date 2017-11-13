pragma solidity ^0.4.15;


import 'zeppelin-solidity/contracts/token/StandardToken.sol';
import "zeppelin-solidity/contracts/ownership/Ownable.sol";

contract TSCToken is StandardToken, Ownable {

    string public name = "TSC test token";

    string public symbol = "tsc";

    uint public decimals = 18;

    uint constant TOKEN_LIMIT = 444 * 10e6 * 10e18;

    function TSCToken(){
        totalSupply = TOKEN_LIMIT;
        owner = msg.sender;
        balances[owner] = TOKEN_LIMIT;
    }
}
