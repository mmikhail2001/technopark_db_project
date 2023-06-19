// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package dto

import (
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

func easyjsonFcc75609DecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryUserDto(in *jlexer.Lexer, out *UserUpdateResponseDTO) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "about":
			out.About = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "fullname":
			out.Fullname = string(in.String())
		case "nickname":
			out.Nickname = string(in.String())
		default:
			in.AddError(&jlexer.LexerError{
				Offset: in.GetPos(),
				Reason: "unknown field",
				Data:   key,
			})
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonFcc75609EncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryUserDto(out *jwriter.Writer, in UserUpdateResponseDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"about\":"
		out.RawString(prefix[1:])
		out.String(string(in.About))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"fullname\":"
		out.RawString(prefix)
		out.String(string(in.Fullname))
	}
	{
		const prefix string = ",\"nickname\":"
		out.RawString(prefix)
		out.String(string(in.Nickname))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserUpdateResponseDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFcc75609EncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryUserDto(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserUpdateResponseDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFcc75609EncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryUserDto(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserUpdateResponseDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFcc75609DecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryUserDto(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserUpdateResponseDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFcc75609DecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryUserDto(l, v)
}
func easyjsonFcc75609DecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryUserDto1(in *jlexer.Lexer, out *UserUpdateRequestDTO) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Nickname_":
			out.Nickname_ = string(in.String())
		case "fullname":
			out.Fullname = string(in.String())
		case "about":
			out.About = string(in.String())
		case "email":
			out.Email = string(in.String())
		default:
			in.AddError(&jlexer.LexerError{
				Offset: in.GetPos(),
				Reason: "unknown field",
				Data:   key,
			})
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonFcc75609EncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryUserDto1(out *jwriter.Writer, in UserUpdateRequestDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Nickname_\":"
		out.RawString(prefix[1:])
		out.String(string(in.Nickname_))
	}
	{
		const prefix string = ",\"fullname\":"
		out.RawString(prefix)
		out.String(string(in.Fullname))
	}
	{
		const prefix string = ",\"about\":"
		out.RawString(prefix)
		out.String(string(in.About))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserUpdateRequestDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFcc75609EncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryUserDto1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserUpdateRequestDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFcc75609EncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryUserDto1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserUpdateRequestDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFcc75609DecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryUserDto1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserUpdateRequestDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFcc75609DecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryUserDto1(l, v)
}