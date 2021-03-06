package global

import (
	"github.com/hiro942/elden-client/model"
	"github.com/tjfoc/gmsm/sm2"
)

const (
	TimeTemplate = "2006-01-02 15:04:05"
)

const (
	// FabricAppBaseUrl fabric app 地址
	FabricAppBaseUrl = "http://39.107.126.155:8080"

	// DefaultAuthenticationPort 认证服务默认端口
	DefaultAuthenticationPort = "19999"
)

const (
	// CryptoPath 加密材料存储路径
	CryptoPath             = "./.crypto/"
	PrivateKeyPemFileName  = "id_sm2"
	PublicKeyPemFileName   = "id_sm2.pub"
	SessionRecordsFileName = "session_records.json"

	// DefaultFilePerm 文件权限
	DefaultFilePerm = 0777

	// DefaultSessionKeyAge 会话密钥默认寿命
	DefaultSessionKeyAge = 3600 * 24
)

var (
	SIMCardExist = false
	MyHashedIMSI string

	MyMacAddr string

	PrivateKey             *sm2.PrivateKey
	PublicKey              *sm2.PublicKey
	PrivateKeyPath         string
	PublicKeyPath          string
	SessionRecordsFilePath string

	// PrivateKeyPwd 私钥加密密码
	PrivateKeyPwd = []byte("elden")
)

var (
	// SatellitePubKeys 卫星的公钥
	SatellitePubKeys = map[string]*sm2.PublicKey{}

	// SatelliteSockets 卫星套接字
	SatelliteSockets = map[string]string{}

	// SatelliteId 接收到的卫星广播ID
	SatelliteId string

	// HandoverSatellite 交接卫星的ID
	HandoverSatellite string

	CurrentSession = model.Session{}
)
