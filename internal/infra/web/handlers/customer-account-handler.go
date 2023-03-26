package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	dto "github.com/MaiconGiehl/API/internal/dto"
	"github.com/MaiconGiehl/API/internal/infra/database"
	"github.com/MaiconGiehl/API/internal/usecase"
	"github.com/go-chi/chi/v5"
)

type CustomerAccountHandler struct {
	Ctx 										context.Context
	CustomerRepository 			*database.CustomerRepository
	PersonRepository 				*database.PersonRepository
	AccountRepository 			*database.AccountRepository
}

func NewCustomerAccountHandler(
	ctx 										context.Context,
	customerRepository 			*database.CustomerRepository,
	personRepository 				*database.PersonRepository,
	accountRepository 			*database.AccountRepository,
	) *CustomerAccountHandler {
	return &CustomerAccountHandler{
		Ctx: 									ctx,
		CustomerRepository: 	customerRepository,
		PersonRepository: 		personRepository,
		AccountRepository: 		accountRepository,
	}
}

// CreateCustomer godoc
// @Summary      			Add customer
// @Description  			Create new customer
// @Tags         			Customer
// @Accept       			json
// @Produce      			json
// @Param        			request   				body      dto.CustomerAccountInputDTO  true  "Customer Info"
// @Success      			200  											{object}   object
// @Failure      			404
// @Router       			/customer [post]
func (h *CustomerAccountHandler) CreateCustomerAccount(w http.ResponseWriter, r *http.Request) {
	input, err := h.getCreateInput(w, r)
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	usecase := usecase.NewCreateCustomerAccountUseCase(*h.CustomerRepository, *h.PersonRepository, *h.AccountRepository) 
	err = usecase.Execute(input)
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("New customer created")
}

// Login godoc
// @Summary      			Delete a specific bus
// @Description  			Get a bus
// @Tags         			Customer
// @Accept       			json
// @Produce      			json
// @Param        			email   										path      		string  true  "Account ID"
// @Param        			password   									path      		string  true  "Account ID"
// @Success      			202  										{object}   		object
// @Failure      			404
// @Router       			/customer/{email}/{password} [get]
func (h *CustomerAccountHandler) GetCustomerAccount(w http.ResponseWriter, r *http.Request) {
	input, err := h.getLoginInfo(w, r)
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	usecase := usecase.NewGetCustomerAccountUseCase(*h.AccountRepository) 
	output, err := usecase.Execute(input)
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}


func (h *CustomerAccountHandler) getCreateInput(w http.ResponseWriter, r *http.Request) (*dto.CustomerAccountInputDTO, error) {
	var customerAccount dto.CustomerAccountInputDTO
	err := json.NewDecoder(r.Body).Decode(&customerAccount)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return &customerAccount, err
	}
	return &customerAccount, nil
}

func (h *CustomerAccountHandler) getLoginInfo(w http.ResponseWriter, r *http.Request) (*dto.LoginCustomerInputDTO, error) {
	return &dto.LoginCustomerInputDTO{
		Email: chi.URLParam(r, "email"),
		Password: chi.URLParam(r, "password"),
	}, nil
}