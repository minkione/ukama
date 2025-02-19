package msgbus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuild(t *testing.T) {

	t.Run("basic_usage", func(t *testing.T) {
		rk, err := NewRoutingKeyBuilder().SetEventType().SetCloudSource().SetObject("some-obj").
			SetActionCreate().SetCloudSource().SetContainer("some_container").Build()
		assert.NoError(t, err)
		assert.Equal(t, "event.cloud.some_container.some-obj.create", rk)
	})

	t.Run("use_star_segment", func(t *testing.T) {
		rk, err := NewRoutingKeyBuilder().SetEventType().SetCloudSource().SetObject("some-obj").
			SetAction("*").SetCloudSource().SetContainer("some_container").Build()
		assert.NoError(t, err)
		assert.Equal(t, "event.cloud.some_container.some-obj.*", rk)
	})

	t.Run("error_missing_segment", func(t *testing.T) {
		_, err := NewRoutingKeyBuilder().SetEventType().SetCloudSource().
			SetAction("*").SetCloudSource().SetContainer("some_container").Build()
		assert.Error(t, err, "")
		assert.EqualErrorf(t, err, "object segment is not set", "")

	})

	t.Run("make_sure_new_instace_is_created", func(t *testing.T) {
		rk := NewRoutingKeyBuilder()
		rk1 := rk.SetEventType().SetCloudSource().
			SetAction("*").SetCloudSource().SetContainer("container1")

		rk2 := rk.SetEventType().SetCloudSource().
			SetAction("*").SetCloudSource().SetContainer("container2")

		assert.NotEqual(t, rk1, rk2)
	})
}
