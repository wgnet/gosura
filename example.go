package gosura

import "log"

func example() {
	client := NewHasuraClient().
		URL("https://hasura.example.com").
		Endpoint("/location_behind_proxy/v1/query").
		SetAdminSecret("super_secret").
		SkipTLSVerify(true)

	query := NewRunSqlQuery()
	if err := query.SetArgs(RunSqlArgs{
		SQL: "select count(*) from public.posts",
	}); err != nil {
		panic(err)
	}

	data, err := client.Do(query)
	if err != nil {
		log.Fatal(err)
	}

	if data != nil {
		log.Printf("Result type: %s\nResult: %+v",
			data.(RunSqlResponse).ResultType,
			data.(RunSqlResponse).Result,
		)
	}
}
