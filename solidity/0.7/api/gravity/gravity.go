// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gravity

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

// GravityABI is the input ABI used to generate the binding from.
const GravityABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newConsuls\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"newBftValue\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"bftValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"consuls\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConsuls\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newConsuls\",\"type\":\"address[]\"}],\"name\":\"hashNewConsuls\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rounds\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newConsuls\",\"type\":\"address[]\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"newLastRound\",\"type\":\"uint256\"}],\"name\":\"updateConsuls\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GravityFuncSigs maps the 4-byte function signature to its string representation.
var GravityFuncSigs = map[string]string{
	"3cec1bdd": "bftValue()",
	"a2c2c617": "consuls(uint256)",
	"ad595b1a": "getConsuls()",
	"4dea5eba": "hashNewConsuls(address[])",
	"8c65c81f": "rounds(uint256)",
	"92c388ab": "updateConsuls(address[],uint8[],bytes32[],bytes32[],uint256)",
}

// GravityBin is the compiled bytecode used for deploying new contracts.
var GravityBin = "0x608060405234801561001057600080fd5b506040516109343803806109348339818101604052604081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b90830190602082018581111561006857600080fd5b825186602082028301116401000000008211171561008557600080fd5b82525081516020918201928201910280838360005b838110156100b257818101518382015260200161009a565b50505050919091016040525060209081015184519093506100d992506001918501906100e3565b5060025550610167565b828054828255906000526020600020908101928215610138579160200282015b8281111561013857825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190610103565b50610144929150610148565b5090565b5b808211156101445780546001600160a01b0319168155600101610149565b6107be806101766000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80633cec1bdd146100675780634dea5eba146100815780638c65c81f1461012257806392c388ab14610153578063a2c2c6171461037e578063ad595b1a146103b7575b600080fd5b61006f61040f565b60408051918252519081900360200190f35b61006f6004803603602081101561009757600080fd5b810190602081018135600160201b8111156100b157600080fd5b8201836020820111156100c357600080fd5b803590602001918460208302840111600160201b831117156100e457600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610415945050505050565b61013f6004803603602081101561013857600080fd5b50356104d1565b604080519115158252519081900360200190f35b61037c600480360360a081101561016957600080fd5b810190602081018135600160201b81111561018357600080fd5b82018360208201111561019557600080fd5b803590602001918460208302840111600160201b831117156101b657600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561020557600080fd5b82018360208201111561021757600080fd5b803590602001918460208302840111600160201b8311171561023857600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561028757600080fd5b82018360208201111561029957600080fd5b803590602001918460208302840111600160201b831117156102ba57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561030957600080fd5b82018360208201111561031b57600080fd5b803590602001918460208302840111600160201b8311171561033c57600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955050913592506104e6915050565b005b61039b6004803603602081101561039457600080fd5b503561067b565b604080516001600160a01b039092168252519081900360200190f35b6103bf6106a2565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156103fb5781810151838201526020016103e3565b505050509050019250505060405180910390f35b60025481565b6000606060005b83518110156104c2578184828151811061043257fe5b60200260200101516040516020018083805190602001908083835b6020831061046c5780518252601f19909201916020918201910161044d565b6001836020036101000a038019825116818451168082178552505050505050905001826001600160a01b031660601b8152601401925050506040516020818303038152906040529150808060010191505061041c565b50805160209091012092915050565b60006020819052908152604090205460ff1681565b6000806104f287610415565b905060005b6001548110156105fa576001818154811061050e57fe5b9060005260206000200160009054906101000a90046001600160a01b03166001600160a01b031660018389848151811061054457fe5b602002602001015189858151811061055857fe5b602002602001015189868151811061056c57fe5b602002602001015160405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa1580156105c8573d6000803e3d6000fd5b505050602060405103516001600160a01b0316146105e75760006105ea565b60015b60ff1692909201916001016104f7565b50600254821015610646576040805162461bcd60e51b81526020600482015260116024820152701a5b9d985b1a590818999d0818dbdd5b9d607a1b604482015290519081900360640190fd5b86516106599060019060208a0190610704565b5050506000908152602081905260409020805460ff1916600117905550505050565b6001818154811061068857fe5b6000918252602090912001546001600160a01b0316905081565b606060018054806020026020016040519081016040528092919081815260200182805480156106fa57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116106dc575b5050505050905090565b828054828255906000526020600020908101928215610759579160200282015b8281111561075957825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190610724565b50610765929150610769565b5090565b5b808211156107655780546001600160a01b031916815560010161076a56fea26469706673582212204b7174b5136a9266dbee6537f004fc1825e28e856e3391f88f46fe93c115433464736f6c63430007010033"

// DeployGravity deploys a new Ethereum contract, binding an instance of Gravity to it.
func DeployGravity(auth *bind.TransactOpts, backend bind.ContractBackend, newConsuls []common.Address, newBftValue *big.Int) (common.Address, *types.Transaction, *Gravity, error) {
	parsed, err := abi.JSON(strings.NewReader(GravityABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GravityBin), backend, newConsuls, newBftValue)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Gravity{GravityCaller: GravityCaller{contract: contract}, GravityTransactor: GravityTransactor{contract: contract}, GravityFilterer: GravityFilterer{contract: contract}}, nil
}

// Gravity is an auto generated Go binding around an Ethereum contract.
type Gravity struct {
	GravityCaller     // Read-only binding to the contract
	GravityTransactor // Write-only binding to the contract
	GravityFilterer   // Log filterer for contract events
}

// GravityCaller is an auto generated read-only Go binding around an Ethereum contract.
type GravityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GravityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GravityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GravityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GravityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GravitySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GravitySession struct {
	Contract     *Gravity          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GravityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GravityCallerSession struct {
	Contract *GravityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// GravityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GravityTransactorSession struct {
	Contract     *GravityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GravityRaw is an auto generated low-level Go binding around an Ethereum contract.
type GravityRaw struct {
	Contract *Gravity // Generic contract binding to access the raw methods on
}

// GravityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GravityCallerRaw struct {
	Contract *GravityCaller // Generic read-only contract binding to access the raw methods on
}

// GravityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GravityTransactorRaw struct {
	Contract *GravityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGravity creates a new instance of Gravity, bound to a specific deployed contract.
func NewGravity(address common.Address, backend bind.ContractBackend) (*Gravity, error) {
	contract, err := bindGravity(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gravity{GravityCaller: GravityCaller{contract: contract}, GravityTransactor: GravityTransactor{contract: contract}, GravityFilterer: GravityFilterer{contract: contract}}, nil
}

// NewGravityCaller creates a new read-only instance of Gravity, bound to a specific deployed contract.
func NewGravityCaller(address common.Address, caller bind.ContractCaller) (*GravityCaller, error) {
	contract, err := bindGravity(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GravityCaller{contract: contract}, nil
}

// NewGravityTransactor creates a new write-only instance of Gravity, bound to a specific deployed contract.
func NewGravityTransactor(address common.Address, transactor bind.ContractTransactor) (*GravityTransactor, error) {
	contract, err := bindGravity(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GravityTransactor{contract: contract}, nil
}

// NewGravityFilterer creates a new log filterer instance of Gravity, bound to a specific deployed contract.
func NewGravityFilterer(address common.Address, filterer bind.ContractFilterer) (*GravityFilterer, error) {
	contract, err := bindGravity(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GravityFilterer{contract: contract}, nil
}

// bindGravity binds a generic wrapper to an already deployed contract.
func bindGravity(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GravityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gravity *GravityRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Gravity.Contract.GravityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gravity *GravityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gravity.Contract.GravityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gravity *GravityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gravity.Contract.GravityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gravity *GravityCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Gravity.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gravity *GravityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gravity.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gravity *GravityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gravity.Contract.contract.Transact(opts, method, params...)
}

// BftValue is a free data retrieval call binding the contract method 0x3cec1bdd.
//
// Solidity: function bftValue() view returns(uint256)
func (_Gravity *GravityCaller) BftValue(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Gravity.contract.Call(opts, out, "bftValue")
	return *ret0, err
}

// BftValue is a free data retrieval call binding the contract method 0x3cec1bdd.
//
// Solidity: function bftValue() view returns(uint256)
func (_Gravity *GravitySession) BftValue() (*big.Int, error) {
	return _Gravity.Contract.BftValue(&_Gravity.CallOpts)
}

// BftValue is a free data retrieval call binding the contract method 0x3cec1bdd.
//
// Solidity: function bftValue() view returns(uint256)
func (_Gravity *GravityCallerSession) BftValue() (*big.Int, error) {
	return _Gravity.Contract.BftValue(&_Gravity.CallOpts)
}

// Consuls is a free data retrieval call binding the contract method 0xa2c2c617.
//
// Solidity: function consuls(uint256 ) view returns(address)
func (_Gravity *GravityCaller) Consuls(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Gravity.contract.Call(opts, out, "consuls", arg0)
	return *ret0, err
}

// Consuls is a free data retrieval call binding the contract method 0xa2c2c617.
//
// Solidity: function consuls(uint256 ) view returns(address)
func (_Gravity *GravitySession) Consuls(arg0 *big.Int) (common.Address, error) {
	return _Gravity.Contract.Consuls(&_Gravity.CallOpts, arg0)
}

// Consuls is a free data retrieval call binding the contract method 0xa2c2c617.
//
// Solidity: function consuls(uint256 ) view returns(address)
func (_Gravity *GravityCallerSession) Consuls(arg0 *big.Int) (common.Address, error) {
	return _Gravity.Contract.Consuls(&_Gravity.CallOpts, arg0)
}

// GetConsuls is a free data retrieval call binding the contract method 0xad595b1a.
//
// Solidity: function getConsuls() view returns(address[])
func (_Gravity *GravityCaller) GetConsuls(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Gravity.contract.Call(opts, out, "getConsuls")
	return *ret0, err
}

// GetConsuls is a free data retrieval call binding the contract method 0xad595b1a.
//
// Solidity: function getConsuls() view returns(address[])
func (_Gravity *GravitySession) GetConsuls() ([]common.Address, error) {
	return _Gravity.Contract.GetConsuls(&_Gravity.CallOpts)
}

// GetConsuls is a free data retrieval call binding the contract method 0xad595b1a.
//
// Solidity: function getConsuls() view returns(address[])
func (_Gravity *GravityCallerSession) GetConsuls() ([]common.Address, error) {
	return _Gravity.Contract.GetConsuls(&_Gravity.CallOpts)
}

// HashNewConsuls is a free data retrieval call binding the contract method 0x4dea5eba.
//
// Solidity: function hashNewConsuls(address[] newConsuls) pure returns(bytes32)
func (_Gravity *GravityCaller) HashNewConsuls(opts *bind.CallOpts, newConsuls []common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Gravity.contract.Call(opts, out, "hashNewConsuls", newConsuls)
	return *ret0, err
}

// HashNewConsuls is a free data retrieval call binding the contract method 0x4dea5eba.
//
// Solidity: function hashNewConsuls(address[] newConsuls) pure returns(bytes32)
func (_Gravity *GravitySession) HashNewConsuls(newConsuls []common.Address) ([32]byte, error) {
	return _Gravity.Contract.HashNewConsuls(&_Gravity.CallOpts, newConsuls)
}

// HashNewConsuls is a free data retrieval call binding the contract method 0x4dea5eba.
//
// Solidity: function hashNewConsuls(address[] newConsuls) pure returns(bytes32)
func (_Gravity *GravityCallerSession) HashNewConsuls(newConsuls []common.Address) ([32]byte, error) {
	return _Gravity.Contract.HashNewConsuls(&_Gravity.CallOpts, newConsuls)
}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(bool)
func (_Gravity *GravityCaller) Rounds(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Gravity.contract.Call(opts, out, "rounds", arg0)
	return *ret0, err
}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(bool)
func (_Gravity *GravitySession) Rounds(arg0 *big.Int) (bool, error) {
	return _Gravity.Contract.Rounds(&_Gravity.CallOpts, arg0)
}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(bool)
func (_Gravity *GravityCallerSession) Rounds(arg0 *big.Int) (bool, error) {
	return _Gravity.Contract.Rounds(&_Gravity.CallOpts, arg0)
}

// UpdateConsuls is a paid mutator transaction binding the contract method 0x92c388ab.
//
// Solidity: function updateConsuls(address[] newConsuls, uint8[] v, bytes32[] r, bytes32[] s, uint256 newLastRound) returns()
func (_Gravity *GravityTransactor) UpdateConsuls(opts *bind.TransactOpts, newConsuls []common.Address, v []uint8, r [][32]byte, s [][32]byte, newLastRound *big.Int) (*types.Transaction, error) {
	return _Gravity.contract.Transact(opts, "updateConsuls", newConsuls, v, r, s, newLastRound)
}

// UpdateConsuls is a paid mutator transaction binding the contract method 0x92c388ab.
//
// Solidity: function updateConsuls(address[] newConsuls, uint8[] v, bytes32[] r, bytes32[] s, uint256 newLastRound) returns()
func (_Gravity *GravitySession) UpdateConsuls(newConsuls []common.Address, v []uint8, r [][32]byte, s [][32]byte, newLastRound *big.Int) (*types.Transaction, error) {
	return _Gravity.Contract.UpdateConsuls(&_Gravity.TransactOpts, newConsuls, v, r, s, newLastRound)
}

// UpdateConsuls is a paid mutator transaction binding the contract method 0x92c388ab.
//
// Solidity: function updateConsuls(address[] newConsuls, uint8[] v, bytes32[] r, bytes32[] s, uint256 newLastRound) returns()
func (_Gravity *GravityTransactorSession) UpdateConsuls(newConsuls []common.Address, v []uint8, r [][32]byte, s [][32]byte, newLastRound *big.Int) (*types.Transaction, error) {
	return _Gravity.Contract.UpdateConsuls(&_Gravity.TransactOpts, newConsuls, v, r, s, newLastRound)
}

// ModelsABI is the input ABI used to generate the binding from.
const ModelsABI = "[]"

// ModelsBin is the compiled bytecode used for deploying new contracts.
var ModelsBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122084fdf702da12b33e22071d696066a96a6c69f9155dc2362034b4ae72b8e62a8f64736f6c63430007010033"

// DeployModels deploys a new Ethereum contract, binding an instance of Models to it.
func DeployModels(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Models, error) {
	parsed, err := abi.JSON(strings.NewReader(ModelsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ModelsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Models{ModelsCaller: ModelsCaller{contract: contract}, ModelsTransactor: ModelsTransactor{contract: contract}, ModelsFilterer: ModelsFilterer{contract: contract}}, nil
}

// Models is an auto generated Go binding around an Ethereum contract.
type Models struct {
	ModelsCaller     // Read-only binding to the contract
	ModelsTransactor // Write-only binding to the contract
	ModelsFilterer   // Log filterer for contract events
}

// ModelsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ModelsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ModelsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ModelsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ModelsSession struct {
	Contract     *Models           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModelsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ModelsCallerSession struct {
	Contract *ModelsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ModelsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ModelsTransactorSession struct {
	Contract     *ModelsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModelsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ModelsRaw struct {
	Contract *Models // Generic contract binding to access the raw methods on
}

// ModelsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ModelsCallerRaw struct {
	Contract *ModelsCaller // Generic read-only contract binding to access the raw methods on
}

// ModelsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ModelsTransactorRaw struct {
	Contract *ModelsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewModels creates a new instance of Models, bound to a specific deployed contract.
func NewModels(address common.Address, backend bind.ContractBackend) (*Models, error) {
	contract, err := bindModels(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Models{ModelsCaller: ModelsCaller{contract: contract}, ModelsTransactor: ModelsTransactor{contract: contract}, ModelsFilterer: ModelsFilterer{contract: contract}}, nil
}

// NewModelsCaller creates a new read-only instance of Models, bound to a specific deployed contract.
func NewModelsCaller(address common.Address, caller bind.ContractCaller) (*ModelsCaller, error) {
	contract, err := bindModels(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ModelsCaller{contract: contract}, nil
}

// NewModelsTransactor creates a new write-only instance of Models, bound to a specific deployed contract.
func NewModelsTransactor(address common.Address, transactor bind.ContractTransactor) (*ModelsTransactor, error) {
	contract, err := bindModels(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ModelsTransactor{contract: contract}, nil
}

// NewModelsFilterer creates a new log filterer instance of Models, bound to a specific deployed contract.
func NewModelsFilterer(address common.Address, filterer bind.ContractFilterer) (*ModelsFilterer, error) {
	contract, err := bindModels(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ModelsFilterer{contract: contract}, nil
}

// bindModels binds a generic wrapper to an already deployed contract.
func bindModels(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ModelsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Models *ModelsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Models.Contract.ModelsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Models *ModelsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Models.Contract.ModelsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Models *ModelsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Models.Contract.ModelsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Models *ModelsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Models.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Models *ModelsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Models.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Models *ModelsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Models.Contract.contract.Transact(opts, method, params...)
}

// QueueLibABI is the input ABI used to generate the binding from.
const QueueLibABI = "[]"

// QueueLibFuncSigs maps the 4-byte function signature to its string representation.
var QueueLibFuncSigs = map[string]string{
	"9d6ad84b": "drop(QueueLib.Queue storage,bytes32)",
	"a506d954": "push(QueueLib.Queue storage,bytes32)",
}

// QueueLibBin is the compiled bytecode used for deploying new contracts.
var QueueLibBin = "0x610198610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100405760003560e01c80639d6ad84b14610045578063a506d95414610077575b600080fd5b81801561005157600080fd5b506100756004803603604081101561006857600080fd5b50803590602001356100a7565b005b81801561008357600080fd5b506100756004803603604081101561009a57600080fd5b5080359060200135610114565b6000818152600383016020908152604080832054600286019092529091205481156100e457600082815260028501602052604090208190556100e8565b8084555b8015610106576000818152600385016020526040902082905561010e565b600184018290555b50505050565b8154610129578082556001820181905561015e565b600182018054600090815260028401602081815260408084208690558454868552600388018352818520559190528120558190555b505056fea26469706673582212203029a86e3996cd91fc9d88cb412f535c77b14a42a11d51369d9032d6481fef9c64736f6c63430007010033"

// DeployQueueLib deploys a new Ethereum contract, binding an instance of QueueLib to it.
func DeployQueueLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *QueueLib, error) {
	parsed, err := abi.JSON(strings.NewReader(QueueLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(QueueLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &QueueLib{QueueLibCaller: QueueLibCaller{contract: contract}, QueueLibTransactor: QueueLibTransactor{contract: contract}, QueueLibFilterer: QueueLibFilterer{contract: contract}}, nil
}

// QueueLib is an auto generated Go binding around an Ethereum contract.
type QueueLib struct {
	QueueLibCaller     // Read-only binding to the contract
	QueueLibTransactor // Write-only binding to the contract
	QueueLibFilterer   // Log filterer for contract events
}

// QueueLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type QueueLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QueueLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QueueLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QueueLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QueueLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QueueLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QueueLibSession struct {
	Contract     *QueueLib         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QueueLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QueueLibCallerSession struct {
	Contract *QueueLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// QueueLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QueueLibTransactorSession struct {
	Contract     *QueueLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// QueueLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type QueueLibRaw struct {
	Contract *QueueLib // Generic contract binding to access the raw methods on
}

// QueueLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QueueLibCallerRaw struct {
	Contract *QueueLibCaller // Generic read-only contract binding to access the raw methods on
}

// QueueLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QueueLibTransactorRaw struct {
	Contract *QueueLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQueueLib creates a new instance of QueueLib, bound to a specific deployed contract.
func NewQueueLib(address common.Address, backend bind.ContractBackend) (*QueueLib, error) {
	contract, err := bindQueueLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QueueLib{QueueLibCaller: QueueLibCaller{contract: contract}, QueueLibTransactor: QueueLibTransactor{contract: contract}, QueueLibFilterer: QueueLibFilterer{contract: contract}}, nil
}

// NewQueueLibCaller creates a new read-only instance of QueueLib, bound to a specific deployed contract.
func NewQueueLibCaller(address common.Address, caller bind.ContractCaller) (*QueueLibCaller, error) {
	contract, err := bindQueueLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QueueLibCaller{contract: contract}, nil
}

// NewQueueLibTransactor creates a new write-only instance of QueueLib, bound to a specific deployed contract.
func NewQueueLibTransactor(address common.Address, transactor bind.ContractTransactor) (*QueueLibTransactor, error) {
	contract, err := bindQueueLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QueueLibTransactor{contract: contract}, nil
}

// NewQueueLibFilterer creates a new log filterer instance of QueueLib, bound to a specific deployed contract.
func NewQueueLibFilterer(address common.Address, filterer bind.ContractFilterer) (*QueueLibFilterer, error) {
	contract, err := bindQueueLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QueueLibFilterer{contract: contract}, nil
}

// bindQueueLib binds a generic wrapper to an already deployed contract.
func bindQueueLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(QueueLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QueueLib *QueueLibRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _QueueLib.Contract.QueueLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QueueLib *QueueLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QueueLib.Contract.QueueLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QueueLib *QueueLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QueueLib.Contract.QueueLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QueueLib *QueueLibCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _QueueLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QueueLib *QueueLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QueueLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QueueLib *QueueLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QueueLib.Contract.contract.Transact(opts, method, params...)
}

