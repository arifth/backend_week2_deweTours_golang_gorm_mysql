package handlers

import (
	"encoding/json"
	"fmt"
	authdto "gorm-imp/dto/auth"
	dto "gorm-imp/dto/result"
	"gorm-imp/models"
	"gorm-imp/pkg/bcrypt"
	jwtToken "gorm-imp/pkg/jwt"
	"gorm-imp/repositories"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}

}

func (h *handlerAuth) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// fmt.Println(r.Body)
	request := new(authdto.RegisterRequest)

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()

	err = validation.Struct(request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	pass, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	//   put name, email and hashed pass to model user
	user := models.User{
		FullName: request.FullName,
		Email:    request.Email,
		Password: pass,
		Phone:    request.Phone,
		Address:  request.Address,
	}

	fmt.Println(user)

	data, err := h.AuthRepository.Register(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerAuth) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// fmt.Println(r.Body)
	request := new(authdto.LoginRequest)

	//  write request body into variable request container
	err := json.NewDecoder(r.Body).Decode(&request)

	// NOTE: dont know why,but then we put request into a new variable somehow,which assigning value from request into user with datatype models.User

	user := models.User{
		Email:    request.Email,
		Password: request.Password,
	}

	// NOTE: in this code block , we do email and password validation

	// email validation

	// fetch username and pass from db
	user, err = h.AuthRepository.Login(user.Email)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// password validation
	// compare both hashed pass from db and hashed pass retrieved from r.password

	isValidPass := bcrypt.CheckPasswordHash(request.Password, user.Password)

	if !isValidPass {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "wrong email or password"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// generate jwt token

	claims := jwt.MapClaims{}
	claims["id"] = user.ID

	// set jwt expiration time to 2 hours
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	// check isi dari claims jwt token
	fmt.Println(claims)

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		fmt.Println("unauthorized ")

		return

	}

	loginResponse := authdto.LoginRequest{
		FullName: user.FullName,
		Email:    user.Email,
		Password: user.Password,
		Token:    token,
	}

	// after all validation succes , set header response ,add value needed to a new var ,then
	// write data to response and encode it to json format
	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Code: http.StatusOK, Data: loginResponse}
	json.NewEncoder(w).Encode(response)

}
