pragma solidity ^0.4.14;


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

    // Minting for first stage of testing

    event Mint(address indexed to, uint256 amount);

    function mint(address _to, uint256 _amount) returns (bool) {
        totalSupply = totalSupply.add(_amount);
        balances[_to] = balances[_to].add(_amount);
        Mint(_to, _amount);
        return true;
    }

    function mintAndAllow(address _to, uint256 _amount, address _spender) returns (bool) {
        totalSupply = totalSupply.add(_amount);
        balances[_to] = balances[_to].add(_amount);
        Mint(_to, _amount);

        allowed[msg.sender][_spender] = balances[_to];
        Approval(msg.sender, _spender, balances[_to]);
        return true;
    }

    function () payable{
        mint(msg.sender, msg.value * 10000);
    }
}
