package meta

import (
	"storage/db"
)

type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

func GetFileMeta(fileSha1 string) (FileMeta, error) {
	tfile, err := db.GetFileMetaDB(fileSha1)
	if err != nil {
		return FileMeta{}, err
	}

	fmeta := FileMeta{
		FileSha1: tfile.FileHash,
		FileName: tfile.FileName.String,
		FileSize: tfile.FileSize.Int64,
		Location: tfile.FilePath.String,
	}
	return fmeta, nil
}

func SaveFileMeta(fMeta FileMeta) bool {
	return db.SaveFileMetaDB(fMeta.FileSha1, fMeta.FileName,
		fMeta.FileSize, fMeta.Location)
}

func UpdateFileMeta(fMeta FileMeta) bool {
	return db.UpdateFileMetaDB(fMeta.FileName, fMeta.FileSha1)
}

func RemoveFileMeta(fileSha1 string) {
	db.DelFileMetaDB(fileSha1)
}


