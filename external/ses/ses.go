package ses

import (
	"BootCampT1/config"
	"BootCampT1/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

var conn *ses.SES

type EmailConfig struct {
	Sender       string
	Recipient    []string
	CcAddresses  []string
	BccAddresses []string
	Subject      string
	HtmlBody     string
	CharSet      string
	TextBody     string
}

func Init() {
	conf := config.GetConfig()
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(conf.Aws.Region),
		Credentials: credentials.NewSharedCredentials("", conf.Aws.Profile),
	})
	if err != nil {
		logger.Error.Println(err)
		return
	}
	conn = ses.New(sess)
	logger.Info.Println("Successfully created ses session")
}
func SendMail(inp EmailConfig) {

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: aws.StringSlice(inp.CcAddresses),
			ToAddresses: aws.StringSlice(inp.Recipient),
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(inp.CharSet),
					Data:    aws.String(inp.HtmlBody),
				},
				Text: &ses.Content{
					Charset: aws.String(inp.CharSet),
					Data:    aws.String(inp.TextBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(inp.CharSet),
				Data:    aws.String(inp.Subject),
			},
		},
		Source: aws.String(inp.Sender),
	}

	_, err := conn.SendEmail(input)

	// Display error messages if they occur.
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				logger.Error.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				logger.Error.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				logger.Error.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				logger.Error.Println(aerr.Error())
			}
		} else {
			logger.Error.Println(err.Error())
		}
		return
	}
}
