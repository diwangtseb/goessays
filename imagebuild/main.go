package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

const DockerClientVersion = "1.37"

var imageName string

func main() {
	var tarFile string
	flag.StringVar(&tarFile, "tar", "", "tar file")
	flag.StringVar(&imageName, "image", "", "image name")
	flag.Parse()
	if tarFile == "" || imageName == "" {
		flag.Usage()
		return
	}

	//create image build's client obj
	imageBuildClient, err := client.NewClientWithOpts(client.WithVersion(DockerClientVersion))
	if err != nil {
		log.Fatal(err)
		return
	}
	// open tar file
	tarFileFp, err := os.Open(tarFile)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer tarFileFp.Close()

	// send build req
	ctx := context.Background()

	imageBuildResp, err := imageBuildClient.ImageBuild(ctx, tarFileFp, types.ImageBuildOptions{
		Tags:       []string{imageName},
		Dockerfile: "Dockerfile",
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	defer imageBuildResp.Body.Close()

	//print build output
	_, err = io.Copy(os.Stdout, imageBuildResp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	var input string
	fmt.Println("your images are ready to be pushed,yes or no?  ")
	fmt.Scanln(&input)
	if input != "yes" {
		fmt.Println("bye")
		return
	}
	var inputUserName string
	fmt.Println("please input your username:")
	fmt.Scanf("%s:\n", &inputUserName)
	var inputPwd string
	fmt.Println("please input your password:")
	fmt.Scanf("%s:\n", &inputPwd)
	if inputUserName == "" || inputPwd == "" {
		fmt.Println("username or password is empty")
		return
	}
	DockerPush(imageName, inputUserName, inputPwd)
}

func DockerPush(imageName string, username string, password string) {
	if imageName == "" {
		fmt.Println("Err: no image name specified")
		return
	}

	// 创建镜像推送的 Client 对象
	imagePushClient, err := client.NewClientWithOpts(client.WithVersion(DockerClientVersion))
	if err != nil {
		fmt.Println("Err: create docker push client error,", err.Error())
		return
	}

	// 构建镜像推送的鉴权信息
	imagePushAuthConfig := types.AuthConfig{
		Username: username,
		Password: password,
	}
	imagePushAuth, _ := json.Marshal(&imagePushAuthConfig)

	// 发送镜像推送的请求
	ctx := context.Background()
	imagePushResp, err := imagePushClient.ImagePush(ctx, imageName, types.ImagePushOptions{
		RegistryAuth: base64.URLEncoding.EncodeToString(imagePushAuth),
	})
	if err != nil {
		fmt.Println("Err: send image push request error,", err.Error())
		return
	}
	defer imagePushResp.Close()

	// 打印镜像推送的输出
	_, err = io.Copy(os.Stdout, imagePushResp)
	if err != nil {
		fmt.Println("Err: read image push response error,", err.Error())
		return
	}
}
