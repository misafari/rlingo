package main

func main() {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// dbConf := database.Config{
	// 	Host:           "localhost",
	// 	Port:           5455,
	// 	User:           "admin",
	// 	Password:       "admin",
	// 	Database:       "rlingo",
	// 	MaxConnections: 25,
	// 	MinConnections: 5,
	// }

	// postgresqlDatabaseConnectionPool, err := database.Connect(ctx, dbConf)
	// if err != nil {
	// 	log.Fatalf("Error connecting to database: %v", err)
	// }

	// defer database.Close()

	// queries := generated.New(postgresqlDatabaseConnectionPool)

	// projectRepository := project.NewRepository(queries, postgresqlDatabaseConnectionPool)

	// projectService := project.NewService(projectRepository)

	// app := fiber.New(fiber.Config{
	//ErrorHandler: middleware.ErrorHandler,
	// 	AppName: "Rlingo API v1.0",
	// })

	// app.Use(logger.New(), cors.New(cors.Config{
	// 	AllowOrigins: "*",
	// 	AllowHeaders: "Origin, Content-Type, Accept",
	// 	AllowMethods: "GET, POST, PUT, PATCH, DELETE, OPTIONS",
	// }))

	// api := app.Group("/api/v1")

	// projectHttpHandler := handler.NewProjectHandler(projectCrudUseCase)
	// papi := api.Group("/projects")
	// papi.Get("/", projectHttpHandler.FetchAll)
	// papi.Post("/", projectHttpHandler.Create)
	// papi.Delete("/:id", projectHttpHandler.DeleteOneById)
	// papi.Put("/:id", projectHttpHandler.Update)

	// if err = app.Listen(":8000"); err != nil {
	// 	log.Fatal(err)
	// }
}
