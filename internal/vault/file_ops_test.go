package vault

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "vaulttest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	tmpfilePath := tmpfile.Name()
	defer tmpfile.Close()

	data := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}

	passphrase := "secretpassword"

	// Encode the data
	err = Encode(tmpfilePath, passphrase, data)
	if err != nil {
		t.Errorf("Encode failed: %v", err)
	}

	// Decode the data
	decodedData, err := Decode(tmpfilePath, passphrase)
	if err != nil {
		t.Errorf("Decode failed: %v", err)
	}

	var decodedMap map[string]interface{}
	if err := json.Unmarshal(decodedData, &decodedMap); err != nil {
		t.Errorf("Failed to unmarshal decoded data: %v", err)
	}

	var expectedMap map[string]interface{}
	expectedJSON := []byte(`{"key1":"value1","key2":"value2"}`)
	if err := json.Unmarshal(expectedJSON, &expectedMap); err != nil {
		t.Errorf("Failed to unmarshal expected data: %v", err)
	}

	if !reflect.DeepEqual(decodedMap, expectedMap) {
		t.Errorf("Decoded data doesn't match expected data")
	}
}
