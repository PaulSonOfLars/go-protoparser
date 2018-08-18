package protoparser

import (
	"reflect"
	"testing"
)

func TestParseOneof(t *testing.T) {
	tests := []struct {
		name              string
		input             string
		wantOneof         *Oneof
		wantRecentScanned string
		wantErr           bool
	}{
		{
			name:      "parsing an empty",
			wantOneof: &Oneof{},
			wantErr:   true,
		},
		{
			name: "parsing a normal oneof",
			input: `
oneof foo {
    // name is the foo's name.
    string name = 4;
    // sub_messages are the optional messages.
    repeated SubMessage sub_messages = 9;
}
            `,
			wantOneof: &Oneof{
				Name: "foo",
				Fields: []*Field{
					&Field{
						Name: "name",
						Type: &Type{
							Name: "string",
						},
						Comments: []string{
							`// name is the foo's name.`,
						},
					},
					&Field{
						Name: "sub_messages",
						Type: &Type{
							Name:       "SubMessage",
							IsRepeated: true,
						},
						Comments: []string{
							`// sub_messages are the optional messages.`,
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			lex := lex(test.input)
			got, err := parseOneof(lex)
			switch {
			case test.wantErr:
				if err == nil {
					t.Errorf("got err nil, but want err")
				}
				return
			case !test.wantErr && err != nil:
				t.Errorf("got err %v, but want nil", err)
				return
			}

			if !reflect.DeepEqual(got, test.wantOneof) {
				t.Errorf("got %v, but want %v", got, test.wantOneof)
			}
			if lex.text() != test.wantRecentScanned {
				t.Errorf("got %v, but want %v", lex.text(), test.wantRecentScanned)
			}
		})
	}
}