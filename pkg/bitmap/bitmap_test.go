package bitmap

import "testing"

func TestBitmap_Set(t *testing.T) {
	type fields struct {
		bits []byte
		size int
	}
	type args struct {
		id string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "",
			fields: fields{
				bits: make([]byte, 256),
				size: 256 * 8,
			},
			args: args{"0x0000000000000001"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bitmap{
				bits: tt.fields.bits,
				size: tt.fields.size,
			}
			b.Set(tt.args.id)
		})
	}
}
