package main

import (
	"context"
	"fmt"
	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("hello dgraph...")
	cxt := context.Background()
	d, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	defer d.Close()

	if err != nil {
		log.Fatal(err)
	}
	dgraph := dgo.NewDgraphClient(api.NewDgraphClient(d))
	txn := dgraph.NewReadOnlyTxn()
	defer txn.Discard(cxt)

	const q = `
		{
  			all_students (func:allofterms(nd_name,"student")){
    			uid
    			s_name
  			}
		}
	`
	resp, err := txn.Query(cxt, q)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.String())
}
