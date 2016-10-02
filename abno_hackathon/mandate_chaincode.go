
package main

import (
	"encoding/base64"
	"errors"
  "fmt"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/op/go-logging"
)

var myLogger = logging.MustGetLogger("mandate_mgm")

type MandateChaincode struct {
}

func main() {
	//primitives.SetSecurityLevel("SHA3", 256)
	err := shim.Start(new(MandateChaincode))
	if err != nil {
		fmt.Printf("Error starting MandateChaincode: %s", err)
	}
}

func (t *MandateChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	myLogger.Debug("Init Chaincode...")
	if len(args) != 5 {
		return nil, errors.New("Incorrect number of arguments. Expecting 5")
	}

  var err error
  var myTcert,cust1Tcert,cust2Tcert,supplierTcert,gridOperatorTcert []byte
  //store peer certificate as myCert. Each peer node can use myCert to invoke and query the chaincode
  myTcert, err = base64.StdEncoding.DecodeString(args[0])
  if err != nil {
		return nil, errors.New("Failed decoding peer certificate")
	}

	//store the peer node's certificate
  stub.PutState("myTcert", myTcert)

	//store the Auditor's certificate (EDSN)
	edsnTcert,err := base64.StdEncoding.DecodeString("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNVekNDQWZpZ0F3SUJBZ0lSQVBkc2hLOHpRVTU3dTFMcjNJTk5NUUF3Q2dZSUtvWkl6ajBFQXdNd01URUwKTUFrR0ExVUVCaE1DVlZNeEZEQVNCZ05WQkFvVEMwaDVjR1Z5YkdWa1oyVnlNUXd3Q2dZRFZRUURFd04wWTJFdwpIaGNOTVRZeE1EQXhNVGN6T0RRMFdoY05NVFl4TWpNd01UY3pPRFEwV2pCRk1Rc3dDUVlEVlFRR0V3SlZVekVVCk1CSUdBMVVFQ2hNTFNIbHdaWEpzWldSblpYSXhJREFlQmdOVkJBTVRGMVJ5WVc1ellXTjBhVzl1SUVObGNuUnAKWm1sallYUmxNRmt3RXdZSEtvWkl6ajBDQVFZSUtvWkl6ajBEQVFjRFFnQUUrck0wazQwMmFsOUI5VlhiS2V4VApKbTZHL0Y5WHdPM0VVYjRRRUJMNkFvMWV2OE1FZ2orMFhvMDREL0lzeGpJa0M1V1I4RFNRbW42ZUxRa2hyQXJPCjNhT0IzRENCMlRBT0JnTlZIUThCQWY4RUJBTUNCNEF3REFZRFZSMFRBUUgvQkFJd0FEQU5CZ05WSFE0RUJnUUUKQVFJREJEQVBCZ05WSFNNRUNEQUdnQVFCQWdNRU1FMEdCaW9EQkFVR0J3RUIvd1JBWURtMkV2a0cxMDBodEJ5agpDSktyNEtsdDZkMU5lRU0wNmQ2UjlmZURFQ1JZcXlHYk5ybitZMWorOXBRK0hYaFBjeVRuSmR0Y0xpQURwREZNCmg0bTl2REJLQmdZcUF3UUZCZ2dFUUx3dUF5UnhHbU5udDkya1Z3bUZDcUZsamFERURCcitjdFVTNHZ5eU1rVkQKK0MzZ0gvMDNhVXM0ZHdKQU0ydHJRMUk4a0RLL2VOcTVYcTk2aTZMdEFlc3dDZ1lJS29aSXpqMEVBd01EU1FBdwpSZ0loQUp1NFlYTE8raWp2VTJ3TkM4RXh0N3B3T3d2c2l2QTZ6a1RXZ2dlNjFvdG9BaUVBam13UGs5cHJXcEpZCnlvYVVkQjlUdnhQTGNxYkR5cVNsSU1kQjZMYzNuZWs9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K")
	if err != nil {
		return nil, errors.New("Failed decoding edsn certificate")
	}
	stub.PutState("edsnTcert",edsnTcert)

	// Create CAR table
	err = stub.CreateTable("CAR", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "custId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "custTcert", Type: shim.ColumnDefinition_BYTES, Key: false},
    &shim.ColumnDefinition{Name: "custMandate", Type: shim.ColumnDefinition_BYTES, Key: false},
    &shim.ColumnDefinition{Name: "supplierTcert", Type: shim.ColumnDefinition_BYTES, Key: false},
    &shim.ColumnDefinition{Name: "gridOperatorTcert", Type: shim.ColumnDefinition_BYTES, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating CAR table.")
	}

  //update two dummy customer
  custMandate := "{'mandate':[{'op1':'no','startDate':'10/10/10','endDate':'10/10/10'},{'op2':'no','startDate':'10/10/10','endDate':'10/10/10'},{'op3':'no','startDate':'10/10/10','endDate':'10/10/10'}]}"

  cust1Tcert, err = base64.StdEncoding.DecodeString(args[1])
  if err != nil {
		return nil, errors.New("Failed decoding customer 1 certificate")
	}

  cust2Tcert, err = base64.StdEncoding.DecodeString(args[2])
  if err != nil {
		return nil, errors.New("Failed decoding customer 2 certificate")
	}

  supplierTcert, err = base64.StdEncoding.DecodeString(args[3])
  if err != nil {
		return nil, errors.New("Failed decoding supplier certificate")
	}

  gridOperatorTcert, err = base64.StdEncoding.DecodeString(args[4])
  if err != nil {
		return nil, errors.New("Failed decoding grid operator certificate")
	}

	custId := "1"
  ok, err := stub.InsertRow("CAR", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: custId}},
			&shim.Column{Value: &shim.Column_Bytes{Bytes: cust1Tcert}},
      &shim.Column{Value: &shim.Column_Bytes{Bytes: []byte(custMandate)}},
      &shim.Column{Value: &shim.Column_Bytes{Bytes: supplierTcert}},
      &shim.Column{Value: &shim.Column_Bytes{Bytes: gridOperatorTcert}}},
	})

	if err != nil{
		return nil, err
	}
	if !ok && err == nil {
		return nil, errors.New("Unable to insert database.")
	}

	custId = "2"
  ok, err = stub.InsertRow("CAR", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: custId}},
			&shim.Column{Value: &shim.Column_Bytes{Bytes: cust2Tcert}},
      &shim.Column{Value: &shim.Column_Bytes{Bytes: []byte(custMandate)}},
      &shim.Column{Value: &shim.Column_Bytes{Bytes: supplierTcert}},
      &shim.Column{Value: &shim.Column_Bytes{Bytes: gridOperatorTcert}}},
	})

	if err != nil{
		return nil, err
	}
	if !ok && err == nil {
		return nil, errors.New("Unable to insert database.")
	}

	myLogger.Debug("Init...done!")

  return nil,err

}

func (t *MandateChaincode) updateMandate(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	myLogger.Debug("updateMandate...")

	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3")
	}

  custId := args[0]

  var columns []shim.Column
  col1 := shim.Column{Value: &shim.Column_String_{String_: custId}}
  columns = append(columns, col1)

  row, err := stub.GetRow("CAR", columns)
  if err != nil {
    return nil, fmt.Errorf("Failed retrieving customer [%s]: [%s]", custId, err)
  }

  //get the corresponding supplier cert and grid operator cert
  supplierTcert := row.Columns[3].GetBytes()
  gridOperatorTcert := row.Columns[4].GetBytes()
  myTcert,err := stub.GetState("myTcert")
  if err != nil {
		return nil, errors.New("Failed fetching myTcert")
	}

  if (string(myTcert) != string(supplierTcert) && string(myTcert) != string(gridOperatorTcert)){
    return nil, fmt.Errorf("[%s]: You don't have permissions to do this operation.", myTcert)
  }

  custTcert := row.Columns[1].GetBytes()
  custProvidedTcert, err := base64.StdEncoding.DecodeString(args[1])
  if err != nil {
		return nil, errors.New("Failed decoding customer provided certificate")
	}

  if (string(custTcert) != string(custProvidedTcert)){
    return nil, fmt.Errorf("[%s]: Customer certificate mis-match.", custProvidedTcert)
  }

  err = stub.DeleteRow(
    "CAR",
    []shim.Column{shim.Column{Value: &shim.Column_String_{String_: custId}}},
  )
  if err != nil {
    return nil, errors.New("Failed deliting row.")
  }

  _, err = stub.InsertRow(
    "CAR",
    shim.Row{
      Columns: []*shim.Column{
  			&shim.Column{Value: &shim.Column_String_{String_: custId}},
  			&shim.Column{Value: &shim.Column_Bytes{Bytes: custTcert}},
        &shim.Column{Value: &shim.Column_Bytes{Bytes: []byte(args[2])}},
        &shim.Column{Value: &shim.Column_Bytes{Bytes: supplierTcert}},
        &shim.Column{Value: &shim.Column_Bytes{Bytes: gridOperatorTcert}}},
    })
  if err != nil {
    return nil, errors.New("Failed inserting row.")
  }

  myLogger.Debug("updateMandate...done")

	return nil, nil

}


func (t *MandateChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {

	// Handle different functions
	if function == "updateMandate" {
		// Update mandate information of user
		return t.updateMandate(stub, args)
	}

	return nil, errors.New("Received unknown function invocation")
}

func (t *MandateChaincode) getCustomerDetails(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	myLogger.Debug("getCustomerDetails...")
	fmt.Printf("getCustomerDetails method started")

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}

  custId := args[0]
	fmt.Printf("Received custId: %s", custId)
	myLogger.Debug("custId: [%s]",custId)

	//testing
	//return []byte("before GetRow"),nil

  var columns []shim.Column
  col1 := shim.Column{Value: &shim.Column_String_{String_: custId}}
  columns = append(columns, col1)
	fmt.Printf("database query started")
	myLogger.Debug("database query started...")

  row, err := stub.GetRow("CAR", columns)
  if err != nil {
    return nil, fmt.Errorf("Failed retrieving customer [%s]: [%s]", custId, err)
  }

	fmt.Printf("database query completed")
	myLogger.Debug("database query completed...")
	myLogger.Debug("accessing columns...")

	myLogger.Debug("column 0...[%s]",row.Columns[0].GetBytes())
	myLogger.Debug("column 1...[%s]",row.Columns[1].GetBytes())
	myLogger.Debug("column 2...[%s]",row.Columns[2].GetBytes())

  //get the corresponding supplier cert and grid operator cert
  supplierTcert := row.Columns[3].GetBytes()
	fmt.Printf("supplier tcert: %s", supplierTcert)
	myLogger.Debug("supplier tcert...%s",supplierTcert)

  gridOperatorTcert := row.Columns[4].GetBytes()
	fmt.Printf("grid operator tcert: %s", gridOperatorTcert)
	myLogger.Debug("grid tcert...%s",gridOperatorTcert)

  myTcert,err := stub.GetState("myTcert")
  if err != nil {
		return nil, errors.New("Failed fetching myTcert")

	}
	fmt.Printf("ny tcert: %s", myTcert)
	myLogger.Debug("my tcert...%s",myTcert)

  if (string(myTcert) != string(supplierTcert) && string(myTcert) != string(gridOperatorTcert)){
    return nil, fmt.Errorf("[%s]: You don't have permissions to do this operation.", myTcert)
  }

  custTcert := row.Columns[1].GetBytes()
	fmt.Printf("customer tcert: %s", custTcert)
	myLogger.Debug("customer tcert...%s",custTcert)

  custProvidedTcert, err := base64.StdEncoding.DecodeString(args[1])
  if err != nil {
		return nil, errors.New("Failed decoding customer provided certificate")
	}
	fmt.Printf("customer provided tcert: %s", custProvidedTcert)
	myLogger.Debug("cust provided tcert...%s",custProvidedTcert)

  if (string(custTcert) != string(custProvidedTcert)){
    return nil, fmt.Errorf("[%s]: Customer certificate mis-match.", custProvidedTcert)
  }

	fmt.Printf("return data: %s", row.Columns[2].GetBytes())
	myLogger.Debug("return data...%s",row.Columns[2].GetBytes())

  return row.Columns[2].GetBytes(),nil

}


func (t *MandateChaincode) getAllCustomerDetails(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	myLogger.Debug("getAllCustomerDetails...")
	fmt.Printf("getAllCustomerDetails method started")

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	receivedTcert, err := base64.StdEncoding.DecodeString(args[0])
  if err != nil {
		return nil, errors.New("Failed decoding received certificate")
	}

	myLogger.Debug("received tcert...%s",receivedTcert)

	edsnTcert,err := stub.GetState("edsnTcert")
  if err != nil {
		return nil, errors.New("Failed fetching edsnTcert")

	}

	if (string(receivedTcert) != string(edsnTcert)){
    return nil, fmt.Errorf("[%s]: You don't have suffucient privilege for this operation.", receivedTcert)
  }

	var columns []shim.Column
	//col1 := shim.Column{Value: &shim.Column_String_{String_: custId}}
	//columns = append(columns, col1)
	fmt.Printf("database query started")
	myLogger.Debug("database query started...")

	rowChannel, err := stub.GetRows("CAR", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed retrieving all customers: [%s]", err)
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

	var s string
	s += "{"
	for i := 0;  i< len(rows); i++ {
		s += "'" + rows[i].Columns[0].GetString_() + "':"
		myLogger.Debug("data: [%s]",s)

		s += string(rows[i].Columns[2].GetBytes()) + ","
		myLogger.Debug("data: [%s]",s)
	}
	s = strings.TrimRight(s, ",")
	s += "}"

	return []byte(s),nil

}

func (t *MandateChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {

	// Handle different functions
	if function == "single" {
		// Get customer's details
		return t.getCustomerDetails(stub, args)
	}else if function == "all" {
		// Return all customer details
		return t.getAllCustomerDetails(stub, args)
	}

	return nil, errors.New("Received unknown function invocation")
}
