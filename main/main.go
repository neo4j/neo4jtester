package main

import (
	"context"
	"fmt"
	"github.com/neo4jtester"
	"os"
)

func main() {

	checkArguments()
	uri := os.Args[1]
	username := os.Args[2]
	password := os.Args[3]
	licenseType := os.Args[4]

	db := neo4jtester.NewNeo4jDatabase(uri, username, password)
	
	err := db.CreateMoviesDataSet(context.Background())
	if err != nil {
		fmt.Println("error seen while creating movies dataset %v", err)
		os.Exit(1)
	}

	err = db.VerifyMoviesDataSet(context.Background())
	if err != nil {
		fmt.Println("failed while verifying movies dataset %v", err)
		os.Exit(1)
	}

	if licenseType == "Evaluation" {
		err = db.CheckEvaluationLicense(context.Background())
		if err != nil {
			fmt.Println("failed while verifying evaluation license type %v", err)
			os.Exit(1)
		}
	}

}

func checkArguments() {
	if len(os.Args) != 5 {
		fmt.Println("Please provide uri , username, password and license type. (4 command line arguments expected)")
		os.Exit(1)
	}
}
