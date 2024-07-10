package main

import (
    "database/sql"
    "fmt"
    "log"
    "github.com/bryanwahyu/test-golang/api"
    "github.com/bryanwahyu/test-golang/internal/grpc"
    "github.com/bryanwahyu/test-golang/internal/repository"
    "github.com/bryanwahyu/test-golang/internal/service"
    "os"

    _ "github.com/lib/pq"
    "gopkg.in/yaml.v2"
    "google.golang.org/grpc"
    "net"
)

type Config struct {
    Server struct {
        Port int `yaml:"port"`
    } `yaml:"server"`
    Database struct {
        Host     string `yaml:"host"`
        Port     int    `yaml:"port"`
        User     string `yaml:"user"`
        Password string `yaml:"password"`
        DBName   string `yaml:"dbname"`
    } `yaml:"database"`
}

func loadConfig() (*Config, error) {
    f, err := os.Open("configs/config.yaml")
    if err != nil {
        return nil, err
    }
    defer f.Close()

    var cfg Config
    decoder := yaml.NewDecoder(f)
    if err := decoder.Decode(&cfg); err != nil {
        return nil, err
    }
    return &cfg, nil
}

func main() {
    cfg, err := loadConfig()
    if err != nil {
        log.Fatalf("failed to load config: %v", err)
    }

    connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName)
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

    repo := &repository.EntityRepository{DB: db}
    entityService := &service.EntityService{Repo: repo}
    grpcServer := grpc.NewServer()
    api.RegisterMyServiceServer(grpcServer, &grpc.Server{EntityService: entityService})

    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.Port))
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    log.Printf("server listening at %v", lis.Addr())
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
