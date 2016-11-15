/*
Shim interface changed <https://github.com/hyperledger/fabric-chaintool/pull/25/files?diff=split>
Below code is based on latest chaincode shim. This might not work in IBM Bluemix
*/
//Redeploying
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/op/go-logging"
)

var myLogger = logging.MustGetLogger("safety_device")

// Chaincode type
type SafetyDeviceChaincode struct {
}

func main() {
	err := shim.Start(new(SafetyDeviceChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}

//Init Method***********************************************************************************************************
func (t *SafetyDeviceChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	myLogger.Debug("Init Chaincode...")

	var err error

	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments, Expecting 0")
	}

	// Create technician table
	err = stub.CreateTable("TechnicianDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "phone", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "name", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "email", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "deviceId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "deviceMake", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "deviceModel", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lastUpdatedOn", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating TechnicianDetails table.")
	}

	// Create device owner table
	err = stub.CreateTable("OwnerDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "phone", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "name", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "email", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "deviceId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "deviceMake", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "deviceModel", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "technicianId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lastUpdatedOn", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating OwnerDetails table.")
	}

	// Create Passenger details table
	err = stub.CreateTable("PassengerDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "phone", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "name", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "email", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "deviceId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "deviceMake", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "deviceModel", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lastUpdatedOn", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating PassengerDetails table.")
	}

	// Create Passenger's relative details table
	err = stub.CreateTable("PassengerRelativeDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "passengerPhone", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "Phone", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "name", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "isVIP", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lastUpdatedOn", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating PassengerRelativeDetails table.")
	}

	// Create Driver details table
	err = stub.CreateTable("DriverDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "phone", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "name", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "email", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "deviceId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "ownerId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "picLink", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "dlNumber", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "aadharId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lastUpdatedOn", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating DriverDetails table.")
	}

	// Create Vehicle details table
	err = stub.CreateTable("VehicleDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "registrationNumber", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "vehicleType", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "make", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "model", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "colour", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "ownerId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lastUpdatedOn", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating VehicleDetails table.")
	}

	// Create Panic Device details table
	err = stub.CreateTable("PanicDeviceDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "deviceId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "ownerId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "vehicleRegistrationNumber", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "deviceType", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "devicePosition", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lastUpdatedOn", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating PanicDeviceDetails table.")
	}

	// Create Panic Sequence details table
	err = stub.CreateTable("PanicSequenceDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "incidentId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "latitude", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "longitude", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "vehicleRegistrationNumber", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "videoLink", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "audioLink", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "passengerDeviceId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "panicDeviceId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lastUpdatedOn", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating PanicSequenceDetails table.")
	}

	myLogger.Debug("Init Chaincode...done")

	return nil, nil
}

//Panic Sequence Update Method*****************************************************************************************
func (t *SafetyDeviceChaincode) panicSequenceUpdate(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	myLogger.Debug("panicSequenceUpdate...")

	if len(args) != 8 {
		return nil, errors.New("Incorrect number of arguments. Expecting 8")
		//incidentId,latitude,longitude,vehicleRegistrationNumber,videoLink,audioLink,passengerDeviceId,panicDeviceId
	}

	incidentId := args[0]

	status, err := stub.ReplaceRow(
		"PanicSequenceDetails",
		shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: incidentId}},
				&shim.Column{Value: &shim.Column_String_{String_: args[1]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[2]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[3]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[4]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[5]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[6]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[7]}},
				&shim.Column{Value: &shim.Column_String_{String_: time.Now().UTC().String()}},
			},
		})
	if err != nil {
		return nil, errors.New("Failed replacing row.")
	}

	if status == false { //row doesn't exist
		_, err = stub.InsertRow(
			"PanicSequenceDetails",
			shim.Row{
				Columns: []*shim.Column{
					&shim.Column{Value: &shim.Column_String_{String_: incidentId}},
					&shim.Column{Value: &shim.Column_String_{String_: args[1]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[2]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[3]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[4]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[5]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[6]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[7]}},
					&shim.Column{Value: &shim.Column_String_{String_: time.Now().UTC().String()}},
				},
			})
		if err != nil {
			return nil, errors.New("Failed inserting row.")
		}
	}

	myLogger.Debug("panicSequenceUpdate...Done")
	return nil, nil

}

//Panic Device Update Method*****************************************************************************************
func (t *SafetyDeviceChaincode) panicDeviceUpdate(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	myLogger.Debug("panicDeviceUpdate...")

	if len(args) != 5 {
		return nil, errors.New("Incorrect number of arguments. Expecting 5")
		//deviceId,ownerId,vehicleRegistrationNumber,deviceType,devicePosition
	}

	deviceId := args[0]

	status, err := stub.ReplaceRow(
		"PanicDeviceDetails",
		shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: deviceId}},
				&shim.Column{Value: &shim.Column_String_{String_: args[1]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[2]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[3]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[4]}},
				&shim.Column{Value: &shim.Column_String_{String_: time.Now().UTC().String()}},
			},
		})
	if err != nil {
		return nil, errors.New("Failed replacing row.")
	}

	if status == false { //row doesn't exist
		_, err = stub.InsertRow(
			"PanicDeviceDetails",
			shim.Row{
				Columns: []*shim.Column{
					&shim.Column{Value: &shim.Column_String_{String_: deviceId}},
					&shim.Column{Value: &shim.Column_String_{String_: args[1]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[2]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[3]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[4]}},
					&shim.Column{Value: &shim.Column_String_{String_: time.Now().UTC().String()}},
				},
			})
		if err != nil {
			return nil, errors.New("Failed inserting row.")
		}
	}

	myLogger.Debug("panicDeviceUpdate...Done")
	return nil, nil

}

//Vehicle Update Method*****************************************************************************************
func (t *SafetyDeviceChaincode) vehicleUpdate(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	myLogger.Debug("vehicleUpdate...")

	if len(args) != 6 {
		return nil, errors.New("Incorrect number of arguments. Expecting 6")
		//registrationNumber,vehicleType,make,model,colour,ownerId
	}

	registrationNumber := args[0]

	status, err := stub.ReplaceRow(
		"VehicleDetails",
		shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: registrationNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: args[1]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[2]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[3]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[4]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[5]}},
				&shim.Column{Value: &shim.Column_String_{String_: time.Now().UTC().String()}},
			},
		})
	if err != nil {
		return nil, errors.New("Failed replacing row.")
	}

	if status == false { //row doesn't exist
		_, err = stub.InsertRow(
			"VehicleDetails",
			shim.Row{
				Columns: []*shim.Column{
					&shim.Column{Value: &shim.Column_String_{String_: registrationNumber}},
					&shim.Column{Value: &shim.Column_String_{String_: args[1]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[2]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[3]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[4]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[5]}},
					&shim.Column{Value: &shim.Column_String_{String_: time.Now().UTC().String()}},
				},
			})
		if err != nil {
			return nil, errors.New("Failed inserting row.")
		}
	}

	myLogger.Debug("vehicleUpdate...Done")
	return nil, nil

}

//Driver Update Method*****************************************************************************************
//NOTE: Needs to be throughly tested. update by dlNumber or aadharId
func (t *SafetyDeviceChaincode) driverUpdate(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	myLogger.Debug("driverUpdate...")

	if len(args) != 8 {
		return nil, errors.New("Incorrect number of arguments. Expecting 8")
		//phone,name,email,deviceId,ownerId,picLink,dlNumber,aadharId
	}

	phone := args[0]

	status, err := stub.ReplaceRow(
		"DriverDetails",
		shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: phone}},
				&shim.Column{Value: &shim.Column_String_{String_: args[1]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[2]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[3]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[4]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[5]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[6]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[7]}},
				&shim.Column{Value: &shim.Column_String_{String_: time.Now().UTC().String()}},
			},
		})
	if err != nil {
		return nil, errors.New("Failed replacing row.")
	}

	if status == false { //row doesn't exist
		_, err = stub.InsertRow(
			"DriverDetails",
			shim.Row{
				Columns: []*shim.Column{
					&shim.Column{Value: &shim.Column_String_{String_: phone}},
					&shim.Column{Value: &shim.Column_String_{String_: args[1]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[2]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[3]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[4]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[5]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[6]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[7]}},
					&shim.Column{Value: &shim.Column_String_{String_: time.Now().UTC().String()}},
				},
			})
		if err != nil {
			return nil, errors.New("Failed inserting row.")
		}
	}

	myLogger.Debug("driverUpdate...Done")
	return nil, nil

}

//Passenger Update Method*****************************************************************************************
func (t *SafetyDeviceChaincode) passengerUpdate(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	myLogger.Debug("passengerUpdate...")

	if len(args) != 6 {
		return nil, errors.New("Incorrect number of arguments. Expecting 6")
		//phone,name,email,deviceId,deviceMake,deviceModel
	}

	phone := args[0]

	status, err := stub.ReplaceRow(
		"PassengerDetails",
		shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: phone}},
				&shim.Column{Value: &shim.Column_String_{String_: args[1]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[2]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[3]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[4]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[5]}},
				&shim.Column{Value: &shim.Column_String_{String_: time.Now().UTC().String()}},
			},
		})
	if err != nil {
		return nil, errors.New("Failed replacing row.")
	}

	if status == false { //row doesn't exist
		_, err = stub.InsertRow(
			"PassengerDetails",
			shim.Row{
				Columns: []*shim.Column{
					&shim.Column{Value: &shim.Column_String_{String_: phone}},
					&shim.Column{Value: &shim.Column_String_{String_: args[1]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[2]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[3]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[4]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[5]}},
					&shim.Column{Value: &shim.Column_String_{String_: time.Now().UTC().String()}},
				},
			})
		if err != nil {
			return nil, errors.New("Failed inserting row.")
		}
	}

	myLogger.Debug("passengerUpdate...Done")
	return nil, nil

}

//Passenger's Relative Update Method*****************************************************************************************
func (t *SafetyDeviceChaincode) passengerRelativeUpdate(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	myLogger.Debug("passengerRelativeUpdate...")

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
		//passengerPhone,phone,name,isVIP
	}

	passengerPhone := args[0]

	status, err := stub.ReplaceRow(
		"PassengerRelativeDetails",
		shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: passengerPhone}},
				&shim.Column{Value: &shim.Column_String_{String_: args[1]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[2]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[3]}},
				&shim.Column{Value: &shim.Column_String_{String_: time.Now().UTC().String()}},
			},
		})
	if err != nil {
		return nil, errors.New("Failed replacing row.")
	}

	if status == false { //row doesn't exist
		_, err = stub.InsertRow(
			"PassengerRelativeDetails",
			shim.Row{
				Columns: []*shim.Column{
					&shim.Column{Value: &shim.Column_String_{String_: passengerPhone}},
					&shim.Column{Value: &shim.Column_String_{String_: args[1]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[2]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[3]}},
					&shim.Column{Value: &shim.Column_String_{String_: time.Now().UTC().String()}},
				},
			})
		if err != nil {
			return nil, errors.New("Failed inserting row.")
		}
	}

	myLogger.Debug("passengerRelativeUpdate...Done")
	return nil, nil

}

//Technician Update Method***********************************************************************************************
func (t *SafetyDeviceChaincode) technicianUpdate(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	myLogger.Debug("technicianUpdate...")

	if len(args) != 6 {
		return nil, errors.New("Incorrect number of arguments. Expecting 6")
		//phone,name,email,deviceId,deviceMake,deviceModel
	}

	phone := args[0]

	status, err := stub.ReplaceRow(
		"TechnicianDetails",
		shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: phone}},
				&shim.Column{Value: &shim.Column_String_{String_: args[1]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[2]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[3]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[4]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[5]}},
				&shim.Column{Value: &shim.Column_String_{String_: time.Now().UTC().String()}},
			},
		})
	if err != nil {
		return nil, errors.New("Failed replacing row.")
	}

	if status == false { //row doesn't exist
		_, err = stub.InsertRow(
			"TechnicianDetails",
			shim.Row{
				Columns: []*shim.Column{
					&shim.Column{Value: &shim.Column_String_{String_: phone}},
					&shim.Column{Value: &shim.Column_String_{String_: args[1]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[2]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[3]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[4]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[5]}},
					&shim.Column{Value: &shim.Column_String_{String_: time.Now().UTC().String()}},
				},
			})
		if err != nil {
			return nil, errors.New("Failed inserting row.")
		}
	}

	myLogger.Debug("technicianUpdate...Done")
	return nil, nil

}

//Owner Update Method*****************************************************************************************
func (t *SafetyDeviceChaincode) ownerUpdate(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	myLogger.Debug("ownerUpdate...")

	if len(args) != 7 {
		return nil, errors.New("Incorrect number of arguments. Expecting 7")
		//phone,name,email,deviceId,deviceMake,deviceModel,technicianId
	}

	phone := args[0]

	status, err := stub.ReplaceRow(
		"OwnerDetails",
		shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: phone}},
				&shim.Column{Value: &shim.Column_String_{String_: args[1]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[2]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[3]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[4]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[5]}},
				&shim.Column{Value: &shim.Column_String_{String_: args[6]}},
				&shim.Column{Value: &shim.Column_String_{String_: time.Now().UTC().String()}},
			},
		})
	if err != nil {
		return nil, errors.New("Failed replacing row.")
	}

	if status == false { //row doesn't exist
		_, err = stub.InsertRow(
			"OwnerDetails",
			shim.Row{
				Columns: []*shim.Column{
					&shim.Column{Value: &shim.Column_String_{String_: phone}},
					&shim.Column{Value: &shim.Column_String_{String_: args[1]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[2]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[3]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[4]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[5]}},
					&shim.Column{Value: &shim.Column_String_{String_: args[6]}},
					&shim.Column{Value: &shim.Column_String_{String_: time.Now().UTC().String()}},
				},
			})
		if err != nil {
			return nil, errors.New("Failed inserting row.")
		}
	}

	myLogger.Debug("ownerUpdate...Done")
	return nil, nil

}

//Invoke Method***********************************************************************************************************
func (t *SafetyDeviceChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	myLogger.Debug("Invoke Chaincode...started")

	// Handle different functions
	if function == "technicianUpdate" {
		return t.technicianUpdate(stub, args)
	} else if function == "ownerUpdate" {
		return t.ownerUpdate(stub, args)
	} else if function == "passengerUpdate" {
		return t.passengerUpdate(stub, args)
	} else if function == "passengerRelativeUpdate" {
		return t.passengerRelativeUpdate(stub, args)
	} else if function == "driverUpdate" {
		return t.driverUpdate(stub, args)
	} else if function == "vehicleUpdate" {
		return t.vehicleUpdate(stub, args)
	} else if function == "panicDeviceUpdate" {
		return t.panicDeviceUpdate(stub, args)
	} else if function == "panicSequenceUpdate" {
		return t.panicSequenceUpdate(stub, args)
	}

	return nil, errors.New("Received unknown function invocation")
}

//Owner Query Method*****************************************************************************************
func (t *SafetyDeviceChaincode) ownerQuery(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	myLogger.Debug("ownerQuery...")

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
		//phone
	}

	phone := args[0]

	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: phone}}
	columns = append(columns, col1)
	row, err := stub.GetRow("OwnerDetails", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed retrieving owner details [%s]: [%s]", phone, err)
	}

	//contentsBytes, err := json.Marshal(row)
	//if err != nil {
	//	return nil, fmt.Errorf("Failed marshal all contents: %v", err)
	//}
	//fmt.Println(string(contentsBytes))

	var m map[string]string
	var mapBytes []byte
	m = make(map[string]string)

	if len(row.Columns) == 0 {
		mapBytes, err = json.Marshal(m)
		if err != nil {
			return nil, fmt.Errorf("Failed marshal all map contents: %v", err)
		}
		fmt.Println(string(mapBytes))
	} else {

		m["phone"] = row.Columns[0].GetString_()
		m["name"] = row.Columns[1].GetString_()
		m["email"] = row.Columns[2].GetString_()
		m["deviceId"] = row.Columns[3].GetString_()
		m["deviceMake"] = row.Columns[4].GetString_()
		m["deviceModel"] = row.Columns[5].GetString_()
		m["technicianId"] = row.Columns[6].GetString_()
		m["lastUpdatedOn"] = row.Columns[7].GetString_()

		mapBytes, err = json.Marshal(m)
		if err != nil {
			return nil, fmt.Errorf("Failed marshal all map contents: %v", err)
		}
		fmt.Println(string(mapBytes))
	}

	myLogger.Debug("ownerQuery...DONE")
	return mapBytes, nil

}

//Technician Query Method*****************************************************************************************
func (t *SafetyDeviceChaincode) technicianQuery(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	myLogger.Debug("technicianQuery...")

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
		//phone
	}

	phone := args[0]

	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: phone}}
	columns = append(columns, col1)
	row, err := stub.GetRow("TechnicianDetails", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed retrieving technician details [%s]: [%s]", phone, err)
	}

	//contentsBytes, err := json.Marshal(row)
	//if err != nil {
	//return nil, fmt.Errorf("Failed marshal all contents: %v", err)
	//}
	//fmt.Println(string(contentsBytes))

	var m map[string]string
	var mapBytes []byte
	m = make(map[string]string)

	if len(row.Columns) == 0 {
		mapBytes, err = json.Marshal(m)
		if err != nil {
			return nil, fmt.Errorf("Failed marshal all map contents: %v", err)
		}
		fmt.Println(string(mapBytes))
	} else {

		m["phone"] = row.Columns[0].GetString_()
		m["name"] = row.Columns[1].GetString_()
		m["email"] = row.Columns[2].GetString_()
		m["deviceId"] = row.Columns[3].GetString_()
		m["deviceMake"] = row.Columns[4].GetString_()
		m["deviceModel"] = row.Columns[5].GetString_()
		m["lastUpdatedOn"] = row.Columns[6].GetString_()

		mapBytes, err = json.Marshal(m)
		if err != nil {
			return nil, fmt.Errorf("Failed marshal all map contents: %v", err)
		}
		fmt.Println(string(mapBytes))
	}

	myLogger.Debug("technicianQuery...DONE")
	return mapBytes, nil

}

//Passenger Query Method*****************************************************************************************
func (t *SafetyDeviceChaincode) passengerQuery(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	myLogger.Debug("passengerQuery...")

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
		//phone
	}

	phone := args[0]

	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: phone}}
	columns = append(columns, col1)
	row, err := stub.GetRow("PassengerDetails", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed retrieving passenger details [%s]: [%s]", phone, err)
	}

	//contentsBytes, err := json.Marshal(row)
	//if err != nil {
	//return nil, fmt.Errorf("Failed marshal all contents: %v", err)
	//}
	//fmt.Println(string(contentsBytes))

	var m map[string]string
	var mapBytes []byte
	m = make(map[string]string)

	if len(row.Columns) == 0 {
		mapBytes, err = json.Marshal(m)
		if err != nil {
			return nil, fmt.Errorf("Failed marshal all map contents: %v", err)
		}
		fmt.Println(string(mapBytes))
	} else {

		m["phone"] = row.Columns[0].GetString_()
		m["name"] = row.Columns[1].GetString_()
		m["email"] = row.Columns[2].GetString_()
		m["deviceId"] = row.Columns[3].GetString_()
		m["deviceMake"] = row.Columns[4].GetString_()
		m["deviceModel"] = row.Columns[5].GetString_()
		m["lastUpdatedOn"] = row.Columns[6].GetString_()

		mapBytes, err = json.Marshal(m)
		if err != nil {
			return nil, fmt.Errorf("Failed marshal all map contents: %v", err)
		}
		fmt.Println(string(mapBytes))
	}

	myLogger.Debug("passengerQuery...DONE")
	return mapBytes, nil

}

//Passenger's Relative Query Method*****************************************************************************************
func (t *SafetyDeviceChaincode) passengerRelativeQuery(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	myLogger.Debug("passengerRelativeQuery...")

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
		//passengerPhone
	}

	passengerPhone := args[0]

	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: passengerPhone}}
	columns = append(columns, col1)
	row, err := stub.GetRow("PassengerRelativeDetails", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed retrieving passenger details [%s]: [%s]", passengerPhone, err)
	}

	//contentsBytes, err := json.Marshal(row)
	//if err != nil {
	//return nil, fmt.Errorf("Failed marshal all contents: %v", err)
	//}
	//fmt.Println(string(contentsBytes))

	var m map[string]string
	var mapBytes []byte
	m = make(map[string]string)

	if len(row.Columns) == 0 {
		mapBytes, err = json.Marshal(m)
		if err != nil {
			return nil, fmt.Errorf("Failed marshal all map contents: %v", err)
		}
		fmt.Println(string(mapBytes))
	} else {

		m["passengerPhone"] = row.Columns[0].GetString_()
		m["phone"] = row.Columns[1].GetString_()
		m["name"] = row.Columns[2].GetString_()
		m["isVIP"] = row.Columns[3].GetString_()
		m["lastUpdatedOn"] = row.Columns[4].GetString_()

		mapBytes, err = json.Marshal(m)
		if err != nil {
			return nil, fmt.Errorf("Failed marshal all map contents: %v", err)
		}
		fmt.Println(string(mapBytes))
	}

	myLogger.Debug("passengerRelativeQuery...DONE")
	return mapBytes, nil

}

//Driver Query Method*****************************************************************************************
func (t *SafetyDeviceChaincode) driverQuery(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	myLogger.Debug("driverQuery...")

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
		//phone
	}

	phone := args[0]

	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: phone}}
	columns = append(columns, col1)
	row, err := stub.GetRow("DriverDetails", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed retrieving driver details [%s]: [%s]", phone, err)
	}

	//contentsBytes, err := json.Marshal(row)
	//if err != nil {
	//return nil, fmt.Errorf("Failed marshal all contents: %v", err)
	//}
	//fmt.Println(string(contentsBytes))

	var m map[string]string
	var mapBytes []byte
	m = make(map[string]string)

	if len(row.Columns) == 0 {
		mapBytes, err = json.Marshal(m)
		if err != nil {
			return nil, fmt.Errorf("Failed marshal all map contents: %v", err)
		}
		fmt.Println(string(mapBytes))
	} else {

		m["phone"] = row.Columns[0].GetString_()
		m["name"] = row.Columns[1].GetString_()
		m["email"] = row.Columns[2].GetString_()
		m["deviceId"] = row.Columns[3].GetString_()
		m["ownerId"] = row.Columns[4].GetString_()
		m["picLink"] = row.Columns[5].GetString_()
		m["dlNumber"] = row.Columns[6].GetString_()
		m["aadharId"] = row.Columns[7].GetString_()
		m["lastUpdatedOn"] = row.Columns[8].GetString_()

		mapBytes, err = json.Marshal(m)
		if err != nil {
			return nil, fmt.Errorf("Failed marshal all map contents: %v", err)
		}
		fmt.Println(string(mapBytes))
	}

	myLogger.Debug("driverQuery...DONE")
	return mapBytes, nil

}

//Vehicle Query Method*****************************************************************************************
func (t *SafetyDeviceChaincode) vehicleQuery(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	myLogger.Debug("vehicleQuery...")

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
		//registrationNumber
	}

	registrationNumber := args[0]

	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: registrationNumber}}
	columns = append(columns, col1)
	row, err := stub.GetRow("VehicleDetails", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed retrieving vehicle details [%s]: [%s]", registrationNumber, err)
	}

	//contentsBytes, err := json.Marshal(row)
	//if err != nil {
	//return nil, fmt.Errorf("Failed marshal all contents: %v", err)
	//}
	//fmt.Println(string(contentsBytes))

	var m map[string]string
	var mapBytes []byte
	m = make(map[string]string)

	if len(row.Columns) == 0 {
		mapBytes, err = json.Marshal(m)
		if err != nil {
			return nil, fmt.Errorf("Failed marshal all map contents: %v", err)
		}
		fmt.Println(string(mapBytes))
	} else {

		m["registrationNumber"] = row.Columns[0].GetString_()
		m["vehicleType"] = row.Columns[1].GetString_()
		m["make"] = row.Columns[2].GetString_()
		m["model"] = row.Columns[3].GetString_()
		m["colour"] = row.Columns[4].GetString_()
		m["ownerId"] = row.Columns[5].GetString_()
		m["lastUpdatedOn"] = row.Columns[6].GetString_()

		mapBytes, err = json.Marshal(m)
		if err != nil {
			return nil, fmt.Errorf("Failed marshal all map contents: %v", err)
		}
		fmt.Println(string(mapBytes))
	}

	myLogger.Debug("vehicleQuery...DONE")
	return mapBytes, nil

}

//Panic Device Query Method*****************************************************************************************
func (t *SafetyDeviceChaincode) panicDeviceQuery(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	myLogger.Debug("panicDeviceQuery...")

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
		//deviceId
	}

	deviceId := args[0]

	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: deviceId}}
	columns = append(columns, col1)
	row, err := stub.GetRow("PanicDeviceDetails", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed retrieving panic device details [%s]: [%s]", deviceId, err)
	}

	//contentsBytes, err := json.Marshal(row)
	//if err != nil {
	//return nil, fmt.Errorf("Failed marshal all contents: %v", err)
	//}
	//fmt.Println(string(contentsBytes))

	var m map[string]string
	var mapBytes []byte
	m = make(map[string]string)

	if len(row.Columns) == 0 {
		mapBytes, err = json.Marshal(m)
		if err != nil {
			return nil, fmt.Errorf("Failed marshal all map contents: %v", err)
		}
		fmt.Println(string(mapBytes))
	} else {

		m["deviceId"] = row.Columns[0].GetString_()
		m["ownerId"] = row.Columns[1].GetString_()
		m["vehicleRegistrationNumber"] = row.Columns[2].GetString_()
		m["deviceType"] = row.Columns[3].GetString_()
		m["devicePosition"] = row.Columns[4].GetString_()
		m["lastUpdatedOn"] = row.Columns[5].GetString_()

		mapBytes, err = json.Marshal(m)
		if err != nil {
			return nil, fmt.Errorf("Failed marshal all map contents: %v", err)
		}
		fmt.Println(string(mapBytes))
	}

	myLogger.Debug("panicDeviceQuery...DONE")
	return mapBytes, nil

}

//Panic Sequence Query Method*****************************************************************************************
func (t *SafetyDeviceChaincode) panicSequenceQuery(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	myLogger.Debug("panicSequenceQuery...")

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
		//incidentId : -1 for query_all
	}

	if args[0] == "-1" {
		var columns []shim.Column
		rowChannel, err := stub.GetRows("PanicSequenceDetails", columns)
		if err != nil {
			return nil, fmt.Errorf("Failed retrieving all panic sequences: [%s]", err)
		}

		// 2. Fetch all the rows
		var rows []*shim.Row
		for {
			select {
			case row, ok := <-rowChannel:
				if !ok {
					rowChannel = nil
				} else {
					rows = append(rows, &row)
				}
			}

			if rowChannel == nil {
				break
			}
		}

		var sliceBytes []byte
		var m map[string]string
		var s []map[string]string

		m = make(map[string]string)
		if len(rows) == 0 {
			sliceBytes, err = json.Marshal(m)
			if err != nil {
				return nil, fmt.Errorf("Failed marshal all slice contents: %v", err)
			}
			fmt.Println(string(sliceBytes))
		} else {
			s = make([]map[string]string, len(rows))
			for i := 0; i < len(rows); i++ {
				m = make(map[string]string)
				m["incidentId"] = rows[i].Columns[0].GetString_()
				m["latitude"] = rows[i].Columns[1].GetString_()
				m["longitude"] = rows[i].Columns[2].GetString_()
				m["vehicleRegistrationNumber"] = rows[i].Columns[3].GetString_()
				m["videoLink"] = rows[i].Columns[4].GetString_()
				m["audioLink"] = rows[i].Columns[5].GetString_()
				m["passengerDeviceId"] = rows[i].Columns[6].GetString_()
				m["panicDeviceId"] = rows[i].Columns[7].GetString_()
				m["lastUpdatedOn"] = rows[i].Columns[8].GetString_()

				//s = append(s, m)
				s[i] = m

			}
			sliceBytes, err = json.Marshal(s)
			if err != nil {
				return nil, fmt.Errorf("Failed marshal all slice contents: %v", err)
			}
			fmt.Println(string(sliceBytes))
		}

		myLogger.Debug("panicSequenceQuery...DONE")
		return sliceBytes, nil

	} else {
		incidentId := args[0]

		var columns []shim.Column
		col1 := shim.Column{Value: &shim.Column_String_{String_: incidentId}}
		columns = append(columns, col1)
		row, err := stub.GetRow("PanicSequenceDetails", columns)
		if err != nil {
			return nil, fmt.Errorf("Failed retrieving panic sequence details [%s]: [%s]", incidentId, err)
		}

		//contentsBytes, err := json.Marshal(row)
		//if err != nil {
		//return nil, fmt.Errorf("Failed marshal all contents: %v", err)
		//}
		//fmt.Println(string(contentsBytes))

		var m map[string]string
		var mapBytes []byte
		m = make(map[string]string)

		if len(row.Columns) == 0 {
			mapBytes, err = json.Marshal(m)
			if err != nil {
				return nil, fmt.Errorf("Failed marshal all map contents: %v", err)
			}
			fmt.Println(string(mapBytes))
		} else {

			m["incidentId"] = row.Columns[0].GetString_()
			m["latitude"] = row.Columns[1].GetString_()
			m["longitude"] = row.Columns[2].GetString_()
			m["vehicleRegistrationNumber"] = row.Columns[3].GetString_()
			m["videoLink"] = row.Columns[4].GetString_()
			m["audioLink"] = row.Columns[5].GetString_()
			m["passengerDeviceId"] = row.Columns[6].GetString_()
			m["panicDeviceId"] = row.Columns[7].GetString_()
			m["lastUpdatedOn"] = row.Columns[8].GetString_()

			mapBytes, err = json.Marshal(m)
			if err != nil {
				return nil, fmt.Errorf("Failed marshal all map contents: %v", err)
			}
			fmt.Println(string(mapBytes))
		}
		myLogger.Debug("panicSequenceQuery...DONE")
		return mapBytes, nil
	}

}

//Query Method***********************************************************************************************************
func (t *SafetyDeviceChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	myLogger.Debug("Query Chaincode...started")

	// Handle different functions
	if function == "technicianQuery" {
		return t.technicianQuery(stub, args)
	} else if function == "ownerQuery" {
		return t.ownerQuery(stub, args)
	} else if function == "passengerQuery" {
		return t.passengerQuery(stub, args)
	} else if function == "passengerRelativeQuery" {
		return t.passengerRelativeQuery(stub, args)
	} else if function == "driverQuery" {
		return t.driverQuery(stub, args)
	} else if function == "vehicleQuery" {
		return t.vehicleQuery(stub, args)
	} else if function == "panicDeviceQuery" {
		return t.panicDeviceQuery(stub, args)
	} else if function == "panicSequenceQuery" {
		return t.panicSequenceQuery(stub, args)
	}

	myLogger.Debug("Query Chaincode...DONE")
	return nil, errors.New("Received unknown function invocation")
}
