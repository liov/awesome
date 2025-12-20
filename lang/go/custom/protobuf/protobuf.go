package protobuf

import (
	"fmt"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"time"
	"unicode"

	stringsx "github.com/hopeio/gox/strings"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
)

var (
	mutex        sync.RWMutex
	pkgCache     = map[string]*descriptorpb.FileDescriptorProto{}
	messageCache = map[reflect.Type]*descriptorpb.DescriptorProto{}
	fileCache    = map[reflect.Type]protoreflect.FieldDescriptor{}
	anon         = 0
)

// StructToDescriptor untest
func StructToDescriptor(v any) (protoreflect.MessageDescriptor, error) {
	if pb, ok := v.(proto.Message); ok {
		return pb.ProtoReflect().Descriptor(), nil
	}

	t := reflect.TypeOf(v)
	if t == nil {
		return nil, fmt.Errorf("nil value")
	}
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("not struct: %s", t.Kind())
	}
	return buildMessage(t), nil
}

func buildMessage(t reflect.Type, field *reflect.StructField, parent *descriptorpb.DescriptorProto) *descriptorpb.DescriptorProto {
	pkgPath := t.PkgPath()
	if pkgPath == "" {
		pkgPath = "dynamic"
	}
	file, ok := pkgCache[pkgPath]
	if !ok {
		file = &descriptorpb.FileDescriptorProto{
			Name:        proto.String(filepath.Base(pkgPath) + ".proto"),
			Package:     proto.String(strings.ReplaceAll(pkgPath, "/", ".")),
			MessageType: []*descriptorpb.DescriptorProto{},
			Syntax:      proto.String("proto3"),
		}
		pkgCache[pkgPath] = file
	}

	if d, ok := messageCache[t]; ok {
		return d
	}
	d := &descriptorpb.DescriptorProto{
		Name: proto.String(t.Name()),
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		d.Field = append(d.Field, fdp)
	}
	return d
}

func buildField(t reflect.Type, f *reflect.StructField, parent *descriptorpb.DescriptorProto) *descriptorpb.FieldDescriptorProto {
	if f.PkgPath != "" {
		return nil
	}
	fn := deriveFieldName(f)

	lab := descriptorpb.FieldDescriptorProto_LABEL_REQUIRED
	typ := descriptorpb.FieldDescriptorProto_TYPE_STRING
	var typeName string
	ft := f.Type
	for ft.Kind() == reflect.Ptr {
		ft = ft.Elem()
	}
	switch ft.Kind() {
	case reflect.Bool:
		typ = descriptorpb.FieldDescriptorProto_TYPE_BOOL
	case reflect.Int8, reflect.Int16, reflect.Int32:
		typ = descriptorpb.FieldDescriptorProto_TYPE_INT32
	case reflect.Int, reflect.Int64:
		typ = descriptorpb.FieldDescriptorProto_TYPE_INT64
	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		typ = descriptorpb.FieldDescriptorProto_TYPE_UINT32
	case reflect.Uint, reflect.Uint64:
		typ = descriptorpb.FieldDescriptorProto_TYPE_UINT64
	case reflect.Float32:
		typ = descriptorpb.FieldDescriptorProto_TYPE_FLOAT
	case reflect.Float64:
		typ = descriptorpb.FieldDescriptorProto_TYPE_DOUBLE
	case reflect.String:
		typ = descriptorpb.FieldDescriptorProto_TYPE_STRING
	case reflect.Slice:
		if ft.Elem().Kind() == reflect.Uint8 {
			typ = descriptorpb.FieldDescriptorProto_TYPE_BYTES
		} else {
			lab = descriptorpb.FieldDescriptorProto_LABEL_REPEATED
			et := ft.Elem()
			for et.Kind() == reflect.Ptr {
				et = et.Elem()
			}
			sub := buildField(et, f, parent)
			if sub.GetType() == descriptorpb.FieldDescriptorProto_TYPE_MESSAGE {
				typeName = sub.GetTypeName()
			}
		}
	case reflect.Map:
		lab = descriptorpb.FieldDescriptorProto_LABEL_REPEATED
		entry := buildMapEntry(*parent.Name, fn, ft, c)
		typ = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE
		typeName = entry.GetName()
	case reflect.Struct:
		typ = descriptorpb.FieldDescriptorProto_TYPE_MESSAGE
		if ft == reflect.TypeOf(time.Time{}) {
			typeName = ".google.protobuf.Timestamp"
		} else {
			child := buildMessage(ft, f, parent)
			typeName = child.GetName()
		}
	case reflect.Ptr:
		lab = descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL
		if ft.Elem().Kind() == reflect.Uint8 {
			typ = descriptorpb.FieldDescriptorProto_TYPE_BYTES
		} else {
			lab = descriptorpb.FieldDescriptorProto_LABEL_REPEATED
			et := ft.Elem()
			for et.Kind() == reflect.Ptr {
				et = et.Elem()
			}
			sub := buildField(et, f, parent)
			if sub.GetType() == descriptorpb.FieldDescriptorProto_TYPE_MESSAGE {
				typeName = sub.GetTypeName()
			}
		}
	default:
		typ = descriptorpb.FieldDescriptorProto_TYPE_STRING
	}
	fdp := &descriptorpb.FieldDescriptorProto{
		Name:   proto.String(fn),
		Number: proto.Int32(int32(len(d.Field) + 1)),
		Label:  &lab,
		Type:   &typ,
	}
	if typeName != "" {
		fdp.TypeName = proto.String(typeName)
	}
	return fdp
}

func buildMapEntry(rootName, fieldName string, ft reflect.Type, c *ctx) *descriptorpb.DescriptorProto {
	en := rootName + "_" + stringsx.CamelCase(fieldName) + "Entry"
	if d, ok := c.byName[en]; ok {
		return d
	}
	d := &descriptorpb.DescriptorProto{}
	d.Name = proto.String(en)
	b := true
	d.Options = &descriptorpb.MessageOptions{MapEntry: &b}
	ktyp := ft.Key()
	vtyp := ft.Elem()
	for vtyp.Kind() == reflect.Ptr {
		vtyp = vtyp.Elem()
	}
	kLab := descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL
	kNum := int32(1)
	kFd := &descriptorpb.FieldDescriptorProto{}
	kFd.Name = proto.String("key")
	kFd.Number = proto.Int32(kNum)
	kFd.Label = &kLab
	kFd.Type = enumForScalar(ktyp)
	vLab := descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL
	vNum := int32(2)
	vFd := &descriptorpb.FieldDescriptorProto{}
	vFd.Name = proto.String("value")
	vFd.Number = proto.Int32(vNum)
	vFd.Label = &vLab

	d.Field = []*descriptorpb.FieldDescriptorProto{kFd, vFd}
	c.nested = append(c.nested, d)
	c.byName[en] = d
	return d
}

func enumForScalar(t reflect.Type) *descriptorpb.FieldDescriptorProto_Type {
	switch t.Kind() {
	case reflect.Bool:
		tt := descriptorpb.FieldDescriptorProto_TYPE_BOOL
		return &tt
	case reflect.Int8, reflect.Int16, reflect.Int32:
		tt := descriptorpb.FieldDescriptorProto_TYPE_INT32
		return &tt
	case reflect.Int, reflect.Int64:
		tt := descriptorpb.FieldDescriptorProto_TYPE_INT64
		return &tt
	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		tt := descriptorpb.FieldDescriptorProto_TYPE_UINT32
		return &tt
	case reflect.Uint, reflect.Uint64:
		tt := descriptorpb.FieldDescriptorProto_TYPE_UINT64
		return &tt
	case reflect.Float32:
		tt := descriptorpb.FieldDescriptorProto_TYPE_FLOAT
		return &tt
	case reflect.Float64:
		tt := descriptorpb.FieldDescriptorProto_TYPE_DOUBLE
		return &tt
	case reflect.String:
		tt := descriptorpb.FieldDescriptorProto_TYPE_STRING
		return &tt
	default:
		tt := descriptorpb.FieldDescriptorProto_TYPE_STRING
		return &tt
	}
}

func typeName(t reflect.Type) string {
	if t.Name() != "" {
		return t.Name()
	}
	return "Struct"
}

func typeName(t reflect.Type) string {
	if t.Name() != "" {
		p := t.PkgPath()
		if p == "" {
			return t.Name()
		}
		return strings.NewReplacer("/", "_", ".", "_", "-", "_").Replace(p + "_" + t.Name())
	}
	anon++
	return "AnonStruct" + fmt.Sprintf("%d", anon)
}

func deriveFieldName(f *reflect.StructField) string {
	n := f.Name
	tag := f.Tag.Get("json")
	if tag != "" {
		p := strings.Split(tag, ",")[0]
		if p != "" && p != "-" {
			n = p
		}
	}
	return n
}
