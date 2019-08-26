package main

import (
	"fmt"

    "github.com/nautilus/gateway"
    "github.com/nautilus/graphql"
)

// the schemas of each of our services
const schema1 = `
	type User {
		firstName: String!
		id: ID!
	}

	type Query {
		allUsers: [User!]!
	}
`
const schema2 = `
	type User {
		market: Market
	}

	type Market {
		name: String!
		logo: Image!
	}

	type Image! {
		id: ID!
	}
`
const schema3 = `
	type Image {
		thumbnailURL: String!
	}
`

// the query we are going to fire
const query = `
	query {
		allUsers {
			...UserList
		}
	}

	fragment UserList on Query {
		firstName
		market {
			...MarketDetails
		}
	}

	fragment MarketDetails on Market {
		name
		logo {
			thumbnailURL
		}
	}	
`


func main() {
	// build up a list of remote services
	remoteServices := []*graphql.RemoteSchema{}
	for i, schemaString := range []string{schema1, schema2, schema3} {
		// parse the schema
		schema, err := graphql.LoadSchema(schemaString)
		if err != nil {
			panic(err)
		}

		// add the schema to the list 
		remoteServices = append(remoteServices, &graphql.RemoteSchema{
			Schema: schema,
			URL: fmt.Sprintf("url-%i", i),
		})
	}

	// instantiate the gateway we'll time
	gw, err := gateway.New(remoteServices)
	if err != nil {
		panic(err)
	}


}