package cli_test

import (
	"github.com/lpreaux/projector-go/pkg/cli"
	"testing"
)

func getData() *cli.Data {
	return &cli.Data{
		Projector: map[string]map[string]string{
			"/": {
				"foo": "bar1",
				"fem": "is_great",
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
			Config:    "Hello, Frontend Masters",
			Pwd:       pwd,
		},
		data,
	)
}

func test(t *testing.T, proj *cli.Projector, key, value string) {
	v, ok := proj.GetValue(key)
	if !ok {
		t.Errorf("expected to find value \"%v\"", value)
	}
	if value != v {
		t.Errorf("expected to find %v but received %v", value, v)
	}
}

func TestGetValue(t *testing.T) {
	data := getData()
	proj := getProjector("/foo/bar", data)

	test(t, proj, "foo", "bar3")
	test(t, proj, "fem", "is_great")
}

func TestSetValue(t *testing.T) {
	data := getData()
	proj := getProjector("/foo/bar", data)

	test(t, proj, "foo", "bar3")
	test(t, proj, "fem", "is_great")
	proj.SetValue("foo", "baz")
	proj.SetValue("fem", "is_super_great")
	test(t, proj, "foo", "baz")
	test(t, proj, "fem", "is_super_great")

	proj = getProjector("/", data)
	test(t, proj, "fem", "is_great")
}

func TestRemoveValue(t *testing.T) {
	data := getData()
	proj := getProjector("/foo/bar", data)

	test(t, proj, "foo", "bar3")
	proj.RemoveValue("foo")
	test(t, proj, "foo", "bar2")

	proj.RemoveValue("fem")
	test(t, proj, "fem", "is_great")
}
