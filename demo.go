package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://zxc1998:zxc1998@cluster0.kfqzy.mongodb.net/Task?retryWrites=true&w=majority")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	//collection := client.Database("Task").Collection("parking")

var plot *pk.ParkingLot

func processString(s string) {
	vals := strings.Split(s, " ")
	switch vals[0] {
	case "create_parking_lot":
		n, err := strconv.Atoi(vals[1])
		if err != nil {
			fmt.Println("CREATE: Invalid value for number of slots.")
		} else {
			plot = pk.NewParkingLot(n)
		}
	case "park":
		plot.Park(vals[1], vals[2])
	case "leave":
		i, err := strconv.Atoi(vals[1])
		if err != nil {
			fmt.Println("LEAVE: Invalid value for slot number.")
		} else {
			plot.Leave(i)
		}
	case "status":
		plot.Status()
	case "registration_numbers_for_cars_with_colour":
		plot.GetRegNoOfColor(vals[1])
	case "slot_numbers_for_cars_with_colour":
		plot.GetSlotNoOfColor(vals[1])
	case "slot_number_for_registration_number":
		plot.GetSlotNoOfCar(vals[1])
	default:
		fmt.Println("ACTION: Invalid action")
	}
}

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")

}
