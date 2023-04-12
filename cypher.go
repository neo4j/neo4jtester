package neo4jtester

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
)

func (db *Neo4jDatabase) CheckEvaluationLicense(ctx context.Context) error {
	log.Println("Checking License Type")
	driver, err := neo4j.NewDriverWithContext(db.uri, neo4j.BasicAuth(db.username, db.password, ""))
	if err != nil {
		return err
	}
	defer driver.Close(ctx)

	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	value, err := session.ExecuteWrite(ctx, func(transaction neo4j.ManagedTransaction) (any, error) {
		result, err := transaction.Run(ctx,
			EVALUATION_LICENSE_CYPHER,
			nil)
		if err != nil {
			return nil, err
		}

		records, err := result.Collect(ctx)
		if err != nil {
			return nil, err
		}
		if len(records) > 1 {
			return nil, fmt.Errorf("expected only one record , recieved more %v", records)
		}
		if len(records) == 0 {
			return nil, fmt.Errorf("expected one record , received none")
		}
		value, found := records[0].Get("value")
		if !found {
			return nil, fmt.Errorf("cannot find the key 'value'")
		}
		return value, nil
	})

	if err != nil {
		return err
	}

	if value.(string) != "30" {
		return fmt.Errorf("exepcted '30' as evaluation license details , got %v", value)
	}
	log.Println("Verified Evaluation License Type!!")
	return nil
}
