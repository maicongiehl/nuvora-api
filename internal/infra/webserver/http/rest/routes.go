package rest

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	di "github.com/maicongiehl/nuvora-api/configs/di"
	"github.com/maicongiehl/nuvora-api/internal/infra/webserver/http/rest/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Router(app *di.App) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/docs/nuvora/v1*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/nuvora/v1/doc.json")))
	
	customerHandler := handlers.NewCustomerHandler(app) 
	r.Route("/customer", func (r chi.Router) {
		r.Post("/",  customerHandler.Login)
		r.Get("/last-purchases/{id}", customerHandler.LastPurchases)
	})
	
	travelHandler := handlers.NewTravelHandler(app) 
	r.Route("/travel", func (r chi.Router) {
		r.Get("/avaiables/{id}", travelHandler.CustomerPossibleTravels)
	})

	companyHandler := handlers.NewCompanyHandler(app)
	r.Route("/company", func (r chi.Router) {
		r.Post("/", companyHandler.Login)
	})

	return r
}