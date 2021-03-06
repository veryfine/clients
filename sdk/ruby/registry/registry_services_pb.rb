# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: registry/registry.proto for package 'go.micro.registry'

require 'grpc'
require 'registry/registry_pb'

module Go
  module Micro
    module Registry
      module Registry
        class Service

          include GRPC::GenericService

          self.marshal_class_method = :encode
          self.unmarshal_class_method = :decode
          self.service_name = 'go.micro.registry.Registry'

          rpc :GetService, GetRequest, GetResponse
          rpc :Register, Service, EmptyResponse
          rpc :Deregister, Service, EmptyResponse
          rpc :ListServices, ListRequest, ListResponse
          rpc :Watch, WatchRequest, stream(Result)
        end

        Stub = Service.rpc_stub_class
      end
    end
  end
end
