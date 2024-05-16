package main

import (
	"log"
	"net"

	"go-grpc-skeleton/config"
	"go-grpc-skeleton/internal/handler"
	"go-grpc-skeleton/internal/pkg/elasticsearch"
	"go-grpc-skeleton/internal/pkg/mysql"
	"go-grpc-skeleton/internal/pkg/redis"
	repo "go-grpc-skeleton/internal/repository"
	"go-grpc-skeleton/internal/service"
	pb "go-grpc-skeleton/proto"

	"go-grpc-skeleton/global"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {

	// 加载配置
	etcdEndpoints := []string{"localhost:2379"}
	configKey := "/config/myapp"
	cfg, err := config.LoadConfigFromEtcd(etcdEndpoints, configKey)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// 初始化日志记录器
	logger := global.Logger()
	// sugar := global.Sugar()

	// 初始化 GORM MySQL 客户端
	db, err := mysql.NewGormClient(cfg.MySQL)
	if err != nil {
		logger.Error("Failed to initialize GORM MySQL client", zap.Error(err))
	}

	// 初始化 Redis 客户端
	redisClient, err := redis.NewRedisClient(cfg.Redis)
	if err != nil {
		logger.Error("failed to connect to Redis", zap.Error(err))
	}
	defer redisClient.Close()

	// 初始化 Elasticsearch 客户端
	esClient, err := elasticsearch.NewElasticsearchClient(cfg.Elasticsearch)
	if err != nil {
		logger.Error("failed to connect to Elasticsearch: ", zap.Error(err))
	}

	// 初始化 product 模块
	productRepo := repo.NewProductRepository(db, redisClient, esClient)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	// 初始化 category 模块
	categoryRepo := repo.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, productHandler)
	pb.RegisterCategoryServiceServer(grpcServer, categoryHandler)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Error("failed to listen: ", zap.Error(err))
	}

	logger.Info("gRPC server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		logger.Error("failed to serve: ", zap.Error(err))
	}
}
