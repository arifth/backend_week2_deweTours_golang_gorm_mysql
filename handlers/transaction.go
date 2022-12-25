package handlers

import (
	"encoding/json"
	"fmt"
	dto "gorm-imp/dto/result"
	transactiondto "gorm-imp/dto/transaction"
	tripdto "gorm-imp/dto/trip"
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
	dataCountry, _ := strconv.Atoi(r.FormValue("country"))
	dataNight, _ := strconv.Atoi(r.FormValue("night"))
	dataDay, _ := strconv.Atoi(r.FormValue("day"))
	dataPrice, _ := strconv.Atoi(r.FormValue("price"))
	dataQuota, _ := strconv.Atoi(r.FormValue("quota"))

	request := tripdto.CreateTripRequest{
		Title:          r.FormValue("title"),
		Country:        dataCountry,
		Accomodation:   r.FormValue("accomodation"),
		Transportation: r.FormValue("transportation"),
		Eat:            r.FormValue("eat"),
		Day:            dataDay,
		Night:          dataNight,
		DateTrip:       r.FormValue("date_trip"),
		Price:          dataPrice,
		Quota:          dataQuota,
		Description:    r.FormValue("description"),
		Image:          filename,
	}

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

	// trip := models.Trip{
	// 	Title:          request.Title,
	// 	CountryId:      request.Country,
	// 	Accomodation:   request.Transportation,
	// 	Transportation: request.Transportation,
	// 	Eat:            request.Eat,
	// 	Day:            request.Day,
	// 	Night:          request.Night,
	// 	DateTrip:       request.DateTrip,
	// 	Price:          request.Price,
	// 	Quota:          request.Quota,
	// 	Description:    request.Description,
	// 	Image:          filename,
	// }

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// trip, _ = h.TripRepository.FindSingleTrip(id)

	// fmt.Println(request.Country)
	// return

	trans := models.Transaction{}

	// check all field for emptieness

	// if request.Title != "" {
	// 	trans.Title = request.Title
	// }
	// if request.Accomodation != "" {
	// 	trans.Accomodation = request.Accomodation
	// }
	// if request.Country != 0 {
	// 	trans.CountryId = request.Country
	// }
	// if request.Transportation != "" {
	// 	trans.Transportation = request.Transportation
	// }
	// if request.Eat != "" {
	// 	trans.Eat = request.Eat
	// }
	// if request.Day != 0 {
	// 	trans.Day = request.Day
	// }
	// if request.Night != 0 {
	// 	trans.Night = request.Night
	// }
	// if request.DateTrip != "" {
	// 	trans.DateTrip = request.DateTrip
	// }
	// if request.Price != 0 {
	// 	trans.Price = request.Price
	// }
	// if request.Quota != 0 {
	// 	trans.Quota = request.Quota
	// }
	// if request.Description != "" {
	// 	trans.Description = request.Description
	// }
	// if request.Image != "" {
	// 	trans.Image = request.Image
	// }

	// fmt.Println(request)

	data, err := h.TransactionRepository.UpdateTrans(trans, id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())

	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)

}
