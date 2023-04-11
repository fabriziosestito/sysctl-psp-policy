// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package main

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

func easyjson6601e8cdDecodeTmpEasyjson(in *jlexer.Lexer, out *RawSettings) {
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
		case "allowedUnsafeSysctls":
			if in.IsNull() {
				in.Skip()
				out.AllowedUnsafeSysctls = nil
			} else {
				in.Delim('[')
				if out.AllowedUnsafeSysctls == nil {
					if !in.IsDelim(']') {
						out.AllowedUnsafeSysctls = make([]string, 0, 4)
					} else {
						out.AllowedUnsafeSysctls = []string{}
					}
				} else {
					out.AllowedUnsafeSysctls = (out.AllowedUnsafeSysctls)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.AllowedUnsafeSysctls = append(out.AllowedUnsafeSysctls, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "forbiddenSysctls":
			if in.IsNull() {
				in.Skip()
				out.ForbiddenSysctls = nil
			} else {
				in.Delim('[')
				if out.ForbiddenSysctls == nil {
					if !in.IsDelim(']') {
						out.ForbiddenSysctls = make([]string, 0, 4)
					} else {
						out.ForbiddenSysctls = []string{}
					}
				} else {
					out.ForbiddenSysctls = (out.ForbiddenSysctls)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.ForbiddenSysctls = append(out.ForbiddenSysctls, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson6601e8cdEncodeTmpEasyjson(out *jwriter.Writer, in RawSettings) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"allowedUnsafeSysctls\":"
		out.RawString(prefix[1:])
		if in.AllowedUnsafeSysctls == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.AllowedUnsafeSysctls {
				if v3 > 0 {
					out.RawByte(',')
				}
				out.String(string(v4))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"forbiddenSysctls\":"
		out.RawString(prefix)
		if in.ForbiddenSysctls == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.ForbiddenSysctls {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RawSettings) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeTmpEasyjson(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RawSettings) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeTmpEasyjson(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RawSettings) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeTmpEasyjson(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RawSettings) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeTmpEasyjson(l, v)
}