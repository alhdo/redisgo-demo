package core

import (
	"reflect"
	"testing"
)

func TestNewPubSubService(t *testing.T) {
	type args struct {
		pool       *redis.Pool
		pubChannel string
		subChannel string
	}
	tests := []struct {
		name string
		args args
		want *PubSubService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPubSubService(tt.args.pool, tt.args.pubChannel, tt.args.subChannel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPubSubService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPubSubService_PublishMessage(t *testing.T) {
	type fields struct {
		pool       *redis.Pool
		subChannel string
		pubChannel string
	}
	type args struct {
		message string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &PubSubService{
				pool:       tt.fields.pool,
				subChannel: tt.fields.subChannel,
				pubChannel: tt.fields.pubChannel,
			}
			if err := ps.PublishMessage(tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("PublishMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPubSubService_Subscribe(t *testing.T) {
	type fields struct {
		pool       *redis.Pool
		subChannel string
		pubChannel string
	}
	type args struct {
		messages chan string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &PubSubService{
				pool:       tt.fields.pool,
				subChannel: tt.fields.subChannel,
				pubChannel: tt.fields.pubChannel,
			}
			ps.Subscribe(tt.args.messages)
		})
	}
}
