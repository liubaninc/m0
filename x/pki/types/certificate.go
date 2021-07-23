package types

func (cert *Certificate) Index() string {
	return cert.Issuer + "/" + cert.SerialNumber
}

func (cert *CertificateIdentifier) Index() string {
	return cert.Subject + "/" + cert.SubjectKeyID
}

func NewRootCertificate(pemCert string, subject string, subjectKeyID string,
	serialNumber string, creator string) Certificate {
	return Certificate{
		PemCert:      pemCert,
		Subject:      subject,
		SubjectKeyID: subjectKeyID,
		SerialNumber: serialNumber,
		IsRoot:       true,
		Creator:      creator,
	}
}

func NewNonRootCertificate(pemCert string, subject string, subjectKeyID string, serialNumber string,
	issuer string, authorityKeyID string,
	rootSubject string, rootSubjectKeyID string,
	creator string) Certificate {
	return Certificate{
		PemCert:          pemCert,
		Subject:          subject,
		SubjectKeyID:     subjectKeyID,
		SerialNumber:     serialNumber,
		Issuer:           issuer,
		AuthorityKeyID:   authorityKeyID,
		RootSubject:      rootSubject,
		RootSubjectKeyID: rootSubjectKeyID,
		IsRoot:           false,
		Creator:          creator,
	}
}