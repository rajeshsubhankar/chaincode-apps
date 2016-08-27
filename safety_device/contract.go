/*


*/
package main

import(
  "fmt"
  "errors"
  "strconv"
  "encoding/json"

  "github.com/hyperledger/fabric/core/chaincode/shim"
)

// ContractChaincode example contract Chaincode implementation
type ContractChaincode struct {
}

func main() {
    err := shim.Start(new(ContractChaincode))
    if err != nil {
      fmt.Printf("Error starting chaincode: %s",err)
    }
}

func (t *ContractChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error)  {

    if len(args) != 1 {
      return nil, errors.New("Incorrect number of arguments, Expecting 1")
    }

    err := stub.PutState("owner",[]byte(args[0]))
    if err != nil {
      return nil,err
    }

    err = stub.PutState("eventId",[]byte(strconv.Itoa(1)))
    if err != nil {
      return nil,err
    }

    return nil,nil
}

func (t *ContractChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error){
    fmt.Println("Invoke is running" + function)

    //Handle different function
    if function == "init" {
      return t.Init(stub, "init", args)
    } else if function == "write" {
      return t.write(stub, args)
    }

    fmt.Println("Invoke didn't find func: " + function)

    return nil, errors.New("Received unknown function invocation")
}

func (t *ContractChaincode) write(stub *shim.ChaincodeStub, args[]string) ([]byte, error)  {
  if len(args) != 10 {
    return nil, errors.New("Incorrect number of arguments, Expecting 10")
  }

    var key int
    var value map[string]string
    var err error
    var data []byte

    key,err = t.getEventId(stub)
    if err != nil {
      return nil, err
    }

    value = make(map[string]string)

    value["deviceId"] = args[0]
    value["deviceType"] = args[1]
    value["timestamp"] = args[2]
    value["passengerId"] = args[3]
    value["passengerDetails"] = args[4]
    value["driverId"] = args[5]
    value["driverDetails"] = args[6]
    value["vehicleDetails"] = args[7]
    value["gps"] = args[8]
    value["videoHash"] = args[9]

    data, err = json.Marshal(value)

    err = stub.PutState(strconv.Itoa(key), data)
    if err != nil {
      return nil, err
    }

    //successful blockchain update
    key = key +1

    return t.updateEventId(stub, key)
}

func (t *ContractChaincode) getEventId(stub *shim.ChaincodeStub) (int, error) {
    eventId,err := stub.GetState("eventId")
    if err != nil {
      return -1, err
    }
    s, err1 := strconv.Atoi(string(eventId))
    if err1 != nil {
      return -1, err1
    }

    return s, nil
}

func (t *ContractChaincode) updateEventId(stub * shim.ChaincodeStub, key int) ([]byte, error)  {
    err := stub.PutState("eventId", []byte(strconv.Itoa(key)))

    if err != nil {
      return nil, err
    }

    return []byte(strconv.Itoa(key)),nil

}

func (t *ContractChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error)  {
  fmt.Println("Query is running " + function)

  if function == "read" {
    return t.read(stub, args)
  }
  fmt.Println("Query didn't find function: " + function)

  return nil, errors.New("Received unknown function error")

}

func (t *ContractChaincode)read(stub *shim.ChaincodeStub, args []string) ([]byte, error)  {
  if len(args) != 1 {
      return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
  }

  key := args[0]
  valAsbytes, err := stub.GetState(key)
  if err != nil {
      jsonResp := "{\"Error\":\"Failed to get state for " + key + "\"}"
      return nil, errors.New(jsonResp)
    }

    return valAsbytes, nil

}
