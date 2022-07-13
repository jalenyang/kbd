package realm

import "testing"

func TestOperator(t *testing.T) {
	tcs := []struct {
		name string
		args []string
	}{
		{"list realm", []string{"ls"}},
		{"set realm", []string{"set", "-name", "test", "-desc", "test", "-kubeconfig", "/tmp/kubeconfig"}},
		{"use realm", []string{"use", "-name", "test"}},
		{"remove realm", []string{"rm", "-name", "test"}},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			if err := Operator(tc.args); err != nil {
				t.Fail()
			}
		})
	}
}
