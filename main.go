package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/sha3"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

const (
	EncryptionProviderConfigShake256SecretName      = "encryption-provider-config-shake256"
	EncryptionProviderConfigShake256SecretNamespace = "kube-system"

	EnvNodeName = "NODE_NAME"
)

var (
	scheme = runtime.NewScheme()
)

func init() {
	_ = clientgoscheme.AddToScheme(scheme)

	//+kubebuilder:scaffold:scheme
}

func main() {
	ctx := context.TODO()
	// flags
	var encryptionConfigFilePath string
	flag.StringVar(&encryptionConfigFilePath, "encryption-config-file-path", "/etc/kubernetes/encryption/config.yaml", "The path to the encryption config file.")

	// read the encryption config file
	f, err := ioutil.ReadFile(encryptionConfigFilePath)
	if err != nil {
		fmt.Printf("ERROR: failed to read file %s\n", encryptionConfigFilePath)
		panic(err)
	}

	// calculate the sum
	configShake256Sum := shake256Sum(f)

	// get node name
	nodeName := os.Getenv(EnvNodeName)
	if nodeName == "" {
		fmt.Printf("ERROR: '%s' env cannot be empty", EnvNodeName)
		panic(errors.New(fmt.Sprintf(" '%s' env cannot be empty", EnvNodeName)))
	}

	ctrlClient, err := ctrlclient.New(config.GetConfigOrDie(), ctrlclient.Options{})
	if err != nil {
		fmt.Printf("ERROR: failed to init k8s client\n")
		panic(err)
	}

	// fetch the secret
	var secret v1.Secret
	err = ctrlClient.Get(ctx,
		ctrlclient.ObjectKey{
			Name:      EncryptionProviderConfigShake256SecretName,
			Namespace: EncryptionProviderConfigShake256SecretNamespace,
		},
		&secret)

	if err != nil {
		fmt.Printf("ERROR: failed to fetch secret %s/%s\n", EncryptionProviderConfigShake256SecretNamespace, EncryptionProviderConfigShake256SecretName)
		panic(err)
	}
	// update the node
	secret.Data[nodeName] = []byte(configShake256Sum)

	// update the sum for this node in the secret
	err = ctrlClient.Update(ctx, &secret)
	if err != nil {
		fmt.Printf("ERROR: failed to fetch secret %s/%s\n", EncryptionProviderConfigShake256SecretNamespace, EncryptionProviderConfigShake256SecretName)
		panic(err)
	}

	// the file do not change during lifetime of a machine so no need to try multiple time
	// wait forever (daemonSets pod cannot exit as they would be restarted by the controller)
	select{}
}

func shake256Sum(buf []byte) string {
	h := make([]byte, 64)
	// Compute a 64-byte hash of buf and put it in h.
	sha3.ShakeSum256(h, buf)
	return fmt.Sprintf("%x\n", h)
}
