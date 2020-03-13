package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

var (
	metricNamespace string
	metricName      string
	url             string
	secretName      string
)

type Request struct{}

type HttpBasicAuth struct {
	Username string
	Password string
}

func HandleRequest(ctx context.Context, req Request) (string, error) {
	metricNamespace = os.Getenv("CW_METRIC_NAMESPACE")
	url = os.Getenv("TARGET_URL")
	metricName = os.Getenv("CW_METRIC_NAME")
	result := webIsReachable(url)

	if result == true {
		pushMetric(1)
	} else {
		pushMetric(0)
	}

	return "Finished", nil
}

func pushMetric(value float64) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("REGION"))},
	)
	if err != nil {
		log.Fatal("Error creating session")
	}

	svc := cloudwatch.New(sess)
	now := time.Now()

	_, err = svc.PutMetricData(&cloudwatch.PutMetricDataInput{
		MetricData: []*cloudwatch.MetricDatum{
			&cloudwatch.MetricDatum{
				MetricName: &metricName,
				Value:      &value,
				Timestamp:  &now,
			},
		},
		Namespace: &metricNamespace,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func webIsReachable(web string) bool {

	auth := getSecret()

	client := &http.Client{}
	req, err := http.NewRequest("GET", web, nil)
	if err != nil {
		panic(err)
	}

	if auth != nil {
		req.SetBasicAuth(auth.Username, auth.Password)
	}

	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		_, netErrors := http.Get("https://www.google.com")
		if netErrors != nil {
			fmt.Fprintf(os.Stderr, "no internet\n")
			os.Exit(1)
		}
		return false
	}

	if response.StatusCode == 200 {
		log.Print("Target website returned ", response.Status)
		return true
	} else {
		log.Print("Target website returned ", response.Status)
	}

	return false
}

func getSecret() *HttpBasicAuth {
	region := os.Getenv("REGION")
	secretName = os.Getenv("SECRET_NAME")

	//Create a Secrets Manager client
	svc := secretsmanager.New(session.New(),
		aws.NewConfig().WithRegion(region))
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	// In this sample we only handle the specific exceptions for the 'GetSecretValue' API.
	// See https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html

	result, err := svc.GetSecretValue(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeDecryptionFailure:
				// Secrets Manager can't decrypt the protected secret text using the provided KMS key.
				fmt.Println(secretsmanager.ErrCodeDecryptionFailure, aerr.Error())

			case secretsmanager.ErrCodeInternalServiceError:
				// An error occurred on the server side.
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())

			case secretsmanager.ErrCodeInvalidParameterException:
				// You provided an invalid value for a parameter.
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())

			case secretsmanager.ErrCodeInvalidRequestException:
				// You provided a parameter value that is not valid for the current state of the resource.
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())

			case secretsmanager.ErrCodeResourceNotFoundException:
				// We can't find the resource that you asked for.
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return nil
	}

	// Decrypts secret using the associated KMS CMK.
	// Depending on whether the secret is a string or binary, one of these fields will be populated.
	var secretString string
	if result.SecretString != nil {
		secretString = *result.SecretString
	}

	var httpBasicAuth HttpBasicAuth
	json.Unmarshal([]byte(secretString), &httpBasicAuth)

	return &httpBasicAuth
}

func main() {
	HandleRequest(nil, Request{})
	// lambda.Start(HandleRequest)
}
