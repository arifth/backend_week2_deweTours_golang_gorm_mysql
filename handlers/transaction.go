package handlers

import (
	"encoding/json"
	"fmt"
	dto "gorm-imp/dto/result"
	transactiondto "gorm-imp/dto/transaction"
	"gorm-imp/models"
	"gorm-imp/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Transactionhandler struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *Transactionhandler {
	return &Transactionhandler{TransactionRepository}
}

func (h *Transactionhandler) FindTrans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	trans, err := h.TransactionRepository.FindTrans()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}
	if err != nil {
		fmt.Println("errornya adalah", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: trans}
	json.NewEncoder(w).Encode(response)
}

func (h *Transactionhandler) FindTran(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// user, err := h.UserRepository.GetUser(id)

	trip, err := h.TransactionRepository.FindTran(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{
		Code: http.StatusOK,
		Data: trip,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Transactionhandler) CreateTrans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContext := r.Context().Value("dataFile")

	// assign nama file ke variable filename

	filename := dataContext.(string)

	// NOTE: face error caused by key value in postman using whitespace after it , DONT DO THAT !!
	// fmt.Println(reflect.TypeOf(r.FormValue("day")))
	// fmt.Println(r.FormValue("night"))
	// fmt.Println(r.FormValue("date_trip"))
	// fmt.Println(reflect.TypeOf(r.FormValue("price")))
	// fmt.Println(r.FormValue("quota"))
	// fmt.Println(r.FormValue("description"))

	dataCounter, _ := strconv.Atoi(r.FormValue("counter_qty"))
	dataTotal, _ := strconv.Atoi(r.FormValue("total"))
	dataTrip, _ := strconv.Atoi(r.FormValue("trip_id"))

	request := transactiondto.CreateTransactionRequest{
		CounterQty: dataCounter,
		Total:      dataTotal,
		Status:     r.FormValue("status"),
		Attachment: filename,
		TripId:     dataTrip,
	}

	// fmt.Println(request)
	// return

	// validate request against struct form created
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// countryId := strconv.Atoi()

	trans := models.Transaction{
		// Model:      gorm.Model{},
		CounterQty: request.CounterQty,
		Total:      request.Total,
		Status:     request.Status,
		Attachment: request.Attachment,
		TripId:     request.TripId,
		// Trip:       models.TripResponse{},
		// UserId:     0,
		// User:       models.UserResponse{},
	}

	data, err := h.TransactionRepository.CreateTrans(trans)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())

	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)

}

// /* Handler for request ,logic => get all values from req body ,decode it , put it in user variable which it's datatype is user model,
// then write it to DB with UpdateUser() method
// then return response to user with succesCode and data written with NewEncoder().Encode()
// */

// // COMMENT: able to insert Name ,but Email and password isn't included
// // solved , caused by typo in dto.SuccessResult

// NOTE: create new value in database,instead updating it/ SOLVED

func (h *Transactionhandler) UpdateTrans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContext := r.Context().Value("dataFile")

	// assign nama file ke variable filename
	filename := dataContext.(string)

	// NOTE: face error caused by key value in postman using whitespace after it , DONT DO THAT !!

	// get data country convrt ke int
	dataQty, _ := strconv.Atoi(r.FormValue("counter_qty"))
	dataTotal, _ := strconv.Atoi(r.FormValue("total"))
	dataTrip, _ := strconv.Atoi(r.FormValue("trip_id"))
	// dataPrice, _ := strconv.Atoi(r.FormValue("price"))
	// dataQuota, _ := strconv.Atoi(r.FormValue("quota"))

	request := transactiondto.CreateTransactionRequest{
		CounterQty: dataQty,
		Total:      dataTotal,
		Status:     r.FormValue("status"),
		Attachment: filename,
		TripId:     dataTrip,
	}

	// fmt.Println(request)
	// return

	// if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	// validate request against struct form created

	// fmt.Println("baris ke 203 trip handler ", request)
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// countryId := strconv.Atoi()

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// fmt.Println(id)
	// return

	// trip, _ = h.TripRepository.FindSingleTrip(id)

	// fmt.Println(request.Country)
	// return

	trans := models.Transaction{}

	// check all field for emptieness

	if request.Attachment != "" {
		trans.Attachment = request.Attachment
	}
	if request.Status != "" {
		trans.Status = request.Status
	}

	if request.CounterQty != 0 {
		trans.CounterQty = request.CounterQty
	}

	if request.Total != 0 {
		trans.Total = request.Total
	}

	if request.TripId != 0 {
		trans.TripId = request.TripId
	}

	// fmt.Println(request)
	// return

	data, err := h.TransactionRepository.UpdateTrans(trans, id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())

	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)

}
