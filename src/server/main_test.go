package main

import (
	"os/exec"
	"testing"
)

func TestIntegrationA4(t *testing.T) {

	serverListPath := "./servers.txt"
	file := "./javaTestClients/a4_2024_dummy_tests_v1.jar"
	err := testJarWithPort(file, serverListPath)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	// TODO: Read Log file and check for word FAILED in it

}

func testJarWithPort(file string, serverListPath string) error {
	cmd := exec.Command("java", "-jar", file, "--servers-list", serverListPath)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
