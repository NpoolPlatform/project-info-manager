package api

import (
	"testing"

	_ "github.com/NpoolPlatform/go-service-framework/pkg/version"
)

func TestVersion(t *testing.T) {
	// if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
	// 	return
	// }

	// cli := resty.New()
	// resp, err := cli.R().
	// 	Post("http://localhost:50310/version")
	// if assert.Nil(t, err) {
	// 	assert.Equal(t, 200, resp.StatusCode())
	// 	// we should compare body, but we cannot do here
	// 	_, err := version.GetVersion()
	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, ver, string(resp.Body()))
	// }
}
