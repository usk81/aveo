package aveo

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_lookuper_Lookup(t *testing.T) {
	type fields struct {
		env Env
	}
	type args struct {
		key string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantVal    string
		wantExists bool
	}{
		{
			name: "exists",
			fields: fields{
				env: &MapEnv{
					store: map[string]string{
						"foo": "bar",
					},
				},
			},
			args: args{
				key: "foo",
			},
			wantVal:    "bar",
			wantExists: true,
		},
		{
			name: "non-exists",
			fields: fields{
				env: &MapEnv{
					store: map[string]string{
						"foo": "bar",
					},
				},
			},
			args: args{
				key: "fizz",
			},
			wantVal:    "",
			wantExists: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			l := NewLookuper(tt.fields.env)
			gotVal, gotExists := l.Lookup(tt.args.key)
			if gotVal != tt.wantVal {
				t.Errorf("lookuper.Lookup() gotVal = %v, want %v", gotVal, tt.wantVal)
			}
			if gotExists != tt.wantExists {
				t.Errorf("lookuper.Lookup() gotExists = %v, want %v", gotExists, tt.wantExists)
			}
		})
	}
}

func TestProcess(t *testing.T) {
	type args struct {
		env    Env
		prefix string
		spec   interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// string
		{
			name: "string",
			args: args{
				env: &MapEnv{
					store: map[string]string{
						"FIELD": "foo",
					},
				},
				prefix: "",
				spec: &struct {
					Field string `env:"FIELD"`
				}{},
			},
			want: &struct {
				Field string `env:"FIELD"`
			}{
				Field: "foo",
			},
			wantErr: false,
		},
		// bool
		{
			name: "bool/true",
			args: args{
				env: &MapEnv{
					store: map[string]string{
						"FIELD": "true",
					},
				},
				prefix: "",
				spec: &struct {
					Field bool `env:"FIELD"`
				}{},
			},
			want: &struct {
				Field bool `env:"FIELD"`
			}{
				Field: true,
			},
			wantErr: false,
		},
		{
			name: "bool/false",
			args: args{
				env: &MapEnv{
					store: map[string]string{
						"FIELD": "false",
					},
				},
				prefix: "",
				spec: &struct {
					Field bool `env:"FIELD"`
				}{},
			},
			want: &struct {
				Field bool `env:"FIELD"`
			}{
				Field: false,
			},
			wantErr: false,
		},
		{
			name: "bool/error",
			args: args{
				env: &MapEnv{
					store: map[string]string{
						"FIELD": "not a valid bool",
					},
				},
				prefix: "",
				spec: &struct {
					Field bool `env:"FIELD"`
				}{},
			},
			want: &struct {
				Field bool `env:"FIELD"`
			}{},
			wantErr: true,
		},
		// float64
		{
			name: "float64/valid",
			args: args{
				env: &MapEnv{
					store: map[string]string{
						"FIELD": "1.234",
					},
				},
				prefix: "",
				spec: &struct {
					Field float64 `env:"FIELD"`
				}{},
			},
			want: &struct {
				Field float64 `env:"FIELD"`
			}{
				Field: 1.234,
			},
			wantErr: false,
		},
		{
			name: "float64/error",
			args: args{
				env: &MapEnv{
					store: map[string]string{
						"FIELD": "not a valid float64",
					},
				},
				prefix: "",
				spec: &struct {
					Field float64 `env:"FIELD"`
				}{},
			},
			want: &struct {
				Field float64 `env:"FIELD"`
			}{},
			wantErr: true,
		},
		// float32
		{
			name: "float32/valid",
			args: args{
				env: &MapEnv{
					store: map[string]string{
						"FIELD": "1.234",
					},
				},
				prefix: "",
				spec: &struct {
					Field float32 `env:"FIELD"`
				}{},
			},
			want: &struct {
				Field float32 `env:"FIELD"`
			}{
				Field: 1.234,
			},
			wantErr: false,
		},
		{
			name: "float32/error",
			args: args{
				env: &MapEnv{
					store: map[string]string{
						"FIELD": "not a valid float32",
					},
				},
				prefix: "",
				spec: &struct {
					Field float32 `env:"FIELD"`
				}{},
			},
			want: &struct {
				Field float32 `env:"FIELD"`
			}{},
			wantErr: true,
		},
		// int
		{
			name: "int/valid",
			args: args{
				env: &MapEnv{
					store: map[string]string{
						"FIELD": "1234",
					},
				},
				prefix: "",
				spec: &struct {
					Field int `env:"FIELD"`
				}{},
			},
			want: &struct {
				Field int `env:"FIELD"`
			}{
				Field: 1234,
			},
			wantErr: false,
		},
		{
			name: "int/error",
			args: args{
				env: &MapEnv{
					store: map[string]string{
						"FIELD": "not a valid int",
					},
				},
				prefix: "",
				spec: &struct {
					Field int `env:"FIELD"`
				}{},
			},
			want: &struct {
				Field int `env:"FIELD"`
			}{},
			wantErr: true,
		},
		// int64
		{
			name: "int64/valid",
			args: args{
				env: &MapEnv{
					store: map[string]string{
						"FIELD": "1234",
					},
				},
				prefix: "",
				spec: &struct {
					Field int64 `env:"FIELD"`
				}{},
			},
			want: &struct {
				Field int64 `env:"FIELD"`
			}{
				Field: 1234,
			},
			wantErr: false,
		},
		{
			name: "int64/error",
			args: args{
				env: &MapEnv{
					store: map[string]string{
						"FIELD": "not a valid int",
					},
				},
				prefix: "",
				spec: &struct {
					Field int64 `env:"FIELD"`
				}{},
			},
			want: &struct {
				Field int64 `env:"FIELD"`
			}{},
			wantErr: true,
		},
		// int32
		{
			name: "int32/valid",
			args: args{
				env: &MapEnv{
					store: map[string]string{
						"FIELD": "1234",
					},
				},
				prefix: "",
				spec: &struct {
					Field int32 `env:"FIELD"`
				}{},
			},
			want: &struct {
				Field int32 `env:"FIELD"`
			}{
				Field: 1234,
			},
			wantErr: false,
		},
		{
			name: "int32/error",
			args: args{
				env: &MapEnv{
					store: map[string]string{
						"FIELD": "not a valid int",
					},
				},
				prefix: "",
				spec: &struct {
					Field int32 `env:"FIELD"`
				}{},
			},
			want: &struct {
				Field int32 `env:"FIELD"`
			}{},
			wantErr: true,
		},
		// Lookup prefixes
		{
			name: "string",
			args: args{
				env: &MapEnv{
					store: map[string]string{
						"PREFIX_FIELD": "foo",
					},
				},
				prefix: "PREFIX_",
				spec: &struct {
					Field string `env:"FIELD"`
				}{},
			},
			want: &struct {
				Field string `env:"FIELD"`
			}{
				Field: "foo",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Process(tt.args.env, context.Background(), tt.args.prefix, tt.args.spec); (err != nil) != tt.wantErr {
				t.Errorf("Process() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Printf("%s : %#v\n", tt.name, tt.args.spec)

			if diff := cmp.Diff(tt.args.spec, tt.want); diff != "" {
				t.Errorf("mismatch (-want, +got):\n%s", diff)
			}
		})
	}
}
