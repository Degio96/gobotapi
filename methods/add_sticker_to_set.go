// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"github.com/Squirrel-Network/gobotapi/types"
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
)

// AddStickerToSet Use this method to add a new sticker to a set created by the bot
// You must use exactly one of the fields png_sticker, tgs_sticker, or webm_sticker
// Animated stickers can be added to animated sticker sets and only to them
// Animated sticker sets can have up to 50 stickers
// Static sticker sets can have up to 120 stickers
// Returns True on success.
type AddStickerToSet struct {
	Emojis string `json:"emojis"`
	MaskPosition *types.MaskPosition `json:"mask_position,omitempty"`
	Name string `json:"name"`
	PngSticker rawTypes.InputFile `json:"png_sticker,omitempty"`
	TgsSticker rawTypes.InputFile `json:"tgs_sticker,omitempty"`
	UserID int64 `json:"user_id"`
	WebmSticker rawTypes.InputFile `json:"webm_sticker,omitempty"`
}

func (entity *AddStickerToSet) Files() map[string]rawTypes.InputFile {
	files := make(map[string]rawTypes.InputFile)
	switch entity.PngSticker.(type) {
		case types.InputFile:
			files["png_sticker"] = entity.PngSticker
			entity.PngSticker = nil
	}
	switch entity.TgsSticker.(type) {
		case types.InputFile:
			files["tgs_sticker"] = entity.TgsSticker
			entity.TgsSticker = nil
	}
	switch entity.WebmSticker.(type) {
		case types.InputFile:
			files["webm_sticker"] = entity.WebmSticker
			entity.WebmSticker = nil
	}
	return files
}

func (AddStickerToSet) MethodName() string {
	return "addStickerToSet"
}

func (AddStickerToSet) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result bool `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result {
		Kind: types.TypeBoolean,
		Result: x1.Result,
	}
	return &result, nil
}
