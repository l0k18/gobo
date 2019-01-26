package args

import "testing"

func TestDeclaration(t *testing.T) {
	_ = []interface{}{
		[]CmdS{
			Cmd(
				"ctl", "c",
				func(a []interface{}) error {
					return nil
				},
				[]string{
					"node",
				},
				[]interface{}{
					Trigger(
						"version", "v",
						func(a []interface{}) error {
							return nil
						},
						[]string{
							"",
						},
					),
				},
			),
			Cmd(
				"node", "n",
				func(a []interface{}) error {
					return nil
				},
				[]string{
					"cmd",
				},
				[]interface{}{
					Trigger(
						"version", "v",
						func(a []interface{}) error {
							return nil
						},
						[]string{
							"",
						},
					),
				},
			),
		},
	}

}
