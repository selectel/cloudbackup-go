## Getting started

### Installation

You can install needed `cloudbackup-go` packages via `go get` command:

```bash
go get github.com/selectel/cloudbackup-go
```

### Authentication

To work with the Selectel Cloud Backup API you first need to:

* Create a Selectel account: [registration page](https://my.selectel.ru/registration).
* Create a project in Selectel Cloud Platform [projects](https://my.selectel.ru/vpc/projects).
* Retrieve a token for your project via API or [go-selvpcclient](https://github.com/selectel/go-selvpcclient).

### Endpoints

You can find available endpoints [here](https://docs.selectel.ru/en/api/urls/).

### Usage example

```go
package main

import (
	"context"
	"fmt"
	"log"

	cloudbackup "github.com/selectel/cloudbackup-go/pkg/v2"
)

func main() {
	// Token to work with Selectel Cloud project.
	token := "gAAAAABeVNzu-..."

	// Cloud backup endpoint to work with.
	endpoint := "https://ru-3.cloud.api.selcloud.ru/data-protect/v2/"

	// Create the client.
	client := cloudbackup.NewClientV2(token,endpoint)

	// Get the plans with the name "plan-name".
	plans, _,err := client.Plans(context.Background(), &cloudbackup.PlansQuery{Name: "plan-name"} )
	if err != nil {
		log.Fatal(err)
	}

	// Print the plans.
	for idx, plan := range plans {
		fmt.Printf("Plan %d: %+v", idx, plan)
	}
}

```