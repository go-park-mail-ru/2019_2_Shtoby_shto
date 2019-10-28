// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package cardGroup

import (
	customType "2019_2_Shtoby_shto/src/customType"
	card "2019_2_Shtoby_shto/src/dicts/card"
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC80ae7adDecode20192ShtobyShtoSrcDictsCardGroup(in *jlexer.Lexer, out *CardGroup) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "board_id":
			out.BoardID = customType.StringUUID(in.String())
		case "cards":
			if in.IsNull() {
				in.Skip()
				out.Cards = nil
			} else {
				in.Delim('[')
				if out.Cards == nil {
					if !in.IsDelim(']') {
						out.Cards = make([]card.Card, 0, 1)
					} else {
						out.Cards = []card.Card{}
					}
				} else {
					out.Cards = (out.Cards)[:0]
				}
				for !in.IsDelim(']') {
					var v1 card.Card
					(v1).UnmarshalEasyJSON(in)
					out.Cards = append(out.Cards, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "id":
			out.ID = customType.StringUUID(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC80ae7adEncode20192ShtobyShtoSrcDictsCardGroup(out *jwriter.Writer, in CardGroup) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"board_id\":"
		out.RawString(prefix)
		out.String(string(in.BoardID))
	}
	{
		const prefix string = ",\"cards\":"
		out.RawString(prefix)
		if in.Cards == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Cards {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.String(string(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CardGroup) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncode20192ShtobyShtoSrcDictsCardGroup(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CardGroup) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncode20192ShtobyShtoSrcDictsCardGroup(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CardGroup) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecode20192ShtobyShtoSrcDictsCardGroup(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CardGroup) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecode20192ShtobyShtoSrcDictsCardGroup(l, v)
}