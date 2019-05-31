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

var tokenqb string = "Bearer eyJlbmMiOiJBMTI4Q0JDLUhTMjU2IiwiYWxnIjoiZGlyIn0..20_VOud066J33yhfJOS_dQ.00YAssM9uwGdtj-6Kmm8egmOiLI5GHRI3xD_0_9AtcNqg7lHZ99rfq5icNtwAUl0YsPEMD_yueIqFHRzHM3JUBp39-NeDAxZaR8V7hkgSpategrnXRIkBycavgckkJrNYpeXvpT0_oo8Ok58ymxLBC7bPHKQa6wzdj06AyE9ynW_2VS_xeO_aQmfjDhsWTe277vnqrYiwPfbze1_Y-VH1oAk5gSXUqHZvoP2LwsNFXsR17_1OUw9LkEdC8VaEMnkdsDM-DB_syANDG77Zbuhsgnb3XHMjIC7TRrYH86ylPip2b-OIMIkkVKtAYrIa06ruWtXecUX2953_9Vd58AJSxI5mjKDdyoqbYTNzhqxgLDZkmzD5xYbZbgevuTfzGAnrgflbc_l3mB1o6Lgg7BgyNemOMXJdYAXaLC603qWXEON2ww8iqN3h0RRnEfUIvFb6MdbGfZ8wiJSPpcNWnZKmSwe4sHJWoAkqnCCPIjdTU2mRJA8YHRsas-bV7E5_-j2G0BtkFPHFyH-QhpTyrf2g5a5khkq7iQoOoI-zlScAyWjBZvhRMMxBpgQjMUd0-vb1fcVdggB-WxRA5yYSYI1PehGCdFNht4JgW0wk2Z7nkYr7R70MMHJZQtfTYPe_eZY46Fz18zY_4xbQpDlSwygdozESY21PvTOQP3xevjSlknz47Et_5hVwLY9fVdiyb_sbFwq-ogkEbVoYatwCoNzq0plK0fTOUg7ZEV3nx3iVHjHCWLaB1pvu4ptIBohrT1nBU3CEYJuivBYJWEYB2xCqdCeMMSXAWeIslVpFIo9vvo_1OD0GU6lMW_CTX7B1jBf.uwJpIOtsS-pXsmaVXsiwog"
var tokenxero string = "gfdgfd"
var tokenfb string = "fdfdff"

//struct decleration
//struct to be decleard for accepting json request

type accountyServer struct {
	//add store
	store pb.AccoutingEmployeeLinkStore
	//add other clients
}

type Employee1 struct {
	Employee     *Layer1 `json:"Employee"`
	XeroEmployee *Layer1 `json:"Employees"`
}

type Customer1 struct {
	QbCustomer *Layer1 `json:"Customer"`
}

type Timeactivity1 struct {
	BusinessHrs *Layer1 `json:"TimeActivity"`
	TimeEntryFB *Layer1 `json:"time_entry"`
}

type Payment1 struct {
	PaymentQb *Layer1 `json:"Payment"`
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

func ConnectApiEndpoint(urlqb, payloadqbstr, token string) ([]byte, error) {

	//msg := "No app synced with accounty or end point refused to connect"

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

func ConnectToXeroAPI(urlqb, payloadqbstr string) ([]byte, error) {

	//msg := "No app synced with accounty or end point refused to connect"

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

func CreateAccountLedger(AccType, Name, Classification, CompanyId string) string {
	url := "https://SANDBOX-QUICKBOOKS.API.INTUIT.COM/v3/company/" + CompanyId + "/account?minorversion=37"
	payload := "{\n  \"Classification\": \"" + Classification + "\",  \n  \"AccountType\": \"" + AccType + "\",\n  \"Name\": \"" + Name + "\"\n}"
	body, _ := ConnectApiEndpoint(url, payload, tokenqb)
	fmt.Println(string(body))
	return "done"
}

//CreateEmployee ########################################################################################################################################################################################################
func (s *accountyServer) CreateEmployee(ctx context.Context, in *pb.CreateEmployeeRequest) (*pb.Employee, error) {

	sendemp := pb.Employee{Id: "No App synced with accounty service of appointy"}

	if in.AppType == 1 {
		acc := CreateAccountLedger("Accounts Receivable", "MasterAcc", "Asset", in.CompanyId)
		fmt.Println(acc)
		urlqb := "https://SANDBOX-QUICKBOOKS.API.INTUIT.COM/v3/company/" + in.CompanyId + "/employee?minorversion=37"
		CountrySubDivisionCode := in.Employee.Metadata["countrycode"]
		City := in.Employee.Metadata["city"]
		PostalCode := in.Employee.Metadata["pincode"]
		//Creating a staff in quickbook database
		payloadqbstr := "{\n  \"GivenName\": \"" + in.Employee.FirstName + "\", \n  \"SSN\": \"XXX-XX-XXXX\", \n  \"PrimaryAddr\": {\n    \"CountrySubDivisionCode\": \"" + CountrySubDivisionCode + "\", \n    \"City\": \"" + City + "\", \n    \"PostalCode\": \"" + PostalCode + "\", \n    \"Line1\": \"null\"\n  }, \n  \"PrimaryPhone\": {\n    \"FreeFormNumber\": \"" + in.Employee.PhoneNumber + "\"\n  }, \n  \"FamilyName\": \"" + in.Employee.LastName + "\"\n}"

		//calling the API
		//token := "Bearer eyJlbmMiOiJBMTI4Q0JDLUhTMjU2IiwiYWxnIjoiZGlyIn0..IWcaW7YI7XSDvYH7M3Pq-Q.MOJCtljnKVE23wvhEo9nmnKujtHkVDCJi64HNJl-kr1EvP7P4jsGwt_VipR35J9DrURTqsBwcGmWH7rChnly9m9oEjkoFEySGDb5x8E7f5PunI7cyvOGocoVBCaE16GSnfoL0nzvKOTtCF2LBEodGhnDg5AinBJnfp_tFOom9oTJHYI2q5Do2m1KEWsjSzxtg6iMvNa0naDUjvzJPfQIuLXnoiwQbrR-O_TNBowV7EZdpvvqZpE8BDsNwtTNpFgLQduxM0n4ASMAIqP589nHAF1a7h-cX0cyID54IlU92Jyv5cVPJzOMOphlubgxQ0KpfOCkzIRjiqExRp-GLD8n_XSlmtQqWtUF3eSpPnhTcTqoTsQYRzHzV-_VGq8gnUqv9naU59KzIoc_F8ZtTG2_FVy6NsIhdQlkNone-ZFUvmHuAO9ZMP8iqKuV3D_Bo9Y0v0T51OPEsclWhWaUhLH0nGBXVfbV1WKJ55I87iMJDuOKdzRHwwTw_hBHVio6P_2DiZUSOrdWlV7IBppsQ3XXOghTEyBI1ruF1zsooEaTx04Imijzv8-1L3dnPQyTfJzoIy3x-mkI5Z1dwq9mkvOBBPwGPHJ-XZa3ny8_bQaKAKCeZ4swHJFWoNO0u6tChveOZcm1bmmGdGDKfP4kyn-s4baX2Fx3IvCFtOFMGlpxmT_s2QPWj2OrOX1oTvmLxtPNHeOIQm4GKKn9TlUDscjgT01gLGfL-K3mS3UdxCTgGr2B9W3CiPKN6X-mqd0WaR22lM3UZFhbG3DLUTvBLIU7PbP2vGcuWWaRWxPFYRz_tTmmhrte84mxpO7hXr5yEqFu.0xm1_AVBzBGaRBNVW3gZhQ"
		body, _ := ConnectApiEndpoint(urlqb, payloadqbstr, tokenqb)

		//Unmarshelling the json response in Employee1 struct jsut created.(doubt)
		Emp1 := Employee1{}
		jsonErr := json.Unmarshal(body, &Emp1)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		log.Println(Emp1)
		//in.Employee.Id = Employee1.Id

		//Inserting Id generated into chaku database
		ids, _ := s.store.CreateAccoutingEmployeeLinks(ctx, &pb.AccoutingEmployeeLink{
			AppointyId: in.Employee.AppointyId,
			ExternalId: Emp1.Employee.Id,
			AppType:    1,
		})

		//log.Println(len(ids))
		//Returning poulated employee struct with Id of the database back to program for further use
		in.Employee.Id = ids[0]
		sendemp = pb.Employee{Id: in.Employee.Id}
		return &sendemp, nil
	}

	//	Xero employee add
	if in.AppType == 2 {
		urlqb := "https://api.xero.com/api.xro/2.0/Employees"
		payloadqbstr := "{\n  \"Employees\": [\n    {\n      \"FirstName\": \"" + in.Employee.FirstName + "\",\n      \"LastName\": \"" + in.Employee.LastName + "\"\n    }\n  ]\n}"
		//calling the API

		body, _ := ConnectToXeroAPI(urlqb, payloadqbstr)

		//Unmarshelling the json response in Employee1 struct jsut created.(doubt)
		/*Emp2 := Employee1{}
		jsonErr := json.Unmarshal(body, &Emp2)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		*/
		log.Println(string(body))
		//in.Employee.Id = Employee1.Id

		//Inserting Id generated into chaku database
		ids, _ := s.store.CreateAccoutingEmployeeLinks(ctx, &pb.AccoutingEmployeeLink{
			AppointyId: in.Employee.AppointyId,
			ExternalId: "4",
			AppType:    2,
		})

		//log.Println(len(ids))
		//Returning poulated employee struct with Id of the database back to program for further use
		in.Employee.Id = ids[0]
		sendemp = pb.Employee{Id: in.Employee.Id}
		return &sendemp, nil
	}
	return &sendemp, nil

}

//DeleteEmployee #################################################################################################################################################################################################################
func (s *accountyServer) DeleteEmployee(ctx context.Context, in *pb.DeleteEmployeeRequest) (*empty.Empty, error) {

	empl, err := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id", "app_type"}, pb.AccoutingEmployeeLinkAppointyIdEq{in.AppointyEmployeeId})

	//extract employee first name and last name from appointy database and pass it to the below string
	urlqb := "https://SANDBOX-QUICKBOOKS.API.INTUIT.COM/v3/company/" + in.CompanyId + "/employee"
	payloadqbstr := "{\t\n  \"Active\": false, \n   \"Id\": \"" + empl.ExternalId + "\",\n   \"GivenName\": \"from appointy database\",\n    \"FamilyName\": \"from appointy database\"\n  }"

	//calling the API
	//token := "Bearer eyJlbmMiOiJBMTI4Q0JDLUhTMjU2IiwiYWxnIjoiZGlyIn0..IWcaW7YI7XSDvYH7M3Pq-Q.MOJCtljnKVE23wvhEo9nmnKujtHkVDCJi64HNJl-kr1EvP7P4jsGwt_VipR35J9DrURTqsBwcGmWH7rChnly9m9oEjkoFEySGDb5x8E7f5PunI7cyvOGocoVBCaE16GSnfoL0nzvKOTtCF2LBEodGhnDg5AinBJnfp_tFOom9oTJHYI2q5Do2m1KEWsjSzxtg6iMvNa0naDUjvzJPfQIuLXnoiwQbrR-O_TNBowV7EZdpvvqZpE8BDsNwtTNpFgLQduxM0n4ASMAIqP589nHAF1a7h-cX0cyID54IlU92Jyv5cVPJzOMOphlubgxQ0KpfOCkzIRjiqExRp-GLD8n_XSlmtQqWtUF3eSpPnhTcTqoTsQYRzHzV-_VGq8gnUqv9naU59KzIoc_F8ZtTG2_FVy6NsIhdQlkNone-ZFUvmHuAO9ZMP8iqKuV3D_Bo9Y0v0T51OPEsclWhWaUhLH0nGBXVfbV1WKJ55I87iMJDuOKdzRHwwTw_hBHVio6P_2DiZUSOrdWlV7IBppsQ3XXOghTEyBI1ruF1zsooEaTx04Imijzv8-1L3dnPQyTfJzoIy3x-mkI5Z1dwq9mkvOBBPwGPHJ-XZa3ny8_bQaKAKCeZ4swHJFWoNO0u6tChveOZcm1bmmGdGDKfP4kyn-s4baX2Fx3IvCFtOFMGlpxmT_s2QPWj2OrOX1oTvmLxtPNHeOIQm4GKKn9TlUDscjgT01gLGfL-K3mS3UdxCTgGr2B9W3CiPKN6X-mqd0WaR22lM3UZFhbG3DLUTvBLIU7PbP2vGcuWWaRWxPFYRz_tTmmhrte84mxpO7hXr5yEqFu.0xm1_AVBzBGaRBNVW3gZhQ"
	body, Ferr := ConnectApiEndpoint(urlqb, payloadqbstr, tokenqb)

	if Ferr != nil {
		log.Fatal(Ferr)
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
		body, _ := ConnectApiEndpoint(urlqb[p], payloadqbstr[p], token[p])
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

			//Unmarshelling the json response in Employee1 struct jsut created.(doubt)
			bshr := Timeactivity1{}
			jsonErr := json.Unmarshal(body, &bshr)
			if jsonErr != nil {
				log.Fatal(jsonErr)
			}
			returnBHrs = pb.BusinessHour{Name: bshr.BusinessHrs.EmpRef.Name}

		}

		if p == 2 {

			//Unmarshelling the json response in Employee1 struct jsut created.(doubt )
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

//Add prepaid appointment###########################################################################################################################################3
func (s *accountyServer) CreatePrepaidAppointment(ctx context.Context, in *pb.CreatePrepaidAppointmentRequest) (*pb.PrepaidAppointment, error) {

	var urlqb [3]string
	var payloadqbstr [3]string
	returnPPA := pb.PrepaidAppointment{Id: "No App synced with accounty service."}
	var j pb.AccountingIntegrationType
	var p int = int(j)
	var token [3]string
	for p = 0; p < 3; p++ {

		j = j + 1
		ppa, err := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id", "app_type"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{j}, pb.AccoutingEmployeeLinkAppointyIdEq{in.PrepaidAppointment.AppointyId}})
		if ppa == nil {
			fmt.Println(err)
			continue
		}
		if p == 0 {
			urlqb[p] = "https://SANDBOX-QUICKBOOKS.API.INTUIT.COM/v3/company/" + in.CompanyId + "/payment?minorversion=37"
			payloadqbstr[p] = "{\n  \"TotalAmt\": " + in.PrepaidAppointment.TotalAmount + ", \n  \"CustomerRef\": {\n    \"value\": \"" + ppa.ExternalId + "\"\n  }\n}"
			token[p] = tokenqb
		}
		/*if j == 1 {
			urlqb[j] = ""
			payloadqbstr[1] = "{\n    \"NameOf\": \"Employee\",\n    \"EmployeeRef\": {\n        \"value\": \"" + empl.ExternalId + "\",\n        \"name\": \"" + in.BusinessHour.Name + "\"\n    },\n    \"StartTime\": \"2011-07-05T17:00:00-08:00\",\n    \"EndTime\": \"2013-07-05T17:00:00-08:00\"\n}\n"
			token[j] = tokenxero
		}
		if j == 2 {
			urlqb[j] = "https://api.freshbooks.com/timetracking/business/" + in.CompanyId + "/time_entries"
			payloadqbstr[j] = "{\n    \"time_entry\": {\n        \"is_logged\": true,\n        \"duration\": " + in.BusinessHour.TotalTime + ",\n        \"note\": \"" + in.BusinessHour.Description + "\",\n        \"started_at\": \"2016-08-16T20:00:00.000Z\",\n        \"client_id\": " + empl.ExternalId + ",\n        \"project_id\": {{projectId}}\n    }\n}"
			token[j] = tokenfb
		}*/
		body, _ := ConnectApiEndpoint(urlqb[p], payloadqbstr[p], token[p])
		if p == 0 {

			newppa := Payment1{}
			jsonErr := json.Unmarshal(body, &newppa)
			if jsonErr != nil {
				log.Fatal(jsonErr)
			}

			returnPPA = pb.PrepaidAppointment{Id: newppa.PaymentQb.Id}
		}

		/*if j == 1 {
			body, _ := ConnectApiEndpoint(urlqb[0], payloadqbstr[0], token[0])
			fmt.Printf(string(body))
			//Unmarshelling the json response in Employee1 struct jsut created.(doubt)
			BusinessHour1 := Timeactivity1{}
			jsonErr := json.Unmarshal(body, &BusinessHour1)
			if jsonErr != nil {
				log.Fatal(jsonErr)
			}
			returnBHrs = pb.BusinessHour{Name: BusinessHour1.BusinessHrs.EmpRef.Name}
		}

		if j == 2 {
			//calling the API
			body, _ := ConnectApiEndpoint(urlqb[0], payloadqbstr[0], token[0])
			fmt.Printf(string(body))
			//Unmarshelling the json response in Employee1 struct jsut created.(doubt )
			BusinessHourfb := Timeactivity1{}
			jsonErr := json.Unmarshal(body, &BusinessHourfb)
			if jsonErr != nil {
				log.Fatal(jsonErr)
			}

			returnBHrs = pb.BusinessHour{Name: BusinessHourfb.TimeEntryFB.IdFBTimeEntry}
		}*/

		fmt.Printf(string(body))
		return &returnPPA, nil
	}
	return &returnPPA, nil

}

//cancel and refund of the prepaid appointments####################################################################################################################################################################################

func (s *accountyServer) CancelNRefPrepaidAppointment(ctx context.Context, in *pb.CancelNRefPrepaidAppointmentRequest) (*pb.PrepaidAppointment, error) {
	panic(`impliment me`)
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
		//in.Employee.Id = Employee1.Id

		//Inserting Id generated into chaku database
		ids, _ := s.store.CreateAccoutingEmployeeLinks(ctx, &pb.AccoutingEmployeeLink{
			AppointyId: in.Customer.AppointyId,
			ExternalId: cus1.QbCustomer.Id,
			AppType:    1,
		})

		//log.Println(len(ids))
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
		}
		*/
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

	var urlqb [3]string
	var payloadqbstr [3]string

	var j pb.AccountingIntegrationType = 0
	var p int = int(j)
	var token [3]string
	for p = 0; p < 3; p++ {
		j = j + 1
		empl, err := s.store.GetAccoutingEmployeeLink(ctx, []string{"external_id", "app_type"}, pb.AccoutingEmployeeLinkAnd{pb.AccoutingEmployeeLinkAppTypeEq{j}, pb.AccoutingEmployeeLinkAppointyIdEq{in.Customer.AppointyId}})
		if empl == nil {
			fmt.Println(err)
			continue
		}
		if p == 0 {
			urlqb[p] = "https://SANDBOX-QUICKBOOKS.API.INTUIT.COM/v3/company/" + in.CompanyId + "/customer"
			payloadqbstr[p] = "{\n    \"domain\": \"QBO\",\n    \"sparse\": true,\n    \"Id\": \"" + empl.ExternalId + "\",\n   \n    \"Active\": false\n}"
			token[p] = tokenqb
		}
		/*if p == 1 {
			urlqb[p] = ""
			payloadqbstr[p] = "{\n    \"NameOf\": \"Employee\",\n    \"EmployeeRef\": {\n        \"value\": \"" + empl.ExternalId + "\",\n        \"name\": \"" + in.BusinessHour.Name + "\"\n    },\n    \"StartTime\": \"2011-07-05T17:00:00-08:00\",\n    \"EndTime\": \"2013-07-05T17:00:00-08:00\"\n}\n"
			token[p] = tokenxero
		}
		if p == 2 {
			urlqb[p] = "https://api.freshbooks.com/timetracking/business/" + in.CompanyId + "/time_entries"
			payloadqbstr[p] = "{\n    \"time_entry\": {\n        \"is_logged\": true,\n        \"duration\": " + in.BusinessHour.TotalTime + ",\n        \"note\": \"" + in.BusinessHour.Description + "\",\n        \"started_at\": \"2016-08-16T20:00:00.000Z\",\n        \"client_id\": " + empl.ExternalId + ",\n        \"project_id\": {{projectId}}\n    }\n}"
			token[p] = tokenfb
		}*/
		body, _ := ConnectApiEndpoint(urlqb[p], payloadqbstr[p], token[p])

		fmt.Printf(string(body))
	}
	return nil, nil

}

//update customer account ########################################################################################################################################################################################################################################################################################################
func (s *accountyServer) UpdateCustomerAccount(ctx context.Context, in *pb.UpdateCustomerRequest) (*pb.Customer, error) {
	panic(`impliment me`)
}

//create an inventory/addon ######################################################################################################################################################################################################################################################################################################

func (s *accountyServer) CreateInventory(ctx context.Context, in *pb.CreateInventoryRequest) (*pb.Inventory, error) {

	panic(`impliment me`)
	/*
		sendcus := pb.Customer{Id: "No app is synced with accounty API"}

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
			//in.Employee.Id = Employee1.Id

			//Inserting Id generated into chaku database
			ids, _ := s.store.CreateAccoutingEmployeeLinks(ctx, &pb.AccoutingEmployeeLink{
				AppointyId: in.Customer.AppointyId,
				ExternalId: cus1.QbCustomer.Id,
				AppType:    1,
			})

			//log.Println(len(ids))
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
			}

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
		return &sendcus, nil*/
}
