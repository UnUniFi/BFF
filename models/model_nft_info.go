package models

type NftInfo struct {

	// nft_class_id/nft_id
	Id NftId `json:"id"`

	NftClass CosmosNftClass `json:"nft_class,omitempty"`

	Nft CosmosNftNft `json:"nft,omitempty"`

	Meta NftMeta `json:"meta,omitempty"`

	Listing Listing `json:"listing,omitempty"`
}
