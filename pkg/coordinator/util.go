package coordinator

import "os/exec"

func generateUUID() (string, error) {
	uuid, err := exec.Command("uuidgen").Output()
	if err != nil {
		return "", err
	}

	return string(uuid), nil
}
