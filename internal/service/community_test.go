package service

import (
	"sublinks/sublinks-federation/internal/db"
	"sublinks/sublinks-federation/internal/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func compareGroup(t *testing.T, got, want model.Group) {
	if got.Id != want.Id {
		t.Errorf("Id = %v, want %v", got.Id, want.Id)
	}
	if got.Username != want.Username {
		t.Errorf("Username = %v, want %v", got.Username, want.Username)
	}
	if got.Bio != want.Bio {
		t.Errorf("Bio = %v, want %v", got.Bio, want.Bio)
	}
	if got.PublicKey != want.PublicKey {
		t.Errorf("PublicKey = %v, want %v", got.PublicKey, want.PublicKey)
	}
	if got.PrivateKey != want.PrivateKey {
		t.Errorf("PrivateKey = %v, want %v", got.PrivateKey, want.PrivateKey)
	}
}

func TestCommunityService_GetById(t *testing.T) {
	type fields struct {
		db db.Database
	}
	type args struct {
		id string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *model.Group
	}{
		{
			name: "TestCommunityService_GetById",
			fields: fields{
				db: NewFakeDB(),
			},
			args: args{
				id: "https://sublinks.com/c/somecommunity",
			},
			want: &model.Group{
				Id:         "https://sublinks.com/c/somecommunity",
				Username:   "somecommunity",
				Bio:        "some bio",
				PublicKey:  "some key",
				PrivateKey: "some private",
				Posts:      []model.Post{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := CommunityService{
				db: tt.fields.db,
			}
			if err := tt.fields.db.Connect(); err != nil {
				t.Errorf("Error connecting to database: %v", err)
			}
			groupRows := sqlmock.NewRows([]string{"id", "username", "bio", "public_key", "private_key"}).AddRow(tt.want.Id, tt.want.Username, tt.want.Bio, tt.want.PublicKey, tt.want.PrivateKey)
			tt.fields.db.(*fakeDB).Mock.ExpectQuery("^SELECT (.+) FROM \"groups\" WHERE \"groups\".\"id\" = (.+)$").WithArgs(tt.args.id).WillReturnRows(groupRows)
			postRows := sqlmock.NewRows([]string{})
			tt.fields.db.(*fakeDB).Mock.ExpectQuery("^SELECT (.+) FROM \"posts\" WHERE \"posts\".\"author_id\" = ? AND \"posts\".\"deleted_at\" IS NULL$").WithArgs(tt.args.id).WillReturnRows(postRows)
			got := a.GetById(tt.args.id)
			compareGroup(t, *got, *tt.want)
		})
	}
}

func TestCommunityService_GetByUsername(t *testing.T) {
	type fields struct {
		db db.Database
	}
	type args struct {
		username string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *model.Group
	}{
		{
			name: "TestCommunityService_GetByUsername",
			fields: fields{
				db: NewFakeDB(),
			},
			args: args{
				username: "somecommunity",
			},
			want: &model.Group{
				Id:         "https://sublinks.com/c/somecommunity",
				Username:   "somecommunity",
				Bio:        "some bio",
				PublicKey:  "some key",
				PrivateKey: "some private",
				Posts:      []model.Post{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := CommunityService{
				db: tt.fields.db,
			}
			if err := tt.fields.db.Connect(); err != nil {
				t.Errorf("Error connecting to database: %v", err)
			}
			groupRows := sqlmock.NewRows([]string{"id", "username", "bio", "public_key", "private_key"}).AddRow(tt.want.Id, tt.want.Username, tt.want.Bio, tt.want.PublicKey, tt.want.PrivateKey)
			tt.fields.db.(*fakeDB).Mock.ExpectQuery("^SELECT (.+) FROM \"groups\" WHERE \"groups\".\"username\" = (.+)$").WithArgs(tt.args.username).WillReturnRows(groupRows)
			postRows := sqlmock.NewRows([]string{})
			tt.fields.db.(*fakeDB).Mock.ExpectQuery("^SELECT (.+) FROM \"posts\" WHERE \"posts\".\"author_id\" = ? AND \"posts\".\"deleted_at\" IS NULL$").WithArgs(tt.want.Id).WillReturnRows(postRows)
			got := a.GetByUsername(tt.args.username)
			compareGroup(t, *got, *tt.want)
		})
	}
}

func TestCommunityService_Save(t *testing.T) {
	type fields struct {
		db db.Database
	}
	type args struct {
		person *model.Group
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "TestCommunityService_Save",
			fields: fields{
				db: NewFakeDB(),
			},
			args: args{
				person: &model.Group{
					Id:         "https://sublinks.com/c/somecommunity",
					Username:   "somecommunity",
					Bio:        "some bio",
					PublicKey:  "some key",
					PrivateKey: "some private",
					Posts:      []model.Post{},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := CommunityService{
				db: tt.fields.db,
			}
			if err := tt.fields.db.Connect(); err != nil {
				t.Errorf("Error connecting to database: %v", err)
			}
			tt.fields.db.(*fakeDB).Mock.ExpectExec("^UPDATE \"groups\" SET \"groups\".\"usrname\" = (.+) WHERE (.+)$").WillReturnResult(sqlmock.NewResult(1, 1))
			tt.fields.db.(*fakeDB).Mock.ExpectExec("^INSERT INTO \"groups\" (.+)$").WillReturnResult(sqlmock.NewResult(1, 1))
			if got := a.Save(tt.args.person); got != tt.want {
				t.Errorf("CommunityService.Save() = %v, want %v", got, tt.want)
			}
		})
	}
}
