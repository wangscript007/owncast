package models

//StorageProvider is how a chunk storage provider should be implemented
type StorageProvider interface {
	Setup() error
	Save(filePath string, retryCount int) (string, error)
	Cleanup()

	SegmentWritten(localFilePath string)
	VariantPlaylistWritten(localFilePath string)
	MasterPlaylistWritten(localFilePath string)
}
