// GENERATED CODE -- DO NOT EDIT!

// package: nhentai
// file: nhentai.proto

import * as nhentai_pb from "./nhentai_pb";
import * as grpc from "grpc";

interface IAPIService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  gallery: grpc.MethodDefinition<nhentai_pb.GalleryRequest, nhentai_pb.GalleryResponse>;
}

export const APIService: IAPIService;

export class APIClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  gallery(argument: nhentai_pb.GalleryRequest, callback: grpc.requestCallback<nhentai_pb.GalleryResponse>): grpc.ClientUnaryCall;
  gallery(argument: nhentai_pb.GalleryRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<nhentai_pb.GalleryResponse>): grpc.ClientUnaryCall;
  gallery(argument: nhentai_pb.GalleryRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<nhentai_pb.GalleryResponse>): grpc.ClientUnaryCall;
}
