//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package generated

const (
	ModuleName    = "azkeys"
	ModuleVersion = "v0.5.0"
)

// ActionType - The type of the action.
type ActionType string

const (
	// ActionTypeRotate - Rotate the key based on the key policy.
	ActionTypeRotate ActionType = "rotate"
	// ActionTypeNotify - Trigger event grid events. For preview, the notification time is not configurable and it is default
	// to 30 days before expiry.
	ActionTypeNotify ActionType = "notify"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeRotate,
		ActionTypeNotify,
	}
}

// DeletionRecoveryLevel - Reflects the deletion recovery level currently in effect for keys in the current vault. If it contains
// 'Purgeable' the key can be permanently deleted by a privileged user; otherwise, only the system
// can purge the key, at the end of the retention interval.
type DeletionRecoveryLevel string

const (
	// DeletionRecoveryLevelCustomizedRecoverable - Denotes a vault state in which deletion is recoverable without the possibility
	// for immediate and permanent deletion (i.e. purge when 7<= SoftDeleteRetentionInDays < 90).This level guarantees the recoverability
	// of the deleted entity during the retention interval and while the subscription is still available.
	DeletionRecoveryLevelCustomizedRecoverable DeletionRecoveryLevel = "CustomizedRecoverable"
	// DeletionRecoveryLevelCustomizedRecoverableProtectedSubscription - Denotes a vault and subscription state in which deletion
	// is recoverable, immediate and permanent deletion (i.e. purge) is not permitted, and in which the subscription itself cannot
	// be permanently canceled when 7<= SoftDeleteRetentionInDays < 90. This level guarantees the recoverability of the deleted
	// entity during the retention interval, and also reflects the fact that the subscription itself cannot be cancelled.
	DeletionRecoveryLevelCustomizedRecoverableProtectedSubscription DeletionRecoveryLevel = "CustomizedRecoverable+ProtectedSubscription"
	// DeletionRecoveryLevelCustomizedRecoverablePurgeable - Denotes a vault state in which deletion is recoverable, and which
	// also permits immediate and permanent deletion (i.e. purge when 7<= SoftDeleteRetentionInDays < 90). This level guarantees
	// the recoverability of the deleted entity during the retention interval, unless a Purge operation is requested, or the subscription
	// is cancelled.
	DeletionRecoveryLevelCustomizedRecoverablePurgeable DeletionRecoveryLevel = "CustomizedRecoverable+Purgeable"
	// DeletionRecoveryLevelPurgeable - Denotes a vault state in which deletion is an irreversible operation, without the possibility
	// for recovery. This level corresponds to no protection being available against a Delete operation; the data is irretrievably
	// lost upon accepting a Delete operation at the entity level or higher (vault, resource group, subscription etc.)
	DeletionRecoveryLevelPurgeable DeletionRecoveryLevel = "Purgeable"
	// DeletionRecoveryLevelRecoverable - Denotes a vault state in which deletion is recoverable without the possibility for immediate
	// and permanent deletion (i.e. purge). This level guarantees the recoverability of the deleted entity during the retention
	// interval(90 days) and while the subscription is still available. System wil permanently delete it after 90 days, if not
	// recovered
	DeletionRecoveryLevelRecoverable DeletionRecoveryLevel = "Recoverable"
	// DeletionRecoveryLevelRecoverableProtectedSubscription - Denotes a vault and subscription state in which deletion is recoverable
	// within retention interval (90 days), immediate and permanent deletion (i.e. purge) is not permitted, and in which the subscription
	// itself cannot be permanently canceled. System wil permanently delete it after 90 days, if not recovered
	DeletionRecoveryLevelRecoverableProtectedSubscription DeletionRecoveryLevel = "Recoverable+ProtectedSubscription"
	// DeletionRecoveryLevelRecoverablePurgeable - Denotes a vault state in which deletion is recoverable, and which also permits
	// immediate and permanent deletion (i.e. purge). This level guarantees the recoverability of the deleted entity during the
	// retention interval (90 days), unless a Purge operation is requested, or the subscription is cancelled. System wil permanently
	// delete it after 90 days, if not recovered
	DeletionRecoveryLevelRecoverablePurgeable DeletionRecoveryLevel = "Recoverable+Purgeable"
)

// PossibleDeletionRecoveryLevelValues returns the possible values for the DeletionRecoveryLevel const type.
func PossibleDeletionRecoveryLevelValues() []DeletionRecoveryLevel {
	return []DeletionRecoveryLevel{
		DeletionRecoveryLevelCustomizedRecoverable,
		DeletionRecoveryLevelCustomizedRecoverableProtectedSubscription,
		DeletionRecoveryLevelCustomizedRecoverablePurgeable,
		DeletionRecoveryLevelPurgeable,
		DeletionRecoveryLevelRecoverable,
		DeletionRecoveryLevelRecoverableProtectedSubscription,
		DeletionRecoveryLevelRecoverablePurgeable,
	}
}

// JSONWebKeyCurveName - Elliptic curve name. For valid values, see JsonWebKeyCurveName.
type JSONWebKeyCurveName string

const (
	// JSONWebKeyCurveNameP256 - The NIST P-256 elliptic curve, AKA SECG curve SECP256R1.
	JSONWebKeyCurveNameP256 JSONWebKeyCurveName = "P-256"
	// JSONWebKeyCurveNameP256K - The SECG SECP256K1 elliptic curve.
	JSONWebKeyCurveNameP256K JSONWebKeyCurveName = "P-256K"
	// JSONWebKeyCurveNameP384 - The NIST P-384 elliptic curve, AKA SECG curve SECP384R1.
	JSONWebKeyCurveNameP384 JSONWebKeyCurveName = "P-384"
	// JSONWebKeyCurveNameP521 - The NIST P-521 elliptic curve, AKA SECG curve SECP521R1.
	JSONWebKeyCurveNameP521 JSONWebKeyCurveName = "P-521"
)

// PossibleJSONWebKeyCurveNameValues returns the possible values for the JSONWebKeyCurveName const type.
func PossibleJSONWebKeyCurveNameValues() []JSONWebKeyCurveName {
	return []JSONWebKeyCurveName{
		JSONWebKeyCurveNameP256,
		JSONWebKeyCurveNameP256K,
		JSONWebKeyCurveNameP384,
		JSONWebKeyCurveNameP521,
	}
}

// JSONWebKeyEncryptionAlgorithm - algorithm identifier
type JSONWebKeyEncryptionAlgorithm string

const (
	JSONWebKeyEncryptionAlgorithmA128CBC    JSONWebKeyEncryptionAlgorithm = "A128CBC"
	JSONWebKeyEncryptionAlgorithmA128CBCPAD JSONWebKeyEncryptionAlgorithm = "A128CBCPAD"
	JSONWebKeyEncryptionAlgorithmA128GCM    JSONWebKeyEncryptionAlgorithm = "A128GCM"
	JSONWebKeyEncryptionAlgorithmA128KW     JSONWebKeyEncryptionAlgorithm = "A128KW"
	JSONWebKeyEncryptionAlgorithmA192CBC    JSONWebKeyEncryptionAlgorithm = "A192CBC"
	JSONWebKeyEncryptionAlgorithmA192CBCPAD JSONWebKeyEncryptionAlgorithm = "A192CBCPAD"
	JSONWebKeyEncryptionAlgorithmA192GCM    JSONWebKeyEncryptionAlgorithm = "A192GCM"
	JSONWebKeyEncryptionAlgorithmA192KW     JSONWebKeyEncryptionAlgorithm = "A192KW"
	JSONWebKeyEncryptionAlgorithmA256CBC    JSONWebKeyEncryptionAlgorithm = "A256CBC"
	JSONWebKeyEncryptionAlgorithmA256CBCPAD JSONWebKeyEncryptionAlgorithm = "A256CBCPAD"
	JSONWebKeyEncryptionAlgorithmA256GCM    JSONWebKeyEncryptionAlgorithm = "A256GCM"
	JSONWebKeyEncryptionAlgorithmA256KW     JSONWebKeyEncryptionAlgorithm = "A256KW"
	JSONWebKeyEncryptionAlgorithmRSA15      JSONWebKeyEncryptionAlgorithm = "RSA1_5"
	JSONWebKeyEncryptionAlgorithmRSAOAEP    JSONWebKeyEncryptionAlgorithm = "RSA-OAEP"
	JSONWebKeyEncryptionAlgorithmRSAOAEP256 JSONWebKeyEncryptionAlgorithm = "RSA-OAEP-256"
)

// PossibleJSONWebKeyEncryptionAlgorithmValues returns the possible values for the JSONWebKeyEncryptionAlgorithm const type.
func PossibleJSONWebKeyEncryptionAlgorithmValues() []JSONWebKeyEncryptionAlgorithm {
	return []JSONWebKeyEncryptionAlgorithm{
		JSONWebKeyEncryptionAlgorithmA128CBC,
		JSONWebKeyEncryptionAlgorithmA128CBCPAD,
		JSONWebKeyEncryptionAlgorithmA128GCM,
		JSONWebKeyEncryptionAlgorithmA128KW,
		JSONWebKeyEncryptionAlgorithmA192CBC,
		JSONWebKeyEncryptionAlgorithmA192CBCPAD,
		JSONWebKeyEncryptionAlgorithmA192GCM,
		JSONWebKeyEncryptionAlgorithmA192KW,
		JSONWebKeyEncryptionAlgorithmA256CBC,
		JSONWebKeyEncryptionAlgorithmA256CBCPAD,
		JSONWebKeyEncryptionAlgorithmA256GCM,
		JSONWebKeyEncryptionAlgorithmA256KW,
		JSONWebKeyEncryptionAlgorithmRSA15,
		JSONWebKeyEncryptionAlgorithmRSAOAEP,
		JSONWebKeyEncryptionAlgorithmRSAOAEP256,
	}
}

// JSONWebKeyOperation - JSON web key operations. For more information, see JsonWebKeyOperation.
type JSONWebKeyOperation string

const (
	JSONWebKeyOperationDecrypt   JSONWebKeyOperation = "decrypt"
	JSONWebKeyOperationEncrypt   JSONWebKeyOperation = "encrypt"
	JSONWebKeyOperationExport    JSONWebKeyOperation = "export"
	JSONWebKeyOperationImport    JSONWebKeyOperation = "import"
	JSONWebKeyOperationSign      JSONWebKeyOperation = "sign"
	JSONWebKeyOperationUnwrapKey JSONWebKeyOperation = "unwrapKey"
	JSONWebKeyOperationVerify    JSONWebKeyOperation = "verify"
	JSONWebKeyOperationWrapKey   JSONWebKeyOperation = "wrapKey"
)

// PossibleJSONWebKeyOperationValues returns the possible values for the JSONWebKeyOperation const type.
func PossibleJSONWebKeyOperationValues() []JSONWebKeyOperation {
	return []JSONWebKeyOperation{
		JSONWebKeyOperationDecrypt,
		JSONWebKeyOperationEncrypt,
		JSONWebKeyOperationExport,
		JSONWebKeyOperationImport,
		JSONWebKeyOperationSign,
		JSONWebKeyOperationUnwrapKey,
		JSONWebKeyOperationVerify,
		JSONWebKeyOperationWrapKey,
	}
}

// JSONWebKeySignatureAlgorithm - The signing/verification algorithm identifier. For more information on possible algorithm
// types, see JsonWebKeySignatureAlgorithm.
type JSONWebKeySignatureAlgorithm string

const (
	// JSONWebKeySignatureAlgorithmES256 - ECDSA using P-256 and SHA-256, as described in https://tools.ietf.org/html/rfc7518.
	JSONWebKeySignatureAlgorithmES256 JSONWebKeySignatureAlgorithm = "ES256"
	// JSONWebKeySignatureAlgorithmES256K - ECDSA using P-256K and SHA-256, as described in https://tools.ietf.org/html/rfc7518
	JSONWebKeySignatureAlgorithmES256K JSONWebKeySignatureAlgorithm = "ES256K"
	// JSONWebKeySignatureAlgorithmES384 - ECDSA using P-384 and SHA-384, as described in https://tools.ietf.org/html/rfc7518
	JSONWebKeySignatureAlgorithmES384 JSONWebKeySignatureAlgorithm = "ES384"
	// JSONWebKeySignatureAlgorithmES512 - ECDSA using P-521 and SHA-512, as described in https://tools.ietf.org/html/rfc7518
	JSONWebKeySignatureAlgorithmES512 JSONWebKeySignatureAlgorithm = "ES512"
	// JSONWebKeySignatureAlgorithmPS256 - RSASSA-PSS using SHA-256 and MGF1 with SHA-256, as described in https://tools.ietf.org/html/rfc7518
	JSONWebKeySignatureAlgorithmPS256 JSONWebKeySignatureAlgorithm = "PS256"
	// JSONWebKeySignatureAlgorithmPS384 - RSASSA-PSS using SHA-384 and MGF1 with SHA-384, as described in https://tools.ietf.org/html/rfc7518
	JSONWebKeySignatureAlgorithmPS384 JSONWebKeySignatureAlgorithm = "PS384"
	// JSONWebKeySignatureAlgorithmPS512 - RSASSA-PSS using SHA-512 and MGF1 with SHA-512, as described in https://tools.ietf.org/html/rfc7518
	JSONWebKeySignatureAlgorithmPS512 JSONWebKeySignatureAlgorithm = "PS512"
	// JSONWebKeySignatureAlgorithmRS256 - RSASSA-PKCS1-v1_5 using SHA-256, as described in https://tools.ietf.org/html/rfc7518
	JSONWebKeySignatureAlgorithmRS256 JSONWebKeySignatureAlgorithm = "RS256"
	// JSONWebKeySignatureAlgorithmRS384 - RSASSA-PKCS1-v1_5 using SHA-384, as described in https://tools.ietf.org/html/rfc7518
	JSONWebKeySignatureAlgorithmRS384 JSONWebKeySignatureAlgorithm = "RS384"
	// JSONWebKeySignatureAlgorithmRS512 - RSASSA-PKCS1-v1_5 using SHA-512, as described in https://tools.ietf.org/html/rfc7518
	JSONWebKeySignatureAlgorithmRS512 JSONWebKeySignatureAlgorithm = "RS512"
	// JSONWebKeySignatureAlgorithmRSNULL - Reserved
	JSONWebKeySignatureAlgorithmRSNULL JSONWebKeySignatureAlgorithm = "RSNULL"
)

// PossibleJSONWebKeySignatureAlgorithmValues returns the possible values for the JSONWebKeySignatureAlgorithm const type.
func PossibleJSONWebKeySignatureAlgorithmValues() []JSONWebKeySignatureAlgorithm {
	return []JSONWebKeySignatureAlgorithm{
		JSONWebKeySignatureAlgorithmES256,
		JSONWebKeySignatureAlgorithmES256K,
		JSONWebKeySignatureAlgorithmES384,
		JSONWebKeySignatureAlgorithmES512,
		JSONWebKeySignatureAlgorithmPS256,
		JSONWebKeySignatureAlgorithmPS384,
		JSONWebKeySignatureAlgorithmPS512,
		JSONWebKeySignatureAlgorithmRS256,
		JSONWebKeySignatureAlgorithmRS384,
		JSONWebKeySignatureAlgorithmRS512,
		JSONWebKeySignatureAlgorithmRSNULL,
	}
}

// JSONWebKeyType - JsonWebKey Key Type (kty), as defined in https://tools.ietf.org/html/draft-ietf-jose-json-web-algorithms-40.
type JSONWebKeyType string

const (
	// JSONWebKeyTypeEC - Elliptic Curve.
	JSONWebKeyTypeEC JSONWebKeyType = "EC"
	// JSONWebKeyTypeECHSM - Elliptic Curve with a private key which is stored in the HSM.
	JSONWebKeyTypeECHSM JSONWebKeyType = "EC-HSM"
	// JSONWebKeyTypeOct - Octet sequence (used to represent symmetric keys)
	JSONWebKeyTypeOct JSONWebKeyType = "oct"
	// JSONWebKeyTypeOctHSM - Octet sequence (used to represent symmetric keys) which is stored the HSM.
	JSONWebKeyTypeOctHSM JSONWebKeyType = "oct-HSM"
	// JSONWebKeyTypeRSA - RSA (https://tools.ietf.org/html/rfc3447)
	JSONWebKeyTypeRSA JSONWebKeyType = "RSA"
	// JSONWebKeyTypeRSAHSM - RSA with a private key which is stored in the HSM.
	JSONWebKeyTypeRSAHSM JSONWebKeyType = "RSA-HSM"
)

// PossibleJSONWebKeyTypeValues returns the possible values for the JSONWebKeyType const type.
func PossibleJSONWebKeyTypeValues() []JSONWebKeyType {
	return []JSONWebKeyType{
		JSONWebKeyTypeEC,
		JSONWebKeyTypeECHSM,
		JSONWebKeyTypeOct,
		JSONWebKeyTypeOctHSM,
		JSONWebKeyTypeRSA,
		JSONWebKeyTypeRSAHSM,
	}
}

// KeyEncryptionAlgorithm - The encryption algorithm to use to protected the exported key material
type KeyEncryptionAlgorithm string

const (
	KeyEncryptionAlgorithmCKMRSAAESKEYWRAP KeyEncryptionAlgorithm = "CKM_RSA_AES_KEY_WRAP"
	KeyEncryptionAlgorithmRSAAESKEYWRAP256 KeyEncryptionAlgorithm = "RSA_AES_KEY_WRAP_256"
	KeyEncryptionAlgorithmRSAAESKEYWRAP384 KeyEncryptionAlgorithm = "RSA_AES_KEY_WRAP_384"
)

// PossibleKeyEncryptionAlgorithmValues returns the possible values for the KeyEncryptionAlgorithm const type.
func PossibleKeyEncryptionAlgorithmValues() []KeyEncryptionAlgorithm {
	return []KeyEncryptionAlgorithm{
		KeyEncryptionAlgorithmCKMRSAAESKEYWRAP,
		KeyEncryptionAlgorithmRSAAESKEYWRAP256,
		KeyEncryptionAlgorithmRSAAESKEYWRAP384,
	}
}
