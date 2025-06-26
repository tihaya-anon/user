package payload_util

import (
	"MVC_DI/global/infra/schema"
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func protoToNative(msg proto.Message) (map[string]any, error) {
	m := make(map[string]any)
	pb := msg.ProtoReflect()

	pb.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		name := string(fd.Name())
		switch {
		case fd.IsList():
			var arr []any
			list := v.List()
			for i := 0; i < list.Len(); i++ {
				elem := list.Get(i)
				nativeElem, _ := singleValueToNative(fd, elem)
				arr = append(arr, nativeElem)
			}
			m[name] = arr
			return true

		case fd.IsMap():
			nativeMap := make(map[string]any)
			mp := v.Map()
			mp.Range(func(k protoreflect.MapKey, v protoreflect.Value) bool {
				keyStr := k.String()
				valDesc := fd.MapValue()
				nativeVal, _ := singleValueToNative(valDesc, v)
				nativeMap[keyStr] = nativeVal
				return true
			})
			m[name] = nativeMap
			return true
		}

		nativeVal, _ := singleValueToNative(fd, v)
		m[name] = nativeVal
		return true
	})

	return m, nil
}

func singleValueToNative(fd protoreflect.FieldDescriptor, v protoreflect.Value) (any, error) {
	switch fd.Kind() {
	case protoreflect.BoolKind:
		return v.Bool(), nil
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return int32(v.Int()), nil
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return v.Int(), nil
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return int32(v.Uint()), nil
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return int64(v.Uint()), nil
	case protoreflect.FloatKind:
		return float32(v.Float()), nil
	case protoreflect.DoubleKind:
		return v.Float(), nil
	case protoreflect.StringKind:
		return v.String(), nil
	case protoreflect.BytesKind:
		return v.Bytes(), nil
	case protoreflect.EnumKind:
		desc := fd.Enum().Values().ByNumber(v.Enum())
		return string(desc.Name()), nil
	case protoreflect.MessageKind, protoreflect.GroupKind:
		return protoToNative(v.Message().Interface())
	default:
		return nil, fmt.Errorf("unsupported kind %v", fd.Kind())
	}
}

func ProtoToAvroBytes(msg proto.Message) ([]byte, error) {
	codec, err := schema.SchemaManager.GetOrLoadCodecByObject(msg)
	if err != nil {
		return nil, err
	}
	native, err := protoToNative(msg)
	if err != nil {
		return nil, err
	}
	return codec.BinaryFromNative(nil, native)
}
