package queue_test

import (
	"testing"

	"dijkstra-shortest-path/queue"

	"github.com/stretchr/testify/require"
)

type Num int

func NewNum(i int) *Num {
	n := Num(i)
	return &n
}

func (n Num) SortableValue() Num {
	return n
}

func TestMinPriorityQueue(t *testing.T) {
	type args struct {
		newItems []*Num
	}
	type want struct {
		items []*Num
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "tc1",
			args: args{newItems: []*Num{
				NewNum(1),
				NewNum(2),
				NewNum(3),
				NewNum(4),
				NewNum(5),
			}},
			want: want{items: []*Num{
				NewNum(1),
				NewNum(2),
				NewNum(3),
				NewNum(4),
				NewNum(5),
			}},
		},
		{
			name: "tc2",
			args: args{newItems: []*Num{
				NewNum(5),
				NewNum(4),
				NewNum(3),
				NewNum(2),
				NewNum(1),
			}},
			want: want{items: []*Num{
				NewNum(1),
				NewNum(2),
				NewNum(3),
				NewNum(4),
				NewNum(5),
			}},
		},
		{
			name: "tc3",
			args: args{newItems: []*Num{
				NewNum(1),
				NewNum(5),
				NewNum(2),
				NewNum(4),
				NewNum(3),
			}},
			want: want{items: []*Num{
				NewNum(1),
				NewNum(2),
				NewNum(3),
				NewNum(4),
				NewNum(5),
			}},
		},
		{
			name: "tc4",
			args: args{newItems: []*Num{
				NewNum(5),
				NewNum(1),
				NewNum(4),
				NewNum(2),
				NewNum(3),
			}},
			want: want{items: []*Num{
				NewNum(1),
				NewNum(2),
				NewNum(3),
				NewNum(4),
				NewNum(5),
			}},
		},
		{
			name: "tc5",
			args: args{newItems: []*Num{}},
			want: want{items: []*Num{}},
		},
		{
			name: "tc6",
			args: args{newItems: []*Num{
				NewNum(1),
			}},
			want: want{items: []*Num{
				NewNum(1),
			}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require := require.New(t)

			q := queue.NewMinPriority(func(n *Num) Num { return *n })

			// empty
			require.Equal(0, q.Size())
			require.True(q.Empty())

			// enqueue nodes
			for _, it := range tc.args.newItems {
				q.Enqueue(it)
			}

			// not empty if nodes enqueued
			require.Equal(len(tc.args.newItems), q.Size())
			if len(tc.args.newItems) > 0 {
				require.False(q.Empty())
			}

			// dequeue all nodes
			for _, expdq := range tc.want.items {
				dq := q.Dequeue()
				require.Equal(expdq, dq)
			}
			require.Nil(q.Dequeue())

			// empty after all nodes dequeued
			require.Equal(0, q.Size())
			require.True(q.Empty())
		})
	}
}
