// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msp/msp_config.proto

package msp // import "github.com/Matrix-Zhang/fabric-gm/protos/msp"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// MSPConfig collects all the configuration information for
// an MSP. The Config field should be unmarshalled in a way
// that depends on the Type
type MSPConfig struct {
	// Type holds the type of the MSP; the default one would
	// be of type FABRIC implementing an X.509 based provider
	Type int32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	// Config is MSP dependent configuration info
	Config               []byte   `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MSPConfig) Reset()         { *m = MSPConfig{} }
func (m *MSPConfig) String() string { return proto.CompactTextString(m) }
func (*MSPConfig) ProtoMessage()    {}
func (*MSPConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_msp_config_c4b502fff5780348, []int{0}
}
func (m *MSPConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MSPConfig.Unmarshal(m, b)
}
func (m *MSPConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MSPConfig.Marshal(b, m, deterministic)
}
func (dst *MSPConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MSPConfig.Merge(dst, src)
}
func (m *MSPConfig) XXX_Size() int {
	return xxx_messageInfo_MSPConfig.Size(m)
}
func (m *MSPConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MSPConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MSPConfig proto.InternalMessageInfo

func (m *MSPConfig) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *MSPConfig) GetConfig() []byte {
	if m != nil {
		return m.Config
	}
	return nil
}

// FabricMSPConfig collects all the configuration information for
// a Fabric MSP.
// Here we assume a default certificate validation policy, where
// any certificate signed by any of the listed rootCA certs would
// be considered as valid under this MSP.
// This MSP may or may not come with a signing identity. If it does,
// it can also issue signing identities. If it does not, it can only
// be used to validate and verify certificates.
type FabricMSPConfig struct {
	// Name holds the identifier of the MSP; MSP identifier
	// is chosen by the application that governs this MSP.
	// For example, and assuming the default implementation of MSP,
	// that is X.509-based and considers a single Issuer,
	// this can refer to the Subject OU field or the Issuer OU field.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// List of root certificates trusted by this MSP
	// they are used upon certificate validation (see
	// comment for IntermediateCerts below)
	RootCerts [][]byte `protobuf:"bytes,2,rep,name=root_certs,json=rootCerts,proto3" json:"root_certs,omitempty"`
	// List of intermediate certificates trusted by this MSP;
	// they are used upon certificate validation as follows:
	// validation attempts to build a path from the certificate
	// to be validated (which is at one end of the path) and
	// one of the certs in the RootCerts field (which is at
	// the other end of the path). If the path is longer than
	// 2, certificates in the middle are searched within the
	// IntermediateCerts pool
	IntermediateCerts [][]byte `protobuf:"bytes,3,rep,name=intermediate_certs,json=intermediateCerts,proto3" json:"intermediate_certs,omitempty"`
	// Identity denoting the administrator of this MSP
	Admins [][]byte `protobuf:"bytes,4,rep,name=admins,proto3" json:"admins,omitempty"`
	// Identity revocation list
	RevocationList [][]byte `protobuf:"bytes,5,rep,name=revocation_list,json=revocationList,proto3" json:"revocation_list,omitempty"`
	// SigningIdentity holds information on the signing identity
	// this peer is to use, and which is to be imported by the
	// MSP defined before
	SigningIdentity *SigningIdentityInfo `protobuf:"bytes,6,opt,name=signing_identity,json=signingIdentity,proto3" json:"signing_identity,omitempty"`
	// OrganizationalUnitIdentifiers holds one or more
	// fabric organizational unit identifiers that belong to
	// this MSP configuration
	OrganizationalUnitIdentifiers []*FabricOUIdentifier `protobuf:"bytes,7,rep,name=organizational_unit_identifiers,json=organizationalUnitIdentifiers,proto3" json:"organizational_unit_identifiers,omitempty"`
	// FabricCryptoConfig contains the configuration parameters
	// for the cryptographic algorithms used by this MSP
	CryptoConfig *FabricCryptoConfig `protobuf:"bytes,8,opt,name=crypto_config,json=cryptoConfig,proto3" json:"crypto_config,omitempty"`
	// List of TLS root certificates trusted by this MSP.
	// They are returned by GetTLSRootCerts.
	TlsRootCerts [][]byte `protobuf:"bytes,9,rep,name=tls_root_certs,json=tlsRootCerts,proto3" json:"tls_root_certs,omitempty"`
	// List of TLS intermediate certificates trusted by this MSP;
	// They are returned by GetTLSIntermediateCerts.
	TlsIntermediateCerts [][]byte `protobuf:"bytes,10,rep,name=tls_intermediate_certs,json=tlsIntermediateCerts,proto3" json:"tls_intermediate_certs,omitempty"`
	// fabric_node_ous contains the configuration to distinguish clients from peers from orderers
	// based on the OUs.
	FabricNodeOus        *FabricNodeOUs `protobuf:"bytes,11,opt,name=fabric_node_ous,json=fabricNodeOus,proto3" json:"fabric_node_ous,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FabricMSPConfig) Reset()         { *m = FabricMSPConfig{} }
func (m *FabricMSPConfig) String() string { return proto.CompactTextString(m) }
func (*FabricMSPConfig) ProtoMessage()    {}
func (*FabricMSPConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_msp_config_c4b502fff5780348, []int{1}
}
func (m *FabricMSPConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FabricMSPConfig.Unmarshal(m, b)
}
func (m *FabricMSPConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FabricMSPConfig.Marshal(b, m, deterministic)
}
func (dst *FabricMSPConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FabricMSPConfig.Merge(dst, src)
}
func (m *FabricMSPConfig) XXX_Size() int {
	return xxx_messageInfo_FabricMSPConfig.Size(m)
}
func (m *FabricMSPConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FabricMSPConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FabricMSPConfig proto.InternalMessageInfo

func (m *FabricMSPConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FabricMSPConfig) GetRootCerts() [][]byte {
	if m != nil {
		return m.RootCerts
	}
	return nil
}

func (m *FabricMSPConfig) GetIntermediateCerts() [][]byte {
	if m != nil {
		return m.IntermediateCerts
	}
	return nil
}

func (m *FabricMSPConfig) GetAdmins() [][]byte {
	if m != nil {
		return m.Admins
	}
	return nil
}

func (m *FabricMSPConfig) GetRevocationList() [][]byte {
	if m != nil {
		return m.RevocationList
	}
	return nil
}

func (m *FabricMSPConfig) GetSigningIdentity() *SigningIdentityInfo {
	if m != nil {
		return m.SigningIdentity
	}
	return nil
}

func (m *FabricMSPConfig) GetOrganizationalUnitIdentifiers() []*FabricOUIdentifier {
	if m != nil {
		return m.OrganizationalUnitIdentifiers
	}
	return nil
}

func (m *FabricMSPConfig) GetCryptoConfig() *FabricCryptoConfig {
	if m != nil {
		return m.CryptoConfig
	}
	return nil
}

func (m *FabricMSPConfig) GetTlsRootCerts() [][]byte {
	if m != nil {
		return m.TlsRootCerts
	}
	return nil
}

func (m *FabricMSPConfig) GetTlsIntermediateCerts() [][]byte {
	if m != nil {
		return m.TlsIntermediateCerts
	}
	return nil
}

func (m *FabricMSPConfig) GetFabricNodeOus() *FabricNodeOUs {
	if m != nil {
		return m.FabricNodeOus
	}
	return nil
}

// FabricCryptoConfig contains configuration parameters
// for the cryptographic algorithms used by the MSP
// this configuration refers to
type FabricCryptoConfig struct {
	// SignatureHashFamily is a string representing the hash family to be used
	// during sign and verify operations.
	// Allowed values are "SHA2" and "SHA3".
	SignatureHashFamily string `protobuf:"bytes,1,opt,name=signature_hash_family,json=signatureHashFamily,proto3" json:"signature_hash_family,omitempty"`
	// IdentityIdentifierHashFunction is a string representing the hash function
	// to be used during the computation of the identity identifier of an MSP identity.
	// Allowed values are "SHA256", "SHA384" and "SHA3_256", "SHA3_384".
	IdentityIdentifierHashFunction string   `protobuf:"bytes,2,opt,name=identity_identifier_hash_function,json=identityIdentifierHashFunction,proto3" json:"identity_identifier_hash_function,omitempty"`
	XXX_NoUnkeyedLiteral           struct{} `json:"-"`
	XXX_unrecognized               []byte   `json:"-"`
	XXX_sizecache                  int32    `json:"-"`
}

func (m *FabricCryptoConfig) Reset()         { *m = FabricCryptoConfig{} }
func (m *FabricCryptoConfig) String() string { return proto.CompactTextString(m) }
func (*FabricCryptoConfig) ProtoMessage()    {}
func (*FabricCryptoConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_msp_config_c4b502fff5780348, []int{2}
}
func (m *FabricCryptoConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FabricCryptoConfig.Unmarshal(m, b)
}
func (m *FabricCryptoConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FabricCryptoConfig.Marshal(b, m, deterministic)
}
func (dst *FabricCryptoConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FabricCryptoConfig.Merge(dst, src)
}
func (m *FabricCryptoConfig) XXX_Size() int {
	return xxx_messageInfo_FabricCryptoConfig.Size(m)
}
func (m *FabricCryptoConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FabricCryptoConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FabricCryptoConfig proto.InternalMessageInfo

func (m *FabricCryptoConfig) GetSignatureHashFamily() string {
	if m != nil {
		return m.SignatureHashFamily
	}
	return ""
}

func (m *FabricCryptoConfig) GetIdentityIdentifierHashFunction() string {
	if m != nil {
		return m.IdentityIdentifierHashFunction
	}
	return ""
}

// IdemixMSPConfig collects all the configuration information for
// an Idemix MSP.
type IdemixMSPConfig struct {
	// Name holds the identifier of the MSP
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// ipk represents the (serialized) issuer public key
	Ipk []byte `protobuf:"bytes,2,opt,name=ipk,proto3" json:"ipk,omitempty"`
	// signer may contain crypto material to configure a default signer
	Signer *IdemixMSPSignerConfig `protobuf:"bytes,3,opt,name=signer,proto3" json:"signer,omitempty"`
	// revocation_pk is the public key used for revocation of credentials
	RevocationPk []byte `protobuf:"bytes,4,opt,name=revocation_pk,json=revocationPk,proto3" json:"revocation_pk,omitempty"`
	// epoch represents the current epoch (time interval) used for revocation
	Epoch                int64    `protobuf:"varint,5,opt,name=epoch,proto3" json:"epoch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdemixMSPConfig) Reset()         { *m = IdemixMSPConfig{} }
func (m *IdemixMSPConfig) String() string { return proto.CompactTextString(m) }
func (*IdemixMSPConfig) ProtoMessage()    {}
func (*IdemixMSPConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_msp_config_c4b502fff5780348, []int{3}
}
func (m *IdemixMSPConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdemixMSPConfig.Unmarshal(m, b)
}
func (m *IdemixMSPConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdemixMSPConfig.Marshal(b, m, deterministic)
}
func (dst *IdemixMSPConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdemixMSPConfig.Merge(dst, src)
}
func (m *IdemixMSPConfig) XXX_Size() int {
	return xxx_messageInfo_IdemixMSPConfig.Size(m)
}
func (m *IdemixMSPConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_IdemixMSPConfig.DiscardUnknown(m)
}

var xxx_messageInfo_IdemixMSPConfig proto.InternalMessageInfo

func (m *IdemixMSPConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *IdemixMSPConfig) GetIpk() []byte {
	if m != nil {
		return m.Ipk
	}
	return nil
}

func (m *IdemixMSPConfig) GetSigner() *IdemixMSPSignerConfig {
	if m != nil {
		return m.Signer
	}
	return nil
}

func (m *IdemixMSPConfig) GetRevocationPk() []byte {
	if m != nil {
		return m.RevocationPk
	}
	return nil
}

func (m *IdemixMSPConfig) GetEpoch() int64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

// IdemixMSPSIgnerConfig contains the crypto material to set up an idemix signing identity
type IdemixMSPSignerConfig struct {
	// cred represents the serialized idemix credential of the default signer
	Cred []byte `protobuf:"bytes,1,opt,name=cred,proto3" json:"cred,omitempty"`
	// sk is the secret key of the default signer, corresponding to credential Cred
	Sk []byte `protobuf:"bytes,2,opt,name=sk,proto3" json:"sk,omitempty"`
	// organizational_unit_identifier defines the organizational unit the default signer is in
	OrganizationalUnitIdentifier string `protobuf:"bytes,3,opt,name=organizational_unit_identifier,json=organizationalUnitIdentifier,proto3" json:"organizational_unit_identifier,omitempty"`
	// role defines whether the default signer is admin, peer, member or client
	Role int32 `protobuf:"varint,4,opt,name=role,proto3" json:"role,omitempty"`
	// enrollment_id contains the enrollment id of this signer
	EnrollmentId string `protobuf:"bytes,5,opt,name=enrollment_id,json=enrollmentId,proto3" json:"enrollment_id,omitempty"`
	// credential_revocation_information contains a serialized CredentialRevocationInformation
	CredentialRevocationInformation []byte   `protobuf:"bytes,6,opt,name=credential_revocation_information,json=credentialRevocationInformation,proto3" json:"credential_revocation_information,omitempty"`
	XXX_NoUnkeyedLiteral            struct{} `json:"-"`
	XXX_unrecognized                []byte   `json:"-"`
	XXX_sizecache                   int32    `json:"-"`
}

func (m *IdemixMSPSignerConfig) Reset()         { *m = IdemixMSPSignerConfig{} }
func (m *IdemixMSPSignerConfig) String() string { return proto.CompactTextString(m) }
func (*IdemixMSPSignerConfig) ProtoMessage()    {}
func (*IdemixMSPSignerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_msp_config_c4b502fff5780348, []int{4}
}
func (m *IdemixMSPSignerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdemixMSPSignerConfig.Unmarshal(m, b)
}
func (m *IdemixMSPSignerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdemixMSPSignerConfig.Marshal(b, m, deterministic)
}
func (dst *IdemixMSPSignerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdemixMSPSignerConfig.Merge(dst, src)
}
func (m *IdemixMSPSignerConfig) XXX_Size() int {
	return xxx_messageInfo_IdemixMSPSignerConfig.Size(m)
}
func (m *IdemixMSPSignerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_IdemixMSPSignerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_IdemixMSPSignerConfig proto.InternalMessageInfo

func (m *IdemixMSPSignerConfig) GetCred() []byte {
	if m != nil {
		return m.Cred
	}
	return nil
}

func (m *IdemixMSPSignerConfig) GetSk() []byte {
	if m != nil {
		return m.Sk
	}
	return nil
}

func (m *IdemixMSPSignerConfig) GetOrganizationalUnitIdentifier() string {
	if m != nil {
		return m.OrganizationalUnitIdentifier
	}
	return ""
}

func (m *IdemixMSPSignerConfig) GetRole() int32 {
	if m != nil {
		return m.Role
	}
	return 0
}

func (m *IdemixMSPSignerConfig) GetEnrollmentId() string {
	if m != nil {
		return m.EnrollmentId
	}
	return ""
}

func (m *IdemixMSPSignerConfig) GetCredentialRevocationInformation() []byte {
	if m != nil {
		return m.CredentialRevocationInformation
	}
	return nil
}

// SigningIdentityInfo represents the configuration information
// related to the signing identity the peer is to use for generating
// endorsements
type SigningIdentityInfo struct {
	// PublicSigner carries the public information of the signing
	// identity. For an X.509 provider this would be represented by
	// an X.509 certificate
	PublicSigner []byte `protobuf:"bytes,1,opt,name=public_signer,json=publicSigner,proto3" json:"public_signer,omitempty"`
	// PrivateSigner denotes a reference to the private key of the
	// peer's signing identity
	PrivateSigner        *KeyInfo `protobuf:"bytes,2,opt,name=private_signer,json=privateSigner,proto3" json:"private_signer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SigningIdentityInfo) Reset()         { *m = SigningIdentityInfo{} }
func (m *SigningIdentityInfo) String() string { return proto.CompactTextString(m) }
func (*SigningIdentityInfo) ProtoMessage()    {}
func (*SigningIdentityInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_msp_config_c4b502fff5780348, []int{5}
}
func (m *SigningIdentityInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SigningIdentityInfo.Unmarshal(m, b)
}
func (m *SigningIdentityInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SigningIdentityInfo.Marshal(b, m, deterministic)
}
func (dst *SigningIdentityInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SigningIdentityInfo.Merge(dst, src)
}
func (m *SigningIdentityInfo) XXX_Size() int {
	return xxx_messageInfo_SigningIdentityInfo.Size(m)
}
func (m *SigningIdentityInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SigningIdentityInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SigningIdentityInfo proto.InternalMessageInfo

func (m *SigningIdentityInfo) GetPublicSigner() []byte {
	if m != nil {
		return m.PublicSigner
	}
	return nil
}

func (m *SigningIdentityInfo) GetPrivateSigner() *KeyInfo {
	if m != nil {
		return m.PrivateSigner
	}
	return nil
}

// KeyInfo represents a (secret) key that is either already stored
// in the bccsp/keystore or key material to be imported to the
// bccsp key-store. In later versions it may contain also a
// keystore identifier
type KeyInfo struct {
	// Identifier of the key inside the default keystore; this for
	// the case of Software BCCSP as well as the HSM BCCSP would be
	// the SKI of the key
	KeyIdentifier string `protobuf:"bytes,1,opt,name=key_identifier,json=keyIdentifier,proto3" json:"key_identifier,omitempty"`
	// KeyMaterial (optional) for the key to be imported; this is
	// properly encoded key bytes, prefixed by the type of the key
	KeyMaterial          []byte   `protobuf:"bytes,2,opt,name=key_material,json=keyMaterial,proto3" json:"key_material,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyInfo) Reset()         { *m = KeyInfo{} }
func (m *KeyInfo) String() string { return proto.CompactTextString(m) }
func (*KeyInfo) ProtoMessage()    {}
func (*KeyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_msp_config_c4b502fff5780348, []int{6}
}
func (m *KeyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyInfo.Unmarshal(m, b)
}
func (m *KeyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyInfo.Marshal(b, m, deterministic)
}
func (dst *KeyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyInfo.Merge(dst, src)
}
func (m *KeyInfo) XXX_Size() int {
	return xxx_messageInfo_KeyInfo.Size(m)
}
func (m *KeyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_KeyInfo proto.InternalMessageInfo

func (m *KeyInfo) GetKeyIdentifier() string {
	if m != nil {
		return m.KeyIdentifier
	}
	return ""
}

func (m *KeyInfo) GetKeyMaterial() []byte {
	if m != nil {
		return m.KeyMaterial
	}
	return nil
}

// FabricOUIdentifier represents an organizational unit and
// its related chain of trust identifier.
type FabricOUIdentifier struct {
	// Certificate represents the second certificate in a certification chain.
	// (Notice that the first certificate in a certification chain is supposed
	// to be the certificate of an identity).
	// It must correspond to the certificate of root or intermediate CA
	// recognized by the MSP this message belongs to.
	// Starting from this certificate, a certification chain is computed
	// and bound to the OrganizationUnitIdentifier specified
	Certificate []byte `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate,omitempty"`
	// OrganizationUnitIdentifier defines the organizational unit under the
	// MSP identified with MSPIdentifier
	OrganizationalUnitIdentifier string   `protobuf:"bytes,2,opt,name=organizational_unit_identifier,json=organizationalUnitIdentifier,proto3" json:"organizational_unit_identifier,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	XXX_unrecognized             []byte   `json:"-"`
	XXX_sizecache                int32    `json:"-"`
}

func (m *FabricOUIdentifier) Reset()         { *m = FabricOUIdentifier{} }
func (m *FabricOUIdentifier) String() string { return proto.CompactTextString(m) }
func (*FabricOUIdentifier) ProtoMessage()    {}
func (*FabricOUIdentifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_msp_config_c4b502fff5780348, []int{7}
}
func (m *FabricOUIdentifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FabricOUIdentifier.Unmarshal(m, b)
}
func (m *FabricOUIdentifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FabricOUIdentifier.Marshal(b, m, deterministic)
}
func (dst *FabricOUIdentifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FabricOUIdentifier.Merge(dst, src)
}
func (m *FabricOUIdentifier) XXX_Size() int {
	return xxx_messageInfo_FabricOUIdentifier.Size(m)
}
func (m *FabricOUIdentifier) XXX_DiscardUnknown() {
	xxx_messageInfo_FabricOUIdentifier.DiscardUnknown(m)
}

var xxx_messageInfo_FabricOUIdentifier proto.InternalMessageInfo

func (m *FabricOUIdentifier) GetCertificate() []byte {
	if m != nil {
		return m.Certificate
	}
	return nil
}

func (m *FabricOUIdentifier) GetOrganizationalUnitIdentifier() string {
	if m != nil {
		return m.OrganizationalUnitIdentifier
	}
	return ""
}

// FabricNodeOUs contains configuration to tell apart clients from peers from orderers
// based on OUs. If NodeOUs recognition is enabled then an msp identity
// that does not contain any of the specified OU will be considered invalid.
type FabricNodeOUs struct {
	// If true then an msp identity that does not contain any of the specified OU will be considered invalid.
	Enable bool `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	// OU Identifier of the clients
	ClientOuIdentifier *FabricOUIdentifier `protobuf:"bytes,2,opt,name=client_ou_identifier,json=clientOuIdentifier,proto3" json:"client_ou_identifier,omitempty"`
	// OU Identifier of the peers
	PeerOuIdentifier *FabricOUIdentifier `protobuf:"bytes,3,opt,name=peer_ou_identifier,json=peerOuIdentifier,proto3" json:"peer_ou_identifier,omitempty"`
	// OU Identifier of the admins
	AdminOuIdentifier *FabricOUIdentifier `protobuf:"bytes,4,opt,name=admin_ou_identifier,json=adminOuIdentifier,proto3" json:"admin_ou_identifier,omitempty"`
	// OU Identifier of the orderers
	OrdererOuIdentifier  *FabricOUIdentifier `protobuf:"bytes,5,opt,name=orderer_ou_identifier,json=ordererOuIdentifier,proto3" json:"orderer_ou_identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *FabricNodeOUs) Reset()         { *m = FabricNodeOUs{} }
func (m *FabricNodeOUs) String() string { return proto.CompactTextString(m) }
func (*FabricNodeOUs) ProtoMessage()    {}
func (*FabricNodeOUs) Descriptor() ([]byte, []int) {
	return fileDescriptor_msp_config_c4b502fff5780348, []int{8}
}
func (m *FabricNodeOUs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FabricNodeOUs.Unmarshal(m, b)
}
func (m *FabricNodeOUs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FabricNodeOUs.Marshal(b, m, deterministic)
}
func (dst *FabricNodeOUs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FabricNodeOUs.Merge(dst, src)
}
func (m *FabricNodeOUs) XXX_Size() int {
	return xxx_messageInfo_FabricNodeOUs.Size(m)
}
func (m *FabricNodeOUs) XXX_DiscardUnknown() {
	xxx_messageInfo_FabricNodeOUs.DiscardUnknown(m)
}

var xxx_messageInfo_FabricNodeOUs proto.InternalMessageInfo

func (m *FabricNodeOUs) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *FabricNodeOUs) GetClientOuIdentifier() *FabricOUIdentifier {
	if m != nil {
		return m.ClientOuIdentifier
	}
	return nil
}

func (m *FabricNodeOUs) GetPeerOuIdentifier() *FabricOUIdentifier {
	if m != nil {
		return m.PeerOuIdentifier
	}
	return nil
}

func (m *FabricNodeOUs) GetAdminOuIdentifier() *FabricOUIdentifier {
	if m != nil {
		return m.AdminOuIdentifier
	}
	return nil
}

func (m *FabricNodeOUs) GetOrdererOuIdentifier() *FabricOUIdentifier {
	if m != nil {
		return m.OrdererOuIdentifier
	}
	return nil
}

func init() {
	proto.RegisterType((*MSPConfig)(nil), "msp.MSPConfig")
	proto.RegisterType((*FabricMSPConfig)(nil), "msp.FabricMSPConfig")
	proto.RegisterType((*FabricCryptoConfig)(nil), "msp.FabricCryptoConfig")
	proto.RegisterType((*IdemixMSPConfig)(nil), "msp.IdemixMSPConfig")
	proto.RegisterType((*IdemixMSPSignerConfig)(nil), "msp.IdemixMSPSignerConfig")
	proto.RegisterType((*SigningIdentityInfo)(nil), "msp.SigningIdentityInfo")
	proto.RegisterType((*KeyInfo)(nil), "msp.KeyInfo")
	proto.RegisterType((*FabricOUIdentifier)(nil), "msp.FabricOUIdentifier")
	proto.RegisterType((*FabricNodeOUs)(nil), "msp.FabricNodeOUs")
}

func init() { proto.RegisterFile("msp/msp_config.proto", fileDescriptor_msp_config_c4b502fff5780348) }

var fileDescriptor_msp_config_c4b502fff5780348 = []byte{
	// 878 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xcd, 0x6e, 0x23, 0x45,
	0x10, 0x96, 0x7f, 0x77, 0x5d, 0x1e, 0xdb, 0xd9, 0xce, 0x0f, 0x23, 0xc4, 0xee, 0x3a, 0x03, 0x08,
	0x5f, 0x70, 0xa4, 0x2c, 0x12, 0x12, 0xe2, 0xb4, 0x81, 0x85, 0x61, 0x09, 0x89, 0x3a, 0xca, 0x85,
	0xcb, 0xa8, 0x3d, 0xd3, 0xb6, 0x5b, 0x9e, 0xe9, 0x1e, 0x75, 0xf7, 0x44, 0x18, 0x71, 0xe6, 0x05,
	0x78, 0x07, 0xce, 0xbc, 0x09, 0xaf, 0x84, 0xfa, 0x27, 0xf6, 0x38, 0x89, 0xcc, 0xde, 0xaa, 0xab,
	0xbe, 0xfa, 0xa6, 0xfa, 0xab, 0xaa, 0x1e, 0x38, 0x2a, 0x54, 0x79, 0x56, 0xa8, 0x32, 0x49, 0x05,
	0x9f, 0xb3, 0xc5, 0xb4, 0x94, 0x42, 0x0b, 0xd4, 0x2a, 0x54, 0x19, 0x7d, 0x0d, 0xbd, 0xcb, 0x9b,
	0xeb, 0x0b, 0xeb, 0x47, 0x08, 0xda, 0x7a, 0x5d, 0xd2, 0xb0, 0x31, 0x6e, 0x4c, 0x3a, 0xd8, 0xda,
	0xe8, 0x04, 0xba, 0x2e, 0x2b, 0x6c, 0x8e, 0x1b, 0x93, 0x00, 0xfb, 0x53, 0xf4, 0x4f, 0x1b, 0x46,
	0xef, 0xc8, 0x4c, 0xb2, 0x74, 0x27, 0x9f, 0x93, 0xc2, 0xe5, 0xf7, 0xb0, 0xb5, 0xd1, 0x4b, 0x00,
	0x29, 0x84, 0x4e, 0x52, 0x2a, 0xb5, 0x0a, 0x9b, 0xe3, 0xd6, 0x24, 0xc0, 0x3d, 0xe3, 0xb9, 0x30,
	0x0e, 0xf4, 0x25, 0x20, 0xc6, 0x35, 0x95, 0x05, 0xcd, 0x18, 0xd1, 0xd4, 0xc3, 0x5a, 0x16, 0xf6,
	0xa2, 0x1e, 0x71, 0xf0, 0x13, 0xe8, 0x92, 0xac, 0x60, 0x5c, 0x85, 0x6d, 0x0b, 0xf1, 0x27, 0xf4,
	0x05, 0x8c, 0x24, 0xbd, 0x13, 0x29, 0xd1, 0x4c, 0xf0, 0x24, 0x67, 0x4a, 0x87, 0x1d, 0x0b, 0x18,
	0x6e, 0xdd, 0x3f, 0x33, 0xa5, 0xd1, 0x05, 0x1c, 0x28, 0xb6, 0xe0, 0x8c, 0x2f, 0x12, 0x96, 0x51,
	0xae, 0x99, 0x5e, 0x87, 0xdd, 0x71, 0x63, 0xd2, 0x3f, 0x0f, 0xa7, 0x85, 0x2a, 0xa7, 0x37, 0x2e,
	0x18, 0xfb, 0x58, 0xcc, 0xe7, 0x02, 0x8f, 0xd4, 0xae, 0x13, 0x25, 0xf0, 0x5a, 0xc8, 0x05, 0xe1,
	0xec, 0x77, 0x4b, 0x4c, 0xf2, 0xa4, 0xe2, 0x4c, 0x7b, 0xc2, 0x39, 0xa3, 0x52, 0x85, 0xcf, 0xc6,
	0xad, 0x49, 0xff, 0xfc, 0x23, 0xcb, 0xe9, 0x64, 0xba, 0xba, 0x8d, 0x37, 0x71, 0xfc, 0x72, 0x37,
	0xff, 0x96, 0x33, 0xbd, 0x8d, 0x2a, 0xf4, 0x2d, 0x0c, 0x52, 0xb9, 0x2e, 0xb5, 0xf0, 0x1d, 0x0b,
	0x9f, 0xdb, 0x12, 0xeb, 0x74, 0x17, 0x36, 0xee, 0x84, 0xc7, 0x41, 0x5a, 0x3b, 0xa1, 0xcf, 0x60,
	0xa8, 0x73, 0x95, 0xd4, 0x64, 0xef, 0x59, 0x2d, 0x02, 0x9d, 0x2b, 0xbc, 0x51, 0xfe, 0x2b, 0x38,
	0x31, 0xa8, 0x27, 0xd4, 0x07, 0x8b, 0x3e, 0xd2, 0xb9, 0x8a, 0x1f, 0x35, 0xe0, 0x1b, 0x18, 0xcd,
	0xed, 0xf7, 0x13, 0x2e, 0x32, 0x9a, 0x88, 0x4a, 0x85, 0x7d, 0x5b, 0x1b, 0xaa, 0xd5, 0xf6, 0x8b,
	0xc8, 0xe8, 0xd5, 0xad, 0xc2, 0x83, 0xf9, 0xf6, 0x58, 0xa9, 0xe8, 0xaf, 0x06, 0xa0, 0xc7, 0xc5,
	0xa3, 0x73, 0x38, 0x36, 0x02, 0x13, 0x5d, 0x49, 0x9a, 0x2c, 0x89, 0x5a, 0x26, 0x73, 0x52, 0xb0,
	0x7c, 0xed, 0xc7, 0xe8, 0x70, 0x13, 0xfc, 0x91, 0xa8, 0xe5, 0x3b, 0x1b, 0x42, 0x31, 0x9c, 0xde,
	0xb7, 0xaf, 0x26, 0xbb, 0xcf, 0xae, 0x78, 0x6a, 0x64, 0xb5, 0x03, 0xdb, 0xc3, 0xaf, 0xee, 0x81,
	0x5b, 0x81, 0x2d, 0x91, 0x47, 0x45, 0x7f, 0x37, 0x60, 0x14, 0x67, 0xb4, 0x60, 0xbf, 0xed, 0x1f,
	0xe4, 0x03, 0x68, 0xb1, 0x72, 0xe5, 0xb7, 0xc0, 0x98, 0xe8, 0x1c, 0xba, 0xa6, 0x36, 0x2a, 0xc3,
	0x96, 0x95, 0xe0, 0x63, 0x2b, 0xc1, 0x86, 0xeb, 0xc6, 0xc6, 0x7c, 0x87, 0x3c, 0x12, 0x7d, 0x0a,
	0x83, 0xda, 0xa0, 0x96, 0xab, 0xb0, 0x6d, 0xf9, 0x82, 0xad, 0xf3, 0x7a, 0x85, 0x8e, 0xa0, 0x43,
	0x4b, 0x91, 0x2e, 0xc3, 0xce, 0xb8, 0x31, 0x69, 0x61, 0x77, 0x88, 0xfe, 0x6c, 0xc2, 0xf1, 0x93,
	0xe4, 0xa6, 0xdc, 0x54, 0xd2, 0xcc, 0x96, 0x1b, 0x60, 0x6b, 0xa3, 0x21, 0x34, 0xd5, 0x7d, 0xb5,
	0x4d, 0xb5, 0x42, 0xdf, 0xc1, 0xab, 0xfd, 0x33, 0x6b, 0x2f, 0xd1, 0xc3, 0x9f, 0xec, 0x9b, 0x4c,
	0xf3, 0x25, 0x29, 0x72, 0x6a, 0xab, 0xee, 0x60, 0x6b, 0x9b, 0x2b, 0x51, 0x2e, 0x45, 0x9e, 0x17,
	0x94, 0x1b, 0x42, 0x5b, 0x75, 0x0f, 0x07, 0x5b, 0x67, 0x9c, 0xa1, 0x9f, 0xe0, 0xd4, 0x94, 0x65,
	0x88, 0x48, 0x9e, 0xd4, 0x24, 0x60, 0x7c, 0x2e, 0x64, 0x61, 0x6d, 0xbb, 0x88, 0x01, 0x7e, 0xbd,
	0x05, 0xe2, 0x0d, 0x2e, 0xde, 0xc2, 0x22, 0x01, 0x87, 0x4f, 0xac, 0xa9, 0xa9, 0xa3, 0xac, 0x66,
	0x39, 0x4b, 0x13, 0xdf, 0x15, 0x27, 0x47, 0xe0, 0x9c, 0x4e, 0x30, 0xf4, 0x06, 0x86, 0xa5, 0x64,
	0x77, 0x66, 0xd8, 0x3d, 0xaa, 0x69, 0x7b, 0x17, 0xd8, 0xde, 0xbd, 0xa7, 0x6e, 0xe3, 0x07, 0x1e,
	0xe3, 0x92, 0xa2, 0x1b, 0x78, 0xe6, 0x23, 0xe8, 0x73, 0x18, 0xae, 0x68, 0x7d, 0xe6, 0xfc, 0x8c,
	0x0c, 0x56, 0xb4, 0x36, 0x60, 0xe8, 0x14, 0x02, 0x03, 0x2b, 0x88, 0xa6, 0x92, 0x91, 0xdc, 0xf7,
	0xa1, 0xbf, 0xa2, 0xeb, 0x4b, 0xef, 0x8a, 0xfe, 0xb8, 0x5f, 0x86, 0xfa, 0xc3, 0x80, 0xc6, 0xd0,
	0x37, 0x4b, 0xc8, 0xe6, 0x2c, 0x25, 0x9a, 0xfa, 0x2b, 0xd4, 0x5d, 0x1f, 0xd0, 0xc8, 0xe6, 0xff,
	0x37, 0x32, 0xfa, 0xb7, 0x09, 0x83, 0x9d, 0x65, 0x35, 0x4f, 0x2b, 0xe5, 0x64, 0x96, 0xbb, 0x8f,
	0x3e, 0xc7, 0xfe, 0x84, 0x62, 0x38, 0x4a, 0x73, 0x66, 0x5a, 0x2b, 0xaa, 0x87, 0x5f, 0xd9, 0xf3,
	0xc2, 0x21, 0x97, 0x74, 0x55, 0xd5, 0x2e, 0xf7, 0x3d, 0xa0, 0x92, 0x52, 0xf9, 0x80, 0xa8, 0xb5,
	0x9f, 0xe8, 0xc0, 0xa4, 0xec, 0xd0, 0xfc, 0x00, 0x87, 0xf6, 0xd9, 0x7f, 0xc0, 0xd3, 0xde, 0xcf,
	0xf3, 0xc2, 0xe6, 0xec, 0x10, 0xbd, 0x87, 0x63, 0x21, 0x33, 0x2a, 0x1f, 0x95, 0xd4, 0xd9, 0x4f,
	0x75, 0xe8, 0xb3, 0xea, 0x64, 0x6f, 0x13, 0x38, 0x15, 0x72, 0x31, 0x5d, 0xae, 0x4b, 0x2a, 0x73,
	0x9a, 0x2d, 0xa8, 0x9c, 0xba, 0xe7, 0xcf, 0xfd, 0x6e, 0x95, 0x21, 0x7b, 0x7b, 0x70, 0xa9, 0x4a,
	0xb7, 0xb4, 0xd7, 0x24, 0x5d, 0x91, 0x05, 0xfd, 0x75, 0xb2, 0x60, 0x7a, 0x59, 0xcd, 0xa6, 0xa9,
	0x28, 0xce, 0x6a, 0xb9, 0x67, 0x2e, 0xf7, 0xcc, 0xe5, 0x9a, 0x9f, 0xf7, 0xac, 0x6b, 0xed, 0x37,
	0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x69, 0x49, 0xd1, 0x51, 0xce, 0x07, 0x00, 0x00,
}
