package utility

func S3FileUrl(path string) string {
	if path == "" {
		return ""
	}
	return S3Url + path
}
