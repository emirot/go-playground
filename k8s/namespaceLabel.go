package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

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
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	ns, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	for _, n := range ns.Items {
		fmt.Println(n.Name)
		ns, _ := clientset.CoreV1().Namespaces().Get(context.TODO(), n.Name, metav1.GetOptions{})
		m := ns.Labels
		m["net_addr"] = n.Name
		ns.Labels = m
		res, _ := clientset.CoreV1().Namespaces().Update(context.TODO(), ns, metav1.UpdateOptions{})
		fmt.Println(res.Labels)
	}
}
