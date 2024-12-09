package bot

import (
	"fmt"
	"testing"
	"time"

	"github.com/avast/retry-go/v4"
)

func TestRetry(t *testing.T) {
	n := 3
	i := 0
	begin := time.Now()
	err := retry.Do(func() error {
		fmt.Printf("%s\n", time.Since(begin))
		begin = time.Now()
		if i != n {
			fmt.Printf("i: %d\n", i)
			i++
			return fmt.Errorf("not equal")
		}
		return nil
	},
		retry.Delay(time.Second),
	)
	if err != nil {
		t.Fatal(err)
	}
}
