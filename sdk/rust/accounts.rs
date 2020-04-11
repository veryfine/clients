// This file is generated by rust-protobuf 2.13.0. Do not edit
// @generated

// https://github.com/rust-lang/rust-clippy/issues/702
#![allow(unknown_lints)]
#![allow(clippy::all)]

#![cfg_attr(rustfmt, rustfmt_skip)]

#![allow(box_pointers)]
#![allow(dead_code)]
#![allow(missing_docs)]
#![allow(non_camel_case_types)]
#![allow(non_snake_case)]
#![allow(non_upper_case_globals)]
#![allow(trivial_casts)]
#![allow(unsafe_code)]
#![allow(unused_imports)]
#![allow(unused_results)]
//! Generated file from `auth/accounts.proto`

use protobuf::Message as Message_imported_for_functions;
use protobuf::ProtobufEnum as ProtobufEnum_imported_for_functions;

/// Generated files are compatible only with the same version
/// of protobuf runtime.
// const _PROTOBUF_VERSION_CHECK: () = ::protobuf::VERSION_2_13_0;

#[derive(PartialEq,Clone,Default)]
pub struct ListAccountsRequest {
    // special fields
    pub unknown_fields: ::protobuf::UnknownFields,
    pub cached_size: ::protobuf::CachedSize,
}

impl<'a> ::std::default::Default for &'a ListAccountsRequest {
    fn default() -> &'a ListAccountsRequest {
        <ListAccountsRequest as ::protobuf::Message>::default_instance()
    }
}

impl ListAccountsRequest {
    pub fn new() -> ListAccountsRequest {
        ::std::default::Default::default()
    }
}

impl ::protobuf::Message for ListAccountsRequest {
    fn is_initialized(&self) -> bool {
        true
    }

    fn merge_from(&mut self, is: &mut ::protobuf::CodedInputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        while !is.eof()? {
            let (field_number, wire_type) = is.read_tag_unpack()?;
            match field_number {
                _ => {
                    ::protobuf::rt::read_unknown_or_skip_group(field_number, wire_type, is, self.mut_unknown_fields())?;
                },
            };
        }
        ::std::result::Result::Ok(())
    }

    // Compute sizes of nested messages
    #[allow(unused_variables)]
    fn compute_size(&self) -> u32 {
        let mut my_size = 0;
        my_size += ::protobuf::rt::unknown_fields_size(self.get_unknown_fields());
        self.cached_size.set(my_size);
        my_size
    }

    fn write_to_with_cached_sizes(&self, os: &mut ::protobuf::CodedOutputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        os.write_unknown_fields(self.get_unknown_fields())?;
        ::std::result::Result::Ok(())
    }

    fn get_cached_size(&self) -> u32 {
        self.cached_size.get()
    }

    fn get_unknown_fields(&self) -> &::protobuf::UnknownFields {
        &self.unknown_fields
    }

    fn mut_unknown_fields(&mut self) -> &mut ::protobuf::UnknownFields {
        &mut self.unknown_fields
    }

    fn as_any(&self) -> &dyn (::std::any::Any) {
        self as &dyn (::std::any::Any)
    }
    fn as_any_mut(&mut self) -> &mut dyn (::std::any::Any) {
        self as &mut dyn (::std::any::Any)
    }
    fn into_any(self: Box<Self>) -> ::std::boxed::Box<dyn (::std::any::Any)> {
        self
    }

    fn descriptor(&self) -> &'static ::protobuf::reflect::MessageDescriptor {
        Self::descriptor_static()
    }

    fn new() -> ListAccountsRequest {
        ListAccountsRequest::new()
    }

    fn descriptor_static() -> &'static ::protobuf::reflect::MessageDescriptor {
        static mut descriptor: ::protobuf::lazy::Lazy<::protobuf::reflect::MessageDescriptor> = ::protobuf::lazy::Lazy::INIT;
        unsafe {
            descriptor.get(|| {
                let fields = ::std::vec::Vec::new();
                ::protobuf::reflect::MessageDescriptor::new_pb_name::<ListAccountsRequest>(
                    "ListAccountsRequest",
                    fields,
                    file_descriptor_proto()
                )
            })
        }
    }

    fn default_instance() -> &'static ListAccountsRequest {
        static mut instance: ::protobuf::lazy::Lazy<ListAccountsRequest> = ::protobuf::lazy::Lazy::INIT;
        unsafe {
            instance.get(ListAccountsRequest::new)
        }
    }
}

impl ::protobuf::Clear for ListAccountsRequest {
    fn clear(&mut self) {
        self.unknown_fields.clear();
    }
}

impl ::std::fmt::Debug for ListAccountsRequest {
    fn fmt(&self, f: &mut ::std::fmt::Formatter<'_>) -> ::std::fmt::Result {
        ::protobuf::text_format::fmt(self, f)
    }
}

impl ::protobuf::reflect::ProtobufValue for ListAccountsRequest {
    fn as_ref(&self) -> ::protobuf::reflect::ReflectValueRef {
        ::protobuf::reflect::ReflectValueRef::Message(self)
    }
}

#[derive(PartialEq,Clone,Default)]
pub struct ListAccountsResponse {
    // message fields
    pub accounts: ::protobuf::RepeatedField<super::auth::Account>,
    // special fields
    pub unknown_fields: ::protobuf::UnknownFields,
    pub cached_size: ::protobuf::CachedSize,
}

impl<'a> ::std::default::Default for &'a ListAccountsResponse {
    fn default() -> &'a ListAccountsResponse {
        <ListAccountsResponse as ::protobuf::Message>::default_instance()
    }
}

impl ListAccountsResponse {
    pub fn new() -> ListAccountsResponse {
        ::std::default::Default::default()
    }

    // repeated .go.micro.auth.Account accounts = 1;


    pub fn get_accounts(&self) -> &[super::auth::Account] {
        &self.accounts
    }
    pub fn clear_accounts(&mut self) {
        self.accounts.clear();
    }

    // Param is passed by value, moved
    pub fn set_accounts(&mut self, v: ::protobuf::RepeatedField<super::auth::Account>) {
        self.accounts = v;
    }

    // Mutable pointer to the field.
    pub fn mut_accounts(&mut self) -> &mut ::protobuf::RepeatedField<super::auth::Account> {
        &mut self.accounts
    }

    // Take field
    pub fn take_accounts(&mut self) -> ::protobuf::RepeatedField<super::auth::Account> {
        ::std::mem::replace(&mut self.accounts, ::protobuf::RepeatedField::new())
    }
}

impl ::protobuf::Message for ListAccountsResponse {
    fn is_initialized(&self) -> bool {
        for v in &self.accounts {
            if !v.is_initialized() {
                return false;
            }
        };
        true
    }

    fn merge_from(&mut self, is: &mut ::protobuf::CodedInputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        while !is.eof()? {
            let (field_number, wire_type) = is.read_tag_unpack()?;
            match field_number {
                1 => {
                    ::protobuf::rt::read_repeated_message_into(wire_type, is, &mut self.accounts)?;
                },
                _ => {
                    ::protobuf::rt::read_unknown_or_skip_group(field_number, wire_type, is, self.mut_unknown_fields())?;
                },
            };
        }
        ::std::result::Result::Ok(())
    }

    // Compute sizes of nested messages
    #[allow(unused_variables)]
    fn compute_size(&self) -> u32 {
        let mut my_size = 0;
        for value in &self.accounts {
            let len = value.compute_size();
            my_size += 1 + ::protobuf::rt::compute_raw_varint32_size(len) + len;
        };
        my_size += ::protobuf::rt::unknown_fields_size(self.get_unknown_fields());
        self.cached_size.set(my_size);
        my_size
    }

    fn write_to_with_cached_sizes(&self, os: &mut ::protobuf::CodedOutputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        for v in &self.accounts {
            os.write_tag(1, ::protobuf::wire_format::WireTypeLengthDelimited)?;
            os.write_raw_varint32(v.get_cached_size())?;
            v.write_to_with_cached_sizes(os)?;
        };
        os.write_unknown_fields(self.get_unknown_fields())?;
        ::std::result::Result::Ok(())
    }

    fn get_cached_size(&self) -> u32 {
        self.cached_size.get()
    }

    fn get_unknown_fields(&self) -> &::protobuf::UnknownFields {
        &self.unknown_fields
    }

    fn mut_unknown_fields(&mut self) -> &mut ::protobuf::UnknownFields {
        &mut self.unknown_fields
    }

    fn as_any(&self) -> &dyn (::std::any::Any) {
        self as &dyn (::std::any::Any)
    }
    fn as_any_mut(&mut self) -> &mut dyn (::std::any::Any) {
        self as &mut dyn (::std::any::Any)
    }
    fn into_any(self: Box<Self>) -> ::std::boxed::Box<dyn (::std::any::Any)> {
        self
    }

    fn descriptor(&self) -> &'static ::protobuf::reflect::MessageDescriptor {
        Self::descriptor_static()
    }

    fn new() -> ListAccountsResponse {
        ListAccountsResponse::new()
    }

    fn descriptor_static() -> &'static ::protobuf::reflect::MessageDescriptor {
        static mut descriptor: ::protobuf::lazy::Lazy<::protobuf::reflect::MessageDescriptor> = ::protobuf::lazy::Lazy::INIT;
        unsafe {
            descriptor.get(|| {
                let mut fields = ::std::vec::Vec::new();
                fields.push(::protobuf::reflect::accessor::make_repeated_field_accessor::<_, ::protobuf::types::ProtobufTypeMessage<super::auth::Account>>(
                    "accounts",
                    |m: &ListAccountsResponse| { &m.accounts },
                    |m: &mut ListAccountsResponse| { &mut m.accounts },
                ));
                ::protobuf::reflect::MessageDescriptor::new_pb_name::<ListAccountsResponse>(
                    "ListAccountsResponse",
                    fields,
                    file_descriptor_proto()
                )
            })
        }
    }

    fn default_instance() -> &'static ListAccountsResponse {
        static mut instance: ::protobuf::lazy::Lazy<ListAccountsResponse> = ::protobuf::lazy::Lazy::INIT;
        unsafe {
            instance.get(ListAccountsResponse::new)
        }
    }
}

impl ::protobuf::Clear for ListAccountsResponse {
    fn clear(&mut self) {
        self.accounts.clear();
        self.unknown_fields.clear();
    }
}

impl ::std::fmt::Debug for ListAccountsResponse {
    fn fmt(&self, f: &mut ::std::fmt::Formatter<'_>) -> ::std::fmt::Result {
        ::protobuf::text_format::fmt(self, f)
    }
}

impl ::protobuf::reflect::ProtobufValue for ListAccountsResponse {
    fn as_ref(&self) -> ::protobuf::reflect::ReflectValueRef {
        ::protobuf::reflect::ReflectValueRef::Message(self)
    }
}

static file_descriptor_proto_data: &'static [u8] = b"\
    \n\x13auth/accounts.proto\x12\rgo.micro.auth\x1a\nauth.proto\"\x15\n\x13\
    ListAccountsRequest\"J\n\x14ListAccountsResponse\x122\n\x08accounts\x18\
    \x01\x20\x03(\x0b2\x16.go.micro.auth.AccountR\x08accounts2]\n\x08Account\
    s\x12Q\n\x04List\x12\".go.micro.auth.ListAccountsRequest\x1a#.go.micro.a\
    uth.ListAccountsResponse\"\0B\x0bZ\tauth;authJ\x80\x02\n\x06\x12\x04\0\0\
    \x11\x01\n\x08\n\x01\x0c\x12\x03\0\0\x12\n\x08\n\x01\x02\x12\x03\x02\x08\
    \x15\n\x08\n\x01\x08\x12\x03\x04\0\x20\n\t\n\x02\x08\x0b\x12\x03\x04\0\
    \x20\n\t\n\x02\x03\0\x12\x03\x06\x07\x13\n\n\n\x02\x06\0\x12\x04\x08\0\n\
    \x01\n\n\n\x03\x06\0\x01\x12\x03\x08\x08\x10\n\x0b\n\x04\x06\0\x02\0\x12\
    \x03\t\x02A\n\x0c\n\x05\x06\0\x02\0\x01\x12\x03\t\x06\n\n\x0c\n\x05\x06\
    \0\x02\0\x02\x12\x03\t\x0b\x1e\n\x0c\n\x05\x06\0\x02\0\x03\x12\x03\t)=\n\
    \n\n\x02\x04\0\x12\x04\x0c\0\r\x01\n\n\n\x03\x04\0\x01\x12\x03\x0c\x08\
    \x1b\n\n\n\x02\x04\x01\x12\x04\x0f\0\x11\x01\n\n\n\x03\x04\x01\x01\x12\
    \x03\x0f\x08\x1c\n\x0b\n\x04\x04\x01\x02\0\x12\x03\x10\x02\x20\n\x0c\n\
    \x05\x04\x01\x02\0\x04\x12\x03\x10\x02\n\n\x0c\n\x05\x04\x01\x02\0\x06\
    \x12\x03\x10\x0b\x12\n\x0c\n\x05\x04\x01\x02\0\x01\x12\x03\x10\x13\x1b\n\
    \x0c\n\x05\x04\x01\x02\0\x03\x12\x03\x10\x1e\x1fb\x06proto3\
";

static mut file_descriptor_proto_lazy: ::protobuf::lazy::Lazy<::protobuf::descriptor::FileDescriptorProto> = ::protobuf::lazy::Lazy::INIT;

fn parse_descriptor_proto() -> ::protobuf::descriptor::FileDescriptorProto {
    ::protobuf::parse_from_bytes(file_descriptor_proto_data).unwrap()
}

pub fn file_descriptor_proto() -> &'static ::protobuf::descriptor::FileDescriptorProto {
    unsafe {
        file_descriptor_proto_lazy.get(|| {
            parse_descriptor_proto()
        })
    }
}