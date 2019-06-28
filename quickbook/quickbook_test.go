package accounty_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
//	"go.appointy.com/calendar/gcalendar"
	"go.appointy.com/calendar/pb"
	"go.appointy.com/chaku/driver"
	"google.golang.org/grpc"

	// "google.golang.org/api/calendar/v3"
	// "google.golang.org/api/option"
	integrationpb "go.appointy.com/integration/pb"
	"google.golang.org/genproto/protobuf/field_mask"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"testing"
	"time"
)

var (
	testCalendarId = "sh4itv0fcil3bogfs844do15vk@group.calendar.google.com"
	user1          = "user1"
)

func TestGcalendarServer_ListCalendar(t *testing.T) {
	type args struct {
		ctx context.Context
		req *pb.ListCalendarRequest
	}

	type wants struct {
		resp    *pb.ListCalendarResponse
		wantErr bool
	}

	tests := []struct {
		name  string
		a     *args
		w     *wants
		setup func(a *args, w *wants)
	}{
		{
			name: "List",
			a: &args{
				ctx: context.Background(),
				req: &pb.ListCalendarRequest{UserId: user1},
			},
			w: &wants{
				resp:    nil,
				wantErr: false,
			},
			setup: func(a *args, w *wants) {
				list, err := ListCurrentCalendars()
				if err != nil {
					t.Error("ListCurrentCalendars err", err.Error())
					return
				}
				w.resp = &pb.ListCalendarResponse{Calendars: list}
			},
		},
	}

	srv, srvErr := getServer()
	if srvErr != nil {
		t.Errorf("getting server failed error = %v,", srvErr)
		return
	}
	for _, tt := range tests {
		tt.setup(tt.a, tt.w)
		got, err := srv.ListCalendar(tt.a.ctx, tt.a.req)
		if (err != nil) != tt.w.wantErr {
			t.Errorf("srv.GetEvent() error = %v, wantErr %v", err, tt.w.wantErr)
			return
		}

		if !tt.w.wantErr {
			if !reflect.DeepEqual(got, tt.w.resp) {
				t.Errorf("srv.GetEvent() = %v, want %v", got, tt.w.resp)
			}
		}
	}
}

func ListCurrentCalendars() ([]*pb.Calendar, error) {
	srv, srvErr := getServer()
	if srvErr != nil {
		return nil, srvErr
	}

	got, err := srv.ListCalendar(context.Background(), &pb.ListCalendarRequest{UserId: user1})
	if err != nil {
		return nil, err
	}
	return got.Calendars, nil
}


func getServer() (pb.CalendarsServer, error) {
	//calendarService, err := calendar.NewService(ctx, option.WithTokenSource(googleOauthConfig.TokenSource(ctx, token)))
	ctx := context.Background()
	subscriptionStore := createStores()
	if err := subscriptionStore.CreateAccoutingEmployeeLinkPGStore(ctx); err != nil {
		return nil, err
	}

	// get Interation Client
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	integrationClient := integrationpb.NewIntegrationsClient(conn)

	// get Appointment Client
	return gcalendar.NewCalendarsServer(subscriptionStore, integrationClient, nil), nil
}

func createStores() pb.AccoutingEmployeeLinkStore {
	config := getConfig()

	db, err := sql.Open("postgres", config)
	if err != nil {
		panic(fmt.Errorf("connection not open | %v", err.Error()))
	}
	if err = db.Ping(); err != nil {
		panic("ping  " + err.Error())
	}
	return pb.NewPostgresAccoutingEmployeeLinkStore(db, driver.GetUserId)
}

func getConfig() string {
	config := "host=localhost port=5432 user=postgres password=shashank dbname=chaku sslmode=disable"

	// create pgconfig.json and put your credentials in it
	pg, err := os.Open("pgconfig.json")
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
	}
	return config
}
