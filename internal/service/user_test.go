package service

import (
	"sublinks/sublinks-federation/internal/db"
	"sublinks/sublinks-federation/internal/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func comparePerson(t *testing.T, got, want model.Person) {
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

func TestUserService_GetById(t *testing.T) {
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
		want   *model.Person
	}{
		{
			name: "TestUserService_GetById",
			fields: fields{
				db: NewFakeDB(),
			},
			args: args{
				id: "https://sublinks.com/u/someuser",
			},
			want: &model.Person{
				Id:         "https://sublinks.com/u/someuser",
				Username:   "someuser",
				Bio:        "some bio",
				PublicKey:  "some key",
				PrivateKey: "some private",
				Posts:      []model.Post{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := UserService{
				db: tt.fields.db,
			}
			if err := tt.fields.db.Connect(); err != nil {
				t.Errorf("Error connecting to database: %v", err)
			}
			peopleRows := sqlmock.NewRows([]string{"id", "username", "bio", "public_key", "private_key"}).AddRow(tt.want.Id, tt.want.Username, tt.want.Bio, tt.want.PublicKey, tt.want.PrivateKey)
			tt.fields.db.(*fakeDB).Mock.ExpectQuery("^SELECT (.+) FROM \"people\" WHERE \"people\".\"id\" = (.+)$").WithArgs(tt.args.id).WillReturnRows(peopleRows)
			postRows := sqlmock.NewRows([]string{})
			tt.fields.db.(*fakeDB).Mock.ExpectQuery("^SELECT (.+) FROM \"posts\" WHERE \"posts\".\"author_id\" = ? AND \"posts\".\"deleted_at\" IS NULL$").WithArgs(tt.args.id).WillReturnRows(postRows)
			got := a.GetById(tt.args.id)
			comparePerson(t, *got, *tt.want)
		})
	}
}

func TestUserService_GetByUsername(t *testing.T) {
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
		want   *model.Person
	}{
		{
			name: "TestUserService_GetByUsername",
			fields: fields{
				db: NewFakeDB(),
			},
			args: args{
				username: "someuser",
			},
			want: &model.Person{
				Id:         "https://sublinks.com/u/someuser",
				Username:   "someuser",
				Bio:        "some bio",
				PublicKey:  "some key",
				PrivateKey: "some private",
				Posts:      []model.Post{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := UserService{
				db: tt.fields.db,
			}
			if err := tt.fields.db.Connect(); err != nil {
				t.Errorf("Error connecting to database: %v", err)
			}
			peopleRows := sqlmock.NewRows([]string{"id", "username", "bio", "public_key", "private_key"}).AddRow(tt.want.Id, tt.want.Username, tt.want.Bio, tt.want.PublicKey, tt.want.PrivateKey)
			tt.fields.db.(*fakeDB).Mock.ExpectQuery("^SELECT (.+) FROM \"people\" WHERE \"people\".\"username\" = (.+)$").WithArgs(tt.args.username).WillReturnRows(peopleRows)
			postRows := sqlmock.NewRows([]string{})
			tt.fields.db.(*fakeDB).Mock.ExpectQuery("^SELECT (.+) FROM \"posts\" WHERE \"posts\".\"author_id\" = ? AND \"posts\".\"deleted_at\" IS NULL$").WithArgs(tt.want.Id).WillReturnRows(postRows)
			got := a.GetByUsername(tt.args.username)
			comparePerson(t, *got, *tt.want)
		})
	}
}

func TestUserService_Save(t *testing.T) {
	type fields struct {
		db db.Database
	}
	type args struct {
		person *model.Person
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "TestUserService_Save",
			fields: fields{
				db: NewFakeDB(),
			},
			args: args{
				person: &model.Person{
					Id:         "https://sublinks.com/u/someuser",
					Username:   "someuser",
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
			a := UserService{
				db: tt.fields.db,
			}
			if err := tt.fields.db.Connect(); err != nil {
				t.Errorf("Error connecting to database: %v", err)
			}
			tt.fields.db.(*fakeDB).Mock.ExpectExec("^UPDATE \"people\" SET \"people\".\"usrname\" = (.+) WHERE (.+)$").WillReturnResult(sqlmock.NewResult(1, 1))
			tt.fields.db.(*fakeDB).Mock.ExpectExec("^INSERT INTO \"people\" (.+)$").WillReturnResult(sqlmock.NewResult(1, 1))
			if got := a.Save(tt.args.person); got != tt.want {
				t.Errorf("UserService.Save() = %v, want %v", got, tt.want)
			}
		})
	}
}
