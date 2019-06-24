package accounty

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
	pb "go.appointy.com/accounty/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errInternal = status.Error(codes.Internal, `Oops! Something went wrong`) // Generic Error to be returned to client to hide possible sensitive information
)

//variable decleration

var tokenqb string = "Bearer eyJlbmMiOiJBMTI4Q0JDLUhTMjU2IiwiYWxnIjoiZGlyIn0..08NMJlKy9bfhZWqsPPvqSg.zFID7O3BZLyuUU3i13Bm0Ru1bGUJLbBgKn_m2DWzt7YZTE7kGtlUgEMr14gxXQE6VIH4P57TVe0aDUUQrsJxhWFGWaNL_FzHsbd67xNfDPtHJv1khpLuPDZ75GuffyAGhU6vwahjFlF6tb6dNA7lMt7d_UKWiarPlnkjL2QAlYn62ROlE6UWI6uoQNv37RklyhIMtzRzM0SoyDKJilJj-V1Pz_919nlwUiCW3pD1QEJWzqWAwm3QJX-BmVW6Yey9dGU64cJJUgTcN8K2kAXO1l0TfNNspKKKCcSxwQHmD0zlUCeb9Hedxo_OBpAPa-sZMZCEab7sdHhBkWJAhkvoaE5rxJtm7TTCkBOodyn0bV5Q1tHyPtgnQ_8gr1lawK1HFjk8JtKBWrDfIc084qy9XCcHB_oU10YE8NKekrMBDpWw1FRyRrLta7zZtcA00tIpTF5p7C5AHA15xN7fVPIsQTAFhqn0LVpUK0H0xr1-WW8vLmp8qb9RweKjFSFlnz7mu-C0PAua94dNhr-mFjmynpOw6JnM5jCuEyP8fpgvEOaMfu2XpR-2TctkYMh92aBeZ-tackuvynNU__zd_a_A9Bwpr0WePlXusFr52iDpHOHHktgZ05WFkf8xxu8Y_ezpFIV3aqP0R1QeyrhRa66C2-vrSDdnS_zO5Uog4LeMKr34qVCTzaWx0R1z2RSD-8M96QzZvzX97O_VsJp1PH5AWVozSNX7Y1k1rhB7su-DJQoF52TiR9uPw0csaApIZP1GfjJG4JPaLSeOcnnkEpxH26XlafLb9F365xnn2eUuDpMlOnhyYzLDHbnUQQFQK0JO.BUdSTAZZTHqPEK_oPVVOvQ"
var tokenxero string = "xero token will go here"
var tokenfb string = "freshbook token will go here"

type accountyServer struct {
	//add store
	store pb.AccoutingEmployeeLinkStore
	//add other clients
}

//struct decleration
//struct to be decleard for accepting json request
//these decleared struct will be transfered to separate file for ease of management

type Response struct {
	QuerryResponse  *NextResponse `json:"QueryResponse"`
	TimeForDBAccess string        `json:"time"`
}

type NextResponse struct {
	Account *[]NextLedger `json:"Account"`
}

type NextLedger struct {
	CurBal string `json:"CurrentBalance"`
}

type Employee1 struct {
	Employee     *Layer1 `json:"Employee"`
	XeroEmployee *Layer1 `json:"Employees"`
}

type Customer1 struct {
	QbCustomer *Layer1 `json:"Customer"`
	QbVendor   *Layer1 `json:"Vendor"`
}

type Timeactivity1 struct {
	BusinessHrs *Layer1 `json:"TimeActivity"`
	TimeEntryFB *Layer1 `json:"time_entry"`
}

type Payment1 struct {
	PaymentQb    *Layer1 `json:"Payment"`
	BillQb       *Layer1 `json:"Bill"`
	SaleReciptQb *Layer1 `json:"SalesReceipt"`
	InvoiceQb    *Layer1 `json:"Invoice"`
	ItemQb       *Layer1 `json:"Item"`
}

type Item1 struct {
	ItemQb *Layer1 `json:"Item"`
}

type Layer1 struct {
	Id             string  `json:"Id"`
	EmpRef         *Layer2 `json:"EmployeeRef"`
	IdFBTimeEntry  string  `json:"client_id"`
	XeroEmployeeId string  `json:"EmployeeID"`
}

type Layer2 struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

//NewquickbookServerServer returns a quickbookServerServer implementation with core business logic
func NewquickbookServerServer(Store pb.AccoutingEmployeeLinkStore) pb.QuickbookServerServer {
	return &accountyServer{store: Store}
}

//This function contains boiler plate code for calling API of Quickbook, Freshbook and Xero.
//External function for OAuth verification purposes will replace this function.
func ConnectApiEndpoint(urlqb, payloadqbstr, token string) ([]byte, error) {

	payloadqb := strings.NewReader(payloadqbstr)
	req, _ := http.NewRequest("POST", urlqb, payloadqb)
	req.Header.Add("User-Agent", "QBOV3-OAuth2-Postman-Collection")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)
	//req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "7380246b-5997-4989-afdf-223116950cc1")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	return body, err
}

//function for Xero API connection.

func ConnectToXeroAPI(urlqb, payloadqbstr string) ([]byte, error) {

	payloadqb := strings.NewReader(payloadqbstr)
	req, _ := http.NewRequest("POST", urlqb, payloadqb)
	req.Header.Add("Authorization", `,OAuth oauth_consumer_key="54BT1LQ29XWN7N7PMCLOZ5KBE0WTFY",oauth_token="R9CFAZZODQVLLMORPZBSHATDVNR4S5",oauth_signature_method="HMAC-SHA1",oauth_timestamp="1559040766",oauth_nonce="DCzne7kxHU0",oauth_version="1.0",oauth_signature="AEoDCYAq2qj3Slsp3U41uQvRIlw%3D"`)
	req.Header.Add("accept", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "b4df501f-ebdb-4007-b7d4-b9f42a310924")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return body, err
}

//This is the function to create our required account in quickbook but it will be redundent as the accounts will be matched by the users at the time of syncing.
/*func CreateAccountLedger(CompanyId string) string {

	url := "https://SANDBOX-QUICKBOOKS.API.INTUIT.COM/v3/company/" + CompanyId + "/query?query=select%20%2A%20from%20Account%20where%20Name=%27Accounts%20Receivable%20%28A/R%29%27&minorversion=37"
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("User-Agent", "QBOV3-OAuth2-Postman-Collection")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", tokenqb)
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "1a60fc8f-6859-4248-954e-3ff2ea592fe7")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	Account := Response{}
	jsonErr := json.Unmarshal(body, &Account)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(Account.QuerryResponse.Account.)
	/*payload := "{\n  \"Classification\": \"" + Classification + "\",  \n  \"AccountType\": \"" + AccType + "\",\n  \"Name\": \"" + Name + "\"\n}"
	body, _ := ConnectApiEndpoint(url, payload, tokenqb)
	fmt.Println(string(body))
	account := Account1{}
	jsonErr := json.Unmarshal(body, &Emp1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return "done"

}*/

//***********************************************************************************************************************************************************************************************************************
//###############################################core account modification RPCs(Quickbook) #########################################################################################################################################
//***********************************************************************************************************************************************************************************************************************
//sale recipt API for payment##########################################################################################################################################################################################

func SaleRecipt(CompanyId, AppointyId, DepositAccType, id, Description, TotalAmount string, Store pb.AccoutingEmployeeLinkStore) (string, error) {
	url := "https://SANDBOX-QUICKBOOKS.API.INTUIT.COM/v3/company/" + CompanyId + "/salesreceipt?minorversion=37"
	payload := "{   \n\t\"DepositToAccountRef\": {\n      \"name\": \"" + DepositAccType + "\"\n    }, \n    \"Line\": [{\n        \"Id\": \"" + id + "\",\n        \"Description\": \"" + Description + "\",\n        \"Amount\": " + TotalAmount + ",\n        \"DetailType\": \"SalesItemLineDetail\",\n        \"SalesItemLineDetail\": {\n            \n            \"TaxCodeRef\": {\n                \"value\": \"NON\"\n            }\n        }\n    }]\n}\n"
	body, err := ConnectApiEndpoint(url, payload, tokenqb)
	NewAppointmentSale := Payment1{}
	jsonErr := json.Unmarshal(body, &NewAppointmentSale)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	SaleReciptAppointyId := "SaleRecipt" + AppointyId
	ids, errDB1 := Store.CreateAccoutingEmployeeLinks(context.Background(), &pb.AccoutingEmployeeLink{
		AppointyId: SaleReciptAppointyId,
		ExternalId: NewAppointmentSale.SaleReciptQb.Id,
		AppType:    1,
	})
	if errDB1 != nil {
		return "null", errInternal
	}

	return ids[0], err

}

func DeleteSaleRecipt(CompanyId, Id string) ([]byte, error) {
	url := "https://SANDBOX-QUICKBOOKS.API.INTUIT.COM/v3/company/" + CompanyId + "/salesreceipt?operation=delete"
	payload := "{\n    \"Id\": \"" + Id + "\",\n    \"SyncToken\": \"1\"\n}"
	body, err := ConnectApiEndpoint(url, payload, tokenqb)
	return body, err
}

func CashMemo(CompanyId, DepositAccType, id, Description, TotalAmount string) ([]byte, error) {
	url := "https://SANDBOX-QUICKBOOKS.API.INTUIT.COM/v3/company/" + CompanyId + "/salesreceipt?minorversion=37"
	payload := "{   \n\t\"DepositToAccountRef\": {\n      \"name\": \"" + DepositAccType + "\"\n    }, \n    \"Line\": [{\n        \"Id\": \"" + id + "\",\n        \"Description\": \"" + Description + "\",\n        \"Amount\": " + TotalAmount + ",\n        \"DetailType\": \"SalesItemLineDetail\",\n        \"SalesItemLineDetail\": {\n            \n            \"TaxCodeRef\": {\n                \"value\": \"NON\"\n            }\n        }\n    }]\n}\n"
	body, err := ConnectApiEndpoint(url, payload, tokenqb)
	return body, err
}

//Bill payment API for payment##########################################################################################################################################################################################

/*func (s *accountyServer) BillPayment(ctx context.Context, req *pb.CreateBillPaymentRequest) (*pb.PaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BillPayment not implemented")
}
*/
//Bill API for payment##########################################################################################################################################################################################

func Bill(CompanyId, AppointyId, DepositAccType, Amount, Customer string, Store pb.AccoutingEmployeeLinkStore) (string, error) {
	url := "https://SANDBOX-QUICKBOOKS.API.INTUIT.COM/v3/company/" + CompanyId + "/bill?minorversion=37"
	payload := "{\n\t\"APAccountRef\": {\n      \"name\": \"" + DepositAccType + "\"\n     \n    }, \n    \"Line\":[\n        {\n             \"Amount\": " + Amount + ",\n            \"DetailType\":\"AccountBasedExpenseLineDetail\",\n            \"AccountBasedExpenseLineDetail\":\n            {\n                \"AccountRef\":\n                {\n                    \"value\":\"7\"\n                }\n            }\n        } \n    ],\n    \"VendorRef\":\n    {\n        \"value\":\"" + Customer + "\"\n    }\n}"
	body, err := ConnectApiEndpoint(url, payload, tokenqb)
	NewAppointmentBill := Payment1{}
	jsonErr := json.Unmarshal(body, &NewAppointmentBill)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	BillAppointyId := "Bill" + AppointyId
	ids, errDB2 := Store.CreateAccoutingEmployeeLinks(context.Background(), &pb.AccoutingEmployeeLink{
		AppointyId: BillAppointyId,
		ExternalId: NewAppointmentBill.BillQb.Id,
		AppType:    1,
	})
	if errDB2 != nil {
		return "null", errInternal
	}

	return ids[0], err
}

func BillDelete(CompanyId, BillId string) ([]byte, error) {
	url := "https://SANDBOX-QUICKBOOKS.API.INTUIT.COM/v3/company/" + CompanyId + "/bill?operation=delete"
	payload := "{\n    \"Id\": \"" + BillId + "\",\n    \"SyncToken\": \"0\"\n}\n"
	body, err := ConnectApiEndpoint(url, payload, tokenqb)
	return body, err
}

//Payment API for payment##########################################################################################################################################################################################
//used to reduce amount to cash
func Payment(CompanyId, AppointyId, DepositAccType, TotalAmt, CustId string, Store pb.AccoutingEmployeeLinkStore) (string, error) {
	url := "https://SANDBOX-QUICKBOOKS.API.INTUIT.COM/v3/company/" + CompanyId + "/payment?minorversion=37"
	payload := "{\n    \"ARAccountRef\":{\n    \t\"name\":\"" + DepositAccType + "\"\n    },\n    \"CustomerRef\":\n    {\n        \"value\": \"" + CustId + "\" \n    },\n    \"TotalAmt\": " + TotalAmt + "\n    \n}"
	body, err := ConnectApiEndpoint(url, payload, tokenqb)
	NewAppointmentPay := Payment1{}
	jsonErr := json.Unmarshal(body, &NewAppointmentPay)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	PaymentAppointyId := "Payment" + AppointyId
	ids, errDB := Store.CreateAccoutingEmployeeLinks(context.Background(), &pb.AccoutingEmployeeLink{
		AppointyId: PaymentAppointyId,
		ExternalId: NewAppointmentPay.PaymentQb.Id,
		AppType:    1,
	})
	if errDB != nil {
		return "null", errInternal
	}
	return ids[0], err
}

func PaymentDelete(CompanyId, PaymentId string) ([]byte, error) {
	url := "https://SANDBOX-QUICKBOOKS.API.INTUIT.COM/v3/company/" + CompanyId + "/payment?operation=delete"
	payload := "{\n    \"Id\": \"" + PaymentId + "\",\n    \"SyncToken\": \"1\"\n}"
	body, err := ConnectApiEndpoint(url, payload, tokenqb)
	return body, err
}

//Invoice payment method for payment##########################################################################################################################################################################################
//used to add amount to cash account
func InvoicePayment(CompanyID, AppointyId, Amount, CustomerID string, Store pb.AccoutingEmployeeLinkStore) (string, error) {
	url := "https://SANDBOX-QUICKBOOKS.API.INTUIT.COM/v3/company/" + CompanyID + "/invoice?minorversion=37"
	payload := "{\n  \"Line\": [\n    {\n      \"Amount\": " + Amount + ",\n      \"DetailType\": \"SalesItemLineDetail\",\n      \"SalesItemLineDetail\": {\n        \"ItemRef\": {\n          \"value\": \"1\",\n          \"name\": \"Services\"\n        }\n      }\n    }\n  ],\n  \"CustomerRef\": {\n    \"value\": \"" + CustomerID + "\"\n  }\n}"
	body, err := ConnectApiEndpoint(url, payload, tokenqb)
	NewAppointmentPay := Payment1{}
	jsonErr := json.Unmarshal(body, &NewAppointmentPay)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	PaymentAppointyId := "Invoice" + AppointyId
	ids, errDB3 := Store.CreateAccoutingEmployeeLinks(context.Background(), &pb.AccoutingEmployeeLink{
		AppointyId: PaymentAppointyId,
		ExternalId: NewAppointmentPay.InvoiceQb.Id,
		AppType:    1,
	})
	if errDB3 != nil {
		return "null", errInternal
	}
	return ids[0], err
}

//Item API for payment##########################################################################################################################################################################################

func Item(CompanyId, ItemName, CustId, Qnt, TotalAmt string, paymentdone int) ([]byte, error) {
	url := "https://SANDBOX-QUICKBOOKS.API.INTUIT.COM/v3/company/" + CompanyId + "/item?minorversion=37"
	payload := "{\n  \"Name\": \"" + ItemName + "\",\n  \"IncomeAccountRef\": {\n    \"value\": \"79\",\n    \"name\": \"Sales of Product Income\"\n  },\n  \"ExpenseAccountRef\": {\n    \"value\": \"80\",\n    \"name\": \"Cost of Goods Sold\"\n  },\n  \"AssetAccountRef\": {\n    \"value\": \"81\",\n    \"name\": \"Inventory Asset\"\n  },\n  \"Type\": \"Inventory\",\n  \"TrackQtyOnHand\": true,\n  \"QtyOnHand\": " + Qnt + ",\n  \"InvStartDate\": \"2015-01-01\"\n}"
	body1, err := ConnectApiEndpoint(url, payload, tokenqb)
	if paymentdone == 1 {
		url = "https://SANDBOX-QUICKBOOKS.API.INTUIT.COM/v3/company/" + CompanyId + "/transfer?minorversion=37"
		//from account object querry Ids of checking account and prepaid expense(for now) account
		payload = "{\n  \"FromAccountRef\": {\n        \"value\": \"35\",\n        \"name\": \"checking\"\n    },\n    \"ToAccountRef\": {\n        \"value\": \"3\",\n        \"name\": \"Prepaid Expenses\"\n    },\n    \"Amount\": \"" + TotalAmt + "\"\n}\n"
		/*url = "https://SANDBOX-QUICKBOOKS.API.INTUIT.COM/v3/company/" + CompanyId + "/payment?minorversion=37"
		payload = "{\n    \"ARAccountRef\":{\n    \t\"name\":\"Accounts Recieved (A/R)\"\n    },\n    \"CustomerRef\":\n    {\n        \"value\": \"" + CustId + "\" \n    },\n    \"TotalAmt\": " + TotalAmt + "\n    \n}"*/
		_, err = ConnectApiEndpoint(url, payload, tokenqb)
	}
	if paymentdone == 2 {
		url = "https://SANDBOX-QUICKBOOKS.API.INTUIT.COM/v3/company/" + CompanyId + "/bill?minorversion=37"
		payload = "{\n\t\"APAccountRef\": {\n      \"name\": \"Accounts Payable (A/P)\"\n     \n    }, \n    \"Line\":[\n        {\n             \"Amount\": " + TotalAmt + ",\n            \"DetailType\":\"AccountBasedExpenseLineDetail\",\n            \"AccountBasedExpenseLineDetail\":\n            {\n                \"AccountRef\":\n                {\n                    \"value\":\"7\"\n                }\n            }\n        } \n    ],\n    \"VendorRef\":\n    {\n        \"value\":\"" + CustId + "\"\n    }\n}"
		_, err = ConnectApiEndpoint(url, payload, tokenqb)
	}
	return body1, err
}

/*func DeleteItem(CompanyId, ItemId, PaymentId, Qnt, TotalAmt string, paymentdone int) ([]byte, error) {
url := "https://SANDBOX-QUICKBOOKS.API.INTUIT.COM/v3/company/" + CompanyId + "/item?minorversion=37"
payload := "{\n    \n    \"Active\": false,\n    \n    \n   \n    \"Type\": \"Inventory\",\n    \"IncomeAccountRef\": {\n      \"value\": \"79\",\n      \"name\": \"Sales of Product Income\"\n    },\n   \n    \"ExpenseAccountRef\": {\n      \"value\": \"80\",\n      \"name\": \"Cost of Goods Sold\"\n    },\n    \"AssetAccountRef\": {\n      \"value\": \"81\",\n      \"name\": \"Inventory Asset\"\n    },\n    \"Id\": \"" + ItemId + "\"\n    \n  }"
body1, err := ConnectApiEndpoint(url, payload, tokenqb)
if paymentdone == 1 {
	url = "https://SANDBOX-QUICKBOOKS.API.INTUIT.COM/v3/company/" + CompanyId + "/transfer?minorversion=37"
	//from account object querry Ids of checking account and prepaid expense(for now) account
	payload = "{\n  \"FromAccountRef\": {\n        \"value\": \"35\",\n        \"name\": \"checking\"\n    },\n    \"ToAccountRef\": {\n        \"value\": \"3\",\n        \"name\": \"Prepaid Expenses\"\n    },\n    \"Amount\": \"" + TotalAmt + "\"\n}\n"
	/*url = "https://SANDBOX-QUICKBOOKS.API.INTUIT.COM/v3/company/" + CompanyId + "/payment?minorversion=37"
	payload = "{\n    \"ARAccountRef\":{\n    \t\"name\":\"Accounts Recieved (A/R)\"\n    },\n    \"CustomerRef\":\n    {\n        \"value\": \"" + CustId + "\" \n    },\n    \"TotalAmt\": " + TotalAmt + "\n    \n}"*/
/*	_, err = ConnectApiEndpoint(url, payload, tokenqb)
	}
	if paymentdone == 2 {
		url = "https://SANDBOX-QUICKBOOKS.API.INTUIT.COM/v3/company/" + CompanyId + "/bill?minorversion=37"
		payload = "{\n\t\"APAccountRef\": {\n      \"name\": \"Accounts Payable (A/P)\"\n     \n    }, \n    \"Line\":[\n        {\n             \"Amount\": " + TotalAmt + ",\n            \"DetailType\":\"AccountBasedExpenseLineDetail\",\n            \"AccountBasedExpenseLineDetail\":\n            {\n                \"AccountRef\":\n                {\n                    \"value\":\"7\"\n                }\n            }\n        } \n    ],\n    \"VendorRef\":\n    {\n        \"value\":\"" + CustId + "\"\n    }\n}"
		_, err = ConnectApiEndpoint(url, payload, tokenqb)
	}
	return body1, err
}*/

//Vendor API for payment##########################################################################################################################################################################################

/*func (s *accountyServer) Vendor(ctx context.Context, req *pb.CreateVendorRequest) (*pb.PaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Vendor not implemented")
}*/

//***********************************************************************************************************************************************************************************************************************
//###############################################Basic accounting RPCs using core account balencing RPCs (common for Quickbook, freshbook and Xero)#########################################################################################################################################
//***********************************************************************************************************************************************************************************************************************

//CreateEmployee ########################################################################################################################################################################################################
func (s *accountyServer) CreateEmployee(ctx context.Context, in *pb.CreateEmployeeRequest) (*pb.Employee, error) {

	sendemp := pb.Employee{Id: "No App synced with accounty service of appointy"}

	//Quickbook employee add.
	if in.AppType == 1 {

		urlqb := "https://SANDBOX-QUICKBOOKS.API.INTUIT.COM/v3/company/" + in.CompanyId + "/employee?minorversion=37"
		CountrySubDivisionCode := in.Employee.Metadata["countrycode"]
		City := in.Employee.Metadata["city"]
		PostalCode := in.Employee.Metadata["pincode"]
		//Creating a staff in quickbook database
		payloadqbstr := "{\n  \"GivenName\": \"" + in.Employee.FirstName + "\", \n  \"SSN\": \"XXX-XX-XXXX\", \n  \"PrimaryAddr\": {\n    \"CountrySubDivisionCode\": \"" + CountrySubDivisionCode + "\", \n    \"City\": \"" + City + "\", \n    \"PostalCode\": \"" + PostalCode + "\", \n    \"Line1\": \"null\"\n  }, \n  \"PrimaryPhone\": {\n    \"FreeFormNumber\": \"" + in.Employee.PhoneNumber + "\"\n  }, \n  \"FamilyName\": \"" + in.Employee.LastName + "\"\n}"

		//calling the API
		body, APIerr := ConnectApiEndpoint(urlqb, payloadqbstr, tokenqb)
		if APIerr != nil {
			log.Fatal(APIerr)
			sendemp = pb.Employee{Id: "No account found in Quickbook"}
		}

		//Unmarshelling the json response in Employee1 struct jsut created.(doubt)
		Emp1 := Employee1{}
		jsonErr := json.Unmarshal(body, &Emp1)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		log.Println(Emp1)
		//Inserting Id generated into chaku database
		ids, _ := s.store.CreateAccoutingEmployeeLinks(ctx, &pb.AccoutingEmployeeLink{
			AppointyId: in.Employee.AppointyId,
			ExternalId: Emp1.Employee.Id,
			AppType:    1,
		})

		//Returning poulated employee struct with Id of the database back to program for further use.
		in.Employee.Id = ids[0]
		sendemp = pb.Employee{Id: in.Employee.Id}
		return &sendemp, nil
	}

	//	Create Xero employee
	if in.AppType == 2 {
		urlqb := "https://api.xero.com/api.xro/2.0/Employees"
		payloadqbstr := "{\n  \"Employees\": [\n    {\n      \"FirstName\": \"" + in.Employee.FirstName + "\",\n      \"LastName\": \"" + in.Employee.LastName + "\"\n    }\n  ]\n}"

		//calling the API
		body, APIerr := ConnectToXeroAPI(urlqb, payloadqbstr)
		if APIerr != nil {
			log.Fatal(APIerr)
			sendemp = pb.Employee{Id: "No account found in Xero"}
		}

		//Unmarshelling the json response in Employee1 struct jsut created.(doubt)
		/*Emp2 := Employee1{}
		jsonErr := json.Unmarshal(body, &Emp2)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		*/
		log.Println(string(body))

		//Inserting Id generated into chaku database
		ids, _ := s.store.CreateAccoutingEmployeeLinks(ctx, &pb.AccoutingEmployeeLink{
			AppointyId: in.Employee.AppointyId,
			ExternalId: "4",
			AppType:    2,
		})

		//Returning poulated employee struct with Id of the database back to program for further use
		in.Employee.Id = ids[0]
		sendemp = pb.Employee{Id: in.Employee.Id}
		return &sendemp, nil
	}

	if in.AppType == 3 {
		//Apply Freshbook logic
		fmt.Println("Freshbook create employee API end point.")
	}
	return &sendemp, nil

}

//DeleteEmployee #################################################################################################################################################################################################################
//Only implimented for Quickbook.
func (s *accountyServer) DeleteEmployee(ctx context.Context, in *pb.DeleteEmployeeRequest) (*empty.Empty, error) {

	empl, err := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id", "app_type"}, pb.AccoutingEmployeeLinkAppointyIdEq{in.AppointyEmployeeId})
	if err != nil {
		return nil, errInternal
	}
	//extract employee information from quickbook using ExternalId and pass it to the below string

	//Calling Delete API in Quickbook
	urlqb := "https://SANDBOX-QUICKBOOKS.API.INTUIT.COM/v3/company/" + in.CompanyId + "/employee"
	payloadqbstr := "{\t\n  \"Active\": false, \n   \"Id\": \"" + empl.ExternalId + "\",\n   \"GivenName\": \"from appointy database\",\n    \"FamilyName\": \"from appointy database\"\n  }"

	//calling the API
	body, Ferr := ConnectApiEndpoint(urlqb, payloadqbstr, tokenqb)

	if Ferr != nil {
		return nil, Ferr
	}
	fmt.Println(string(body))
	return nil, err

}

//CreateBusinessHour ######################################################################################################################################################################################################
func (s *accountyServer) CreateBusinessHour(ctx context.Context, in *pb.CreateBusinessHourRequest) (*pb.BusinessHour, error) {

	var urlqb [3]string
	var payloadqbstr [3]string
	returnBHrs := pb.BusinessHour{Name: "No App synced with accounty service."}
	var j pb.AccountingIntegrationType = 0
	var p int = int(j)
	var token [3]string
	for p = 0; p < 3; p++ {
		j = j + 1
		empl, err := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id", "app_type"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{j}, pb.AccoutingEmployeeLinkAppointyIdEq{in.BusinessHour.AppointyEmployeeId}})
		if empl == nil {
			fmt.Println(err)
			continue
		}
		if p == 0 {
			urlqb[p] = "https://SANDBOX-QUICKBOOKS.API.INTUIT.COM/v3/company/" + in.CompanyId + "/timeactivity?minorversion=37"
			payloadqbstr[p] = "{\n    \"NameOf\": \"Employee\",\n    \"EmployeeRef\": {\n        \"value\": \"" + empl.ExternalId + "\",\n        \"name\": \"" + in.BusinessHour.Name + "\"\n    },\n    \"StartTime\": \"2011-07-05T17:00:00-08:00\",\n    \"EndTime\": \"2013-07-05T17:00:00-08:00\"\n}\n"
			token[p] = tokenqb
		}
		if p == 1 {
			urlqb[p] = ""
			payloadqbstr[p] = "{\n    \"NameOf\": \"Employee\",\n    \"EmployeeRef\": {\n        \"value\": \"" + empl.ExternalId + "\",\n        \"name\": \"" + in.BusinessHour.Name + "\"\n    },\n    \"StartTime\": \"2011-07-05T17:00:00-08:00\",\n    \"EndTime\": \"2013-07-05T17:00:00-08:00\"\n}\n"
			token[p] = tokenxero
		}
		if p == 2 {
			urlqb[p] = "https://api.freshbooks.com/timetracking/business/" + in.CompanyId + "/time_entries"
			payloadqbstr[p] = "{\n    \"time_entry\": {\n        \"is_logged\": true,\n        \"duration\": " + in.BusinessHour.TotalTime + ",\n        \"note\": \"" + in.BusinessHour.Description + "\",\n        \"started_at\": \"2016-08-16T20:00:00.000Z\",\n        \"client_id\": " + empl.ExternalId + ",\n        \"project_id\": {{projectId}}\n    }\n}"
			token[p] = tokenfb
		}
		body, APIerr := ConnectApiEndpoint(urlqb[p], payloadqbstr[p], token[p])
		if APIerr != nil {
			return nil, APIerr
		}
		if p == 0 {

			BusinessHour1 := Timeactivity1{}
			jsonErr := json.Unmarshal(body, &BusinessHour1)
			if jsonErr != nil {
				log.Fatal(jsonErr)
			}
			//returnBHrs = pb.BusinessHour{Description: BusinessHour1.BusinessHrs.EmpRef.Value}
			returnBHrs = pb.BusinessHour{Name: "here bro its done partially"}

		}

		if p == 1 {

			//Unmarshelling the json response in Employee1 struct jsut created.
			bshr := Timeactivity1{}
			jsonErr := json.Unmarshal(body, &bshr)
			if jsonErr != nil {
				log.Fatal(jsonErr)
			}
			returnBHrs = pb.BusinessHour{Name: bshr.BusinessHrs.EmpRef.Name}

		}

		if p == 2 {

			//Unmarshelling the json response in Employee1 struct jsut created.
			BusinessHourfb := Timeactivity1{}
			jsonErr := json.Unmarshal(body, &BusinessHourfb)
			if jsonErr != nil {
				log.Fatal(jsonErr)
			}

			returnBHrs = pb.BusinessHour{Description: BusinessHourfb.TimeEntryFB.IdFBTimeEntry}

		}
		fmt.Printf(string(body))
		return &returnBHrs, nil
	}
	return &returnBHrs, nil
}

//Delete businesshour ###############################################################################################################################################################33
//only implimented for Quickbook.
func (s *accountyServer) DeleteBusinessHours(ctx context.Context, in *pb.DeleteBusinessHourRequest) (*empty.Empty, error) {

	empl, err := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id", "app_type"}, pb.AccoutingEmployeeLinkAppointyIdEq{in.AppointyEmployeeId})

	//extract employee first name and last name from appointy database and pass it to the below string
	urlqb := "https://SANDBOX-QUICKBOOKS.API.INTUIT.COM/v3/company/" + in.CompanyId + "/timeactivity?operation=delete"
	payloadqbstr := "{\n    \"Id\": \"" + empl.ExternalId + "\",\n    \"SyncToken\": \"1\"\n}"

	//calling the API
	//token := "Bearer eyJlbmMiOiJBMTI4Q0JDLUhTMjU2IiwiYWxnIjoiZGlyIn0..IWcaW7YI7XSDvYH7M3Pq-Q.MOJCtljnKVE23wvhEo9nmnKujtHkVDCJi64HNJl-kr1EvP7P4jsGwt_VipR35J9DrURTqsBwcGmWH7rChnly9m9oEjkoFEySGDb5x8E7f5PunI7cyvOGocoVBCaE16GSnfoL0nzvKOTtCF2LBEodGhnDg5AinBJnfp_tFOom9oTJHYI2q5Do2m1KEWsjSzxtg6iMvNa0naDUjvzJPfQIuLXnoiwQbrR-O_TNBowV7EZdpvvqZpE8BDsNwtTNpFgLQduxM0n4ASMAIqP589nHAF1a7h-cX0cyID54IlU92Jyv5cVPJzOMOphlubgxQ0KpfOCkzIRjiqExRp-GLD8n_XSlmtQqWtUF3eSpPnhTcTqoTsQYRzHzV-_VGq8gnUqv9naU59KzIoc_F8ZtTG2_FVy6NsIhdQlkNone-ZFUvmHuAO9ZMP8iqKuV3D_Bo9Y0v0T51OPEsclWhWaUhLH0nGBXVfbV1WKJ55I87iMJDuOKdzRHwwTw_hBHVio6P_2DiZUSOrdWlV7IBppsQ3XXOghTEyBI1ruF1zsooEaTx04Imijzv8-1L3dnPQyTfJzoIy3x-mkI5Z1dwq9mkvOBBPwGPHJ-XZa3ny8_bQaKAKCeZ4swHJFWoNO0u6tChveOZcm1bmmGdGDKfP4kyn-s4baX2Fx3IvCFtOFMGlpxmT_s2QPWj2OrOX1oTvmLxtPNHeOIQm4GKKn9TlUDscjgT01gLGfL-K3mS3UdxCTgGr2B9W3CiPKN6X-mqd0WaR22lM3UZFhbG3DLUTvBLIU7PbP2vGcuWWaRWxPFYRz_tTmmhrte84mxpO7hXr5yEqFu.0xm1_AVBzBGaRBNVW3gZhQ"
	body, Ferr := ConnectApiEndpoint(urlqb, payloadqbstr, tokenqb)

	if Ferr != nil {
		log.Fatal(Ferr)
	}
	fmt.Println(string(body))
	return nil, err

}

//Add Staff Appointments ############################################################################################################################################

func (s *accountyServer) CreateStaffAppointments(ctx context.Context, in *pb.CreateStaffAppointmentRequest) (*pb.PrepaidAppointment, error) {
	StaffAppointmentReturn := pb.PrepaidAppointment{Id: "No account found on Quickbook, Xero or Freshbook"}

	CustId, err := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{in.StaffAppointment.AppointyId}})
	VendorAppointyId := "Vendor" + in.StaffAppointment.AppointyId
	VenCustId, err1 := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{VendorAppointyId}})
	if CustId == nil || VenCustId == nil {
		if err != nil || err1 != nil {
			return nil, errInternal
		}
		log.Println("Customer donot exist in database.")
		StaffAppointmentReturn = pb.PrepaidAppointment{Id: "Customer donot exist in database."}
	}
	//Adding sale or amount to income account
	SaleReciptLineId := "1111" + in.StaffAppointment.AppointyId
	id, SaleReciptErr := SaleRecipt(in.CompanyID, in.StaffAppointment.AppointyId, "Checking", SaleReciptLineId, "adding sale amount to income account for booking staff appointment", in.StaffAppointment.TotalAmount, s.store)
	if SaleReciptErr != nil {
		log.Println("error! working with salerecipt API.")
		StaffAppointmentReturn = pb.PrepaidAppointment{Id: "error! working with salerecipt API."}
	}
	fmt.Println(id)

	//Adding amount to Expense account which now we use as advertising but will map with user choice expense type acount in the production
	id, BillErr := Bill(in.CompanyID, in.StaffAppointment.AppointyId, "Advertising", in.StaffAppointment.TotalAmount, VenCustId.ExternalId, s.store)
	if BillErr != nil {
		log.Println("error! working with Bll API.")
		StaffAppointmentReturn = pb.PrepaidAppointment{Id: "error! working with Bill API."}
	}

	fmt.Println(id)
	if in.StaffAppointment.PolicyApplied == true {
		//removing amount from AP (liability).
		AdvBillAppointyId := "AdvanceBill" + in.StaffAppointment.AppointyId
		AdvBillId, AdvBillerr := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{AdvBillAppointyId}})
		if AdvBillId == nil {
			if AdvBillerr != nil {
				return nil, errInternal
			}
		}
		_, BillDelErr := BillDelete(in.CompanyID, AdvBillId.ExternalId)
		if BillDelErr != nil {
			StaffAppointmentReturn = pb.PrepaidAppointment{Id: "Bill Delete API not working properly. Error!"}
		}
	}
	if in.StaffAppointment.PolicyApplied == false {
		//adding payment in AR (cash) using invoice API
		id, PayErr := InvoicePayment(in.CompanyID, in.StaffAppointment.AppointyId, in.StaffAppointment.TotalAmount, CustId.ExternalId, s.store)
		fmt.Println(id)
		if PayErr != nil {
			log.Println("error! working with invoice API.")
			StaffAppointmentReturn = pb.PrepaidAppointment{Id: "error! working with invoice API."}
		}

		fmt.Println(id)
	}
	//Adding amount to
	StaffAppointmentReturn = pb.PrepaidAppointment{Id: "Staff appointment created. AP account, AR account and income account (Advertising) modified successfully."}
	return &StaffAppointmentReturn, nil
}

//Add prepaid appointment###########################################################################################################################################3
func (s *accountyServer) CreatePrepaidAppointment(ctx context.Context, in *pb.CreatePrepaidAppointmentRequest) (*pb.PrepaidAppointment, error) {

	returnPPA := pb.PrepaidAppointment{Id: "No account found on Quickbook, Xero or Freshbook"}
	ppa, err := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{in.PrepaidAppointment.AppointyId}})
	VendorAppointyId := "Vendor" + in.PrepaidAppointment.AppointyId
	ppa1, err1 := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{VendorAppointyId}})

	//If any AppNo represented by this index is not found in our database against any AppointyID then skip this index.
	if ppa == nil || ppa1 == nil {
		if err != nil || err1 != nil {
			return nil, errInternal
		}
		fmt.Println("Account not found or registered.")
	}

	//Payment added as cash and payment Id saved in chaku db
	id, PayErr := InvoicePayment(in.CompanyId, in.PrepaidAppointment.AppointyId, in.PrepaidAppointment.TotalAmount, ppa.ExternalId, s.store)
	if PayErr != nil {
		log.Println("InvoicePayment API not working.")
	}
	fmt.Println(id)

	//Create liability or add advance and save the billId generated in the chaku database
	id, BillErr := Bill(in.CompanyId, in.PrepaidAppointment.AppointyId, "Accounts Payable (A/P)", in.PrepaidAppointment.TotalAmount, ppa1.ExternalId, s.store)
	if BillErr != nil {
		log.Println("bill API not working.")
	}

	returnPPA = pb.PrepaidAppointment{Id: "Prepaid appointment created and AP and AR accounts modified"}

	return &returnPPA, nil

}

//cancel and refund of the prepaid appointments####################################################################################################################################################################################

func (s *accountyServer) CancelNRefPrepaidAppointment(ctx context.Context, in *pb.CancelNRefPrepaidAppointmentRequest) (*pb.PrepaidAppointment, error) {

	returnPPA := pb.PrepaidAppointment{Id: "No appointment made with this business or the payment donot match."}

	CustomerId, err := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{in.PrepaidAppointment.AppointyId}})
	BillAppointyId := "Bill" + in.PrepaidAppointment.AppointyId
	BillId, err1 := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{BillAppointyId}})
	if CustomerId == nil || BillId == nil {
		if err != nil || err1 != nil {
			return nil, errInternal
		}
		fmt.Println("No Id registered in our database.")

	}

	//Reduce from cash using payment API
	id1, PayErr := Payment(in.CompanyId, in.PrepaidAppointment.AppointyId, "Accounts Receivable (A/R)", in.PrepaidAppointment.TotalAmount, CustomerId.ExternalId, s.store)
	if PayErr != nil {
		log.Println("Payment API not working.")
	}
	fmt.Println(id1)
	//Remove liability or subtract advance
	body2, _ := BillDelete(in.CompanyId, BillId.ExternalId)
	fmt.Println(string(body2))
	//putting cancellation charges into other income using sale recipt api and saving SaleRecipt Id for further use
	SaleReciptId := "1111" + in.PrepaidAppointment.AppointyId
	id, _ := SaleRecipt(in.CompanyId, in.PrepaidAppointment.AppointyId, "Checking", SaleReciptId, "cancellation charge", in.PrepaidAppointment.TotalAmount, s.store)
	fmt.Print(id)

	returnPPA = pb.PrepaidAppointment{Id: "Prepaid Appointment cancelled and amount Refunded. A/P Account, A/R account and Checking account (SaleReciptID : " + id + " ) modified."}

	return &returnPPA, nil

}

//Update prepaid appointment #######################################################################################################################################################################################################

func (s *accountyServer) UpdatePrepaidAppointment(ctx context.Context, in *pb.UpdatePrepaidAppointmentRequest) (*pb.PrepaidAppointment, error) {

	//fetching all the required ids for making the updates.
	returnUPA := pb.PrepaidAppointment{Id: "No account synced with Accounty API"}

	CustomerId, err := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{in.PrepaidAppointment.AppointyId}})
	CustVendId := "Vendor" + in.PrepaidAppointment.AppointyId
	CustomerAsVendorId, err1 := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{CustVendId}})
	if CustomerId == nil || CustomerAsVendorId == nil {
		if err != nil || err1 != nil {
			return nil, errInternal
		}
		returnUPA = pb.PrepaidAppointment{Id: "No payment or bill history found related to this prepaid appointment."}
	}

	//Putting modification charges into other income account
	SaleReciptId := "1111" + in.PrepaidAppointment.AppointyId
	id, SaleErr := SaleRecipt(in.CompanyId, in.PrepaidAppointment.AppointyId, "Checking", SaleReciptId, "modification charge", in.PrepaidAppointment.TotalAmount, s.store)
	if SaleErr != nil {
		returnUPA = pb.PrepaidAppointment{Id: "SaleRecipt API not working properly"}
	}
	fmt.Println(id)

	//Add or Remove the difference from cash account (AR) and liability account AP (if it was created at the time of appointment) depending on the modification made
	SaleBefore := in.PrepaidAppointment.Metadata["SaleBefore"]
	SaleAfter := in.PrepaidAppointment.Metadata["SaleAfter"]
	if SaleAfter > SaleBefore {
		//Add the difference to cash account (AR) and if liability was created add amount to AP.
		DiffInSale := string(SaleAfter - SaleBefore)
		id, InPayErr := InvoicePayment(in.CompanyId, in.PrepaidAppointment.AppointyId, DiffInSale, CustomerId.ExternalId, s.store)
		if InPayErr != nil {
			log.Println("Invoice API not working.")
		}
		fmt.Println(id)

		//fetch the liability id from chaku db using its correspoding AppointyId
		LiabilityAppointyId := "Bill" + in.PrepaidAppointment.AppointyId
		LiabilityID, err3 := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{LiabilityAppointyId}})
		if LiabilityID == nil {
			if err3 != nil {
				return nil, errInternal
			}
			returnUPA = pb.PrepaidAppointment{Id: "Modification charge added in the checking account, cash added to AR account and no liability was created."}
			return &returnUPA, nil
		}

		//Add amount to liability account.
		id, BillErr := Bill(in.CompanyId, in.PrepaidAppointment.AppointyId, "Accounts Payable (A/P)", DiffInSale, CustomerAsVendorId.ExternalId, s.store)
		if BillErr != nil {
			log.Println("bill API not working.")
		}
		fmt.Println(id)
		returnUPA = pb.PrepaidAppointment{Id: "Modification charge added in the checking account(income account), cash added to AR account and AP account (liability)."}

	} else if SaleAfter < SaleBefore {
		DiffInSale := string(SaleBefore - SaleAfter)
		//Remove the difference from (AR) account using payment API
		id, PayErr := Payment(in.CompanyId, in.PrepaidAppointment.AppointyId, "Accounts Receivable (A/R)", DiffInSale, CustomerId.ExternalId, s.store)
		if PayErr != nil {
			log.Println("Payment API not working.")
		}
		fmt.Println(id)
		//Remove amount from liability (AP)account or delete Bill by fetching bill Id.
		body2, _ := BillDelete(in.CompanyId, CustomerAsVendorId.ExternalId)
		fmt.Println(string(body2))
	} else {
		returnUPA = pb.PrepaidAppointment{Id: "No modification made in the prepaid appointment."}
	}
	return &returnUPA, nil
}

//No show in a prepaid appointment ###############################################################################################################################################################################################

func (s *accountyServer) NoShowPrepaidAppointment(ctx context.Context, in *pb.NoShowPrepaidAppointmentRequest) (*pb.PrepaidAppointment, error) {
	returnNSPPA := pb.PrepaidAppointment{Id: "No account synced with accounty API."}

	//fetching all the required ids for making the updates.
	CustomerId, err0 := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{in.PrepaidAppointment.AppointyId}})
	BillAppointyId := "Bill" + in.PrepaidAppointment.AppointyId
	BillId, err1 := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{BillAppointyId}})
	BillForModificAppointyId := "Billmodified" + in.PrepaidAppointment.AppointyId
	ModifiedBillId, err2 := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{BillForModificAppointyId}})
	if CustomerId == nil || BillId == nil || ModifiedBillId == nil {
		if err0 != nil || err1 != nil || err2 != nil {
			return nil, errInternal
		}
		returnNSPPA = pb.PrepaidAppointment{Id: "No bill payment was found in our database."}
	}
	//Remove amount from liability (AP) account by deleting bill and modifiedbill (created if prepaid appointment modified)
	body1, BillErr1 := BillDelete(in.CompanyId, BillId.ExternalId)
	if BillErr1 != nil {
		returnNSPPA = pb.PrepaidAppointment{Id: "BillDelete API not working properly. Error!"}
	}
	fmt.Println(string(body1))
	body2, BillErr2 := BillDelete(in.CompanyId, ModifiedBillId.ExternalId)
	if BillErr2 != nil {
		returnNSPPA = pb.PrepaidAppointment{Id: "BillDelete API not working properly. Error!"}
	}
	fmt.Println(string(body2))

	//Reduce amount from cash (AR) account using payment.
	id1, PayErr := Payment(in.CompanyId, in.PrepaidAppointment.AppointyId, "Accounts Receivable (A/R)", in.PrepaidAppointment.TotalAmount, CustomerId.ExternalId, s.store)
	if PayErr != nil {
		log.Println("Payment API not working.")
	}
	fmt.Println(id1)
	//add no show charge amount to checking account by sale recipt.
	SaleReciptId := "1111" + in.PrepaidAppointment.AppointyId
	NoShowCharge := in.PrepaidAppointment.MetadataString["NoShowCharge"]
	id2, SaleErr := SaleRecipt(in.CompanyId, in.PrepaidAppointment.AppointyId, "Checking", SaleReciptId, "No show for prepaid appointment charge", NoShowCharge, s.store)
	if SaleErr != nil {
		returnNSPPA = pb.PrepaidAppointment{Id: "SaleRecipt API not working properly"}
	}
	fmt.Println(id2)
	return &returnNSPPA, nil
}

//Create Partpaid appointment ########################################################################################################################################################################################################

func (s *accountyServer) CreatePartpaidAppointment(ctx context.Context, in *pb.CreatePartpaidAppointmentRequest) (*pb.PartpaidAppointment, error) {

	returnPrtPA := pb.PartpaidAppointment{Id: "No account synced with accounty API!"}
	//Fetch required IDs for modifying involved accounts
	CustomerId, err1 := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{in.PartpaidAppointment.AppointyId}})
	CustVendId := "Vendor" + in.PartpaidAppointment.AppointyId
	CustomerAsVendorId, err2 := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{CustVendId}})
	if CustomerId == nil || CustomerAsVendorId == nil {
		if err1 != nil || err2 != nil {
			return nil, errInternal
		}
		returnPrtPA = pb.PartpaidAppointment{Id: "No Customer found with this synced with accounty API!"}
	}

	//Add advance to the AP (liability) using bill API
	id1, BillErr := Bill(in.CompanyId, in.PartpaidAppointment.AppointyId, "Accounts Payable (A/P)", in.PartpaidAppointment.TotalAmount, CustomerAsVendorId.ExternalId, s.store)
	if BillErr != nil {
		returnPrtPA = pb.PartpaidAppointment{Id: "Bill API not working properly."}
	}
	fmt.Println(id1)

	//Add amount to AR (cash) account using invoice API
	id2, InPayErr := InvoicePayment(in.CompanyId, in.PartpaidAppointment.AppointyId, in.PartpaidAppointment.TotalAmount, CustomerId.ExternalId, s.store)
	if InPayErr != nil {
		log.Println("error! working with invoice API.")
		returnPrtPA = pb.PartpaidAppointment{Id: "error! working with invoice API."}
	}
	fmt.Println(id2)
	return &returnPrtPA, nil
}

//Cancel and refund Partpaid appointment ########################################################################################################################################################################################################
func (s *accountyServer) CancelNRefPartpaidAppointment(ctx context.Context, in *pb.CancelNRefPartpaidAppointmentRequest) (*pb.PartpaidAppointment, error) {

	returnCNRPrtPA := pb.PartpaidAppointment{Id: "No account synced with accounty API!"}
	//Fetch required IDs for modifying involved accounts
	BillAppointyId := "BillPrtPA" + in.PartpaidAppointment.AppointyId
	BillId, err1 := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{BillAppointyId}})
	if BillId == nil {
		if err1 != nil {
			return nil, errInternal
		}
	}
	//Remove amount from AP (liability) account using delete bill API
	body1, BillErr1 := BillDelete(in.CompanyId, BillId.ExternalId)
	fmt.Println(string(body1))
	if BillErr1 != nil {
		returnCNRPrtPA = pb.PartpaidAppointment{Id: "BillDelete API not working properly. Error!"}
	}

	//Add refund charge amount to other income account (checking) using sale recipt API
	SaleReciptId := "1111" + in.PartpaidAppointment.AppointyId
	CancelNRefundCharge := in.PartpaidAppointment.MetadataString["CancelNRefundCharge"]
	id1, SaleErr := SaleRecipt(in.CompanyId, in.PartpaidAppointment.AppointyId, "Checking", SaleReciptId, "Cancel N Refund charge of partpaid appointment charge", CancelNRefundCharge, s.store)
	if SaleErr != nil {
		returnCNRPrtPA = pb.PartpaidAppointment{Id: "SaleRecipt API not working properly"}
	}
	fmt.Println(id1)

	//Remove amount from AR (cash) account using Cash Memo API used for refund.

	returnCNRPrtPA = pb.PartpaidAppointment{Id: "Partpaid appointment cancelled and amount refunded. AP(liability), other income (checking) and AR(cash) account modified"}
	return &returnCNRPrtPA, nil

}

//Update Partpaid appointment ########################################################################################################################################################################################################
func (s *accountyServer) UpdatePartpaidAppointment(ctx context.Context, in *pb.UpdatePartpaidAppointmentRequest) (*pb.PartpaidAppointment, error) {

	returnUPrtPA := pb.PartpaidAppointment{Id: "No account synced with accounty API!"}
	//Fetch required IDs for modifying involved accounts
	CustomerId, err := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{in.PartpaidAppointment.AppointyId}})
	CustVendId := "Vendor" + in.PartpaidAppointment.AppointyId
	CustomerAsVendorId, err1 := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{CustVendId}})
	if CustomerId == nil || CustomerAsVendorId == nil {
		if err != nil || err1 != nil {
			return nil, errInternal
		}
		returnUPrtPA = pb.PartpaidAppointment{Id: "No payment or bill history found related to this prepaid appointment."}
	}

	//Add amount to other income (checking) account using sale recipt API
	SaleReciptId := "1111" + in.PartpaidAppointment.AppointyId
	UpdatePrtPACharge := in.PartpaidAppointment.MetadataString["UpdateCharge"]
	id1, SaleErr := SaleRecipt(in.CompanyId, in.PartpaidAppointment.AppointyId, "Checking", SaleReciptId, "Modification of partpaid appointment charge", UpdatePrtPACharge, s.store)
	if SaleErr != nil {
		returnUPrtPA = pb.PartpaidAppointment{Id: "SaleRecipt API not working properly"}
	}
	fmt.Println(id1)

	//Add or Remove the difference from cash account (AR) and liability account AP (if it was created at the time of appointment) depending on the modification made
	SaleBefore := in.PartpaidAppointment.Metadata["SaleBefore"]
	SaleAfter := in.PartpaidAppointment.Metadata["SaleAfter"]
	if SaleAfter > SaleBefore {
		//Add the difference to cash account (AR) and if liability was created add amount to AP.
		DiffInSale := string(SaleAfter - SaleBefore)
		id2, InPayErr := InvoicePayment(in.CompanyId, in.PartpaidAppointment.AppointyId, DiffInSale, CustomerId.ExternalId, s.store)
		if InPayErr != nil {
			log.Println("Invoice Payment API not working.")
		}
		fmt.Println(id2)

		//fetch the liability id from chaku db using its correspoding AppointyId
		LiabilityAppointyId := "Bill" + in.PartpaidAppointment.AppointyId
		LiabilityID, err2 := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{LiabilityAppointyId}})
		if LiabilityID == nil {
			if err2 != nil {
				return nil, errInternal
			}
			returnUPrtPA = pb.PartpaidAppointment{Id: "Modification charge added in the checking account, cash added to AR account and no liability was found."}
			return &returnUPrtPA, nil
		}

		//Add amount to liability account.
		id3, BillErr := Bill(in.CompanyId, in.PartpaidAppointment.AppointyId, "Accounts Payable (A/P)", DiffInSale, CustomerAsVendorId.ExternalId, s.store)
		if BillErr != nil {
			log.Println("bill API not working.")
		}
		fmt.Println(id3)
		returnUPrtPA = pb.PartpaidAppointment{Id: "Modification charge added in the checking account(income account), cash added to AR account and AP account (liability)."}

	} else if SaleAfter < SaleBefore {
		DiffInSale := string(SaleBefore - SaleAfter)
		//Remove the difference from AR cash account using payment API
		id4, PayErr := Payment(in.CompanyId, in.PartpaidAppointment.AppointyId, "Accounts Recieved (A/R)", DiffInSale, CustomerId.ExternalId, s.store)
		if PayErr != nil {
			returnUPrtPA = pb.PartpaidAppointment{Id: "Payment API not working properly."}
		}
		fmt.Println(id4)

		//Remove amount from liability (AP)account or delete Bill by fetching bill Id.
		body5, DelBillErr := BillDelete(in.CompanyId, CustomerAsVendorId.ExternalId)
		if DelBillErr != nil {
			returnUPrtPA = pb.PartpaidAppointment{Id: "Delete bill API not working."}
		}
		fmt.Println(string(body5))
	} else {
		returnUPrtPA = pb.PartpaidAppointment{Id: "Partpay appointment modified. Cash account (AR), Liability account (AP) and other income (checking) account modified."}
	}
	return &returnUPrtPA, nil

}

//No show case in Partpaid appointment ########################################################################################################################################################################################################
func (s *accountyServer) NoShowPartpaidAppointment(ctx context.Context, in *pb.NoShowPartpaidAppointmentRequest) (*pb.PartpaidAppointment, error) {

	returnNSPrtPA := pb.PartpaidAppointment{Id: "No account synced with accounty API!"}
	//Fetch required IDs for account modification.
	CustomerId, err := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{in.PartpaidAppointment.AppointyId}})
	CustVendId := "Vendor" + in.PartpaidAppointment.AppointyId
	CustomerAsVendorId, err1 := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{CustVendId}})
	if CustomerId == nil || CustomerAsVendorId == nil {
		if err != nil || err1 != nil {
			return nil, errInternal
		}
		returnNSPrtPA = pb.PartpaidAppointment{Id: "No Customer found related to this prepaid appointment."}
	}

	//Remove amount from liability (AP) account using Delete bill API
	LiabilityAppointyId := "Bill" + in.PartpaidAppointment.AppointyId
	LiabilityID, err2 := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{LiabilityAppointyId}})
	if LiabilityID == nil {
		if err2 != nil {
			return nil, errInternal
		}
		returnNSPrtPA = pb.PartpaidAppointment{Id: "No liability found related to this partpaid appointment."}
	} else {
		body1, DelBillErr := BillDelete(in.CompanyId, LiabilityID.ExternalId)
		if DelBillErr != nil {
			returnNSPrtPA = pb.PartpaidAppointment{Id: "Delete bill API not working."}
		}
		fmt.Println(string(body1))
	}
	//Remove amount from cash (AR) account using credit memo API
	id1, PayErr := Payment(in.CompanyId, in.PartpaidAppointment.AppointyId, "Accounts Recieved (A/R)", in.PartpaidAppointment.TotalAmount, CustomerId.ExternalId, s.store)
	if PayErr != nil {
		returnNSPrtPA = pb.PartpaidAppointment{Id: "Payment API not working properly."}
	}
	fmt.Println(id1)

	//Add amount to other income (Checking) account using Sale Recipt API
	SaleReciptId := "1111" + in.PartpaidAppointment.AppointyId
	NoShowPrtPACharge := in.PartpaidAppointment.MetadataString["NoShowCharge"]
	id2, SaleErr := SaleRecipt(in.CompanyId, in.PartpaidAppointment.AppointyId, "Checking", SaleReciptId, "Modification of partpaid appointment charge", NoShowPrtPACharge, s.store)
	if SaleErr != nil {
		returnNSPrtPA = pb.PartpaidAppointment{Id: "SaleRecipt API not working properly"}
	}
	fmt.Println(id2)

	returnNSPrtPA = pb.PartpaidAppointment{Id: "No show for partpaid appointment. AP account AR account and other income (checking) account modified."}
	return &returnNSPrtPA, nil

}

//create an inventory/addon ######################################################################################################################################################################################################################################################################################################

/*func (s *accountyServer) CreateInventory(ctx context.Context, in *pb.CreateInventoryRequest) (*pb.Inventory, error) {

	returnAddon := pb.Inventory{Id: "No App synced with accounty service."}
	ppa, err := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{in.Inventory.VendorAppointyId}})
	CustIdasVen := "Vendor" + in.Inventory.VendorAppointyId
	ppa2, err2 := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{CustIdasVen}})
	if ppa == nil || ppa2 == nil {
		if err != nil || err2 != nil {
			return nil, errInternal
		}
	}
	if in.Inventory.PaymentDone == true {
		//Add item to inventory and reduce amount from cash
		body1, ItemErr := Item(in.CompanyID, in.Inventory.Name, ppa.ExternalId, in.Inventory.Qnt, in.Inventory.ItemCost, 1)
		if ItemErr != nil {
			returnAddon = pb.Inventory{InventoryId: "Error! Item API not working properly"}
		}
		fmt.Println(string(body1))
		newadd := Payment1{}
		jsonErr := json.Unmarshal(body1, &newadd)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		ItemAppointyId := "Item" + in.Inventory.VendorAppointyId
		_, errDB1 := s.store.CreateAccoutingEmployeeLinks(ctx, &pb.AccoutingEmployeeLink{
			AppointyId: ItemAppointyId,
			ExternalId: newadd.ItemQb.Id,
			AppType:    1,
		})
		if errDB1 != nil {
			return nil, errInternal
		}

	}
	if in.Inventory.PaymentDone == false {
		//Add item to inventory
		body1, ItemErr := Item(in.CompanyID, in.Inventory.Name, ppa2.ExternalId, in.Inventory.Qnt, in.Inventory.ItemCost, 2)
		if ItemErr != nil {
			returnAddon = pb.Inventory{InventoryId: "ErrOr! Item API not working properly"}
		}
		fmt.Println(string(body1))
		newadd := Payment1{}
		jsonErr := json.Unmarshal(body1, &newadd)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		ItemAppointyId := "Item" + in.Inventory.VendorAppointyId
		_, errDB1 := s.store.CreateAccoutingEmployeeLinks(ctx, &pb.AccoutingEmployeeLink{
			AppointyId: ItemAppointyId,
			ExternalId: newadd.ItemQb.Id,
			AppType:    1,
		})
		if errDB1 != nil {
			return nil, errInternal
		}

	}
	return &returnAddon, nil

}

//Remove from inventory (AS in the case of loss of item)##########################################################################################################################################################################################################

func (s *accountyServer) RemoveInventory(ctx context.Context, in *pb.RemoveInventoryRequest) (*pb.Inventory, error) {

	returnRemoveAddon := pb.Inventory{Id: "No App synced with accounty service."}
	//Fetch all the required IDs for account modification.
	ppa, err := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{in.Inventory.VendorAppointyId}})
	CustIdasVen := "Vendor" + in.Inventory.VendorAppointyId
	ppa2, err2 := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{CustIdasVen}})
	if ppa == nil || ppa2 == nil {
		if err != nil || err2 != nil {
			return nil, errInternal
		}
	}

	//Delete from inventory asset account.

	//Item returned or sold
	if StateOfItem == true {
		if PaymentDone == true {
			//Add to cash (AR) case of returned or sold using invoice API
		}
		if PaymentDone == false {
			//Remove amount from liability AP account using Delete bill API
		}
	}

	//item lost and total loss occoured.
	if StateOfItem == false {
		//Add amount to loss account (created at the time of account syncing)

	}
}*/

//Remove from inventory ##########################################################################################################################################################################################################

func (s *accountyServer) SaleNRemoveInventory(ctx context.Context, in *pb.SaleNRemoveInventoryRequest) (*pb.Inventory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaleNRemoveInventory not implemented")
}

//Remove from inventory ##########################################################################################################################################################################################################

func (s *accountyServer) ReturnNAddInventory(ctx context.Context, in *pb.ReturnNAddInventoryRequest) (*pb.Inventory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReturnNAddInventory not implemented")
}

//Sell gift certificate #########################################################################################################################################################################################################

func (s *accountyServer) SellGiftCertificate(ctx context.Context, in *pb.SellGiftCertificateRequest) (*pb.PrepaidAppointment, error) {

	returnSGC := pb.PrepaidAppointment{Id: "No accont synced with accounty API."}
	//fetch IDs for required accounts modification.
	CustomerId, err := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{in.PrepaidAppointment.AppointyId}})
	CustVendId := "Vendor" + in.PrepaidAppointment.AppointyId
	CustomerAsVendorId, err1 := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{CustVendId}})
	if CustomerId == nil || CustomerAsVendorId == nil {
		if err != nil || err1 != nil {
			return nil, errInternal
		}
		returnSGC = pb.PrepaidAppointment{Id: "No Customer found related to this Gift Certificate."}
	}

	//Add amount to cash account AR using Invoice API
	id1, InPayErr := InvoicePayment(in.CompanyId, in.PrepaidAppointment.AppointyId, in.PrepaidAppointment.TotalAmount, CustomerId.ExternalId, s.store)
	if InPayErr != nil {
		log.Println("Invoice Payment API not working.")
	}
	fmt.Println(id1)

	//Add amount to AP (liability) account
	id2, BillErr := Bill(in.CompanyId, in.PrepaidAppointment.AppointyId, "Accounts Payable (A/P)", in.PrepaidAppointment.TotalAmount, CustomerAsVendorId.ExternalId, s.store)
	if BillErr != nil {
		log.Println("bill API not working.")
	}
	fmt.Println(id2)
	returnSGC = pb.PrepaidAppointment{Id: "Gift Certificate sold. AP account and AR account modified."}
	return &returnSGC, nil
}

//Redeem gift certificates ##################################################################################################################################################################################################################

func (s *accountyServer) RedeemGiftCertificate(ctx context.Context, in *pb.RedeemGiftCertificateRequest) (*pb.PrepaidAppointment, error) {

	ReturnRGC := pb.PrepaidAppointment{Id: "Sorry! No account synced with accounty API."}
	//Fetching required Ids for account modifications

	//adding amount to sale account (checking)
	SaleReciptLineId := "1111" + in.PrepaidAppointment.AppointyId
	id1, SaleReciptErr := SaleRecipt(in.CompanyId, in.PrepaidAppointment.AppointyId, "Checking", SaleReciptLineId, "adding sale amount to income account for redeeming gift certificate. ", in.PrepaidAppointment.TotalAmount, s.store)
	if SaleReciptErr != nil {
		ReturnRGC = pb.PrepaidAppointment{Id: "error! working with salerecipt API."}
	}
	fmt.Println(id1)

	//Remove liability if created
	LiabilityAppointyId := "Bill" + in.PrepaidAppointment.AppointyId
	LiabilityID, err2 := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{LiabilityAppointyId}})
	if LiabilityID == nil {
		if err2 != nil {
			return nil, errInternal
		}
		ReturnRGC = pb.PrepaidAppointment{Id: "No liability found related to this Gift certificate."}
	} else {
		body1, DelBillErr := BillDelete(in.CompanyId, LiabilityID.ExternalId)
		if DelBillErr != nil {
			ReturnRGC = pb.PrepaidAppointment{Id: "Delete bill API not working."}
		}
		fmt.Println(string(body1))
	}
	ReturnRGC = pb.PrepaidAppointment{Id: "Gift certificate redeemed. Sale account (checking), AP account modified accordingly."}
	return &ReturnRGC, nil
}

//Cancel Gift certificates #################################################################################################################################################################################################################################

func (s *accountyServer) CancelGiftCertificate(ctx context.Context, in *pb.CancelGiftCertificateRequest) (*pb.PrepaidAppointment, error) {

	ReturnCGC := pb.PrepaidAppointment{Id: "No App synced with accounty service of appointy"}
	//Fetching Ids required for modifying accounts involved
	CustomerId, err := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{in.PrepaidAppointment.AppointyId}})
	if CustomerId == nil {
		if err != nil {
			return nil, errInternal
		}
		ReturnCGC = pb.PrepaidAppointment{Id: "couldnot find the user in the database"}

	}
	//Adding processing fee amount to other income (checking) account using Sale recipt API
	//processing fee is the total fee reduced by refund amount
	SaleReciptLineId := "1111" + in.PrepaidAppointment.AppointyId
	CancellationCharge := string(in.PrepaidAppointment.Metadata["CancellationCharge"])
	id1, SaleReciptErr := SaleRecipt(in.CompanyId, in.PrepaidAppointment.AppointyId, "Checking", SaleReciptLineId, "adding cancellation charge to income account for cancelling gift certificate ", CancellationCharge, s.store)
	if SaleReciptErr != nil {
		ReturnCGC = pb.PrepaidAppointment{Id: "error! working with salerecipt API."}
	}
	fmt.Println(id1)

	//Reduce amount from liability account (AP) using bill delete API
	LiabilityAppointyId := "Bill" + in.PrepaidAppointment.AppointyId
	LiabilityID, err2 := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{LiabilityAppointyId}})
	if LiabilityID == nil {
		if err2 != nil {
			return nil, errInternal
		}
		ReturnCGC = pb.PrepaidAppointment{Id: "No liability found related to this Gift certificate."}
	} else {
		body1, DelBillErr := BillDelete(in.CompanyId, LiabilityID.ExternalId)
		if DelBillErr != nil {
			ReturnCGC = pb.PrepaidAppointment{Id: "Delete bill API not working."}
		}
		fmt.Println(string(body1))
	}

	//Reduce refunded amount from AR account (cash) using payment API
	RefundedAmt := string(in.PrepaidAppointment.Metadata["TotalAmount"] - in.PrepaidAppointment.Metadata["CancellationCharge"])
	id2, PayErr := Payment(in.CompanyId, in.PrepaidAppointment.AppointyId, "Accounts Recieved (A/R)", RefundedAmt, CustomerId.ExternalId, s.store)
	if PayErr != nil {
		ReturnCGC = pb.PrepaidAppointment{Id: "Payment API not working properly."}
	}
	fmt.Println(id2)
	ReturnCGC = pb.PrepaidAppointment{Id: "Gift certificate cancelled and payment refunded. Checking account, AP account and AR account modified."}
	return &ReturnCGC, nil

}

//Gift certificate expires ########################################################################################################################################################################################################

func (s *accountyServer) ExpireGiftCertificate(ctx context.Context, in *pb.ExpireGiftCertificateRequest) (*pb.PrepaidAppointment, error) {

	ReturnEGC := pb.PrepaidAppointment{Id: "No App synced with accounty service of appointy"}
	//Fetching required IDs for account modifications

	//Add amount to other income account (checking) using sale recipt
	SaleReciptLineId := "1111" + in.PrepaidAppointment.AppointyId
	ExpiryCharge := string(in.PrepaidAppointment.Metadata["ExpiryCharge"])
	id1, SaleReciptErr := SaleRecipt(in.CompanyId, in.PrepaidAppointment.AppointyId, "Checking", SaleReciptLineId, "adding Expiry charge to income account for gift certificate getting expired", ExpiryCharge, s.store)
	if SaleReciptErr != nil {
		ReturnEGC = pb.PrepaidAppointment{Id: "error! working with salerecipt API."}
	}
	fmt.Println(id1)
	ReturnEGC = pb.PrepaidAppointment{Id: "Gift certificate expired and checking account (other income) modified accordingly."}
	return &ReturnEGC, nil
}

//Add customer account###########################################################################################################################################################################################################
func (s *accountyServer) CreateCustomerAccount(ctx context.Context, in *pb.CreateCustomerRequest) (*pb.Customer, error) {

	sendcus := pb.Customer{Id: "No App synced with accounty service of appointy"}

	if in.AppType == 1 {
		urlqb := "https://SANDBOX-QUICKBOOKS.API.INTUIT.COM/v3/company/" + in.CompanyId + "/customer?minorversion=37"
		CountrySubDivisionCode := in.Customer.Metadata["countrycode"]
		City := in.Customer.Metadata["city"]
		BaseAddress := in.Customer.Metadata["baseaddr"]

		payloadqbstr := "{\n    \"BillAddr\": {\n        \"Line1\": \"" + BaseAddress + "\",\n        \"City\": \"" + City + "\",\n        \"Country\": \"" + CountrySubDivisionCode + "\",\n        \"CountrySubDivisionCode\": \"CA\",\n        \"PostalCode\": \"94042\"\n    },\n    \"Notes\": \"Here are other details.\",\n    \"DisplayName\": \"" + in.Customer.FirstName + "\",\n    \"PrimaryPhone\": {\n        \"FreeFormNumber\": \"" + in.Customer.PhoneNumber + "\"\n    }\n    \n}\n"
		//calling the API
		//token := "Bearer eyJlbmMiOiJBMTI4Q0JDLUhTMjU2IiwiYWxnIjoiZGlyIn0..IWcaW7YI7XSDvYH7M3Pq-Q.MOJCtljnKVE23wvhEo9nmnKujtHkVDCJi64HNJl-kr1EvP7P4jsGwt_VipR35J9DrURTqsBwcGmWH7rChnly9m9oEjkoFEySGDb5x8E7f5PunI7cyvOGocoVBCaE16GSnfoL0nzvKOTtCF2LBEodGhnDg5AinBJnfp_tFOom9oTJHYI2q5Do2m1KEWsjSzxtg6iMvNa0naDUjvzJPfQIuLXnoiwQbrR-O_TNBowV7EZdpvvqZpE8BDsNwtTNpFgLQduxM0n4ASMAIqP589nHAF1a7h-cX0cyID54IlU92Jyv5cVPJzOMOphlubgxQ0KpfOCkzIRjiqExRp-GLD8n_XSlmtQqWtUF3eSpPnhTcTqoTsQYRzHzV-_VGq8gnUqv9naU59KzIoc_F8ZtTG2_FVy6NsIhdQlkNone-ZFUvmHuAO9ZMP8iqKuV3D_Bo9Y0v0T51OPEsclWhWaUhLH0nGBXVfbV1WKJ55I87iMJDuOKdzRHwwTw_hBHVio6P_2DiZUSOrdWlV7IBppsQ3XXOghTEyBI1ruF1zsooEaTx04Imijzv8-1L3dnPQyTfJzoIy3x-mkI5Z1dwq9mkvOBBPwGPHJ-XZa3ny8_bQaKAKCeZ4swHJFWoNO0u6tChveOZcm1bmmGdGDKfP4kyn-s4baX2Fx3IvCFtOFMGlpxmT_s2QPWj2OrOX1oTvmLxtPNHeOIQm4GKKn9TlUDscjgT01gLGfL-K3mS3UdxCTgGr2B9W3CiPKN6X-mqd0WaR22lM3UZFhbG3DLUTvBLIU7PbP2vGcuWWaRWxPFYRz_tTmmhrte84mxpO7hXr5yEqFu.0xm1_AVBzBGaRBNVW3gZhQ"
		body, _ := ConnectApiEndpoint(urlqb, payloadqbstr, tokenqb)

		//Unmarshelling the json response in Employee1 struct jsut created.(doubt)
		cus1 := Customer1{}
		jsonErr := json.Unmarshal(body, &cus1)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		fmt.Println(string(body))

		//Inserting Id generated into chaku database
		ids, _ := s.store.CreateAccoutingEmployeeLinks(ctx, &pb.AccoutingEmployeeLink{
			AppointyId: in.Customer.AppointyId,
			ExternalId: cus1.QbCustomer.Id,
			AppType:    1})

		//making a vendor in the name of customer.
		urlqbvendor := "https://SANDBOX-QUICKBOOKS.API.INTUIT.COM/v3/company/" + in.CompanyId + "/vendor?minorversion=37"
		payloadqbvendor := "{\n    \"BillAddr\": {\n        \"Line1\": \"" + BaseAddress + "\",\n        \"City\": \"" + City + "\",\n        \"Country\": \"" + CountrySubDivisionCode + "\",\n        \"CountrySubDivisionCode\": \"CA\",\n        \"PostalCode\": \"94030\"\n    },\n    \"Title\": \"Ms/Mr.\",\n    \"GivenName\": \"Vendor" + in.Customer.FirstName + "\",\n    \"FamilyName\": \"" + in.Customer.LastName + "\",\n    \"Suffix\": \"Sr.\",\n    \n    \"DisplayName\": \"" + in.Customer.FirstName + " " + in.Customer.LastName + "\"\n}\n"
		body1, _ := ConnectApiEndpoint(urlqbvendor, payloadqbvendor, tokenqb)

		cus2 := Customer1{}
		jsonErr2 := json.Unmarshal(body1, &cus2)
		if jsonErr2 != nil {
			log.Fatal(jsonErr2)
		}
		fmt.Println(string(body1))
		VendorAppointyId := "Vendor" + in.Customer.AppointyId

		ids1, _ := s.store.CreateAccoutingEmployeeLinks(ctx, &pb.AccoutingEmployeeLink{
			AppointyId: VendorAppointyId,
			ExternalId: cus2.QbVendor.Id,
			AppType:    1})
		fmt.Println(ids1)

		//Returning poulated employee struct with Id of the database back to program for further use
		in.Customer.Id = ids[0]
		sendcus = pb.Customer{Id: in.Customer.Id}
		return &sendcus, nil
	}

	//	Xero employee add
	if in.AppType == 2 {
		urlqb := "https://api.xero.com/api.xro/2.0/Employees"
		payloadqbstr := "{\n  \"Employees\": [\n    {\n      \"FirstName\": \"" + in.Customer.FirstName + "\",\n      \"LastName\": \"" + in.Customer.LastName + "\"\n    }\n  ]\n}"
		//calling the API

		body, _ := ConnectToXeroAPI(urlqb, payloadqbstr)

		//Unmarshelling the json response in Employee1 struct jsut created.(doubt)
		/*Emp2 := Employee1{}
		jsonErr := json.Unmarshal(body, &Emp2)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}*/
		log.Println(string(body))
		//in.Employee.Id = Employee1.Id

		//Inserting Id generated into chaku database
		ids, _ := s.store.CreateAccoutingEmployeeLinks(ctx, &pb.AccoutingEmployeeLink{
			AppointyId: in.Customer.AppointyId,
			ExternalId: "4",
			AppType:    2,
		})

		//log.Println(len(ids))
		//Returning poulated employee struct with Id of the database back to program for further use
		in.Customer.Id = ids[0]
		sendcus = pb.Customer{Id: in.Customer.Id}
		return &sendcus, nil
	}
	return &sendcus, nil

}

//remove customer account #############################################################################################################################################################################################################

func (s *accountyServer) RemoveCustomerAccount(ctx context.Context, in *pb.RemoveCustomerRequest) (*empty.Empty, error) {

	empl, err := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id", "app_type"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{in.Customer.AppointyId}})
	if empl == nil {
		if err != nil {
			return nil, errInternal
		}
		fmt.Println("Account donot exist in Quickbook")
	}
	urlqb := "https://SANDBOX-QUICKBOOKS.API.INTUIT.COM/v3/company/" + in.CompanyId + "/customer"
	payloadqbstr := "{\n    \"domain\": \"QBO\",\n    \"sparse\": true,\n    \"Id\": \"" + empl.ExternalId + "\",\n   \n    \"Active\": false\n}"
	body, _ := ConnectApiEndpoint(urlqb, payloadqbstr, tokenqb)
	fmt.Printf(string(body))
	return nil, nil

}

//update customer account ########################################################################################################################################################################################################################################################################################################
func (s *accountyServer) UpdateCustomerAccount(ctx context.Context, in *pb.UpdateCustomerRequest) (*pb.Customer, error) {
	panic(`impliment me`)
}

//Create Manual Payment ###########################################################################################################################################################################################################################################################################################################

func (s *accountyServer) CreateManualPayment(ctx context.Context, in *pb.CreateManualPaymentRequest) (*pb.PartpaidAppointment, error) {

	returnCMP := pb.PartpaidAppointment{Id: "Sorry! no account synced with Accounty API"}
	//fetch Ids required for account modification
	CustomerId, err := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{in.PrepaidAppointment.AppointyId}})
	if CustomerId == nil {
		if err != nil {
			return nil, errInternal
		}
	}
	//Add amount to sale account (checking) using sale recipt API
	SaleReciptLineId := "1111" + in.PrepaidAppointment.AppointyId
	id1, SaleReciptErr := SaleRecipt(in.CompanyId, in.PrepaidAppointment.AppointyId, "Checking", SaleReciptLineId, "adding cancellation charge to income account for cancelling gift certificate ", in.PrepaidAppointment.TotalAmount, s.store)
	if SaleReciptErr != nil {
		returnCMP = pb.PartpaidAppointment{Id: "error! working with salerecipt API."}
	}
	fmt.Println(id1)

	//Add amount to cash account (AR) using invoice API
	id2, InPayErr := InvoicePayment(in.CompanyId, in.PrepaidAppointment.AppointyId, in.PrepaidAppointment.TotalAmount, CustomerId.ExternalId, s.store)
	if InPayErr != nil {
		log.Println("Invoice Payment API not working.")
	}
	fmt.Println(id2)
	returnCMP = pb.PartpaidAppointment{Id: "Manual payment added. sale account (checking) and cash account (AR) modified."}
	return &returnCMP, nil
}

//Remove manual payment ##########################################################################################################################################################################################################################################################################################################

func (s *accountyServer) RemoveManualPayment(ctx context.Context, in *pb.RemoveManualPaymentRequest) (*pb.PartpaidAppointment, error) {

	returnRMP := pb.PartpaidAppointment{Id: "No account synced with accounty API"}
	//fetch Ids required for account modification
	SaleReciptAccountyIdDB := "SaleRecipt" + in.PrepaidAppointment.AppointyId
	SaleReciptId, err := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{SaleReciptAccountyIdDB}})
	CustomerId, err1 := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{in.PrepaidAppointment.AppointyId}})
	if SaleReciptId == nil || CustomerId == nil {
		if err != nil || err1 != nil {
			return nil, errInternal
		}
	}
	//Remove amount from sale account (checking) using delete sale recipt API
	_, SaleDelErr := DeleteSaleRecipt(in.CompanyId, SaleReciptId.ExternalId)
	if SaleDelErr != nil {
		returnRMP = pb.PartpaidAppointment{Id: "Delete SaleRecipt API not working."}
		return &returnRMP, SaleDelErr
	}
	//Add cancellation fee to other income account (checking) using sale recipt API
	CancellationFee := string(in.PrepaidAppointment.Metadata["CancellationFee"])
	SaleReciptLineId := "1111" + in.PrepaidAppointment.AppointyId
	id1, SaleReciptErr := SaleRecipt(in.CompanyId, in.PrepaidAppointment.AppointyId, "Checking", SaleReciptLineId, "adding cancellation charge to income account for cancelling gift certificate ", CancellationFee, s.store)
	if SaleReciptErr != nil {
		returnRMP = pb.PartpaidAppointment{Id: "error! working with salerecipt API."}
	}
	fmt.Println(id1)
	return &returnRMP, nil
}

//Update manual payment ##########################################################################################################################################################################################################################################################################################################

func (s *accountyServer) UpdateManualPayment(ctx context.Context, in *pb.UpdateManualPaymentRequest) (*pb.PartpaidAppointment, error) {

	returnUMP := pb.PartpaidAppointment{Id: "No account synced with accounty API"}
	//fetch Ids required for account modification
	SaleReciptAccountyIdDB := "SaleRecipt" + in.PrepaidAppointment.AppointyId
	SaleReciptId, err := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{SaleReciptAccountyIdDB}})
	if SaleReciptId == nil {
		if err != nil {
			return nil, errInternal
		}
	}

	//Add payment modification charge to income account (checking) using sale recipt API
	ModificationFee := string(in.PrepaidAppointment.Metadata["ModificationFee"])
	SaleReciptLineId := "1111" + in.PrepaidAppointment.AppointyId
	id1, SaleReciptErr := SaleRecipt(in.CompanyId, in.PrepaidAppointment.AppointyId, "Checking", SaleReciptLineId, "adding modification charge to income account for modifying a payment", ModificationFee, s.store)
	if SaleReciptErr != nil {
		returnUMP = pb.PartpaidAppointment{Id: "error! working with salerecipt API."}
	}
	fmt.Println(id1)

	//AddOrRemove amount ToOrFrom income account (API) using invoice API accordingly
	NewPayment := string(in.PrepaidAppointment.Metadata["NewPayment"])

	id1, SaleReciptErr1 := SaleRecipt(in.CompanyId, in.PrepaidAppointment.AppointyId, "Checking", SaleReciptLineId, "adding new payment to Revenue", NewPayment, s.store)
	if SaleReciptErr1 != nil {
		returnUMP = pb.PartpaidAppointment{Id: "error! working with salerecipt API."}
	}
	fmt.Println(id1)

	//Remove old payment from sale account (checking) using delete sale recipt API
	_, SaleDelErr := DeleteSaleRecipt(in.CompanyId, SaleReciptId.ExternalId)
	if SaleDelErr != nil {
		returnUMP = pb.PartpaidAppointment{Id: "Delete SaleRecipt API not working."}
		return &returnUMP, SaleDelErr
	}
	return &returnUMP, nil
}

//Create Discount Coupon Enabled Sale ##########################################################################################################################################################################################################################################################################################################

func (s *accountyServer) CreateDiscountCouponEnabledSale(ctx context.Context, in *pb.CreateDiscountCouponEnabledSaleRequest) (*pb.PartpaidAppointment, error) {

	returnCDCES := pb.PartpaidAppointment{Id: "No app sysnced usnig accounty API."}
	//fetch Ids required for account modification
	CustVendId := "Vendor" + in.PrepaidAppointment.AppointyId
	CustomerAsVendorId, err1 := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{CustVendId}})
	CustomerId, err := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{in.PrepaidAppointment.AppointyId}})
	if CustomerAsVendorId == nil || CustomerId == nil {
		if err != nil || err1 != nil {
			return nil, errInternal
		}
		returnCDCES = pb.PartpaidAppointment{Id: "error! Customer not found."}
		return &returnCDCES, nil
	}

	//Add total amount to sale account (checking) using sale Recipt API
	TotalAmountSale := string(in.PrepaidAppointment.Metadata["TotalAmount"])
	SaleReciptLineId := "1111" + in.PrepaidAppointment.AppointyId
	id1, SaleReciptErr := SaleRecipt(in.CompanyId, in.PrepaidAppointment.AppointyId, "Checking", SaleReciptLineId, "adding Total sale amount in a discount enabled sale to income account", TotalAmountSale, s.store)
	if SaleReciptErr != nil {
		returnCDCES = pb.PartpaidAppointment{Id: "error! working with salerecipt API."}
	}
	fmt.Println(id1)

	//Add amount to cash acount (AR) using invoice API
	CashAdded := string(in.PrepaidAppointment.Metadata["TotalAmount"] - in.PrepaidAppointment.Metadata["DiscountAmount"])
	id1, InPayErr := InvoicePayment(in.CompanyId, in.PrepaidAppointment.AppointyId, CashAdded, CustomerId.ExternalId, s.store)
	if InPayErr != nil {
		log.Println("Invoice Payment API not working.")
	}
	fmt.Println(id1)

	//Adding discounted amount to Expense account (which now we use as advertising but will map with Discount expense type acount in the production)
	DiscountAmount := string(in.PrepaidAppointment.Metadata["DiscountAmount"])
	id, BillErr := Bill(in.CompanyId, in.PrepaidAppointment.AppointyId, "Advertising", DiscountAmount, CustomerAsVendorId.ExternalId, s.store)
	if BillErr != nil {
		log.Println("error! working with Bll API.")
		returnCDCES = pb.PartpaidAppointment{Id: "error! working with Bill API."}
	}
	fmt.Println(id)
	return &returnCDCES, nil
}

//Remove Discount Coupon Enabled Sale ##########################################################################################################################################################################################################################################################################################################

func (s *accountyServer) RemoveDiscountCouponEnabledSale(ctx context.Context, in *pb.RemoveDiscountCouponEnabledSaleRequest) (*pb.PartpaidAppointment, error) {

	returnRDCES := pb.PartpaidAppointment{Id: "No app sysnced usnig accounty API."}

	//fetch Ids required for account modification

	SaleReciptAccountyIdDB := "SaleRecipt" + in.PrepaidAppointment.AppointyId
	SaleReciptId, err := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{SaleReciptAccountyIdDB}})
	CustomerId, err1 := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{in.PrepaidAppointment.AppointyId}})
	if SaleReciptId == nil || CustomerId == nil {
		if err != nil || err1 != nil {
			return nil, errInternal
		}
		returnRDCES = pb.PartpaidAppointment{Id: "No account found."}
		return &returnRDCES, nil
	}

	//Reduce from sale account using delete Sale reciot API.
	_, SaleDelErr := DeleteSaleRecipt(in.CompanyId, SaleReciptId.ExternalId)
	if SaleDelErr != nil {
		returnRDCES = pb.PartpaidAppointment{Id: "Delete SaleRecipt API not working."}
		return &returnRDCES, SaleDelErr
	}

	//Add cancellation charges to other income using sale recipt API
	CancellationCharge := string(in.PrepaidAppointment.Metadata["CancellationCharge"])
	SaleReciptLineId := "1111" + in.PrepaidAppointment.AppointyId
	id1, SaleReciptErr := SaleRecipt(in.CompanyId, in.PrepaidAppointment.AppointyId, "Checking", SaleReciptLineId, "Add cancellation charges to income account for cancelling Discount Coupon Enabled Sale ", CancellationCharge, s.store)
	if SaleReciptErr != nil {
		returnRDCES = pb.PartpaidAppointment{Id: "error! working with salerecipt API."}
	}
	fmt.Println(id1)

	//reduce refund from cash using payment API
	RefundedAmt := string(in.PrepaidAppointment.Metadata["TotalAmount"] - in.PrepaidAppointment.Metadata["CancellationCharge"])
	id2, PayErr := Payment(in.CompanyId, in.PrepaidAppointment.AppointyId, "Accounts Recieved (A/R)", RefundedAmt, CustomerId.ExternalId, s.store)
	if PayErr != nil {
		returnRDCES = pb.PartpaidAppointment{Id: "Payment API not working properly."}
	}
	fmt.Println(id2)

	//Remove amount from Expense (addvertisemrnt for now) account add earlier.

	return &returnRDCES, nil
}

//Update Discount Coupon Enabled Sale ##########################################################################################################################################################################################################################################################################################################

func (s *accountyServer) UpdateDiscountCouponEnabledSale(ctx context.Context, in *pb.UpdateDiscountCouponEnabledSaleRequest) (*pb.PartpaidAppointment, error) {

	returnUDCES := pb.PartpaidAppointment{Id: "No app sysnced usnig accounty API."}
	//Fetch required ids for account modification
	SaleReciptAccountyIdDB := "SaleRecipt" + in.PrepaidAppointment.AppointyId
	SaleReciptId, err := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{SaleReciptAccountyIdDB}})
	CustVendId := "Vendor" + in.PrepaidAppointment.AppointyId
	CustomerAsVendorId, err1 := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{1}, pb.AccoutingEmployeeLinkAppointyIdEq{CustVendId}})
	if SaleReciptId == nil || CustomerAsVendorId == nil {
		if err != nil || err1 != nil {
			return nil, errInternal
		}
		returnUDCES = pb.PartpaidAppointment{Id: "No account found."}
		return &returnUDCES, nil
	}

	//Remove Earlier sale amount using delete sale recipt API and then add new sale amount to income account using sale recipt API.
	_, SaleDelErr := DeleteSaleRecipt(in.CompanyId, SaleReciptId.ExternalId)
	if SaleDelErr != nil {
		returnUDCES = pb.PartpaidAppointment{Id: "Delete SaleRecipt API not working."}
		return &returnUDCES, SaleDelErr
	}

	NewSaleAmount := string(in.PrepaidAppointment.Metadata["NewSaleAmount"])
	SaleReciptLineId := "1111" + in.PrepaidAppointment.AppointyId
	id1, SaleReciptErr := SaleRecipt(in.CompanyId, in.PrepaidAppointment.AppointyId, "Checking", SaleReciptLineId, "Add New sale to income account after updating Discount Coupon Enabled Sale ", NewSaleAmount, s.store)
	if SaleReciptErr != nil {
		returnUDCES = pb.PartpaidAppointment{Id: "error! working with salerecipt API."}
	}
	fmt.Println(id1)

	//Add modification charges to income account (checking) using sale recipt API
	ModificationFee := string(in.PrepaidAppointment.Metadata["ModificationFee"])
	id1, SaleReciptErr1 := SaleRecipt(in.CompanyId, in.PrepaidAppointment.AppointyId, "Checking", SaleReciptLineId, "adding modification charge to income account for modifying a payment", ModificationFee, s.store)
	if SaleReciptErr1 != nil {
		returnUDCES = pb.PartpaidAppointment{Id: "error! working with salerecipt API."}
	}
	fmt.Println(id1)

	//If the change is less than the old one then add the difference to expense account (advertising)
	DifferenceAmount := string(in.PrepaidAppointment.Metadata["OldSaleAmount"] - in.PrepaidAppointment.Metadata["NewSaleAmount"])
	id, BillErr := Bill(in.CompanyId, in.PrepaidAppointment.AppointyId, "Advertising", DifferenceAmount, CustomerAsVendorId.ExternalId, s.store)
	if BillErr != nil {
		log.Println("error! working with Bll API.")
		returnUDCES = pb.PartpaidAppointment{Id: "error! working with Bill API."}
	}
	fmt.Println(id)
	return &returnUDCES, nil

}
