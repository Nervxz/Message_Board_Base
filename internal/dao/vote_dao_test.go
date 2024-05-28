package dao

import (
	"strconv"
	"strings"
	"testing"
)

func TestVoteDAO_genInsertSQL(t *testing.T) {
	tests := []struct {
		name     string
		numVotes int
		want     string
	}{
		{
			name:     "single vote",
			numVotes: 1,
			want: strings.TrimSpace(`
insert into msg_board.votes(by, topic_id, created_at) values ($1,$2,$3) returning id;
`),
		},
		{
			name:     "multiple votes",
			numVotes: 3,
			want: strings.TrimSpace(`
insert into msg_board.votes(by, topic_id, created_at) values ($1,$2,$3),($4,$5,$6),($7,$8,$9) returning id;
`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := VoteDAO{}
			if got := d.genInsertSQL(tt.numVotes); got != tt.want {
				t.Errorf("genInsertSQL()=\n%v\n want\n%v", got, tt.want)
			}
		})
	}
}

func BenchmarkVoteDAO_genInsertSQL(b *testing.B) {
	d := VoteDAO{}
	for n := 1; n < 5; n++ {
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			b.ReportAllocs()

			for j := 0; j < b.N; j++ {
				d.genInsertSQL(n)
			}
		})
	}
}
