package main

import "testing"

func TestGenerateKey(t *testing.T){
    _ , err := generateKey()
    if err != nil {
        t.Fatalf(err.Error())
    }
}
//Tests if secret computation is correct
func TestSecretComputation(t *testing.T){
    alice, _ := generateKey()
    bob, _ := generateKey()
    aliceSecretX, aliceSecretY := getSecret(alice,bob.PublicKey)
    bobSecretX, bobSecretY := getSecret(bob,alice.PublicKey)
    t.Log(aliceSecretX)
    if aliceSecretX.Cmp(bobSecretX)!=0 {
        t.Fatalf("Multiplication Failed")
    }
    if aliceSecretY.Cmp(bobSecretY)!=0 {
        t.Fatalf("Muliplication Failed")
    }
}
