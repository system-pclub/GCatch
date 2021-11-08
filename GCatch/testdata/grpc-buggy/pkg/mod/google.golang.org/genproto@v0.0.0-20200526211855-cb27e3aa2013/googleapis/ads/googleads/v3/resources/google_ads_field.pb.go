// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.2
// source: google/ads/googleads/v3/resources/google_ads_field.proto

package resources

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v3/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// A field or resource (artifact) used by GoogleAdsService.
type GoogleAdsField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the artifact.
	// Artifact resource names have the form:
	//
	// `googleAdsFields/{name}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The name of the artifact.
	Name *wrappers.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The category of the artifact.
	Category enums.GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory `protobuf:"varint,3,opt,name=category,proto3,enum=google.ads.googleads.v3.enums.GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory" json:"category,omitempty"`
	// Output only. Whether the artifact can be used in a SELECT clause in search
	// queries.
	Selectable *wrappers.BoolValue `protobuf:"bytes,4,opt,name=selectable,proto3" json:"selectable,omitempty"`
	// Output only. Whether the artifact can be used in a WHERE clause in search
	// queries.
	Filterable *wrappers.BoolValue `protobuf:"bytes,5,opt,name=filterable,proto3" json:"filterable,omitempty"`
	// Output only. Whether the artifact can be used in a ORDER BY clause in search
	// queries.
	Sortable *wrappers.BoolValue `protobuf:"bytes,6,opt,name=sortable,proto3" json:"sortable,omitempty"`
	// Output only. The names of all resources, segments, and metrics that are selectable with
	// the described artifact.
	SelectableWith []*wrappers.StringValue `protobuf:"bytes,7,rep,name=selectable_with,json=selectableWith,proto3" json:"selectable_with,omitempty"`
	// Output only. The names of all resources that are selectable with the described
	// artifact. Fields from these resources do not segment metrics when included
	// in search queries.
	//
	// This field is only set for artifacts whose category is RESOURCE.
	AttributeResources []*wrappers.StringValue `protobuf:"bytes,8,rep,name=attribute_resources,json=attributeResources,proto3" json:"attribute_resources,omitempty"`
	// Output only. At and beyond version V1 this field lists the names of all metrics that are
	// selectable with the described artifact when it is used in the FROM clause.
	// It is only set for artifacts whose category is RESOURCE.
	//
	// Before version V1 this field lists the names of all metrics that are
	// selectable with the described artifact. It is only set for artifacts whose
	// category is either RESOURCE or SEGMENT
	Metrics []*wrappers.StringValue `protobuf:"bytes,9,rep,name=metrics,proto3" json:"metrics,omitempty"`
	// Output only. At and beyond version V1 this field lists the names of all artifacts,
	// whether a segment or another resource, that segment metrics when included
	// in search queries and when the described artifact is used in the FROM
	// clause. It is only set for artifacts whose category is RESOURCE.
	//
	// Before version V1 this field lists the names of all artifacts, whether a
	// segment or another resource, that segment metrics when included in search
	// queries. It is only set for artifacts of category RESOURCE, SEGMENT or
	// METRIC.
	Segments []*wrappers.StringValue `protobuf:"bytes,10,rep,name=segments,proto3" json:"segments,omitempty"`
	// Output only. Values the artifact can assume if it is a field of type ENUM.
	//
	// This field is only set for artifacts of category SEGMENT or ATTRIBUTE.
	EnumValues []*wrappers.StringValue `protobuf:"bytes,11,rep,name=enum_values,json=enumValues,proto3" json:"enum_values,omitempty"`
	// Output only. This field determines the operators that can be used with the artifact
	// in WHERE clauses.
	DataType enums.GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType `protobuf:"varint,12,opt,name=data_type,json=dataType,proto3,enum=google.ads.googleads.v3.enums.GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType" json:"data_type,omitempty"`
	// Output only. The URL of proto describing the artifact's data type.
	TypeUrl *wrappers.StringValue `protobuf:"bytes,13,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// Output only. Whether the field artifact is repeated.
	IsRepeated *wrappers.BoolValue `protobuf:"bytes,14,opt,name=is_repeated,json=isRepeated,proto3" json:"is_repeated,omitempty"`
}

func (x *GoogleAdsField) Reset() {
	*x = GoogleAdsField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v3_resources_google_ads_field_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleAdsField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleAdsField) ProtoMessage() {}

func (x *GoogleAdsField) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v3_resources_google_ads_field_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleAdsField.ProtoReflect.Descriptor instead.
func (*GoogleAdsField) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v3_resources_google_ads_field_proto_rawDescGZIP(), []int{0}
}

func (x *GoogleAdsField) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *GoogleAdsField) GetName() *wrappers.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *GoogleAdsField) GetCategory() enums.GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory {
	if x != nil {
		return x.Category
	}
	return enums.GoogleAdsFieldCategoryEnum_UNSPECIFIED
}

func (x *GoogleAdsField) GetSelectable() *wrappers.BoolValue {
	if x != nil {
		return x.Selectable
	}
	return nil
}

func (x *GoogleAdsField) GetFilterable() *wrappers.BoolValue {
	if x != nil {
		return x.Filterable
	}
	return nil
}

func (x *GoogleAdsField) GetSortable() *wrappers.BoolValue {
	if x != nil {
		return x.Sortable
	}
	return nil
}

func (x *GoogleAdsField) GetSelectableWith() []*wrappers.StringValue {
	if x != nil {
		return x.SelectableWith
	}
	return nil
}

func (x *GoogleAdsField) GetAttributeResources() []*wrappers.StringValue {
	if x != nil {
		return x.AttributeResources
	}
	return nil
}

func (x *GoogleAdsField) GetMetrics() []*wrappers.StringValue {
	if x != nil {
		return x.Metrics
	}
	return nil
}

func (x *GoogleAdsField) GetSegments() []*wrappers.StringValue {
	if x != nil {
		return x.Segments
	}
	return nil
}

func (x *GoogleAdsField) GetEnumValues() []*wrappers.StringValue {
	if x != nil {
		return x.EnumValues
	}
	return nil
}

func (x *GoogleAdsField) GetDataType() enums.GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType {
	if x != nil {
		return x.DataType
	}
	return enums.GoogleAdsFieldDataTypeEnum_UNSPECIFIED
}

func (x *GoogleAdsField) GetTypeUrl() *wrappers.StringValue {
	if x != nil {
		return x.TypeUrl
	}
	return nil
}

func (x *GoogleAdsField) GetIsRepeated() *wrappers.BoolValue {
	if x != nil {
		return x.IsRepeated
	}
	return nil
}

var File_google_ads_googleads_v3_resources_google_ads_field_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v3_resources_google_ads_field_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x61, 0x64, 0x73, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x33, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x3d, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5f, 0x61, 0x64, 0x73, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5f, 0x61, 0x64, 0x73, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x08, 0x0a, 0x0e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x54, 0x0a, 0x0d, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x2f, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x29, 0x0a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x35, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x71, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x50, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x33, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x45,
	0x6e, 0x75, 0x6d, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x3f, 0x0a, 0x0a, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x73,
	0x6f, 0x72, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x08,
	0x73, 0x6f, 0x72, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x4a, 0x0a, 0x0f, 0x73, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x0e, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x57, 0x69, 0x74, 0x68, 0x12, 0x52, 0x0a, 0x13, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x12, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x3b, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x07, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x3d, 0x0a, 0x08, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x08, 0x73, 0x65, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x42, 0x0a, 0x0b, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x65, 0x6e,
	0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x72, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x50, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79,
	0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3c, 0x0a, 0x08,
	0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x07, 0x74, 0x79, 0x70, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x40, 0x0a, 0x0b, 0x69, 0x73,
	0x5f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x0a, 0x69, 0x73, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x3a, 0x50, 0xea, 0x41,
	0x4d, 0x0a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2f, 0x7b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5f, 0x61, 0x64, 0x73, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x7d, 0x42, 0x80,
	0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x13, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x4a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41,
	0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x33, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41,
	0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x33, 0x5c,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x56, 0x33, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v3_resources_google_ads_field_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v3_resources_google_ads_field_proto_rawDescData = file_google_ads_googleads_v3_resources_google_ads_field_proto_rawDesc
)

func file_google_ads_googleads_v3_resources_google_ads_field_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v3_resources_google_ads_field_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v3_resources_google_ads_field_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v3_resources_google_ads_field_proto_rawDescData)
	})
	return file_google_ads_googleads_v3_resources_google_ads_field_proto_rawDescData
}

var file_google_ads_googleads_v3_resources_google_ads_field_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v3_resources_google_ads_field_proto_goTypes = []interface{}{
	(*GoogleAdsField)(nil),                                       // 0: google.ads.googleads.v3.resources.GoogleAdsField
	(*wrappers.StringValue)(nil),                                 // 1: google.protobuf.StringValue
	(enums.GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory)(0), // 2: google.ads.googleads.v3.enums.GoogleAdsFieldCategoryEnum.GoogleAdsFieldCategory
	(*wrappers.BoolValue)(nil),                                   // 3: google.protobuf.BoolValue
	(enums.GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType)(0), // 4: google.ads.googleads.v3.enums.GoogleAdsFieldDataTypeEnum.GoogleAdsFieldDataType
}
var file_google_ads_googleads_v3_resources_google_ads_field_proto_depIdxs = []int32{
	1,  // 0: google.ads.googleads.v3.resources.GoogleAdsField.name:type_name -> google.protobuf.StringValue
	2,  // 1: google.ads.googleads.v3.resources.GoogleAdsField.category:type_name -> google.ads.googleads.v3.enums.GoogleAdsFieldCategoryEnum.GoogleAdsFieldCategory
	3,  // 2: google.ads.googleads.v3.resources.GoogleAdsField.selectable:type_name -> google.protobuf.BoolValue
	3,  // 3: google.ads.googleads.v3.resources.GoogleAdsField.filterable:type_name -> google.protobuf.BoolValue
	3,  // 4: google.ads.googleads.v3.resources.GoogleAdsField.sortable:type_name -> google.protobuf.BoolValue
	1,  // 5: google.ads.googleads.v3.resources.GoogleAdsField.selectable_with:type_name -> google.protobuf.StringValue
	1,  // 6: google.ads.googleads.v3.resources.GoogleAdsField.attribute_resources:type_name -> google.protobuf.StringValue
	1,  // 7: google.ads.googleads.v3.resources.GoogleAdsField.metrics:type_name -> google.protobuf.StringValue
	1,  // 8: google.ads.googleads.v3.resources.GoogleAdsField.segments:type_name -> google.protobuf.StringValue
	1,  // 9: google.ads.googleads.v3.resources.GoogleAdsField.enum_values:type_name -> google.protobuf.StringValue
	4,  // 10: google.ads.googleads.v3.resources.GoogleAdsField.data_type:type_name -> google.ads.googleads.v3.enums.GoogleAdsFieldDataTypeEnum.GoogleAdsFieldDataType
	1,  // 11: google.ads.googleads.v3.resources.GoogleAdsField.type_url:type_name -> google.protobuf.StringValue
	3,  // 12: google.ads.googleads.v3.resources.GoogleAdsField.is_repeated:type_name -> google.protobuf.BoolValue
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v3_resources_google_ads_field_proto_init() }
func file_google_ads_googleads_v3_resources_google_ads_field_proto_init() {
	if File_google_ads_googleads_v3_resources_google_ads_field_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v3_resources_google_ads_field_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleAdsField); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v3_resources_google_ads_field_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v3_resources_google_ads_field_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v3_resources_google_ads_field_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v3_resources_google_ads_field_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v3_resources_google_ads_field_proto = out.File
	file_google_ads_googleads_v3_resources_google_ads_field_proto_rawDesc = nil
	file_google_ads_googleads_v3_resources_google_ads_field_proto_goTypes = nil
	file_google_ads_googleads_v3_resources_google_ads_field_proto_depIdxs = nil
}
