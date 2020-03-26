package main

import (
	"context"
	"fmt"
	"k8s.io/client-go/tools/clientcmd"
	"os"

	"flag"
	"io/ioutil"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"path/filepath"
	"sigs.k8s.io/yaml"
	"strings"
)

func currentQuota(namespace string, clientset *kubernetes.Clientset) corev1.ResourceQuota {
	quotas, err := clientset.CoreV1().ResourceQuotas(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, v := range quotas.Items {
		return v
	}
	return corev1.ResourceQuota{}

}

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func readFile(filename string) (resourceQuota corev1.ResourceQuota, err error) {

	yamlByte, err := ioutil.ReadFile(filename)

	if err != nil {
		return resourceQuota, err
	}

	err = yaml.Unmarshal(yamlByte, &resourceQuota)
	if err != nil {
		return resourceQuota, err
	}
	return resourceQuota, nil
}

func compareQuota(ask corev1.ResourceQuota, namespace string, clientset *kubernetes.Clientset, clusterCapcity Capcity) (approve bool) {

	askRequestMEM := ask.Spec.Hard[corev1.ResourceRequestsMemory]

	current := currentQuota(namespace, clientset)
	currentRequestMEM := current.Spec.Hard[corev1.ResourceRequestsMemory]
	if askRequestMEM.Value() <= currentRequestMEM.Value() {
		approve = true
	} else {
		difference := askRequestMEM.Value() - currentRequestMEM.Value()
		clusterdiff := clusterCapcity.AvailableMemoryRequestTotal - difference
		if clusterdiff >= 0 {
			approve = true
		} else {
			fmt.Printf("Request Denied for namespace %s, Not enough memory to fulfil. %d\n", namespace, clusterdiff)
		}
	}
	return approve
}

func huntDir(directoryName string, clientset *kubernetes.Clientset, clusterCapcity Capcity) {
	err := filepath.Walk(directoryName,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// Is it a file?
			if info.Mode().IsRegular() {
				parentDir := strings.Split(path, "/")
				element := len(parentDir) - 2
				// Do not read files in the initial root git directory, as it will contain no resourceQuotas
				if element >= 0 {
					if parentDir[element] == directoryName {
						return nil
					}
				} else {
					return nil
				}
				quota, err := readFile(path)
				if err != nil {
					fmt.Printf("Error reading file %s, skipping\n", path)
					return nil
				}
				// We got a good resourceQuota from the file
				approve := compareQuota(quota, parentDir[element], clientset, clusterCapcity)
				if approve {
					fmt.Printf("Quota appproved for namespace: %s\n", parentDir[element])
				}
			}
			return nil
		})
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	check(err)

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	check(err)

	nodeLabel := ""
	clusterInfo := gatherInfo(clientset, &nodeLabel)
	clusterCapcity := getCapcity(clusterInfo)

	huntDir("exampleQuotaGitRepo", clientset, clusterCapcity)
}
