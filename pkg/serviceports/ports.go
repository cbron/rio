package serviceports

import (
	"fmt"
	"strings"

	riov1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func Protocol(proto riov1.Protocol) (protocol v1.Protocol) {
	switch proto {
	case riov1.ProtocolUDP:
		protocol = v1.ProtocolUDP
	case riov1.ProtocolSCTP:
		protocol = v1.ProtocolSCTP
	default:
		protocol = v1.ProtocolTCP
	}

	return
}

//func ContainerPorts(w riov1.Wrangler) []riov1.ContainerPort {
//	var (
//		ports   []riov1.ContainerPort
//		portMap = map[string]bool{}
//	)
//
//	for _, container := range services.ToNamedContainers(w) {
//		for _, port := range container.Ports {
//			port = stringers.NormalizeContainerPort(port)
//
//			if port.Port == 0 {
//				continue
//			}
//
//			key := fmt.Sprintf("%v/%v", port.Port, port.Protocol)
//			if portMap[key] {
//				continue
//			}
//			portMap[key] = true
//
//			ports = append(ports, port)
//		}
//	}
//
//	return ports
//}

func ServiceNamedPorts(w riov1.Wrangler) (servicePorts []v1.ServicePort) {
	//for _, port := range ContainerPorts(w) {
	for _, port := range w.GetSpec().Ports {
		servicePort := v1.ServicePort{
			Name:     port.Name,
			Port:     port.Port,
			Protocol: Protocol(port.Protocol),
			TargetPort: intstr.IntOrString{
				IntVal: port.TargetPort,
			},
		}

		if servicePort.Name == "" {
			servicePort.Name = strings.ToLower(fmt.Sprintf("%s-%d", port.Protocol, port.Port))
		}

		servicePorts = append(servicePorts, servicePort)
	}

	return
}
