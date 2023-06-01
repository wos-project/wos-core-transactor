package nearChain

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

	x := viper.GetString("chains.near.contractAddress")
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

func TestNearChains(t *testing.T) {

	chain := NearChain{}

	// mint NFT
	nftContractAddress, txHashHex, cost, err := chain.AirdropErc721(
		"",
		viper.GetString("chains.near.senderAccountId"),
		"questori.testnet",
		GenerateRandomCid()) 

	assert.Nil(t, err)
	assert.True(t, len(txHashHex) > 0)
	assert.True(t, len(nftContractAddress) > 0)
	assert.Equal(t, 1, cost.Cmp(big.NewInt(0)))
	fmt.Printf("NFT contract addr=%s\n", nftContractAddress)
	fmt.Printf("NFT txhash=%s\n", txHashHex)
	fmt.Printf(`near view questori.testnet nft_tokens_for_owner '{"account_id": "'questori.testnet'"}'`)

	time.Sleep(2 * time.Second)

	// mint second NFT
	nftContractAddress, txHashHex, cost, err = chain.AirdropErc721(
		"",
		viper.GetString("chains.near.senderAccountId"),
		"questori.testnet",
		GenerateRandomCid()) 

	assert.Nil(t, err)
	assert.True(t, len(txHashHex) > 0)
	assert.True(t, len(nftContractAddress) > 0)
	assert.Equal(t, 1, cost.Cmp(big.NewInt(0)))
	fmt.Printf("NFT contract addr=%s\n", nftContractAddress)
	fmt.Printf("NFT txhash=%s\n", txHashHex)
}