package internal


type FileData struct{
	filename string
	filetype string
	filesize int64
}

func NewFileData(filename string, filetype string, filesize int64) *FileData{
	filedata := FileData{
		filename: filename,
		filetype: filetype,
		filesize: filesize,
	}

	return &filedata
}

