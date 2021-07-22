package repositories

import (
	_ "github.com/go-sql-driver/mysql"
	"mysql-metadata/pkg/proto_models"
)

var (
	messageStr proto_models.StringList
	messageId  proto_models.IdList
)

// Banner
func (db *Db) GetBannerByCampaignIdPlacement(campaignId int, adplacement string) ([]proto_models.Banner, error) {
	res, err := db.SqlDb.Query("SELECT b.id, b.ecpm, b.campaign_id, b.type_id, b.subtype_id, b.name, b.data, "+
		"b.weightage FROM dhx_daedalus_banners AS b INNER JOIN dhx_banner_adplacement_map AS bm ON bm.banner_id = b.id"+
		" WHERE b.campaign_id = ? AND bm.adplacement = ? AND b.status_id = 1  AND b.deleted_at is NULL ORDER By b.id",
		campaignId, adplacement)
	if err != nil {
		return nil, err
	}
	return ScanBanner(res)
}

// Carousel Banner
func (db *Db) GetCarouselBanner(bannerId int) ([]proto_models.CarouselBanner, error) {
	res, err := db.SqlDb.Query("SELECT id, banner_id, name, data, status_id, created_at, updated_at"+
		" FROM dhx_daedalus_carousel_banners WHERE banner_id=? and deleted_at is NULL",
		bannerId)
	if err != nil {
		return nil, err
	}
	return ScanCarouselBanner(res)
}

// IDs
func (db *Db) GetBannersByCampaignIds(campaignId int) (proto_models.IdList, error) {
	res, err := db.SqlDb.Query("SELECT id from dhx_daedalus_banners WHERE campaign_id = ?", campaignId)
	if err != nil {
		return messageId, err
	}
	return ScanIds(res)
}

// IDs
func (db *Db) GetAllAppIds() (proto_models.IdList, error) {
	res, err := db.SqlDb.Query("SELECT id from dhx_apps")
	if err != nil {
		return messageId, err
	}
	return ScanIds(res)
}

// Strings
func (db *Db) GetContextMetaStringData() (proto_models.StringList, error) {
	res, err := db.SqlDb.Query("SELECT data FROM dhx_context_meta WHERE keyname='allContextMeta' LIMIT 1")
	if err != nil {
		return messageStr, err
	}
	return ScanStrings(res)
}

// Strings
func (db *Db) GetContextTargetingParametersFromAllDeliveringCampaignStringData() (proto_models.StringList, error) {
	res, err := db.SqlDb.Query("SELECT context_targeting_parameters FROM dhx_requirements WHERE status_id=5")
	if err != nil {
		return messageStr, err
	}
	return ScanStrings(res)
}

// Strings
func (db *Db) GetContentContextMetaStringData(appId string) (proto_models.StringList, error) {
	res, err := db.SqlDb.Query("SELECT data from dhx_context_meta where keyname = 'entityContextMeta' and host_app_id = ? LIMIT 1", appId)
	if err != nil {
		return messageStr, err
	}
	return ScanStrings(res)
}

// DhxAffiliateOffers
func (db *Db) GetAffiliatePostbackData() ([]proto_models.Affiliate, error) {
	res, err := db.SqlDb.Query("select offers.requirement_id,offers.user_id,offers.event_info from " +
		"dhx_affiliate_offers as offers join dhx_requirements as req on offers.requirement_id = req.id " +
		"where  now() <=  DATE_ADD(end_date, INTERVAL 60 DAY)")
	if err != nil {
		return nil, err
	}
	return ScanDhxAffiliateOffers(res)
}

// Strings
func (db *Db) GetCardSpacingDetailStringData() (proto_models.StringList, error) {
	res, err := db.SqlDb.Query("SELECT rules FROM dhx_ad_placements_settings")
	if err != nil {
		return messageStr, err
	}
	return ScanStrings(res)
}

// IDs
func (db *Db) GetAllLeadEntityIds() (proto_models.IdList, error) {
	res, err := db.SqlDb.Query("select id from dhx_leads")
	if err != nil {
		return messageId, err
	}
	return ScanIds(res)
}

// IDs
func (db *Db) GetAffiliateCampaignIds() (proto_models.IdList, error) {
	res, err := db.SqlDb.Query("select id from dhx_requirements where order_id=-4")
	if err != nil {
		return messageId, err
	}
	return ScanIds(res)
}

// DhxEntityInfo
func (db *Db) GetLeadEntityInfo(leadId int) (proto_models.EntityInfo, error) {
	var message proto_models.EntityInfo
	res, err := db.SqlDb.Query("select id,primary_event_names from dhx_leads where id=?", leadId)
	if err != nil {
		return message, err
	}
	return ScanDhxEntityInfo(res)
}

// IDs
func (db *Db) GetAllOptimizationEntityIds() (proto_models.IdList, error) {
	res, err := db.SqlDb.Query("select id from dhx_optimization_entities")
	if err != nil {
		return messageId, err
	}
	return ScanIds(res)
}

// DhxEntityInfo
func (db *Db) GetOptimizationEntityInfo(id int) (proto_models.EntityInfo, error) {
	var message proto_models.EntityInfo
	res, err := db.SqlDb.Query("select id,primary_event_names from dhx_optimization_entities where id=?", id)
	if err != nil {
		return message, err
	}
	return ScanDhxEntityInfo(res)
}

// Banner
func (db *Db) GetBannerByBannerId(adId int) ([]proto_models.Banner, error) {
	res, err := db.SqlDb.Query("SELECT * from dhx_daedalus_banners WHERE id = ?", adId)
	if err != nil {
		return nil, err
	}
	return ScanBanner(res)
}

// CaimpaignData
func (db *Db) GetCampaignDataByCampaignId(campaignId int) ([]proto_models.Campaign, error) {
	res, err := db.SqlDb.Query("SELECT * FROM dhx_requirements WHERE id = ? limit 1", campaignId)
	if err != nil {
		return nil, err
	}
	return ScanCampaign(res)
}

//func GetContextTargetingParametersFromHubIdsQuery(db sql.DB, hubId int) (*sql.Rows, error) {
//
//	res, err := db.Query("SELECT evaluation_rules FROM dhx_content_contexts WHERE ownership_type_id=2 AND hub_id IN (" . implode(",", hubIds) . ")")
//
//	if err != nil {
//		print(err.Error())
//	}
//
//	return res, err
//}
