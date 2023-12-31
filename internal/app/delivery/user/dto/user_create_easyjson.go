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

func easyjsonCe93b82aDecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryUserDto(in *jlexer.Lexer, out *UsersDTO) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(UsersDTO, 0, 1)
			} else {
				*out = UsersDTO{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 UserDTO
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCe93b82aEncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryUserDto(out *jwriter.Writer, in UsersDTO) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v UsersDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCe93b82aEncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryUserDto(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UsersDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCe93b82aEncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryUserDto(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UsersDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCe93b82aDecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryUserDto(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UsersDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCe93b82aDecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryUserDto(l, v)
}
func easyjsonCe93b82aDecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryUserDto1(in *jlexer.Lexer, out *UserDTO) {
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
func easyjsonCe93b82aEncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryUserDto1(out *jwriter.Writer, in UserDTO) {
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
func (v UserDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCe93b82aEncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryUserDto1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCe93b82aEncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryUserDto1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCe93b82aDecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryUserDto1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCe93b82aDecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryUserDto1(l, v)
}
func easyjsonCe93b82aDecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryUserDto2(in *jlexer.Lexer, out *UserCreateRequestDTO) {
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
		case "Nickname":
			out.Nickname = string(in.String())
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
func easyjsonCe93b82aEncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryUserDto2(out *jwriter.Writer, in UserCreateRequestDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Nickname\":"
		out.RawString(prefix[1:])
		out.String(string(in.Nickname))
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
func (v UserCreateRequestDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCe93b82aEncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryUserDto2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserCreateRequestDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCe93b82aEncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryUserDto2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserCreateRequestDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCe93b82aDecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryUserDto2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserCreateRequestDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCe93b82aDecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryUserDto2(l, v)
}
