package core

import (
	"relayer/appchains/fisco/store"
	"strings"
)

// HandleInterchainRequest handles the interchain request
func (r *Relayer) HandleInterchainRequest(chainID string, request InterchainRequest, txHash string) error {
	r.Logger.Infof("got the interchain request on %s: %+v", chainID, request)



	request.TxHash = txHash

	callback := func(icRequestID string, response ResponseI) {
		r.Logger.Infof(
			"got the response of the interchain request on %s: %+v",
			r.HubChain.GetChainID(),
			response,
		)

		// TODO
		//mysql.OnInterchainRequestHandled()

		err := r.AppChains[chainID].SendResponse(request.ID, response)
		if err != nil {
			r.Logger.Errorf(
				"failed to send the response to %s: %s",
				chainID,
				err,
			)

			return
		}

		r.Logger.Infof(
			"response sent to %s successfully",
			chainID,
		)
	}

	reqInfo,err := r.HubChain.SendInterchainRequest(request, callback)
	if err != nil {

		if  ! strings.Contains(err.Error(),"duplicated request sequence"){
			store.InitRelayerTransRecord(request.ID,chainID,txHash,request.DestChainID,reqInfo.HubReqTxId,reqInfo.IcRequestId,store.TxStatus_Error,err.Error())
		}else {
			r.Logger.Infof("duplicated request sequence ! not record trans")
		}

		r.Logger.Errorf(
			"failed to handle the interchain request %+v on %s: %s",
			request,
			r.HubChain.GetChainID(),
			err,
		)

		return err
	}

	//mysql.OnInterchainRequestReceived(request.ID, chainID, txHash)
	store.InitRelayerTransRecord(request.ID,chainID,txHash,request.DestChainID,reqInfo.HubReqTxId,reqInfo.IcRequestId,store.TxStatus_Unknow,"")
	r.Logger.Infof("HandleInterchainRequest is End !!!")
	return nil
}
