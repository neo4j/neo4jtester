# neo4jtester

This is a Go application written to execute tests on Neo4j Database.
It expects four arguments:

    Neo4j URI <neo4j://<>:7687)
    Username
    Password
    License Type (Evaluation/Enterprise)

The tests it performs so far are:

    Create a Movie DataSet
    Verify the Movie DataSet (By verifying the number of records)

### Build
In order to build this application for linux  (linux is what we use in Github Actions) run the following command
```
cd marketplace/neo4jtester
env GOOS=linux GOARCH=amd64 go build -o ../neo4jtester_linux main/main.go
```
In order to build this application for Mac OS
```
cd marketplace/neo4jtester
go build -o ../neo4jtester main/main.go
```

