package secretsmanager

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

var (
	sm SM
)

type Config struct {
	Region                  string
	DirectorySaltSecretName string
}

type SM struct {
	directoryClientSalt string
	directoryServerSalt string
}

type directorySalt struct {
	SaltS string `json:"saltS"`
	SaltC string `json:"saltC"`
}

func GetSM() *SM {
	return &sm
}

func (s *SM) setDirectoryClientSalt(salt string) {
	s.directoryClientSalt = salt
}

func (s *SM) setDirectoryServerSalt(salt string) {
	s.directoryServerSalt = salt
}

func (s *SM) GetDirectoryClientSalt() string {
	return s.directoryClientSalt
}

func (s *SM) GetDirectoryServerSalt() string {
	return s.directoryServerSalt
}

func Init(c Config) *SM {
	secretName := c.DirectorySaltSecretName
	region := c.Region

	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatal(err)
	}

	// Create Secrets Manager client
	svc := secretsmanager.NewFromConfig(config)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		// For a list of exceptions thrown, see
		// https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html
		log.Fatal(err.Error())
	}

	// Decrypts secret using the associated KMS key.
	var secretString string = *result.SecretString
	var secret directorySalt
	err = json.Unmarshal([]byte(secretString), &secret)
	if err != nil {
		log.Fatal(err.Error())
	}
	sm.setDirectoryClientSalt(secret.SaltC)
	sm.setDirectoryServerSalt(secret.SaltS)

	return &sm
}
