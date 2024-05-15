package config

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"gopkg.in/yaml.v2"
)

type Config struct {
	MySQL         MySQLConfig         `yaml:"mysql"`
	Redis         RedisConfig         `yaml:"redis"`
	Kafka         KafkaConfig         `yaml:"kafka"`
	Elasticsearch ElasticsearchConfig `yaml:"elasticsearch"`
}

type MySQLConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type KafkaConfig struct {
	Brokers []string `yaml:"brokers"`
}

type ElasticsearchConfig struct {
	Addr string `yaml:"addr"`
}

func LoadConfigFromEtcd(endpoints []string, key string) (*Config, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		return nil, err
	}

	if len(resp.Kvs) == 0 {
		return nil, fmt.Errorf("key %s not found in etcd", key)
	}

	var config Config
	err = yaml.Unmarshal(resp.Kvs[0].Value, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
