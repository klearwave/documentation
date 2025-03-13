package models

import (
	"reflect"
	"testing"

	"github.com/klearwave/service-info/pkg/models"
)

func TestVersion_setVersions(t *testing.T) {
	t.Parallel()

	type fields struct {
		Model           models.Model
		VersionBase     VersionBase
		Stable          bool
		XVersion        int
		YVersion        int
		ZVersion        int
		BuildVersion    string
		ContainerImages []*ContainerImage
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
		want    *Version
	}{
		{
			name: "fail: missing version id",
			fields: fields{
				VersionBase: VersionBase{
					VersionId: "",
				},
			},
			wantErr: true,
		},
		{
			name: "fail: invalid version id",
			fields: fields{
				VersionBase: VersionBase{
					VersionId: "v1.2",
				},
			},
			wantErr: true,
		},
		{
			name: "success: x/y/z versions are set",
			fields: fields{
				VersionBase: VersionBase{
					VersionId: "v1.2.3",
				},
			},
			want: &Version{
				VersionBase: VersionBase{
					VersionId: "v1.2.3",
				},
				XVersion: 1,
				YVersion: 2,
				ZVersion: 3,
				Stable:   true,
			},
		},
		{
			name: "success: build version is set",
			fields: fields{
				VersionBase: VersionBase{
					VersionId: "v1.2.3-prerelease.1",
				},
			},
			want: &Version{
				VersionBase: VersionBase{
					VersionId: "v1.2.3-prerelease.1",
				},
				XVersion:     1,
				YVersion:     2,
				ZVersion:     3,
				BuildVersion: "prerelease.1",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			version := &Version{
				Model:           tt.fields.Model,
				VersionBase:     tt.fields.VersionBase,
				Stable:          tt.fields.Stable,
				XVersion:        tt.fields.XVersion,
				YVersion:        tt.fields.YVersion,
				ZVersion:        tt.fields.ZVersion,
				BuildVersion:    tt.fields.BuildVersion,
				ContainerImages: tt.fields.ContainerImages,
			}
			if err := version.setVersioning(); (err != nil) != tt.wantErr {
				t.Errorf("Version.setVersions() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.want != nil && !reflect.DeepEqual(*tt.want, *version) {
				t.Errorf("want %+v, got %+v", tt.want, version)
			}
		})
	}
}
