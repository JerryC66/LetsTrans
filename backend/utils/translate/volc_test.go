package translate

import (
	"os"
	"testing"
)

func TestVolc_Translate(t *testing.T) {
	type args struct {
		text       string
		sourceLang string
		targetLang string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"n1",
			args{
				"hello",
				"en",
				"zh",
			},
			"你好",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Volc{}
			got, err := v.Translate(tt.args.text, tt.args.sourceLang, tt.args.targetLang)
			t.Logf("%v %v\n", os.Getenv("VOLC_AK"), os.Getenv("VOLC_SK"))

			if (err != nil) != tt.wantErr {
				t.Errorf("Translate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Translate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
