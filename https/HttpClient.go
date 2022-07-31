package httpclient

import (
  "net"
  "net/http"
  "time"

  "github.com/illublank/go-common/config"
)

// NewHttpClient todo
func NewHttpClient(config config.Config) *http.Client {
  transport := &http.Transport{
    MaxIdleConnsPerHost:   config.GetInt("Transport_MaxIdleConnsPerHost", 256),
    IdleConnTimeout:       time.Duration(config.GetInt64("Transport_IdleConnTimeout", 90)) * time.Second,
    DisableCompression:    config.GetBool("Transport_DisableCompression", true),
    MaxIdleConns:          config.GetInt("Transport_MaxIdleConns", 100),
    TLSHandshakeTimeout:   time.Duration(config.GetInt64("Transport_TLSHandshakeTimeout", 10)) * time.Second,
    ExpectContinueTimeout: time.Duration(config.GetInt64("Transport_ExpectContinueTimeout", 1)) * time.Second,
  }
  transport.DialContext = (&net.Dialer{
    Timeout:   time.Duration(config.GetInt64("Transport_DialContext_Timeout", 30)) * time.Second,
    KeepAlive: time.Duration(config.GetInt64("Transport_DialContext_KeepAlive", 30)) * time.Second,
    DualStack: config.GetBool("Transport_DialContext_DualStack", true),
  }).DialContext

  timeout := time.Duration(config.GetInt64("Timeout", 100)) * time.Second
  return &http.Client{
    Timeout:   timeout,
    Transport: transport,
  }
}
