import grpc
import test_pb2_grpc
import test_pb2
import grpc

if __name__ == "__main__":
    channel = grpc.insecure_channel('127.0.0.1:1200')
    stub = test_pb2_grpc.SomeStructStub(channel)
    msg = test_pb2.MSG1(A="keyur")
    print stub.SomeFunction(msg)
