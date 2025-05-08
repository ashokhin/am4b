package utils

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/ashokhin/am4bot/internal/model"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

// intFromString deletes all non-digit values like words, letters, signs, spaces etc. and returns Integer value.
func intFromString(str string) (int, error) {
	var i int
	var err error

	allNumRegex := regexp.MustCompile("[0-9]+")
	str = strings.Join(allNumRegex.FindAllString(str, -1), "")
	i, err = strconv.Atoi(str)
	if err != nil {
		slog.Warn("error in utils.intFromString", "string", str, "error", err)

		return i, err
	}

	return i, nil
}

// floatFromString deletes all non-digit values like words, letters, signs, spaces etc. and returns Integer value.
func floatFromString(str string) (float64, error) {
	var f float64
	var err error

	allNumRegex := regexp.MustCompile(`[0-9]+(\.[0-9]+)?`)
	str = strings.Join(allNumRegex.FindAllString(str, -1), "")
	f, err = strconv.ParseFloat(str, 64)
	if err != nil {
		slog.Warn("error in utils.floatFromString", "string", str, "error", err)

		return f, err
	}

	return f, nil
}

func getCallerFunctionName() string {
	pc, _, _, _ := runtime.Caller(2)
	f := runtime.FuncForPC(pc)

	if f == nil {
		return ""
	}

	sliceOfFuncPath := strings.Split(f.Name(), ".")
	funcName := sliceOfFuncPath[len(sliceOfFuncPath)-1]

	return funcName
}

func RefreshPage() chromedp.Tasks {
	slog.Debug("refresh page")

	return chromedp.Tasks{
		chromedp.Reload(),
		chromedp.WaitNotVisible(model.OVERLAY_LOADING, chromedp.ByQuery),
	}
}

func Screenshot() chromedp.Tasks {
	var buf []byte

	slog.Debug("take a screenshot")

	callerFunc := getCallerFunctionName()
	scrName := fmt.Sprintf("screenshot_%s_%s.png", callerFunc, time.Now().Format("2006-01-02T15-04-05"))

	slog.Debug("save screenshot", "file", scrName)

	return chromedp.Tasks{
		chromedp.FullScreenshot(&buf, 100),
		chromedp.ActionFunc(func(ctx context.Context) error {
			if err := os.WriteFile(scrName, buf, 0644); err != nil {
				return err
			}

			return nil
		}),
	}
}

// GetIntFromElement is an element query action that retrieves the visible text of the first element
// node matching the selector and converts it to Integer.
func GetIntFromElement(sel string, resultInt *int) chromedp.Tasks {
	var resultStr string
	var err error

	slog.Debug("get integer from element", "element", sel)

	return chromedp.Tasks{
		chromedp.Text(sel, &resultStr, chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			*resultInt, err = intFromString(resultStr)
			if err != nil {
				slog.Warn("error in utils.GetIntFromElement > utils.intFromString",
					"string", resultStr, "error", err)

				return err
			}

			return nil
		}),
	}
}

func GetIntFromChildElement(sel string, resultInt *int, node *cdp.Node) chromedp.Tasks {
	var resultStr string
	var err error

	slog.Debug("get integer from child element", "element", sel)

	return chromedp.Tasks{
		chromedp.Text(sel, &resultStr, chromedp.ByQuery, chromedp.FromNode(node)),
		chromedp.ActionFunc(func(ctx context.Context) error {
			*resultInt, err = intFromString(resultStr)
			if err != nil {
				slog.Warn("error in utils.GetIntFromElement > utils.intFromString",
					"string", resultStr, "error", err)

				return err
			}

			return nil
		}),
	}
}

func GetFloatFromElement(sel string, resultFloat *float64) chromedp.Tasks {
	var resultStr string
	var err error

	slog.Debug("get float from element", "element", sel)

	return chromedp.Tasks{
		chromedp.Text(sel, &resultStr, chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			*resultFloat, err = floatFromString(resultStr)
			if err != nil {
				slog.Warn("error in utils.GetIntFromElement > utils.floatFromString",
					"string", resultStr, "error", err)

				return err
			}

			return nil
		}),
	}
}

func GetFloatFromChildElement(sel string, resultFloat *float64, node *cdp.Node) chromedp.Tasks {
	var resultStr string
	var resultInt int
	var err error

	slog.Debug("get float from child element", "element", sel)

	return chromedp.Tasks{
		chromedp.Text(sel, &resultStr, chromedp.ByQuery, chromedp.FromNode(node)),
		chromedp.ActionFunc(func(ctx context.Context) error {
			resultInt, err = intFromString(resultStr)
			if err != nil {
				slog.Warn("error in utils.GetIntFromElement > utils.intFromString",
					"string", resultStr, "error", err)

				return err
			}

			*resultFloat = float64(resultInt)

			return nil
		}),
	}
}

// ClickElement is an element query action that sends a mouse click event to the first element
// node matching the selector and waits for 1sec.
func ClickElement(sel string) chromedp.Tasks {
	slog.Debug("click element", "element", sel)

	return chromedp.Tasks{
		chromedp.Click(sel, chromedp.ByQuery),
		chromedp.Sleep(1 * time.Second),
	}
}

func DoClickElement(ctx context.Context, sel string) error {
	slog.Debug("click element", "element", sel)

	if err := chromedp.Run(ctx,
		chromedp.Click(sel, chromedp.ByQuery),
		chromedp.Sleep(1*time.Second),
	); err != nil {
		slog.Warn("error in utils.DoClickElement", "error", err)

		return err
	}

	return nil
}

func IsElementVisible(ctx context.Context, sel string) bool {
	var nodesList []*cdp.Node

	slog.Debug("check if element is visible", "element", sel)

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	slog.Debug("init nodesList", "len", len(nodesList))

	if err := chromedp.Run(ctx,
		chromedp.Nodes(sel, &nodesList, chromedp.ByQueryAll),
	); err != nil {
		// if not found for the ctx timeout then return false - element is not visible
		slog.Debug("error in utils.IsElementVisible", "error", err, "selector", sel)

		return false
	}

	slog.Debug("current nodesList", "len", len(nodesList))

	// if 1 or more elements found then return true - element is visible
	return len(nodesList) > 0

}
