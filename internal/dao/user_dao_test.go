package dao

import (
	"strconv"
	"strings"
	"testing"
)

func TestUserDAO_genInsertSQL(t *testing.T) {
	tests := []struct {
		name     string
		numUsers int
		want     string
	}{
		{
			name:     "single user",
			numUsers: 1,
			want: strings.TrimSpace(`
insert into msg_board.users(username, password, created_at) values ($1,$2,$3) returning id;
`),
		},

		{
			name:     "single user",
			numUsers: 3,
			want: strings.TrimSpace(`
insert into msg_board.users(username, password, created_at) values ($1,$2,$3),($4,$5,$6),($7,$8,$9) returning id;
`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := UserDAO{}
			if got := d.genInsertSQL(tt.numUsers); got != tt.want {
				t.Errorf("genInsertSQL()=\n%v\n want\n%v", got, tt.want)
			}
		})
	}
}

func BenchmarkUserDAO_genInsertSQL(b *testing.B) {
	d := UserDAO{}
	for n := 1; n < 5; n++ {
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			b.ReportAllocs()

			for j := 0; j < b.N; j++ {
				d.genInsertSQL(n)
			}
		})
	}
}
