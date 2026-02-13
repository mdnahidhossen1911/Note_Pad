package cmd

import (
	"fmt"
	"note_pad/config"
	"note_pad/infra/db"
	"note_pad/routes"
	"os"
)

func Serve() {

	cfg := config.GetConfig()

	// 2. Connect DB
	dbConn, err := db.NewDBConnection(cfg.DbConfig)
	if err != nil {
		fmt.Println("DB connection failed:", err)
		os.Exit(1)
	}

	// 3. Migrate
	if err := db.MigrateDB(dbConn); err != nil {
		fmt.Println("Migration failed:", err)
		os.Exit(1)
	}

	// 4. Boot router
	router := routes.SetupRouter(cfg, dbConn)

	port := fmt.Sprintf(":%s", cfg.Port)
	fmt.Printf("\n🚀 %s (MVC) running at http://localhost%s\n\n", cfg.ServiceName, port)
	router.Run(port)
}
