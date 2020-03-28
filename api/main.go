package main

import (
	"cloud.google.com/go/profiler"
	"gitlab.com/oneplanet/corona-backend/api/controllers"
	"gitlab.com/oneplanet/corona-backend/api/utils/log"
	"os"
)

func setupGceProfiler() {
	// Profiler initialization, best done as early as possible.
	if err := profiler.Start(profiler.Config{
		Service:        "oneplanet-corona-backend-api",
		ServiceVersion: "1.0.0",
		// ProjectID must be set if not running on GCP.
		// ProjectID: "oneplanet-corona-backend",
	}); err != nil {
		log.Error("Error in creating the profiler")
	}
}

func main() {
	log.Setup()
	setupGceProfiler()
	//models.Init()
	r := controllers.SetupRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	_ = r.Run(":" + port)
}
