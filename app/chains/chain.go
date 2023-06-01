package chains

import (
	"fmt"
	"math/big"

	"github.com/spf13/viper"
	"github.com/wos-project/wos-core-transactor/app/chains/ethereum_chain"
	"github.com/wos-project/wos-core-transactor/app/chains/near_chain"
)

const ChainKind_Ethereum = "ethereum"
const ChainKind_Near = "near"

type ChainFunctions interface {

	// ERC20
	AirdropErc20(chainUri string, payerPrivateHexAddress string, toPublicHexAddress string, contractAddress string, quantity int64) (txHashHex string, cost *big.Int, err error)
	DeployErc20(chainUri string, payerPrivateHexAddress string, initialQuantity int64) (contractHexAddress string, cost *big.Int, err error)

	// ERC721
	AirdropErc721(chainUri string, payerPrivateHexAddress string, recipientHexAddress string, ipfsCid string) (contractAddressHex string, txHashHex string, cost *big.Int, err error)
}

func AirdropErc20(chainKind string, recipientHexAddress string, quantity int64) (txHashHex string, cost *big.Int, err error) {

	var c ethereumChain.EthereumChain
	if chainKind == ChainKind_Ethereum {
		c = ethereumChain.EthereumChain{}
	} else {
		return "", nil, fmt.Errorf("unknown chain type")
	}

	txHashHex, cost, err = _AirdropErc20(c, recipientHexAddress, quantity)
	return txHashHex, cost, err
}

func _AirdropErc20(cfuncs ChainFunctions, recipientHexAddress string, quantity int64) (txHashHex string, cost *big.Int, err error) {
	txHashHex, cost, err = cfuncs.AirdropErc20(
		viper.GetString("chains.ethereum.uri"),
		viper.GetString("chains.ethereum.privateWalletHexAddress"),
		recipientHexAddress,
		viper.GetString("chains.ethereum.erc20ContractHexAddress"),
		quantity) 
	return txHashHex, cost, err
}

func AirdropErc721(chainKind string, recipientHexAddress string, ipfsCid string) (contractAddressHex string, txHashHex string, cost *big.Int, err error) {

	if chainKind == ChainKind_Ethereum {
		c := ethereumChain.EthereumChain{}
		contractAddressHex, txHashHex, cost, err = c.AirdropErc721(
			viper.GetString("chains.ethereum.uri"),
			viper.GetString("chains.ethereum.privateWalletHexAddress"),
			recipientHexAddress, 
			ipfsCid)
		return contractAddressHex, txHashHex, cost, err
	} else if chainKind == ChainKind_Near {
		c := nearChain.NearChain{}
		contractAddressHex, txHashHex, cost, err = c.AirdropErc721(
			viper.GetString(""),
			viper.GetString("chains.near.senderAccountId"),
			recipientHexAddress, 
			ipfsCid)
		return contractAddressHex, txHashHex, cost, err
	} else {
		return "", "", nil, fmt.Errorf("unknown chain type")
	}
}
