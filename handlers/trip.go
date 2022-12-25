package handlers

import (
	"encoding/json"
	"fmt"
	dto "gorm-imp/dto/result"
	tripdto "gorm-imp/dto/trip"
	"gorm-imp/models"
	"gorm-imp/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Triphandler struct {
	TripRepository repositories.TripRepository
}

func HandlerTrip(TripRepository repositories.TripRepository) *Triphandler {
	return &Triphandler{TripRepository}
}

func (h *Triphandler) FindTrips(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	trip, err := h.TripRepository.FindTrip()
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
	response := dto.SuccessResult{Code: http.StatusOK, Data: trip}
	json.NewEncoder(w).Encode(response)
}

func (h *Triphandler) FindTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// user, err := h.UserRepository.GetUser(id)

	trip, err := h.TripRepository.FindSingleTrip(id)

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

func (h *Triphandler) CreateTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	// userId := int(userInfp["id"].(float64))

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
		// Image:          filename,
	}

	// if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

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

	trip := models.Trip{
		Title:          request.Title,
		Country_id:     request.Country,
		Accomodation:   request.Transportation,
		Transportation: request.Transportation,
		Eat:            request.Eat,
		Day:            request.Day,
		Night:          request.Night,
		DateTrip:       request.DateTrip,
		Price:          request.Price,
		Quota:          request.Quota,
		Description:    request.Description,
		Image:          filename,
	}

	data, err := h.TripRepository.CreateTrip(trip)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return

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

// FIXED: create new value in database,instead updating it
func (h *Triphandler) UpdateTrip(w http.ResponseWriter, r *http.Request) {
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

	// NOTES: lakukan pengecekan isian form dengan if

	// fmt.Println("baris ke 203 trip handler ", request)
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// trip, _ = h.TripRepository.FindSingleTrip(id)

	// fmt.Println(request.Country)
	// return

	trip := models.Trip{}

	// check all field for emptieness
	if request.Title != "" {
		trip.Title = request.Title
	}
	if request.Accomodation != "" {
		trip.Accomodation = request.Accomodation
	}
	if request.Country != 0 {
		trip.Country_id = request.Country
	}
	if request.Transportation != "" {
		trip.Transportation = request.Transportation
	}
	if request.Eat != "" {
		trip.Eat = request.Eat
	}
	if request.Day != 0 {
		trip.Day = request.Day
	}
	if request.Night != 0 {
		trip.Night = request.Night
	}
	if request.DateTrip != "" {
		trip.DateTrip = request.DateTrip
	}
	if request.Price != 0 {
		trip.Price = request.Price
	}
	if request.Quota != 0 {
		trip.Quota = request.Quota
	}
	if request.Description != "" {
		trip.Description = request.Description
	}
	if request.Image != "" {
		trip.Image = request.Image
	}

	// fmt.Println(request)

	data, err := h.TripRepository.UpdateTrip(trip, id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())

	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)

}

func (h *Triphandler) DeleteTrip(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	trip, err := h.TripRepository.FindSingleTrip(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.TripRepository.DeleteTrip(trip, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)

}
