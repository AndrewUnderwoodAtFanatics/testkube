// +build e2e

package main

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/kubeshop/kubtest/pkg/rand"
	"github.com/kubeshop/kubtest/test/e2e/kubtest"
	"github.com/stretchr/testify/require"
)

func TestE2E(t *testing.T) {
	a := require.New(t)
	test := kubtest.NewKubtest()
	scriptName := fmt.Sprintf("script-%s", rand.Name())
	collectionFile := "test.postman_collection.json"

	t.Run("install test", func(t *testing.T) {
		// given
		test.Output = "json"

		// uninstall first before installing
		test.Uninstall()

		// TODO change to watch
		sleep(t, 10*time.Second)

		// when
		out, err := test.Install()

		// then
		a.NoError(err)
		a.Contains(string(out), "STATUS: deployed")
		a.Contains(string(out), "Visit http://127.0.0.1:8080 to use your application")
	})

	// TODO change to watch for changes
	sleep(t, time.Minute)

	t.Run("scripts management", func(t *testing.T) {
		// given
		out, err := test.CreateScript(scriptName, collectionFile)
		a.NoError(err)
		a.Contains(string(out), "Script created")

		// when
		out, err = test.List()
		a.NoError(err)

		// then
		a.Contains(string(out), scriptName)
	})

	t.Run("scripts run", func(t *testing.T) {
		// given
		executionName := rand.Name()

		// when
		out, err := test.StartScript(scriptName, executionName)
		a.NoError(err)

		// then check if info about collection steps exists somewhere in output
		a.Contains(string(out), "Kasia.in Homepage")
		a.Contains(string(out), "Google")

		// then check if scripts completed with success
		a.Contains(string(out), "Script execution completed with sucess")

		executionID := GetExecutionID(out)
		t.Logf("Execution completed ID: %s", executionID)
		a.NotEmpty(executionID)

		out, err = test.Execution(scriptName, executionID)
		// check tests results for postman collection
		a.Contains(string(out), "Google")
		a.Contains(string(out), "Successful GET request")
		// check tests results for postman collection
		a.Contains(string(out), "Kasia.in Homepage")
		a.Contains(string(out), "Body matches string")
	})

	sleep(t, time.Second)

	// t.Run("cleaning helm release", func(t *testing.T) {
	// 	out, err := test.Uninstall()
	// 	a.NoError(err)
	// 	a.Contains(string(out), "uninstalled")
	// })

}

func sleep(t *testing.T, d time.Duration) {
	t.Logf("Waiting for changes for %s (because I can't watch yet :P)", d)
	time.Sleep(d)
}

func GetExecutionID(out []byte) string {
	r := regexp.MustCompile("kubectl kubtest scripts execution test ([0-9a-zA-Z]+)")
	matches := r.FindStringSubmatch(string(out))
	if len(matches) == 2 {
		return matches[1]
	}
	return ""
}
