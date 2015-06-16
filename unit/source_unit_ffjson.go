// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: source_unit.go
// DO NOT EDIT!

package unit

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *ID2) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *ID2) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"Type":`)
	fflib.WriteJsonString(buf, string(mj.Type))
	buf.WriteString(`,"Name":`)
	fflib.WriteJsonString(buf, string(mj.Name))
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_ID2base = iota
	ffj_t_ID2no_such_key

	ffj_t_ID2_Type

	ffj_t_ID2_Name
)

var ffj_key_ID2_Type = []byte("Type")

var ffj_key_ID2_Name = []byte("Name")

func (uj *ID2) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *ID2) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_ID2base
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_ID2no_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'N':

					if bytes.Equal(ffj_key_ID2_Name, kn) {
						currentKey = ffj_t_ID2_Name
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'T':

					if bytes.Equal(ffj_key_ID2_Type, kn) {
						currentKey = ffj_t_ID2_Type
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffj_key_ID2_Name, kn) {
					currentKey = ffj_t_ID2_Name
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_ID2_Type, kn) {
					currentKey = ffj_t_ID2_Type
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_ID2no_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_ID2_Type:
					goto handle_Type

				case ffj_t_ID2_Name:
					goto handle_Name

				case ffj_t_ID2no_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_Type:

	/* handler: uj.Type type=string kind=string */

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			uj.Type = string(fs.Output.String())

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Name:

	/* handler: uj.Name type=string kind=string */

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			uj.Name = string(fs.Output.String())

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:
	return nil
}

func (mj *Key) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Key) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"Repo":`)
	fflib.WriteJsonString(buf, string(mj.Repo))
	buf.WriteString(`,"CommitID":`)
	fflib.WriteJsonString(buf, string(mj.CommitID))
	buf.WriteString(`,"UnitType":`)
	fflib.WriteJsonString(buf, string(mj.UnitType))
	buf.WriteString(`,"Unit":`)
	fflib.WriteJsonString(buf, string(mj.Unit))
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_Keybase = iota
	ffj_t_Keyno_such_key

	ffj_t_Key_Repo

	ffj_t_Key_CommitID

	ffj_t_Key_UnitType

	ffj_t_Key_Unit
)

var ffj_key_Key_Repo = []byte("Repo")

var ffj_key_Key_CommitID = []byte("CommitID")

var ffj_key_Key_UnitType = []byte("UnitType")

var ffj_key_Key_Unit = []byte("Unit")

func (uj *Key) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *Key) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_Keybase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_Keyno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'C':

					if bytes.Equal(ffj_key_Key_CommitID, kn) {
						currentKey = ffj_t_Key_CommitID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'R':

					if bytes.Equal(ffj_key_Key_Repo, kn) {
						currentKey = ffj_t_Key_Repo
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'U':

					if bytes.Equal(ffj_key_Key_UnitType, kn) {
						currentKey = ffj_t_Key_UnitType
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Key_Unit, kn) {
						currentKey = ffj_t_Key_Unit
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffj_key_Key_Unit, kn) {
					currentKey = ffj_t_Key_Unit
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Key_UnitType, kn) {
					currentKey = ffj_t_Key_UnitType
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Key_CommitID, kn) {
					currentKey = ffj_t_Key_CommitID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Key_Repo, kn) {
					currentKey = ffj_t_Key_Repo
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_Keyno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_Key_Repo:
					goto handle_Repo

				case ffj_t_Key_CommitID:
					goto handle_CommitID

				case ffj_t_Key_UnitType:
					goto handle_UnitType

				case ffj_t_Key_Unit:
					goto handle_Unit

				case ffj_t_Keyno_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_Repo:

	/* handler: uj.Repo type=string kind=string */

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			uj.Repo = string(fs.Output.String())

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CommitID:

	/* handler: uj.CommitID type=string kind=string */

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			uj.CommitID = string(fs.Output.String())

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_UnitType:

	/* handler: uj.UnitType type=string kind=string */

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			uj.UnitType = string(fs.Output.String())

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Unit:

	/* handler: uj.Unit type=string kind=string */

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			uj.Unit = string(fs.Output.String())

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:
	return nil
}

func (mj *SourceUnit) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *SourceUnit) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "Name":`)
	fflib.WriteJsonString(buf, string(mj.Name))
	buf.WriteString(`,"Type":`)
	fflib.WriteJsonString(buf, string(mj.Type))
	buf.WriteByte(',')
	if len(mj.Repo) != 0 {
		buf.WriteString(`"Repo":`)
		fflib.WriteJsonString(buf, string(mj.Repo))
		buf.WriteByte(',')
	}
	if len(mj.CommitID) != 0 {
		buf.WriteString(`"CommitID":`)
		fflib.WriteJsonString(buf, string(mj.CommitID))
		buf.WriteByte(',')
	}
	if len(mj.Globs) != 0 {
		buf.WriteString(`"Globs":`)
		if mj.Globs != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Globs {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.WriteJsonString(buf, string(v))
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	buf.WriteString(`"Files":`)
	if mj.Files != nil {
		buf.WriteString(`[`)
		for i, v := range mj.Files {
			if i != 0 {
				buf.WriteString(`,`)
			}
			fflib.WriteJsonString(buf, string(v))
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteByte(',')
	if len(mj.Dir) != 0 {
		buf.WriteString(`"Dir":`)
		fflib.WriteJsonString(buf, string(mj.Dir))
		buf.WriteByte(',')
	}
	if len(mj.Dependencies) != 0 {
		buf.WriteString(`"Dependencies":`)
		if mj.Dependencies != nil {
			buf.WriteString(`[`)
			for i, v := range mj.Dependencies {
				if i != 0 {
					buf.WriteString(`,`)
				}
				/* Interface types must use runtime reflection. type=interface {} kind=interface */
				err = buf.Encode(v)
				if err != nil {
					return err
				}
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if mj.Info != nil {
		if true {
			/* Struct fall back. type=unit.Info kind=struct */
			buf.WriteString(`"Info":`)
			err = buf.Encode(mj.Info)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	if mj.Data != nil {
		buf.WriteString(`"Data":`)
		/* Interface types must use runtime reflection. type=interface {} kind=interface */
		err = buf.Encode(mj.Data)
		if err != nil {
			return err
		}
		buf.WriteByte(',')
	}
	if len(mj.Config) != 0 {
		buf.WriteString(`"Config":`)
		/* Falling back. type=map[string]interface {} kind=map */
		err = buf.Encode(mj.Config)
		if err != nil {
			return err
		}
		buf.WriteByte(',')
	}
	if len(mj.Ops) != 0 {
		buf.WriteString(`"Ops":`)
		/* Falling back. type=map[string]*srclib.ToolRef kind=map */
		err = buf.Encode(mj.Ops)
		if err != nil {
			return err
		}
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_SourceUnitbase = iota
	ffj_t_SourceUnitno_such_key

	ffj_t_SourceUnit_Name

	ffj_t_SourceUnit_Type

	ffj_t_SourceUnit_Repo

	ffj_t_SourceUnit_CommitID

	ffj_t_SourceUnit_Globs

	ffj_t_SourceUnit_Files

	ffj_t_SourceUnit_Dir

	ffj_t_SourceUnit_Dependencies

	ffj_t_SourceUnit_Info

	ffj_t_SourceUnit_Data

	ffj_t_SourceUnit_Config

	ffj_t_SourceUnit_Ops
)

var ffj_key_SourceUnit_Name = []byte("Name")

var ffj_key_SourceUnit_Type = []byte("Type")

var ffj_key_SourceUnit_Repo = []byte("Repo")

var ffj_key_SourceUnit_CommitID = []byte("CommitID")

var ffj_key_SourceUnit_Globs = []byte("Globs")

var ffj_key_SourceUnit_Files = []byte("Files")

var ffj_key_SourceUnit_Dir = []byte("Dir")

var ffj_key_SourceUnit_Dependencies = []byte("Dependencies")

var ffj_key_SourceUnit_Info = []byte("Info")

var ffj_key_SourceUnit_Data = []byte("Data")

var ffj_key_SourceUnit_Config = []byte("Config")

var ffj_key_SourceUnit_Ops = []byte("Ops")

func (uj *SourceUnit) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *SourceUnit) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_SourceUnitbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_SourceUnitno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'C':

					if bytes.Equal(ffj_key_SourceUnit_CommitID, kn) {
						currentKey = ffj_t_SourceUnit_CommitID
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_SourceUnit_Config, kn) {
						currentKey = ffj_t_SourceUnit_Config
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'D':

					if bytes.Equal(ffj_key_SourceUnit_Dir, kn) {
						currentKey = ffj_t_SourceUnit_Dir
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_SourceUnit_Dependencies, kn) {
						currentKey = ffj_t_SourceUnit_Dependencies
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_SourceUnit_Data, kn) {
						currentKey = ffj_t_SourceUnit_Data
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'F':

					if bytes.Equal(ffj_key_SourceUnit_Files, kn) {
						currentKey = ffj_t_SourceUnit_Files
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'G':

					if bytes.Equal(ffj_key_SourceUnit_Globs, kn) {
						currentKey = ffj_t_SourceUnit_Globs
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'I':

					if bytes.Equal(ffj_key_SourceUnit_Info, kn) {
						currentKey = ffj_t_SourceUnit_Info
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'N':

					if bytes.Equal(ffj_key_SourceUnit_Name, kn) {
						currentKey = ffj_t_SourceUnit_Name
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'O':

					if bytes.Equal(ffj_key_SourceUnit_Ops, kn) {
						currentKey = ffj_t_SourceUnit_Ops
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'R':

					if bytes.Equal(ffj_key_SourceUnit_Repo, kn) {
						currentKey = ffj_t_SourceUnit_Repo
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'T':

					if bytes.Equal(ffj_key_SourceUnit_Type, kn) {
						currentKey = ffj_t_SourceUnit_Type
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffj_key_SourceUnit_Ops, kn) {
					currentKey = ffj_t_SourceUnit_Ops
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_SourceUnit_Config, kn) {
					currentKey = ffj_t_SourceUnit_Config
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_SourceUnit_Data, kn) {
					currentKey = ffj_t_SourceUnit_Data
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_SourceUnit_Info, kn) {
					currentKey = ffj_t_SourceUnit_Info
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_SourceUnit_Dependencies, kn) {
					currentKey = ffj_t_SourceUnit_Dependencies
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_SourceUnit_Dir, kn) {
					currentKey = ffj_t_SourceUnit_Dir
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_SourceUnit_Files, kn) {
					currentKey = ffj_t_SourceUnit_Files
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_SourceUnit_Globs, kn) {
					currentKey = ffj_t_SourceUnit_Globs
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_SourceUnit_CommitID, kn) {
					currentKey = ffj_t_SourceUnit_CommitID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_SourceUnit_Repo, kn) {
					currentKey = ffj_t_SourceUnit_Repo
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_SourceUnit_Type, kn) {
					currentKey = ffj_t_SourceUnit_Type
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_SourceUnit_Name, kn) {
					currentKey = ffj_t_SourceUnit_Name
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_SourceUnitno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_SourceUnit_Name:
					goto handle_Name

				case ffj_t_SourceUnit_Type:
					goto handle_Type

				case ffj_t_SourceUnit_Repo:
					goto handle_Repo

				case ffj_t_SourceUnit_CommitID:
					goto handle_CommitID

				case ffj_t_SourceUnit_Globs:
					goto handle_Globs

				case ffj_t_SourceUnit_Files:
					goto handle_Files

				case ffj_t_SourceUnit_Dir:
					goto handle_Dir

				case ffj_t_SourceUnit_Dependencies:
					goto handle_Dependencies

				case ffj_t_SourceUnit_Info:
					goto handle_Info

				case ffj_t_SourceUnit_Data:
					goto handle_Data

				case ffj_t_SourceUnit_Config:
					goto handle_Config

				case ffj_t_SourceUnit_Ops:
					goto handle_Ops

				case ffj_t_SourceUnitno_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_Name:

	/* handler: uj.Name type=string kind=string */

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			uj.Name = string(fs.Output.String())

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Type:

	/* handler: uj.Type type=string kind=string */

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			uj.Type = string(fs.Output.String())

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Repo:

	/* handler: uj.Repo type=string kind=string */

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			uj.Repo = string(fs.Output.String())

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CommitID:

	/* handler: uj.CommitID type=string kind=string */

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			uj.CommitID = string(fs.Output.String())

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Globs:

	/* handler: uj.Globs type=[]string kind=slice */

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Globs = nil
		} else {

			uj.Globs = make([]string, 0)

			wantVal := true

			for {

				var v string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: v type=string kind=string */

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						v = string(fs.Output.String())

					}
				}

				uj.Globs = append(uj.Globs, v)
				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Files:

	/* handler: uj.Files type=[]string kind=slice */

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Files = nil
		} else {

			uj.Files = make([]string, 0)

			wantVal := true

			for {

				var v string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: v type=string kind=string */

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						v = string(fs.Output.String())

					}
				}

				uj.Files = append(uj.Files, v)
				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Dir:

	/* handler: uj.Dir type=string kind=string */

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			uj.Dir = string(fs.Output.String())

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Dependencies:

	/* handler: uj.Dependencies type=[]interface {} kind=slice */

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Dependencies = nil
		} else {

			uj.Dependencies = make([]interface{}, 0)

			wantVal := true

			for {

				var v interface{}

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: v type=interface {} kind=interface */

				{
					/* Falling back. type=interface {} kind=interface */
					tbuf, err := fs.CaptureField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}

					err = json.Unmarshal(tbuf, &v)
					if err != nil {
						return fs.WrapErr(err)
					}
				}

				uj.Dependencies = append(uj.Dependencies, v)
				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Info:

	/* handler: uj.Info type=unit.Info kind=struct */

	{
		/* Falling back. type=unit.Info kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.Info)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Data:

	/* handler: uj.Data type=interface {} kind=interface */

	{
		/* Falling back. type=interface {} kind=interface */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.Data)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Config:

	/* handler: uj.Config type=map[string]interface {} kind=map */

	{
		/* Falling back. type=map[string]interface {} kind=map */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.Config)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Ops:

	/* handler: uj.Ops type=map[string]*srclib.ToolRef kind=map */

	{
		/* Falling back. type=map[string]*srclib.ToolRef kind=map */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.Ops)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:
	return nil
}
