HOST_SYSTEM = $(shell uname | cut -f 1 -d_)
SYSTEM ?= $(HOST_SYSTEM)
CXX = g++
CPPFLAGS += `pkg-config --cflags protobuf grpc`
CXXFLAGS += -std=c++11 -I .
ifeq ($(SYSTEM),Darwin)
LDFLAGS += -L/usr/local/lib `pkg-config --libs protobuf grpc++ grpc`\
           -lgrpc++_reflection\
           -ldl
else
LDFLAGS += -L/usr/local/lib `pkg-config --libs protobuf grpc++ grpc`\
           -Wl,--no-as-needed -lgrpc++_reflection -Wl,--as-needed\
           -ldl
endif
PROTOC = protoc
GRPC_CPP_PLUGIN = grpc_cpp_plugin
GRPC_CPP_PLUGIN_PATH ?= `which $(GRPC_CPP_PLUGIN)`

PROTOS_PATH = ./proto/

vpath %.proto $(PROTOS_PATH)

all: greeter-server greeter-client

greeter-server: helloworld.pb.o helloworld.grpc.pb.o server/server.o
	$(CXX) $^ $(LDFLAGS) -o $@

greeter-client: helloworld.pb.o helloworld.grpc.pb.o client/client.o
	$(CXX) $^ $(LDFLAGS) -o $@

.PRECIOUS: %.grpc.pb.cc
%.grpc.pb.cc: %.proto
	$(PROTOC) -I $(PROTOS_PATH) --grpc_out=. --plugin=protoc-gen-grpc=$(GRPC_CPP_PLUGIN_PATH) $<

.PRECIOUS: %.pb.cc
%.pb.cc: %.proto
	$(PROTOC) -I $(PROTOS_PATH) --cpp_out=. $<

clean:
	rm -f *.o client/*.o server/*.o *.pb *.pb.cc *.pb.h

compile-service:
	g++ -std=c++14 summer.cpp summerMain.cpp -lgtest -lgtest_main -pthread -o sumProgram

test-service: 
	./sumProgram

# compile orchestration
comp-orch:
	g++ -std=c++14 main.cpp train.cpp setup.cpp create_loaders.cpp model_checkpoint.cpp -o delphi-inferential-service

# execute orchestration 
exec-orch:
	./delphi-inferential-service