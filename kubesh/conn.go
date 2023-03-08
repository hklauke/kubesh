package kubesh

// import (
// 	"context"
// 	"os"

// 	"golang.org/x/term"
// 	"k8s.io/client-go/kubernetes"
// 	"k8s.io/client-go/kubernetes/scheme"
// 	"k8s.io/client-go/tools/remotecommand"
// )

// func StartConn(ctx context.Context, namespace string, client kubernetes.Interface , config kubernetes.Interface, pod string, container string) {

// 	req := client.RESTClient().
// 		Post().
// 		Namespace(namespace).
// 		Resource("pods").
// 		Name(pod).
// 		SubResource("exec").
// 		VersionedParams(&corev1.PodExecOptions{
// 			Container: container,
// 			Command:   []string{"nsenter", "--target", "1", "--mount", "--uts", "--ipc", "--net", "--pid", "--", "/bin/bash"},
// 			Stdin:     true,
// 			Stdout:    true,
// 			Stderr:    true,
// 			TTY:       true,
// 		}, scheme.ParameterCodec)

// 	exec, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
// 	if err != nil {
// 		panic(err)
// 	}

// 	oldState, err := term.MakeRaw(0)
// 	if err != nil {
// 		panic(err)
// 	}

// 	termWidth, termHeight, _ := term.GetSize(0)
// 	termSize := remotecommand.TerminalSize{Width: uint16(termWidth), Height: uint16(termHeight)}
// 	s := make(remotecommand.TerminalSize, 1)
// 	s <- termSize

// 	defer func() {
// 		err := term.Restore(0, oldState)
// 		if err != nil {
// 			panic(err)
// 		}
// 	}()

// 	err = exec.Stream(remotecommand.StreamOptions{
// 		Stdin:             os.Stdin,
// 		Stdout:            os.Stdout,
// 		Stderr:            os.Stderr,
// 		Tty:               true,
// 		TerminalSizeQueue: s,
// 	})
// 	if err != nil {
// 		panic(err)
// 	}

// }
