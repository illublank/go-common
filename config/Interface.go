package config

type Config interface {

  /*
  	1. filter config item which start with prefix
  	2. config item would be cutted out prefix
  */
  WithoutPrefix(prefix string) Config

  /*
  	1. filter config item which start with prefix (same to WithoutPrefix)
  	2. config item would be retained prefix
  */
  WithPrefix(prefix string) Config

  /*
     get and return config value and exists or not
  */
  Get(string) (interface{}, bool)

  /*
     get and return config int value if exists, but default if not
  */
  GetInt(string, int) int

  /*
     get and return config int64 value if exists, but default if not
  */
  GetInt64(string, int64) int64

  /*
     get and return config bool value if exists, but default if not
  */
  GetBool(string, bool) bool

  /*
     get and return config string value if exists, but default if not
  */
  GetString(string, string) string
}
