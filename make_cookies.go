package app

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

type Dough struct {
	Ready bool
}

type Cookie struct {
	Baked bool
}

func MakeCookies(ctx workflow.Context) ([]Cookie, error) {
	plate := []Cookie{}

	dough := Dough{}

	ao := workflow.ActivityOptions{StartToCloseTimeout: time.Second * 30}
	ctx = workflow.WithActivityOptions(ctx, ao)

	err := workflow.ExecuteActivity(ctx, MakeDough).Get(ctx, &dough)
	if err != nil {
		return plate, err
	}

	err = workflow.ExecuteActivity(ctx, ShapeCookies, dough).Get(ctx, &plate)
	if err != nil {
		return plate, err
	}

	err = workflow.ExecuteActivity(ctx, BakeCookies, plate).Get(ctx, &plate)
	if err != nil {
		return plate, err
	}

	for _, cookie := range plate {
		if cookie.Baked {
			println("\033[31m", "This cookie is baked", "\033[0m")
		} else {
			println("\033[31m", "This cookie is not baked", "\033[0m")
		}
	}
	println("\033[31m", fmt.Sprintf("The plate contains %d cookies", len(plate)), "\033[0m")

	return plate, nil
}

func MakeDough(ctx context.Context) (Dough, error) {
	println("\033[31m", "1) Making dough !", "\033[0m")
	dough := Dough{Ready: true}
	time.Sleep(time.Second * 5)
	println("\033[31m", "1) Dough is ready", "\033[0m")
	return dough, nil
}

func ShapeCookies(ctx context.Context, dough Dough) ([]Cookie, error) {
	//return []Cookie{}, errors.New("shape_error")
	//os.Exit(-1)
	cookies := []Cookie{}
	println("\033[31m", "2) Splitting cookies !", "\033[0m")
	if dough.Ready {
		for i := 0; i < 12; i++ {
			cookies = append(cookies, Cookie{})
		}
		time.Sleep(time.Second * 5)
		println("\033[31m", "2) cookies are splitted", "\033[0m")
		return cookies, nil
	}
	return []Cookie{}, errors.New("shape_error")
}

func BakeCookies(ctx context.Context, plate []Cookie) ([]Cookie, error) {
	println("\033[31m", "3) Baking cookies !", "\033[0m")
	for i, cookie := range plate {
		cookie.Baked = true
		plate[i] = cookie
	}
	time.Sleep(time.Second * 5)
	println("\033[31m", "3) Cookies are baked", "\033[0m")
	return plate, nil
}
