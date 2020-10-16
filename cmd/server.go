package cmd

import (
	"fmt"
	"log"
	"net"
	"tag/controller"
	"tag/env"
	pbdemo "tag/grpc"
	"tag/router"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Start gRPC Server on Port : %v", env.PORT)
		startgRPC()
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&env.PORT, "port", "p", "5000", "grpc server port")
}

func run() {
	fmt.Println("tag run start")

	env.Step()

	// 實例 echo 對象
	e := echo.New()

	// 把echo對象，加入 "middleware.Logger()" middleware
	e.Use(middleware.Logger())

	// 設定 endPoint
	router.Set(e)

	// 啟動服務
	e.Start(":" + env.PORT)
}

func startgRPC() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", env.PORT))
	if err != nil {
		log.Fatalf("start grpc server error : %v", err)
	}
	l := logrus.New().WithField("service", "demo")

	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_logrus.UnaryServerInterceptor(l),
				grpc_recovery.UnaryServerInterceptor(),
			),
		),
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				grpc_logrus.StreamServerInterceptor(l),
				grpc_recovery.StreamServerInterceptor(),
			),
		),
	)

	reflection.Register(s)

	// grpc 設定

	// 註冊服務
	pbdemo.RegisterTagServer(s, &controller.TagServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to start server : %v", err)
	}

	return nil
}
