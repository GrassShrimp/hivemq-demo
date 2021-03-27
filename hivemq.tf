resource "helm_release" "hivemq-operator" {
  name              = "hivemq-operator"
  repository        = "https://hivemq.github.io/helm-charts" 
  chart             = "hivemq-operator"
  version           = "0.8.3"
  namespace         = "hivemq-operator"

  values = [
  <<EOF
  hivemq:
    nodeCount: "3"
    ports:
    - name: "mqtt"
      port: 1883
      expose: true
      patch:
      - '[{"op":"add","path":"/spec/selector/hivemq.com~1node-offline","value":"false"},{"op":"add","path":"/metadata/annotations","value":{"service.spec.externalTrafficPolicy":"Local"}}]'
      # If you want Kubernetes to expose the MQTT port to external traffic
      # - '[{"op":"add","path":"/spec/type","value":"LoadBalancer"}]'
    - name: "cc"
      port: 8080
      expose: true
      patch:
      - '[{"op":"add","path":"/spec/sessionAffinity","value":"ClientIP"}]'
      # If you want Kubernetes to expose the HiveMQ control center via load balancer.
      # Warning: You should consider configuring proper security and TLS beforehand. Ingress may be a better option here.
      # - '[{"op":"add","path":"/spec/type","value":"LoadBalancer"}]'
  EOF
  ]

  create_namespace  = true

  depends_on = [ kind_cluster.hivemq ]
}

resource "null_resource" "install_mysql_route" {
  provisioner "local-exec" {
    command = "kubectl apply -f ./hivemq-route.yaml -n ${helm_release.hivemq-operator.namespace}"
  }

  depends_on = [
    null_resource.installing-istio
  ]
}