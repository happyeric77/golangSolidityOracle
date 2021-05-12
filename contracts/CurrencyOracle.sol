pragma solidity ^0.8.0;


contract CurrencyOracle {

    uint256 public currencyPrice;

    event RequestCurrencyPrice (uint256 currency);

    function requestCurrencyPrice (uint256 _currency) public {
        emit RequestCurrencyPrice(_currency);
    }

    function updateCurrencyPrice (uint256 _price) public {
        currencyPrice = _price;
    }    
}