package unit

import (
	"news/internal/data"
	"news/models"
	"reflect"
	"testing"
)

func TestDB_News(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		db   *data.DB
		args args
		want []models.Post
	}{
		{
			name: "last element from three",
			db: func() *data.DB {
				db := data.New()
				new1 := models.Post{
					Title: "сорокопят",
				}
				new2 := models.Post{
					Title: "азимут",
				}
				new3 := models.Post{
					Title: "эта новость",
				}
				db.AddNews(new1)
				db.AddNews(new2)
				db.AddNews(new3)

				return db
			}(),
			args: args{1},
			want: []models.Post{
				{
					Title: "эта новость",
				},
			},
		},
		{
			name: "Empty DB",
			db:   data.New(),
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := tt.db.News(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.News() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_AddNews(t *testing.T) {
	db := data.New()
	new := models.Post{
		Title: "Живьем не брать",
	}
	db.AddNews(new)

	news, err := db.News(1)
	if err != nil {
		t.Errorf("Ошибка получения %v", err)
	}
	if len(news) != 1 {
		t.Errorf("Ошибка длины БД, должно быть 1,а по факту %d", len(news))
	}
}
