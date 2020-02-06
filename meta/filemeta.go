package meta

type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas = make(map[string]FileMeta)


func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}


func UpdateFileMeta(fMeta FileMeta) {
	fileMetas[fMeta.FileSha1] = fMeta
}

func RemoveFileMeta(fileSha1 string) {
	delete(fileMetas, fileSha1)
}


