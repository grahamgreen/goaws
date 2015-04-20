package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"

	"github.com/m8ncman/goutils"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
)

func main() {
	auth, err := aws.EnvAuth()
	goutils.Check(err)

	//TODO check the len of the args
	//TODO use multi-put for faster upload

	bucket_name := os.Getenv("AWS_BUCKET")
	goutils.NotEmpty(bucket_name)
	file_name := os.Args[1]
	goutils.NotEmpty(file_name)

	path := os.Args[2]
	path = goutils.DefaultIfEmpty(path, "NeedAName.sad")

	//log.Print(file_name)
	//log.Print(bucket_name)

	file, err := os.Open(file_name)
	goutils.Check(err)

	defer file.Close()

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	bytes := make([]byte, size)

	buffer := bufio.NewReader(file)
	_, err = buffer.Read(bytes)

	filetype := http.DetectContentType(bytes)

	client := s3.New(auth, aws.USEast)
	bucket := client.Bucket(bucket_name)

	err = bucket.Put(path, bytes, filetype, s3.ACL("public-read"))
	goutils.Check(err)

	fmt.Printf("Uploaded to %s with %v bytes to S3.\n\n", path, size)
	//buk := client.Bucket("vagrant-clearcare")

	//log.Print(buk.Name)
	//contents, err := buk.GetBucketContents()
	//util.check(err)

	//for key, _ := range *contents {
	//	log.Print(key)
	//}
}
