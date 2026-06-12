package ports

type ObjectStorage interface {
	Upload(key string , data []byte)(string, error)
	Download(key string)([]byte , error)
}