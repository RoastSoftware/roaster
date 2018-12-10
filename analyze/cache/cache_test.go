package cache_test

import (
	"testing"
	"time"

	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/analyze/cache"
	"github.com/LuleaUniversityOfTechnology/2018-project-roaster/model"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

const code = `
def test(ξ):
    ξ is None 
    thisfunctiondoesnotexist("Expect an error.")

def too_complex():
    def b():
        def c():
            pass
        c()
    b()
`

var result = model.RoastResult{
	Errors: []model.RoastError{
		{
			RoastMessage: model.RoastMessage{
				Hash:        uuid.UUID{},
				Row:         0x4,
				Column:      0x5,
				Engine:      "flake8",
				Name:        "F821",
				Description: "undefined name 'thisfunctiondoesnotexist'",
			},
		},
	},
	Warnings: []model.RoastWarning{
		{
			RoastMessage: model.RoastMessage{
				Hash:        uuid.UUID{},
				Row:         0x3,
				Column:      0xe,
				Engine:      "flake8",
				Name:        "W291",
				Description: "trailing whitespace",
			},
		},
	},
}

func TestCacheSimpleGetAndSet(t *testing.T) {
	c := cache.New(1*time.Hour, 1*time.Hour)

	r, ok := c.Get(code)
	assert.Equal(t, false, ok)

	c.Set(code, result)
	r, ok = c.Get(code)
	assert.Equal(t, true, ok)
	assert.Equal(t, r, result)
}

func TestCacheCleanUp(t *testing.T) {
	// Set a _very_ low cache retention and clean up interval.
	c := cache.New(10*time.Millisecond, 1*time.Millisecond)
	c.Set(code, result)

	// Allow the clean routine to remove the object.
	time.Sleep(15 * time.Millisecond)

	_, ok := c.Get(code)
	assert.Equal(t, false, ok) // Expect no object in the cache.
}

func BenchmarkCacheGet(b *testing.B) {
	c := cache.New(1*time.Minute, 1*time.Second)
	c.Set(code, result)

	for n := 0; n < b.N; n++ {
		_, ok := c.Get(code)
		if !ok {
			b.Error("failed to lookup key in cache")
		}
	}
}

func BenchmarkCacheSet(b *testing.B) {
	c := cache.New(1*time.Minute, 1*time.Second)

	for n := 0; n < b.N; n++ {
		c.Set(code+string(n), result)
	}
}
