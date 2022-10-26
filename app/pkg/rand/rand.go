package rand

import (
	"crypto/rand"
	"math/big"

	"github.com/sirupsen/logrus"
)

func GetRandomCoordinates(w, h int) (x, y int, err error) {
	x, err = getRandomNumber(w)
	if err != nil {
		return 0, 0, err
	}

	y, err = getRandomNumber(h)
	if err != nil {
		return 0, 0, err
	}

	return x, y, nil
}

func getRandomNumber(n int) (int, error) {
	r, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	if err != nil {
		logrus.Error(err)
		return 0, err
	}
	return int(r.Int64()), nil
}
