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

func easyjson5bb86ef4DecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryPostDto(in *jlexer.Lexer, out *PostsResponseDTO) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(PostsResponseDTO, 0, 8)
			} else {
				*out = PostsResponseDTO{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 *PostResponseDTO
			if in.IsNull() {
				in.Skip()
				v1 = nil
			} else {
				if v1 == nil {
					v1 = new(PostResponseDTO)
				}
				(*v1).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson5bb86ef4EncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryPostDto(out *jwriter.Writer, in PostsResponseDTO) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			if v3 == nil {
				out.RawString("null")
			} else {
				(*v3).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v PostsResponseDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5bb86ef4EncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryPostDto(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostsResponseDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5bb86ef4EncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryPostDto(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostsResponseDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5bb86ef4DecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryPostDto(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostsResponseDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5bb86ef4DecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryPostDto(l, v)
}
func easyjson5bb86ef4DecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryPostDto1(in *jlexer.Lexer, out *PostsRequestDTO) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(PostsRequestDTO, 0, 1)
			} else {
				*out = PostsRequestDTO{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v4 PostRequestDTO
			(v4).UnmarshalEasyJSON(in)
			*out = append(*out, v4)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson5bb86ef4EncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryPostDto1(out *jwriter.Writer, in PostsRequestDTO) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v5, v6 := range in {
			if v5 > 0 {
				out.RawByte(',')
			}
			(v6).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v PostsRequestDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5bb86ef4EncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryPostDto1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostsRequestDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5bb86ef4EncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryPostDto1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostsRequestDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5bb86ef4DecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryPostDto1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostsRequestDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5bb86ef4DecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryPostDto1(l, v)
}
func easyjson5bb86ef4DecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryPostDto2(in *jlexer.Lexer, out *PostResponseDTO) {
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
		case "parent":
			out.Parent = int(in.Int())
		case "author":
			out.Author = string(in.String())
		case "message":
			out.Message = string(in.String())
		case "isEdited":
			out.IsEdited = bool(in.Bool())
		case "forum":
			out.Forum = string(in.String())
		case "thread":
			out.Thread = int(in.Int())
		case "created":
			out.Created = string(in.String())
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
func easyjson5bb86ef4EncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryPostDto2(out *jwriter.Writer, in PostResponseDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"parent\":"
		out.RawString(prefix)
		out.Int(int(in.Parent))
	}
	{
		const prefix string = ",\"author\":"
		out.RawString(prefix)
		out.String(string(in.Author))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"isEdited\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsEdited))
	}
	{
		const prefix string = ",\"forum\":"
		out.RawString(prefix)
		out.String(string(in.Forum))
	}
	{
		const prefix string = ",\"thread\":"
		out.RawString(prefix)
		out.Int(int(in.Thread))
	}
	{
		const prefix string = ",\"created\":"
		out.RawString(prefix)
		out.String(string(in.Created))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostResponseDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5bb86ef4EncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryPostDto2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostResponseDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5bb86ef4EncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryPostDto2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostResponseDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5bb86ef4DecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryPostDto2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostResponseDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5bb86ef4DecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryPostDto2(l, v)
}
func easyjson5bb86ef4DecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryPostDto3(in *jlexer.Lexer, out *PostRequestDTO) {
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
		case "parent":
			out.Parent = int(in.Int())
		case "author":
			out.Author = string(in.String())
		case "message":
			out.Message = string(in.String())
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
func easyjson5bb86ef4EncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryPostDto3(out *jwriter.Writer, in PostRequestDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"parent\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Parent))
	}
	{
		const prefix string = ",\"author\":"
		out.RawString(prefix)
		out.String(string(in.Author))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostRequestDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5bb86ef4EncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryPostDto3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostRequestDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5bb86ef4EncodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryPostDto3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostRequestDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5bb86ef4DecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryPostDto3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostRequestDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5bb86ef4DecodeGithubComMmikhail2001TechnoparkDbProjectInternalAppDeliveryPostDto3(l, v)
}