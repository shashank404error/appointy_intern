package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	alphaa "go.appointy.com/accounty"
	"go.appointy.com/accounty/pb"
	"go.appointy.com/chaku/driver"
)

var ctx context.Context = context.Background()

//var tknqb string="Bearer eyJlbmMiOiJBMTI4Q0JDLUhTMjU2IiwiYWxnIjoiZGlyIn0..Ba_hmmFaoSs4NEE_MTuUSg.o5fi6K6u2EdyGaCNWx1gsYm3TWYkxQg8ToDfbiX7WhcLZ2Cebq3NSCC7wTteiUM1EtabZBtp1RloCQrgstI8svRvgbre3S9vT9QSc1Z_rc_b_A0ff78tF3mv3SDiANm0LsHbP5BLRFIw8DE37qaZoO9stPuhaWisjGoej8sbIxST-TXyB8imHIHZezoMQ82uW7sG1OnVJXOJsADpD8RWS0qVWxvO4SRkckqSifZpp6L6UjwICsAKjH16yLlYHTKexuxhd20tyKJQ_dVhXmaqCrM7ws9Q1p1ll0M1hWRn358C87E4gDgtju3dP7W4UxdqtUnF-NwHty1MfkYck1fdkFsUY6F1Sid7Ti-FeIuzcbjwgAqSSVLIfY_JWX9XxQ9TkaGnbGG6Proo7ugGInXxmYSyXFhM3K0B6cgT9LWkmFUPL1VmzqLyNjjHBEnF_URqvFvdI9wWy8s2a1uwYGM36hzbUNh-U61IyBBYARjv7x54ztmVuUBCgrN-DJrDooyRxdpfzNzYNvxa7XPGD119Ax4xD5ORgg16gDH4qMYjt2TdGePbS91jri2CCJWKR6hvlzT325E5YnYocYJAZFaRASqH1BKpespjI_9f0_mBm5PCPSw48T2uZwgDJ-DqpSa-d_jfESC4CFkRJzP6ky2S_WPD225E5wS27XjAOAiQO-6TG1B3qN6V6newQ1i4hSqn9hWwcK7qsf28D4TGkFRulL9-nQmdWJY20v8bQMDWS6q9-IIK4yK_VEOIG5ekNBnzz1tyI9nEWIiJ0CqWb_RMMz6P8jK2AaqujyRwY-d7W8JAu2dLutejxRFHgQg7-r1q.GlCzMv6rC23bgKc_v4dxiQ"
//var tknxero string=",OAuth oauth_consumer_key=\"54BT1LQ29XWN7N7PMCLOZ5KBE0WTFY\",oauth_token=\"PLV5ALOCORQAD5J1BJXIA5VM4FBMXM\",oauth_signature_method=\"HMAC-SHA1\",oauth_timestamp=\"1558615444\",oauth_nonce=\"faScuBF6ZgT\",oauth_version=\"1.0\",oauth_signature=\"GvFgqc9FoZXMlv%2B1NmRCpbVZ%2Bgk%3D\""

func getServer() pb.QuickbookServerServer {

	store := createStore()

	if err := store.CreateAccoutingEmployeeLinkPGStore(context.Background()); err != nil {
		return nil
	}
	return alphaa.NewquickbookServerServer(store)
}

func createStore() pb.AccoutingEmployeeLinkStore {
	config := getConfig()

	db, err := sql.Open("postgres", config)
	if err != nil {
		panic(fmt.Errorf("connection not open | %v", err.Error()))
	}
	if err = db.Ping(); err != nil {

		panic("ping  " + err.Error())
	}
	return pb.NewPostgresAccoutingEmployeeLinkStore(db, driver.GetUserName)
}

func getConfig() string {
	config := "host=localhost port=5432 user=postgres password=shashank dbname=chaku sslmode=disable"

	// create pgconfig.json and put your credentials in it
	/*pg, err := os.Open("pgconfig.json")
	if err == nil {
		pgc, err := ioutil.ReadAll(pg)
		if err != nil {
			log.Println(err)
		}

		mp := make(map[string]string, 0)
		err = json.Unmarshal(pgc, &mp)
		if err != nil {
			log.Fatalln(err)
		}
		config = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			mp["host"], mp["port"], mp["user"], mp["password"], mp["dbname"])
	}*/
	return config
}

func AddStaff(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering AddStaff")
	emp := &pb.CreateEmployeeRequest{CompanyId: "123146508006334", AppType: 1, Employee: &pb.Employee{AppointyId: "1111", FirstName: "testing creating acc", LastName: "in quickbook", PhoneNumber: "9079528686", Metadata: map[string]string{"countrycode": "91", "city": "patna", "pincode": "1234567"}}}

	aa := getServer()
	val, _ := aa.CreateEmployee(ctx, emp)
	fmt.Println(val)
	head := "<h1><center><font color=blue><u><ul><li> " + val.Id + " </li></ul></u></font></center></h1>"
	fmt.Println("here")
	fmt.Fprintf(w, head)
	log.Println("Exiting AddStaff")
}

func RemoveStaff(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering RemoveStaff ")
	emp := &pb.DeleteEmployeeRequest{CompanyId: "123146508006334", AppType: 1, AppointyEmployeeId: "1005"}
	aa := getServer()
	_, err := aa.DeleteEmployee(ctx, emp)
	if err != nil {
		log.Fatal(err)
	}
	head := "<h1><center><font color=blue><u><ul><li>" + emp.AppointyEmployeeId + " Deleted</li></ul></u></font></center></h1>"
	fmt.Fprintf(w, head)
	log.Println("Exiting RemoveStaff ")
}

func AddStaffHours(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering addStaffhours ")
	timeentry := &pb.CreateBusinessHourRequest{CompanyId: "123146508006334", AppType: 1, BusinessHour: &pb.BusinessHour{AppointyEmployeeId: "pb1013", Name: "finaltest13 sholay6"}}
	aa := getServer()
	val, _ := aa.CreateBusinessHour(ctx, timeentry)
	fmt.Println(val)
	head := "<h1><center><font color=blue><u><ul><li>" + val.Name + " Added Business Hours</li></ul></u></font></center></h1>"
	fmt.Fprintf(w, head)
	log.Println("Exiting addStaffhours ")
}

func AddCustomer(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering addcustomer ")
	addCustomer := &pb.CreateCustomerRequest{CompanyId: "123146508006334", AppType: 1, Customer: &pb.Customer{AppointyId: "pb1013", FirstName: "finaltest13", LastName: "sholay6", PhoneNumber: "1234561222", Metadata: map[string]string{"countrycode": "91", "city": "patna", "baseaddr": "1234567"}}}
	aa := getServer()
	val, _ := aa.CreateCustomerAccount(ctx, addCustomer)
	fmt.Println(val)
	head := "<h1><center><font color=blue><u><ul><li>" + val.Id + " Added as customer in quickbook</li></ul></u></font></center></h1>"
	fmt.Fprintf(w, head)
	log.Println("Exiting addCustomer")
}

func AddPrepaidApp(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering addprepaidappointment")
	addCustomer := &pb.CreatePrepaidAppointmentRequest{CompanyId: "123146508006334", PrepaidAppointment: &pb.PrepaidAppointment{AppointyId: "pb1013", TotalAmount: "25.0"}}
	aa := getServer()
	val, _ := aa.CreatePrepaidAppointment(ctx, addCustomer)
	fmt.Println(val)
	head := "<h1><center><font color=blue><u><ul><li>" + val.Id + " Added as customer in quickbook</li></ul></u></font></center></h1>"
	fmt.Fprintf(w, head)
	log.Println("Exiting addprepaidappointment")
}

/*
func RemoveStaffHours(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering RemoveStaffhours ")
	timeentry := &pb.TimeEntryRequest{Tokenqb: tknqb, Tokenxero: tknxero, Tokenfb: "ddfdf", AppNo: "", Timeentry: &pb.TimeEntry{Id: "59", Description: "0"}}
	aa := alphaa.Newaccounting_apiServer()
	val, _ := aa.RemoveStaffHours(ctx, timeentry)
	fmt.Println(val)
	head := "<h1><center><font color=blue><u><ul><li>" + val.Response + "</li></ul></u></font></center></h1>"
	fmt.Fprintf(w, head)
	log.Println("Exiting RemoveStaffhours ")
}*/

func main() {

	fmt.Println("Accounty Listning on port:6060")
	//register handler routes
	http.HandleFunc("/addstaff", AddStaff)
	http.HandleFunc("/removestaff", RemoveStaff)
	http.HandleFunc("/addstaffhours", AddStaffHours)
	http.HandleFunc("/addcustomer", AddCustomer)
	http.HandleFunc("/addppa", AddPrepaidApp)
	//http.HandleFunc("/removestaffhours", RemoveStaffHours)

	http.ListenAndServe(":6060", nil)
	//log and start server
	//log.Println("running server on ", config.OAuthConfig.Port)
	//log.Fatal(http.ListenAndServe(config.OAuthConfig.Port, nil))

}
