package specs

import (
	"fmt"
	"os"

	"github.com/graphql-go/graphql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/willsams/go-hotel-reservation-service/api"
)

var _ = ginkgo.Describe("GetAvailableRooms", func() {
	var db *sqlx.DB

	ginkgo.BeforeEach(func() {
		var err error

		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWD")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbName := "hotel_test"

		dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			dbHost, dbPort, dbUser, dbPassword, dbName)

		db, err = sqlx.Connect("postgres", dataSourceName)
		gomega.Expect(err).To(gomega.BeNil())
	})

	ginkgo.AfterEach(func() {
		err := db.Close()
		gomega.Expect(err).To(gomega.BeNil())
	})

	ginkgo.It("returns rooms that match the specified criteria", func() {
		resolve := api.GetAvailableRooms(db)
		params := graphql.ResolveParams{
			Args: map[string]interface{}{
				"startDate":    "2022-06-01",
				"endDate":      "2022-06-02",
				"numBeds":      2,
				"allowSmoking": false,
			},
		}

		rooms, err := resolve(params)
		gomega.Expect(err).To(gomega.BeNil())

		// Check that the returned rooms match the criteria
		for _, room := range rooms.([]api.Room) {
			gomega.Expect(room.NumBeds).To(gomega.Equal(2))
			gomega.Expect(room.AllowSmoking).To(gomega.BeFalse())
		}
	})
})
