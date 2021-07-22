package cmd

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
	"mysql-metadata/internal/repositories"
	"mysql-metadata/internal/storage"
)

func StartApplication() {
	//banners, err := repositories.DbConn.GetBannerByCampaignIdPlacement(1, "pgi")
	//bannersByte := marshalProtoBanner(banners)
	//storage.Upload("active-cacher-bkt", "BannerByCampaignIdPlacement", string(getDataType(
	//	"BannerByCampaignIdPlacement")), bannersByte)
	//events.PushMsgToTopic("OBJECT_FINALIZE", "active-cacher-bkt",
	//	"BannerByCampaignIdPlacement", string(getDataType("BannerByCampaignIdPlacement")))
	//
	//CarBanners, err := repositories.DbConn.GetCarouselBanner(277)
	//CarBannersByte := marshalProtoCarouselBanner(CarBanners)
	//storage.Upload("active-cacher-bkt", "CarouselBanner", string(getDataType("CarouselBanner")),
	//	CarBannersByte)
	//events.PushMsgToTopic("OBJECT_FINALIZE", "active-cacher-bkt", "CarouselBanner",
	//	string(getDataType("CarouselBanner")))
	//
	bannerIds, err := repositories.DbConn.GetBannersByCampaignIds(1)
	bannerIdByte, err := proto.Marshal(&bannerIds)
	storage.Upload("active-cacher-bkt", "BannersByCampaignIds", string(getDataType(
		"BannersByCampaignIds")), bannerIdByte)
	fmt.Println(bannerIds)
	//events.PushMsgToTopic("OBJECT_FINALIZE", "active-cacher-bkt", "BannersByCampaignIds",
	//	string(getDataType("BannersByCampaignIds")))
	//
	//appIds, err := repositories.DbConn.GetAllAppIds()
	//appIdByte, err := proto.Marshal(&appIds)
	//storage.Upload("active-cacher-bkt", "AllAppIds", string(getDataType(
	//	"AllAppIds")), appIdByte)
	//events.PushMsgToTopic("OBJECT_FINALIZE", "active-cacher-bkt", "AllAppIds",
	//	string(getDataType("AllAppIds")))
	//
	//contextMeta, err := repositories.DbConn.GetContextMetaStringData()
	//contextMetaByte, err := proto.Marshal(&contextMeta)
	//storage.Upload("active-cacher-bkt", "ContextMetaStringData", string(getDataType(
	//	"ContextMetaStringData")), contextMetaByte)
	//events.PushMsgToTopic("OBJECT_FINALIZE", "active-cacher-bkt", "ContextMetaStringData",
	//	string(getDataType("ContextMetaStringData")))
	//
	//contextTargeting, err := repositories.DbConn.GetContextTargetingParametersFromAllDeliveringCampaignStringData()
	//contextTargetingByte, err := proto.Marshal(&contextTargeting)
	//storage.Upload("active-cacher-bkt", "ContextTargetingParametersFromAllDeliveringCampaignStringData",
	//	string(getDataType("ContextTargetingParametersFromAllDeliveringCampaignStringData")), contextTargetingByte)
	//events.PushMsgToTopic("OBJECT_FINALIZE", "active-cacher-bkt",
	//	"ContextTargetingParametersFromAllDeliveringCampaignStringData",
	//	string(getDataType("ContextTargetingParametersFromAllDeliveringCampaignStringData")))
	//
	//contentContext, err := repositories.DbConn.GetContentContextMetaStringData("DH_APP")
	//contentContextByte, err := proto.Marshal(&contentContext)
	//storage.Upload("active-cacher-bkt", "ContentContextMetaStringData", string(getDataType(
	//	"ContentContextMetaStringData")), contentContextByte)
	//events.PushMsgToTopic("OBJECT_FINALIZE", "active-cacher-bkt",
	//	"ContentContextMetaStringData", string(getDataType("ContentContextMetaStringData")))
	//
	//affPostBack, err := repositories.DbConn.GetAffiliatePostbackData()
	//affPostBackByte := marshalProtoDhxAffiliate(affPostBack)
	//storage.Upload("active-cacher-bkt", "AffiliatePostbackData", string(getDataType(
	//	"AffiliatePostbackData")), affPostBackByte)
	//events.PushMsgToTopic("OBJECT_FINALIZE", "active-cacher-bkt",
	//	"AffiliatePostbackData", string(getDataType("AffiliatePostbackData")))
	//
	//cardSpacing, err := repositories.DbConn.GetCardSpacingDetailStringData()
	//cardSpacingByte, err := proto.Marshal(&cardSpacing)
	//storage.Upload("active-cacher-bkt", "CardSpacingDetailStringData", string(getDataType(
	//	"CardSpacingDetailStringData")), cardSpacingByte)
	//events.PushMsgToTopic("OBJECT_FINALIZE", "active-cacher-bkt",
	//	"CardSpacingDetailStringData", string(getDataType("CardSpacingDetailStringData")))
	//
	//leadEntityIds, err := repositories.DbConn.GetAllLeadEntityIds()
	//leadEntityIdsByte, err := proto.Marshal(&leadEntityIds)
	//storage.Upload("active-cacher-bkt", "AllLeadEntityIds", string(getDataType("AllLeadEntityIds")),
	//	leadEntityIdsByte)
	//events.PushMsgToTopic("OBJECT_FINALIZE", "active-cacher-bkt", "AllLeadEntityIds",
	//	string(getDataType("AllLeadEntityIds")))
	//
	//affCampaignIds, err := repositories.DbConn.GetAffiliateCampaignIds()
	//affCampaignIdsByte, err := proto.Marshal(&affCampaignIds)
	//storage.Upload("active-cacher-bkt", "AffiliateCampaignIds", string(getDataType(
	//	"AffiliateCampaignIds")), affCampaignIdsByte)
	//events.PushMsgToTopic("OBJECT_FINALIZE", "active-cacher-bkt", "AffiliateCampaignIds",
	//	string(getDataType("AffiliateCampaignIds")))
	//
	//leadEntityInfo, err := repositories.DbConn.GetLeadEntityInfo(1)
	//leadEntityInfoByte, err := proto.Marshal(&leadEntityInfo)
	//storage.Upload("active-cacher-bkt", "LeadEntityInfo", string(getDataType("LeadEntityInfo")),
	//	leadEntityInfoByte)
	//events.PushMsgToTopic("OBJECT_FINALIZE", "active-cacher-bkt", "LeadEntityInfo",
	//	string(getDataType("LeadEntityInfo")))
	//
	//optEntityIds, err := repositories.DbConn.GetAllOptimizationEntityIds()
	//optEntityIdsByte, err := proto.Marshal(&optEntityIds)
	//storage.Upload("active-cacher-bkt", "AllOptimizationEntityIds", string(getDataType(
	//	"AllOptimizationEntityIds")), optEntityIdsByte)
	//events.PushMsgToTopic("OBJECT_FINALIZE", "active-cacher-bkt", "AllOptimizationEntityIds",
	//	string(getDataType("AllOptimizationEntityIds")))
	//
	//optEntityInfo, err := repositories.DbConn.GetOptimizationEntityInfo(1)
	//optEntityInfoByte, err := proto.Marshal(&optEntityInfo)
	//storage.Upload("active-cacher-bkt", "OptimizationEntityInfo", string(getDataType(
	//	"OptimizationEntityInfo")), optEntityInfoByte)
	//events.PushMsgToTopic("OBJECT_FINALIZE", "active-cacher-bkt", "OptimizationEntityInfo",
	//	string(getDataType("OptimizationEntityInfo")))
	//
	//bannersById, err := repositories.DbConn.GetBannerByBannerId(1)
	//bannersByIdByte := marshalProtoBanner(bannersById)
	//storage.Upload("active-cacher-bkt", "BannerByBannerId", string(getDataType("BannerByBannerId")),
	//	bannersByIdByte)
	//events.PushMsgToTopic("OBJECT_FINALIZE", "active-cacher-bkt", "BannerByBannerId",
	//	string(getDataType("BannerByBannerId")))
	//
	//campaign, err := repositories.DbConn.GetCampaignDataByCampaignId(1)
	//campaignByte := marshalProtoCampaignData(campaign)
	//storage.Upload("active-cacher-bkt", "CampaignDataByCampaignId", string(getDataType(
	//	"CampaignDataByCampaignId")), campaignByte)
	//events.PushMsgToTopic("OBJECT_FINALIZE", "active-cacher-bkt", "CampaignDataByCampaignId",
	//	string(getDataType("CampaignDataByCampaignId")))

	if err != nil {
		log.Println("Failed to fetch data ", err.Error())
	}

	repositories.DbConn.SqlDb.Close()
}
