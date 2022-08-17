package restFul

import "github.com/cucumber/godog"

func iSendRequestTo(arg1, arg2 string) error {
	return godog.ErrPending
}

func nothing() error {
	return godog.ErrPending
}

func theResponseCodeShouldBe(arg1 int) error {
	return godog.ErrPending
}

func theResponseShouldMatchJson(arg1 *godog.DocString) error {
	return godog.ErrPending
}

func theResponseTimeShouldBeLessThanSeconds(arg1 int) error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I send "([^"]*)" request to "([^"]*)"$`, iSendRequestTo)
	ctx.Step(`^nothing$`, nothing)
	ctx.Step(`^the response code should be (\d+)$`, theResponseCodeShouldBe)
	ctx.Step(`^the response should match json:$`, theResponseShouldMatchJson)
	ctx.Step(`^the response time should be less than (\d+) seconds$`, theResponseTimeShouldBeLessThanSeconds)
}
