package api

import (
	"fmt"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	_ "github.com/nikhilnarayanan623/x-tention-crew/user-servcie/cmd/api/docs"
	"github.com/nikhilnarayanan623/x-tention-crew/user-servcie/pkg/api/handler/interfaces"
	"github.com/nikhilnarayanan623/x-tention-crew/user-servcie/pkg/api/routes"
	"github.com/nikhilnarayanan623/x-tention-crew/user-servcie/pkg/config"
	"github.com/nikhilnarayanan623/x-tention-crew/user-servcie/pkg/pb"
	"github.com/nikhilnarayanan623/x-tention-crew/user-servcie/pkg/utils"
	"google.golang.org/grpc"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	basePath = "/api"
)

type Server struct {
	restServer *gin.Engine
	restPort   string

	grpcServer *grpc.Server
	ls         net.Listener
}

func NewServer(cfg config.Config,
	serviceServer pb.UserServiceServer, userHandler interfaces.UserHandler) (*Server, error) {

	// register rest server
	engine := gin.New()
	engine.Use(gin.Logger())

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// register all REST api routes
	routes.RegisterRoutes(engine.Group(basePath), userHandler)

	// register grpc server
	addrs := fmt.Sprintf("%s:%s", cfg.UserServiceGrpcHost, cfg.UserServiceGrpcPort)
	ls, err := net.Listen("tcp", addrs)
	if err != nil {
		return nil, utils.PrependMessageToError(err, "failed to create net listener")
	}
	// create new grpc server
	gs := grpc.NewServer()
	// register service server with grpc server
	pb.RegisterUserServiceServer(gs, serviceServer)

	return &Server{
		restServer: engine,
		restPort:   cfg.UserServiceRestPort,

		grpcServer: gs,
		ls:         ls,
	}, nil
}

func (s *Server) Start() {

	errChan := make(chan error)
	// run grpc server on separate goroutine
	go func() {
		fmt.Println("user grpc server running..")
		errChan <- s.grpcServer.Serve(s.ls)
	}()
	// run REST server on separate goroutine
	go func() {
		errChan <- s.restServer.Run(":" + s.restPort)
	}()

	// wait for any of one error in main goroutine
	err := <-errChan
	log.Fatal("failed to run server: ", err)
}
