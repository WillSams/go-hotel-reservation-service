package specs

/*
- When a room is reserved, it cannot be reserved by another guest on overlapping dates.
- Whenever there are multiple available rooms for a request, the room with the lower final price is assigned.
- Whenever a request is made for a single room, a double bed room may be assigned (if no single is available?).
- Smokers are not placed in non-smoking rooms.
- Non-smokers are not placed in allowed smoking rooms.
- Final price for reservations are determined by daily price * num of days requested, plus the cleaning fee.
*/

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/willsams/go-hotel-reservation-service/api"
)

var _ = ginkgo.Describe("Reservations API", func() {
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

	ginkgo.Describe("CreateReservation", func() {
		ginkgo.It("assigns the room with the lower final price", func() {
			createReservation := api.CreateReservation(db) // Create two rooms with different daily prices

			// Create two rooms with different final prices
			room1ID := uuid.New().String()
			room1DailyPrice := 100.0
			room1CleaningFee := 10.0
			_, err := db.Exec("INSERT INTO rooms (id, daily_rate, cleaning_fee) VALUES ($1, $2, $3)", room1ID, room1DailyPrice, room1CleaningFee)
			gomega.Expect(err).ToNot(gomega.HaveOccurred())

			room2ID := uuid.New().String()
			room2DailyPrice := 50.0
			room2CleaningFee := 5.0
			_, err = db.Exec("INSERT INTO rooms (id, daily_rate, cleaning_fee) VALUES ($1, $2, $3)", room2ID, room2DailyPrice, room2CleaningFee)
			gomega.Expect(err).ToNot(gomega.HaveOccurred())

			// Create a reservation request with the checkin and checkout dates overlapping with the rooms
			checkinDate := "2023-01-01"
			checkoutDate := "2023-01-02"
			numOfDays := 2
			params := graphql.ResolveParams{
				Args: map[string]interface{}{
					"rommId":       "room-1",
					"checkinDate":  checkinDate,
					"checkoutDate": checkoutDate,
					"totalCharge":  100.00,
				},
			}

			// Call the CreateReservation function
			reservation, err := createReservation(params)
			gomega.Expect(err).ToNot(gomega.HaveOccurred())

			// Check that the reservation was assigned to the room with the lower final price
			room1FinalPrice := room1DailyPrice*float64(numOfDays) + room1CleaningFee
			room2FinalPrice := room2DailyPrice*float64(numOfDays) + room2CleaningFee
			if room1FinalPrice < room2FinalPrice {
				gomega.Expect(reservation.(api.Reservation).RoomID).To(gomega.Equal(room1ID))
			} else {
				gomega.Expect(reservation.(api.Reservation).RoomID).To(gomega.Equal(room2ID))
			}
		})

		ginkgo.It("prevents overlapping reservations for a room", func() {
			resolve := api.CreateReservation(db)
			params := graphql.ResolveParams{
				Args: map[string]interface{}{
					"roomId":       "room-1",
					"checkinDate":  "2022-06-01",
					"checkoutDate": "2022-06-02",
					"totalCharge":  100.00,
				},
			}

			_, err := resolve(params)
			gomega.Expect(err).To(gomega.BeNil())

			// Attempt to reserve the same room for overlapping dates
			params = graphql.ResolveParams{
				Args: map[string]interface{}{
					"roomId":       "room-1",
					"checkinDate":  "2022-06-02",
					"checkoutDate": "2022-06-03",
					"totalCharge":  100.00,
				},
			}

			_, err = resolve(params)
			gomega.Expect(err).ToNot(gomega.BeNil())
			gomega.Expect(err.Error()).To(gomega.Equal("room-1 is not available for the requested dates"))
		})
	})

})
