package repositories

import (
	"database/sql"
	"mysql-metadata/pkg/proto_models"
)

func ScanIds(res *sql.Rows) (proto_models.IdList, error) {
	var message proto_models.IdList
	for res.Next() {
		var val int64
		err := res.Scan(&val)
		if err != nil {
			return message, err
		}
		message.Id = append(message.Id, val)
	}
	return message, nil
}

func ScanStrings(res *sql.Rows) (proto_models.StringList, error) {
	var message proto_models.StringList
	for res.Next() {
		var val string
		err := res.Scan(&val)
		if err != nil {
			return message, err
		}
		message.Str = append(message.Str, val)
	}
	return message, nil
}

func ScanCarouselBanner(res *sql.Rows) ([]proto_models.CarouselBanner, error) {
	var message []proto_models.CarouselBanner
	for res.Next() {
		var val proto_models.CarouselBanner
		err := res.Scan(&val.Id, &val.BannerId, &val.Name, &val.Data, &val.StatusId, &val.CreatedAt, &val.UpdatedAt)
		if err != nil {
			return message, err
		}
		message = append(message, val)
	}
	return message, nil
}

func ScanDhxAffiliateOffers(res *sql.Rows) ([]proto_models.Affiliate, error) {
	var message []proto_models.Affiliate
	for res.Next() {
		var val proto_models.Affiliate
		err := res.Scan(&val.RequirementId, &val.UserId, &val.EventInfo)
		if err != nil {
			return message, err
		}
		message = append(message, val)
	}
	return message, nil
}

func ScanDhxEntityInfo(res *sql.Rows) (proto_models.EntityInfo, error) {
	var message proto_models.EntityInfo
	for res.Next() {
		var val int64
		var str string
		err := res.Scan(&val, &str)
		if err != nil {
			return message, err
		}
		message.Id = append(message.Id, val)
		message.PrimaryEventName = append(message.PrimaryEventName, str)
	}
	return message, nil
}

func ScanBanner(res *sql.Rows) ([]proto_models.Banner, error) {
	var message []proto_models.Banner
	for res.Next() {
		var val proto_models.Banner
		err := res.Scan(&val.Id, &val.CampaignId, &val.TypeId, &val.SubtypeId, &val.Name, &val.Data, &val.Ecpm, &val.Weightage)
		if err != nil {
			return nil, err
		}
		message = append(message, val)
	}
	return message, nil
}

func ScanCampaign(res *sql.Rows) ([]proto_models.Campaign, error) {
	var message []proto_models.Campaign
	for res.Next() {
		var val proto_models.Campaign
		err := res.Scan(&val.Id, &val.Name, &val.OrderId, &val.GroupId, &val.ReportGroupId, &val.StatusId, &val.HasRep,
			&val.BillOn, &val.BillingTypeId, &val.BillingAmount, &val.BookedInventory, &val.AddOnPercentage,
			&val.AddOnReason, &val.RejectionReason, &val.InventoryToDeliver, &val.BillingRate, &val.BudgetType,
			&val.BudgetValue, &val.DailyCap, &val.DailyImpressionCap, &val.DailyClickCap, &val.DailyInstallCap,
			&val.FrequencyCap, &val.FrequencyCapAuto, &val.ResetTimerSecondsAuto, &val.ReleaseDate, &val.EndDate,
			&val.Ecpm, &val.VEcpm, &val.VCpi, &val.Iep, &val.Cep, &val.CepMetadata, &val.HasCep, &val.IsValueAdd,
			&val.Notes, &val.SubtypeId, &val.IsBooked, &val.TargetingParameters, &val.SubtargetingParameters,
			&val.SystemTargetingParameters, &val.CampaignProperties, &val.VersionCapEnabled, &val.SubSlots,
			&val.IsSelfserve, &val.IsPaymentSettled, &val.PaymentAmount, &val.CreditAmount, &val.CouponAmount,
			&val.UsedPaymentAmount, &val.UsedCreditAmount, &val.UsedCouponAmount, &val.BilledInventory,
			&val.BilledPayment, &val.RequiresApproval, &val.IsRedFlagged, &val.DeliveryType, &val.DeliveryPriority,
			&val.Version, &val.CrId, &val.Data, &val.NthImp, &val.Variance, &val.Mean, &val.DeliverySlabCtr,
			&val.CreatedAt, &val.UpdatedAt, &val.DeletedAt, &val.OnboardedAt, &val.CreatedBy, &val.CurrentCtr,
			&val.ZoneCtr, &val.OptimizationTypeId, &val.ContextTargetingParameters, &val.ThirdPartySource, &val.Tqi,
			&val.IntradayOsi, &val.HasKeywordTargeting, &val.ClicksToDeliver, &val.IsContentCampaign,
			&val.CampaignClassificationId, val.CxPacedReachMaximised, &val.SecondaryGoalBillingRate, &val.HostAppId,
			&val.OptimizationGroupId)
		if err != nil {
			return nil, err
		}
		message = append(message, val)
	}
	return message, nil
}
