package conf

import (
	"testing"
)

const path = "../../configs/config.yaml"

func TestReadConfig(t *testing.T) {
	notify := make(chan bool, 1)
	v := NewLoadConf(notify)
	err := v.Load(path)
	if err != nil {
		t.Log(err)
	}
	c := &Conf{}
	c, err = v.Parse(c)
	if err != nil {
		t.Log(err)
	}
	t.Log(c.Db)
}
