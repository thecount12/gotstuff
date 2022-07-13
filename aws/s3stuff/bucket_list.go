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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// initialize session to load ~/.aws/credentials

func main() {

//  default creds
sess := session.Must(session.NewSessionWithOptions(session.Options{
    SharedConfigState: session.SharedConfigEnable,
}))
// s3 client
svc := s3.New(sess)

// List, passing nil means no filters applied
result, err := svc.ListBuckets(nil)
check(err)

fmt.Println("Buckets:")

for _, b := range result.Buckets {
    fmt.Printf("* %s created on %s\n",
        aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}

}
