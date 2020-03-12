package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

var (
	metricNamespace string
	metricName      string
	url             string
)

type Request struct{}

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
	response, errors := http.Get(web)
	if errors != nil {
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
		log.Fatal("Target website returned ", response.Status)
	}

	return false
}

func main() {
	lambda.Start(HandleRequest)
}
