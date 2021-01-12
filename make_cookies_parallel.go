package app

import (
	"errors"
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

func MakeCookiesParallelDough(ctx workflow.Context) error {
	plate := []Cookie{}

	dough := Dough{}

	ao := workflow.ActivityOptions{StartToCloseTimeout: time.Second * 30}
	ctx = workflow.WithActivityOptions(ctx, ao)

	err := workflow.ExecuteActivity(ctx, MakeDough).Get(ctx, &dough)
	if err != nil {
		return err
	}

	// Start Parallel executions
	futures := []workflow.Future{}
	for i := 0; i < 3; i++ {
		futures = append(futures, workflow.ExecuteActivity(ctx, ShapeCookies, dough))
	}

	// Get Parallel executions
	plates := [][]Cookie{}
	for _, f := range futures {
		if f.Get(ctx, &plate) != nil {
			return err
		}
		plates = append(plates, plate)
	}

	// Start Parallel executions
	futures = []workflow.Future{}
	for _, plate := range plates {
		futures = append(futures, workflow.ExecuteActivity(ctx, BakeCookies, plate))
	}

	// Get Parallel executions
	plates = [][]Cookie{}
	for _, f := range futures {
		if f.Get(ctx, &plate) != nil {
			return err
		}
		plates = append(plates, plate)
	}

	// Check that all cookies in all plates are baked
	for _, plate := range plates {
		for _, cookie := range plate {
			if !cookie.Baked {
				println("\033[31m", "At least one cookie is not baked")
				return errors.New("not_all_cookies_baked")
			}
		}
	}

	println("\033[31m", fmt.Sprintf("%d cookies have been produced", len(plates)*len(plate)), "\033[0m")

	return nil
}
