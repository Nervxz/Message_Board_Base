package dao

import (
	"strconv"
	"strings"
	"testing"
)

func TestCommentDAO_genInsertSQL(t *testing.T) {
	tests := []struct {
		name        string
		numComments int
		want        string
	}{
		{
			name:        "single comment",
			numComments: 1,
			want: strings.TrimSpace(`
insert into msg_board.comments(by, topic_id, body, created_at) values ($1,$2,$3,$4) returning id;
`),
		},
		{
			name:        "multiple comments",
			numComments: 3,
			want: strings.TrimSpace(`
insert into msg_board.comments(by, topic_id, body, created_at) values ($1,$2,$3,$4),($5,$6,$7,$8),($9,$10,$11,$12) returning id;
`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := CommentDAO{}
			if got := d.genInsertSQL(tt.numComments); got != tt.want {
				t.Errorf("genInsertSQL()=\n%v\n want\n%v", got, tt.want)
			}
		})
	}
}

func BenchmarkCommentDAO_genInsertSQL(b *testing.B) {
	d := CommentDAO{}
	for n := 1; n <= 5; n++ {
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			b.ReportAllocs()

			for j := 0; j < b.N; j++ {
				d.genInsertSQL(n)
			}
		})
	}
}
