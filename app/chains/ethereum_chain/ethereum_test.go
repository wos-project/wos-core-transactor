package ethereumChain

import (
	"fmt"
	"math/big"
	"flag"
	"crypto/rand"
	"time"

	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/ipfs/go-cid"
	mh "github.com/multiformats/go-multihash"

	"github.com/wos-project/wos-core-transactor/app/config"
)

var _ = func() bool {
	testing.Init()
	return true
}()

func init() {
	config.ConfigPath = flag.String("config", "../../config.yaml", "path to YAML config file")
	config.InitializeConfiguration()
	flag.Parse()

	x := viper.GetString("chains.ethereum.uri")
	fmt.Println(x)

}

// Generate ipfs style cid
func GenerateRandomCid() (string) {
	pref := cid.Prefix{
		Version: 0,
		Codec: cid.Raw,
		MhType: mh.SHA2_256,
		MhLength: -1, // default length
	}
	
	// And then feed it some data
	blk := make([]byte, 32)
	rand.Read(blk)

	c, err := pref.Sum(blk)
	if err != nil {
		return ""
	}
	return c.String()
}

func TestCids(t *testing.T) {

	s := GenerateRandomCid()

	cid, _ := cid.Decode(s)

	b := cid.Bytes()

	fmt.Println(b)

}

func TestChains(t *testing.T) {

	chain := EthereumChain{}

	/* SKIP FOR NOW
	// deploy an ERC-20
	contractHexAddress, cost, err := chain.DeployErc20(
		viper.GetString("chains.ethereum.uri"),
		viper.GetString("chains.ethereum.privateWalletHexAddress"),
		100)
	
	assert.Nil(t, err)
	assert.True(t, len(contractHexAddress) > 0)
	assert.Equal(t, 1, cost.Cmp(big.NewInt(0)))
	fmt.Printf("ERC-20 contract addr=%s\n", contractHexAddress)

	// transfer ERC-20
	txHashHex, cost, err := chain.AirdropErc20(
		viper.GetString("chains.ethereum.uri"),
		viper.GetString("chains.ethereum.privateWalletHexAddress"),
		"0xff4d7946CabE6662EEBc12d74db83194ca72d18d",
		contractHexAddress, 
		1)

	assert.Nil(t, err)
	assert.True(t, len(txHashHex) > 0)
	assert.Equal(t, 1, cost.Cmp(big.NewInt(0)))
	fmt.Printf("ERC-20 transfer cost=%v\n", cost)
	*/

	// mint NFT
	nftContractAddress, txHashHex, cost, err := chain.AirdropErc721(
		viper.GetString("chains.ethereum.uri"),
		viper.GetString("chains.ethereum.privateWalletHexAddress"),
		"0xff4d7946CabE6662EEBc12d74db83194ca72d18d",
		GenerateRandomCid()) 

	assert.Nil(t, err)
	assert.True(t, len(txHashHex) > 0)
	assert.True(t, len(nftContractAddress) > 0)
	assert.Equal(t, 1, cost.Cmp(big.NewInt(0)))
	fmt.Printf("NFT contract addr=%s\n", nftContractAddress)
	fmt.Printf("NFT txhash=%s\n", txHashHex)

	time.Sleep(2 * time.Second)

	// mint second NFT
	nftContractAddress, txHashHex, cost, err = chain.AirdropErc721(
		viper.GetString("chains.ethereum.uri"),
		viper.GetString("chains.ethereum.privateWalletHexAddress"),
		"0xff4d7946CabE6662EEBc12d74db83194ca72d18d",
		GenerateRandomCid()) 

	assert.Nil(t, err)
	assert.True(t, len(txHashHex) > 0)
	assert.True(t, len(nftContractAddress) > 0)
	assert.Equal(t, 1, cost.Cmp(big.NewInt(0)))
	fmt.Printf("NFT contract addr=%s\n", nftContractAddress)
	fmt.Printf("NFT txhash=%s\n", txHashHex)
}