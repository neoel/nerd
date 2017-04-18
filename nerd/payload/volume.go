package payload

//VolumeName contains the volume Name field and its validation rules
type VolumeName struct {
	Name string `json:"name" valid:"matches(^[a-z][a-z-_]{0,31}$)"`
}

//VolumeCreateInput is used as input to volume creation
type VolumeCreateInput struct {
	VolumeName
}

//VolumeCreateOutput is returned from creating a volume
type VolumeCreateOutput struct {
	Volume
}

//Volume is a volume in the list output
type Volume struct {
	VolumeName
	ProjectID string `json:"project_id"`
	Root      string `json:"root"`
}
