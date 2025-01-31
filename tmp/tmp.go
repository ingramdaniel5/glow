package tmp

const (
	TX_DEPLOY_CONTRACT = `
        transaction(name: String, code: String) {
            prepare(signer: AuthAccount) {
                signer.contracts.add(name: name, code: code.decodeHex())
            }
        }
    `

	TX_CREATE_ACCOUNT = `
		transaction(publicKey: String) {
			prepare(signer: AuthAccount) {
				let account = AuthAccount(payer: signer)
				let accountKey = PublicKey(
					publicKey: publicKey.decodeHex(),
					signatureAlgorithm: SignatureAlgorithm.ECDSA_P256
				)
				account.keys.add(
					publicKey: accountKey,
					hashAlgorithm: HashAlgorithm.SHA3_256,
					weight: 1000.0
				)
			}
		}
	`

	// todo...
)
