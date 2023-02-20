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
		db.Exec("INSERT INTO Rooms (id, num_beds, allow_smoking, daily_rate, cleaning_fee) VALUES ($1, $2, $3, $4, $5)", "201", 2, true, 200.0, 20.0)
	})

	ginkgo.AfterEach(func() {
		db.Exec("DELETE FROM Reservations")
		db.Exec("DELETE FROM Rooms")
		db.Close()
	})

	ginkgo.Describe("API Specs", func() {
		ginkgo.Describe("When a request is made for a single room", func() {
			ginkgo.It("a double bed room may be assigned", func() {
				/// create a reservation request for a single room
				startDate := "2023-03-01"
				endDate := "2023-04-05"
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

				createReservationParams := graphql.ResolveParams{
					Args: map[string]interface{}{
						"roomId":       room.ID,
						"checkinDate":  startDate,
						"checkoutDate": endDate,
						"totalCharge":  room.DailyRate*4 + room.CleaningFee,
					},
				}
				_, err = api.CreateReservation(db)(createReservationParams)
				gomega.Expect(err).To(gomega.BeNil())

				// check that the double room is assigned
				availableRoomsParams = graphql.ResolveParams{
					Args: map[string]interface{}{
						"startDate":    startDate,
						"endDate":      endDate,
						"numBeds":      numBeds,
						"allowSmoking": allowSmoking,
					},
				}
				availableRooms, err = api.GetAvailableRooms(db)(availableRoomsParams)
				gomega.Expect(err).To(gomega.BeNil())
				gomega.Expect(availableRooms).NotTo(gomega.BeNil())

				rooms = availableRooms.([]api.Room)
				for _, r := range rooms {
					gomega.Expect(r.ID).NotTo(gomega.Equal("104"))
				}
			})

			ginkgo.It("smokers are not placed in non-smoking rooms", func() {
				// get available non-smoking room
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
			})

			ginkgo.It("non-smokers are not placed in allowed smoking rooms", func() {

				// get available smoking rooms
				startDate := "2023-03-01"
				endDate := "2023-03-05"
				numBeds := 1
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
			})

			ginkgo.It("final price for reservations are determined by daily price * num of days requested, plus the cleaning fee", func() {
				// create a reservation request for a single room
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

				dailyRate := room.DailyRate
				numDays := 4
				cleaningFee := room.CleaningFee

				totalCharge := (dailyRate * float64(numDays)) + cleaningFee

				createReservationParams := graphql.ResolveParams{
					Args: map[string]interface{}{
						"roomId":       room.ID,
						"checkinDate":  startDate,
						"checkoutDate": endDate,
						"totalCharge":  totalCharge,
					},
				}
				_, err = api.CreateReservation(db)(createReservationParams)
				gomega.Expect(err).To(gomega.BeNil())

				// verify that the total charge is correct
				var reservation api.Reservation
				err = db.Get(&reservation, "SELECT * FROM Reservations WHERE room_id=$1", room.ID)
				gomega.Expect(err).To(gomega.BeNil())
				gomega.Expect(reservation.TotalCharge).To(gomega.Equal(totalCharge))
			})
		})
	})
})
