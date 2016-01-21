package web
import "testing"

func Test_Generate(t *testing.T) {
	cert_path := "cert.pem"
	key_path := "key.pem"

	err := Generate(cert_path, key_path, "localhost:8080")
	if err != nil {
		t.Log(err)
	}
}
