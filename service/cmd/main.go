package main

import (
	"basic-personal-financial-tracking-api/service/database"
	"basic-personal-financial-tracking-api/service/module/repository"
	"basic-personal-financial-tracking-api/service/module/use_case"
	"net"

	"errors"
	"fmt"
	"log"

	grpcDelivery "basic-personal-financial-tracking-api/service/module/delivery/grpc"
	restfulDelivery "basic-personal-financial-tracking-api/service/module/delivery/restful"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	envDATABASE_HOST     = "DATABASE_HOST"
	envDATABASE_NAME     = "DATABASE_NAME"
	envDATABASE_USER     = "DATABASE_USER"
	envDATABASE_PASSWORD = "DATABASE_PASSWORD"
)

var (
	errFailedToBindingEnvironment = errors.New("failed to binding open-telemetry-jaeger environment")
	envList                       = []string{

		envDATABASE_HOST,
		envDATABASE_NAME,
		envDATABASE_USER,
		envDATABASE_PASSWORD,
	}
)

func init() {
	var err error

	for i := 0; i < len(envList); i++ {
		if err = viper.BindEnv(envList[i]); err != nil {
			log.Fatalf("%s: %v", errFailedToBindingEnvironment, err)
		}
	}

}

func main() {

	var err error

	dbConn, err := connectDB()

	if err != nil {
		log.Panicln("failed to connect database", err)
	} else {
		fmt.Println("Connect ok", dbConn)
	}
	engine := gin.New()

	repo := repository.NewRepository(dbConn)
	use := use_case.NewUseCase(repo)

	restfulDelivery.NewHandler(engine, use)

	lis, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	grpcDelivery.NewServerGrpc(server, use)

	if err = server.Serve(lis); err != nil {
		panic(err)
	}

	fmt.Println("started")

	err = engine.Run("localhost:8080")

	if err != nil {
		log.Fatalln("failed to run", err)
	}
}

func connectDB() (client *gorm.DB, err error) {

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", viper.GetString("DATABASE_HOST"), "5432", viper.GetString("DATABASE_USER"), viper.GetString("DATABASE_NAME"), viper.GetString("DATABASE_PASSWORD"), "disable")
	client, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	fmt.Println("Running AutoMigrate...")
	err = database.MigrateDB(client)
	if err != nil {
		log.Fatal("failed to migrate db", err)
	}
	fmt.Println("AutoMigrate done")
	return client, err
}
