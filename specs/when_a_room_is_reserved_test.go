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

var _ = ginkgo.Describe("Go Hotel Reservations Example", func() {
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

		// Add some available rooms to the AvailableRooms table
		db.Exec("INSERT INTO Rooms (id, num_beds, allow_smoking , daily_rate , cleaning_fee) VALUES ($1, $2, $3, $4, $5)", "101", 1, false, 100.0, 10.0)
		db.Exec("INSERT INTO Rooms (id, num_beds, allow_smoking , daily_rate , cleaning_fee) VALUES ($1, $2, $3, $4, $5)", "102", 1, false, 120.0, 10.0)
		db.Exec("INSERT INTO Rooms (id, num_beds, allow_smoking , daily_rate , cleaning_fee) VALUES ($1, $2, $3, $4, $5)", "103", 2, false, 150.0, 15.0)
		db.Exec("INSERT INTO Rooms (id, num_beds, allow_smoking, daily_rate, cleaning_fee) VALUES ($1, $2, $3, $4, $5)", "201", 2, true, 200.0, 20.0)
		db.Exec("INSERT INTO Rooms (id, num_beds, allow_smoking, daily_rate, cleaning_fee) VALUES ($1, $2, $3, $4, $5)", "202", 2, true, 150.0, 15.0)
		db.Exec("INSERT INTO Rooms (id, num_beds, allow_smoking, daily_rate, cleaning_fee) VALUES ($1, $2, $3, $4, $5)", "104", 1, false, 80.0, 8.0)
		db.Exec("INSERT INTO Rooms (id, num_beds, allow_smoking, daily_rate, cleaning_fee) VALUES ($1, $2, $3, $4, $5)", "105", 2, false, 120.0, 12.0)

		// Add some reservations to the Reservations table
		db.Exec("INSERT INTO Reservations (room_id, checkin_date, checkout_date, total_charge) VALUES ($1, $2, $3, $4)", "101", "2023-03-02", "2023-03-05", 330.0)
		db.Exec("INSERT INTO Reservations (room_id, checkin_date, checkout_date, total_charge) VALUES ($1, $2, $3, $4)", "102", "2023-03-05", "2023-03-08", 360.0)
	})

	ginkgo.AfterEach(func() {
		_, err := db.Exec("DELETE FROM Reservations")
		gomega.Expect(err).To(gomega.BeNil())

		db.Exec("DELETE FROM Rooms")
		gomega.Expect(err).To(gomega.BeNil())

		err = db.Close()
		gomega.Expect(err).To(gomega.BeNil())
	})

	ginkgo.Describe("When a room is reserved", func() {
		ginkgo.It("cannot be reserved by another guest on overlapping dates", func() {
			// Reserve a room for some dates
			startDate := "2023-03-01"
			endDate := "2023-03-05"
			numBeds := 1
			allowSmoking := false
			availableRoomsParams := graphql.ResolveParams{
				Args: map[string]interface{}{
					"startDate":    startDate,
					"endDate":      endDate,
					"numBeds":      numBeds,
					"allowSmoking": allowSmoking,
				},
			}
			availableRooms, err := api.GetAvailableRooms(db)(availableRoomsParams)
			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(availableRooms).NotTo(gomega.BeNil())

			rooms := availableRooms.([]api.Room)
			room := rooms[0]

			// Try to reserve the same room twice for overlapping dates
			overlappingStartDate := "2023-03-04"
			overlappingEndDate := "2023-03-08"
			createReservationParams := graphql.ResolveParams{
				Args: map[string]interface{}{
					"roomId":       room.ID,
					"checkinDate":  overlappingStartDate,
					"checkoutDate": overlappingEndDate,
					"totalCharge":  room.DailyRate*4 + room.CleaningFee,
				},
			}
			gomega.Expect(err).To(gomega.BeNil())

			api.CreateReservation(db)(createReservationParams)
			createReservationParams = graphql.ResolveParams{
				Args: map[string]interface{}{
					"roomId":       room.ID,
					"checkinDate":  overlappingStartDate,
					"checkoutDate": overlappingEndDate,
					"totalCharge":  room.DailyRate*4 + room.CleaningFee,
				},
			}

			_, err = api.CreateReservation(db)(createReservationParams)
			gomega.Expect(err).NotTo(gomega.BeNil())
		})
	})

	ginkgo.Describe("When there are multiple available rooms for a request", func() {
		ginkgo.It("the room with the lower final price is assigned", func() {
			// create a reservation request with the checkin and checkout dates overlapping with the rooms
			startDate := "2023-03-01"
			endDate := "2023-03-05"
			numBeds := 2
			allowSmoking := true
			availableRoomsParams := graphql.ResolveParams{
				Args: map[string]interface{}{
					"startDate":    startDate,
					"endDate":      endDate,
					"numBeds":      numBeds,
					"allowSmoking": allowSmoking,
				},
			}
			availableRooms, err := api.GetAvailableRooms(db)(availableRoomsParams)
			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(availableRooms).NotTo(gomega.BeNil())

			rooms := availableRooms.([]api.Room)

			// check that the room with the lower final price is assigned
			gomega.Expect(rooms[0].ID).To(gomega.Equal("202"))
		})
	})
})
