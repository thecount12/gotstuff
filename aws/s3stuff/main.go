package main

/*
 * https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html
 */

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    //"github.com/aws/aws-sdk-go/service/s3/s3manager"
    "fmt"
    "os"
)

// catch error stuff
func exitErrorf(msg string, args ...interface{}) {
    fmt.Fprintf(os.Stderr, msg+"\n", args...)
    os.Exit(1)
}

// initialize session to load ~/.aws/credentials

func main() {

// default region and default creds
sess := session.Must(session.NewSessionWithOptions(session.Options{
    SharedConfigState: session.SharedConfigEnable,
}))


// specific region
/*
sess, err := session.NewSession(&aws.Config{
    Region: aws.String("us-west-2")},
)
*/

// profile
/*
os.Setenv("AWS_PROFILE", "prod")
sess, err := session.NewSession(&aws.Config{
    Region: aws.String("us-west-2"),
    Credentials: credentials.NewSharedCredentials("", "prod"),
})
*/

// manual creds
/*
sess, err := session.NewSession(&aws.Config{
    Region:      aws.String("us-west-2"),
    Credentials: credentials.NewStaticCredentials("AKID", "SECRET_KEY", "TOKEN"),
})
*/

// Create S3 service client
svc := s3.New(sess)

// List, passing nil means no filters applied
result, err := svc.ListBuckets(nil)
if err != nil {
    exitErrorf("Unable to list buckets, %v", err)
}

fmt.Println("Buckets:")

for _, b := range result.Buckets {
    fmt.Printf("* %s created on %s\n",
        aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}

}
