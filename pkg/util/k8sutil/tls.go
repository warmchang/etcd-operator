// Copyright 2017 The etcd-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package k8sutil

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"github.com/on2itsecurity/etcd-operator/pkg/util/etcdutil"
)

type TLSData struct {
	CertData []byte
	KeyData  []byte
	CAData   []byte
}

// GetTLSDataFromSecret retrives the kubernete secret that contain etcd tls certs and put them into TLSData.
func GetTLSDataFromSecret(ctx context.Context, kubecli kubernetes.Interface, ns, se string) (*TLSData, error) {
	secret, err := kubecli.CoreV1().Secrets(ns).Get(ctx, se, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	if secret.Type == v1.SecretTypeTLS {
		return &TLSData{
			CertData: secret.Data["tls.crt"],
			KeyData:  secret.Data["tls.key"],
			CAData:   secret.Data["ca.crt"],
		}, nil
	}

	return &TLSData{
		CertData: secret.Data[etcdutil.CliCertFile],
		KeyData:  secret.Data[etcdutil.CliKeyFile],
		CAData:   secret.Data[etcdutil.CliCAFile],
	}, nil
}
