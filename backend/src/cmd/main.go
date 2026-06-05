package main

import (
	"TouristAgencyApp/src/config"
	"TouristAgencyApp/src/database"
	"TouristAgencyApp/src/handler"
	"TouristAgencyApp/src/repository"
	"TouristAgencyApp/src/routes"
	"TouristAgencyApp/src/service"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	_ = godotenv.Load()

	cfg := config.Load()

	db, err := database.NewPostgresConnection(cfg)
	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOriginFunc: func(origin string) (bool, error) {
			return true, nil
		},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			"X-Requested-With",
		},
		ExposeHeaders: []string{
			echo.HeaderAuthorization,
		},
		AllowCredentials: false,
		MaxAge:           86400,
	}))

	authRepository := repository.NewAuthRepository(db)
	aclRepository := repository.NewACLRepository(db)
	touristCategoryRepository := repository.NewTouristCategoryRepository(db)
	cargoTypeRepository := repository.NewCargoTypeRepository(db)
	flightTypeRepository := repository.NewFlightTypeRepository(db)
	excursionAgencyRepository := repository.NewExcursionAgencyRepository(db)
	excursionRepository := repository.NewExcursionRepository(db)
	hotelRepository := repository.NewHotelRepository(db)
	hotelRoomRepository := repository.NewHotelRoomRepository(db)
	financialCategoryRepository := repository.NewFinancialCategoryRepository(db)
	financialOperationsRepository := repository.NewFinancialOperationRepository(db)
	touristsRepository := repository.NewTouristRepository(db)
	identityDocumentsRepository := repository.NewIdentityDocumentRepository(db)
	touristGroupRepository := repository.NewTouristGroupRepository(db)
	groupMembersRepository := repository.NewGroupMemberRepository(db)
	visaRepository := repository.NewVisaRepository(db)
	accommodationRepository := repository.NewAccommodationRepository(db)
	excursionScheduleRepository := repository.NewExcursionScheduleRepository(db)
	excursionBookingRepository := repository.NewExcursionBookingRepository(db)
	flightRepository := repository.NewFlightRepository(db)
	cargoStatementRepository := repository.NewCargoStatementRepository(db)
	cargoItemRepository := repository.NewCargoItemRepository(db)
	cargoShipmentRepository := repository.NewCargoShipmentRepository(db)
	childCompanionRepository := repository.NewChildCompanionRepository(db)
	userRepository := repository.NewUserRepository(db)
	roleRepository := repository.NewRoleRepository(db)
	permissionRepository := repository.NewPermissionRepository(db)
	rolePermissionRepository := repository.NewRolePermissionRepository(db)
	userRoleRepository := repository.NewUserRoleRepository(db)
	reportRepository := repository.NewReportRepository(db)
	meRepository := repository.NewMeRepository(db)

	authService := service.NewAuthService(authRepository, cfg.JWTSecret, cfg.JWTAccessTTLMinutes)
	aclService := service.NewACLService(aclRepository)
	touristCategoryService := service.NewTouristCategoryService(touristCategoryRepository)
	cargoTypeService := service.NewCargoTypeService(cargoTypeRepository)
	flightTypeService := service.NewFlightTypeService(flightTypeRepository)
	excursionAgencyService := service.NewExcursionAgencyService(excursionAgencyRepository)
	excursionService := service.NewExcursionService(excursionRepository)
	hotelService := service.NewHotelService(hotelRepository)
	hotelRoomService := service.NewHotelRoomService(hotelRoomRepository)
	financialCategoryService := service.NewFinancialCategoryService(financialCategoryRepository)
	financialOperationsService := service.NewFinancialOperationService(financialOperationsRepository)
	touristsService := service.NewTouristService(touristsRepository)
	identityDocumentsService := service.NewIdentityDocumentService(identityDocumentsRepository)
	touristGroupService := service.NewTouristGroupService(touristGroupRepository)
	groupMembersService := service.NewGroupMemberService(groupMembersRepository)
	visaService := service.NewVisaService(visaRepository)
	accommodationService := service.NewAccommodationService(accommodationRepository)
	excursionScheduleService := service.NewExcursionScheduleService(excursionScheduleRepository)
	excursionBookingService := service.NewExcursionBookingService(excursionBookingRepository)
	flightService := service.NewFlightService(flightRepository)
	cargoStatementService := service.NewCargoStatementService(cargoStatementRepository)
	cargoItemService := service.NewCargoItemService(cargoItemRepository)
	cargoShipmentService := service.NewCargoShipmentService(cargoShipmentRepository)
	childCompanionService := service.NewChildCompanionService(childCompanionRepository)
	userService := service.NewUserService(userRepository)
	roleService := service.NewRoleService(roleRepository)
	permissionService := service.NewPermissionService(permissionRepository)
	rolePermissionService := service.NewRolePermissionService(rolePermissionRepository)
	userRoleService := service.NewUserRoleService(userRoleRepository)
	reportService := service.NewReportService(reportRepository)
	meService := service.NewMeService(meRepository)

	healthHandler := handler.NewHealthHandler(db)
	authHandler := handler.NewAuthHandler(authService)
	touristCategoryHandler := handler.NewTouristCategoryHandler(touristCategoryService)
	cargoTypeHandler := handler.NewCargoTypeHandler(cargoTypeService)
	flightTypeHandler := handler.NewFlightTypeHandler(flightTypeService)
	excursionAgencyHandler := handler.NewExcursionAgencyHandler(excursionAgencyService)
	excursionHandler := handler.NewExcursionHandler(excursionService)
	hotelHandler := handler.NewHotelHandler(hotelService)
	hotelRoomHandler := handler.NewHotelRoomHandler(hotelRoomService)
	financialCategoryHandler := handler.NewFinancialCategoryHandler(financialCategoryService)
	financialOperationsHandler := handler.NewFinancialOperationHandler(financialOperationsService)
	touristsHandler := handler.NewTouristHandler(touristsService)
	identityDocumentsHandler := handler.NewIdentityDocumentHandler(identityDocumentsService)
	touristGroupHandler := handler.NewTouristGroupHandler(touristGroupService)
	groupMembersHandler := handler.NewGroupMemberHandler(groupMembersService)
	visaHandler := handler.NewVisaHandler(visaService)
	accommodationHandler := handler.NewAccommodationHandler(accommodationService)
	excursionScheduleHandler := handler.NewExcursionScheduleHandler(excursionScheduleService)
	excursionBookingHandler := handler.NewExcursionBookingHandler(excursionBookingService)
	flightHandler := handler.NewFlightHandler(flightService)
	cargoStatementHandler := handler.NewCargoStatementHandler(cargoStatementService)
	cargoItemHandler := handler.NewCargoItemHandler(cargoItemService)
	cargoShipmentHandler := handler.NewCargoShipmentHandler(cargoShipmentService)
	childCompanionHandler := handler.NewChildCompanionHandler(childCompanionService)
	userHandler := handler.NewUserHandler(userService)
	roleHandler := handler.NewRoleHandler(roleService)
	permissionHandler := handler.NewPermissionHandler(permissionService)
	rolePermissionHandler := handler.NewRolePermissionHandler(rolePermissionService)
	userRoleHandler := handler.NewUserRoleHandler(userRoleService)
	reportHandler := handler.NewReportHandler(reportService)
	meHandler := handler.NewMeHandler(meService)

	router.RegisterRoutes(e,
		healthHandler, authHandler, touristCategoryHandler, aclService, cargoTypeHandler,
		flightTypeHandler, excursionAgencyHandler, excursionHandler, hotelHandler, hotelRoomHandler,
		financialCategoryHandler, financialOperationsHandler, touristsHandler, identityDocumentsHandler,
		touristGroupHandler, groupMembersHandler, visaHandler, accommodationHandler, excursionScheduleHandler,
		excursionBookingHandler, flightHandler, cargoStatementHandler, cargoItemHandler, cargoShipmentHandler,
		childCompanionHandler, userHandler, roleHandler, permissionHandler, rolePermissionHandler, userRoleHandler,
		reportHandler, meHandler,
		cfg.JWTSecret)

	e.Logger.Fatal(e.Start(":" + cfg.AppPort))
}
