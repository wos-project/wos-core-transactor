package ethereumChain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ipfs/go-cid"
	"github.com/spf13/viper"
	"github.com/golang/glog"
)

type EthereumChain struct{}

// AirdropErc721 airdrops an NFT to wallet address recipientPublicHexAddress
func (e EthereumChain) AirdropErc721(chainUri string, spenderPrivateHexAddress string, recipientWalletID string, ipfsCid string) (contractAddressHex string, txHashHex string, cost *big.Int, err error) {

	ctx := context.Background()

	client, err := ethclient.Dial(chainUri)
	if err != nil {
		return "", "", big.NewInt(0), err
	}

	privateKey, err := crypto.HexToECDSA(spenderPrivateHexAddress)
	if err != nil {
		return "", "", big.NewInt(0), err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", "", big.NewInt(0), fmt.Errorf("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	fmt.Printf("NFT nonce1=%v\n", nonce)
	if err != nil {
		return "", "", big.NewInt(0), err
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return "", "", big.NewInt(0), err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return "", "", big.NewInt(0), err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(viper.GetInt64("chains.ethereum.gasLimit")) // in units
	auth.GasPrice = big.NewInt(viper.GetInt64("chains.ethereum.gasPrice"))

	glog.Infof("deploy NFT from=%s nonce=%v ipfs=%s", fromAddress, nonce, ipfsCid)

	contractAddress, _, instance, err := DeployERC721(auth, client)
	if err != nil {
		return "", "", big.NewInt(0), err
	}

	// TODO: wait for completion
	time.Sleep(2 * time.Second)
	nat, _ := client.NonceAt(ctx, fromAddress, nil)
	fmt.Printf("NFT nonce1 confirm=%v\n", nat)

	nonce, err = client.PendingNonceAt(ctx, fromAddress)
	fmt.Printf("NFT nonce2=%v\n", nonce)
	if err != nil {
		return "", "", big.NewInt(0), err
	}
	auth.Nonce = big.NewInt(int64(nonce))

	ipfSlice, err := cid.Decode(ipfsCid)
	if err != nil{
		return "", "", big.NewInt(0), fmt.Errorf("ipfs cid decode error1 %s %v", ipfsCid, err)
	}
	var ipfsArr [32]byte
	c := copy(ipfsArr[:], ipfSlice.Bytes())
	if c != 32 {
		return "", "", big.NewInt(0), fmt.Errorf("ipfs cid decode error2 %s %v", ipfsCid, c)
	}

	glog.Infof("mint NFT from=%s to=%s nonce=%v ipfs=%s", fromAddress, recipientWalletID, nonce, ipfsCid)

	reply, err := instance.Mint(auth, common.HexToAddress(recipientWalletID), ipfsArr)
	if err != nil {
		return "", "", big.NewInt(0), err
	}

	nat, _ = client.NonceAt(ctx, fromAddress, nil)
	fmt.Printf("NFT nonce2 confirm=%v\n", nat)

	return contractAddress.Hex(), reply.Hash().Hex(), reply.Cost(), nil
}

// AirdropErc20 drops tokens to wallet address recipientWalletID
// Good Ropsten contract address to use: 0x576Bf5838b6F91d6366Ef2228f7D6173e5102668
func (e EthereumChain) AirdropErc20(chainUri string, payerPrivateHexAddress string, recipientWalletID string, contractHexAddress string, quantity int64) (txHashHex string, cost *big.Int, err error) {

	client, err := ethclient.Dial(chainUri)
	if err != nil {
		return "", big.NewInt(0), err
	}

	privateKey, err := crypto.HexToECDSA(payerPrivateHexAddress)
	if err != nil {
		return "", big.NewInt(0), err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", big.NewInt(0), fmt.Errorf("invalid private key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	fmt.Printf("ERC-20 airdrop nonce=%v\n", nonce)
	if err != nil {
		return "", big.NewInt(0), err
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return "", big.NewInt(0), err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return "", big.NewInt(0), err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(viper.GetInt64("chains.ethereum.gasLimit")) // in units
	auth.GasPrice = big.NewInt(viper.GetInt64("chains.ethereum.gasPrice"))
  
	t, err := NewERC20Basic(common.HexToAddress(recipientWalletID), client)

	bigQuantity := big.NewInt(quantity)

	reply, err := t.Transfer(auth, common.HexToAddress(recipientWalletID), bigQuantity)

	if err != nil {
		return "", big.NewInt(0), err
	}

	return reply.Hash().Hex(), reply.Cost(), nil
}

// DeployErc20 creates an ERC20 contract
func (e EthereumChain) DeployErc20(chainUri string, payerPrivateHexAddress string, initialQuantity int64) (contractHexAddress string, cost *big.Int, err error) {

	client, err := ethclient.Dial(chainUri)
	if err != nil {
		return "", big.NewInt(0), err
	}

	privateKey, err := crypto.HexToECDSA(payerPrivateHexAddress)
	if err != nil {
		return "", big.NewInt(0), err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", big.NewInt(0), err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	fmt.Printf("ERC-20 deploy nonce=%v\n", nonce)
	if err != nil {
		return "", big.NewInt(0), err
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return "", big.NewInt(0), err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return "", big.NewInt(0), err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(viper.GetInt64("chains.ethereum.gasLimit")) // in units
	auth.GasPrice = big.NewInt(viper.GetInt64("chains.ethereum.gasPrice"))

	contractAddress, tx, _, err := DeployERC20Basic(auth, client, big.NewInt(initialQuantity))
	if err != nil {
		return "", big.NewInt(0), err
	}

	return contractAddress.Hex(), tx.Cost(), nil
}
