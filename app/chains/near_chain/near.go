package nearChain

import (
	"fmt"
	"math/big"
	"os/exec"
	"crypto/rand"
	"regexp"

	"github.com/spf13/viper"
)

type NearChain struct{}

// AirdropErc721 airdrops an NFT to wallet address recipientPublicHexAddress
func (e NearChain) AirdropErc721(chainUri string, spenderPrivateHexAddress string, recipientPublicHexAddress string, ipfsCid string) (contractAddressHex string, txHashHex string, cost *big.Int, err error) {

	// follow prerequisites in README, install node, near cli, build wasm contract, deploy
	contractAddress := viper.GetString("chains.near.contractAddress")
	gasLimit := viper.GetFloat64("chains.near.gasLimit")
	ipfsUri := fmt.Sprintf("https://ipfs.io/ipfs/%s", ipfsCid)
	nBig, err := rand.Int(rand.Reader, big.NewInt(100000000))
	if err != nil {
		return "", "", big.NewInt(0), fmt.Errorf("cannot generate random tokenId %v", err)
	}
	tokenId := nBig.Int64()

	jsonStr := fmt.Sprintf(`{"token_id": "%d", "receiver_id": "%s", "token_metadata": { "title": "Questori NFT", "description": "Questori NFT containing media and metadata", "media": "%s", "copies": 1}}`,
		tokenId, recipientPublicHexAddress, ipfsUri)

	// mint
	out, err := exec.Command("near", "call", contractAddress, "nft_mint", jsonStr, "--accountId", spenderPrivateHexAddress, "--deposit", fmt.Sprintf("%f", gasLimit)).Output()
	if err != nil {
		return "", "", big.NewInt(0), fmt.Errorf("near mint cli returned error %v", err)
	}

	// parse out and get Transaction Id: xxxx
	r := regexp.MustCompile(`Transaction Id (.*)`)
	m := r.FindStringSubmatch(string(out))

	if len(m) == 0 {
		return "", "", big.NewInt(0), fmt.Errorf("cannot find Transaction Id in Near tx, %s", out)
	}
	transactionId := m[1]

	// TODO: get cost, not sure where to find that???

	return contractAddress, transactionId, big.NewInt(1), nil
}

// AirdropErc20 drops tokens to wallet address recipientPublicHexAddress
// Good Ropsten contract address to use: 0x576Bf5838b6F91d6366Ef2228f7D6173e5102668
func (e NearChain) AirdropErc20(chainUri string, payerPrivateHexAddress string, recipientPublicHexAddress string, contractHexAddress string, quantity int64) (txHashHex string, cost *big.Int, err error) {

	return "", big.NewInt(0), fmt.Errorf("not supported")
}

// DeployErc20 creates an ERC20 contract
func (e NearChain) DeployErc20(chainUri string, payerPrivateHexAddress string, initialQuantity int64) (contractHexAddress string, cost *big.Int, err error) {

	return "", big.NewInt(0), fmt.Errorf("not supported")
}
