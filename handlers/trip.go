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
		CountryId:      request.Country,
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
		CountryId:      request.Country,
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

	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		fmt.Println("param tidak ada ")
		fmt.Println(err)
	}

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
		// Image:          filename,
	}

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
		CountryId:      request.Country,
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

	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		fmt.Println("param tidak ada ")
		fmt.Println(err)
	}

	data, err := h.TripRepository.DeleteTrip(trip, id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())

	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)

}
