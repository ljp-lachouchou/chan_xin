package configserver

import (
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/conf"
)

var ErrNotSetConfig = errors.New("未设置配置信息")

type OnChange func([]byte) error
type ConfigServer interface {
	Build() error
	SetOnChange(OnChange)
	FromJsonBytes() ([]byte, error)
}
type configServer struct {
	ConfigServer
	configFile string
}

func NewConfigServer(configFile string, s ConfigServer) *configServer {
	return &configServer{
		ConfigServer: s,
		configFile:   configFile,
	}
}
func (s *configServer) MustLoad(v any, onChange OnChange) error {
	if s.configFile == "" && s.ConfigServer == nil {
		return ErrNotSetConfig
	}
	if s.ConfigServer == nil {
		//加载go-zero默认
		conf.MustLoad(s.configFile, v)
		return nil
	}
	if onChange != nil {
		s.SetOnChange(onChange)
	}
	if err := s.Build(); err != nil {
		return err
	}
	data, err := s.ConfigServer.FromJsonBytes()
	if err != nil {
		return err
	}
	return LoadFromJsonBytes(data, v)
}
func LoadFromJsonBytes(data []byte, v any) error {
	return conf.LoadFromJsonBytes(data, v)
}
