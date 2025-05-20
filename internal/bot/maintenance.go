package bot

import (
	"context"
	"log/slog"
	"sort"
	"strconv"
	"strings"

	"github.com/ashokhin/am4bot/internal/model"
	"github.com/ashokhin/am4bot/internal/utils"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func (b *Bot) maintenance(ctx context.Context) error {
	slog.Info("start aircraft maintenance")

	slog.Debug("open pop-up window", "window", "maintenance")
	// open the "Maintenance" pop-up
	utils.DoClickElement(ctx, model.BUTTON_MAIN_MAINTENANCE)

	defer utils.DoClickElement(ctx, model.BUTTON_COMMON_CLOSE_POPUP)

	if err := b.aCheckAllAircraft(ctx); err != nil {
		slog.Warn("error in Bot.maintenance > Bot.aCheckAllAircraft", "error", err)

		return err
	}

	if err := b.repairAllAircraft(ctx); err != nil {
		slog.Warn("error in Bot.maintenance > Bot.repairAllAircraft", "error", err)

		return err
	}

	if err := b.modifyAllAircraft(ctx); err != nil {
		slog.Warn("error in Bot.maintenance > Bot.modifyAllAircraft", "error", err)

		return err
	}

	return nil
}

func (b *Bot) maintenanceAcByType(ctx context.Context, ac model.Aircraft, mntType model.MaintenanceType) (bool, error) {
	var mntOperationStr string
	var mntOperationButton string
	var mntOperationPlanButton string
	var mntOperationCostText string
	var mntOperationCost float64
	var mntOperationPerformed bool
	var acWebElemNode *cdp.Node

	switch mntType {
	case model.A_CHECK:
		mntOperationStr = "a-check"
		mntOperationButton = model.BUTTON_MAINTENANCE_A_CHECK
		mntOperationPlanButton = model.BUTTON_MAINTENANCE_PLAN_CHECK
		mntOperationCostText = model.TEXT_MAINTENANCE_A_CHECK_TOTAL_COST
	case model.REPAIR:
		mntOperationStr = "repair"
		mntOperationButton = model.BUTTON_MAINTENANCE_REPAIR
		mntOperationPlanButton = model.BUTTON_MAINTENANCE_PLAN_REPAIR
		mntOperationCostText = model.TEXT_MAINTENANCE_REPAIR_TOTAL_COST
	case model.MODIFY:
		mntOperationStr = "modify"
		mntOperationButton = model.BUTTON_MAINTENANCE_MODIFY
		mntOperationPlanButton = model.BUTTON_MAINTENANCE_PLAN_MODIFY
		mntOperationCostText = model.TEXT_MAINTENANCE_MODIFY_TOTAL_COST

	}

	slog.Debug("maintenance aircraft", "operation", mntOperationStr, "reg.number", strings.ToUpper(ac.RegNumber), "button", mntOperationButton)
	slog.Debug("get aircraft rows")

	var aircraftElemList []*cdp.Node

	if err := chromedp.Run(ctx,
		utils.ClickElement(model.BUTTON_COMMON_TAB2),
		chromedp.Nodes(model.LIST_MAINTENANCE_AC_LIST, &aircraftElemList, chromedp.ByQueryAll),
	); err != nil {
		slog.Warn("error in Bot.maintenanceAcByType > get aircraftElements list", "error", err)

		return mntOperationPerformed, err
	}

	slog.Debug("search aircraft row")

	for _, acElem := range aircraftElemList {
		if ac.RegNumber == acElem.AttributeValue("data-reg") {
			slog.Debug("row found")

			acWebElemNode = acElem

			break
		}
	}

	if acWebElemNode == nil {
		slog.Warn("aircraft row not found", "operation", mntOperationStr, "reg.number", strings.ToUpper(ac.RegNumber))

		return mntOperationPerformed, nil
	}

	slog.Debug("get cost for aircraft operation", "operation", mntOperationStr, "reg.number", strings.ToUpper(ac.RegNumber))

	// open operation window
	if err := chromedp.Run(ctx,
		chromedp.Click(mntOperationButton, chromedp.ByQuery, chromedp.FromNode(acWebElemNode)),
	); err != nil {
		slog.Warn("error in Bot.maintenanceAcByType > open operation window", "error", err)

		return mntOperationPerformed, err
	}

	// if operation is "modify" then flag all modifications
	if mntType == model.MODIFY {
		if err := chromedp.Run(ctx,
			chromedp.Click(model.CHECKBOX_MAINTENANCE_MODIFY_MOD1, chromedp.ByQuery),
			chromedp.Click(model.CHECKBOX_MAINTENANCE_MODIFY_MOD2, chromedp.ByQuery),
			chromedp.Click(model.CHECKBOX_MAINTENANCE_MODIFY_MOD3, chromedp.ByQuery),
		); err != nil {
			slog.Warn("error in Bot.maintenanceAcByType > flag 'modify' options", "error", err)

			return mntOperationPerformed, err
		}
	}

	// get final cost for maintenance operation
	if err := chromedp.Run(ctx,
		utils.GetFloatFromElement(mntOperationCostText, &mntOperationCost),
	); err != nil {
		slog.Warn("error in Bot.maintenanceAcByType > get operation cost", "error", err)

		return mntOperationPerformed, err
	}

	slog.Debug("maintenance cost", "cost", int(mntOperationCost), "operation", mntOperationStr, "reg.number", strings.ToUpper(ac.RegNumber))

	if mntOperationCost == 0 {
		slog.Debug("maintenance cost is $0")

		return mntOperationPerformed, nil
	}

	if mntOperationCost > b.Conf.BudgetMoney.Maintenance {
		slog.Warn("maintenance is too expensive", "cost", int(mntOperationCost),
			"budget", int(b.Conf.BudgetMoney.Maintenance), "operation", mntOperationStr,
			"reg.number", strings.ToUpper(ac.RegNumber))

		return mntOperationPerformed, nil
	}

	slog.Info("plan maintenance", "operation", mntOperationStr, "reg.number", strings.ToUpper(ac.RegNumber))

	if err := chromedp.Run(ctx,
		utils.ClickElement(mntOperationPlanButton),
	); err != nil {
		slog.Warn("error in Bot.maintenanceAcByType > plan maintenance operation", "error", err)

		return mntOperationPerformed, err
	}

	b.Conf.BudgetMoney.Maintenance -= mntOperationCost
	b.AccountBalance -= mntOperationCost
	mntOperationPerformed = true

	return mntOperationPerformed, nil
}

func (b *Bot) aCheckAllAircraft(ctx context.Context) error {
	var aircraftPlaned int
	var aircraftNeedACheck []model.Aircraft
	var aircraftElemList []*cdp.Node

	slog.Info("search aircraft which need A-Check")
	slog.Debug("get list of aircraftElements")

	if err := chromedp.Run(ctx,
		// open "Plan +" tab
		utils.ClickElement(model.BUTTON_COMMON_TAB2),
		// click on "Base only" button
		utils.ClickElement(model.BUTTON_MAINTENANCE_BASE_ONLY),
		// sort by "A-Check"
		utils.ClickElement(model.BUTTON_MAINTENANCE_SORT_BY_CHECK),
		// search all "aircraft" rows
		chromedp.Nodes(model.LIST_MAINTENANCE_AC_LIST, &aircraftElemList, chromedp.ByQueryAll),
	); err != nil {
		slog.Warn("error in Bot.aCheckAircraft > get aircraftElements list", "error", err)

		return err
	}

	// the "Maintenance list" element is dynamic, it means that we have to search
	// every aircraft individually

	// collect list of aircraft which need a-check
	for _, aircraftElem := range aircraftElemList {
		acACheckHours, _ := strconv.ParseFloat(aircraftElem.AttributeValue("data-hours"), 64)

		if acACheckHours > b.Conf.AircraftMaxHoursToCheck {
			slog.Debug("skip aircraft")

			continue
		}

		var aircraft model.Aircraft
		aircraft.RegNumber = aircraftElem.AttributeValue("data-reg")
		aircraft.AcType = aircraftElem.AttributeValue("data-type")
		aircraft.HoursACheck, _ = strconv.ParseFloat(aircraftElem.AttributeValue("data-hours"), 64)

		slog.Debug("add aircraft for a-check", "aircraft", aircraft)

		aircraftNeedACheck = append(aircraftNeedACheck, aircraft)
	}

	if len(aircraftNeedACheck) == 0 {
		slog.Info("no aircraft need A-Check")

		return nil
	}

	slog.Info("found aircraft for a-check", "count", len(aircraftNeedACheck))

	for _, aircraft := range aircraftNeedACheck {
		slog.Debug("try to a-check aircraft", "aircraft", aircraft)

		if mntOperationPerformed, err := b.maintenanceAcByType(ctx, aircraft, model.A_CHECK); err != nil {
			slog.Warn("error in Bot.aCheckAllAircraft > Bot.maintenanceAcByType", "error", err)

			return err
		} else if mntOperationPerformed {
			aircraftPlaned++
		}
	}

	if aircraftPlaned > 0 {
		slog.Info("aircraft a-check planed", "count", aircraftPlaned)
	}

	return nil
}

func (b *Bot) repairAllAircraft(ctx context.Context) error {
	var aircraftPlaned int
	var aircraftNeedRepair []model.Aircraft
	var aircraftElemList []*cdp.Node

	slog.Info("search aircraft which need repair")
	slog.Debug("get list of aircraftElements")

	if err := chromedp.Run(ctx,
		// open "Plan +" tab
		utils.ClickElement(model.BUTTON_COMMON_TAB2),
		// sort by "Wear"
		utils.ClickElement(model.BUTTON_MAINTENANCE_SORT_BY_WEAR),
		// search all "aircraft" rows
		chromedp.Nodes(model.LIST_MAINTENANCE_AC_LIST, &aircraftElemList, chromedp.ByQueryAll),
	); err != nil {
		slog.Warn("error in Bot.repairAllAircraft > get aircraftElements list", "error", err)

		return err
	}

	// the "Maintenance list" element is dynamic, it means that we have to search
	// every aircraft individually

	// collect list of aircraft which need repair
	for _, aircraftElem := range aircraftElemList {
		acWearPercent, _ := strconv.ParseFloat(aircraftElem.AttributeValue("data-wear"), 64)

		if acWearPercent < b.Conf.AircraftWearPercent {
			slog.Debug("skip aircraft")

			continue
		}

		var aircraft model.Aircraft
		aircraft.RegNumber = aircraftElem.AttributeValue("data-reg")
		aircraft.AcType = aircraftElem.AttributeValue("data-type")
		aircraft.WearPercent, _ = strconv.ParseFloat(aircraftElem.AttributeValue("data-wear"), 64)

		slog.Debug("add aircraft for repair", "aircraft", aircraft)

		aircraftNeedRepair = append(aircraftNeedRepair, aircraft)
	}

	if len(aircraftNeedRepair) == 0 {
		slog.Info("no aircraft need repair")

		return nil
	}

	slog.Info("found aircraft for repair", "count", len(aircraftNeedRepair))

	for _, aircraft := range aircraftNeedRepair {
		slog.Debug("try to repair aircraft", "aircraft", aircraft)

		if mntOperationPerformed, err := b.maintenanceAcByType(ctx, aircraft, model.REPAIR); err != nil {
			slog.Warn("error in Bot.repairAllAircraft > Bot.maintenanceAcByType", "error", err)

			return err
		} else if mntOperationPerformed {
			aircraftPlaned++
		}
	}

	if aircraftPlaned > 0 {
		slog.Info("aircraft repair planed", "count", aircraftPlaned)
	}

	return nil
}

func (b *Bot) modifyAllAircraft(ctx context.Context) error {
	var aircraftPlaned int
	var aircraftNeedModify []model.Aircraft
	var aircraftElemList []*cdp.Node

	slog.Info("search aircraft which need modify")
	slog.Debug("get list of aircraftElements")

	if err := chromedp.Run(ctx,
		// open "Plan +" tab
		utils.ClickElement(model.BUTTON_COMMON_TAB2),
		// click on "Base only" button
		utils.ClickElement(model.BUTTON_MAINTENANCE_BASE_ONLY),
		// search all "aircraft" rows
		chromedp.Nodes(model.LIST_MAINTENANCE_AC_LIST, &aircraftElemList, chromedp.ByQueryAll),
	); err != nil {
		slog.Warn("error in Bot.modifyAllAircraft > get aircraftElements list", "error", err)

		return err
	}

	// the "Maintenance list" element is dynamic, it means that we have to search
	// every aircraft individually

	// create "aircraft" list
	for _, aircraftElem := range aircraftElemList {

		var aircraft model.Aircraft
		aircraft.RegNumber = aircraftElem.AttributeValue("data-reg")
		aircraft.AcType = aircraftElem.AttributeValue("data-type")

		slog.Debug("add aircraft for modify", "aircraft", aircraft)

		aircraftNeedModify = append(aircraftNeedModify, aircraft)
	}

	slog.Debug("sort and slice aircraft for modify list")

	// sort "aircraft" list by reg.number and get only last "Conf.AircraftModifyLimit" number of aircraft
	sort.Slice(aircraftNeedModify, func(i, j int) bool {
		return aircraftNeedModify[i].RegNumber < aircraftNeedModify[j].RegNumber
	})

	aircraftNeedModify = aircraftNeedModify[len(aircraftNeedModify)-int(b.Conf.AircraftModifyLimit):]

	slog.Debug("sorted and sliced aircraft for modify list", "list", aircraftNeedModify)

	for _, aircraft := range aircraftNeedModify {
		slog.Debug("try to modify aircraft", "aircraft", aircraft)

		if mntOperationPerformed, err := b.maintenanceAcByType(ctx, aircraft, model.MODIFY); err != nil {
			slog.Warn("error in Bot.modifyAllAircraft > Bot.maintenanceAcByType", "error", err)

			return err
		} else if mntOperationPerformed {
			aircraftPlaned++
		}
	}

	if aircraftPlaned > 0 {
		slog.Info("aircraft modify planed", "count", aircraftPlaned)
	} else {
		slog.Info("no aircraft need modification")
	}

	return nil
}
