package aveo

import (
	"reflect"
	"testing"
)

func sliceStringEqual(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for _, x := range s1 {
		ok := false
		for _, y := range s2 {
			if x == y {
				ok = true
			}
		}
		if !ok {
			return false
		}
	}
	return true
}

func TestNewMap(t *testing.T) {
	tests := []struct {
		name string
		want Env
	}{
		{
			name: "new",
			want: &MapEnv{
				store: map[string]string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapEnv_Environ(t *testing.T) {
	type fields struct {
		store map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "single",
			fields: fields{
				store: map[string]string{
					"foo": "bar",
				},
			},
			want: []string{
				"foo=bar",
			},
		},
		{
			name: "single",
			fields: fields{
				store: map[string]string{
					"foo":  "bar",
					"fizz": "bizz",
				},
			},
			want: []string{
				"foo=bar",
				"fizz=bizz",
			},
		},
		{
			name: "empty",
			fields: fields{
				store: map[string]string{},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m := &MapEnv{
				store: tt.fields.store,
			}
			if got := m.Environ(); !sliceStringEqual(got, tt.want) {
				t.Errorf("MapEnv.Environ() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapEnv_Clearenv(t *testing.T) {
	type fields struct {
		store map[string]string
	}
	tests := []struct {
		name         string
		fields       fields
		beforeLength int
		afterLength  int
	}{
		{
			name: "single",
			fields: fields{
				store: map[string]string{
					"foo": "bar",
				},
			},
			beforeLength: 1,
			afterLength:  0,
		},
		{
			name: "multiple",
			fields: fields{
				store: map[string]string{
					"foo":  "bar",
					"fizz": "bizz",
				},
			},
			beforeLength: 2,
			afterLength:  0,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m := &MapEnv{
				store: tt.fields.store,
			}
			if l := len(m.Environ()); l != tt.beforeLength {
				t.Errorf("Clearenv : beforeLength: want %d, actual: %d", tt.beforeLength, l)
			}
			m.Clearenv()
			if l := len(m.Environ()); l != tt.afterLength {
				t.Errorf("Clearenv : afterLength: want %d, actual: %d", tt.afterLength, l)
			}
		})
	}
}

func TestMapEnv_ExpandEnv(t *testing.T) {
	type fields struct {
		store map[string]string
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "expand",
			fields: fields{
				store: map[string]string{
					"golang": "awesome!!",
				},
			},
			args: args{
				s: "Golang is $golang",
			},
			want: "Golang is awesome!!",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m := &MapEnv{
				store: tt.fields.store,
			}
			if got := m.ExpandEnv(tt.args.s); got != tt.want {
				t.Errorf("MapEnv.ExpandEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapEnv_Getenv(t *testing.T) {
	type fields struct {
		store map[string]string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "existing key",
			fields: fields{
				store: map[string]string{
					"foo": "bar",
				},
			},
			args: args{
				key: "foo",
			},
			want: "bar",
		},
		{
			name: "existing key",
			fields: fields{
				store: map[string]string{
					"fizz": "bizz",
				},
			},
			args: args{
				key: "foo",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MapEnv{
				store: tt.fields.store,
			}
			if got := m.Getenv(tt.args.key); got != tt.want {
				t.Errorf("MapEnv.Getenv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapEnv_LookupEnv(t *testing.T) {
	type fields struct {
		store map[string]string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue string
		wantExist bool
	}{
		{
			name: "existing key",
			fields: fields{
				store: map[string]string{
					"foo": "bar",
				},
			},
			args: args{
				key: "foo",
			},
			wantValue: "bar",
			wantExist: true,
		},
		{
			name: "existing key",
			fields: fields{
				store: map[string]string{
					"fizz": "bizz",
				},
			},
			args: args{
				key: "foo",
			},
			wantValue: "",
			wantExist: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MapEnv{
				store: tt.fields.store,
			}
			got, got1 := m.LookupEnv(tt.args.key)
			if got != tt.wantValue {
				t.Errorf("MapEnv.LookupEnv() got = %v, want %v", got, tt.wantValue)
			}
			if got1 != tt.wantExist {
				t.Errorf("MapEnv.LookupEnv() got1 = %v, want %v", got1, tt.wantExist)
			}
		})
	}
}

func TestMapEnv_MapEnvs(t *testing.T) {
	type fields struct {
		store map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		{
			name: "single",
			fields: fields{
				store: map[string]string{
					"foo": "bar",
				},
			},
			want: map[string]string{
				"foo": "bar",
			},
		},
		{
			name: "single",
			fields: fields{
				store: map[string]string{
					"foo":  "bar",
					"fizz": "bizz",
				},
			},
			want: map[string]string{
				"foo":  "bar",
				"fizz": "bizz",
			},
		},
		{
			name: "empty",
			fields: fields{
				store: map[string]string{},
			},
			want: map[string]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MapEnv{
				store: tt.fields.store,
			}
			if got := m.MapEnvs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapEnv.MapEnvs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapEnv_Setenv(t *testing.T) {
	type fields struct {
		store map[string]string
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "new value",
			fields: fields{
				store: map[string]string{},
			},
			args: args{
				key:   "foo",
				value: "bar",
			},
			wantErr: false,
		},
		{
			name: "overwrite",
			fields: fields{
				store: map[string]string{
					"golang is": "bad",
				},
			},
			args: args{
				key:   "golang is",
				value: "awesome",
			},
			wantErr: false,
		},
		{
			name: "empty value",
			fields: fields{
				store: map[string]string{},
			},
			args: args{
				key:   "key",
				value: "",
			},
			wantErr: false,
		},
		{
			name: "empty key",
			fields: fields{
				store: map[string]string{},
			},
			args: args{
				key:   "",
				value: "value",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			m := &MapEnv{
				store: tt.fields.store,
			}
			if err := m.Setenv(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("MapEnv.Setenv() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if got := m.Getenv(tt.args.key); got != tt.args.value {
				t.Errorf("MapEnv.MapEnvs() = %v, want %v", got, tt.args.value)
			}
		})
	}
}

func TestMapEnv_Unsetenv(t *testing.T) {
	type fields struct {
		store map[string]string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		checkKey string
		exist    bool
		wantErr  bool
	}{
		{
			name: "existing key",
			fields: fields{
				store: map[string]string{
					"foo": "bar",
				},
			},
			args: args{
				key: "foo",
			},
			checkKey: "foo",
			exist:    false,
			wantErr:  false,
		},
		{
			name: "unexisting key",
			fields: fields{
				store: map[string]string{
					"foo": "bar",
				},
			},
			args: args{
				key: "fizz",
			},
			checkKey: "foo",
			exist:    true,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m := &MapEnv{
				store: tt.fields.store,
			}
			if err := m.Unsetenv(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("MapEnv.Unsetenv() error = %v, wantErr %v", err, tt.wantErr)
			}
			if _, got := m.LookupEnv(tt.checkKey); got != tt.exist {
				t.Errorf("MapEnv.Unsetenv()  %v, want %v", got, tt.exist)
			}
		})
	}
}
