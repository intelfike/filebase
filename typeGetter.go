package filebase

// null & nil?
import "encoding/json"

//  fmt.Stringer interface
func (f Filebase) String() string {
	var b []byte
	if len(f.Indent) == 0 {
		b, _ = f.Bytes()
	} else {
		b, _ = f.BytesIndent()
	}
	return string(b)
}

func (f *Filebase) Bytes() ([]byte, error) {
	return json.Marshal(f.Interface())
}
func (f *Filebase) BytesIndent() ([]byte, error) {
	return json.MarshalIndent(f.Interface(), "", f.Indent)
}

// Assert string.
func (f *Filebase) ToString() string {
	return f.Interface().(string)
}

func (f *Filebase) ToBytes() []byte {
	return []byte(f.Interface().(string))
}

func (f *Filebase) ToBool() bool {
	return f.Interface().(bool)
}

func (f *Filebase) ToInt() int64 {
	return f.Interface().(int64)
}

func (f *Filebase) ToUint() uint64 {
	return f.Interface().(uint64)
}

func (f *Filebase) ToFloat() float64 {
	return f.Interface().(float64)
}

func (f *Filebase) ToArray() []*Filebase {
	a := make([]*Filebase, f.Len())
	fc, _ := f.Clone()
	for n, _ := range fc.Interface().([]interface{}) {
		a[n] = fc.Child(n)
	}
	return a
}

func (f *Filebase) ToMap() map[string]*Filebase {
	m := map[string]*Filebase{}
	fc, _ := f.Clone()
	for k, _ := range fc.Interface().(map[string]interface{}) {
		m[k] = fc.Child(k)
	}
	return m
}

func (f *Filebase) Interface() interface{} {
	i, err := f.GetInterfacePt()
	if err != nil {
		return nil
	}
	return *i
}
