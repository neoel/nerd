package payload

//VolumeCreateInput is used as input to volume creation
type VolumeCreateInput struct {
	Name string `json:"name"`
}

//VolumeCreateOutput is returned from creating a volume
type VolumeCreateOutput struct {
	Volume
}

//Volume is a volume in the list output
type Volume struct {
	ProjectID string `json:"project_id"`
	Name      string `json:"name"`
	Root      string `json:"root"`
}
