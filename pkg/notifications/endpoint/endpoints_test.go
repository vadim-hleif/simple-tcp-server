package endpoint

import (
	"reflect"
	"testing"

	"github.com/go-kit/kit/endpoint"

	"simple-tcp-server/pkg/notifications"
)

func TestMakeEndpoints(t *testing.T) {
	type args struct {
		storage notifications.UsersStorage
	}
	tests := []struct {
		name string
		args args
		want Endpoints
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeEndpoints(tt.args.storage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeEndpoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeUserLoginEndpoint(t *testing.T) {
	type args struct {
		storage notifications.UsersStorage
	}
	tests := []struct {
		name string
		args args
		want endpoint.Endpoint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeUserLoginEndpoint(tt.args.storage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeUserLoginEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeUserLogoutEndpoint(t *testing.T) {
	type args struct {
		storage notifications.UsersStorage
	}
	tests := []struct {
		name string
		args args
		want endpoint.Endpoint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeUserLogoutEndpoint(tt.args.storage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeUserLogoutEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
