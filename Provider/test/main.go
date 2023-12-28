package main

import (
	"fmt"

	air "github.com/afsabel/client"
)

func main() {

	HostURL := "http://localhost:5001"
	ClientId := ""
	ClientSecret := "n148Q~mDL5BK3iLQ0VXXeKmmVedWFjcsgZbDRdv-"

	c, err := air.NewClient(&HostURL, &ClientId, &ClientSecret)

	if err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
	}

	applications, err := c.GetApplications()

	if err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
	} else {
		fmt.Println(applications)
	}
}
