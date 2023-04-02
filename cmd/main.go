package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"github.com/maicongiehl/nuvera-api/config"
	_ "github.com/maicongiehl/nuvera-api/docs"
	"github.com/maicongiehl/nuvera-api/internal/infra/database"
	"github.com/maicongiehl/nuvera-api/internal/infra/web/handlers"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           TechTur Service
// @version         1.0
// @description     Product API with auhtentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   TechTur
// @contact.url    https://techtur.com.br
// @contact.email  atendimento@techtur.com.br

// @license.name   Acerta Promotora License
// @license.url    https://techtur.com.br

// @host      localhost:8080
// @BasePath  /
func main() {

	dbconfig := config.LoadConfig()
	psqlinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbconfig.DBHost, dbconfig.DBPort, dbconfig.DBUser, dbconfig.DBPassword, dbconfig.DBName)
	
	db, err := sql.Open("postgres", psqlinfo)
	if err != nil {
	  panic(err)
	}
	defer db.Close()

	
	fmt.Print("DB connected\n")
	http.ListenAndServe(":8080", createRouter(db))
}

func createRouter(db *sql.DB) *chi.Mux {
	var ctx context.Context
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	
	busRepository := database.NewBusRepository(db)
	busHandler := handlers.NewBusHandler(ctx, busRepository)
	
	r.Route("/bus", func(r chi.Router) {
		r.Get("/", busHandler.GetAll)
		r.Post("/", busHandler.CreateBus)
		r.Get("/{id}", busHandler.GetBus)
		r.Delete("/{id}", busHandler.DeleteBus)
		r.Patch("/{id}", busHandler.UpdateBus)
	})

	ticketRepository := database.NewTicketRepository(db)
	ticketHandler := handlers.NewTicketHandler(ctx, ticketRepository)

	r.Route("/ticket", func(r chi.Router) {
		r.Post("/", ticketHandler.CreateTicket)
		r.Delete("/{id}", ticketHandler.DeleteTicket)
	})

	personRepository := database.NewPersonRepository(db)
	accountHandler := database.NewAccountRepository(db)

	customerRepository := database.NewCustomerRepository(db)
	
	customerAccountHandler := handlers.NewCustomerAccountHandler(ctx, customerRepository, personRepository, accountHandler)
	
	r.Route("/customer", func(r chi.Router) {
		r.Get("/{email}/{password}", customerAccountHandler.Login)
	})

	companyRepository := database.NewCompanyRepository(db)
	companyHandler := handlers.NewCompanyHandler(ctx, personRepository, companyRepository, accountHandler, customerRepository)

	r.Route("/company", func(r chi.Router) {
		r.Get("/", companyHandler.GetAll)
		r.Get("/{id}/employees", companyHandler.GetEmployees)
		r.Get("/{id}/employees", companyHandler.GetEmployees)
		r.Get("/{id}/employees/tickets", companyHandler.GetEmployeesTickets)
		r.Get("/{email}/{password}", companyHandler.Login)
		r.Post("/{id}/employees", companyHandler.CreateEmployee)
	})

	travelRepository := database.NewTravelRepository(db)
	travelHandler := handlers.NewTravelHandler(ctx, travelRepository)
	r.Route("/travel", func(r chi.Router) {
		r.Post("/", travelHandler.CreateTravel)
		r.Get("/{departure_city_id}/{arrival_city_id}", travelHandler.GetAllPossibleTravesl)
	})
	
	cityRepository := database.NewCityRepository(db)
	cityHandler := handlers.NewCityHandler(ctx, cityRepository)
	r.Route("/city", func(r chi.Router) {
		r.Post("/{name}", cityHandler.CreateCity)
	})

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/doc.json")))
	
	return r
}