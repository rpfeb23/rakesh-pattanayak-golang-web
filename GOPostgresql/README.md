`/Applications/Postgres.app/Contents/Versions/12/bin/psql -p5432 "postgres"` will connect to the Database `postgress`

    - Give your DB name instead of `postgres` .
    - If you dont have DB, just run `/Applications/Postgres.app/Contents/Versions/12/bin/psql -p5432`
    
 -  In this package instead of doing `go get github.com/lib/pq` I have added go.mod file to manage dependency
    -   since we are on GOPATH and I am in go 1.12 I had to do `export GO111MODULE=on` but in go 1.13 the  `GO111MODULE=on` is default.
    - then `go mod init .` meaning create my go.mod in current directory
    - Now run 'go run main.go' first time the imports not resolved will be extracted and put in the pkg
    
