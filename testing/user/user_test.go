package user

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	db2 "iw2/testing/user/db"
	"iw2/testing/user/db/mock"
	"iw2/testing/user/model"
	"reflect"
	"testing"
)

func TestGetUserById(t *testing.T) {
	type args struct {
		db db2.Db
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    model.User
		wantErr error
	}{
		{
			name: "Select simple",
			args: args{
				db: nil,
				id: 18,
			},
			want:    model.User{Id: 18},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mock := mock.NewMockDb(gomock.NewController(t))
			mock.EXPECT().QueryUser("SELECT * FROM users WHERE id = ?", 18)

			got, err := GetUserById(tt.args.db, tt.args.id)

			if tt.wantErr != nil {
				assert.Equal(t, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserById() got = %v, want %v", got, tt.want)
			}
		})
	}
}
