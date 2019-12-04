package eventhub

import (
	"testing"
	"fmt"
)

var seqsm = map[int][]int{}

func TestSubscribe(t *testing.T) {
	fmt.Printf("Subscribe start!\n")
	for i := 0; i < 100; i++ {
		seq := Subscribe(i, func(args ...interface{}) {
			fmt.Printf("Execute handle args=%v\n", args)
		})
		seqsm[i] = append(seqsm[i], seq)
	}
	fmt.Printf("Subscribe end!\n")
}

func TestUnsubscribe(t *testing.T) {
	fmt.Printf("Unsubscribe start!\n")
	for i := 30; i <= 80; i++ {
		if seqs, ok := seqsm[i]; ok {
			for _, seq := range seqs {
				Unsubscribe(i, seq)
				fmt.Printf("Unsubscribe eventId=%d seq=%d\n", i, seq)
			}
		}
	}
	fmt.Printf("Unsubscribe start!\n")
}

func TestPublish(t *testing.T) {
	fmt.Printf("Publish start!\n")
	for i := 0; i < 100; i++ {
		msg := fmt.Sprintf("eventhId=%d", i)
		Publish(i, msg)
	}
	fmt.Printf("Publish end!\n")
}