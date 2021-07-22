package cmd

import (
	"bytes"
	"google.golang.org/protobuf/proto"
	"log"
	"mysql-metadata/pkg/proto_models"
)

type ProtoType string

const (
	Banner         = "Banner"
	Campaign       = "Campaign"
	CarouselBanner = "CarouselBanner"
	Affiliate      = "Affiliate"
	StringList     = "StringList"
	IdList         = "IdList"
	EntityInfo     = "EntityInfo"
)

func getDataType(key string) ProtoType {
	switch key {
	case "BannerByCampaignIdPlacement":
		return Banner
	case "CarouselBanner":
		return CarouselBanner
	case "BannersByCampaignIds":
		return IdList
	case "AllAppIds":
		return IdList
	case "ContextMetaStringData":
		return StringList
	case "ContextTargetingParametersFromAllDeliveringCampaignStringData":
		return StringList
	case "ContentContextMetaStringData":
		return StringList
	case "AffiliatePostbackData":
		return Affiliate
	case "CardSpacingDetailStringData":
		return StringList
	case "AllLeadEntityIds":
		return IdList
	case "AffiliateCampaignIds":
		return IdList
	case "LeadEntityInfo":
		return EntityInfo
	case "AllOptimizationEntityIds":
		return IdList
	case "OptimizationEntityInfo":
		return EntityInfo
	case "BannerByBannerId":
		return Banner
	case "CampaignDataByCampaignId":
		return Campaign
	}
	return ""
}

func marshalProtoBanner(protoArr []proto_models.Banner) []byte {
	var byteArr [][]byte
	for i := range protoArr {
		byte, err := proto.Marshal(&protoArr[i])
		if err != nil {
			log.Println("unable to marshal ", err.Error())
		}
		byteArr = append(byteArr, byte)
	}
	result := bytes.Join(byteArr, nil)
	return result
}

func marshalProtoCarouselBanner(protoArr []proto_models.CarouselBanner) []byte {
	var byteArr [][]byte
	for i := range protoArr {
		byte, err := proto.Marshal(&protoArr[i])
		if err != nil {
			log.Println("unable to marshal ", err.Error())
		}
		byteArr = append(byteArr, byte)
	}
	result := bytes.Join(byteArr, nil)
	return result
}

func marshalProtoDhxAffiliate(protoArr []proto_models.Affiliate) []byte {
	var byteArr [][]byte
	for i := range protoArr {
		byte, err := proto.Marshal(&protoArr[i])
		if err != nil {
			log.Println("unable to marshal ", err.Error())
		}
		byteArr = append(byteArr, byte)
	}
	result := bytes.Join(byteArr, nil)
	return result
}

func marshalProtoCampaignData(protoArr []proto_models.Campaign) []byte {
	var byteArr [][]byte
	for i := range protoArr {
		byte, err := proto.Marshal(&protoArr[i])
		if err != nil {
			log.Println("unable to marshal ", err.Error())
		}
		byteArr = append(byteArr, byte)
	}
	result := bytes.Join(byteArr, nil)
	return result
}
