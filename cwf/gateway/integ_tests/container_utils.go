package integration

import (
	"archive/tar"
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	dockerTypes "github.com/docker/docker/api/types"
	dockerFilters "github.com/docker/docker/api/types/filters"
	dockerClient "github.com/docker/docker/client"
	"github.com/stretchr/testify/assert"
)

const (
	configDir       string = "/var/opt/magma/configs/"
	mconfigFileName string = "gateway.mconfig"
)

//RestartService adds ability to restart a particular service managed by docker
func (tr *TestRunner) RestartService(serviceName string) error {
	fmt.Printf("Restarting docker container: %v\n", serviceName)
	ctx := context.Background()
	cli, containerID, err := tr.getDockerClientAndContainerID(serviceName)
	if err != nil {
		return err
	}
	timeout := 30 * time.Second
	err = cli.ContainerRestart(ctx, containerID, &timeout)
	return err
}

//StopService adds ability to stop a particular service managed by docker
func (tr *TestRunner) PauseService(serviceName string) error {
	fmt.Printf("Pausing docker container: %v\n", serviceName)
	ctx := context.Background()
	cli, containerID, err := tr.getDockerClientAndContainerID(serviceName)
	if err != nil {
		return err
	}
	err = cli.ContainerPause(ctx, containerID)
	return err
}

// OverwriteMConfig adds ability to overwrite the mconfig file with a specified
// file
func (tr *TestRunner) OverwriteMConfig(mconfigPath string, serviceName string) error {
	fmt.Printf("Overwriting mconfig for %v with %v\n", serviceName, mconfigPath)
	ctx := context.Background()
	cli, containerID, err := tr.getDockerClientAndContainerID(serviceName)
	if err != nil {
		return err
	}

	archive, err := tr.tarArchiveFromPath(mconfigPath, mconfigFileName)
	if err != nil {
		return err
	}

	copyOptions := dockerTypes.CopyToContainerOptions{AllowOverwriteDirWithFile: true}
	err = cli.CopyToContainer(ctx, containerID, configDir, archive, copyOptions)
	if err != nil {
		return err
	}
	return nil
}

//ScanContainerLogs provides ability to scan the container logs for a string
func (tr *TestRunner) ScanContainerLogs(serviceName string, line string) int {
	ctx := context.Background()
	cli, containerID, err := tr.getDockerClientAndContainerID(serviceName)
	if err != nil {
		return 0
	}
	reader, _ := cli.ContainerLogs(ctx, containerID, dockerTypes.ContainerLogsOptions{ShowStdout: true, ShowStderr: true})
	defer reader.Close()

	b, _ := ioutil.ReadAll(reader)
	content := string(b)
	return strings.Count(content, line)
}

func (tr *TestRunner) getDockerClientAndContainerID(serviceName string) (*dockerClient.Client, string, error) {
	cli, err := dockerClient.NewEnvClient()
	if err != nil {
		fmt.Printf("error %v getting a new client \n", err)
		return nil, "", err
	}
	containerID, err := tr.findContainer(cli, serviceName)
	if err != nil {
		fmt.Printf("error %v getting container id \n", err)
		return nil, "", err
	}
	return cli, containerID, nil
}

func (tr *TestRunner) findContainer(cli *dockerClient.Client, serviceName string) (string, error) {
	filter := dockerFilters.NewArgs()
	filter.Add("name", serviceName)
	sessiondContainerFilter := dockerTypes.ContainerListOptions{Filters: filter}
	containers, err := cli.ContainerList(context.Background(), sessiondContainerFilter)
	if err != nil {
		return "", err
	}
	if len(containers) == 0 {
		return "", fmt.Errorf("container %s not found ", serviceName)
	}
	return containers[0].ID, nil
}

// path : path to the file to tar archive
// filename : the desired file name
func (tr *TestRunner) tarArchiveFromPath(path string, filename string) (io.Reader, error) {
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)

	ok := filepath.Walk(path, func(file string, fi os.FileInfo, err error) error {
		assert.NoError(tr.t, err)

		// Create header & overwrite the Name field with desired filename
		header, err := tar.FileInfoHeader(fi, filename)
		assert.NoError(tr.t, err)
		header.Name = filename
		err = tw.WriteHeader(header)
		assert.NoError(tr.t, err)

		f, err := os.Open(file)
		assert.NoError(tr.t, err)
		if fi.IsDir() {
			return nil
		}
		_, err = io.Copy(tw, f)
		assert.NoError(tr.t, err)
		assert.NoError(tr.t, f.Close())

		return nil
	})

	if ok != nil {
		return nil, ok
	}
	ok = tw.Close()
	if ok != nil {
		return nil, ok
	}
	return bufio.NewReader(&buf), nil
}
