package service

import (
	"sublinks/sublinks-federation/internal/db"
	"sublinks/sublinks-federation/internal/model"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func compareComment(t *testing.T, got, want model.Comment) {
	if got.Id != want.Id {
		t.Errorf("Id = %v, want %v", got.Id, want.Id)
	}
	if got.Nsfw != want.Nsfw {
		t.Errorf("Nsfw = %v, want %v", got.Nsfw, want.Nsfw)
	}
	if got.Content != want.Content {
		t.Errorf("Content = %v, want %v", got.Content, want.Content)
	}
	if got.UrlStub != want.UrlStub {
		t.Errorf("UrlStub = %v, want %v", got.UrlStub, want.UrlStub)
	}
	if got.Published != want.Published {
		t.Errorf("Published = %v, want %v", got.Published, want.Published)
	}
}

func TestCommentService_GetById(t *testing.T) {
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
		want   *model.Comment
	}{
		{
			name: "TestCommentService_GetById",
			fields: fields{
				db: NewFakeDB(),
			},
			args: args{
				id: "https://sublinks.com/comment/somecomment",
			},
			want: &model.Comment{
				Id:        "https://sublinks.com/comment/somecomment",
				Content:   "some content",
				Nsfw:      false,
				UrlStub:   "somecomment",
				Published: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := CommentService{
				db: tt.fields.db,
			}
			if err := tt.fields.db.Connect(); err != nil {
				t.Errorf("Error connecting to database: %v", err)
			}
			commentRows := sqlmock.NewRows([]string{"id", "content", "nsfw", "url_stub", "published"}).AddRow(tt.want.Id, tt.want.Content, tt.want.Nsfw, tt.want.UrlStub, tt.want.Published)
			tt.fields.db.(*fakeDB).Mock.ExpectQuery("^SELECT (.+) FROM \"comments\" WHERE \"comments\".\"id\" = (.+)$").WithArgs(tt.args.id).WillReturnRows(commentRows)
			got := a.GetById(tt.args.id)
			compareComment(t, *got, *tt.want)
		})
	}
}

func TestCommentService_Save(t *testing.T) {
	type fields struct {
		db db.Database
	}
	type args struct {
		comment *model.Comment
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "TestCommentService_Save",
			fields: fields{
				db: NewFakeDB(),
			},
			args: args{
				comment: &model.Comment{
					Id:        "https://sublinks.com/comment/somecomment",
					Content:   "some content",
					Nsfw:      false,
					UrlStub:   "somecomment",
					Published: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := CommentService{
				db: tt.fields.db,
			}
			if err := tt.fields.db.Connect(); err != nil {
				t.Errorf("Error connecting to database: %v", err)
			}
			tt.fields.db.(*fakeDB).Mock.ExpectExec("^UPDATE \"comments\" SET \"comments\".\"usrname\" = (.+) WHERE (.+)$").WillReturnResult(sqlmock.NewResult(1, 1))
			tt.fields.db.(*fakeDB).Mock.ExpectExec("^INSERT INTO \"comments\" (.+)$").WillReturnResult(sqlmock.NewResult(1, 1))
			if got := a.Save(tt.args.comment); got != tt.want {
				t.Errorf("CommentService.Save() = %v, want %v", got, tt.want)
			}
		})
	}
}
