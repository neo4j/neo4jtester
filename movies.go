package neo4jtester

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
)

func (db *Neo4jDatabase) CreateMoviesDataSet(ctx context.Context) error {
	log.Println("Creating movies dataset")
	driver, err := neo4j.NewDriverWithContext(db.uri, neo4j.BasicAuth(db.username, db.password, ""))
	if err != nil {
		return err
	}
	defer driver.Close(ctx)

	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	_, err = session.ExecuteWrite(ctx, func(transaction neo4j.ManagedTransaction) (any, error) {
		result, err := transaction.Run(ctx,
			MOVIES_CYPHER,
			nil)
		if err != nil {
			return nil, err
		}

		resultSummary, err := result.Consume(ctx)
		if err != nil {
			return nil, err
		}
		return resultSummary, nil
	})

	if err != nil {
		return err
	}
	log.Println("Movies DataSet created !!")
	return nil
}

func (db *Neo4jDatabase) VerifyMoviesDataSet(ctx context.Context) error {
	log.Println("Verifying movies dataset")
	driver, err := neo4j.NewDriverWithContext(db.uri, neo4j.BasicAuth(db.username, db.password, ""))
	if err != nil {
		return err
	}
	defer driver.Close(ctx)

	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	_, err = session.ExecuteWrite(ctx, func(transaction neo4j.ManagedTransaction) (any, error) {
		result, err := transaction.Run(ctx,
			`MATCH (N) RETURN N;`,
			nil)
		if err != nil {
			return nil, err
		}

		records, err := result.Collect(ctx)
		if err != nil {
			return nil, err
		}
		if len(records) > 0 {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get any records %v", err)
	})

	if err != nil {
		return err
	}
	log.Println("Verification Successful !!")
	return nil
}
