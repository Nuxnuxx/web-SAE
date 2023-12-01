package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func main() {
	ctx := context.Background()

	driver, err := neo4j.NewDriverWithContext("neo4j://localhost:7687",
		neo4j.BasicAuth("neo4j", "Wawa02290", ""))

	defer driver.Close(ctx)

	if err != nil {
		panic(err)
	}

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Print("Connection established")

	session := driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})

	defer session.Close(ctx)

	e := echo.New()
	e.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")

		result, err := session.Run(ctx,
			"MATCH (n:Person) WHERE n.id = 4 RETURN n.name",
			map[string]interface{}{
				"id": id,
			},
		)

		if err != nil {
			panic(err)
		}

		if result.Next(ctx) {
			return c.JSON(http.StatusOK, result.Record())
		}
		return c.JSON(http.StatusOK, "saucisse")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
