package main
import (
    "crypto/rand"
    "crypto/ecdsa"
    "crypto/elliptic"
    "math/big"
    )
//Generates a key pair
func generateKey() (*ecdsa.PrivateKey,error) {
    pubkeyCurve := elliptic.P384()
    return ecdsa.GenerateKey(pubkeyCurve, rand.Reader)
}

func getSecret(priv *ecdsa.PrivateKey, pub ecdsa.PublicKey)(*big.Int, *big.Int){
    return elliptic.P384().ScalarMult(pub.X,pub.Y, priv.D.Bytes())
}