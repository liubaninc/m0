package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/liubaninc/m0/cmd/testnet/docs"
	toml "github.com/pelletier/go-toml"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/tendermint/tendermint/crypto/tmhash"
	tmos "github.com/tendermint/tendermint/libs/os"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	"github.com/tendermint/tendermint/types/time"
	"io/ioutil"
	"net/http"
	"os/exec"
	"path/filepath"
	"strings"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server synced server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 192.168.7.241:8080
// @BasePath /api

var (
	uploadDir = "./mytestnet"
	port      = 8080
)

func main() {
	flag.StringVar(&uploadDir, "home", uploadDir, "upload data directory")
	flag.IntVar(&port, "port", port, "listen port")

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := router.Group("/api")
	v1.POST("/chain", chain)
	v1.GET("/download/:id", download)
	if err := router.Run(fmt.Sprintf(":%d", port)); err != nil {
		panic(err)
	}
}

type ChainRequest struct {
	ChainID               string `json:"chain-id"`                  // 节点个数
	Num                   int64  `json:"num"`                       // 节点个数
	Validator             string `json:"validator-key"`             // 验证者私钥列表
	IP                    string `json:"node-ip"`                   // ip:port列表
	GenesisTime           int64  `json:"genesis-time"`              //创世时间
	ReservedAccount       string `json:"reserved-account-mnemonic"` // 预留账户
	ReservedAccountAmount string `json:"reserved-account-amount"`   // 预留账户
	Algo                  string `json:"algo"`                      // 私钥算法
}

// @生成创世块文件
// @Summary 生成创世块文件
// @Description
// @Accept  json
// @Produce json
// @Param tx body ChainRequest true "链信息"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /chain [post]
func chain(c *gin.Context) {
	var req ChainRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	args := []string{"testnet"}
	if req.Num < 1 {
		c.JSON(http.StatusBadRequest, "num must greater 1")
		return
	}

	if len(req.ChainID) == 0 {
		req.ChainID = "chain-" + tmrand.Str(6)
	}
	args = append(args, fmt.Sprintf("--chain-id=%v", req.ChainID))

	if req.Num == 0 {
		req.Num = 1
	}
	args = append(args, fmt.Sprintf("--num=%v", req.Num))

	if req.GenesisTime == 0 {
		req.GenesisTime = time.Now().Unix()
	}
	args = append(args, fmt.Sprintf("--genesis-time=%v", req.GenesisTime))

	if n := len(req.Validator); n != 0 {
		if int64(len(strings.Split(req.Validator, ","))) != req.Num {
			c.JSON(http.StatusBadRequest, "validator keys mismatch")
			return
		}
		args = append(args, fmt.Sprintf("--node-validator-key=%v", req.Validator))
	}

	if n := len(req.IP); n != 0 {
		if int64(len(strings.Split(req.IP, ","))) != req.Num {
			c.JSON(http.StatusBadRequest, "ips mismatch")
			return
		}
		args = append(args, fmt.Sprintf("--node-ip=%v", req.IP))
	}

	if len(req.Algo) != 0 {
		args = append(args, fmt.Sprintf("--algo=%v", req.Algo))
	}

	if len(req.ReservedAccount) != 0 {
		args = append(args, fmt.Sprintf("--reserved-account-mnemonic=%v", req.ReservedAccount))
	}

	if len(req.ReservedAccountAmount) != 0 {
		args = append(args, fmt.Sprintf("--reserved-coin=%v", req.ReservedAccountAmount))
	}

	hash := fmt.Sprintf("%X", tmhash.Sum([]byte(strings.Join(args, " "))))

	dst := filepath.Join(uploadDir, hash)
	//if err := tmos.EnsureDir(dst, 0755); err != nil {
	//	c.JSON(http.StatusInternalServerError, err.Error())
	//	return
	//}
	if tmos.FileExists(filepath.Join(uploadDir, hash, "node0", ".m0d", "config", "genesis.json")) {
		c.JSON(http.StatusOK, map[string]string{
			"hash":  hash,
			"seeds": seeds(hash),
		})
		return
	}
	args = append(args, fmt.Sprintf("--output-dir=%v", dst))
	cmd := exec.Command("m0d", args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	fmt.Println(cmd.String(), hash)
	if err := cmd.Run(); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error()+":"+stderr.String())
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"hash":  hash,
		"seeds": seeds(hash),
	})
}

// @下载创世块文件
// @Summary 下载创世块文件
// @Description
// @Accept  json
// @Produce json
// @Param id path string true "文件id"
// @Success 200 {object} string
// @Router /download/{id} [get]
func download(c *gin.Context) {
	dst := filepath.Join(uploadDir, c.Param("id"), "node0", ".m0d", "config", "genesis.json")
	fileContentDisposition := "attachment;filename=\"" + "genesis.json\""
	c.Header("Content-Type", "application/json")
	c.Header("Content-Disposition", fileContentDisposition)
	c.File(dst)
}

func seeds(hash string) string {
	dst := filepath.Join(uploadDir, hash, "node0", ".m0d", "config", "config.toml")
	bts, err := ioutil.ReadFile(dst)
	if err != nil {
		panic(err)
	}
	tree, err := toml.LoadBytes(bts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n\n", tree.Get("p2p.persistent_peers"))
	return tree.Get("p2p.persistent_peers").(string)
}
