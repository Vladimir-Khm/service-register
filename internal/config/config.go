package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	LogLevel            string `mapstructure:"LOG_LEVEL"`
	PayloadSignatureKey string `mapstructure:"PAYLOAD_SIGNATURE_KEY"`

	DbHost               string `mapstructure:"DB_HOST"`
	DbPort               string `mapstructure:"DB_PORT"`
	DbUser               string `mapstructure:"DB_USER"`
	DbPassword           string `mapstructure:"DB_PASSWORD"`
	DbName               string `mapstructure:"DB_NAME"`
	DbMaxOpenConns       int    `mapstructure:"DB_MAX_OPEN_CONNS"`
	DbMaxIdleConns       int    `mapstructure:"DB_MAX_IDLE_CONNS"`
	DbConnMaxLifetimeSec int64  `mapstructure:"DB_CONN_MAX_LIFETIME_SEC"`
	DbConnMaxIdleTimeSec int64  `mapstructure:"DB_CONN_MAX_IDLE_TIME_SEC"`

	BotToken   string `mapstructure:"BOT_TOKEN"`
	MiniAppUrl string `mapstructure:"MINI_APP_URL"`

	TonProofPayloadSignatureKey string `mapstructure:"TONPROOF_PAYLOAD_SIGNATURE_KEY"`
	//TonProofPayloadLifeTimeSec  int64  `mapstructure:"TONPROOF_PAYLOAD_LIFETIME_SEC"`
	TonProofLifeTimeSec               int64   `mapstructure:"TONPROOF_PROOF_LIFETIME_SEC"`
	TonProofDomain                    string  `mapstructure:"TONPROOF_DOMAIN"`
	BotID                             int64   `mapstructure:"BOT_ID"`
	AdminsTelegramIDs                 []int64 `mapstructure:"ADMINS_TELEGRAM_IDS"`
	TelegramGroupsRequireVerification []int64 `mapstructure:"GROUPS_REQUIRE_VERIFICATION"`

	GoogleCloudStorageBucket   string `mapstructure:"GOOGLE_CLOUD_STORAGE_BUCKET"`
	GoogleCloudCredentialsJson string `mapstructure:"GOOGLE_CLOUD_CREDENTIALS_JSON"`

	TonIdPartnerID                 string `mapstructure:"PARTNER_ID"`
	TonIdApiKey                    string `mapstructure:"SBT_API_KEY"`
	HistoryImportMessagesRateLimit int    `mapstructure:"HISTORY_IMPORT_MESSAGES_RATE_LIMIT"`
}

func LoadConfig() (*Config, error) {
	config := &Config{}
	viper.SetConfigFile(".env")

	// ignore error to make .env config not mandatory
	viper.ReadInConfig()

	viper.AutomaticEnv()

	err := viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (c *Config) GetGoogleCloudStorageBucket() string {
	return c.GoogleCloudStorageBucket
}
func (c *Config) GetGoogleCloudCredentialsJson() string {
	return c.GoogleCloudCredentialsJson
}
