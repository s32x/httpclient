package cachers /* import "s32x.com/httpclient/cache/cachers" */

import (
	"reflect"
	"testing"
	"time"
)

func TestBadger_Set(t *testing.T) {
	b, err := NewBadger(t.Name(), time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		key        string
		val        []byte
		expiration []time.Duration
	}
	tests := []struct {
		name    string
		badger  *Badger
		args    args
		wantErr string
	}{
		{
			name:   "success",
			badger: b,
			args: args{
				key: "a",
				val: []byte("b"),
			},
		},
		{
			name:   "empty key",
			badger: b,
			args: args{
				key: "",
				val: nil,
			},
			wantErr: "badger db update: Key cannot be empty",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.badger.Set(tt.args.key, tt.args.val, tt.args.expiration...)
			if (err != nil) && err.Error() != tt.wantErr {
				t.Errorf("Badger.Set() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestBadger_Get(t *testing.T) {
	b, err := NewBadger(t.Name(), time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	if err := b.Set("valid-test", []byte("b")); err != nil {
		t.Fatal(err)
	}
	if err := b.Set("expired-test", []byte("y"), 1); err != nil {
		t.Fatal(err)
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		badger  *Badger
		args    args
		want    []byte
		wantErr string
	}{
		{
			name:   "valid test",
			badger: b,
			args: args{
				key: "valid-test",
			},
			want: []byte("b"),
		},
		{
			name:   "empty key",
			badger: b,
			args: args{
				key: "",
			},
			wantErr: "badger db view: Key cannot be empty",
		},
		{
			name:   "not found",
			badger: b,
			args: args{
				key: "expired-test",
			},
			wantErr: "badger db view: key not found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				got, err := tt.badger.Get(tt.args.key)
				if (err != nil) && err.Error() != tt.wantErr {
					t.Errorf("Badger.Get() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Badger.Get() = %v, want %v", got, tt.want)
				}
			})
		})
	}
}
