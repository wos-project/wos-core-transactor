package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/golang/glog"
	"github.com/spf13/viper"
	"github.com/wos-project/wos-core-transactor/app/chains"
	"github.com/wos-project/wos-core-transactor/app/models"
	"gorm.io/gorm"
)

// respTransactionQueuedItem is the enqueued tx that we must process
type respTransactionQueuedItem struct {
	Kind string                 `json:"kind" binding:"required"`
	Spec map[string]interface{} `json:"spec" binding:"required"`
}

// respTransactionQueuedItemAirdropErc721 is the enqueued Erc721 that we must process
type respTransactionQueuedItemAirdropErc721 struct {
	Uid         string `json:"uid" binding:"required"`
	CallbackUri string `json:"callbackUri" binding:"required"`
	WalletAddr  string `json:"walletAddr" binding:"required"`
	WalletKind  string `json:"walletKind" binding:"required"`
	IpfsCid     string `json:"ipfsCid" binding:"required"`
}

// respTransactionQueuedItemAirdropErc20 is the enqueued Erc20 that we must process
type respTransactionQueuedItemAirdropErc20 struct {
	Uid           string `json:"uid" binding:"required"`
	CallbackUri   string `json:"callbackUri" binding:"required"`
	WalletAddr    string `json:"walletAddr" binding:"required"`
	WalletKind    string `json:"walletKind" binding:"required"`
	TokenQuantity int    `json:"tokenQuantity" binding:"required"`
}

// reqTransactionQueuedItemCallback is the callback request back to requestor
type reqTransactionQueuedItemCallback struct {
	Uid           string `json:"uid" binding:"required"`
	TxId          string `json:"txId" binding:"required"`
	ContractAddr  string `json:"contractAddr" binding:"required"`
	Status        string `json:"status" binding:"required"`
	IpfsCid       string `json:"ipfsCid"`
	Cost          string `json:"cost"`
	TokenQuantity int    `json:"tokenQuantity"`
}

func performRequest(method, url string, body string) (*http.Response, error) {

	client := &http.Client{}

	req, _ := http.NewRequest(method, url, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(viper.GetString("services.tx.apiKey.key"), viper.GetString("services.tx.apiKey.value"))

	return client.Do(req)
}

// ServiceAirdrop processes airdrops
func ServiceAirdrop() {

	var cbBody reqTransactionQueuedItemCallback
	var err error
	var b []byte
	var contractAddressHex, txHashHex string
	var spec721 respTransactionQueuedItemAirdropErc721
	var spec20 respTransactionQueuedItemAirdropErc20
	var queuedItem respTransactionQueuedItem
	var tx models.Transaction
	var r *gorm.DB

	glog.V(2).Info("ServiceAirdrop: start")

	// get tx
	resp, err := performRequest("GET", viper.GetString("services.tx.uri"), "")
	if err != nil {
		err = fmt.Errorf("error getting airdrop job %s %v", viper.GetString("services.tx.uri"), err)
		glog.Error(err)
		goto Done
	}
	defer resp.Body.Close()
	if resp.StatusCode == 201 {
		glog.V(2).Info("no transactions")
		goto Done
	}
	glog.Info("processing transactions")
	if resp.StatusCode != 200 {
		err = fmt.Errorf("failure code requesting tx %v", resp.StatusCode)
		glog.Error(err)
		goto Done
	}
	b, _ = ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(b, &queuedItem)
	if err != nil {
		err = fmt.Errorf("error getting request job %v", err)
		glog.Error(err)
		goto Done
	}
	b, err = json.Marshal(queuedItem.Spec)
	if err != nil {
		err = fmt.Errorf("error marshaling request spec %v", err)
		glog.Error(err)
		goto Done
	}

	glog.Infof("processing tx %s", string(b))

	if queuedItem.Kind == "erc721" {
		err = json.Unmarshal([]byte(b), &spec721)
		if err != nil { glog.Error("error unmarshaling %v", err); goto Done}
		if spec721.CallbackUri == "" { glog.Error("error missing CallbackUri"); goto Done}
		if spec721.Uid == "" { glog.Error("error missing Uid"); goto Done}
		if spec721.WalletAddr == "" { glog.Error("error missing WalletAddr"); goto Done}
	//	if spec721.WalletKind == "" { glog.Error("error missing WalletKind"); goto Done}
		if spec721.IpfsCid == "" { glog.Error("error missing IpfsCid"); goto Done}
	} else if queuedItem.Kind == "erc20" {
		err = json.Unmarshal([]byte(b), &spec20)
		if err != nil { glog.Error("error unmarshaling %v", err); goto Done}
		if spec20.CallbackUri == "" { glog.Error("error missing CallbackUri"); goto Done}
		if spec20.Uid == "" { glog.Error("error missing Uid"); goto Done}
		if spec20.WalletAddr == "" { glog.Error("error missing WalletAddr"); goto Done}
	//	if spec20.WalletKind == "" { glog.Error("error missing WalletKind"); goto Done}
		if spec20.TokenQuantity == 0 { glog.Error("error missing TokenQuantity"); goto Done}
	} else {
		err = fmt.Errorf("error unmarshaling - unknown tx Kind %v", err)
		glog.Error(err)
		goto Done
	}

	if err != nil {
		err = fmt.Errorf("error unmarshaling request spec %v", err)
		glog.Error(err)
		goto Done
	}

	// record tx in db, especially cost
	tx = models.Transaction{
		Uid:    spec721.Uid,
		Status: "in-progress",
	}
	r = models.Db.Create(&tx)
	if r.Error != nil {
		glog.Errorf("cannot create tx in db %v", r.Error)
		goto Done
	}

	// TODO: use go routine

	if queuedItem.Kind == "erc721" {
		// mint NFT
		glog.Infof("minting nft wallet=%s ipfs=%s", spec721.WalletAddr, spec721.IpfsCid)
		contractAddressHex, txHashHex, _, err = chains.AirdropErc721(chains.ChainKind_Near, spec721.WalletAddr, spec721.IpfsCid)
		glog.Infof("minted nft contractAddresshex=%s txHashhex=%s", contractAddressHex, txHashHex)
	} else {
		// transfer token
		glog.Infof("token xfer wallet=%s quantity=%d", spec20.WalletAddr, spec20.TokenQuantity)
		txHashHex, _, err = chains.AirdropErc20(chains.ChainKind_Near, spec20.WalletAddr, int64(spec20.TokenQuantity))
	}

	if err != nil {
		tx.Status = fmt.Sprintf("error %v", err)
		r = models.Db.Save(&tx)
		if r.Error != nil {
			glog.Errorf("cannot save tx in db %v", r.Error)
		}
		err = fmt.Errorf("error airdropping 721 request %v", err)
		glog.Error(err)

		cbBody = reqTransactionQueuedItemCallback{
			Uid:          spec721.Uid,
			TxId:         txHashHex,
			ContractAddr: contractAddressHex,
			Status:       fmt.Sprintf("%v", err),
			IpfsCid:      spec721.IpfsCid,
		}
		b, _ = json.Marshal(cbBody)

		glog.Infof("calling wos back %v", string(b))

		resp, err = performRequest("POST", spec721.CallbackUri, string(b))
		if err != nil {
			err = fmt.Errorf("error cb %s %v", spec721.CallbackUri, err)
			glog.Error(err)
		}
		goto Done
	}

	tx.Status = "tx success"
	tx.TxHashHex = txHashHex
	// TODO: tx.Cost
	r = models.Db.Save(&tx)
	if r.Error != nil {
		glog.Errorf("cannot save tx in db %v", r.Error)
		goto Done
	}

	glog.Infof("transaction complete TxHash: %s", tx.TxHashHex)

	// finish by calling cb
	cbBody = reqTransactionQueuedItemCallback{
		Uid:          spec721.Uid,
		TxId:         txHashHex,
		ContractAddr: contractAddressHex,
		Status:       "success",
		IpfsCid:      spec721.IpfsCid,
	}
	b, _ = json.Marshal(cbBody)
	resp, err = performRequest("POST", spec721.CallbackUri, string(b))
	if err != nil {
		err = fmt.Errorf("error cb %s %v", spec721.CallbackUri, err)
		glog.Error(err)

		tx.Status = fmt.Sprintf("cb error %v", err)
		r = models.Db.Save(&tx)
		if r.Error != nil {
			glog.Errorf("cannot save tx in db %v", r.Error)
		}
		goto Done
	}

	tx.Status = "success"
	r = models.Db.Save(&tx)
	if r.Error != nil {
		glog.Errorf("cannot save tx in db %v", r.Error)
	}

	glog.Infof("callback complete uid:%s", tx.Uid)

Done:
	glog.V(2).Info("ServiceAirdrop: stop")
	return
}
