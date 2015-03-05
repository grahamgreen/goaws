package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"

	"github.com/m8ncman/goutils"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/ec2"
)

func main() {
	auth, err := aws.EnvAuth()
	goutils.Check(err)

	client := ec2.New(auth, aws.USEast)
	resp, err := client.Instances()
	goutils.Check(err)

	res := resp.Reservations

}
