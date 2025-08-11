package repository_test

import (
	"testing"

	"github.com/javonnem/web_server/http/internal/myapp/repository"
	"github.com/stretchr/testify/assert"
)


func RepositoryIntegrationTest(t *testing.T) {
	type dependencies struct {
	}
	type want struct {
		error error
	}
	test := []struct {
		name string
		dependencies dependencies
		want want
	} {
		{
			name: "default test",
			dependencies: dependencies{},
			want: want{},
		},
	}

	for _, tt :=range test {
		t.Run(tt.name, func(t *testing.T) {
			repo := repository.NewRepository()
			err := repo.GetSomething()
			if tt.want.error != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
