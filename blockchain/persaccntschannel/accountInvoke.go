package persaccntschannel

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (persAccntsChannelClient *CerberusClient) CreateAccount(accountObject []byte) ([]string, []byte, error) {

	// channel instance -> create
	err := persAccntsChannelClient.setupPersonAccountsChannelClient()

	if err != nil {
		return nil, nil, err
	}
	defer sdkInstance.Close()

	persAccntsChannelClient.channelClient, err = channel.New(persAccntsChannelClient.channelCtx)

	if err != nil {
		return nil, nil, err
	}

	// request -> prepare
	request := channel.Request{
		ChaincodeID: PersonAccountsChannelChainCode,
		Fcn:         "createAccount",
		Args:        [][]byte{accountObject},
	}

	response, err := persAccntsChannelClient.channelClient.Execute(request, channel.WithTargetEndpoints(AnchorPrSipher))

	if err != nil {
		return nil, nil, err
	}

	if response.ChaincodeStatus == 200 {
		fmt.Println("Successfully updated records.")
		fmt.Println("Transaction ID is: " + response.TransactionID)
	}

	return []string{"200", string(response.TransactionID)}, response.Payload, nil
}

func (persAccntsChannelClient *CerberusClient) DeleteAccount(publicId string) ([]string, []byte, error) {

	// channel instance -> create
	err := persAccntsChannelClient.setupPersonAccountsChannelClient()

	if err != nil {
		return nil, nil, err
	}

	defer sdkInstance.Close()

	persAccntsChannelClient.channelClient, err = channel.New(persAccntsChannelClient.channelCtx)

	if err != nil {
		return nil, nil, err
	}

	// request -> prepare
	request := channel.Request{
		ChaincodeID: PersonAccountsChannelChainCode,
		Fcn:         "deleteAccount",
		Args:        [][]byte{[]byte(publicId)},
	}

	response, err := persAccntsChannelClient.channelClient.Execute(request, channel.WithTargetEndpoints(AnchorPrSipher))

	if err != nil {
		return nil, nil, err
	}

	if response.ChaincodeStatus == 200 {
		fmt.Println("Successfully updated records.")
		fmt.Println("Transaction ID is: " + response.TransactionID)
	}

	return []string{"200", string(response.TransactionID)}, response.Payload, nil
}

func (persAccntsChannelClient *CerberusClient) UpdateRecords(updateType, publicId string, updateArgs []string) ([]string, []byte, error) {

	var args [][]byte

	args = append(args, []byte(updateType))
	args = append(args, []byte(publicId))

	for _, arg := range updateArgs {
		args = append(args, []byte(arg))
	}

	// channel instance -> create
	err := persAccntsChannelClient.setupPersonAccountsChannelClient()

	if err != nil {
		return nil, nil, err
	}
	defer sdkInstance.Close()

	persAccntsChannelClient.channelClient, err = channel.New(persAccntsChannelClient.channelCtx)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	// request -> prepare
	request := channel.Request{
		ChaincodeID: PersonAccountsChannelChainCode,
		Fcn:         "updateRecords",
		Args:        args,
	}

	response, err := persAccntsChannelClient.channelClient.Execute(request, channel.WithTargetEndpoints(AnchorPrSipher))

	if err != nil {
		return nil, nil, err
	}

	if response.ChaincodeStatus == 200 {
		fmt.Println("Successfully updated records.")
		fmt.Println("Transaction ID is: " + response.TransactionID)
	}

	return []string{"200", string(response.TransactionID)}, response.Payload, nil
}
