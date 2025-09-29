package main

import (
	grpcDelivery "basic-personal-financial-tracking-api/service/module/delivery/grpc"
	restfulDelivery "basic-personal-financial-tracking-api/service/module/delivery/restful"
	"basic-personal-financial-tracking-api/service/module/repository"
	"basic-personal-financial-tracking-api/service/module/use_case"
	"fmt"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

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
	//TODO: Use ENV
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", "localhost", "5433", "pg", "crud", "pass", "disable")
	client, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}

	err = MigrateDB(client)
	if err != nil {
		log.Fatal(err)
	}

	return client, err
}

func MigrateDB(dbConn *gorm.DB) (err error) {
	// err = dbConn.AutoMigrate(
	// 	&model.TbTodoTableSchema{},
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	return err
}
