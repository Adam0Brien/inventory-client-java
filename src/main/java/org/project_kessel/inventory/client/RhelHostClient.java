package org.project_kessel.inventory.client;

import io.grpc.Channel;
import io.grpc.stub.StreamObserver;
import io.smallrye.mutiny.Uni;
import io.smallrye.mutiny.operators.multi.processors.UnicastProcessor;
import org.project_kessel.api.inventory.v1beta1.CreateRhelHostRequest;
import org.project_kessel.api.inventory.v1beta1.CreateRhelHostResponse;
import org.project_kessel.api.inventory.v1beta1.KesselRhelHostServiceGrpc;



import java.util.logging.Logger;

public class RhelHostClient {
    private static final Logger logger = Logger.getLogger(RhelHostClient.class.getName());

    private final KesselRhelHostServiceGrpc.KesselRhelHostServiceStub asyncStub;
    private final KesselRhelHostServiceGrpc.KesselRhelHostServiceBlockingStub blockingStub;

    RhelHostClient(Channel channel){
        asyncStub = KesselRhelHostServiceGrpc.newStub(channel);
        blockingStub = KesselRhelHostServiceGrpc.newBlockingStub(channel);
    }

    public void CreateRhelHost(CreateRhelHostRequest request, StreamObserver<CreateRhelHostResponse> responseStreamObserver) {
        asyncStub.createRhelHost(request, responseStreamObserver);
    }

    public Uni<CreateRhelHostResponse> CreateRhelHostUni(CreateRhelHostRequest request) {
        final UnicastProcessor<CreateRhelHostResponse> responseProcessor = UnicastProcessor.create();
        var streamObserver = new StreamObserver<CreateRhelHostResponse>() {
            @Override
            public void onNext(CreateRhelHostResponse response) {
                responseProcessor.onNext(response);
            }

            @Override
            public void onError(Throwable t) {
                responseProcessor.onError(t);
            }

            @Override
            public void onCompleted() {
                responseProcessor.onComplete();
            }
        };
        var uni = Uni.createFrom().publisher(responseProcessor);
        CreateRhelHost(request, streamObserver);
        return uni;
    }

    public CreateRhelHostResponse CreateRhelHost(CreateRhelHostRequest request) {
        return blockingStub.createRhelHost(request);
    }
}
