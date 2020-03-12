import grpc

import addressbook_pb2
import addressbook_pb2_grpc


def run():
    # 注意(gRPC Python Team): .close()方法在channel上是可用的。
    # 并且应该在with语句不符合代码需求的情况下使用。
    with grpc.insecure_channel('localhost:8972') as channel:
        stub = addressbook_pb2_grpc.GreeterStub(channel)
        response = stub.SayHello(addressbook_pb2.HelloRequest(name='wei'))
    print("Greeter client received: {}!".format(response.message))


if __name__ == '__main__':
    run()