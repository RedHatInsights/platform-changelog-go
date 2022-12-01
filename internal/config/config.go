package config

import (
	"encoding/json"
	"os"
	"strings"

	_ "embed"

	clowder "github.com/redhatinsights/app-common-go/pkg/api/v1"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

//go:embed services.yaml
var services []byte

//go:embed tenant.yaml
var tenants []byte

type Config struct {
	PublicPort             string
	MetricsPort            string
	MetricsPath            string
	LogLevel               string
	Hostname               string
	CloudwatchConfig       CloudwatchCfg
	DatabaseConfig         DatabaseCfg
	GithubWebhookSecretKey string
	GitlabWebhookSecretKey string
	Services               map[string]Service
	Tenants                map[string]Tenant
	Debug                  bool
	DBImpl                 string
	OpenAPISpec            []byte
}

type DatabaseCfg struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
	DBSSLMode  string
	RDSCa      string
}

type CloudwatchCfg struct {
	CWLogGroup  string
	CWRegion    string
	CWAccessKey string
	CWSecretKey string
}

type Service struct {
	DisplayName string `yaml:"display_name"`
	Tenant      string `yaml:"tenant"`
	GHRepo      string `yaml:"gh_repo,omitempty"`
	GLRepo      string `yaml:"gl_repo,omitempty"`
	Branch      string `yaml:"branch"`
	Namespace   string `yaml:"namespace,omitempty"`
	DeployFile  string `yaml:"deploy_file,omitempty"`
}

type Tenant struct {
	Name string `yaml:"name"`
}

func readOpenAPISpec() []byte {
	var openAPISpec OpenAPISpec
	openAPISpecFile, err := os.Open("schema/openapi.yaml")
	if err != nil {
		panic(err)
	}
	defer openAPISpecFile.Close()
	decoder := yaml.NewDecoder(openAPISpecFile)
	err = decoder.Decode(&openAPISpec)
	if err != nil {
		panic(err)
	}
	openAPISpecJSON, err := json.Marshal(openAPISpec)
	if err != nil {
		panic(err)
	}
	return openAPISpecJSON
}

func Get() *Config {
	options := viper.New()

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	loglevel := os.Getenv("LOGLEVEL")
	if loglevel == "" {
		loglevel = "ERROR"
	}

	dbImpl := os.Getenv("DB_IMPL")
	if dbImpl == "" {
		dbImpl = "impl"
	}

	// global logging
	options.SetDefault("logLevel", loglevel)
	options.SetDefault("Hostname", hostname)
	options.SetDefault("GithubSecretKey", os.Getenv("GITHUB_SECRET_KEY"))
	options.SetDefault("GitlabSecretKey", os.Getenv("GITLAB_SECRET_KEY"))
	options.SetDefault("Debug", os.Getenv("DEBUG") == "true" || os.Getenv("DEBUG") == "1")
	options.SetDefault("db.impl", dbImpl)

	if clowder.IsClowderEnabled() {
		cfg := clowder.LoadedConfig

		// ports
		options.SetDefault("publicPort", cfg.PublicPort)
		options.SetDefault("metricsPort", cfg.MetricsPort)
		options.SetDefault("metricsPath", cfg.MetricsPath)
		// database
		options.SetDefault("db.user", cfg.Database.Username)
		options.SetDefault("db.password", cfg.Database.Password)
		options.SetDefault("db.name", cfg.Database.Name)
		options.SetDefault("db.host", cfg.Database.Hostname)
		options.SetDefault("db.port", cfg.Database.Port)
		options.SetDefault("rdsCa", cfg.Database.RdsCa)
		options.SetDefault("db.sslmode", "verify-full")
		// cloudwatch
		options.SetDefault("logGroup", cfg.Logging.Cloudwatch.LogGroup)
		options.SetDefault("cwRegion", cfg.Logging.Cloudwatch.Region)
		options.SetDefault("cwAccessKey", cfg.Logging.Cloudwatch.AccessKeyId)
		options.SetDefault("cwSecretKey", cfg.Logging.Cloudwatch.SecretAccessKey)
	} else {
		// ports
		options.SetDefault("publicPort", "8000")
		options.SetDefault("metricsPort", "9001")
		options.SetDefault("metricsPath", "/metrics")
		// database
		options.SetDefault("db.user", "crc")
		options.SetDefault("db.password", "crc")
		options.SetDefault("db.name", "gumbaroo")
		options.SetDefault("db.host", "0.0.0.0")
		options.SetDefault("db.port", "5432")
		options.SetDefault("db.sslmode", "disable")
		// cloudwatch
		options.SetDefault("logGroup", "platform-dev")
		options.SetDefault("cwRegion", "us-east-1")
		options.SetDefault("cwAccessKey", os.Getenv("CW_AWS_ACCESS_KEY_ID"))
		options.SetDefault("cwSecretKey", os.Getenv("CW_AWS_SECRET_ACCESS_KEY"))
	}

	options.AutomaticEnv()
	options.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	config := &Config{
		Hostname:               options.GetString("Hostname"),
		LogLevel:               options.GetString("logLevel"),
		GithubWebhookSecretKey: options.GetString("GithubSecretKey"),
		PublicPort:             options.GetString("publicPort"),
		MetricsPort:            options.GetString("metricsPort"),
		MetricsPath:            options.GetString("metricsPath"),
		Debug:                  options.GetBool("Debug"),
		DBImpl:                 options.GetString("db.impl"),
		OpenAPISpec:            readOpenAPISpec(),
		DatabaseConfig: DatabaseCfg{
			DBUser:     options.GetString("db.user"),
			DBPassword: options.GetString("db.password"),
			DBName:     options.GetString("db.name"),
			DBHost:     options.GetString("db.host"),
			DBPort:     options.GetString("db.port"),
			DBSSLMode:  options.GetString("db.sslmode"),
		},
		CloudwatchConfig: CloudwatchCfg{
			CWLogGroup:  options.GetString("logGroup"),
			CWRegion:    options.GetString("cwRegion"),
			CWAccessKey: options.GetString("cwAccessKey"),
			CWSecretKey: options.GetString("cwSecretKey"),
		},
	}

	if clowder.IsClowderEnabled() {

		// write the RDS CA using the app-common-go package
		if clowder.LoadedConfig.Database.RdsCa != nil {
			rdsCAPath, err := clowder.LoadedConfig.RdsCa()

			if err != nil {
				panic("RDS CA Failed to Write")
			}

			config.DatabaseConfig.RDSCa = rdsCAPath
		}
	}

	// read in services.yaml to the config
	err = yaml.Unmarshal(services, config)
	if err != nil {
		panic("Unable to read services.yaml")
	}

	// read in tenant.yaml to the config
	err = yaml.Unmarshal(tenants, config)
	if err != nil {
		panic("Unable to read tenants.yaml")
	}

	return config
}
