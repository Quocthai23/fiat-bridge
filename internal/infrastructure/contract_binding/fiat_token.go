// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract_binding

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ContractBindingMetaData contains all meta data concerning the ContractBinding contract.
var ContractBindingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"coreTxId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FiatBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"coreTxId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FiatMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BURNER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_coreTxId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_coreTxId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"processedTxs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b506040518060400160405280600981526020017f564e4420546f6b656e00000000000000000000000000000000000000000000008152506040518060400160405280600481526020017f564e445400000000000000000000000000000000000000000000000000000000815250816003908161008b9190610454565b50806004908161009b9190610454565b5050506100b05f5f1b336100b660201b60201c565b50610523565b5f6100c783836101ac60201b60201c565b6101a257600160055f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555061013f61021060201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4600190506101a6565b5f90505b92915050565b5f60055f8481526020019081526020015f205f015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b5f33905090565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061029257607f821691505b6020821081036102a5576102a461024e565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026103077fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826102cc565b61031186836102cc565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61035561035061034b84610329565b610332565b610329565b9050919050565b5f819050919050565b61036e8361033b565b61038261037a8261035c565b8484546102d8565b825550505050565b5f5f905090565b61039961038a565b6103a4818484610365565b505050565b5b818110156103c7576103bc5f82610391565b6001810190506103aa565b5050565b601f82111561040c576103dd816102ab565b6103e6846102bd565b810160208510156103f5578190505b610409610401856102bd565b8301826103a9565b50505b505050565b5f82821c905092915050565b5f61042c5f1984600802610411565b1980831691505092915050565b5f610444838361041d565b9150826002028217905092915050565b61045d82610217565b67ffffffffffffffff81111561047657610475610221565b5b610480825461027b565b61048b8282856103cb565b5f60209050601f8311600181146104bc575f84156104aa578287015190505b6104b48582610439565b86555061051b565b601f1984166104ca866102ab565b5f5b828110156104f1578489015182556001820191506020850194506020810190506104cc565b8683101561050e578489015161050a601f89168261041d565b8355505b6001600288020188555050505b505050505050565b611c22806105305f395ff3fe608060405234801561000f575f5ffd5b5060043610610135575f3560e01c806370a08231116100b6578063a217fddf1161007a578063a217fddf14610373578063a9059cbb14610391578063b48272cc146103c1578063d5391393146103dd578063d547741f146103fb578063dd62ed3e1461041757610135565b806370a08231146102a95780637e8816b9146102d957806391d14854146102f55780639395e0eb1461032557806395d89b411461035557610135565b8063248a9ca3116100fd578063248a9ca314610205578063282c51f3146102355780632f2ff15d14610253578063313ce5671461026f57806336568abe1461028d57610135565b806301ffc9a71461013957806306fdde0314610169578063095ea7b31461018757806318160ddd146101b757806323b872dd146101d5575b5f5ffd5b610153600480360381019061014e9190611428565b610447565b604051610160919061146d565b60405180910390f35b6101716104c0565b60405161017e91906114f6565b60405180910390f35b6101a1600480360381019061019c91906115a3565b610550565b6040516101ae919061146d565b60405180910390f35b6101bf610572565b6040516101cc91906115f0565b60405180910390f35b6101ef60048036038101906101ea9190611609565b61057b565b6040516101fc919061146d565b60405180910390f35b61021f600480360381019061021a919061168c565b6105a9565b60405161022c91906116c6565b60405180910390f35b61023d6105c6565b60405161024a91906116c6565b60405180910390f35b61026d600480360381019061026891906116df565b6105ea565b005b61027761060c565b6040516102849190611738565b60405180910390f35b6102a760048036038101906102a291906116df565b610614565b005b6102c360048036038101906102be9190611751565b61068f565b6040516102d091906115f0565b60405180910390f35b6102f360048036038101906102ee91906118a8565b6106d4565b005b61030f600480360381019061030a91906116df565b610813565b60405161031c919061146d565b60405180910390f35b61033f600480360381019061033a9190611914565b610877565b60405161034c919061146d565b60405180910390f35b61035d6108ac565b60405161036a91906114f6565b60405180910390f35b61037b61093c565b60405161038891906116c6565b60405180910390f35b6103ab60048036038101906103a691906115a3565b610942565b6040516103b8919061146d565b60405180910390f35b6103db60048036038101906103d6919061195b565b610964565b005b6103e56109d6565b6040516103f291906116c6565b60405180910390f35b610415600480360381019061041091906116df565b6109fa565b005b610431600480360381019061042c91906119b5565b610a1c565b60405161043e91906115f0565b60405180910390f35b5f7f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806104b957506104b882610a9e565b5b9050919050565b6060600380546104cf90611a20565b80601f01602080910402602001604051908101604052809291908181526020018280546104fb90611a20565b80156105465780601f1061051d57610100808354040283529160200191610546565b820191905f5260205f20905b81548152906001019060200180831161052957829003601f168201915b5050505050905090565b5f5f61055a610b07565b9050610567818585610b0e565b600191505092915050565b5f600254905090565b5f5f610585610b07565b9050610592858285610b20565b61059d858585610bb3565b60019150509392505050565b5f60055f8381526020019081526020015f20600101549050919050565b7f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a84881565b6105f3826105a9565b6105fc81610ca3565b6106068383610cb7565b50505050565b5f6012905090565b61061c610b07565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610680576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61068a8282610da1565b505050565b5f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a66106fe81610ca3565b60068460405161070e9190611a8a565b90815260200160405180910390205f9054906101000a900460ff1615610769576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161076090611aea565b60405180910390fd5b600160068560405161077b9190611a8a565b90815260200160405180910390205f6101000a81548160ff0219169083151502179055506107a98383610e8b565b8273ffffffffffffffffffffffffffffffffffffffff16846040516107ce9190611a8a565b60405180910390207f594bab735eb042168b631578f373069c7d6779a70212b58a02e19ebc53edde758460405161080591906115f0565b60405180910390a350505050565b5f60055f8481526020019081526020015f205f015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b6006818051602081018201805184825260208301602085012081835280955050505050505f915054906101000a900460ff1681565b6060600480546108bb90611a20565b80601f01602080910402602001604051908101604052809291908181526020018280546108e790611a20565b80156109325780601f1061090957610100808354040283529160200191610932565b820191905f5260205f20905b81548152906001019060200180831161091557829003601f168201915b5050505050905090565b5f5f1b81565b5f5f61094c610b07565b9050610959818585610bb3565b600191505092915050565b61096e3382610f0a565b3373ffffffffffffffffffffffffffffffffffffffff16826040516109939190611a8a565b60405180910390207f8fa5ed2fb62bd641a7f1c107ec610ffd186f59f9ffdd2a710e1e2f7581803342836040516109ca91906115f0565b60405180910390a35050565b7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a681565b610a03826105a9565b610a0c81610ca3565b610a168383610da1565b50505050565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b5f33905090565b610b1b8383836001610f89565b505050565b5f610b2b8484610a1c565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811015610bad5781811015610b9e578281836040517ffb8f41b2000000000000000000000000000000000000000000000000000000008152600401610b9593929190611b17565b60405180910390fd5b610bac84848484035f610f89565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610c23575f6040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600401610c1a9190611b4c565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610c93575f6040517fec442f05000000000000000000000000000000000000000000000000000000008152600401610c8a9190611b4c565b60405180910390fd5b610c9e838383611158565b505050565b610cb481610caf610b07565b611371565b50565b5f610cc28383610813565b610d9757600160055f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550610d34610b07565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019050610d9b565b5f90505b92915050565b5f610dac8383610813565b15610e81575f60055f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550610e1e610b07565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a460019050610e85565b5f90505b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610efb575f6040517fec442f05000000000000000000000000000000000000000000000000000000008152600401610ef29190611b4c565b60405180910390fd5b610f065f8383611158565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610f7a575f6040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600401610f719190611b4c565b60405180910390fd5b610f85825f83611158565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610ff9575f6040517fe602df05000000000000000000000000000000000000000000000000000000008152600401610ff09190611b4c565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603611069575f6040517f94280d620000000000000000000000000000000000000000000000000000000081526004016110609190611b4c565b60405180910390fd5b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508015611152578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161114991906115f0565b60405180910390a35b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036111a8578060025f82825461119c9190611b92565b92505081905550611276565b5f5f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905081811015611231578381836040517fe450d38c00000000000000000000000000000000000000000000000000000000815260040161122893929190611b17565b60405180910390fd5b8181035f5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036112bd578060025f8282540392505081905550611307565b805f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161136491906115f0565b60405180910390a3505050565b61137b8282610813565b6113be5780826040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526004016113b5929190611bc5565b60405180910390fd5b5050565b5f604051905090565b5f5ffd5b5f5ffd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b611407816113d3565b8114611411575f5ffd5b50565b5f81359050611422816113fe565b92915050565b5f6020828403121561143d5761143c6113cb565b5b5f61144a84828501611414565b91505092915050565b5f8115159050919050565b61146781611453565b82525050565b5f6020820190506114805f83018461145e565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6114c882611486565b6114d28185611490565b93506114e28185602086016114a0565b6114eb816114ae565b840191505092915050565b5f6020820190508181035f83015261150e81846114be565b905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61153f82611516565b9050919050565b61154f81611535565b8114611559575f5ffd5b50565b5f8135905061156a81611546565b92915050565b5f819050919050565b61158281611570565b811461158c575f5ffd5b50565b5f8135905061159d81611579565b92915050565b5f5f604083850312156115b9576115b86113cb565b5b5f6115c68582860161155c565b92505060206115d78582860161158f565b9150509250929050565b6115ea81611570565b82525050565b5f6020820190506116035f8301846115e1565b92915050565b5f5f5f606084860312156116205761161f6113cb565b5b5f61162d8682870161155c565b935050602061163e8682870161155c565b925050604061164f8682870161158f565b9150509250925092565b5f819050919050565b61166b81611659565b8114611675575f5ffd5b50565b5f8135905061168681611662565b92915050565b5f602082840312156116a1576116a06113cb565b5b5f6116ae84828501611678565b91505092915050565b6116c081611659565b82525050565b5f6020820190506116d95f8301846116b7565b92915050565b5f5f604083850312156116f5576116f46113cb565b5b5f61170285828601611678565b92505060206117138582860161155c565b9150509250929050565b5f60ff82169050919050565b6117328161171d565b82525050565b5f60208201905061174b5f830184611729565b92915050565b5f60208284031215611766576117656113cb565b5b5f6117738482850161155c565b91505092915050565b5f5ffd5b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6117ba826114ae565b810181811067ffffffffffffffff821117156117d9576117d8611784565b5b80604052505050565b5f6117eb6113c2565b90506117f782826117b1565b919050565b5f67ffffffffffffffff82111561181657611815611784565b5b61181f826114ae565b9050602081019050919050565b828183375f83830152505050565b5f61184c611847846117fc565b6117e2565b90508281526020810184848401111561186857611867611780565b5b61187384828561182c565b509392505050565b5f82601f83011261188f5761188e61177c565b5b813561189f84826020860161183a565b91505092915050565b5f5f5f606084860312156118bf576118be6113cb565b5b5f84013567ffffffffffffffff8111156118dc576118db6113cf565b5b6118e88682870161187b565b93505060206118f98682870161155c565b925050604061190a8682870161158f565b9150509250925092565b5f60208284031215611929576119286113cb565b5b5f82013567ffffffffffffffff811115611946576119456113cf565b5b6119528482850161187b565b91505092915050565b5f5f60408385031215611971576119706113cb565b5b5f83013567ffffffffffffffff81111561198e5761198d6113cf565b5b61199a8582860161187b565b92505060206119ab8582860161158f565b9150509250929050565b5f5f604083850312156119cb576119ca6113cb565b5b5f6119d88582860161155c565b92505060206119e98582860161155c565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611a3757607f821691505b602082108103611a4a57611a496119f3565b5b50919050565b5f81905092915050565b5f611a6482611486565b611a6e8185611a50565b9350611a7e8185602086016114a0565b80840191505092915050565b5f611a958284611a5a565b915081905092915050565b7f5472616e73616374696f6e20616c72656164792070726f6365737365640000005f82015250565b5f611ad4601d83611490565b9150611adf82611aa0565b602082019050919050565b5f6020820190508181035f830152611b0181611ac8565b9050919050565b611b1181611535565b82525050565b5f606082019050611b2a5f830186611b08565b611b3760208301856115e1565b611b4460408301846115e1565b949350505050565b5f602082019050611b5f5f830184611b08565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611b9c82611570565b9150611ba783611570565b9250828201905080821115611bbf57611bbe611b65565b5b92915050565b5f604082019050611bd85f830185611b08565b611be560208301846116b7565b939250505056fea264697066735822122042fd292a2c5c9f55b2fdd3acb15609024afb681422cebb7667f190972848d04064736f6c634300081c0033",
}

// ContractBindingABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractBindingMetaData.ABI instead.
var ContractBindingABI = ContractBindingMetaData.ABI

// ContractBindingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractBindingMetaData.Bin instead.
var ContractBindingBin = ContractBindingMetaData.Bin

// DeployContractBinding deploys a new Ethereum contract, binding an instance of ContractBinding to it.
func DeployContractBinding(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractBinding, error) {
	parsed, err := ContractBindingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBindingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractBinding{ContractBindingCaller: ContractBindingCaller{contract: contract}, ContractBindingTransactor: ContractBindingTransactor{contract: contract}, ContractBindingFilterer: ContractBindingFilterer{contract: contract}}, nil
}

// ContractBinding is an auto generated Go binding around an Ethereum contract.
type ContractBinding struct {
	ContractBindingCaller     // Read-only binding to the contract
	ContractBindingTransactor // Write-only binding to the contract
	ContractBindingFilterer   // Log filterer for contract events
}

// ContractBindingCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractBindingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractBindingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractBindingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractBindingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractBindingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractBindingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractBindingSession struct {
	Contract     *ContractBinding  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractBindingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractBindingCallerSession struct {
	Contract *ContractBindingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ContractBindingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractBindingTransactorSession struct {
	Contract     *ContractBindingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ContractBindingRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractBindingRaw struct {
	Contract *ContractBinding // Generic contract binding to access the raw methods on
}

// ContractBindingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractBindingCallerRaw struct {
	Contract *ContractBindingCaller // Generic read-only contract binding to access the raw methods on
}

// ContractBindingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractBindingTransactorRaw struct {
	Contract *ContractBindingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractBinding creates a new instance of ContractBinding, bound to a specific deployed contract.
func NewContractBinding(address common.Address, backend bind.ContractBackend) (*ContractBinding, error) {
	contract, err := bindContractBinding(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractBinding{ContractBindingCaller: ContractBindingCaller{contract: contract}, ContractBindingTransactor: ContractBindingTransactor{contract: contract}, ContractBindingFilterer: ContractBindingFilterer{contract: contract}}, nil
}

// NewContractBindingCaller creates a new read-only instance of ContractBinding, bound to a specific deployed contract.
func NewContractBindingCaller(address common.Address, caller bind.ContractCaller) (*ContractBindingCaller, error) {
	contract, err := bindContractBinding(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractBindingCaller{contract: contract}, nil
}

// NewContractBindingTransactor creates a new write-only instance of ContractBinding, bound to a specific deployed contract.
func NewContractBindingTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractBindingTransactor, error) {
	contract, err := bindContractBinding(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractBindingTransactor{contract: contract}, nil
}

// NewContractBindingFilterer creates a new log filterer instance of ContractBinding, bound to a specific deployed contract.
func NewContractBindingFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractBindingFilterer, error) {
	contract, err := bindContractBinding(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractBindingFilterer{contract: contract}, nil
}

// bindContractBinding binds a generic wrapper to an already deployed contract.
func bindContractBinding(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractBindingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractBinding *ContractBindingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractBinding.Contract.ContractBindingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractBinding *ContractBindingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractBinding.Contract.ContractBindingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractBinding *ContractBindingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractBinding.Contract.ContractBindingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractBinding *ContractBindingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractBinding.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractBinding *ContractBindingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractBinding.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractBinding *ContractBindingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractBinding.Contract.contract.Transact(opts, method, params...)
}

// BURNERROLE is a free data retrieval call binding the contract method 0x282c51f3.
//
// Solidity: function BURNER_ROLE() view returns(bytes32)
func (_ContractBinding *ContractBindingCaller) BURNERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractBinding.contract.Call(opts, &out, "BURNER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BURNERROLE is a free data retrieval call binding the contract method 0x282c51f3.
//
// Solidity: function BURNER_ROLE() view returns(bytes32)
func (_ContractBinding *ContractBindingSession) BURNERROLE() ([32]byte, error) {
	return _ContractBinding.Contract.BURNERROLE(&_ContractBinding.CallOpts)
}

// BURNERROLE is a free data retrieval call binding the contract method 0x282c51f3.
//
// Solidity: function BURNER_ROLE() view returns(bytes32)
func (_ContractBinding *ContractBindingCallerSession) BURNERROLE() ([32]byte, error) {
	return _ContractBinding.Contract.BURNERROLE(&_ContractBinding.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ContractBinding *ContractBindingCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractBinding.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ContractBinding *ContractBindingSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ContractBinding.Contract.DEFAULTADMINROLE(&_ContractBinding.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ContractBinding *ContractBindingCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ContractBinding.Contract.DEFAULTADMINROLE(&_ContractBinding.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ContractBinding *ContractBindingCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractBinding.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ContractBinding *ContractBindingSession) MINTERROLE() ([32]byte, error) {
	return _ContractBinding.Contract.MINTERROLE(&_ContractBinding.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ContractBinding *ContractBindingCallerSession) MINTERROLE() ([32]byte, error) {
	return _ContractBinding.Contract.MINTERROLE(&_ContractBinding.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ContractBinding *ContractBindingCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractBinding.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ContractBinding *ContractBindingSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ContractBinding.Contract.Allowance(&_ContractBinding.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ContractBinding *ContractBindingCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ContractBinding.Contract.Allowance(&_ContractBinding.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ContractBinding *ContractBindingCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractBinding.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ContractBinding *ContractBindingSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ContractBinding.Contract.BalanceOf(&_ContractBinding.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ContractBinding *ContractBindingCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ContractBinding.Contract.BalanceOf(&_ContractBinding.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ContractBinding *ContractBindingCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ContractBinding.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ContractBinding *ContractBindingSession) Decimals() (uint8, error) {
	return _ContractBinding.Contract.Decimals(&_ContractBinding.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ContractBinding *ContractBindingCallerSession) Decimals() (uint8, error) {
	return _ContractBinding.Contract.Decimals(&_ContractBinding.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ContractBinding *ContractBindingCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ContractBinding.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ContractBinding *ContractBindingSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ContractBinding.Contract.GetRoleAdmin(&_ContractBinding.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ContractBinding *ContractBindingCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ContractBinding.Contract.GetRoleAdmin(&_ContractBinding.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ContractBinding *ContractBindingCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ContractBinding.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ContractBinding *ContractBindingSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ContractBinding.Contract.HasRole(&_ContractBinding.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ContractBinding *ContractBindingCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ContractBinding.Contract.HasRole(&_ContractBinding.CallOpts, role, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ContractBinding *ContractBindingCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractBinding.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ContractBinding *ContractBindingSession) Name() (string, error) {
	return _ContractBinding.Contract.Name(&_ContractBinding.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ContractBinding *ContractBindingCallerSession) Name() (string, error) {
	return _ContractBinding.Contract.Name(&_ContractBinding.CallOpts)
}

// ProcessedTxs is a free data retrieval call binding the contract method 0x9395e0eb.
//
// Solidity: function processedTxs(string ) view returns(bool)
func (_ContractBinding *ContractBindingCaller) ProcessedTxs(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _ContractBinding.contract.Call(opts, &out, "processedTxs", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProcessedTxs is a free data retrieval call binding the contract method 0x9395e0eb.
//
// Solidity: function processedTxs(string ) view returns(bool)
func (_ContractBinding *ContractBindingSession) ProcessedTxs(arg0 string) (bool, error) {
	return _ContractBinding.Contract.ProcessedTxs(&_ContractBinding.CallOpts, arg0)
}

// ProcessedTxs is a free data retrieval call binding the contract method 0x9395e0eb.
//
// Solidity: function processedTxs(string ) view returns(bool)
func (_ContractBinding *ContractBindingCallerSession) ProcessedTxs(arg0 string) (bool, error) {
	return _ContractBinding.Contract.ProcessedTxs(&_ContractBinding.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractBinding *ContractBindingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ContractBinding.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractBinding *ContractBindingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ContractBinding.Contract.SupportsInterface(&_ContractBinding.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ContractBinding *ContractBindingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ContractBinding.Contract.SupportsInterface(&_ContractBinding.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ContractBinding *ContractBindingCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractBinding.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ContractBinding *ContractBindingSession) Symbol() (string, error) {
	return _ContractBinding.Contract.Symbol(&_ContractBinding.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ContractBinding *ContractBindingCallerSession) Symbol() (string, error) {
	return _ContractBinding.Contract.Symbol(&_ContractBinding.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ContractBinding *ContractBindingCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractBinding.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ContractBinding *ContractBindingSession) TotalSupply() (*big.Int, error) {
	return _ContractBinding.Contract.TotalSupply(&_ContractBinding.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ContractBinding *ContractBindingCallerSession) TotalSupply() (*big.Int, error) {
	return _ContractBinding.Contract.TotalSupply(&_ContractBinding.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ContractBinding *ContractBindingTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ContractBinding.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ContractBinding *ContractBindingSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ContractBinding.Contract.Approve(&_ContractBinding.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ContractBinding *ContractBindingTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ContractBinding.Contract.Approve(&_ContractBinding.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0xb48272cc.
//
// Solidity: function burn(string _coreTxId, uint256 _amount) returns()
func (_ContractBinding *ContractBindingTransactor) Burn(opts *bind.TransactOpts, _coreTxId string, _amount *big.Int) (*types.Transaction, error) {
	return _ContractBinding.contract.Transact(opts, "burn", _coreTxId, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0xb48272cc.
//
// Solidity: function burn(string _coreTxId, uint256 _amount) returns()
func (_ContractBinding *ContractBindingSession) Burn(_coreTxId string, _amount *big.Int) (*types.Transaction, error) {
	return _ContractBinding.Contract.Burn(&_ContractBinding.TransactOpts, _coreTxId, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0xb48272cc.
//
// Solidity: function burn(string _coreTxId, uint256 _amount) returns()
func (_ContractBinding *ContractBindingTransactorSession) Burn(_coreTxId string, _amount *big.Int) (*types.Transaction, error) {
	return _ContractBinding.Contract.Burn(&_ContractBinding.TransactOpts, _coreTxId, _amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ContractBinding *ContractBindingTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractBinding.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ContractBinding *ContractBindingSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractBinding.Contract.GrantRole(&_ContractBinding.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ContractBinding *ContractBindingTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractBinding.Contract.GrantRole(&_ContractBinding.TransactOpts, role, account)
}

// Mint is a paid mutator transaction binding the contract method 0x7e8816b9.
//
// Solidity: function mint(string _coreTxId, address _to, uint256 _amount) returns()
func (_ContractBinding *ContractBindingTransactor) Mint(opts *bind.TransactOpts, _coreTxId string, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ContractBinding.contract.Transact(opts, "mint", _coreTxId, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x7e8816b9.
//
// Solidity: function mint(string _coreTxId, address _to, uint256 _amount) returns()
func (_ContractBinding *ContractBindingSession) Mint(_coreTxId string, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ContractBinding.Contract.Mint(&_ContractBinding.TransactOpts, _coreTxId, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x7e8816b9.
//
// Solidity: function mint(string _coreTxId, address _to, uint256 _amount) returns()
func (_ContractBinding *ContractBindingTransactorSession) Mint(_coreTxId string, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ContractBinding.Contract.Mint(&_ContractBinding.TransactOpts, _coreTxId, _to, _amount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ContractBinding *ContractBindingTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ContractBinding.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ContractBinding *ContractBindingSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ContractBinding.Contract.RenounceRole(&_ContractBinding.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ContractBinding *ContractBindingTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ContractBinding.Contract.RenounceRole(&_ContractBinding.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ContractBinding *ContractBindingTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractBinding.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ContractBinding *ContractBindingSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractBinding.Contract.RevokeRole(&_ContractBinding.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ContractBinding *ContractBindingTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ContractBinding.Contract.RevokeRole(&_ContractBinding.TransactOpts, role, account)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ContractBinding *ContractBindingTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ContractBinding.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ContractBinding *ContractBindingSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ContractBinding.Contract.Transfer(&_ContractBinding.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ContractBinding *ContractBindingTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ContractBinding.Contract.Transfer(&_ContractBinding.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ContractBinding *ContractBindingTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ContractBinding.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ContractBinding *ContractBindingSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ContractBinding.Contract.TransferFrom(&_ContractBinding.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ContractBinding *ContractBindingTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ContractBinding.Contract.TransferFrom(&_ContractBinding.TransactOpts, from, to, value)
}

// ContractBindingApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ContractBinding contract.
type ContractBindingApprovalIterator struct {
	Event *ContractBindingApproval // Event containing the contract specifics and raw log

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
func (it *ContractBindingApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBindingApproval)
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
		it.Event = new(ContractBindingApproval)
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
func (it *ContractBindingApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBindingApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBindingApproval represents a Approval event raised by the ContractBinding contract.
type ContractBindingApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ContractBinding *ContractBindingFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ContractBindingApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ContractBinding.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ContractBindingApprovalIterator{contract: _ContractBinding.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ContractBinding *ContractBindingFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ContractBindingApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ContractBinding.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBindingApproval)
				if err := _ContractBinding.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ContractBinding *ContractBindingFilterer) ParseApproval(log types.Log) (*ContractBindingApproval, error) {
	event := new(ContractBindingApproval)
	if err := _ContractBinding.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractBindingFiatBurnedIterator is returned from FilterFiatBurned and is used to iterate over the raw logs and unpacked data for FiatBurned events raised by the ContractBinding contract.
type ContractBindingFiatBurnedIterator struct {
	Event *ContractBindingFiatBurned // Event containing the contract specifics and raw log

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
func (it *ContractBindingFiatBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBindingFiatBurned)
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
		it.Event = new(ContractBindingFiatBurned)
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
func (it *ContractBindingFiatBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBindingFiatBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBindingFiatBurned represents a FiatBurned event raised by the ContractBinding contract.
type ContractBindingFiatBurned struct {
	CoreTxId common.Hash
	From     common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFiatBurned is a free log retrieval operation binding the contract event 0x8fa5ed2fb62bd641a7f1c107ec610ffd186f59f9ffdd2a710e1e2f7581803342.
//
// Solidity: event FiatBurned(string indexed coreTxId, address indexed from, uint256 amount)
func (_ContractBinding *ContractBindingFilterer) FilterFiatBurned(opts *bind.FilterOpts, coreTxId []string, from []common.Address) (*ContractBindingFiatBurnedIterator, error) {

	var coreTxIdRule []interface{}
	for _, coreTxIdItem := range coreTxId {
		coreTxIdRule = append(coreTxIdRule, coreTxIdItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ContractBinding.contract.FilterLogs(opts, "FiatBurned", coreTxIdRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &ContractBindingFiatBurnedIterator{contract: _ContractBinding.contract, event: "FiatBurned", logs: logs, sub: sub}, nil
}

// WatchFiatBurned is a free log subscription operation binding the contract event 0x8fa5ed2fb62bd641a7f1c107ec610ffd186f59f9ffdd2a710e1e2f7581803342.
//
// Solidity: event FiatBurned(string indexed coreTxId, address indexed from, uint256 amount)
func (_ContractBinding *ContractBindingFilterer) WatchFiatBurned(opts *bind.WatchOpts, sink chan<- *ContractBindingFiatBurned, coreTxId []string, from []common.Address) (event.Subscription, error) {

	var coreTxIdRule []interface{}
	for _, coreTxIdItem := range coreTxId {
		coreTxIdRule = append(coreTxIdRule, coreTxIdItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ContractBinding.contract.WatchLogs(opts, "FiatBurned", coreTxIdRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBindingFiatBurned)
				if err := _ContractBinding.contract.UnpackLog(event, "FiatBurned", log); err != nil {
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

// ParseFiatBurned is a log parse operation binding the contract event 0x8fa5ed2fb62bd641a7f1c107ec610ffd186f59f9ffdd2a710e1e2f7581803342.
//
// Solidity: event FiatBurned(string indexed coreTxId, address indexed from, uint256 amount)
func (_ContractBinding *ContractBindingFilterer) ParseFiatBurned(log types.Log) (*ContractBindingFiatBurned, error) {
	event := new(ContractBindingFiatBurned)
	if err := _ContractBinding.contract.UnpackLog(event, "FiatBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractBindingFiatMintedIterator is returned from FilterFiatMinted and is used to iterate over the raw logs and unpacked data for FiatMinted events raised by the ContractBinding contract.
type ContractBindingFiatMintedIterator struct {
	Event *ContractBindingFiatMinted // Event containing the contract specifics and raw log

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
func (it *ContractBindingFiatMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBindingFiatMinted)
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
		it.Event = new(ContractBindingFiatMinted)
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
func (it *ContractBindingFiatMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBindingFiatMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBindingFiatMinted represents a FiatMinted event raised by the ContractBinding contract.
type ContractBindingFiatMinted struct {
	CoreTxId common.Hash
	To       common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFiatMinted is a free log retrieval operation binding the contract event 0x594bab735eb042168b631578f373069c7d6779a70212b58a02e19ebc53edde75.
//
// Solidity: event FiatMinted(string indexed coreTxId, address indexed to, uint256 amount)
func (_ContractBinding *ContractBindingFilterer) FilterFiatMinted(opts *bind.FilterOpts, coreTxId []string, to []common.Address) (*ContractBindingFiatMintedIterator, error) {

	var coreTxIdRule []interface{}
	for _, coreTxIdItem := range coreTxId {
		coreTxIdRule = append(coreTxIdRule, coreTxIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ContractBinding.contract.FilterLogs(opts, "FiatMinted", coreTxIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractBindingFiatMintedIterator{contract: _ContractBinding.contract, event: "FiatMinted", logs: logs, sub: sub}, nil
}

// WatchFiatMinted is a free log subscription operation binding the contract event 0x594bab735eb042168b631578f373069c7d6779a70212b58a02e19ebc53edde75.
//
// Solidity: event FiatMinted(string indexed coreTxId, address indexed to, uint256 amount)
func (_ContractBinding *ContractBindingFilterer) WatchFiatMinted(opts *bind.WatchOpts, sink chan<- *ContractBindingFiatMinted, coreTxId []string, to []common.Address) (event.Subscription, error) {

	var coreTxIdRule []interface{}
	for _, coreTxIdItem := range coreTxId {
		coreTxIdRule = append(coreTxIdRule, coreTxIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ContractBinding.contract.WatchLogs(opts, "FiatMinted", coreTxIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBindingFiatMinted)
				if err := _ContractBinding.contract.UnpackLog(event, "FiatMinted", log); err != nil {
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

// ParseFiatMinted is a log parse operation binding the contract event 0x594bab735eb042168b631578f373069c7d6779a70212b58a02e19ebc53edde75.
//
// Solidity: event FiatMinted(string indexed coreTxId, address indexed to, uint256 amount)
func (_ContractBinding *ContractBindingFilterer) ParseFiatMinted(log types.Log) (*ContractBindingFiatMinted, error) {
	event := new(ContractBindingFiatMinted)
	if err := _ContractBinding.contract.UnpackLog(event, "FiatMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractBindingRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ContractBinding contract.
type ContractBindingRoleAdminChangedIterator struct {
	Event *ContractBindingRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ContractBindingRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBindingRoleAdminChanged)
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
		it.Event = new(ContractBindingRoleAdminChanged)
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
func (it *ContractBindingRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBindingRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBindingRoleAdminChanged represents a RoleAdminChanged event raised by the ContractBinding contract.
type ContractBindingRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ContractBinding *ContractBindingFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ContractBindingRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ContractBinding.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ContractBindingRoleAdminChangedIterator{contract: _ContractBinding.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ContractBinding *ContractBindingFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ContractBindingRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ContractBinding.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBindingRoleAdminChanged)
				if err := _ContractBinding.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ContractBinding *ContractBindingFilterer) ParseRoleAdminChanged(log types.Log) (*ContractBindingRoleAdminChanged, error) {
	event := new(ContractBindingRoleAdminChanged)
	if err := _ContractBinding.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractBindingRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ContractBinding contract.
type ContractBindingRoleGrantedIterator struct {
	Event *ContractBindingRoleGranted // Event containing the contract specifics and raw log

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
func (it *ContractBindingRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBindingRoleGranted)
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
		it.Event = new(ContractBindingRoleGranted)
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
func (it *ContractBindingRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBindingRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBindingRoleGranted represents a RoleGranted event raised by the ContractBinding contract.
type ContractBindingRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractBinding *ContractBindingFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractBindingRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ContractBinding.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractBindingRoleGrantedIterator{contract: _ContractBinding.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractBinding *ContractBindingFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ContractBindingRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ContractBinding.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBindingRoleGranted)
				if err := _ContractBinding.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractBinding *ContractBindingFilterer) ParseRoleGranted(log types.Log) (*ContractBindingRoleGranted, error) {
	event := new(ContractBindingRoleGranted)
	if err := _ContractBinding.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractBindingRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ContractBinding contract.
type ContractBindingRoleRevokedIterator struct {
	Event *ContractBindingRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ContractBindingRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBindingRoleRevoked)
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
		it.Event = new(ContractBindingRoleRevoked)
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
func (it *ContractBindingRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBindingRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBindingRoleRevoked represents a RoleRevoked event raised by the ContractBinding contract.
type ContractBindingRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractBinding *ContractBindingFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractBindingRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ContractBinding.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractBindingRoleRevokedIterator{contract: _ContractBinding.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractBinding *ContractBindingFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ContractBindingRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ContractBinding.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBindingRoleRevoked)
				if err := _ContractBinding.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ContractBinding *ContractBindingFilterer) ParseRoleRevoked(log types.Log) (*ContractBindingRoleRevoked, error) {
	event := new(ContractBindingRoleRevoked)
	if err := _ContractBinding.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractBindingTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ContractBinding contract.
type ContractBindingTransferIterator struct {
	Event *ContractBindingTransfer // Event containing the contract specifics and raw log

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
func (it *ContractBindingTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBindingTransfer)
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
		it.Event = new(ContractBindingTransfer)
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
func (it *ContractBindingTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBindingTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBindingTransfer represents a Transfer event raised by the ContractBinding contract.
type ContractBindingTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ContractBinding *ContractBindingFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ContractBindingTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ContractBinding.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractBindingTransferIterator{contract: _ContractBinding.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ContractBinding *ContractBindingFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ContractBindingTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ContractBinding.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBindingTransfer)
				if err := _ContractBinding.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ContractBinding *ContractBindingFilterer) ParseTransfer(log types.Log) (*ContractBindingTransfer, error) {
	event := new(ContractBindingTransfer)
	if err := _ContractBinding.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
