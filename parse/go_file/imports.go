package go_file

var typeRefImport = map[string]string{
	"time.Time": "time",
}

func RegisterType(k, v string) {
	typeRefImport[k] = v
}

func TransferType(k string) string {
	return typeRefImport[k]
}
