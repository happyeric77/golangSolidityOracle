// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package currencyOracle

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CurrencyOracleABI is the input ABI used to generate the binding from.
const CurrencyOracleABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currency\",\"type\":\"uint256\"}],\"name\":\"RequestCurrencyPrice\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"currencyPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_currency\",\"type\":\"uint256\"}],\"name\":\"requestCurrencyPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"updateCurrencyPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CurrencyOracleBin is the compiled bytecode used for deploying new contracts.
var CurrencyOracleBin = "0x608060405234801561001057600080fd5b506101a5806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80636df566d714610046578063da9fa81914610064578063e08812d714610080575b600080fd5b61004e61009c565b60405161005b9190610133565b60405180910390f35b61007e600480360381019061007991906100fb565b6100a2565b005b61009a600480360381019061009591906100fb565b6100ac565b005b60005481565b8060008190555050565b7fefd0eb611d30902f202bdd8e768d4997931078c8fb4999f518e09e58116bdf79816040516100db9190610133565b60405180910390a150565b6000813590506100f581610158565b92915050565b60006020828403121561010d57600080fd5b600061011b848285016100e6565b91505092915050565b61012d8161014e565b82525050565b60006020820190506101486000830184610124565b92915050565b6000819050919050565b6101618161014e565b811461016c57600080fd5b5056fea26469706673582212208738cca126dfda3273655f8150957f5f588f7cec3940bda85c876e24864639d564736f6c63430008040033"

// DeployCurrencyOracle deploys a new Ethereum contract, binding an instance of CurrencyOracle to it.
func DeployCurrencyOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CurrencyOracle, error) {
	parsed, err := abi.JSON(strings.NewReader(CurrencyOracleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CurrencyOracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CurrencyOracle{CurrencyOracleCaller: CurrencyOracleCaller{contract: contract}, CurrencyOracleTransactor: CurrencyOracleTransactor{contract: contract}, CurrencyOracleFilterer: CurrencyOracleFilterer{contract: contract}}, nil
}

// CurrencyOracle is an auto generated Go binding around an Ethereum contract.
type CurrencyOracle struct {
	CurrencyOracleCaller     // Read-only binding to the contract
	CurrencyOracleTransactor // Write-only binding to the contract
	CurrencyOracleFilterer   // Log filterer for contract events
}

// CurrencyOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurrencyOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurrencyOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurrencyOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurrencyOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurrencyOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurrencyOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurrencyOracleSession struct {
	Contract     *CurrencyOracle   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurrencyOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurrencyOracleCallerSession struct {
	Contract *CurrencyOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// CurrencyOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurrencyOracleTransactorSession struct {
	Contract     *CurrencyOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// CurrencyOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurrencyOracleRaw struct {
	Contract *CurrencyOracle // Generic contract binding to access the raw methods on
}

// CurrencyOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurrencyOracleCallerRaw struct {
	Contract *CurrencyOracleCaller // Generic read-only contract binding to access the raw methods on
}

// CurrencyOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurrencyOracleTransactorRaw struct {
	Contract *CurrencyOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurrencyOracle creates a new instance of CurrencyOracle, bound to a specific deployed contract.
func NewCurrencyOracle(address common.Address, backend bind.ContractBackend) (*CurrencyOracle, error) {
	contract, err := bindCurrencyOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurrencyOracle{CurrencyOracleCaller: CurrencyOracleCaller{contract: contract}, CurrencyOracleTransactor: CurrencyOracleTransactor{contract: contract}, CurrencyOracleFilterer: CurrencyOracleFilterer{contract: contract}}, nil
}

// NewCurrencyOracleCaller creates a new read-only instance of CurrencyOracle, bound to a specific deployed contract.
func NewCurrencyOracleCaller(address common.Address, caller bind.ContractCaller) (*CurrencyOracleCaller, error) {
	contract, err := bindCurrencyOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurrencyOracleCaller{contract: contract}, nil
}

// NewCurrencyOracleTransactor creates a new write-only instance of CurrencyOracle, bound to a specific deployed contract.
func NewCurrencyOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*CurrencyOracleTransactor, error) {
	contract, err := bindCurrencyOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurrencyOracleTransactor{contract: contract}, nil
}

// NewCurrencyOracleFilterer creates a new log filterer instance of CurrencyOracle, bound to a specific deployed contract.
func NewCurrencyOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*CurrencyOracleFilterer, error) {
	contract, err := bindCurrencyOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurrencyOracleFilterer{contract: contract}, nil
}

// bindCurrencyOracle binds a generic wrapper to an already deployed contract.
func bindCurrencyOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CurrencyOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurrencyOracle *CurrencyOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurrencyOracle.Contract.CurrencyOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurrencyOracle *CurrencyOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurrencyOracle.Contract.CurrencyOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurrencyOracle *CurrencyOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurrencyOracle.Contract.CurrencyOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurrencyOracle *CurrencyOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurrencyOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurrencyOracle *CurrencyOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurrencyOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurrencyOracle *CurrencyOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurrencyOracle.Contract.contract.Transact(opts, method, params...)
}

// CurrencyPrice is a free data retrieval call binding the contract method 0x6df566d7.
//
// Solidity: function currencyPrice() view returns(uint256)
func (_CurrencyOracle *CurrencyOracleCaller) CurrencyPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurrencyOracle.contract.Call(opts, &out, "currencyPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrencyPrice is a free data retrieval call binding the contract method 0x6df566d7.
//
// Solidity: function currencyPrice() view returns(uint256)
func (_CurrencyOracle *CurrencyOracleSession) CurrencyPrice() (*big.Int, error) {
	return _CurrencyOracle.Contract.CurrencyPrice(&_CurrencyOracle.CallOpts)
}

// CurrencyPrice is a free data retrieval call binding the contract method 0x6df566d7.
//
// Solidity: function currencyPrice() view returns(uint256)
func (_CurrencyOracle *CurrencyOracleCallerSession) CurrencyPrice() (*big.Int, error) {
	return _CurrencyOracle.Contract.CurrencyPrice(&_CurrencyOracle.CallOpts)
}

// RequestCurrencyPrice is a paid mutator transaction binding the contract method 0xe08812d7.
//
// Solidity: function requestCurrencyPrice(uint256 _currency) returns()
func (_CurrencyOracle *CurrencyOracleTransactor) RequestCurrencyPrice(opts *bind.TransactOpts, _currency *big.Int) (*types.Transaction, error) {
	return _CurrencyOracle.contract.Transact(opts, "requestCurrencyPrice", _currency)
}

// RequestCurrencyPrice is a paid mutator transaction binding the contract method 0xe08812d7.
//
// Solidity: function requestCurrencyPrice(uint256 _currency) returns()
func (_CurrencyOracle *CurrencyOracleSession) RequestCurrencyPrice(_currency *big.Int) (*types.Transaction, error) {
	return _CurrencyOracle.Contract.RequestCurrencyPrice(&_CurrencyOracle.TransactOpts, _currency)
}

// RequestCurrencyPrice is a paid mutator transaction binding the contract method 0xe08812d7.
//
// Solidity: function requestCurrencyPrice(uint256 _currency) returns()
func (_CurrencyOracle *CurrencyOracleTransactorSession) RequestCurrencyPrice(_currency *big.Int) (*types.Transaction, error) {
	return _CurrencyOracle.Contract.RequestCurrencyPrice(&_CurrencyOracle.TransactOpts, _currency)
}

// UpdateCurrencyPrice is a paid mutator transaction binding the contract method 0xda9fa819.
//
// Solidity: function updateCurrencyPrice(uint256 _price) returns()
func (_CurrencyOracle *CurrencyOracleTransactor) UpdateCurrencyPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _CurrencyOracle.contract.Transact(opts, "updateCurrencyPrice", _price)
}

// UpdateCurrencyPrice is a paid mutator transaction binding the contract method 0xda9fa819.
//
// Solidity: function updateCurrencyPrice(uint256 _price) returns()
func (_CurrencyOracle *CurrencyOracleSession) UpdateCurrencyPrice(_price *big.Int) (*types.Transaction, error) {
	return _CurrencyOracle.Contract.UpdateCurrencyPrice(&_CurrencyOracle.TransactOpts, _price)
}

// UpdateCurrencyPrice is a paid mutator transaction binding the contract method 0xda9fa819.
//
// Solidity: function updateCurrencyPrice(uint256 _price) returns()
func (_CurrencyOracle *CurrencyOracleTransactorSession) UpdateCurrencyPrice(_price *big.Int) (*types.Transaction, error) {
	return _CurrencyOracle.Contract.UpdateCurrencyPrice(&_CurrencyOracle.TransactOpts, _price)
}

// CurrencyOracleRequestCurrencyPriceIterator is returned from FilterRequestCurrencyPrice and is used to iterate over the raw logs and unpacked data for RequestCurrencyPrice events raised by the CurrencyOracle contract.
type CurrencyOracleRequestCurrencyPriceIterator struct {
	Event *CurrencyOracleRequestCurrencyPrice // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CurrencyOracleRequestCurrencyPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrencyOracleRequestCurrencyPrice)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CurrencyOracleRequestCurrencyPrice)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CurrencyOracleRequestCurrencyPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrencyOracleRequestCurrencyPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrencyOracleRequestCurrencyPrice represents a RequestCurrencyPrice event raised by the CurrencyOracle contract.
type CurrencyOracleRequestCurrencyPrice struct {
	Currency *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRequestCurrencyPrice is a free log retrieval operation binding the contract event 0xefd0eb611d30902f202bdd8e768d4997931078c8fb4999f518e09e58116bdf79.
//
// Solidity: event RequestCurrencyPrice(uint256 currency)
func (_CurrencyOracle *CurrencyOracleFilterer) FilterRequestCurrencyPrice(opts *bind.FilterOpts) (*CurrencyOracleRequestCurrencyPriceIterator, error) {

	logs, sub, err := _CurrencyOracle.contract.FilterLogs(opts, "RequestCurrencyPrice")
	if err != nil {
		return nil, err
	}
	return &CurrencyOracleRequestCurrencyPriceIterator{contract: _CurrencyOracle.contract, event: "RequestCurrencyPrice", logs: logs, sub: sub}, nil
}

// WatchRequestCurrencyPrice is a free log subscription operation binding the contract event 0xefd0eb611d30902f202bdd8e768d4997931078c8fb4999f518e09e58116bdf79.
//
// Solidity: event RequestCurrencyPrice(uint256 currency)
func (_CurrencyOracle *CurrencyOracleFilterer) WatchRequestCurrencyPrice(opts *bind.WatchOpts, sink chan<- *CurrencyOracleRequestCurrencyPrice) (event.Subscription, error) {

	logs, sub, err := _CurrencyOracle.contract.WatchLogs(opts, "RequestCurrencyPrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrencyOracleRequestCurrencyPrice)
				if err := _CurrencyOracle.contract.UnpackLog(event, "RequestCurrencyPrice", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequestCurrencyPrice is a log parse operation binding the contract event 0xefd0eb611d30902f202bdd8e768d4997931078c8fb4999f518e09e58116bdf79.
//
// Solidity: event RequestCurrencyPrice(uint256 currency)
func (_CurrencyOracle *CurrencyOracleFilterer) ParseRequestCurrencyPrice(log types.Log) (*CurrencyOracleRequestCurrencyPrice, error) {
	event := new(CurrencyOracleRequestCurrencyPrice)
	if err := _CurrencyOracle.contract.UnpackLog(event, "RequestCurrencyPrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
