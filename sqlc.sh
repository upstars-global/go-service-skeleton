#!/bin/bash

docker run --rm -v "$(pwd)":/src -w /src kjconroy/sqlc:1.11.0 generate

for FILE in `find internal/repositories/pgsql/requester -type f`
do
  BN=`basename $FILE`
  if [ "$BN" != "db.go" ] && [ "$BN" != "helpers.go" ]
  then
    cat $FILE | \
      sed 's/"database\/sql"/"github\.com\/gocraft\/dbr\/v2"/g' | \
      sed 's/database\.//g' | \
      sed 's/sql.NullBool/dbr.NullBool/g' | \
      sed 's/sql.NullString/dbr.NullString/g' | \
      sed 's/sql.NullInt32/dbr.NullInt64/g' | \
      sed 's/sql.NullInt64/dbr.NullInt64/g' | \
      sed 's/sql.NullTime/dbr.NullTime/g' \
      > $FILE.tmp
    mv $FILE.tmp $FILE
    go fmt $FILE
  fi
done

cat internal/repositories/pgsql/requester/querier.go | \
  sed 's/type Querier interface {/type Querier interface \{\n\tWithTx\(tx \*sql\.Tx\) *Queries/gi' > \
  internal/repositories/pgsql/requester/querier.go.tmp
mv internal/repositories/pgsql/requester/querier.go.tmp internal/repositories/pgsql/requester/querier.go

cat internal/repositories/pgsql/requester/querier.go | \
  sed 's/import (/import \(\n\t\"database\/sql\"/gi' > \
  internal/repositories/pgsql/requester/querier.go.tmp
mv internal/repositories/pgsql/requester/querier.go.tmp internal/repositories/pgsql/requester/querier.go

go fmt internal/repositories/pgsql/requester/querier.go