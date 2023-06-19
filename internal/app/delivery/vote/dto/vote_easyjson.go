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

func easyjsonE3ecfa40DecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryVoteDto(in *jlexer.Lexer, out *VoteThreadResponseDTO) {
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
		case "id":
			out.ID = int(in.Int())
		case "title":
			out.Title = string(in.String())
		case "author":
			out.Author = string(in.String())
		case "forum":
			out.Forum = string(in.String())
		case "slug":
			out.Slug = string(in.String())
		case "message":
			out.Message = string(in.String())
		case "created":
			out.Created = string(in.String())
		case "votes":
			out.SumVotes = int(in.Int())
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
func easyjsonE3ecfa40EncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryVoteDto(out *jwriter.Writer, in VoteThreadResponseDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"author\":"
		out.RawString(prefix)
		out.String(string(in.Author))
	}
	{
		const prefix string = ",\"forum\":"
		out.RawString(prefix)
		out.String(string(in.Forum))
	}
	{
		const prefix string = ",\"slug\":"
		out.RawString(prefix)
		out.String(string(in.Slug))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"created\":"
		out.RawString(prefix)
		out.String(string(in.Created))
	}
	{
		const prefix string = ",\"votes\":"
		out.RawString(prefix)
		out.Int(int(in.SumVotes))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v VoteThreadResponseDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE3ecfa40EncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryVoteDto(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v VoteThreadResponseDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE3ecfa40EncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryVoteDto(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *VoteThreadResponseDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE3ecfa40DecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryVoteDto(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *VoteThreadResponseDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE3ecfa40DecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryVoteDto(l, v)
}
func easyjsonE3ecfa40DecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryVoteDto1(in *jlexer.Lexer, out *VoteRequestDTO) {
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
		case "ThreadSlugOrID":
			out.ThreadSlugOrID = string(in.String())
		case "nickname":
			out.Nickname = string(in.String())
		case "voice":
			out.Vote = int(in.Int())
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
func easyjsonE3ecfa40EncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryVoteDto1(out *jwriter.Writer, in VoteRequestDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ThreadSlugOrID\":"
		out.RawString(prefix[1:])
		out.String(string(in.ThreadSlugOrID))
	}
	{
		const prefix string = ",\"nickname\":"
		out.RawString(prefix)
		out.String(string(in.Nickname))
	}
	{
		const prefix string = ",\"voice\":"
		out.RawString(prefix)
		out.Int(int(in.Vote))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v VoteRequestDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE3ecfa40EncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryVoteDto1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v VoteRequestDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE3ecfa40EncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryVoteDto1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *VoteRequestDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE3ecfa40DecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryVoteDto1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *VoteRequestDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE3ecfa40DecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryVoteDto1(l, v)
}