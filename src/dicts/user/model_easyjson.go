// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package user

import (
	customType "2019_2_Shtoby_shto/src/customType"
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

func easyjsonC80ae7adDecode20192ShtobyShtoSrcDictsUser(in *jlexer.Lexer, out *User) {
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
		case "login":
			out.Login = string(in.String())
		case "password":
			out.Password = string(in.String())
		case "photo_id":
			if in.IsNull() {
				in.Skip()
				out.PhotoID = nil
			} else {
				if out.PhotoID == nil {
					out.PhotoID = new(customType.StringUUID)
				}
				*out.PhotoID = customType.StringUUID(in.String())
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
func easyjsonC80ae7adEncode20192ShtobyShtoSrcDictsUser(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"login\":"
		out.RawString(prefix[1:])
		out.String(string(in.Login))
	}
	if in.Password != "" {
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	if in.PhotoID != nil {
		const prefix string = ",\"photo_id\":"
		out.RawString(prefix)
		out.String(string(*in.PhotoID))
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.String(string(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncode20192ShtobyShtoSrcDictsUser(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncode20192ShtobyShtoSrcDictsUser(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecode20192ShtobyShtoSrcDictsUser(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecode20192ShtobyShtoSrcDictsUser(l, v)
}
