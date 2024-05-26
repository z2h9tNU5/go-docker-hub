package list

import (
	"testing"
)

func TestList(t *testing.T) {
	t.Parallel()

	t.Run("not exist image", func(t *testing.T) {
		t.Parallel()

		res, err := List("abcde")
		if err != nil {
			t.Fatal(err)
		}

		if res.Count != nil {
			t.Errorf("want nil")
		}
	})
	t.Run("exist image", func(t *testing.T) {
		t.Parallel()

		res, err := List("centos")
		if err != nil {
			t.Fatal(err)
		}

		if res.Count == nil || *res.Count < 1 {
			t.Errorf("want not nil")
		}
	})
}
