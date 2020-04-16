package main

import (
	"context"
	"gitlab.com/otis-team/backend/db/client"
	"gitlab.com/otis-team/backend/dtypes/general-merchant/proto"
	"gitlab.com/otis-team/backend/dtypes/partner-merchant/proto"
	"gitlab.com/otis-team/backend/service/merchant/proto/merchant"
	"log"
)


// Handler struct contains the client connection to the DB, to be used by Handler functions.
type Handler struct {
	Client client.RDSClient
}


// CreateGeneralMerchant handles gRPC requests to create a new merchant in the DB.
func (h *Handler) CreateGeneralMerchant(ctx context.Context, req *merchant.MerchantRequest, res *merchant.MerchantsResponse) error {
	log.Print("CreateMerchant handler fired")
	gMerchant := req.GetGeneralMerchant()
	_, err := h.Client.CreateGeneralMerchant(gMerchant)
	res.Executed = err == nil
	res.GeneralMerchants = []*generalmerchant.GeneralMerchant{ gMerchant }
	return err
}

// CreatePartnerMerchant handles gRPC requests to create a new merchant in the DB.
func (h *Handler) CreatePartnerMerchant(ctx context.Context, req *merchant.MerchantRequest, res *merchant.MerchantsResponse) error {
	log.Print("CreateMerchant handler fired")
	pMerchant := req.GetPartnerMerchant()
	_, err := h.Client.CreatePartnerMerchant(pMerchant)
	res.Executed = err == nil
	res.PartnerMerchants = []*partnermerchant.PartnerMerchant{ pMerchant }
	return err
}

// GetGeneralMerchant handles gRPC requests to retrieve one (if Merchant ID is supplied) or many general merchants from the DB.
func (h *Handler) GetGeneralMerchant(ctx context.Context, req *merchant.MerchantQuery, res *merchant.MerchantsResponse) error {
	log.Print("GetGeneralMerchant handler fired")

	var err error
	var merchants []*generalmerchant.GeneralMerchant
	var merchant *generalmerchant.GeneralMerchant

	if len(req.MerchantID) == 0 {
		merchants, err = h.Client.GetAllGeneralMerchants()
		res.GeneralMerchants = merchants
		res.Executed = err == nil
	} else {
		merchant, err = h.Client.GetGeneralMerchantByID(req.MerchantID)
		res.GeneralMerchants = []*generalmerchant.GeneralMerchant{ merchant }
		res.Executed = err == nil
	}
	return err
}

// GetPartnerMerchant handles gRPC requests to retrieve one (if Merchant ID is supplied) or many partner merchants from the DB.
func (h *Handler) GetPartnerMerchant(ctx context.Context, req *merchant.MerchantQuery, res *merchant.MerchantsResponse) error {
	log.Print("GetPartnerMerchant handler fired!")

	var err error
	var merchants []*partnermerchant.PartnerMerchant
	var merchant *partnermerchant.PartnerMerchant

	if len(req.MerchantID) == 0 {
		merchants, err = h.Client.GetAllPartnerMerchants()
		res.PartnerMerchants = merchants
		res.Executed = err != nil
	} else {
		merchant, err = h.Client.GetPartnerMerchantByID(req.MerchantID)
		res.PartnerMerchants = []*partnermerchant.PartnerMerchant{ merchant }
		res.Executed = err != nil
	}
	return err
}

// DeletePartnerMerchant handles gRPC requests to delete a new partner merchant from the DB
func (h *Handler) DeletePartnerMerchant(ctx context.Context, req *merchant.MerchantQuery, res *merchant.MerchantsResponse) error {
	log.Print("DeletePartnerMerchant handler fired")
	err := h.Client.DeletePartnerMerchant(req.MerchantID)
	res.Executed = err == nil
	return err
}

// DeleteGeneralMerchant handles gRPC requests to delete a new general merchant from the DB
func (h *Handler) DeleteGeneralMerchant(ctx context.Context, req *merchant.MerchantQuery, res *merchant.MerchantsResponse) error {
	log.Print("DeleteGeneralMerchant handler fired")
	err := h.Client.DeleteGeneralMerchant(req.MerchantID)
	res.Executed = err == nil
	return err
}