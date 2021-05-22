kubectl delete namespace golang2021-meetup
kubectl create namespace golang2021-meetup

kubectl -n golang2021-meetup create secret generic grpc-server-certificates \
  --from-file=tls.crt=./serverCertificates/grpc-server.crt \
  --from-file=tls.key=./serverCertificates/grpc-server.key \
  --from-file=ca.crt=./certificatesChain/grpc-root-ca-and-grpc-server-ca-chain.crt
