package api

import (
	"github.com/gorilla/mux"
	"github.com/svarlamov/bintrad/context"
	"github.com/svarlamov/bintrad/models"
	"github.com/svarlamov/bintrad/utils"
	"net/http"
	"strconv"
	"time"
)

func V0_API_Start_Contract_Session(w http.ResponseWriter, r *http.Request) {
	user, err := context.GetUserAndCatch(w, r)
	if err != nil {
		Log.WithField("error", err).Error("Error getting user from context")
		return
	}

	completeUser := models.CompleteUser{}
	err = completeUser.PopulateFromUser(user)
	if err != nil {
		utils.JSONInternalError(w, "Error creating a new contract session", "")
		return
	}
	if completeUser.CurrentBalance < 0 {
		utils.JSONBadRequestError(w, "Insufficient funds", "")
		return
	}

	contractPeriod := 1
	ticker := models.Ticker{}
	tickerData, finalTickId, err := ticker.GetRandomTickerAndDataSubset(contractPeriod)
	if err != nil {
		utils.JSONInternalError(w, "Error creating a new contract session", "")
		return
	}
	contractSession := models.ContractSession{
		UserId:      user.Id,
		TickerId:    ticker.Id,
		Period:      int64(contractPeriod),
		Price:       tickerData[len(tickerData)-1].Close,
		Ttl:         130,
		IsClosed:    false,
		DataStart:   tickerData[0].OpensAt,
		DataEnd:     tickerData[len(tickerData)-1].OpensAt,
		FinalTickId: finalTickId,
	}
	err = contractSession.Create()
	if err != nil {
		utils.JSONInternalError(w, "Error creating a new contract session", "")
		return
	}

	respObj := models.CreateContractSessionResponse{
		Id:         contractSession.Id,
		Price:      contractSession.Price,
		Ttl:        contractSession.Ttl,
		Period:     contractSession.Period,
		TickerData: tickerData,
		TickerType: ticker.Type,
	}
	utils.JSONSuccess(w, respObj, "Successfully started contract session")
}

func V0_API_Finalise_Contract_Session(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contractSessionIdStr := vars["sessionId"]
	contractSessionId, err := strconv.Atoi(contractSessionIdStr)
	if err != nil {
		utils.JSONBadRequestError(w, "Unknown contractSessionId", "")
		return
	}
	user, err := context.GetUserAndCatch(w, r)
	if err != nil {
		Log.WithField("error", err).Error("Error getting user from context")
		return
	}
	reqObj := models.FinaliseContractSessionRequest{}
	err = utils.JSONDecodeAndCatch(w, r, &reqObj)
	if err != nil {
		return
	}

	contractSession := models.ContractSession{
		Id:     int64(contractSessionId),
		UserId: user.Id,
	}
	err = contractSession.GetByIdAndUserId()
	if err != nil {
		utils.JSONNotFoundError(w, "Couldn't find contract session", "")
		return
	}
	if contractSession.CreatedAt.Add(time.Duration(contractSession.Ttl) * time.Second).Before(time.Now()) {
		utils.JSONDetailed(w, utils.APIResponse{Message: "Contract session has expired"}, 422)
		return
	}

	completeUser := models.CompleteUser{}
	err = completeUser.PopulateFromUser(user)
	if err != nil {
		utils.JSONInternalError(w, "Error creating a new contract session", "")
		return
	}
	if completeUser.CurrentBalance < reqObj.Bet {
		utils.JSONBadRequestError(w, "Insufficient funds", "")
		return
	}

	contract, err := contractSession.GenerateContractData(reqObj.Bet, reqObj.IsBullish, contractSession.FinalTickId)
	if err != nil {
		utils.JSONInternalError(w, "Error generating contract", "")
		return
	}
	err = contract.Create()
	if err != nil {
		utils.JSONInternalError(w, "Error generating contract", "")
		return
	}

	completeUser = models.CompleteUser{}
	err = completeUser.PopulateFromUser(user)
	if err != nil {
		utils.JSONInternalError(w, "Error getting user data", "")
		return
	}
	respObj := models.FinaliseContractSessionResponse{
		UserData:  completeUser,
		Bet:       contract.Bet,
		Price:     contract.Price,
		IsBullish: contract.IsBullish,
		IsCorrect: contract.IsCorrect,
		Return:    contract.Return,
	}
	// TODO: Send back contract data and user data
	utils.JSONSuccess(w, respObj, "Successfully closed contract")
}
