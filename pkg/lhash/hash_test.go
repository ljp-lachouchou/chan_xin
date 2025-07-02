package lhash

import "testing"

func TestValidatePassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"weak", // 太短
			args{},
			true,
		},
		{"longpassword", // 缺数字和大写
			args{},
			true,
		},
		{"Password123", // 缺特殊字符
			args{},
			true,
		},
		{"Secure!Pass9", // 有效
			args{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidatePassword(tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("ValidatePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
