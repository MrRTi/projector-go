package cli_test

import (
	"testing"

	"github.com/MrRTi/projector-go/pkg/cli"
)

func getData() *cli.Data {
	return &cli.Data{
		Projector: map[string]map[string]string{
			"/": {
				"foo": "bar1",
				"fem": "baz",
			},
			"/foo": {
				"foo": "bar2",
			},
			"/foo/bar": {
				"foo": "bar3",
			},
		},
	}
}

func getProjector(pwd string, data *cli.Data) *cli.Projector {
	return cli.CreateProjector(
		&cli.Config{
			Args:      []string{},
			Operation: cli.Print,
			Pwd:       pwd,
			Config:    "Hello!",
		},
		data,
	)
}

func test(t *testing.T, proj *cli.Projector, key, expValue string) {
	value, ok := proj.GetValue(key)
	if !ok {
		t.Error("expected to find value", key)
	}

	if value != expValue {
		t.Error("expected to find", expValue, "but received", value)
	}

}

func TestGetValue(t *testing.T) {
	data := getData()
	projector := getProjector("/foo/bar", data)

	test(t, projector, "foo", "bar3")
	test(t, projector, "fem", "baz")
}

func TestSetValue(t *testing.T) {
	data := getData()
	projector := getProjector("/foo/bar", data)

	test(t, projector, "foo", "bar3")

	projector.SetValue("foo", "bar4")
	test(t, projector, "foo", "bar4")

	projector.SetValue("fem", "bazbaz")
	test(t, projector, "fem", "bazbaz")

	projector = getProjector("/", data)
	test(t, projector, "fem", "baz")
}

func TestRemoveValue(t *testing.T) {
	data := getData()
	projector := getProjector("/foo/bar", data)

	test(t, projector, "foo", "bar3")
	projector.RemoveValue("foo")
	test(t, projector, "foo", "bar2")

	projector.RemoveValue("fem")
	test(t, projector, "fem", "baz")
}
