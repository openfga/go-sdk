package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/joho/godotenv"
	"github.com/openfga/go-sdk/client"
	"github.com/openfga/go-sdk/credentials"
	"github.com/openfga/go-sdk/telemetry"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

var serviceName = semconv.ServiceNameKey.String("openfga-opentelemetry-example")

func configureOpenTelemetry(ctx context.Context, res *resource.Resource, conn *grpc.ClientConn) (func(context.Context) error, error) {
	metricExporter, _ := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithGRPCConn(conn))

	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(metricExporter)),
		sdkmetric.WithResource(res),
	)
	otel.SetMeterProvider(meterProvider)

	return meterProvider.Shutdown, nil
}

func configureOpenFga() (*client.OpenFgaClient, error) {
	creds := credentials.Credentials{
		Method: credentials.CredentialsMethodClientCredentials,
		Config: &credentials.Config{
			ClientCredentialsClientId:       os.Getenv("FGA_CLIENT_ID"),
			ClientCredentialsClientSecret:   os.Getenv("FGA_CLIENT_SECRET"),
			ClientCredentialsApiAudience:    os.Getenv("FGA_API_AUDIENCE"),
			ClientCredentialsApiTokenIssuer: os.Getenv("FGA_API_TOKEN_ISSUER"),
		},
	}

	otel := telemetry.Configuration{
		Metrics: &telemetry.MetricsConfiguration{
			METRIC_COUNTER_CREDENTIALS_REQUEST: &telemetry.MetricConfiguration{
				ATTR_FGA_CLIENT_REQUEST_CLIENT_ID: &telemetry.AttributeConfiguration{Enabled: true},
				ATTR_HTTP_REQUEST_METHOD:          &telemetry.AttributeConfiguration{Enabled: true},
				ATTR_FGA_CLIENT_REQUEST_MODEL_ID:  &telemetry.AttributeConfiguration{Enabled: true},
				ATTR_FGA_CLIENT_REQUEST_STORE_ID:  &telemetry.AttributeConfiguration{Enabled: true},
				ATTR_FGA_CLIENT_RESPONSE_MODEL_ID: &telemetry.AttributeConfiguration{Enabled: true},
				ATTR_HTTP_HOST:                    &telemetry.AttributeConfiguration{Enabled: true},
				ATTR_HTTP_REQUEST_RESEND_COUNT:    &telemetry.AttributeConfiguration{Enabled: true},
				ATTR_HTTP_RESPONSE_STATUS_CODE:    &telemetry.AttributeConfiguration{Enabled: true},
				ATTR_URL_FULL:                     &telemetry.AttributeConfiguration{Enabled: true},
				ATTR_URL_SCHEME:                   &telemetry.AttributeConfiguration{Enabled: true},
				ATTR_USER_AGENT_ORIGINAL:          &telemetry.AttributeConfiguration{Enabled: true},
			},
			METRIC_HISTOGRAM_REQUEST_DURATION: &telemetry.MetricConfiguration{
				ATTR_FGA_CLIENT_REQUEST_CLIENT_ID: &telemetry.AttributeConfiguration{Enabled: true},
				ATTR_HTTP_REQUEST_METHOD:          &telemetry.AttributeConfiguration{Enabled: true},
				ATTR_FGA_CLIENT_REQUEST_MODEL_ID:  &telemetry.AttributeConfiguration{Enabled: true},
				ATTR_FGA_CLIENT_REQUEST_STORE_ID:  &telemetry.AttributeConfiguration{Enabled: true},
				ATTR_FGA_CLIENT_RESPONSE_MODEL_ID: &telemetry.AttributeConfiguration{Enabled: true},
				ATTR_HTTP_HOST:                    &telemetry.AttributeConfiguration{Enabled: true},
				ATTR_HTTP_REQUEST_RESEND_COUNT:    &telemetry.AttributeConfiguration{Enabled: true},
				ATTR_HTTP_RESPONSE_STATUS_CODE:    &telemetry.AttributeConfiguration{Enabled: true},
				ATTR_URL_FULL:                     &telemetry.AttributeConfiguration{Enabled: true},
				ATTR_URL_SCHEME:                   &telemetry.AttributeConfiguration{Enabled: true},
				ATTR_USER_AGENT_ORIGINAL:          &telemetry.AttributeConfiguration{Enabled: true},
			},
			METRIC_HISTOGRAM_QUERY_DURATION: &telemetry.MetricConfiguration{
				ATTR_FGA_CLIENT_REQUEST_CLIENT_ID: &telemetry.AttributeConfiguration{Enabled: true},
				ATTR_HTTP_REQUEST_METHOD:          &telemetry.AttributeConfiguration{Enabled: true},
				ATTR_FGA_CLIENT_REQUEST_MODEL_ID:  &telemetry.AttributeConfiguration{Enabled: true},
				ATTR_FGA_CLIENT_REQUEST_STORE_ID:  &telemetry.AttributeConfiguration{Enabled: true},
				ATTR_FGA_CLIENT_RESPONSE_MODEL_ID: &telemetry.AttributeConfiguration{Enabled: true},
				ATTR_HTTP_HOST:                    &telemetry.AttributeConfiguration{Enabled: true},
				ATTR_HTTP_REQUEST_RESEND_COUNT:    &telemetry.AttributeConfiguration{Enabled: true},
				ATTR_HTTP_RESPONSE_STATUS_CODE:    &telemetry.AttributeConfiguration{Enabled: true},
				ATTR_URL_FULL:                     &telemetry.AttributeConfiguration{Enabled: true},
				ATTR_URL_SCHEME:                   &telemetry.AttributeConfiguration{Enabled: true},
				ATTR_USER_AGENT_ORIGINAL:          &telemetry.AttributeConfiguration{Enabled: true},
			},
		},
	}

	fgaClient, err := client.NewSdkClient(&client.ClientConfiguration{
		ApiUrl:               os.Getenv("FGA_API_URL"),
		StoreId:              os.Getenv("FGA_STORE_ID"),
		AuthorizationModelId: os.Getenv("FGA_MODEL_ID"),
		Credentials:          &creds,
		Telemetry:            &otel,
	})

	return fgaClient, err
}

func main() {
	godotenv.Load()

	ctx := context.Background()

	conn, _ := grpc.NewClient("localhost:4317",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	res, _ := resource.New(ctx,
		resource.WithAttributes(
			serviceName,
		),
	)

	shutdownMeterProvider, _ := configureOpenTelemetry(ctx, res, conn)

	defer func() {
		if err := shutdownMeterProvider(ctx); err != nil {
			log.Fatalf("failed to shutdown MeterProvider: %s", err)
		}
	}()

	fgaClient, _ := configureOpenFga()

	fmt.Println("Read Authorization Models")
	authModels, _ := fgaClient.ReadAuthorizationModels(ctx).Execute()
	fmt.Printf("Authorization Models Count: %d\n", len(authModels.AuthorizationModels))
}
