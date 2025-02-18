package datastruc

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

type Hashtable struct {
	Keys   []string
	Values []int
}

func hashData(value int) string {
	hashingFunc := sha256.New()
	hashingFunc.Write([]byte(fmt.Sprint(value)))

	key := hashingFunc.Sum(nil)
	return fmt.Sprintf("%x", key)
}

func (h *Hashtable) Put(value int) {
	h.Values = append(h.Values, value)
	h.Keys = append(h.Keys, hashData(value))
}

func (h *Hashtable) Get(key string) (int, error) {
	for i, k := range h.Keys {
		if k == key {
			return h.Values[i], nil
		}
	}
	return -1, errors.New("key not found")
}

func (h *Hashtable) Delete(key string) (int, error) {
	for i, k := range h.Keys {
		if k == key {
			output := h.Values[i]
			if i == 0 && len(h.Values) > 1 {
				h.Values = h.Values[1:]
				h.Keys = h.Keys[1:]
			} else if i == 0 && len(h.Values) <= 1 {
				h.Values = []int{}
				h.Keys = []string{}
			} else if i > 0 && i < len(h.Values)-1 {
				h.Values = append(h.Values[:i], h.Values[i+1:]...)
				h.Keys = append(h.Keys[:i], h.Keys[i+1:]...)
			}
			return output, nil
		}
	}
	return -1, errors.New("key not found")
}
