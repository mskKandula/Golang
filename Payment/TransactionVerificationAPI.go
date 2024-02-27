package main

import (
	"strconv"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/corepkgv2/errormdl"
	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/corepkgv2/loggermdl"
	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/corepkgv2/servicebuildermdl"
)

var (
	privateKey = "ups1234567891231"
)

func TransactionVerificationAPI(ab *servicebuildermdl.AbstractBusinessLogicHolder) (map[string]interface{}, error) {
	// Data get from client side
	MQLRequestData, ok1 := ab.GetDataResultset("MQLRequestData")

	if !ok1 {
		loggermdl.LogError("Key not Found: MQLRequestData")
		return nil, errormdl.Wrap("required data not found")
	}

	clientID, _ := ab.GetGlobalConfigString("TilimiliPaymentClientID")
	purposeID, _ := ab.GetGlobalConfigString("LMSFeePaymentPurposeId")
	paymentModeValue, _ := ab.GetGlobalConfigString("TilimiliPaymentPaymentMode")
	requestFrom, _ := ab.GetGlobalConfigString("RequestFrom")
	transactionNumber := GetTransactionNumber(invoiceData.Get("invoiceNumber").String())

	clientId, clientIdErr := strconv.Atoi(clientID)
	if clientIdErr != nil {
		loggermdl.LogError("PrepareUpsParameter : error in getting clientId : ", clientIdErr)
		return nil, clientIdErr
	}
	// TODO:  transaction purpose id from Harshad Tidke sir
	transactionPurposeId, transactionPurposeIdErr := strconv.Atoi(purposeID)
	if transactionPurposeIdErr != nil {
		loggermdl.LogError("PrepareUpsParameter : error in getting transactionPurposeId : ", transactionPurposeIdErr)
		return nil, transactionPurposeIdErr
	}

	paymentMode, paymentModeErr := strconv.Atoi(paymentModeValue)
	if paymentModeErr != nil {
		loggermdl.LogError("PrepareUpsParameter : error in getting paymentMode : ", paymentModeErr)
		return nil, paymentModeErr
	}
}
