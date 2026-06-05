package router

import (
	"TouristAgencyApp/src/handler"
	appmiddleware "TouristAgencyApp/src/middleware"
	"TouristAgencyApp/src/service"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(
	e *echo.Echo,
	healthHandler *handler.HealthHandler,
	authHandler *handler.AuthHandler,
	touristCategoryHandler *handler.TouristCategoryHandler,
	aclService *service.ACLService,
	cargoTypeHandler *handler.CargoTypeHandler,
	flightTypeHandler *handler.FlightTypeHandler,
	excursionAgencyHandler *handler.ExcursionAgencyHandler,
	excursionHandler *handler.ExcursionHandler,
	hotelHandler *handler.HotelHandler,
	hotelRoomHandler *handler.HotelRoomHandler,
	financialCategoryHandler *handler.FinancialCategoryHandler,
	financialOperationHandler *handler.FinancialOperationHandler,
	touristHandler *handler.TouristHandler,
	identityDocumentHandler *handler.IdentityDocumentHandler,
	touristGroupHandler *handler.TouristGroupHandler,
	groupMemberHandler *handler.GroupMemberHandler,
	visaHandler *handler.VisaHandler,
	accommodationHandler *handler.AccommodationHandler,
	excursionScheduleHandler *handler.ExcursionScheduleHandler,
	excursionBookingHandler *handler.ExcursionBookingHandler,
	flightHandler *handler.FlightHandler,
	cargoStatementHandler *handler.CargoStatementHandler,
	cargoItemHandler *handler.CargoItemHandler,
	cargoShipmentHandler *handler.CargoShipmentHandler,
	childCompanionHandler *handler.ChildCompanionHandler,
	userHandler *handler.UserHandler,
	roleHandler *handler.RoleHandler,
	permissionHandler *handler.PermissionHandler,
	rolePermissionHandler *handler.RolePermissionHandler,
	userRoleHandler *handler.UserRoleHandler,
	reportHandler *handler.ReportHandler,
	meHandler *handler.MeHandler,
	jwtSecret string,
) {
	e.GET("/health", healthHandler.Health)
	e.GET("/health/db", healthHandler.DBHealth)

	auth := e.Group("/auth")
	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)

	api := e.Group("/api/v1")
	api.Use(appmiddleware.JWTAuth(jwtSecret))

	api.GET("/me", authHandler.Me)

	touristCategories := api.Group("/tourist-categories")

	touristCategories.GET(
		"",
		touristCategoryHandler.List,
		appmiddleware.RequirePermission(aclService, "tourist_categories.read"),
	)

	touristCategories.GET(
		"/:id",
		touristCategoryHandler.GetByID,
		appmiddleware.RequirePermission(aclService, "tourist_categories.read"),
	)

	touristCategories.POST(
		"",
		touristCategoryHandler.Create,
		appmiddleware.RequirePermission(aclService, "tourist_categories.create"),
	)

	touristCategories.PUT(
		"/:id",
		touristCategoryHandler.Update,
		appmiddleware.RequirePermission(aclService, "tourist_categories.update"),
	)

	touristCategories.DELETE(
		"/:id",
		touristCategoryHandler.Delete,
		appmiddleware.RequirePermission(aclService, "tourist_categories.delete"),
	)

	cargoTypes := api.Group("/cargo-types")

	cargoTypes.GET(
		"",
		cargoTypeHandler.List,
		appmiddleware.RequirePermission(aclService, "cargo_types.read"),
	)

	cargoTypes.GET(
		"/:id",
		cargoTypeHandler.GetByID,
		appmiddleware.RequirePermission(aclService, "cargo_types.read"),
	)

	cargoTypes.POST(
		"",
		cargoTypeHandler.Create,
		appmiddleware.RequirePermission(aclService, "cargo_types.create"),
	)

	cargoTypes.PUT(
		"/:id",
		cargoTypeHandler.Update,
		appmiddleware.RequirePermission(aclService, "cargo_types.update"),
	)

	cargoTypes.DELETE(
		"/:id",
		cargoTypeHandler.Delete,
		appmiddleware.RequirePermission(aclService, "cargo_types.delete"),
	)

	flightTypes := api.Group("/flight-types")

	flightTypes.GET(
		"",
		flightTypeHandler.List,
		appmiddleware.RequirePermission(aclService, "flight_types.read"),
	)

	flightTypes.GET(
		"/:id",
		flightTypeHandler.GetByID,
		appmiddleware.RequirePermission(aclService, "flight_types.read"),
	)

	flightTypes.POST(
		"",
		flightTypeHandler.Create,
		appmiddleware.RequirePermission(aclService, "flight_types.create"),
	)

	flightTypes.PUT(
		"/:id",
		flightTypeHandler.Update,
		appmiddleware.RequirePermission(aclService, "flight_types.update"),
	)

	flightTypes.DELETE(
		"/:id",
		flightTypeHandler.Delete,
		appmiddleware.RequirePermission(aclService, "flight_types.delete"),
	)

	excursionAgencies := api.Group("/excursion-agencies")

	excursionAgencies.GET(
		"",
		excursionAgencyHandler.List,
		appmiddleware.RequirePermission(aclService, "excursion_agencies.read"),
	)

	excursionAgencies.GET(
		"/:id",
		excursionAgencyHandler.GetByID,
		appmiddleware.RequirePermission(aclService, "excursion_agencies.read"),
	)

	excursionAgencies.POST(
		"",
		excursionAgencyHandler.Create,
		appmiddleware.RequirePermission(aclService, "excursion_agencies.create"),
	)

	excursionAgencies.PUT(
		"/:id",
		excursionAgencyHandler.Update,
		appmiddleware.RequirePermission(aclService, "excursion_agencies.update"),
	)

	excursionAgencies.DELETE(
		"/:id",
		excursionAgencyHandler.Delete,
		appmiddleware.RequirePermission(aclService, "excursion_agencies.delete"),
	)

	excursions := api.Group("/excursions")

	excursions.GET(
		"",
		excursionHandler.List,
		appmiddleware.RequirePermission(aclService, "excursions.read"),
	)

	excursions.GET(
		"/:id",
		excursionHandler.GetByID,
		appmiddleware.RequirePermission(aclService, "excursions.read"),
	)

	excursions.POST(
		"",
		excursionHandler.Create,
		appmiddleware.RequirePermission(aclService, "excursions.create"),
	)

	excursions.PUT(
		"/:id",
		excursionHandler.Update,
		appmiddleware.RequirePermission(aclService, "excursions.update"),
	)

	excursions.DELETE(
		"/:id",
		excursionHandler.Delete,
		appmiddleware.RequirePermission(aclService, "excursions.delete"),
	)

	hotels := api.Group("/hotels")

	hotels.GET(
		"",
		hotelHandler.List,
		appmiddleware.RequirePermission(aclService, "hotels.read"),
	)

	hotels.GET(
		"/:id",
		hotelHandler.GetByID,
		appmiddleware.RequirePermission(aclService, "hotels.read"),
	)

	hotels.POST(
		"",
		hotelHandler.Create,
		appmiddleware.RequirePermission(aclService, "hotels.create"),
	)

	hotels.PUT(
		"/:id",
		hotelHandler.Update,
		appmiddleware.RequirePermission(aclService, "hotels.update"),
	)

	hotels.DELETE(
		"/:id",
		hotelHandler.Delete,
		appmiddleware.RequirePermission(aclService, "hotels.delete"),
	)

	hotelRooms := api.Group("/hotel-rooms")

	hotelRooms.GET(
		"",
		hotelRoomHandler.List,
		appmiddleware.RequirePermission(aclService, "hotel_rooms.read"),
	)

	hotelRooms.GET(
		"/:id",
		hotelRoomHandler.GetByID,
		appmiddleware.RequirePermission(aclService, "hotel_rooms.read"),
	)

	hotelRooms.POST(
		"",
		hotelRoomHandler.Create,
		appmiddleware.RequirePermission(aclService, "hotel_rooms.create"),
	)

	hotelRooms.PUT(
		"/:id",
		hotelRoomHandler.Update,
		appmiddleware.RequirePermission(aclService, "hotel_rooms.update"),
	)

	hotelRooms.DELETE(
		"/:id",
		hotelRoomHandler.Delete,
		appmiddleware.RequirePermission(aclService, "hotel_rooms.delete"),
	)

	financialCategories := api.Group("/financial-categories")

	financialCategories.GET(
		"",
		financialCategoryHandler.List,
		appmiddleware.RequirePermission(aclService, "financial_categories.read"),
	)

	financialCategories.GET(
		"/:id",
		financialCategoryHandler.GetByID,
		appmiddleware.RequirePermission(aclService, "financial_categories.read"),
	)

	financialCategories.POST(
		"",
		financialCategoryHandler.Create,
		appmiddleware.RequirePermission(aclService, "financial_categories.create"),
	)

	financialCategories.PUT(
		"/:id",
		financialCategoryHandler.Update,
		appmiddleware.RequirePermission(aclService, "financial_categories.update"),
	)

	financialCategories.DELETE(
		"/:id",
		financialCategoryHandler.Delete,
		appmiddleware.RequirePermission(aclService, "financial_categories.delete"),
	)

	financialOperations := api.Group("/financial-operations")

	financialOperations.GET(
		"",
		financialOperationHandler.List,
		appmiddleware.RequirePermission(aclService, "financial_operations.read"),
	)

	financialOperations.GET(
		"/:id",
		financialOperationHandler.GetByID,
		appmiddleware.RequirePermission(aclService, "financial_operations.read"),
	)

	financialOperations.POST(
		"",
		financialOperationHandler.Create,
		appmiddleware.RequirePermission(aclService, "financial_operations.create"),
	)

	financialOperations.PUT(
		"/:id",
		financialOperationHandler.Update,
		appmiddleware.RequirePermission(aclService, "financial_operations.update"),
	)

	financialOperations.DELETE(
		"/:id",
		financialOperationHandler.Delete,
		appmiddleware.RequirePermission(aclService, "financial_operations.delete"),
	)

	tourists := api.Group("/tourists")

	tourists.GET(
		"",
		touristHandler.List,
		appmiddleware.RequirePermission(aclService, "tourists.read"),
	)

	tourists.GET(
		"/:id",
		touristHandler.GetByID,
		appmiddleware.RequirePermission(aclService, "tourists.read"),
	)

	tourists.POST(
		"",
		touristHandler.Create,
		appmiddleware.RequirePermission(aclService, "tourists.create"),
	)

	tourists.PUT(
		"/:id",
		touristHandler.Update,
		appmiddleware.RequirePermission(aclService, "tourists.update"),
	)

	tourists.DELETE(
		"/:id",
		touristHandler.Delete,
		appmiddleware.RequirePermission(aclService, "tourists.delete"),
	)

	identityDocuments := api.Group("/identity-documents")

	identityDocuments.GET(
		"",
		identityDocumentHandler.List,
		appmiddleware.RequirePermission(aclService, "identity_documents.read"),
	)

	identityDocuments.GET(
		"/:id",
		identityDocumentHandler.GetByID,
		appmiddleware.RequirePermission(aclService, "identity_documents.read"),
	)

	identityDocuments.POST(
		"",
		identityDocumentHandler.Create,
		appmiddleware.RequirePermission(aclService, "identity_documents.create"),
	)

	identityDocuments.PUT(
		"/:id",
		identityDocumentHandler.Update,
		appmiddleware.RequirePermission(aclService, "identity_documents.update"),
	)

	identityDocuments.DELETE(
		"/:id",
		identityDocumentHandler.Delete,
		appmiddleware.RequirePermission(aclService, "identity_documents.delete"),
	)

	touristGroups := api.Group("/tourist-groups")

	touristGroups.GET(
		"",
		touristGroupHandler.List,
		appmiddleware.RequirePermission(aclService, "tourist_groups.read"),
	)

	touristGroups.GET(
		"/:id",
		touristGroupHandler.GetByID,
		appmiddleware.RequirePermission(aclService, "tourist_groups.read"),
	)

	touristGroups.POST(
		"",
		touristGroupHandler.Create,
		appmiddleware.RequirePermission(aclService, "tourist_groups.create"),
	)

	touristGroups.PUT(
		"/:id",
		touristGroupHandler.Update,
		appmiddleware.RequirePermission(aclService, "tourist_groups.update"),
	)

	touristGroups.DELETE(
		"/:id",
		touristGroupHandler.Delete,
		appmiddleware.RequirePermission(aclService, "tourist_groups.delete"),
	)

	groupMembers := api.Group("/group-members")

	groupMembers.GET(
		"",
		groupMemberHandler.List,
		appmiddleware.RequirePermission(aclService, "group_members.read"),
	)

	groupMembers.GET(
		"/:id",
		groupMemberHandler.GetByID,
		appmiddleware.RequirePermission(aclService, "group_members.read"),
	)

	groupMembers.POST(
		"",
		groupMemberHandler.Create,
		appmiddleware.RequirePermission(aclService, "group_members.create"),
	)

	groupMembers.PUT(
		"/:id",
		groupMemberHandler.Update,
		appmiddleware.RequirePermission(aclService, "group_members.update"),
	)

	groupMembers.DELETE(
		"/:id",
		groupMemberHandler.Delete,
		appmiddleware.RequirePermission(aclService, "group_members.delete"),
	)

	visas := api.Group("/visas")

	visas.GET(
		"",
		visaHandler.List,
		appmiddleware.RequirePermission(aclService, "visas.read"),
	)

	visas.GET(
		"/:id",
		visaHandler.GetByID,
		appmiddleware.RequirePermission(aclService, "visas.read"),
	)

	visas.POST(
		"",
		visaHandler.Create,
		appmiddleware.RequirePermission(aclService, "visas.create"),
	)

	visas.PUT(
		"/:id",
		visaHandler.Update,
		appmiddleware.RequirePermission(aclService, "visas.update"),
	)

	visas.DELETE(
		"/:id",
		visaHandler.Delete,
		appmiddleware.RequirePermission(aclService, "visas.delete"),
	)

	accommodations := api.Group("/accommodations")

	accommodations.GET(
		"",
		accommodationHandler.List,
		appmiddleware.RequirePermission(aclService, "accommodations.read"),
	)

	accommodations.GET(
		"/:id",
		accommodationHandler.GetByID,
		appmiddleware.RequirePermission(aclService, "accommodations.read"),
	)

	accommodations.POST(
		"",
		accommodationHandler.Create,
		appmiddleware.RequirePermission(aclService, "accommodations.create"),
	)

	accommodations.PUT(
		"/:id",
		accommodationHandler.Update,
		appmiddleware.RequirePermission(aclService, "accommodations.update"),
	)

	accommodations.DELETE(
		"/:id",
		accommodationHandler.Delete,
		appmiddleware.RequirePermission(aclService, "accommodations.delete"),
	)

	excursionSchedules := api.Group("/excursion-schedules")

	excursionSchedules.GET(
		"",
		excursionScheduleHandler.List,
		appmiddleware.RequirePermission(aclService, "excursion_schedule.read"),
	)

	excursionSchedules.GET(
		"/:id",
		excursionScheduleHandler.GetByID,
		appmiddleware.RequirePermission(aclService, "excursion_schedule.read"),
	)

	excursionSchedules.POST(
		"",
		excursionScheduleHandler.Create,
		appmiddleware.RequirePermission(aclService, "excursion_schedule.create"),
	)

	excursionSchedules.PUT(
		"/:id",
		excursionScheduleHandler.Update,
		appmiddleware.RequirePermission(aclService, "excursion_schedule.update"),
	)

	excursionSchedules.DELETE(
		"/:id",
		excursionScheduleHandler.Delete,
		appmiddleware.RequirePermission(aclService, "excursion_schedule.delete"),
	)

	excursionBookings := api.Group("/excursion-bookings")

	excursionBookings.GET(
		"",
		excursionBookingHandler.List,
		appmiddleware.RequirePermission(aclService, "excursion_bookings.read"),
	)

	excursionBookings.GET(
		"/:id",
		excursionBookingHandler.GetByID,
		appmiddleware.RequirePermission(aclService, "excursion_bookings.read"),
	)

	excursionBookings.POST(
		"",
		excursionBookingHandler.Create,
		appmiddleware.RequirePermission(aclService, "excursion_bookings.create"),
	)

	excursionBookings.PUT(
		"/:id",
		excursionBookingHandler.Update,
		appmiddleware.RequirePermission(aclService, "excursion_bookings.update"),
	)

	excursionBookings.DELETE(
		"/:id",
		excursionBookingHandler.Delete,
		appmiddleware.RequirePermission(aclService, "excursion_bookings.delete"),
	)

	flights := api.Group("/flights")

	flights.GET(
		"",
		flightHandler.List,
		appmiddleware.RequirePermission(aclService, "flights.read"),
	)

	flights.GET(
		"/:id",
		flightHandler.GetByID,
		appmiddleware.RequirePermission(aclService, "flights.read"),
	)

	flights.POST(
		"",
		flightHandler.Create,
		appmiddleware.RequirePermission(aclService, "flights.create"),
	)

	flights.PUT(
		"/:id",
		flightHandler.Update,
		appmiddleware.RequirePermission(aclService, "flights.update"),
	)

	flights.DELETE(
		"/:id",
		flightHandler.Delete,
		appmiddleware.RequirePermission(aclService, "flights.delete"),
	)

	cargoStatements := api.Group("/cargo-statements")

	cargoStatements.GET(
		"",
		cargoStatementHandler.List,
		appmiddleware.RequirePermission(aclService, "cargo_statements.read"),
	)

	cargoStatements.GET(
		"/:id",
		cargoStatementHandler.GetByID,
		appmiddleware.RequirePermission(aclService, "cargo_statements.read"),
	)

	cargoStatements.POST(
		"",
		cargoStatementHandler.Create,
		appmiddleware.RequirePermission(aclService, "cargo_statements.create"),
	)

	cargoStatements.PUT(
		"/:id",
		cargoStatementHandler.Update,
		appmiddleware.RequirePermission(aclService, "cargo_statements.update"),
	)

	cargoStatements.DELETE(
		"/:id",
		cargoStatementHandler.Delete,
		appmiddleware.RequirePermission(aclService, "cargo_statements.delete"),
	)

	cargoItems := api.Group("/cargo-items")

	cargoItems.GET(
		"",
		cargoItemHandler.List,
		appmiddleware.RequirePermission(aclService, "cargo_items.read"),
	)

	cargoItems.GET(
		"/:id",
		cargoItemHandler.GetByID,
		appmiddleware.RequirePermission(aclService, "cargo_items.read"),
	)

	cargoItems.POST(
		"",
		cargoItemHandler.Create,
		appmiddleware.RequirePermission(aclService, "cargo_items.create"),
	)

	cargoItems.PUT(
		"/:id",
		cargoItemHandler.Update,
		appmiddleware.RequirePermission(aclService, "cargo_items.update"),
	)

	cargoItems.DELETE(
		"/:id",
		cargoItemHandler.Delete,
		appmiddleware.RequirePermission(aclService, "cargo_items.delete"),
	)

	cargoShipments := api.Group("/cargo-shipments")

	cargoShipments.GET(
		"",
		cargoShipmentHandler.List,
		appmiddleware.RequirePermission(aclService, "cargo_shipments.read"),
	)

	cargoShipments.GET(
		"/:id",
		cargoShipmentHandler.GetByID,
		appmiddleware.RequirePermission(aclService, "cargo_shipments.read"),
	)

	cargoShipments.POST(
		"",
		cargoShipmentHandler.Create,
		appmiddleware.RequirePermission(aclService, "cargo_shipments.create"),
	)

	cargoShipments.PUT(
		"/:id",
		cargoShipmentHandler.Update,
		appmiddleware.RequirePermission(aclService, "cargo_shipments.update"),
	)

	cargoShipments.DELETE(
		"/:id",
		cargoShipmentHandler.Delete,
		appmiddleware.RequirePermission(aclService, "cargo_shipments.delete"),
	)

	childCompanions := api.Group("/child-companions")

	childCompanions.GET(
		"",
		childCompanionHandler.List,
		appmiddleware.RequirePermission(aclService, "child_companions.read"),
	)

	childCompanions.GET(
		"/:parent_group_member_id/:child_group_member_id",
		childCompanionHandler.GetByIDs,
		appmiddleware.RequirePermission(aclService, "child_companions.read"),
	)

	childCompanions.POST(
		"",
		childCompanionHandler.Create,
		appmiddleware.RequirePermission(aclService, "child_companions.create"),
	)

	childCompanions.PUT(
		"/:parent_group_member_id/:child_group_member_id",
		childCompanionHandler.Update,
		appmiddleware.RequirePermission(aclService, "child_companions.update"),
	)

	childCompanions.DELETE(
		"/:parent_group_member_id/:child_group_member_id",
		childCompanionHandler.Delete,
		appmiddleware.RequirePermission(aclService, "child_companions.delete"),
	)

	users := api.Group("/users")

	users.GET(
		"",
		userHandler.List,
		appmiddleware.RequirePermission(aclService, "users.read"),
	)

	users.GET(
		"/:id",
		userHandler.GetByID,
		appmiddleware.RequirePermission(aclService, "users.read"),
	)

	users.POST(
		"",
		userHandler.Create,
		appmiddleware.RequirePermission(aclService, "users.create"),
	)

	users.PUT(
		"/:id",
		userHandler.Update,
		appmiddleware.RequirePermission(aclService, "users.update"),
	)

	users.DELETE(
		"/:id",
		userHandler.Delete,
		appmiddleware.RequirePermission(aclService, "users.delete"),
	)

	roles := api.Group("/roles")

	roles.GET(
		"",
		roleHandler.FindAll,
		appmiddleware.RequirePermission(aclService, "roles.read"),
	)

	roles.GET(
		"/:id",
		roleHandler.FindByID,
		appmiddleware.RequirePermission(aclService, "roles.read"),
	)

	roles.POST(
		"",
		roleHandler.Create,
		appmiddleware.RequirePermission(aclService, "roles.create"),
	)

	roles.PUT(
		"/:id",
		roleHandler.Update,
		appmiddleware.RequirePermission(aclService, "roles.update"),
	)

	roles.DELETE(
		"/:id",
		roleHandler.Delete,
		appmiddleware.RequirePermission(aclService, "roles.delete"),
	)

	permissions := api.Group("/permissions")

	permissions.GET(
		"",
		permissionHandler.FindAll,
		appmiddleware.RequirePermission(aclService, "permissions.read"),
	)

	permissions.GET(
		"/:id",
		permissionHandler.FindByID,
		appmiddleware.RequirePermission(aclService, "permissions.read"),
	)

	permissions.POST(
		"",
		permissionHandler.Create,
		appmiddleware.RequirePermission(aclService, "permissions.create"),
	)

	permissions.PUT(
		"/:id",
		permissionHandler.Update,
		appmiddleware.RequirePermission(aclService, "permissions.update"),
	)

	permissions.DELETE(
		"/:id",
		permissionHandler.Delete,
		appmiddleware.RequirePermission(aclService, "permissions.delete"),
	)

	rolePermissions := api.Group("/role-permissions")

	rolePermissions.GET(
		"",
		rolePermissionHandler.FindAll,
		appmiddleware.RequirePermission(aclService, "role_permissions.read"),
	)

	rolePermissions.GET(
		"/:role_id/:permission_id",
		rolePermissionHandler.FindByIDs,
		appmiddleware.RequirePermission(aclService, "role_permissions.read"),
	)

	rolePermissions.POST(
		"",
		rolePermissionHandler.Create,
		appmiddleware.RequirePermission(aclService, "role_permissions.create"),
	)

	rolePermissions.PUT(
		"/:role_id/:permission_id",
		rolePermissionHandler.Update,
		appmiddleware.RequirePermission(aclService, "role_permissions.update"),
	)

	rolePermissions.DELETE(
		"/:role_id/:permission_id",
		rolePermissionHandler.Delete,
		appmiddleware.RequirePermission(aclService, "role_permissions.delete"),
	)

	userRoles := api.Group("/user-roles")

	userRoles.GET(
		"",
		userRoleHandler.FindAll,
		appmiddleware.RequirePermission(aclService, "user_roles.read"),
	)

	userRoles.GET(
		"/by-user/:user_id",
		userRoleHandler.FindByUserID,
		appmiddleware.RequirePermission(aclService, "user_roles.read"),
	)

	userRoles.GET(
		"/:user_id/:role_id",
		userRoleHandler.FindByIDs,
		appmiddleware.RequirePermission(aclService, "user_roles.read"),
	)

	userRoles.POST(
		"",
		userRoleHandler.Create,
		appmiddleware.RequirePermission(aclService, "user_roles.create"),
	)

	userRoles.PUT(
		"/:user_id/:role_id",
		userRoleHandler.Update,
		appmiddleware.RequirePermission(aclService, "user_roles.update"),
	)

	userRoles.DELETE(
		"/:user_id/:role_id",
		userRoleHandler.Delete,
		appmiddleware.RequirePermission(aclService, "user_roles.delete"),
	)

	reports := api.Group("/reports")

	reports.GET(
		"/customs-tourists",
		reportHandler.CustomsTourists,
		appmiddleware.RequirePermission(aclService, "reports.customs_list.read"),
	)

	reports.GET(
		"/accommodation-list",
		reportHandler.AccommodationList,
		appmiddleware.RequirePermission(aclService, "reports.accommodation.read"),
	)

	reports.GET(
		"/tourist-count",
		reportHandler.TouristCount,
		appmiddleware.RequirePermission(aclService, "reports.tourists_count.read"),
	)

	reports.GET(
		"/tourist-info/:tourist_id",
		reportHandler.TouristInfo,
		appmiddleware.RequirePermission(aclService, "reports.tourist_info.read"),
	)

	reports.GET(
		"/hotel-occupancy",
		reportHandler.HotelOccupancy,
		appmiddleware.RequirePermission(aclService, "reports.hotels.read"),
	)

	reports.GET(
		"/excursion-tourist-count",
		reportHandler.ExcursionTouristCount,
		appmiddleware.RequirePermission(aclService, "reports.excursions.read"),
	)

	reports.GET(
		"/excursion-analytics",
		reportHandler.ExcursionAnalytics,
		appmiddleware.RequirePermission(aclService, "reports.excursions.read"),
	)

	reports.GET(
		"/flight-load",
		reportHandler.FlightLoad,
		appmiddleware.RequirePermission(aclService, "reports.flight_load.read"),
	)

	reports.GET(
		"/warehouse-turnover",
		reportHandler.WarehouseTurnover,
		appmiddleware.RequirePermission(aclService, "reports.cargo_turnover.read"),
	)

	reports.GET(
		"/group-financial-report",
		reportHandler.GroupFinancialReport,
		appmiddleware.RequirePermission(aclService, "reports.financial.read"),
	)

	reports.GET(
		"/income-expense",
		reportHandler.IncomeExpense,
		appmiddleware.RequirePermission(aclService, "reports.financial.read"),
	)

	reports.GET(
		"/cargo-type-share",
		reportHandler.CargoTypeShare,
		appmiddleware.RequirePermission(aclService, "reports.cargo_turnover.read"),
	)

	reports.GET(
		"/profitability",
		reportHandler.Profitability,
		appmiddleware.RequirePermission(aclService, "reports.profitability.read"),
	)

	reports.GET(
		"/tourist-category-ratio",
		reportHandler.TouristCategoryRatio,
		appmiddleware.RequirePermission(aclService, "reports.tourist_categories_percent.read"),
	)

	reports.GET(
		"/flight-tourists",
		reportHandler.FlightTourists,
		appmiddleware.RequirePermission(aclService, "reports.flight_load.read"),
	)

	me := api.Group("/me")

	me.GET(
		"",
		authHandler.Me,
		appmiddleware.RequirePermission(aclService, "profile.read"),
	)

	me.GET(
		"/tours",
		meHandler.Tours,
		appmiddleware.RequirePermission(aclService, "own_tours.read"),
	)

	me.GET(
		"/visas",
		meHandler.Visas,
		appmiddleware.RequirePermission(aclService, "own_visas.read"),
	)

	me.GET(
		"/accommodations",
		meHandler.Accommodations,
		appmiddleware.RequirePermission(aclService, "own_accommodations.read"),
	)

	me.GET(
		"/excursions",
		meHandler.Excursions,
		appmiddleware.RequirePermission(aclService, "own_excursions.read"),
	)

	me.GET(
		"/cargo",
		meHandler.Cargo,
		appmiddleware.RequirePermission(aclService, "own_cargo.read"),
	)

	me.GET(
		"/identity-document",
		meHandler.IdentityDocument,
		appmiddleware.RequirePermission(aclService, "profile.read"),
	)

	me.POST(
		"/identity-document",
		meHandler.CreateIdentityDocument,
		appmiddleware.RequirePermission(aclService, "profile.update"),
	)

	me.PUT(
		"/identity-document",
		meHandler.UpdateIdentityDocument,
		appmiddleware.RequirePermission(aclService, "profile.update"),
	)

	me.POST(
		"/excursion-bookings",
		meHandler.CreateExcursionBooking,
		appmiddleware.RequirePermission(aclService, "own_excursions.create"),
	)
}
