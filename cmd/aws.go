package cmd

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/sns"
)

// AWSSessionManager manages the AWS session with lazy initialization
type AWSSessionManager struct {
	mu      sync.Mutex
	session *session.Session
}

// getSession creates or returns an existing AWS session
func (sm *AWSSessionManager) getSession() (*session.Session, error) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	if sm.session != nil {
		return sm.session, nil
	}

	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(os.Getenv("AWS_REGION")),
		},
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return nil, err
	}

	sm.session = sess
	return sm.session, nil
}

// closeSession terminates the AWS session
func (sm *AWSSessionManager) closeSession() {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	// AWS sessions don't need explicit closing, but we can nullify the reference
	sm.session = nil
}

// SNSService manages SNS operations
type SNSService struct {
	sessionManager *AWSSessionManager
	client         *sns.SNS
}

// NewSNSService creates a new SNS service client
func NewSNSService() *SNSService {
	return &SNSService{
		sessionManager: &AWSSessionManager{},
	}
}

// getClient initializes the SNS client if needed
func (s *SNSService) getClient() (*sns.SNS, error) {
	if s.client != nil {
		return s.client, nil
	}

	sess, err := s.sessionManager.getSession()
	if err != nil {
		return nil, err
	}

	s.client = sns.New(sess)
	return s.client, nil
}

// Close releases SNS client resources
func (s *SNSService) Close() {
	s.client = nil
	s.sessionManager.closeSession()
}

// TestConnection verifies SNS connectivity
func (s *SNSService) TestConnection(ctx context.Context) bool {
	client, err := s.getClient()
	if err != nil {
		log.Printf("SNS connection failed: %v", err)
		return false
	}

	_, err = client.ListTopicsWithContext(ctx, &sns.ListTopicsInput{})
	if err != nil {
		log.Printf("SNS connection test failed: %v", err)
		return false
	}
	return true
}

// SESService manages SES operations
type SESService struct {
	sessionManager *AWSSessionManager
	client         *ses.SES
}

// NewSESService creates a new SES service client
func NewSESService() *SESService {
	return &SESService{
		sessionManager: &AWSSessionManager{},
	}
}

// getClient initializes the SES client if needed
func (s *SESService) getClient() (*ses.SES, error) {
	if s.client != nil {
		return s.client, nil
	}

	sess, err := s.sessionManager.getSession()
	if err != nil {
		return nil, err
	}

	s.client = ses.New(sess)
	return s.client, nil
}

// Close releases SES client resources
func (s *SESService) Close() {
	s.client = nil
	s.sessionManager.closeSession()
}

// TestConnection verifies SES connectivity
func (s *SESService) TestConnection(ctx context.Context) bool {
	client, err := s.getClient()
	if err != nil {
		log.Printf("SES connection failed: %v", err)
		return false
	}

	_, err = client.GetSendQuotaWithContext(ctx, &ses.GetSendQuotaInput{})
	if err != nil {
		log.Printf("SES connection test failed: %v", err)
		return false
	}
	return true
}

// LambdaService manages Lambda operations
type LambdaService struct {
	sessionManager *AWSSessionManager
	client         *lambda.Lambda
}

// NewLambdaService creates a new Lambda service client
func NewLambdaService() *LambdaService {
	return &LambdaService{
		sessionManager: &AWSSessionManager{},
	}
}

// getClient initializes the Lambda client if needed
func (l *LambdaService) getClient() (*lambda.Lambda, error) {
	if l.client != nil {
		return l.client, nil
	}

	sess, err := l.sessionManager.getSession()
	if err != nil {
		return nil, err
	}

	l.client = lambda.New(sess)
	return l.client, nil
}

// Close releases Lambda client resources
func (l *LambdaService) Close() {
	l.client = nil
	l.sessionManager.closeSession()
}

// TestConnection verifies Lambda connectivity
func (l *LambdaService) TestConnection(ctx context.Context) bool {
	client, err := l.getClient()
	if err != nil {
		log.Printf("Lambda connection failed: %v", err)
		return false
	}

	_, err = client.ListFunctionsWithContext(ctx, &lambda.ListFunctionsInput{})
	if err != nil {
		log.Printf("Lambda connection test failed: %v", err)
		return false
	}
	return true
}

// DynamoService manages DynamoDB operations
type DynamoService struct {
	sessionManager *AWSSessionManager
	client         *dynamodb.DynamoDB
}

// NewDynamoService creates a new DynamoDB service client
func NewDynamoService() *DynamoService {
	return &DynamoService{
		sessionManager: &AWSSessionManager{},
	}
}

// getClient initializes the DynamoDB client if needed
func (d *DynamoService) getClient() (*dynamodb.DynamoDB, error) {
	if d.client != nil {
		return d.client, nil
	}

	sess, err := d.sessionManager.getSession()
	if err != nil {
		return nil, err
	}

	d.client = dynamodb.New(sess)
	return d.client, nil
}

// Close releases DynamoDB client resources
func (d *DynamoService) Close() {
	d.client = nil
	d.sessionManager.closeSession()
}

// TestConnection verifies DynamoDB connectivity
func (d *DynamoService) TestConnection(ctx context.Context) bool {
	client, err := d.getClient()
	if err != nil {
		log.Printf("DynamoDB connection failed: %v", err)
		return false
	}

	_, err = client.ListTablesWithContext(ctx, &dynamodb.ListTablesInput{})
	if err != nil {
		log.Printf("DynamoDB connection test failed: %v", err)
		return false
	}
	return true
}
