// @author hongjun500
// @date 2023/9/25 14:12
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description:

package main

import "testing"

func TestBuilder(t *testing.T) {
	type args struct {
		name string
		opts []ResourcePoolConfigOptFunc
	}
	tests := []struct {
		name    string
		args    args
		want    *ResourcePoolConfig
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				name: "test",
				opts: []ResourcePoolConfigOptFunc{
					func(option *ResourcePoolConfigOption) {
						option.maxTotal = 20
						option.maxIdle = 10
						option.minIdle = 1
					},
				},
			},
			want: &ResourcePoolConfig{
				name:     "test",
				maxTotal: 20,
				maxIdle:  10,
				minIdle:  1,
			},
		},
		{
			name: "err",
			args: args{
				name: "test",
				opts: []ResourcePoolConfigOptFunc{
					func(option *ResourcePoolConfigOption) {
						option.maxTotal = 10
						option.maxIdle = 20
						option.minIdle = 1
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		_, err := NewResourcePoolConfig(tt.args.name, tt.args.opts...)
		if (err != nil) != tt.wantErr {
			t.Errorf("NewResourcePoolConfig() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
	}
}
