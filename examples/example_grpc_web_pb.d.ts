import * as grpcWeb from 'grpc-web';

import * as example_pb from './example_pb'; // proto import: "example.proto"


export class AuthClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  signIn(
    request: example_pb.SignInRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: example_pb.SignInResponse) => void
  ): grpcWeb.ClientReadableStream<example_pb.SignInResponse>;

  signUp(
    request: example_pb.SignUpRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: example_pb.SignUpResponse) => void
  ): grpcWeb.ClientReadableStream<example_pb.SignUpResponse>;

}

export class AuthPromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  signIn(
    request: example_pb.SignInRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<example_pb.SignInResponse>;

  signUp(
    request: example_pb.SignUpRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<example_pb.SignUpResponse>;

}

